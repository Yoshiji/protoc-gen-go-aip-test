// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: google/spanner/admin/database/v1/common.proto

package database

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Possible encryption types.
type EncryptionInfo_Type int32

const (
	// Encryption type was not specified, though data at rest remains encrypted.
	EncryptionInfo_TYPE_UNSPECIFIED EncryptionInfo_Type = 0
	// The data is encrypted at rest with a key that is
	// fully managed by Google. No key version or status will be populated.
	// This is the default state.
	EncryptionInfo_GOOGLE_DEFAULT_ENCRYPTION EncryptionInfo_Type = 1
	// The data is encrypted at rest with a key that is
	// managed by the customer. The active version of the key. `kms_key_version`
	// will be populated, and `encryption_status` may be populated.
	EncryptionInfo_CUSTOMER_MANAGED_ENCRYPTION EncryptionInfo_Type = 2
)

// Enum value maps for EncryptionInfo_Type.
var (
	EncryptionInfo_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "GOOGLE_DEFAULT_ENCRYPTION",
		2: "CUSTOMER_MANAGED_ENCRYPTION",
	}
	EncryptionInfo_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":            0,
		"GOOGLE_DEFAULT_ENCRYPTION":   1,
		"CUSTOMER_MANAGED_ENCRYPTION": 2,
	}
)

func (x EncryptionInfo_Type) Enum() *EncryptionInfo_Type {
	p := new(EncryptionInfo_Type)
	*p = x
	return p
}

func (x EncryptionInfo_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncryptionInfo_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_spanner_admin_database_v1_common_proto_enumTypes[0].Descriptor()
}

func (EncryptionInfo_Type) Type() protoreflect.EnumType {
	return &file_google_spanner_admin_database_v1_common_proto_enumTypes[0]
}

func (x EncryptionInfo_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncryptionInfo_Type.Descriptor instead.
func (EncryptionInfo_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_spanner_admin_database_v1_common_proto_rawDescGZIP(), []int{2, 0}
}

// Encapsulates progress related information for a Cloud Spanner long
// running operation.
type OperationProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Percent completion of the operation.
	// Values are between 0 and 100 inclusive.
	ProgressPercent int32 `protobuf:"varint,1,opt,name=progress_percent,json=progressPercent,proto3" json:"progress_percent,omitempty"`
	// Time the request was received.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// If set, the time at which this operation failed or was completed
	// successfully.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *OperationProgress) Reset() {
	*x = OperationProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_admin_database_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationProgress) ProtoMessage() {}

func (x *OperationProgress) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_admin_database_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationProgress.ProtoReflect.Descriptor instead.
func (*OperationProgress) Descriptor() ([]byte, []int) {
	return file_google_spanner_admin_database_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *OperationProgress) GetProgressPercent() int32 {
	if x != nil {
		return x.ProgressPercent
	}
	return 0
}

func (x *OperationProgress) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *OperationProgress) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

// Encryption configuration for a Cloud Spanner database.
type EncryptionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Cloud KMS key to be used for encrypting and decrypting
	// the database. Values are of the form
	// `projects/<project>/locations/<location>/keyRings/<key_ring>/cryptoKeys/<kms_key_name>`.
	KmsKeyName string `protobuf:"bytes,2,opt,name=kms_key_name,json=kmsKeyName,proto3" json:"kms_key_name,omitempty"`
}

func (x *EncryptionConfig) Reset() {
	*x = EncryptionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_admin_database_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionConfig) ProtoMessage() {}

func (x *EncryptionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_admin_database_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionConfig.ProtoReflect.Descriptor instead.
func (*EncryptionConfig) Descriptor() ([]byte, []int) {
	return file_google_spanner_admin_database_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *EncryptionConfig) GetKmsKeyName() string {
	if x != nil {
		return x.KmsKeyName
	}
	return ""
}

// Encryption information for a Cloud Spanner database or backup.
type EncryptionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The type of encryption.
	EncryptionType EncryptionInfo_Type `protobuf:"varint,3,opt,name=encryption_type,json=encryptionType,proto3,enum=google.spanner.admin.database.v1.EncryptionInfo_Type" json:"encryption_type,omitempty"`
	// Output only. If present, the status of a recent encrypt/decrypt call on underlying data
	// for this database or backup. Regardless of status, data is always encrypted
	// at rest.
	EncryptionStatus *status.Status `protobuf:"bytes,4,opt,name=encryption_status,json=encryptionStatus,proto3" json:"encryption_status,omitempty"`
	// Output only. A Cloud KMS key version that is being used to protect the database or
	// backup.
	KmsKeyVersion string `protobuf:"bytes,2,opt,name=kms_key_version,json=kmsKeyVersion,proto3" json:"kms_key_version,omitempty"`
}

func (x *EncryptionInfo) Reset() {
	*x = EncryptionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_spanner_admin_database_v1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionInfo) ProtoMessage() {}

