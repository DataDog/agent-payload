// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: github.com/DataDog/agent-payload/v5/proto/autoscaling/kubernetes/recommender.proto

package kubernetes

import (
	_ "github.com/chrusty/protoc-gen-jsonschema"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// WorkloadValuesList represents a list of workload recommendation requests
type WorkloadRecommendationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*WorkloadRecommendationRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *WorkloadRecommendationsRequest) Reset() {
	*x = WorkloadRecommendationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadRecommendationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadRecommendationsRequest) ProtoMessage() {}

func (x *WorkloadRecommendationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadRecommendationsRequest.ProtoReflect.Descriptor instead.
func (*WorkloadRecommendationsRequest) Descriptor() ([]byte, []int) {
	return file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescGZIP(), []int{0}
}

func (x *WorkloadRecommendationsRequest) GetRequests() []*WorkloadRecommendationRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

// WorkloadRecommendationRequest represents a recommendation request
type WorkloadRecommendationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetRef   *WorkloadTargetRef                 `protobuf:"bytes,1,opt,name=targetRef,proto3" json:"targetRef,omitempty"` // TargetRef is the target reference of the workload
	Constraints *WorkloadRecommendationConstraints `protobuf:"bytes,2,opt,name=constraints,proto3" json:"constraints,omitempty"`
	State       *WorkloadState                     `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Targets     []*WorkloadRecommendationTarget    `protobuf:"bytes,4,rep,name=targets,proto3" json:"targets,omitempty"`
	Settings    map[string]*structpb.Value         `protobuf:"bytes,5,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WorkloadRecommendationRequest) Reset() {
	*x = WorkloadRecommendationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadRecommendationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadRecommendationRequest) ProtoMessage() {}

func (x *WorkloadRecommendationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadRecommendationRequest.ProtoReflect.Descriptor instead.
func (*WorkloadRecommendationRequest) Descriptor() ([]byte, []int) {
	return file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescGZIP(), []int{1}
}

func (x *WorkloadRecommendationRequest) GetTargetRef() *WorkloadTargetRef {
	if x != nil {
		return x.TargetRef
	}
	return nil
}

func (x *WorkloadRecommendationRequest) GetConstraints() *WorkloadRecommendationConstraints {
	if x != nil {
		return x.Constraints
	}
	return nil
}

func (x *WorkloadRecommendationRequest) GetState() *WorkloadState {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *WorkloadRecommendationRequest) GetTargets() []*WorkloadRecommendationTarget {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *WorkloadRecommendationRequest) GetSettings() map[string]*structpb.Value {
	if x != nil {
		return x.Settings
	}
	return nil
}

// WorkloadRecommendationConstraints represents the constraints for a workload
type WorkloadRecommendationConstraints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinReplicas int32 `protobuf:"varint,1,opt,name=minReplicas,proto3" json:"minReplicas,omitempty"` // MinReplicas is the minimum number of replicas the workload should have
	MaxReplicas int32 `protobuf:"varint,2,opt,name=maxReplicas,proto3" json:"maxReplicas,omitempty"` // MaxReplicas is the maximum number of replicas the workload should have
}

func (x *WorkloadRecommendationConstraints) Reset() {
	*x = WorkloadRecommendationConstraints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadRecommendationConstraints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadRecommendationConstraints) ProtoMessage() {}

func (x *WorkloadRecommendationConstraints) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadRecommendationConstraints.ProtoReflect.Descriptor instead.
func (*WorkloadRecommendationConstraints) Descriptor() ([]byte, []int) {
	return file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescGZIP(), []int{2}
}

func (x *WorkloadRecommendationConstraints) GetMinReplicas() int32 {
	if x != nil {
		return x.MinReplicas
	}
	return 0
}

func (x *WorkloadRecommendationConstraints) GetMaxReplicas() int32 {
	if x != nil {
		return x.MaxReplicas
	}
	return 0
}

