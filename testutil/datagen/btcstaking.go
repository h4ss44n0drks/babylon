package datagen

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/babylonchain/babylon/btcstaking"
	bbn "github.com/babylonchain/babylon/types"
	bstypes "github.com/babylonchain/babylon/x/btcstaking/types"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func GenRandomBTCValidator(r *rand.Rand) (*bstypes.BTCValidator, error) {
	// key pairs
	btcSK, _, err := GenRandomBTCKeyPair(r)
	if err != nil {
		return nil, err
	}
	return GenRandomBTCValidatorWithBTCSK(r, btcSK)
}

func GenRandomBTCValidatorWithBTCSK(r *rand.Rand, btcSK *btcec.PrivateKey) (*bstypes.BTCValidator, error) {
	bbnSK, _, err := GenRandomSecp256k1KeyPair(r)
	if err != nil {
		return nil, err
	}
	return GenRandomBTCValidatorWithBTCBabylonSKs(r, btcSK, bbnSK)
}

func GenRandomBTCValidatorWithBTCBabylonSKs(r *rand.Rand, btcSK *btcec.PrivateKey, bbnSK cryptotypes.PrivKey) (*bstypes.BTCValidator, error) {
	// commission
	commission := sdk.NewDecWithPrec(int64(RandomInt(r, 49)+1), 2) // [1/100, 50/100]
	// description
	description := stakingtypes.Description{}
	// key pairs
	btcPK := btcSK.PubKey()
	bip340PK := bbn.NewBIP340PubKeyFromBTCPK(btcPK)
	bbnPK := bbnSK.PubKey()
	secp256k1PK, ok := bbnPK.(*secp256k1.PubKey)
	if !ok {
		return nil, fmt.Errorf("failed to assert bbnPK to *secp256k1.PubKey")
	}
	// pop
	pop, err := bstypes.NewPoP(bbnSK, btcSK)
	if err != nil {
		return nil, err
	}
	return &bstypes.BTCValidator{
		Description: &description,
		Commission:  &commission,
		BabylonPk:   secp256k1PK,
		BtcPk:       bip340PK,
		Pop:         pop,
	}, nil
}

func GenRandomBTCDelegation(r *rand.Rand, valBTCPK *bbn.BIP340PubKey, delSK *btcec.PrivateKey, jurySK *btcec.PrivateKey, slashingAddr string, startHeight uint64, endHeight uint64, totalSat uint64) (*bstypes.BTCDelegation, error) {
	net := &chaincfg.SimNetParams
	delPK := delSK.PubKey()
	delBTCPK := bbn.NewBIP340PubKeyFromBTCPK(delPK)
	juryBTCPK := jurySK.PubKey()
	valPK, err := valBTCPK.ToBTCPK()
	if err != nil {
		return nil, err
	}

	// BTC delegation Babylon key pairs
	bbnSK, bbnPK, err := GenRandomSecp256k1KeyPair(r)
	if err != nil {
		return nil, err
	}
	secp256k1PK, ok := bbnPK.(*secp256k1.PubKey)
	if !ok {
		return nil, fmt.Errorf("failed to assert bbnPK to *secp256k1.PubKey")
	}
	// pop
	pop, err := bstypes.NewPoP(bbnSK, delSK)
	if err != nil {
		return nil, err
	}
	// staking/slashing tx
	stakingTx, slashingTx, err := GenBTCStakingSlashingTx(r, net, delSK, valPK, juryBTCPK, uint16(endHeight-startHeight), int64(totalSat), slashingAddr)
	if err != nil {
		return nil, err
	}

	// jury sig and delegator sig
	stakingMsgTx, err := stakingTx.ToMsgTx()
	if err != nil {
		return nil, err
	}
	jurySig, err := slashingTx.Sign(stakingMsgTx, stakingTx.Script, jurySK, net)
	if err != nil {
		return nil, err
	}
	delegatorSig, err := slashingTx.Sign(stakingMsgTx, stakingTx.Script, delSK, net)
	if err != nil {
		return nil, err
	}

	return &bstypes.BTCDelegation{
		BabylonPk:    secp256k1PK,
		BtcPk:        delBTCPK,
		Pop:          pop,
		ValBtcPk:     valBTCPK,
		StartHeight:  startHeight,
		EndHeight:    endHeight,
		TotalSat:     totalSat,
		DelegatorSig: delegatorSig,
		JurySig:      jurySig,
		StakingTx:    stakingTx,
		SlashingTx:   slashingTx,
	}, nil
}

