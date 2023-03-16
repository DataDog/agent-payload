// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: proto/contimage/contimage.proto

package contimage

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// ContainerImagePayload represents the main container image payload
type ContainerImagePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string            `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`     // version represents the ContainerImagePayload message version
	Host    string            `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`           // host contains the hostname
	Source  *string           `protobuf:"bytes,4,opt,name=source,proto3,oneof" json:"source,omitempty"` // use to know the source of the message: agent, other
	Images  []*ContainerImage `protobuf:"bytes,3,rep,name=images,proto3" json:"images,omitempty"`       // Images metadata
}

func (x *ContainerImagePayload) Reset() {
	*x = ContainerImagePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contimage_contimage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerImagePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerImagePayload) ProtoMessage() {}

func (x *ContainerImagePayload) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contimage_contimage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerImagePayload.ProtoReflect.Descriptor instead.
func (*ContainerImagePayload) Descriptor() ([]byte, []int) {
	return file_proto_contimage_contimage_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerImagePayload) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ContainerImagePayload) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ContainerImagePayload) GetSource() string {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return ""
}

func (x *ContainerImagePayload) GetImages() []*ContainerImage {
	if x != nil {
		return x.Images
	}
	return nil
}

// ContainerImage contains the details of a container image
type ContainerImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DdTags      []string                              `protobuf:"bytes,12,rep,name=dd_tags,json=ddTags,proto3" json:"dd_tags,omitempty"`
	Name        string                                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Registry    string                                `protobuf:"bytes,3,opt,name=registry,proto3" json:"registry,omitempty"`
	ShortName   string                                `protobuf:"bytes,4,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	RepoTags    []string                              `protobuf:"bytes,5,rep,name=repo_tags,json=repoTags,proto3" json:"repo_tags,omitempty"`
	Digest      string                                `protobuf:"bytes,6,opt,name=digest,proto3" json:"digest,omitempty"`
	Size        int64                                 `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	RepoDigests []string                              `protobuf:"bytes,8,rep,name=repo_digests,json=repoDigests,proto3" json:"repo_digests,omitempty"`
	Os          *ContainerImage_OperatingSystem       `protobuf:"bytes,9,opt,name=os,proto3" json:"os,omitempty"`
	Layers      []*ContainerImage_ContainerImageLayer `protobuf:"bytes,10,rep,name=layers,proto3" json:"layers,omitempty"`
	BuiltAt     *timestamp.Timestamp                  `protobuf:"bytes,11,opt,name=builtAt,proto3" json:"builtAt,omitempty"`
}

func (x *ContainerImage) Reset() {
	*x = ContainerImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contimage_contimage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerImage) ProtoMessage() {}

func (x *ContainerImage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contimage_contimage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerImage.ProtoReflect.Descriptor instead.
func (*ContainerImage) Descriptor() ([]byte, []int) {
	return file_proto_contimage_contimage_proto_rawDescGZIP(), []int{1}
}

func (x *ContainerImage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContainerImage) GetDdTags() []string {
	if x != nil {
		return x.DdTags
	}
	return nil
}

func (x *ContainerImage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContainerImage) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

func (x *ContainerImage) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *ContainerImage) GetRepoTags() []string {
	if x != nil {
		return x.RepoTags
	}
	return nil
}

func (x *ContainerImage) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *ContainerImage) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ContainerImage) GetRepoDigests() []string {
	if x != nil {
		return x.RepoDigests
	}
	return nil
}

func (x *ContainerImage) GetOs() *ContainerImage_OperatingSystem {
	if x != nil {
		return x.Os
	}
	return nil
}

func (x *ContainerImage) GetLayers() []*ContainerImage_ContainerImageLayer {
	if x != nil {
		return x.Layers
	}
	return nil
}

func (x *ContainerImage) GetBuiltAt() *timestamp.Timestamp {
	if x != nil {
		return x.BuiltAt
	}
	return nil
}

// platform info
type ContainerImage_OperatingSystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version      string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Architecture string `protobuf:"bytes,3,opt,name=architecture,proto3" json:"architecture,omitempty"`
}

func (x *ContainerImage_OperatingSystem) Reset() {
	*x = ContainerImage_OperatingSystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contimage_contimage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerImage_OperatingSystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerImage_OperatingSystem) ProtoMessage() {}

func (x *ContainerImage_OperatingSystem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contimage_contimage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerImage_OperatingSystem.ProtoReflect.Descriptor instead.
func (*ContainerImage_OperatingSystem) Descriptor() ([]byte, []int) {
	return file_proto_contimage_contimage_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ContainerImage_OperatingSystem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContainerImage_OperatingSystem) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ContainerImage_OperatingSystem) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

type ContainerImage_ContainerImageLayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls      []string                                    `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	MediaType string                                      `protobuf:"bytes,2,opt,name=mediaType,proto3" json:"mediaType,omitempty"`
	Digest    string                                      `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	Size      int64                                       `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	History   *ContainerImage_ContainerImageLayer_History `protobuf:"bytes,5,opt,name=history,proto3" json:"history,omitempty"`
}

func (x *ContainerImage_ContainerImageLayer) Reset() {
	*x = ContainerImage_ContainerImageLayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contimage_contimage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerImage_ContainerImageLayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerImage_ContainerImageLayer) ProtoMessage() {}

func (x *ContainerImage_ContainerImageLayer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contimage_contimage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerImage_ContainerImageLayer.ProtoReflect.Descriptor instead.
func (*ContainerImage_ContainerImageLayer) Descriptor() ([]byte, []int) {
	return file_proto_contimage_contimage_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ContainerImage_ContainerImageLayer) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *ContainerImage_ContainerImageLayer) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *ContainerImage_ContainerImageLayer) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *ContainerImage_ContainerImageLayer) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ContainerImage_ContainerImageLayer) GetHistory() *ContainerImage_ContainerImageLayer_History {
	if x != nil {
		return x.History
	}
	return nil
}

type ContainerImage_ContainerImageLayer_History struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created    *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	CreatedBy  string               `protobuf:"bytes,2,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	Author     string               `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Comment    string               `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	EmptyLayer bool                 `protobuf:"varint,5,opt,name=emptyLayer,proto3" json:"emptyLayer,omitempty"`
}

