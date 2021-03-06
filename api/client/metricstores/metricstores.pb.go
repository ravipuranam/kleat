// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: metricstores/metricstores.proto

package metricstores

import (
	proto "github.com/golang/protobuf/proto"
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

// Configuration for the Spinnaker monitoring daemon metric stores.
type MetricStores struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for the Datadog metric store.
	Datadog *Datadog `protobuf:"bytes,1,opt,name=datadog,proto3" json:"datadog,omitempty"`
	// Configuration for the Newrelic metric store.
	Newrelic *Newrelic `protobuf:"bytes,2,opt,name=newrelic,proto3" json:"newrelic,omitempty"`
	// Configuration for the Prometheus metric store.
	Prometheus *Prometheus `protobuf:"bytes,3,opt,name=prometheus,proto3" json:"prometheus,omitempty"`
	// Configuration for the Stackdriver metric store.
	Stackdriver *Stackdriver `protobuf:"bytes,4,opt,name=stackdriver,proto3" json:"stackdriver,omitempty"`
	// Polling period for the monitoring daemon (seconds). Defaults to 30.
	Period int32 `protobuf:"varint,5,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *MetricStores) Reset() {
	*x = MetricStores{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metricstores_metricstores_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricStores) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricStores) ProtoMessage() {}

func (x *MetricStores) ProtoReflect() protoreflect.Message {
	mi := &file_metricstores_metricstores_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricStores.ProtoReflect.Descriptor instead.
func (*MetricStores) Descriptor() ([]byte, []int) {
	return file_metricstores_metricstores_proto_rawDescGZIP(), []int{0}
}

func (x *MetricStores) GetDatadog() *Datadog {
	if x != nil {
		return x.Datadog
	}
	return nil
}

func (x *MetricStores) GetNewrelic() *Newrelic {
	if x != nil {
		return x.Newrelic
	}
	return nil
}

func (x *MetricStores) GetPrometheus() *Prometheus {
	if x != nil {
		return x.Prometheus
	}
	return nil
}

func (x *MetricStores) GetStackdriver() *Stackdriver {
	if x != nil {
		return x.Stackdriver
	}
	return nil
}

func (x *MetricStores) GetPeriod() int32 {
	if x != nil {
		return x.Period
	}
	return 0
}

var File_metricstores_metricstores_proto protoreflect.FileDescriptor

var file_metricstores_metricstores_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x73, 0x1a, 0x1a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f,
	0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02,
	0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x35,
	0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x52, 0x07, 0x64, 0x61,
	0x74, 0x61, 0x64, 0x6f, 0x67, 0x12, 0x38, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x4e, 0x65, 0x77,
	0x72, 0x65, 0x6c, 0x69, 0x63, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x72, 0x65, 0x6c, 0x69, 0x63, 0x12,
	0x3e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68,
	0x65, 0x75, 0x73, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x12,
	0x41, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b,
	0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metricstores_metricstores_proto_rawDescOnce sync.Once
	file_metricstores_metricstores_proto_rawDescData = file_metricstores_metricstores_proto_rawDesc
)

func file_metricstores_metricstores_proto_rawDescGZIP() []byte {
	file_metricstores_metricstores_proto_rawDescOnce.Do(func() {
		file_metricstores_metricstores_proto_rawDescData = protoimpl.X.CompressGZIP(file_metricstores_metricstores_proto_rawDescData)
	})
	return file_metricstores_metricstores_proto_rawDescData
}

var file_metricstores_metricstores_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_metricstores_metricstores_proto_goTypes = []interface{}{
	(*MetricStores)(nil), // 0: proto.metricstores.MetricStores
	(*Datadog)(nil),      // 1: proto.metricstores.Datadog
	(*Newrelic)(nil),     // 2: proto.metricstores.Newrelic
	(*Prometheus)(nil),   // 3: proto.metricstores.Prometheus
	(*Stackdriver)(nil),  // 4: proto.metricstores.Stackdriver
}
var file_metricstores_metricstores_proto_depIdxs = []int32{
	1, // 0: proto.metricstores.MetricStores.datadog:type_name -> proto.metricstores.Datadog
	2, // 1: proto.metricstores.MetricStores.newrelic:type_name -> proto.metricstores.Newrelic
	3, // 2: proto.metricstores.MetricStores.prometheus:type_name -> proto.metricstores.Prometheus
	4, // 3: proto.metricstores.MetricStores.stackdriver:type_name -> proto.metricstores.Stackdriver
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_metricstores_metricstores_proto_init() }
func file_metricstores_metricstores_proto_init() {
	if File_metricstores_metricstores_proto != nil {
		return
	}
	file_metricstores_datadog_proto_init()
	file_metricstores_newrelic_proto_init()
	file_metricstores_prometheus_proto_init()
	file_metricstores_stackdriver_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metricstores_metricstores_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricStores); i {
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
			RawDescriptor: file_metricstores_metricstores_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metricstores_metricstores_proto_goTypes,
		DependencyIndexes: file_metricstores_metricstores_proto_depIdxs,
		MessageInfos:      file_metricstores_metricstores_proto_msgTypes,
	}.Build()
	File_metricstores_metricstores_proto = out.File
	file_metricstores_metricstores_proto_rawDesc = nil
	file_metricstores_metricstores_proto_goTypes = nil
	file_metricstores_metricstores_proto_depIdxs = nil
}