func GenBTCStakingSlashingTxWithOutPoint(
	r *rand.Rand,
	btcNet *chaincfg.Params,
	outPoint *wire.OutPoint,
	stakerSK *btcec.PrivateKey,
	validatorPK *btcec.PublicKey,
	juryPK *btcec.PublicKey,
	stakingTimeBlocks uint16,
	stakingValue int64,
	slashingAddress string,
	withChange bool,
) (*bstypes.BabylonBTCTaprootTx, *bstypes.BTCSlashingTx, error) {
	stakingOutput, stakingScript, err := btcstaking.BuildStakingOutput(
		stakerSK.PubKey(),
		validatorPK,
		juryPK,
		stakingTimeBlocks,
		btcutil.Amount(stakingValue),
		btcNet,
	)
	if err != nil {
		return nil, nil, err
	}

	tx := wire.NewMsgTx(2)
	// add the given tx input
	txIn := wire.NewTxIn(outPoint, nil, nil)
	tx.AddTxIn(txIn)

	tx.AddTxOut(stakingOutput)

	if withChange {
		// 2 outputs for changes and staking output
		changeAddrScript, err := GenRandomPubKeyHashScript(r, btcNet)
		if err != nil {
			return nil, nil, err
		}
		if txscript.GetScriptClass(changeAddrScript) == txscript.NonStandardTy {
			return nil, nil, fmt.Errorf("change address script is non-standard")
		}
		tx.AddTxOut(wire.NewTxOut(10000, changeAddrScript)) // output for change
	}

	// construct staking tx
	var buf bytes.Buffer
	err = tx.Serialize(&buf)
	if err != nil {
		return nil, nil, err
	}
	stakingTx := &bstypes.BabylonBTCTaprootTx{
		Tx:     buf.Bytes(),
		Script: stakingScript,
	}

	// construct slashing tx
	slashingAddrBtc, err := btcutil.DecodeAddress(slashingAddress, btcNet)
	if err != nil {
		return nil, nil, err
	}
	slashingMsgTx, err := btcstaking.BuildSlashingTxFromStakingTxStrict(tx, 0, slashingAddrBtc, 2000, stakingScript, btcNet)
	if err != nil {
		return nil, nil, err
	}
	slashingTx, err := bstypes.NewBTCSlashingTxFromMsgTx(slashingMsgTx)
	if err != nil {
		return nil, nil, err
	}

	return stakingTx, slashingTx, nil
}

func GenBTCStakingSlashingTx(
	r *rand.Rand,
	btcNet *chaincfg.Params,
	stakerSK *btcec.PrivateKey,
	validatorPK *btcec.PublicKey,
	juryPK *btcec.PublicKey,
	stakingTimeBlocks uint16,
	stakingValue int64,
	slashingAddress string,
) (*bstypes.BabylonBTCTaprootTx, *bstypes.BTCSlashingTx, error) {
	// an arbitrary input
	spend := makeSpendableOutWithRandOutPoint(r, btcutil.Amount(stakingValue+1000))
	outPoint := &spend.prevOut
	return GenBTCStakingSlashingTxWithOutPoint(r, btcNet, outPoint, stakerSK, validatorPK, juryPK, stakingTimeBlocks, stakingValue, slashingAddress, true)
}

func GenBTCUnbondingSlashingTx(
	r *rand.Rand,
	btcNet *chaincfg.Params,
	stakerSK *btcec.PrivateKey,
	validatorPK *btcec.PublicKey,
	juryPK *btcec.PublicKey,
	stakingTransactionOutpoint *wire.OutPoint,
	stakingTimeBlocks uint16,
	stakingValue int64,
	slashingAddress string,
) (*bstypes.BabylonBTCTaprootTx, *bstypes.BTCSlashingTx, error) {
	return GenBTCStakingSlashingTxWithOutPoint(r, btcNet, stakingTransactionOutpoint, stakerSK, validatorPK, juryPK, stakingTimeBlocks, stakingValue, slashingAddress, false)
}