func (x *ContainerImage_ContainerImageLayer_History) Reset() {
	*x = ContainerImage_ContainerImageLayer_History{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_contimage_contimage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerImage_ContainerImageLayer_History) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerImage_ContainerImageLayer_History) ProtoMessage() {}

func (x *ContainerImage_ContainerImageLayer_History) ProtoReflect() protoreflect.Message {
	mi := &file_proto_contimage_contimage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerImage_ContainerImageLayer_History.ProtoReflect.Descriptor instead.
func (*ContainerImage_ContainerImageLayer_History) Descriptor() ([]byte, []int) {
	return file_proto_contimage_contimage_proto_rawDescGZIP(), []int{1, 1, 0}
}

func (x *ContainerImage_ContainerImageLayer_History) GetCreated() *timestamp.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ContainerImage_ContainerImageLayer_History) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *ContainerImage_ContainerImageLayer_History) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *ContainerImage_ContainerImageLayer_History) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *ContainerImage_ContainerImageLayer_History) GetEmptyLayer() bool {
	if x != nil {
		return x.EmptyLayer
	}
	return false
}

var File_proto_contimage_contimage_proto protoreflect.FileDescriptor

var file_proto_contimage_contimage_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x06, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x64, 0x6f, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x06, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0xa2, 0x07, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x64, 0x64, 0x54, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x70, 0x6f, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x70, 0x6f, 0x54, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6f,
	0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x4d, 0x0a, 0x06, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x64, 0x6f, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x06, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x62, 0x75, 0x69,
	0x6c, 0x74, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x41, 0x74, 0x1a,
	0x63, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x1a, 0xfe, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x57, 0x0a, 0x07, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x1a, 0xaf, 0x01, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x44, 0x6f, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2d, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x76, 0x35, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_contimage_contimage_proto_rawDescOnce sync.Once
	file_proto_contimage_contimage_proto_rawDescData = file_proto_contimage_contimage_proto_rawDesc
)

func file_proto_contimage_contimage_proto_rawDescGZIP() []byte {
	file_proto_contimage_contimage_proto_rawDescOnce.Do(func() {
		file_proto_contimage_contimage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_contimage_contimage_proto_rawDescData)
	})
	return file_proto_contimage_contimage_proto_rawDescData
}

var file_proto_contimage_contimage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_contimage_contimage_proto_goTypes = []interface{}{
	(*ContainerImagePayload)(nil),                      // 0: datadog.contimage.ContainerImagePayload
	(*ContainerImage)(nil),                             // 1: datadog.contimage.ContainerImage
	(*ContainerImage_OperatingSystem)(nil),             // 2: datadog.contimage.ContainerImage.OperatingSystem
	(*ContainerImage_ContainerImageLayer)(nil),         // 3: datadog.contimage.ContainerImage.ContainerImageLayer
	(*ContainerImage_ContainerImageLayer_History)(nil), // 4: datadog.contimage.ContainerImage.ContainerImageLayer.History
	(*timestamp.Timestamp)(nil),                        // 5: google.protobuf.Timestamp
}
var file_proto_contimage_contimage_proto_depIdxs = []int32{
	1, // 0: datadog.contimage.ContainerImagePayload.images:type_name -> datadog.contimage.ContainerImage
	2, // 1: datadog.contimage.ContainerImage.os:type_name -> datadog.contimage.ContainerImage.OperatingSystem
	3, // 2: datadog.contimage.ContainerImage.layers:type_name -> datadog.contimage.ContainerImage.ContainerImageLayer
	5, // 3: datadog.contimage.ContainerImage.builtAt:type_name -> google.protobuf.Timestamp
	4, // 4: datadog.contimage.ContainerImage.ContainerImageLayer.history:type_name -> datadog.contimage.ContainerImage.ContainerImageLayer.History
	5, // 5: datadog.contimage.ContainerImage.ContainerImageLayer.History.created:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_contimage_contimage_proto_init() }
func file_proto_contimage_contimage_proto_init() {
	if File_proto_contimage_contimage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_contimage_contimage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerImagePayload); i {
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
		file_proto_contimage_contimage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerImage); i {
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
		file_proto_contimage_contimage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerImage_OperatingSystem); i {
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
		file_proto_contimage_contimage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerImage_ContainerImageLayer); i {
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
		file_proto_contimage_contimage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerImage_ContainerImageLayer_History); i {
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
	file_proto_contimage_contimage_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_contimage_contimage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_contimage_contimage_proto_goTypes,
		DependencyIndexes: file_proto_contimage_contimage_proto_depIdxs,
		MessageInfos:      file_proto_contimage_contimage_proto_msgTypes,
	}.Build()
	File_proto_contimage_contimage_proto = out.File
	file_proto_contimage_contimage_proto_rawDesc = nil
	file_proto_contimage_contimage_proto_goTypes = nil
	file_proto_contimage_contimage_proto_depIdxs = nil
}
