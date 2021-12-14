// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: pkg/plugin/proto/hwservice.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PollTimeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollTime uint64 `protobuf:"varint,1,opt,name=pollTime,proto3" json:"pollTime,omitempty"`
}

func (x *PollTimeReply) Reset() {
	*x = PollTimeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_proto_hwservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollTimeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollTimeReply) ProtoMessage() {}

func (x *PollTimeReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_proto_hwservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollTimeReply.ProtoReflect.Descriptor instead.
func (*PollTimeReply) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_proto_hwservice_proto_rawDescGZIP(), []int{0}
}

func (x *PollTimeReply) GetPollTime() uint64 {
	if x != nil {
		return x.PollTime
	}
	return 0
}

type Sensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Sensor) Reset() {
	*x = Sensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_proto_hwservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sensor) ProtoMessage() {}

func (x *Sensor) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_proto_hwservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sensor.ProtoReflect.Descriptor instead.
func (*Sensor) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_proto_hwservice_proto_rawDescGZIP(), []int{1}
}

func (x *Sensor) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Sensor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SensorIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SensorIDRequest) Reset() {
	*x = SensorIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_proto_hwservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SensorIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SensorIDRequest) ProtoMessage() {}

func (x *SensorIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_proto_hwservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SensorIDRequest.ProtoReflect.Descriptor instead.
func (*SensorIDRequest) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_proto_hwservice_proto_rawDescGZIP(), []int{2}
}

func (x *SensorIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Reading struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TypeI    int32   `protobuf:"varint,2,opt,name=typeI,proto3" json:"typeI,omitempty"`
	Type     string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Label    string  `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	Unit     string  `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	Value    float64 `protobuf:"fixed64,6,opt,name=value,proto3" json:"value,omitempty"`
	ValueMin float64 `protobuf:"fixed64,7,opt,name=valueMin,proto3" json:"valueMin,omitempty"`
	ValueMax float64 `protobuf:"fixed64,8,opt,name=valueMax,proto3" json:"valueMax,omitempty"`
	ValueAvg float64 `protobuf:"fixed64,9,opt,name=valueAvg,proto3" json:"valueAvg,omitempty"`
}

func (x *Reading) Reset() {
	*x = Reading{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugin_proto_hwservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reading) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reading) ProtoMessage() {}

func (x *Reading) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugin_proto_hwservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reading.ProtoReflect.Descriptor instead.
func (*Reading) Descriptor() ([]byte, []int) {
	return file_pkg_plugin_proto_hwservice_proto_rawDescGZIP(), []int{3}
}

func (x *Reading) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Reading) GetTypeI() int32 {
	if x != nil {
		return x.TypeI
	}
	return 0
}

func (x *Reading) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Reading) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Reading) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Reading) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Reading) GetValueMin() float64 {
	if x != nil {
		return x.ValueMin
	}
	return 0
}

func (x *Reading) GetValueMax() float64 {
	if x != nil {
		return x.ValueMax
	}
	return 0
}

func (x *Reading) GetValueAvg() float64 {
	if x != nil {
		return x.ValueAvg
	}
	return 0
}

var File_pkg_plugin_proto_hwservice_proto protoreflect.FileDescriptor

var file_pkg_plugin_proto_hwservice_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x68, 0x77, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0d, 0x50, 0x6f, 0x6c, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x21, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xd7, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x49, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x49, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x4d, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x4d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4d,
	0x61, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4d,
	0x61, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x41, 0x76, 0x67, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x41, 0x76, 0x67, 0x32, 0xc0,
	0x01, 0x0a, 0x09, 0x48, 0x57, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x08,
	0x50, 0x6f, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x41,
	0x0a, 0x13, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x46, 0x6f, 0x72, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x68, 0x61, 0x79, 0x6e, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x68,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_plugin_proto_hwservice_proto_rawDescOnce sync.Once
	file_pkg_plugin_proto_hwservice_proto_rawDescData = file_pkg_plugin_proto_hwservice_proto_rawDesc
)

func file_pkg_plugin_proto_hwservice_proto_rawDescGZIP() []byte {
	file_pkg_plugin_proto_hwservice_proto_rawDescOnce.Do(func() {
		file_pkg_plugin_proto_hwservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_plugin_proto_hwservice_proto_rawDescData)
	})
	return file_pkg_plugin_proto_hwservice_proto_rawDescData
}

var file_pkg_plugin_proto_hwservice_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_plugin_proto_hwservice_proto_goTypes = []interface{}{
	(*PollTimeReply)(nil),   // 0: proto.PollTimeReply
	(*Sensor)(nil),          // 1: proto.Sensor
	(*SensorIDRequest)(nil), // 2: proto.SensorIDRequest
	(*Reading)(nil),         // 3: proto.Reading
	(*empty.Empty)(nil),     // 4: google.protobuf.Empty
}
var file_pkg_plugin_proto_hwservice_proto_depIdxs = []int32{
	4, // 0: proto.HWService.PollTime:input_type -> google.protobuf.Empty
	4, // 1: proto.HWService.Sensors:input_type -> google.protobuf.Empty
	2, // 2: proto.HWService.ReadingsForSensorID:input_type -> proto.SensorIDRequest
	0, // 3: proto.HWService.PollTime:output_type -> proto.PollTimeReply
	1, // 4: proto.HWService.Sensors:output_type -> proto.Sensor
	3, // 5: proto.HWService.ReadingsForSensorID:output_type -> proto.Reading
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_plugin_proto_hwservice_proto_init() }
func file_pkg_plugin_proto_hwservice_proto_init() {
	if File_pkg_plugin_proto_hwservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_plugin_proto_hwservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PollTimeReply); i {
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
		file_pkg_plugin_proto_hwservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sensor); i {
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
		file_pkg_plugin_proto_hwservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SensorIDRequest); i {
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
		file_pkg_plugin_proto_hwservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reading); i {
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
			RawDescriptor: file_pkg_plugin_proto_hwservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_plugin_proto_hwservice_proto_goTypes,
		DependencyIndexes: file_pkg_plugin_proto_hwservice_proto_depIdxs,
		MessageInfos:      file_pkg_plugin_proto_hwservice_proto_msgTypes,
	}.Build()
	File_pkg_plugin_proto_hwservice_proto = out.File
	file_pkg_plugin_proto_hwservice_proto_rawDesc = nil
	file_pkg_plugin_proto_hwservice_proto_goTypes = nil
	file_pkg_plugin_proto_hwservice_proto_depIdxs = nil
}