// WorkloadRecommendationTarget represents a target value for a workload
type WorkloadRecommendationTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`                 // Type defines the target value type (e.g. CPU, Memory, custom, etc.)
	TargetValue float64 `protobuf:"fixed64,2,opt,name=targetValue,proto3" json:"targetValue,omitempty"` // TargetValue is the target value of the target
	LowerBound  float64 `protobuf:"fixed64,3,opt,name=lowerBound,proto3" json:"lowerBound,omitempty"`   // LowerBound is the lower bound of the target
	UpperBound  float64 `protobuf:"fixed64,4,opt,name=upperBound,proto3" json:"upperBound,omitempty"`   // UpperBound is the upper bound of the target
}

func (x *WorkloadRecommendationTarget) Reset() {
	*x = WorkloadRecommendationTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadRecommendationTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadRecommendationTarget) ProtoMessage() {}

func (x *WorkloadRecommendationTarget) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadRecommendationTarget.ProtoReflect.Descriptor instead.
func (*WorkloadRecommendationTarget) Descriptor() ([]byte, []int) {
	return file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescGZIP(), []int{3}
}

func (x *WorkloadRecommendationTarget) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *WorkloadRecommendationTarget) GetTargetValue() float64 {
	if x != nil {
		return x.TargetValue
	}
	return 0
}

func (x *WorkloadRecommendationTarget) GetLowerBound() float64 {
	if x != nil {
		return x.LowerBound
	}
	return 0
}

func (x *WorkloadRecommendationTarget) GetUpperBound() float64 {
	if x != nil {
		return x.UpperBound
	}
	return 0
}

// WorkloadRecommendationsReply represents a list of recommendations
type WorkloadRecommendationsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error           *Error                         `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`                     // Error is the error message if new values cannot be calculated
	Recommendations []*WorkloadRecommendationReply `protobuf:"bytes,2,rep,name=recommendations,proto3" json:"recommendations,omitempty"` // Recommendations is the list of recommendations
}

func (x *WorkloadRecommendationsReply) Reset() {
	*x = WorkloadRecommendationsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadRecommendationsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadRecommendationsReply) ProtoMessage() {}

func (x *WorkloadRecommendationsReply) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadRecommendationsReply.ProtoReflect.Descriptor instead.
func (*WorkloadRecommendationsReply) Descriptor() ([]byte, []int) {
	return file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescGZIP(), []int{4}
}

func (x *WorkloadRecommendationsReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *WorkloadRecommendationsReply) GetRecommendations() []*WorkloadRecommendationReply {
	if x != nil {
		return x.Recommendations
	}
	return nil
}

// WorkloadRecommendationReply represents a recommendation reply
type WorkloadRecommendationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error              *Error                          `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`                                  // Error is the error message if new values cannot be calculated
	Timestamp          *timestamppb.Timestamp          `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                          // Timestamp is the time the values were generated
	TargetReplicas     int32                           `protobuf:"varint,3,opt,name=targetReplicas,proto3" json:"targetReplicas,omitempty"`               // TargetReplicas is the number of replicas the workload should have
	LowerBoundReplicas *int32                          `protobuf:"varint,4,opt,name=lowerBoundReplicas,proto3,oneof" json:"lowerBoundReplicas,omitempty"` // LowerBoundReplicas is the number of replicas based on lowerBound input
	UpperBoundReplicas *int32                          `protobuf:"varint,5,opt,name=upperBoundReplicas,proto3,oneof" json:"upperBoundReplicas,omitempty"` // UpperBoundReplicas is the number of replicas based on upperBound input
	ObservedTargets    []*WorkloadRecommendationTarget `protobuf:"bytes,6,rep,name=observedTargets,proto3" json:"observedTargets,omitempty"`              // ObservedTargets is the list of observed targets (only the targetValue will be set)
	Reason             string                          `protobuf:"bytes,7,opt,name=reason,proto3" json:"reason,omitempty"`                                // Reason explains recommender decision
}

func (x *WorkloadRecommendationReply) Reset() {
	*x = WorkloadRecommendationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkloadRecommendationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkloadRecommendationReply) ProtoMessage() {}

