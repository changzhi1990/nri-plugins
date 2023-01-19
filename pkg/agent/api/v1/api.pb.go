//
//Copyright 2019 Intel Corporation
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: pkg/agent/api/v1/api.proto

package v1

import (
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

type GetNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNodeRequest) Reset() {
	*x = GetNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_api_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeRequest) ProtoMessage() {}

func (x *GetNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_api_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeRequest.ProtoReflect.Descriptor instead.
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_agent_api_v1_api_proto_rawDescGZIP(), []int{0}
}

type GetNodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node string `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *GetNodeReply) Reset() {
	*x = GetNodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_api_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeReply) ProtoMessage() {}

func (x *GetNodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_api_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeReply.ProtoReflect.Descriptor instead.
func (*GetNodeReply) Descriptor() ([]byte, []int) {
	return file_pkg_agent_api_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetNodeReply) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

// JsonPatch holds on JSON patch
type JsonPatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	Path  string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *JsonPatch) Reset() {
	*x = JsonPatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_api_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonPatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonPatch) ProtoMessage() {}

func (x *JsonPatch) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_api_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonPatch.ProtoReflect.Descriptor instead.
func (*JsonPatch) Descriptor() ([]byte, []int) {
	return file_pkg_agent_api_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *JsonPatch) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *JsonPatch) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *JsonPatch) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PatchNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of JSON patches to apply on the node
	Patches []*JsonPatch `protobuf:"bytes,1,rep,name=patches,proto3" json:"patches,omitempty"`
}

func (x *PatchNodeRequest) Reset() {
	*x = PatchNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_api_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchNodeRequest) ProtoMessage() {}

func (x *PatchNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_api_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchNodeRequest.ProtoReflect.Descriptor instead.
func (*PatchNodeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_agent_api_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *PatchNodeRequest) GetPatches() []*JsonPatch {
	if x != nil {
		return x.Patches
	}
	return nil
}

type PatchNodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PatchNodeReply) Reset() {
	*x = PatchNodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_api_v1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchNodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchNodeReply) ProtoMessage() {}

func (x *PatchNodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_api_v1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchNodeReply.ProtoReflect.Descriptor instead.
func (*PatchNodeReply) Descriptor() ([]byte, []int) {
	return file_pkg_agent_api_v1_api_proto_rawDescGZIP(), []int{4}
}

type UpdateNodeCapacityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name-value map of status.capacity to update
	Capacities map[string]string `protobuf:"bytes,1,rep,name=capacities,proto3" json:"capacities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UpdateNodeCapacityRequest) Reset() {
	*x = UpdateNodeCapacityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_api_v1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeCapacityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeCapacityRequest) ProtoMessage() {}

func (x *UpdateNodeCapacityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_api_v1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeCapacityRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeCapacityRequest) Descriptor() ([]byte, []int) {
	return file_pkg_agent_api_v1_api_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateNodeCapacityRequest) GetCapacities() map[string]string {
	if x != nil {
		return x.Capacities
	}
	return nil
}

type UpdateNodeCapacityReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateNodeCapacityReply) Reset() {
	*x = UpdateNodeCapacityReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_api_v1_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeCapacityReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeCapacityReply) ProtoMessage() {}

func (x *UpdateNodeCapacityReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_api_v1_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeCapacityReply.ProtoReflect.Descriptor instead.
func (*UpdateNodeCapacityReply) Descriptor() ([]byte, []int) {
	return file_pkg_agent_api_v1_api_proto_rawDescGZIP(), []int{6}
}

type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_api_v1_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_api_v1_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_pkg_agent_api_v1_api_proto_rawDescGZIP(), []int{7}
}

type GetConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string            `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Config   map[string]string `protobuf:"bytes,2,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetConfigReply) Reset() {
	*x = GetConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_api_v1_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigReply) ProtoMessage() {}

func (x *GetConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_api_v1_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigReply.ProtoReflect.Descriptor instead.
func (*GetConfigReply) Descriptor() ([]byte, []int) {
	return file_pkg_agent_api_v1_api_proto_rawDescGZIP(), []int{8}
}

func (x *GetConfigReply) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *GetConfigReply) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_api_v1_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_api_v1_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_pkg_agent_api_v1_api_proto_rawDescGZIP(), []int{9}
}

type HealthCheckReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *HealthCheckReply) Reset() {
	*x = HealthCheckReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_agent_api_v1_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckReply) ProtoMessage() {}

