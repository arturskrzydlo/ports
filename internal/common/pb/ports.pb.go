// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: ports.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	City        string    `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country     string    `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Alias       []string  `protobuf:"bytes,4,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions     []string  `protobuf:"bytes,5,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates []float64 `protobuf:"fixed64,6,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province    string    `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`
	Timezone    string    `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs      []string  `protobuf:"bytes,9,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code        string    `protobuf:"bytes,10,opt,name=code,proto3" json:"code,omitempty"`
	Id          string    `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_ports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_ports_proto_rawDescGZIP(), []int{0}
}

func (x *Port) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Port) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Port) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Port) GetAlias() []string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *Port) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Port) GetCoordinates() []float64 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

func (x *Port) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Port) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Port) GetUnlocs() []string {
	if x != nil {
		return x.Unlocs
	}
	return nil
}

func (x *Port) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Port) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreatePortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port *Port `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *CreatePortRequest) Reset() {
	*x = CreatePortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePortRequest) ProtoMessage() {}

func (x *CreatePortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePortRequest.ProtoReflect.Descriptor instead.
func (*CreatePortRequest) Descriptor() ([]byte, []int) {
	return file_ports_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePortRequest) GetPort() *Port {
	if x != nil {
		return x.Port
	}
	return nil
}

type GetPortsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports []*Port `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
}

func (x *GetPortsResponse) Reset() {
	*x = GetPortsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPortsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPortsResponse) ProtoMessage() {}

func (x *GetPortsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPortsResponse.ProtoReflect.Descriptor instead.
func (*GetPortsResponse) Descriptor() ([]byte, []int) {
	return file_ports_proto_rawDescGZIP(), []int{2}
}

func (x *GetPortsResponse) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

var File_ports_proto protoreflect.FileDescriptor

var file_ports_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x34, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x32,
	0x8e, 0x01, 0x0a, 0x0b, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x72, 0x74, 0x75, 0x72, 0x73, 0x6b, 0x72, 0x7a, 0x79, 0x64, 0x6c, 0x6f, 0x2f, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ports_proto_rawDescOnce sync.Once
	file_ports_proto_rawDescData = file_ports_proto_rawDesc
)

func file_ports_proto_rawDescGZIP() []byte {
	file_ports_proto_rawDescOnce.Do(func() {
		file_ports_proto_rawDescData = protoimpl.X.CompressGZIP(file_ports_proto_rawDescData)
	})
	return file_ports_proto_rawDescData
}

var (
	file_ports_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_ports_proto_goTypes  = []interface{}{
		(*Port)(nil),              // 0: ports.Port
		(*CreatePortRequest)(nil), // 1: ports.CreatePortRequest
		(*GetPortsResponse)(nil),  // 2: ports.GetPortsResponse
		(*emptypb.Empty)(nil),     // 3: google.protobuf.Empty
	}
)
var file_ports_proto_depIdxs = []int32{
	0, // 0: ports.CreatePortRequest.port:type_name -> ports.Port
	0, // 1: ports.GetPortsResponse.ports:type_name -> ports.Port
	1, // 2: ports.PortService.CreatePort:input_type -> ports.CreatePortRequest
	3, // 3: ports.PortService.GetPorts:input_type -> google.protobuf.Empty
	3, // 4: ports.PortService.CreatePort:output_type -> google.protobuf.Empty
	2, // 5: ports.PortService.GetPorts:output_type -> ports.GetPortsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ports_proto_init() }
func file_ports_proto_init() {
	if File_ports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePortRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPortsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ports_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ports_proto_goTypes,
		DependencyIndexes: file_ports_proto_depIdxs,
		MessageInfos:      file_ports_proto_msgTypes,
	}.Build()
	File_ports_proto = out.File
	file_ports_proto_rawDesc = nil
	file_ports_proto_goTypes = nil
	file_ports_proto_depIdxs = nil
}