func (x *WorkloadRecommendationReply) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkloadRecommendationReply.ProtoReflect.Descriptor instead.
func (*WorkloadRecommendationReply) Descriptor() ([]byte, []int) {
	return file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescGZIP(), []int{5}
}

func (x *WorkloadRecommendationReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *WorkloadRecommendationReply) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *WorkloadRecommendationReply) GetTargetReplicas() int32 {
	if x != nil {
		return x.TargetReplicas
	}
	return 0
}

func (x *WorkloadRecommendationReply) GetLowerBoundReplicas() int32 {
	if x != nil && x.LowerBoundReplicas != nil {
		return *x.LowerBoundReplicas
	}
	return 0
}

func (x *WorkloadRecommendationReply) GetUpperBoundReplicas() int32 {
	if x != nil && x.UpperBoundReplicas != nil {
		return *x.UpperBoundReplicas
	}
	return 0
}

func (x *WorkloadRecommendationReply) GetObservedTargets() []*WorkloadRecommendationTarget {
	if x != nil {
		return x.ObservedTargets
	}
	return nil
}

func (x *WorkloadRecommendationReply) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto protoreflect.FileDescriptor

var file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDesc = []byte{
	0x0a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74,
	0x61, 0x44, 0x6f, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x2f, 0x76, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x1a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x61, 0x74, 0x61, 0x44, 0x6f, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x76, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x68, 0x72, 0x75, 0x73, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x1e, 0x57,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x59, 0x0a,
	0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63,
	0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x3a, 0x05, 0xba, 0x46, 0x02, 0x20, 0x01, 0x22,
	0xc7, 0x04, 0x0a, 0x1d, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x56, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x42, 0x05, 0xaa, 0x46, 0x02, 0x10, 0x01, 0x52, 0x09,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x12, 0x6a, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x73, 0x42, 0x05, 0xaa, 0x46, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x43, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x56, 0x0a, 0x07, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x12, 0x67, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x53, 0x0a, 0x0d, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x3a, 0x07, 0xba, 0x46, 0x04, 0x08, 0x01, 0x20, 0x01, 0x22, 0x7e, 0x0a, 0x21, 0x57, 0x6f, 0x72,
	0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x27,
	0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x05, 0xaa, 0x46, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x27, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x05, 0xaa, 0x46,
	0x02, 0x10, 0x01, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73,
	0x3a, 0x07, 0xba, 0x46, 0x04, 0x08, 0x01, 0x20, 0x01, 0x22, 0xa4, 0x01, 0x0a, 0x1c, 0x57, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xaa, 0x46, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x70,
	0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x07, 0xba, 0x46, 0x04, 0x08, 0x01, 0x20, 0x01,
	0x22, 0xc9, 0x01, 0x0a, 0x1c, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x3b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x73,
	0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x65,
	0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f,
	0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x05, 0xba, 0x46, 0x02, 0x20, 0x01, 0x22, 0xeb, 0x03, 0x0a,
	0x1b, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x05, 0xaa, 0x46, 0x02, 0x10, 0x01, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2d, 0x0a, 0x0e, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x05, 0xaa, 0x46, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x33, 0x0a, 0x12, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x12, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x88, 0x01, 0x01, 0x12, 0x33,
	0x0a, 0x12, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x12, 0x75, 0x70,
	0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x66, 0x0a, 0x0f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x0f, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x3a, 0x07, 0xba, 0x46, 0x04, 0x08, 0x01, 0x20, 0x01, 0x42, 0x15, 0x0a, 0x13,
	0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x73, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x44, 0x6f, 0x67,
	0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x76,
	0x35, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescOnce sync.Once
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescData = file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDesc
)

func file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescGZIP() []byte {
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescOnce.Do(func() {
		file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescData)
	})
	return file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDescData
}