func (x *HealthCheckReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_agent_api_v1_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckReply.ProtoReflect.Descriptor instead.
func (*HealthCheckReply) Descriptor() ([]byte, []int) {
	return file_pkg_agent_api_v1_api_proto_rawDescGZIP(), []int{10}
}

func (x *HealthCheckReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_agent_api_v1_api_proto protoreflect.FileDescriptor

var file_pkg_agent_api_v1_api_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x22, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x45, 0x0a, 0x09, 0x4a, 0x73, 0x6f, 0x6e, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3b, 0x0a,
	0x10, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x07, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xa9, 0x01, 0x0a,
	0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0a, 0x63, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x43, 0x61, 0x70,
	0x61, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa0, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a,
	0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x14, 0x0a, 0x12, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x28, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xbf, 0x02, 0x0a, 0x05, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x09, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x52, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_agent_api_v1_api_proto_rawDescOnce sync.Once
	file_pkg_agent_api_v1_api_proto_rawDescData = file_pkg_agent_api_v1_api_proto_rawDesc
)

func file_pkg_agent_api_v1_api_proto_rawDescGZIP() []byte {
	file_pkg_agent_api_v1_api_proto_rawDescOnce.Do(func() {
		file_pkg_agent_api_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_agent_api_v1_api_proto_rawDescData)
	})
	return file_pkg_agent_api_v1_api_proto_rawDescData
}

var file_pkg_agent_api_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pkg_agent_api_v1_api_proto_goTypes = []interface{}{
	(*GetNodeRequest)(nil),            // 0: v1.GetNodeRequest
	(*GetNodeReply)(nil),              // 1: v1.GetNodeReply
	(*JsonPatch)(nil),                 // 2: v1.JsonPatch
	(*PatchNodeRequest)(nil),          // 3: v1.PatchNodeRequest
	(*PatchNodeReply)(nil),            // 4: v1.PatchNodeReply
	(*UpdateNodeCapacityRequest)(nil), // 5: v1.UpdateNodeCapacityRequest
	(*UpdateNodeCapacityReply)(nil),   // 6: v1.UpdateNodeCapacityReply
	(*GetConfigRequest)(nil),          // 7: v1.GetConfigRequest
	(*GetConfigReply)(nil),            // 8: v1.GetConfigReply
	(*HealthCheckRequest)(nil),        // 9: v1.HealthCheckRequest
	(*HealthCheckReply)(nil),          // 10: v1.HealthCheckReply
	nil,                               // 11: v1.UpdateNodeCapacityRequest.CapacitiesEntry
	nil,                               // 12: v1.GetConfigReply.ConfigEntry
}
var file_pkg_agent_api_v1_api_proto_depIdxs = []int32{
	2,  // 0: v1.PatchNodeRequest.patches:type_name -> v1.JsonPatch
	11, // 1: v1.UpdateNodeCapacityRequest.capacities:type_name -> v1.UpdateNodeCapacityRequest.CapacitiesEntry
	12, // 2: v1.GetConfigReply.config:type_name -> v1.GetConfigReply.ConfigEntry
	0,  // 3: v1.Agent.GetNode:input_type -> v1.GetNodeRequest
	3,  // 4: v1.Agent.PatchNode:input_type -> v1.PatchNodeRequest
	5,  // 5: v1.Agent.UpdateNodeCapacity:input_type -> v1.UpdateNodeCapacityRequest
	7,  // 6: v1.Agent.GetConfig:input_type -> v1.GetConfigRequest
	9,  // 7: v1.Agent.HealthCheck:input_type -> v1.HealthCheckRequest
	1,  // 8: v1.Agent.GetNode:output_type -> v1.GetNodeReply
	4,  // 9: v1.Agent.PatchNode:output_type -> v1.PatchNodeReply
	6,  // 10: v1.Agent.UpdateNodeCapacity:output_type -> v1.UpdateNodeCapacityReply
	8,  // 11: v1.Agent.GetConfig:output_type -> v1.GetConfigReply
	10, // 12: v1.Agent.HealthCheck:output_type -> v1.HealthCheckReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_agent_api_v1_api_proto_init() }
func file_pkg_agent_api_v1_api_proto_init() {
	if File_pkg_agent_api_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_agent_api_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeRequest); i {
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
		file_pkg_agent_api_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeReply); i {
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
		file_pkg_agent_api_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonPatch); i {
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
		file_pkg_agent_api_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchNodeRequest); i {
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
		file_pkg_agent_api_v1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchNodeReply); i {
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
		file_pkg_agent_api_v1_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeCapacityRequest); i {
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
		file_pkg_agent_api_v1_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeCapacityReply); i {
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
		file_pkg_agent_api_v1_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigRequest); i {
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
		file_pkg_agent_api_v1_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigReply); i {
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
		file_pkg_agent_api_v1_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_pkg_agent_api_v1_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckReply); i {
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
			RawDescriptor: file_pkg_agent_api_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_agent_api_v1_api_proto_goTypes,
		DependencyIndexes: file_pkg_agent_api_v1_api_proto_depIdxs,
		MessageInfos:      file_pkg_agent_api_v1_api_proto_msgTypes,
	}.Build()
	File_pkg_agent_api_v1_api_proto = out.File
	file_pkg_agent_api_v1_api_proto_rawDesc = nil
	file_pkg_agent_api_v1_api_proto_goTypes = nil
	file_pkg_agent_api_v1_api_proto_depIdxs = nil
}