func (x *EncryptionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_spanner_admin_database_v1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionInfo.ProtoReflect.Descriptor instead.
func (*EncryptionInfo) Descriptor() ([]byte, []int) {
	return file_google_spanner_admin_database_v1_common_proto_rawDescGZIP(), []int{2}
}

func (x *EncryptionInfo) GetEncryptionType() EncryptionInfo_Type {
	if x != nil {
		return x.EncryptionType
	}
	return EncryptionInfo_TYPE_UNSPECIFIED
}

func (x *EncryptionInfo) GetEncryptionStatus() *status.Status {
	if x != nil {
		return x.EncryptionStatus
	}
	return nil
}

func (x *EncryptionInfo) GetKmsKeyVersion() string {
	if x != nil {
		return x.KmsKeyVersion
	}
	return ""
}

var File_google_spanner_admin_database_v1_common_proto protoreflect.FileDescriptor

var file_google_spanner_admin_database_v1_common_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x10, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x48, 0x0a, 0x0c,
	0x6b, 0x6d, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x26, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x52, 0x0a, 0x6b, 0x6d, 0x73, 0x4b,
	0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf6, 0x02, 0x0a, 0x0e, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x64, 0x0a, 0x0f, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52,
	0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x45, 0x0a, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x03, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x59, 0x0a, 0x0f, 0x6b, 0x6d, 0x73, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x31, 0xe2, 0x41, 0x01, 0x03, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b,
	0x6d, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x0d, 0x6b, 0x6d, 0x73, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x5c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1d, 0x0a, 0x19, 0x47, 0x4f, 0x4f, 0x47, 0x4c, 0x45, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c,
	0x54, 0x5f, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x1f,
	0x0a, 0x1b, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47,
	0x45, 0x44, 0x5f, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x42,
	0xa4, 0x04, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73,
	0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73,
	0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0xaa, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x26, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x5c, 0x56, 0x31, 0xea, 0x02, 0x2b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x70, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0xea, 0x41, 0x78, 0x0a, 0x21, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x12, 0x53, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x2f, 0x6b, 0x65, 0x79, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x5f, 0x72,
	0x69, 0x6e, 0x67, 0x7d, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x73, 0x2f,
	0x7b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x7d, 0xea, 0x41, 0xa6, 0x01,
	0x0a, 0x28, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x6d, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x7a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x6b, 0x65, 0x79, 0x52, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x6b, 0x65,
	0x79, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x7d, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65,
	0x79, 0x73, 0x2f, 0x7b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x7d, 0x2f,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_spanner_admin_database_v1_common_proto_rawDescOnce sync.Once
	file_google_spanner_admin_database_v1_common_proto_rawDescData = file_google_spanner_admin_database_v1_common_proto_rawDesc
)

func file_google_spanner_admin_database_v1_common_proto_rawDescGZIP() []byte {
	file_google_spanner_admin_database_v1_common_proto_rawDescOnce.Do(func() {
		file_google_spanner_admin_database_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_spanner_admin_database_v1_common_proto_rawDescData)
	})
	return file_google_spanner_admin_database_v1_common_proto_rawDescData
}

var file_google_spanner_admin_database_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_spanner_admin_database_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_spanner_admin_database_v1_common_proto_goTypes = []interface{}{
	(EncryptionInfo_Type)(0),      // 0: google.spanner.admin.database.v1.EncryptionInfo.Type
	(*OperationProgress)(nil),     // 1: google.spanner.admin.database.v1.OperationProgress
	(*EncryptionConfig)(nil),      // 2: google.spanner.admin.database.v1.EncryptionConfig
	(*EncryptionInfo)(nil),        // 3: google.spanner.admin.database.v1.EncryptionInfo
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*status.Status)(nil),         // 5: google.rpc.Status
}
var file_google_spanner_admin_database_v1_common_proto_depIdxs = []int32{
	4, // 0: google.spanner.admin.database.v1.OperationProgress.start_time:type_name -> google.protobuf.Timestamp
	4, // 1: google.spanner.admin.database.v1.OperationProgress.end_time:type_name -> google.protobuf.Timestamp
	0, // 2: google.spanner.admin.database.v1.EncryptionInfo.encryption_type:type_name -> google.spanner.admin.database.v1.EncryptionInfo.Type
	5, // 3: google.spanner.admin.database.v1.EncryptionInfo.encryption_status:type_name -> google.rpc.Status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_spanner_admin_database_v1_common_proto_init() }
func file_google_spanner_admin_database_v1_common_proto_init() {
	if File_google_spanner_admin_database_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_spanner_admin_database_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationProgress); i {
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
		file_google_spanner_admin_database_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptionConfig); i {
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
		file_google_spanner_admin_database_v1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptionInfo); i {
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
			RawDescriptor: file_google_spanner_admin_database_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_spanner_admin_database_v1_common_proto_goTypes,
		DependencyIndexes: file_google_spanner_admin_database_v1_common_proto_depIdxs,
		EnumInfos:         file_google_spanner_admin_database_v1_common_proto_enumTypes,
		MessageInfos:      file_google_spanner_admin_database_v1_common_proto_msgTypes,
	}.Build()
	File_google_spanner_admin_database_v1_common_proto = out.File
	file_google_spanner_admin_database_v1_common_proto_rawDesc = nil
	file_google_spanner_admin_database_v1_common_proto_goTypes = nil
	file_google_spanner_admin_database_v1_common_proto_depIdxs = nil
}
