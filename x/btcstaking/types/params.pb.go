// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btcstaking/v1/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	github_com_babylonlabs_io_babylon_types "github.com/babylonlabs-io/babylon/types"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the module.
type Params struct {
	// PARAMETERS COVERING STAKING
	// covenant_pks is the list of public keys held by the covenant committee
	// each PK follows encoding in BIP-340 spec on Bitcoin
	CovenantPks []github_com_babylonlabs_io_babylon_types.BIP340PubKey `protobuf:"bytes,1,rep,name=covenant_pks,json=covenantPks,proto3,customtype=github.com/babylonlabs-io/babylon/types.BIP340PubKey" json:"covenant_pks,omitempty"`
	// covenant_quorum is the minimum number of signatures needed for the covenant
	// multisignature
	CovenantQuorum uint32 `protobuf:"varint,2,opt,name=covenant_quorum,json=covenantQuorum,proto3" json:"covenant_quorum,omitempty"`
	// min_staking_value_sat is the minimum of satoshis locked in staking output
	MinStakingValueSat int64 `protobuf:"varint,3,opt,name=min_staking_value_sat,json=minStakingValueSat,proto3" json:"min_staking_value_sat,omitempty"`
	// max_staking_value_sat is the maximum of satoshis locked in staking output
	MaxStakingValueSat int64 `protobuf:"varint,4,opt,name=max_staking_value_sat,json=maxStakingValueSat,proto3" json:"max_staking_value_sat,omitempty"`
	// min_staking_time is the minimum lock time specified in staking output script
	MinStakingTimeBlocks uint32 `protobuf:"varint,5,opt,name=min_staking_time_blocks,json=minStakingTimeBlocks,proto3" json:"min_staking_time_blocks,omitempty"`
	// max_staking_time_blocks is the maximum lock time time specified in staking output script
	MaxStakingTimeBlocks uint32 `protobuf:"varint,6,opt,name=max_staking_time_blocks,json=maxStakingTimeBlocks,proto3" json:"max_staking_time_blocks,omitempty"`
	// PARAMETERS COVERING SLASHING
	// slashing_pk_script is the pk_script expected in slashing output ie. the first
	// output of slashing transaction
	SlashingPkScript []byte `protobuf:"bytes,7,opt,name=slashing_pk_script,json=slashingPkScript,proto3" json:"slashing_pk_script,omitempty"`
	// min_slashing_tx_fee_sat is the minimum amount of tx fee (quantified
	// in Satoshi) needed for the pre-signed slashing tx. It covers both:
	// staking slashing transaction and unbonding slashing transaction
	MinSlashingTxFeeSat int64 `protobuf:"varint,8,opt,name=min_slashing_tx_fee_sat,json=minSlashingTxFeeSat,proto3" json:"min_slashing_tx_fee_sat,omitempty"`
	// slashing_rate determines the portion of the staked amount to be slashed,
	// expressed as a decimal (e.g., 0.5 for 50%). Maximal precion is 2 decimal
	// places
	SlashingRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,9,opt,name=slashing_rate,json=slashingRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"slashing_rate"`
	// PARAMETERS COVERING UNBONDING
	// min_unbonding_time is the minimum time for unbonding transaction timelock in BTC blocks
	MinUnbondingTimeBlocks uint32 `protobuf:"varint,10,opt,name=min_unbonding_time_blocks,json=minUnbondingTimeBlocks,proto3" json:"min_unbonding_time_blocks,omitempty"`
	// unbonding_fee exact fee required for unbonding transaction
	UnbondingFeeSat int64 `protobuf:"varint,11,opt,name=unbonding_fee_sat,json=unbondingFeeSat,proto3" json:"unbonding_fee_sat,omitempty"`
	// PARAMETERS COVERING FINALITY PROVIDERS
	// min_commission_rate is the chain-wide minimum commission rate that a finality provider
	// can charge their delegators expressed as a decimal (e.g., 0.5 for 50%). Maximal precion
	// is 2 decimal places
	MinCommissionRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,12,opt,name=min_commission_rate,json=minCommissionRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"min_commission_rate"`
	// base gas fee for delegation creation
	DelegationCreationBaseGasFee uint64 `protobuf:"varint,13,opt,name=delegation_creation_base_gas_fee,json=delegationCreationBaseGasFee,proto3" json:"delegation_creation_base_gas_fee,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1392776a3e15b9, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetCovenantQuorum() uint32 {
	if m != nil {
		return m.CovenantQuorum
	}
	return 0
}

func (m *Params) GetMinStakingValueSat() int64 {
	if m != nil {
		return m.MinStakingValueSat
	}
	return 0
}

func (m *Params) GetMaxStakingValueSat() int64 {
	if m != nil {
		return m.MaxStakingValueSat
	}
	return 0
}

func (m *Params) GetMinStakingTimeBlocks() uint32 {
	if m != nil {
		return m.MinStakingTimeBlocks
	}
	return 0
}

func (m *Params) GetMaxStakingTimeBlocks() uint32 {
	if m != nil {
		return m.MaxStakingTimeBlocks
	}
	return 0
}

func (m *Params) GetSlashingPkScript() []byte {
	if m != nil {
		return m.SlashingPkScript
	}
	return nil
}

func (m *Params) GetMinSlashingTxFeeSat() int64 {
	if m != nil {
		return m.MinSlashingTxFeeSat
	}
	return 0
}

func (m *Params) GetMinUnbondingTimeBlocks() uint32 {
	if m != nil {
		return m.MinUnbondingTimeBlocks
	}
	return 0
}

func (m *Params) GetUnbondingFeeSat() int64 {
	if m != nil {
		return m.UnbondingFeeSat
	}
	return 0
}

func (m *Params) GetDelegationCreationBaseGasFee() uint64 {
	if m != nil {
		return m.DelegationCreationBaseGasFee
	}
	return 0
}

// StoredParams attach information about the version of stored parameters
type StoredParams struct {
	// version of the stored parameters. Each parameters update
	// increments version number by 1
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// NOTE: Parameters must always be provided
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *StoredParams) Reset()         { *m = StoredParams{} }
func (m *StoredParams) String() string { return proto.CompactTextString(m) }
func (*StoredParams) ProtoMessage()    {}
func (*StoredParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1392776a3e15b9, []int{1}
}
func (m *StoredParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredParams.Merge(m, src)
}
func (m *StoredParams) XXX_Size() int {
	return m.Size()
}
func (m *StoredParams) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredParams.DiscardUnknown(m)
}

var xxx_messageInfo_StoredParams proto.InternalMessageInfo

func (m *StoredParams) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *StoredParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*Params)(nil), "babylon.btcstaking.v1.Params")
	proto.RegisterType((*StoredParams)(nil), "babylon.btcstaking.v1.StoredParams")
}

func init() {
	proto.RegisterFile("babylon/btcstaking/v1/params.proto", fileDescriptor_8d1392776a3e15b9)
}

var fileDescriptor_8d1392776a3e15b9 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x63, 0x1a, 0x52, 0x3a, 0x4d, 0x29, 0x35, 0x2d, 0xb8, 0x05, 0x1c, 0xab, 0x2c, 0xb0,
	0x10, 0xb5, 0x09, 0x2d, 0x12, 0x3f, 0x3b, 0xb7, 0x2a, 0x42, 0x20, 0x14, 0x9c, 0xd2, 0x05, 0x2c,
	0xac, 0xb1, 0x33, 0xb8, 0x23, 0x7b, 0x3c, 0xc6, 0x33, 0x8e, 0x92, 0xb7, 0x60, 0xc9, 0x92, 0x25,
	0x0f, 0xc0, 0x43, 0x74, 0x59, 0xb1, 0x42, 0x5d, 0x54, 0xa8, 0x79, 0x11, 0xe4, 0xf1, 0x38, 0xb1,
	0xaa, 0x2e, 0xba, 0xf3, 0xcc, 0xb9, 0xe7, 0xde, 0xf3, 0x25, 0xbe, 0x06, 0x9b, 0x3e, 0xf4, 0xc7,
	0x31, 0x4d, 0x6c, 0x9f, 0x07, 0x8c, 0xc3, 0x08, 0x27, 0xa1, 0x3d, 0xec, 0xda, 0x29, 0xcc, 0x20,
	0x61, 0x56, 0x9a, 0x51, 0x4e, 0xd5, 0x35, 0x59, 0x63, 0xcd, 0x6a, 0xac, 0x61, 0x77, 0x63, 0x35,
	0xa4, 0x21, 0x15, 0x15, 0x76, 0xf1, 0x54, 0x16, 0x6f, 0xac, 0x07, 0x94, 0x11, 0xca, 0xbc, 0x52,
	0x28, 0x0f, 0xa5, 0xb4, 0xf9, 0xab, 0x05, 0x5a, 0x3d, 0xd1, 0x58, 0xfd, 0x02, 0xda, 0x01, 0x1d,
	0xa2, 0x04, 0x26, 0xdc, 0x4b, 0x23, 0xa6, 0x29, 0xc6, 0x9c, 0xd9, 0x76, 0x5e, 0x9c, 0x9e, 0x75,
	0x76, 0x42, 0xcc, 0x8f, 0x72, 0xdf, 0x0a, 0x28, 0xb1, 0xe5, 0xdc, 0x18, 0xfa, 0x6c, 0x0b, 0xd3,
	0xea, 0x68, 0xf3, 0x71, 0x8a, 0x98, 0xe5, 0xbc, 0xed, 0x6d, 0xef, 0x3c, 0xed, 0xe5, 0xfe, 0x3b,
	0x34, 0x76, 0x17, 0xab, 0x6e, 0xbd, 0x88, 0xa9, 0x8f, 0xc0, 0xf2, 0xb4, 0xf9, 0xb7, 0x9c, 0x66,
	0x39, 0xd1, 0xae, 0x19, 0x8a, 0xb9, 0xe4, 0xde, 0xac, 0xae, 0x3f, 0x8a, 0x5b, 0xb5, 0x0b, 0xd6,
	0x08, 0x4e, 0x3c, 0xc9, 0xe4, 0x0d, 0x61, 0x9c, 0x23, 0x8f, 0x41, 0xae, 0xcd, 0x19, 0x8a, 0x39,
	0xe7, 0xaa, 0x04, 0x27, 0xfd, 0x52, 0x3b, 0x2c, 0xa4, 0x3e, 0xe4, 0xc2, 0x02, 0x47, 0x97, 0x58,
	0x9a, 0xd2, 0x02, 0x47, 0x17, 0x2d, 0xcf, 0xc1, 0xdd, 0xfa, 0x14, 0x8e, 0x09, 0xf2, 0xfc, 0x98,
	0x06, 0x11, 0xd3, 0xae, 0x8b, 0x58, 0xab, 0xb3, 0x39, 0x07, 0x98, 0x20, 0x47, 0x68, 0xc2, 0x56,
	0x9b, 0x54, 0xb7, 0xb5, 0xa4, 0x6d, 0x3a, 0xab, 0x66, 0x7b, 0x02, 0x54, 0x16, 0x43, 0x76, 0x54,
	0x78, 0xd2, 0xc8, 0x63, 0x41, 0x86, 0x53, 0xae, 0xcd, 0x1b, 0x8a, 0xd9, 0x76, 0x6f, 0x55, 0x4a,
	0x2f, 0xea, 0x8b, 0x7b, 0x75, 0x47, 0x66, 0xab, 0x1c, 0x7c, 0xe4, 0x7d, 0x45, 0x25, 0xd0, 0x0d,
	0x01, 0x74, 0xbb, 0xc8, 0x26, 0xd5, 0x83, 0xd1, 0x3e, 0x12, 0x44, 0x87, 0x60, 0x69, 0xea, 0xc8,
	0x20, 0x47, 0xda, 0x82, 0xa1, 0x98, 0x0b, 0x4e, 0xf7, 0xf8, 0xac, 0xd3, 0x38, 0x3d, 0xeb, 0xdc,
	0x2b, 0xff, 0x75, 0x36, 0x88, 0x2c, 0x4c, 0x6d, 0x02, 0xf9, 0x91, 0xf5, 0x1e, 0x85, 0x30, 0x18,
	0xef, 0xa1, 0xe0, 0xcf, 0xef, 0x2d, 0x20, 0x5f, 0x8a, 0x3d, 0x14, 0xb8, 0xed, 0xaa, 0x8f, 0x0b,
	0x39, 0x52, 0x5f, 0x82, 0xf5, 0x22, 0x4d, 0x9e, 0xf8, 0x34, 0x19, 0x5c, 0x84, 0x06, 0x02, 0xfa,
	0x0e, 0xc1, 0xc9, 0xa7, 0x4a, 0xaf, 0x61, 0x3f, 0x06, 0x2b, 0x33, 0x5b, 0x85, 0xb0, 0x28, 0x10,
	0x96, 0xa7, 0x82, 0x8c, 0xdf, 0x07, 0x05, 0x95, 0x17, 0x50, 0x42, 0x30, 0x63, 0x98, 0x26, 0x25,
	0x44, 0x5b, 0x40, 0x3c, 0xbc, 0x02, 0x84, 0xbb, 0x42, 0x70, 0xb2, 0x3b, 0xb5, 0x8b, 0xec, 0xfb,
	0xc0, 0x18, 0xa0, 0x18, 0x85, 0x90, 0x17, 0x0d, 0x83, 0x0c, 0x95, 0x0f, 0x3e, 0x64, 0xc8, 0x0b,
	0x21, 0x2b, 0x32, 0x69, 0x4b, 0x86, 0x62, 0x36, 0xdd, 0xfb, 0xb3, 0xba, 0x5d, 0x59, 0xe6, 0x40,
	0x86, 0xde, 0x40, 0xb6, 0x8f, 0xd0, 0xab, 0xe6, 0x8f, 0x9f, 0x9d, 0xc6, 0x26, 0x02, 0xed, 0x3e,
	0xa7, 0x19, 0x1a, 0xc8, 0x7d, 0xd1, 0xc0, 0xfc, 0x10, 0x65, 0xc5, 0x30, 0x4d, 0x11, 0xbf, 0x43,
	0x75, 0x54, 0x5f, 0x83, 0x56, 0xb9, 0xac, 0xe2, 0x1d, 0x5f, 0x7c, 0xf6, 0xc0, 0xba, 0x74, 0x5b,
	0xad, 0xb2, 0x91, 0xd3, 0x2c, 0xf0, 0x5c, 0x69, 0x71, 0x3e, 0x1c, 0x9f, 0xeb, 0xca, 0xc9, 0xb9,
	0xae, 0xfc, 0x3b, 0xd7, 0x95, 0xef, 0x13, 0xbd, 0x71, 0x32, 0xd1, 0x1b, 0x7f, 0x27, 0x7a, 0xe3,
	0xf3, 0x15, 0xd6, 0x70, 0x54, 0xff, 0x66, 0x88, 0x9d, 0xf4, 0x5b, 0x62, 0xd1, 0xb7, 0xff, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x80, 0xd8, 0x41, 0xc8, 0x56, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DelegationCreationBaseGasFee != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DelegationCreationBaseGasFee))
		i--
		dAtA[i] = 0x68
	}
	{
		size := m.MinCommissionRate.Size()
		i -= size
		if _, err := m.MinCommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if m.UnbondingFeeSat != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UnbondingFeeSat))
		i--
		dAtA[i] = 0x58
	}
	if m.MinUnbondingTimeBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinUnbondingTimeBlocks))
		i--
		dAtA[i] = 0x50
	}
	{
		size := m.SlashingRate.Size()
		i -= size
		if _, err := m.SlashingRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.MinSlashingTxFeeSat != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinSlashingTxFeeSat))
		i--
		dAtA[i] = 0x40
	}
	if len(m.SlashingPkScript) > 0 {
		i -= len(m.SlashingPkScript)
		copy(dAtA[i:], m.SlashingPkScript)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SlashingPkScript)))
		i--
		dAtA[i] = 0x3a
	}
	if m.MaxStakingTimeBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxStakingTimeBlocks))
		i--
		dAtA[i] = 0x30
	}
	if m.MinStakingTimeBlocks != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinStakingTimeBlocks))
		i--
		dAtA[i] = 0x28
	}
	if m.MaxStakingValueSat != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxStakingValueSat))
		i--
		dAtA[i] = 0x20
	}
	if m.MinStakingValueSat != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinStakingValueSat))
		i--
		dAtA[i] = 0x18
	}
	if m.CovenantQuorum != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CovenantQuorum))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CovenantPks) > 0 {
		for iNdEx := len(m.CovenantPks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.CovenantPks[iNdEx].Size()
				i -= size
				if _, err := m.CovenantPks[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *StoredParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Version != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CovenantPks) > 0 {
		for _, e := range m.CovenantPks {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.CovenantQuorum != 0 {
		n += 1 + sovParams(uint64(m.CovenantQuorum))
	}
	if m.MinStakingValueSat != 0 {
		n += 1 + sovParams(uint64(m.MinStakingValueSat))
	}
	if m.MaxStakingValueSat != 0 {
		n += 1 + sovParams(uint64(m.MaxStakingValueSat))
	}
	if m.MinStakingTimeBlocks != 0 {
		n += 1 + sovParams(uint64(m.MinStakingTimeBlocks))
	}
	if m.MaxStakingTimeBlocks != 0 {
		n += 1 + sovParams(uint64(m.MaxStakingTimeBlocks))
	}
	l = len(m.SlashingPkScript)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MinSlashingTxFeeSat != 0 {
		n += 1 + sovParams(uint64(m.MinSlashingTxFeeSat))
	}
	l = m.SlashingRate.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MinUnbondingTimeBlocks != 0 {
		n += 1 + sovParams(uint64(m.MinUnbondingTimeBlocks))
	}
	if m.UnbondingFeeSat != 0 {
		n += 1 + sovParams(uint64(m.UnbondingFeeSat))
	}
	l = m.MinCommissionRate.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.DelegationCreationBaseGasFee != 0 {
		n += 1 + sovParams(uint64(m.DelegationCreationBaseGasFee))
	}
	return n
}

func (m *StoredParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovParams(uint64(m.Version))
	}
	l = m.Params.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CovenantPks", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonlabs_io_babylon_types.BIP340PubKey
			m.CovenantPks = append(m.CovenantPks, v)
			if err := m.CovenantPks[len(m.CovenantPks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CovenantQuorum", wireType)
			}
			m.CovenantQuorum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CovenantQuorum |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinStakingValueSat", wireType)
			}
			m.MinStakingValueSat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinStakingValueSat |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxStakingValueSat", wireType)
			}
			m.MaxStakingValueSat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxStakingValueSat |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinStakingTimeBlocks", wireType)
			}
			m.MinStakingTimeBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinStakingTimeBlocks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxStakingTimeBlocks", wireType)
			}
			m.MaxStakingTimeBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxStakingTimeBlocks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashingPkScript", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SlashingPkScript = append(m.SlashingPkScript[:0], dAtA[iNdEx:postIndex]...)
			if m.SlashingPkScript == nil {
				m.SlashingPkScript = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSlashingTxFeeSat", wireType)
			}
			m.MinSlashingTxFeeSat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSlashingTxFeeSat |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashingRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashingRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUnbondingTimeBlocks", wireType)
			}
			m.MinUnbondingTimeBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinUnbondingTimeBlocks |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingFeeSat", wireType)
			}
			m.UnbondingFeeSat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingFeeSat |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCommissionRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinCommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationCreationBaseGasFee", wireType)
			}
			m.DelegationCreationBaseGasFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelegationCreationBaseGasFee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StoredParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoredParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
