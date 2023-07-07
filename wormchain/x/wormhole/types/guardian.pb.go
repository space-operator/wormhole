// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wormhole/guardian.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type GuardianKey struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *GuardianKey) Reset()         { *m = GuardianKey{} }
func (m *GuardianKey) String() string { return proto.CompactTextString(m) }
func (*GuardianKey) ProtoMessage()    {}
func (*GuardianKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_95afcf26fc23dcb3, []int{0}
}
func (m *GuardianKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GuardianKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GuardianKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GuardianKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuardianKey.Merge(m, src)
}
func (m *GuardianKey) XXX_Size() int {
	return m.Size()
}
func (m *GuardianKey) XXX_DiscardUnknown() {
	xxx_messageInfo_GuardianKey.DiscardUnknown(m)
}

var xxx_messageInfo_GuardianKey proto.InternalMessageInfo

func (m *GuardianKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type GuardianValidator struct {
	GuardianKey   []byte `protobuf:"bytes,1,opt,name=guardianKey,proto3" json:"guardianKey,omitempty"`
	ValidatorAddr []byte `protobuf:"bytes,2,opt,name=validatorAddr,proto3" json:"validatorAddr,omitempty"`
}

func (m *GuardianValidator) Reset()         { *m = GuardianValidator{} }
func (m *GuardianValidator) String() string { return proto.CompactTextString(m) }
func (*GuardianValidator) ProtoMessage()    {}
func (*GuardianValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_95afcf26fc23dcb3, []int{1}
}
func (m *GuardianValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GuardianValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GuardianValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GuardianValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuardianValidator.Merge(m, src)
}
func (m *GuardianValidator) XXX_Size() int {
	return m.Size()
}
func (m *GuardianValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_GuardianValidator.DiscardUnknown(m)
}

var xxx_messageInfo_GuardianValidator proto.InternalMessageInfo

func (m *GuardianValidator) GetGuardianKey() []byte {
	if m != nil {
		return m.GuardianKey
	}
	return nil
}

func (m *GuardianValidator) GetValidatorAddr() []byte {
	if m != nil {
		return m.ValidatorAddr
	}
	return nil
}

type GuardianSet struct {
	Index          uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Keys           [][]byte `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	ExpirationTime uint64   `protobuf:"varint,3,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
}

func (m *GuardianSet) Reset()         { *m = GuardianSet{} }
func (m *GuardianSet) String() string { return proto.CompactTextString(m) }
func (*GuardianSet) ProtoMessage()    {}
func (*GuardianSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_95afcf26fc23dcb3, []int{2}
}
func (m *GuardianSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GuardianSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GuardianSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GuardianSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuardianSet.Merge(m, src)
}
func (m *GuardianSet) XXX_Size() int {
	return m.Size()
}
func (m *GuardianSet) XXX_DiscardUnknown() {
	xxx_messageInfo_GuardianSet.DiscardUnknown(m)
}

var xxx_messageInfo_GuardianSet proto.InternalMessageInfo

func (m *GuardianSet) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *GuardianSet) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *GuardianSet) GetExpirationTime() uint64 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

type ValidatorAllowedAddress struct {
	// the validator/guardian that controls this entry
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// the allowlisted account
	AllowedAddress string `protobuf:"bytes,2,opt,name=allowed_address,json=allowedAddress,proto3" json:"allowed_address,omitempty"`
	// human readable name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ValidatorAllowedAddress) Reset()         { *m = ValidatorAllowedAddress{} }
func (m *ValidatorAllowedAddress) String() string { return proto.CompactTextString(m) }
func (*ValidatorAllowedAddress) ProtoMessage()    {}
func (*ValidatorAllowedAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_95afcf26fc23dcb3, []int{3}
}
func (m *ValidatorAllowedAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorAllowedAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorAllowedAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorAllowedAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorAllowedAddress.Merge(m, src)
}
func (m *ValidatorAllowedAddress) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorAllowedAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorAllowedAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorAllowedAddress proto.InternalMessageInfo

func (m *ValidatorAllowedAddress) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *ValidatorAllowedAddress) GetAllowedAddress() string {
	if m != nil {
		return m.AllowedAddress
	}
	return ""
}