var file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_goTypes = []interface{}{
	(*WorkloadRecommendationsRequest)(nil),    // 0: datadog.autoscaling.kubernetes.WorkloadRecommendationsRequest
	(*WorkloadRecommendationRequest)(nil),     // 1: datadog.autoscaling.kubernetes.WorkloadRecommendationRequest
	(*WorkloadRecommendationConstraints)(nil), // 2: datadog.autoscaling.kubernetes.WorkloadRecommendationConstraints
	(*WorkloadRecommendationTarget)(nil),      // 3: datadog.autoscaling.kubernetes.WorkloadRecommendationTarget
	(*WorkloadRecommendationsReply)(nil),      // 4: datadog.autoscaling.kubernetes.WorkloadRecommendationsReply
	(*WorkloadRecommendationReply)(nil),       // 5: datadog.autoscaling.kubernetes.WorkloadRecommendationReply
	nil,                                       // 6: datadog.autoscaling.kubernetes.WorkloadRecommendationRequest.SettingsEntry
	(*WorkloadTargetRef)(nil),                 // 7: datadog.autoscaling.kubernetes.WorkloadTargetRef
	(*WorkloadState)(nil),                     // 8: datadog.autoscaling.kubernetes.WorkloadState
	(*Error)(nil),                             // 9: datadog.autoscaling.kubernetes.Error
	(*timestamppb.Timestamp)(nil),             // 10: google.protobuf.Timestamp
	(*structpb.Value)(nil),                    // 11: google.protobuf.Value
}
var file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_depIdxs = []int32{
	1,  // 0: datadog.autoscaling.kubernetes.WorkloadRecommendationsRequest.requests:type_name -> datadog.autoscaling.kubernetes.WorkloadRecommendationRequest
	7,  // 1: datadog.autoscaling.kubernetes.WorkloadRecommendationRequest.targetRef:type_name -> datadog.autoscaling.kubernetes.WorkloadTargetRef
	2,  // 2: datadog.autoscaling.kubernetes.WorkloadRecommendationRequest.constraints:type_name -> datadog.autoscaling.kubernetes.WorkloadRecommendationConstraints
	8,  // 3: datadog.autoscaling.kubernetes.WorkloadRecommendationRequest.state:type_name -> datadog.autoscaling.kubernetes.WorkloadState
	3,  // 4: datadog.autoscaling.kubernetes.WorkloadRecommendationRequest.targets:type_name -> datadog.autoscaling.kubernetes.WorkloadRecommendationTarget
	6,  // 5: datadog.autoscaling.kubernetes.WorkloadRecommendationRequest.settings:type_name -> datadog.autoscaling.kubernetes.WorkloadRecommendationRequest.SettingsEntry
	9,  // 6: datadog.autoscaling.kubernetes.WorkloadRecommendationsReply.error:type_name -> datadog.autoscaling.kubernetes.Error
	5,  // 7: datadog.autoscaling.kubernetes.WorkloadRecommendationsReply.recommendations:type_name -> datadog.autoscaling.kubernetes.WorkloadRecommendationReply
	9,  // 8: datadog.autoscaling.kubernetes.WorkloadRecommendationReply.error:type_name -> datadog.autoscaling.kubernetes.Error
	10, // 9: datadog.autoscaling.kubernetes.WorkloadRecommendationReply.timestamp:type_name -> google.protobuf.Timestamp
	3,  // 10: datadog.autoscaling.kubernetes.WorkloadRecommendationReply.observedTargets:type_name -> datadog.autoscaling.kubernetes.WorkloadRecommendationTarget
	11, // 11: datadog.autoscaling.kubernetes.WorkloadRecommendationRequest.SettingsEntry.value:type_name -> google.protobuf.Value
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() {
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_init()
}
func file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_init() {
	if File_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto != nil {
		return
	}
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadRecommendationsRequest); i {
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
		file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadRecommendationRequest); i {
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
		file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadRecommendationConstraints); i {
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
		file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadRecommendationTarget); i {
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
		file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadRecommendationsReply); i {
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
		file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkloadRecommendationReply); i {
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
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_goTypes,
		DependencyIndexes: file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_depIdxs,
		MessageInfos:      file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_msgTypes,
	}.Build()
	File_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto = out.File
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_rawDesc = nil
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_goTypes = nil
	file_github_com_DataDog_agent_payload_v5_proto_autoscaling_kubernetes_recommender_proto_depIdxs = nil
}
