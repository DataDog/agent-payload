// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/network-devices/agent.proto

/*
	Package network_devices is a generated protocol buffer package.

	It is generated from these files:
		proto/network-devices/agent.proto

	It has these top-level messages:
		CollectorNetworkDevice
		CollectorNetworkInterface
		NetworkDevice
		NetworkInterface
*/
package network_devices

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CollectorNetworkDevice struct {
	AutodiscoverySubnet string           `protobuf:"bytes,1,opt,name=autodiscoverySubnet,proto3" json:"autodiscoverySubnet,omitempty"`
	Devices             []*NetworkDevice `protobuf:"bytes,2,rep,name=devices" json:"devices,omitempty"`
	Tags                []string         `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
}

func (m *CollectorNetworkDevice) Reset()                    { *m = CollectorNetworkDevice{} }
func (m *CollectorNetworkDevice) String() string            { return proto.CompactTextString(m) }
func (*CollectorNetworkDevice) ProtoMessage()               {}
func (*CollectorNetworkDevice) Descriptor() ([]byte, []int) { return fileDescriptorAgent, []int{0} }

func (m *CollectorNetworkDevice) GetAutodiscoverySubnet() string {
	if m != nil {
		return m.AutodiscoverySubnet
	}
	return ""
}

func (m *CollectorNetworkDevice) GetDevices() []*NetworkDevice {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *CollectorNetworkDevice) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CollectorNetworkInterface struct {
	DeviceId   string              `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	Interfaces []*NetworkInterface `protobuf:"bytes,2,rep,name=interfaces" json:"interfaces,omitempty"`
}

func (m *CollectorNetworkInterface) Reset()                    { *m = CollectorNetworkInterface{} }
func (m *CollectorNetworkInterface) String() string            { return proto.CompactTextString(m) }
func (*CollectorNetworkInterface) ProtoMessage()               {}
func (*CollectorNetworkInterface) Descriptor() ([]byte, []int) { return fileDescriptorAgent, []int{1} }

