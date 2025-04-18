// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: github.com/DataDog/agent-payload/v5/proto/autoscaling/kubernetes/common.proto

package kubernetes

import (
	_ "github.com/chrusty/protoc-gen-jsonschema"
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

// Error represents an error message
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    *int32 `protobuf:"varint,1,opt,name=code,proto3,oneof" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() int32 {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// WorkloadTargetRef represents the target reference of a workload
type WorkloadTargetRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster    string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`       // Cluster is the name of the cluster
	Namespace  string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`   // Namespace is the namespace of PodAutoscaler object
	Kind       string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`             // Kind is K8S object kind
	Name       string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`             // Name is the name of the K8S object
	ApiVersion string `protobuf:"bytes,5,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"` // APIVersion is the API version of the K8S object
}

func (x *WorkloadTargetRef) Reset() {
	*x = WorkloadTargetRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadTargetRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadTargetRef) ProtoMessage() {}

func (x *WorkloadTargetRef) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadTargetRef.ProtoReflect.Descriptor instead.
func (*WorkloadTargetRef) Descriptor() ([]byte, []int) {
	return file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDescGZIP(), []int{1}
}

func (x *WorkloadTargetRef) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *WorkloadTargetRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *WorkloadTargetRef) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *WorkloadTargetRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkloadTargetRef) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

// WorkloadState represents the state of a workload
type WorkloadState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DesiredReplicas int32  `protobuf:"varint,1,opt,name=desiredReplicas,proto3" json:"desiredReplicas,omitempty"`       // DesiredReplicas is the current number of desired replicas
	CurrentReplicas *int32 `protobuf:"varint,2,opt,name=currentReplicas,proto3,oneof" json:"currentReplicas,omitempty"` // CurrentReplicas is the number of replicas currently running
	ReadyReplicas   *int32 `protobuf:"varint,3,opt,name=readyReplicas,proto3,oneof" json:"readyReplicas,omitempty"`     // ReadyReplicas is the number of ready replicas
}

func (x *WorkloadState) Reset() {
	*x = WorkloadState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadState) ProtoMessage() {}

func (x *WorkloadState) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadState.ProtoReflect.Descriptor instead.
func (*WorkloadState) Descriptor() ([]byte, []int) {
	return file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDescGZIP(), []int{2}
}

func (x *WorkloadState) GetDesiredReplicas() int32 {
	if x != nil {
		return x.DesiredReplicas
	}
	return 0
}

func (x *WorkloadState) GetCurrentReplicas() int32 {
	if x != nil && x.CurrentReplicas != nil {
		return *x.CurrentReplicas
	}
	return 0
}

func (x *WorkloadState) GetReadyReplicas() int32 {
	if x != nil && x.ReadyReplicas != nil {
		return *x.ReadyReplicas
	}
	return 0
}

var File_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto protoreflect.FileDescriptor

var file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74,
	0x61, 0x44, 0x6f, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x2f, 0x76, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x1a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x72, 0x75,
	0x73, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6a,
	0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x3a, 0x07, 0xba, 0x46, 0x04, 0x08, 0x01, 0x20, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x12, 0x1f, 0x0a, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xaa, 0x46,
	0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x05, 0xaa, 0x46, 0x02, 0x10, 0x01, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x19, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x05, 0xaa, 0x46, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xaa, 0x46, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x07, 0xba, 0x46, 0x04, 0x08, 0x01, 0x20, 0x01,
	0x22, 0xc9, 0x01, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x2f, 0x0a, 0x0f, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x05, 0xaa, 0x46, 0x02,
	0x10, 0x01, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x73, 0x12, 0x2d, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x0d, 0x72, 0x65, 0x61,
	0x64, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x88, 0x01, 0x01, 0x3a, 0x07, 0xba,
	0x46, 0x04, 0x08, 0x01, 0x20, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x72,
	0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x44,
	0x6f, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x2f, 0x76, 0x35, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDescOnce sync.Once
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDescData = file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDesc
)

func file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDescGZIP() []byte {
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDescOnce.Do(func() {
		file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDescData)
	})
	return file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDescData
}

var file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_goTypes = []interface{}{
	(*Error)(nil),             // 0: datadog.autoscaling.kubernetes.Error
	(*WorkloadTargetRef)(nil), // 1: datadog.autoscaling.kubernetes.WorkloadTargetRef
	(*WorkloadState)(nil),     // 2: datadog.autoscaling.kubernetes.WorkloadState
}
var file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_init()
}
func file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_init() {
	if File_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadTargetRef); i {
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
		file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadState); i {
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
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_goTypes,
		DependencyIndexes: file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_depIdxs,
		MessageInfos:      file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_msgTypes,
	}.Build()
	File_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto = out.File
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_rawDesc = nil
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_goTypes = nil
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_depIdxs = nil
}