func (m *ValidatorAllowedAddress) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type WasmAllowedContractCodeId struct {
	// bech32 address of the contract that can call wasm instantiate without a VAA
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// reference to the stored WASM code that can be instantiated
	CodeId uint64 `protobuf:"varint,2,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
}

func (m *WasmAllowedContractCodeId) Reset()         { *m = WasmAllowedContractCodeId{} }
func (m *WasmAllowedContractCodeId) String() string { return proto.CompactTextString(m) }
func (*WasmAllowedContractCodeId) ProtoMessage()    {}
func (*WasmAllowedContractCodeId) Descriptor() ([]byte, []int) {
	return fileDescriptor_95afcf26fc23dcb3, []int{4}
}
func (m *WasmAllowedContractCodeId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WasmAllowedContractCodeId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WasmAllowedContractCodeId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WasmAllowedContractCodeId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmAllowedContractCodeId.Merge(m, src)
}
func (m *WasmAllowedContractCodeId) XXX_Size() int {
	return m.Size()
}
func (m *WasmAllowedContractCodeId) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmAllowedContractCodeId.DiscardUnknown(m)
}

var xxx_messageInfo_WasmAllowedContractCodeId proto.InternalMessageInfo

func (m *WasmAllowedContractCodeId) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *WasmAllowedContractCodeId) GetCodeId() uint64 {
	if m != nil {
		return m.CodeId
	}
	return 0
}

func init() {
	proto.RegisterType((*GuardianKey)(nil), "wormhole_foundation.wormchain.wormhole.GuardianKey")
	proto.RegisterType((*GuardianValidator)(nil), "wormhole_foundation.wormchain.wormhole.GuardianValidator")
	proto.RegisterType((*GuardianSet)(nil), "wormhole_foundation.wormchain.wormhole.GuardianSet")
	proto.RegisterType((*ValidatorAllowedAddress)(nil), "wormhole_foundation.wormchain.wormhole.ValidatorAllowedAddress")
	proto.RegisterType((*WasmAllowedContractCodeId)(nil), "wormhole_foundation.wormchain.wormhole.WasmAllowedContractCodeId")
}

func init() { proto.RegisterFile("wormhole/guardian.proto", fileDescriptor_95afcf26fc23dcb3) }

var fileDescriptor_95afcf26fc23dcb3 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0xcd, 0x0e, 0xd2, 0x40,
	0x18, 0x64, 0xa1, 0x62, 0xf8, 0xf8, 0xdf, 0x90, 0x50, 0x3d, 0x94, 0xa6, 0x31, 0x88, 0x31, 0xd2,
	0x83, 0x27, 0xbd, 0x21, 0x07, 0x63, 0xbc, 0x15, 0xa3, 0x89, 0x1e, 0x9a, 0xa5, 0xbb, 0x96, 0x0d,
	0x6d, 0x97, 0xb4, 0x8b, 0xd0, 0xb3, 0x2f, 0xe0, 0x23, 0xf8, 0x38, 0x1e, 0x39, 0x7a, 0x34, 0x70,
	0xf1, 0x31, 0x4c, 0xb7, 0x3f, 0x08, 0xb7, 0xd9, 0xf9, 0xe6, 0x9b, 0x99, 0x4d, 0x3e, 0x18, 0x1f,
	0x44, 0x1c, 0x6e, 0x44, 0xc0, 0x6c, 0x7f, 0x4f, 0x62, 0xca, 0x49, 0x34, 0xdf, 0xc5, 0x42, 0x0a,
	0x3c, 0x2d, 0x07, 0xee, 0x57, 0xb1, 0x8f, 0x28, 0x91, 0x5c, 0x44, 0xf3, 0x8c, 0xf3, 0x36, 0x84,
	0xe7, 0x28, 0x9b, 0x3e, 0x1e, 0xf9, 0xc2, 0x17, 0x6a, 0xc5, 0xce, 0x50, 0xbe, 0x6d, 0x4d, 0xa0,
	0xfd, 0xb6, 0xf0, 0x7b, 0xcf, 0x52, 0x3c, 0x80, 0xc6, 0x96, 0xa5, 0x3a, 0x32, 0xd1, 0xac, 0xe3,
	0x64, 0xd0, 0xfa, 0x02, 0xc3, 0x52, 0xf0, 0x91, 0x04, 0x9c, 0x12, 0x29, 0x62, 0x6c, 0x42, 0xdb,
	0xbf, 0x6e, 0x15, 0xf2, 0xff, 0x29, 0xfc, 0x04, 0xba, 0xdf, 0x4a, 0xf9, 0x82, 0xd2, 0x58, 0xaf,
	0x2b, 0xcd, 0x2d, 0x69, 0xb1, 0x6b, 0xfa, 0x8a, 0x49, 0x3c, 0x82, 0x07, 0x3c, 0xa2, 0xec, 0xa8,
	0x0c, 0xbb, 0x4e, 0xfe, 0xc0, 0x18, 0xb4, 0x2d, 0x4b, 0x13, 0xbd, 0x6e, 0x36, 0x66, 0x1d, 0x47,
	0x61, 0x3c, 0x85, 0x1e, 0x3b, 0xee, 0x78, 0xac, 0x7e, 0xfb, 0x81, 0x87, 0x4c, 0x6f, 0x98, 0x68,
	0xa6, 0x39, 0x77, 0xec, 0x6b, 0xed, 0xef, 0xcf, 0x09, 0xb2, 0xbe, 0x23, 0x18, 0x57, 0xe5, 0x17,
	0x41, 0x20, 0x0e, 0x8c, 0x66, 0xf9, 0x2c, 0x49, 0xf0, 0x73, 0x18, 0x56, 0x9d, 0x5c, 0x92, 0x93,
	0x2a, 0xbf, 0xe5, 0x0c, 0x6e, 0xca, 0x66, 0xe2, 0xa7, 0xd0, 0x27, 0xf9, 0x7a, 0x25, 0xad, 0x2b,
	0x69, 0x8f, 0xdc, 0xba, 0x62, 0xd0, 0x22, 0x52, 0xb4, 0x6a, 0x39, 0x0a, 0x5b, 0x2e, 0x3c, 0xfa,
	0x44, 0x92, 0xb0, 0xc8, 0x5f, 0x8a, 0x48, 0xc6, 0xc4, 0x93, 0x4b, 0x41, 0xd9, 0x3b, 0x8a, 0x9f,
	0xc1, 0xc0, 0x2b, 0x98, 0xbb, 0x16, 0xfd, 0x92, 0x2f, 0xbd, 0xc7, 0xf0, 0xd0, 0x13, 0x94, 0xb9,
	0x9c, 0xaa, 0x70, 0xcd, 0x69, 0x7a, 0xca, 0xe3, 0xcd, 0xea, 0xd7, 0xd9, 0x40, 0xa7, 0xb3, 0x81,
	0xfe, 0x9c, 0x0d, 0xf4, 0xe3, 0x62, 0xd4, 0x4e, 0x17, 0xa3, 0xf6, 0xfb, 0x62, 0xd4, 0x3e, 0xbf,
	0xf2, 0xb9, 0xdc, 0xec, 0xd7, 0x73, 0x4f, 0x84, 0x76, 0x79, 0x10, 0x2f, 0xae, 0xe7, 0x62, 0x57,
	0xe7, 0x62, 0x1f, 0xab, 0xb9, 0x2d, 0xd3, 0x1d, 0x4b, 0xd6, 0x4d, 0x75, 0x27, 0x2f, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x78, 0xe8, 0x16, 0x3f, 0x80, 0x02, 0x00, 0x00,
}

func (this *GuardianSet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GuardianSet)
	if !ok {
		that2, ok := that.(GuardianSet)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if len(this.Keys) != len(that1.Keys) {
		return false
	}
	for i := range this.Keys {
		if !bytes.Equal(this.Keys[i], that1.Keys[i]) {
			return false
		}
	}
	if this.ExpirationTime != that1.ExpirationTime {
		return false
	}
	return true
}
func (m *GuardianKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuardianKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GuardianKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintGuardian(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GuardianValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuardianValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GuardianValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddr) > 0 {
		i -= len(m.ValidatorAddr)
		copy(dAtA[i:], m.ValidatorAddr)
		i = encodeVarintGuardian(dAtA, i, uint64(len(m.ValidatorAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.GuardianKey) > 0 {
		i -= len(m.GuardianKey)
		copy(dAtA[i:], m.GuardianKey)
		i = encodeVarintGuardian(dAtA, i, uint64(len(m.GuardianKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GuardianSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuardianSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GuardianSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpirationTime != 0 {
		i = encodeVarintGuardian(dAtA, i, uint64(m.ExpirationTime))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Keys) > 0 {
		for iNdEx := len(m.Keys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Keys[iNdEx])
			copy(dAtA[i:], m.Keys[iNdEx])
			i = encodeVarintGuardian(dAtA, i, uint64(len(m.Keys[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Index != 0 {
		i = encodeVarintGuardian(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorAllowedAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorAllowedAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorAllowedAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintGuardian(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AllowedAddress) > 0 {
		i -= len(m.AllowedAddress)
		copy(dAtA[i:], m.AllowedAddress)
		i = encodeVarintGuardian(dAtA, i, uint64(len(m.AllowedAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGuardian(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WasmAllowedContractCodeId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WasmAllowedContractCodeId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WasmAllowedContractCodeId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CodeId != 0 {
		i = encodeVarintGuardian(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintGuardian(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGuardian(dAtA []byte, offset int, v uint64) int {
	offset -= sovGuardian(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GuardianKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovGuardian(uint64(l))
	}
	return n
}

func (m *GuardianValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.GuardianKey)
	if l > 0 {
		n += 1 + l + sovGuardian(uint64(l))
	}
	l = len(m.ValidatorAddr)
	if l > 0 {
		n += 1 + l + sovGuardian(uint64(l))
	}
	return n
}

func (m *GuardianSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovGuardian(uint64(m.Index))
	}
	if len(m.Keys) > 0 {
		for _, b := range m.Keys {
			l = len(b)
			n += 1 + l + sovGuardian(uint64(l))
		}
	}
	if m.ExpirationTime != 0 {
		n += 1 + sovGuardian(uint64(m.ExpirationTime))
	}
	return n
}

func (m *ValidatorAllowedAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGuardian(uint64(l))
	}
	l = len(m.AllowedAddress)
	if l > 0 {
		n += 1 + l + sovGuardian(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGuardian(uint64(l))
	}
	return n
}

func (m *WasmAllowedContractCodeId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovGuardian(uint64(l))
	}
	if m.CodeId != 0 {
		n += 1 + sovGuardian(uint64(m.CodeId))
	}
	return n
}

func sovGuardian(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGuardian(x uint64) (n int) {
	return sovGuardian(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GuardianKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuardian
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
			return fmt.Errorf("proto: GuardianKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuardianKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
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
				return ErrInvalidLengthGuardian
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGuardian
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGuardian(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGuardian
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
func (m *GuardianValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuardian
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
			return fmt.Errorf("proto: GuardianValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuardianValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardianKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
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
				return ErrInvalidLengthGuardian
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGuardian
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuardianKey = append(m.GuardianKey[:0], dAtA[iNdEx:postIndex]...)
			if m.GuardianKey == nil {
				m.GuardianKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
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
				return ErrInvalidLengthGuardian
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGuardian
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddr = append(m.ValidatorAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorAddr == nil {
				m.ValidatorAddr = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGuardian(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGuardian
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
func (m *GuardianSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuardian
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
			return fmt.Errorf("proto: GuardianSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuardianSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
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
				return ErrInvalidLengthGuardian
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGuardian
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, make([]byte, postIndex-iNdEx))
			copy(m.Keys[len(m.Keys)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			m.ExpirationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpirationTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGuardian(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGuardian
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
func (m *ValidatorAllowedAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuardian
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
			return fmt.Errorf("proto: ValidatorAllowedAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorAllowedAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
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
				return ErrInvalidLengthGuardian
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGuardian
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
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
				return ErrInvalidLengthGuardian
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGuardian
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
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
				return ErrInvalidLengthGuardian
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGuardian
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGuardian(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGuardian
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
func (m *WasmAllowedContractCodeId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuardian
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
			return fmt.Errorf("proto: WasmAllowedContractCodeId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WasmAllowedContractCodeId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
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
				return ErrInvalidLengthGuardian
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGuardian
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuardian
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGuardian(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGuardian
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
func skipGuardian(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGuardian
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
					return 0, ErrIntOverflowGuardian
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
					return 0, ErrIntOverflowGuardian
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
				return 0, ErrInvalidLengthGuardian
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGuardian
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGuardian
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGuardian        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGuardian          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGuardian = fmt.Errorf("proto: unexpected end of group")
)