func (m *CollectorNetworkInterface) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *CollectorNetworkInterface) GetInterfaces() []*NetworkInterface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type NetworkDevice struct {
	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IpAddress   string   `protobuf:"bytes,4,opt,name=ipAddress,proto3" json:"ipAddress,omitempty"`
	SysObjectId string   `protobuf:"bytes,5,opt,name=sysObjectId,proto3" json:"sysObjectId,omitempty"`
	Profile     string   `protobuf:"bytes,6,opt,name=profile,proto3" json:"profile,omitempty"`
	Vendor      string   `protobuf:"bytes,7,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Tags        []string `protobuf:"bytes,8,rep,name=tags" json:"tags,omitempty"`
}

func (m *NetworkDevice) Reset()                    { *m = NetworkDevice{} }
func (m *NetworkDevice) String() string            { return proto.CompactTextString(m) }
func (*NetworkDevice) ProtoMessage()               {}
func (*NetworkDevice) Descriptor() ([]byte, []int) { return fileDescriptorAgent, []int{2} }

func (m *NetworkDevice) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NetworkDevice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkDevice) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NetworkDevice) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *NetworkDevice) GetSysObjectId() string {
	if m != nil {
		return m.SysObjectId
	}
	return ""
}

func (m *NetworkDevice) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *NetworkDevice) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *NetworkDevice) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type NetworkInterface struct {
	Index       int32  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Alias       string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	MacAddress  string `protobuf:"bytes,5,opt,name=macAddress,proto3" json:"macAddress,omitempty"`
	AdminStatus int32  `protobuf:"varint,6,opt,name=adminStatus,proto3" json:"adminStatus,omitempty"`
	OperStatus  int32  `protobuf:"varint,7,opt,name=operStatus,proto3" json:"operStatus,omitempty"`
}

func (m *NetworkInterface) Reset()                    { *m = NetworkInterface{} }
func (m *NetworkInterface) String() string            { return proto.CompactTextString(m) }
func (*NetworkInterface) ProtoMessage()               {}
func (*NetworkInterface) Descriptor() ([]byte, []int) { return fileDescriptorAgent, []int{3} }

func (m *NetworkInterface) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *NetworkInterface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkInterface) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *NetworkInterface) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NetworkInterface) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *NetworkInterface) GetAdminStatus() int32 {
	if m != nil {
		return m.AdminStatus
	}
	return 0
}

func (m *NetworkInterface) GetOperStatus() int32 {
	if m != nil {
		return m.OperStatus
	}
	return 0
}

func init() {
	proto.RegisterType((*CollectorNetworkDevice)(nil), "network_devices.CollectorNetworkDevice")
	proto.RegisterType((*CollectorNetworkInterface)(nil), "network_devices.CollectorNetworkInterface")
	proto.RegisterType((*NetworkDevice)(nil), "network_devices.NetworkDevice")
	proto.RegisterType((*NetworkInterface)(nil), "network_devices.NetworkInterface")
}
func (m *CollectorNetworkDevice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectorNetworkDevice) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AutodiscoverySubnet) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.AutodiscoverySubnet)))
		i += copy(dAtA[i:], m.AutodiscoverySubnet)
	}
	if len(m.Devices) > 0 {
		for _, msg := range m.Devices {
			dAtA[i] = 0x12
			i++
			i = encodeVarintAgent(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *CollectorNetworkInterface) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectorNetworkInterface) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DeviceId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.DeviceId)))
		i += copy(dAtA[i:], m.DeviceId)
	}
	if len(m.Interfaces) > 0 {
		for _, msg := range m.Interfaces {
			dAtA[i] = 0x12
			i++
			i = encodeVarintAgent(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *NetworkDevice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkDevice) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.IpAddress) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.IpAddress)))
		i += copy(dAtA[i:], m.IpAddress)
	}
	if len(m.SysObjectId) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.SysObjectId)))
		i += copy(dAtA[i:], m.SysObjectId)
	}
	if len(m.Profile) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Profile)))
		i += copy(dAtA[i:], m.Profile)
	}
	if len(m.Vendor) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Vendor)))
		i += copy(dAtA[i:], m.Vendor)
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			dAtA[i] = 0x42
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *NetworkInterface) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkInterface) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAgent(dAtA, i, uint64(m.Index))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Alias) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Alias)))
		i += copy(dAtA[i:], m.Alias)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if len(m.MacAddress) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintAgent(dAtA, i, uint64(len(m.MacAddress)))
		i += copy(dAtA[i:], m.MacAddress)
	}
	if m.AdminStatus != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintAgent(dAtA, i, uint64(m.AdminStatus))
	}
	if m.OperStatus != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintAgent(dAtA, i, uint64(m.OperStatus))
	}
	return i, nil
}

func encodeVarintAgent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CollectorNetworkDevice) Size() (n int) {
	var l int
	_ = l
	l = len(m.AutodiscoverySubnet)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	if len(m.Devices) > 0 {
		for _, e := range m.Devices {
			l = e.Size()
			n += 1 + l + sovAgent(uint64(l))
		}
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovAgent(uint64(l))
		}
	}
	return n
}

func (m *CollectorNetworkInterface) Size() (n int) {
	var l int
	_ = l
	l = len(m.DeviceId)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	if len(m.Interfaces) > 0 {
		for _, e := range m.Interfaces {
			l = e.Size()
			n += 1 + l + sovAgent(uint64(l))
		}
	}
	return n
}

func (m *NetworkDevice) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.IpAddress)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.SysObjectId)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.Profile)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.Vendor)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovAgent(uint64(l))
		}
	}
	return n
}

func (m *NetworkInterface) Size() (n int) {
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovAgent(uint64(m.Index))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.Alias)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	l = len(m.MacAddress)
	if l > 0 {
		n += 1 + l + sovAgent(uint64(l))
	}
	if m.AdminStatus != 0 {
		n += 1 + sovAgent(uint64(m.AdminStatus))
	}
	if m.OperStatus != 0 {
		n += 1 + sovAgent(uint64(m.OperStatus))
	}
	return n
}

func sovAgent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAgent(x uint64) (n int) {
	return sovAgent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CollectorNetworkDevice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CollectorNetworkDevice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectorNetworkDevice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutodiscoverySubnet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AutodiscoverySubnet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Devices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Devices = append(m.Devices, &NetworkDevice{})
			if err := m.Devices[len(m.Devices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
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
func (m *CollectorNetworkInterface) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CollectorNetworkInterface: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectorNetworkInterface: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interfaces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Interfaces = append(m.Interfaces, &NetworkInterface{})
			if err := m.Interfaces[len(m.Interfaces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
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
func (m *NetworkDevice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetworkDevice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkDevice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SysObjectId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SysObjectId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Profile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vendor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vendor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
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
func (m *NetworkInterface) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAgent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetworkInterface: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkInterface: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MacAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAgent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MacAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminStatus", wireType)
			}
			m.AdminStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdminStatus |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperStatus", wireType)
			}
			m.OperStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OperStatus |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAgent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAgent
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
func skipAgent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAgent
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
					return 0, ErrIntOverflowAgent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAgent
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAgent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAgent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAgent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAgent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAgent   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("proto/network-devices/agent.proto", fileDescriptorAgent) }

var fileDescriptorAgent = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xe5, 0x24, 0x9b, 0x34, 0x53, 0xf1, 0x47, 0xa6, 0xaa, 0x0c, 0x42, 0xab, 0x34, 0xa7,
	0x5c, 0x9a, 0x54, 0x70, 0xe1, 0x84, 0x54, 0xc8, 0x25, 0x17, 0x90, 0xb6, 0x37, 0x2e, 0xc8, 0xb1,
	0xa7, 0xc1, 0xb0, 0x6b, 0xaf, 0x6c, 0x6f, 0x20, 0x3c, 0x09, 0x8f, 0xc4, 0x91, 0x13, 0x17, 0x2e,
	0x28, 0x4f, 0x82, 0xd6, 0xbb, 0x1b, 0xb6, 0xdb, 0xf6, 0xe6, 0xf9, 0x7e, 0x33, 0xe3, 0xf9, 0x46,
	0x1a, 0x38, 0xcb, 0xad, 0xf1, 0x66, 0xa1, 0xd1, 0x7f, 0x35, 0xf6, 0xcb, 0xb9, 0xc4, 0xad, 0x12,
	0xe8, 0x16, 0x7c, 0x83, 0xda, 0xcf, 0x03, 0xa3, 0x8f, 0x6a, 0xf8, 0xb1, 0x86, 0xd3, 0x1f, 0x04,
	0x4e, 0xdf, 0x9a, 0x34, 0x45, 0xe1, 0x8d, 0x7d, 0x57, 0xc1, 0x65, 0x60, 0xf4, 0x02, 0x9e, 0xf0,
	0xc2, 0x1b, 0xa9, 0x9c, 0x30, 0x5b, 0xb4, 0xbb, 0xab, 0x62, 0xad, 0xd1, 0x33, 0x32, 0x21, 0xb3,
	0x71, 0x72, 0x17, 0xa2, 0xaf, 0x60, 0x54, 0xf7, 0x65, 0xbd, 0x49, 0x7f, 0x76, 0xfc, 0x22, 0x9e,
	0x77, 0xfe, 0x9b, 0xdf, 0xf8, 0x22, 0x69, 0xd2, 0x29, 0x85, 0x81, 0xe7, 0x1b, 0xc7, 0xfa, 0x93,
	0xfe, 0x6c, 0x9c, 0x84, 0xf7, 0xf4, 0x3b, 0x3c, 0xed, 0x4e, 0xb6, 0xd2, 0x1e, 0xed, 0x35, 0x17,
	0x48, 0x9f, 0xc1, 0x51, 0x55, 0xbb, 0x92, 0xf5, 0x44, 0x87, 0x98, 0x5e, 0x02, 0xa8, 0x26, 0xb1,
	0x99, 0xe4, 0xec, 0xbe, 0x49, 0x0e, 0x2d, 0x93, 0x56, 0xd1, 0xf4, 0x0f, 0x81, 0x07, 0x37, 0xb7,
	0xf1, 0x10, 0x7a, 0xaa, 0xf9, 0xaa, 0xa7, 0x64, 0x39, 0xb1, 0xe6, 0x19, 0xb2, 0x5e, 0x50, 0xc2,
	0x9b, 0x4e, 0xe0, 0x58, 0xa2, 0x13, 0x56, 0xe5, 0x5e, 0x19, 0xcd, 0xfa, 0x01, 0xb5, 0x25, 0xfa,
	0x1c, 0xc6, 0x2a, 0xbf, 0x94, 0xd2, 0xa2, 0x73, 0x6c, 0x10, 0xf8, 0x7f, 0xa1, 0xac, 0x77, 0x3b,
	0xf7, 0x7e, 0xfd, 0x19, 0x85, 0x5f, 0x49, 0x16, 0x55, 0xf5, 0x2d, 0x89, 0x32, 0x18, 0xe5, 0xd6,
	0x5c, 0xab, 0x14, 0xd9, 0x30, 0xd0, 0x26, 0xa4, 0xa7, 0x30, 0xdc, 0xa2, 0x96, 0xc6, 0xb2, 0x51,
	0x00, 0x75, 0x74, 0xd8, 0xec, 0x51, 0x6b, 0xb3, 0xbf, 0x09, 0x3c, 0xbe, 0xb5, 0xd1, 0x13, 0x88,
	0x94, 0x96, 0xf8, 0x2d, 0x78, 0x8c, 0x92, 0x2a, 0xb8, 0xd3, 0xe6, 0x09, 0x44, 0x3c, 0x55, 0xdc,
	0xd5, 0x06, 0xab, 0xa0, 0x6b, 0x7e, 0x70, 0xdb, 0x7c, 0x0c, 0x90, 0x71, 0xd1, 0xb8, 0xaf, 0xdc,
	0xb5, 0x94, 0xb2, 0x03, 0x97, 0x99, 0xd2, 0x57, 0x9e, 0xfb, 0xc2, 0x05, 0x83, 0x51, 0xd2, 0x96,
	0xca, 0x0e, 0x26, 0x47, 0x5b, 0x27, 0x8c, 0x42, 0x42, 0x4b, 0x79, 0xf3, 0xfa, 0xe7, 0x3e, 0x26,
	0xbf, 0xf6, 0x31, 0xf9, 0xbb, 0x8f, 0xc9, 0x87, 0x8b, 0x8d, 0xf2, 0x9f, 0x8a, 0xf5, 0x5c, 0x98,
	0x6c, 0xb1, 0xe4, 0x9e, 0x2f, 0xcd, 0xa6, 0x3a, 0x86, 0xf3, 0x9c, 0xef, 0x52, 0xc3, 0x65, 0xf7,
	0x54, 0xd6, 0xc3, 0x70, 0x25, 0x2f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x13, 0x71, 0x2b, 0x08,
	0x4a, 0x03, 0x00, 0x00,
}
