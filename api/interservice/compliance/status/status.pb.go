// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: interservice/compliance/status/status.proto

package status

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type MigrationStatus_Status int32

const (
	MigrationStatus_UNKNOWN  MigrationStatus_Status = 0
	MigrationStatus_RUNNING  MigrationStatus_Status = 1
	MigrationStatus_FINISHED MigrationStatus_Status = 2
	MigrationStatus_FAILED   MigrationStatus_Status = 3
	MigrationStatus_SKIPPED  MigrationStatus_Status = 4
)

// Enum value maps for MigrationStatus_Status.
var (
	MigrationStatus_Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "RUNNING",
		2: "FINISHED",
		3: "FAILED",
		4: "SKIPPED",
	}
	MigrationStatus_Status_value = map[string]int32{
		"UNKNOWN":  0,
		"RUNNING":  1,
		"FINISHED": 2,
		"FAILED":   3,
		"SKIPPED":  4,
	}
)

func (x MigrationStatus_Status) Enum() *MigrationStatus_Status {
	p := new(MigrationStatus_Status)
	*p = x
	return p
}

func (x MigrationStatus_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MigrationStatus_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_interservice_compliance_status_status_proto_enumTypes[0].Descriptor()
}

func (MigrationStatus_Status) Type() protoreflect.EnumType {
	return &file_interservice_compliance_status_status_proto_enumTypes[0]
}

func (x MigrationStatus_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MigrationStatus_Status.Descriptor instead.
func (MigrationStatus_Status) EnumDescriptor() ([]byte, []int) {
	return file_interservice_compliance_status_status_proto_rawDescGZIP(), []int{0, 0}
}

type ControlIndexMigrationStatus_Status int32

const (
	ControlIndexMigrationStatus_NOTCONFIGURED ControlIndexMigrationStatus_Status = 0
	ControlIndexMigrationStatus_NOTSTARTED    ControlIndexMigrationStatus_Status = 1
	ControlIndexMigrationStatus_INPROGRESS    ControlIndexMigrationStatus_Status = 2
	ControlIndexMigrationStatus_COMPLETED     ControlIndexMigrationStatus_Status = 3
)

// Enum value maps for ControlIndexMigrationStatus_Status.
var (
	ControlIndexMigrationStatus_Status_name = map[int32]string{
		0: "NOTCONFIGURED",
		1: "NOTSTARTED",
		2: "INPROGRESS",
		3: "COMPLETED",
	}
	ControlIndexMigrationStatus_Status_value = map[string]int32{
		"NOTCONFIGURED": 0,
		"NOTSTARTED":    1,
		"INPROGRESS":    2,
		"COMPLETED":     3,
	}
)

func (x ControlIndexMigrationStatus_Status) Enum() *ControlIndexMigrationStatus_Status {
	p := new(ControlIndexMigrationStatus_Status)
	*p = x
	return p
}

func (x ControlIndexMigrationStatus_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ControlIndexMigrationStatus_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_interservice_compliance_status_status_proto_enumTypes[1].Descriptor()
}

func (ControlIndexMigrationStatus_Status) Type() protoreflect.EnumType {
	return &file_interservice_compliance_status_status_proto_enumTypes[1]
}

func (x ControlIndexMigrationStatus_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ControlIndexMigrationStatus_Status.Descriptor instead.
func (ControlIndexMigrationStatus_Status) EnumDescriptor() ([]byte, []int) {
	return file_interservice_compliance_status_status_proto_rawDescGZIP(), []int{1, 0}
}

// MigrationStatus message
type MigrationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     int64                  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty" toml:"total,omitempty" mapstructure:"total,omitempty"`
	Completed int64                  `protobuf:"varint,2,opt,name=completed,proto3" json:"completed,omitempty" toml:"completed,omitempty" mapstructure:"completed,omitempty"`
	Logs      []*LogEntry            `protobuf:"bytes,3,rep,name=logs,proto3" json:"logs,omitempty" toml:"logs,omitempty" mapstructure:"logs,omitempty"`
	Status    MigrationStatus_Status `protobuf:"varint,4,opt,name=status,proto3,enum=chef.automate.domain.compliance.status.MigrationStatus_Status" json:"status,omitempty" toml:"status,omitempty" mapstructure:"status,omitempty"`
}

func (x *MigrationStatus) Reset() {
	*x = MigrationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_compliance_status_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrationStatus) ProtoMessage() {}

func (x *MigrationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_compliance_status_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrationStatus.ProtoReflect.Descriptor instead.
func (*MigrationStatus) Descriptor() ([]byte, []int) {
	return file_interservice_compliance_status_status_proto_rawDescGZIP(), []int{0}
}

func (x *MigrationStatus) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MigrationStatus) GetCompleted() int64 {
	if x != nil {
		return x.Completed
	}
	return 0
}

func (x *MigrationStatus) GetLogs() []*LogEntry {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *MigrationStatus) GetStatus() MigrationStatus_Status {
	if x != nil {
		return x.Status
	}
	return MigrationStatus_UNKNOWN
}

type ControlIndexMigrationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ControlIndexMigrationStatus_Status `protobuf:"varint,1,opt,name=status,proto3,enum=chef.automate.domain.compliance.status.ControlIndexMigrationStatus_Status" json:"status,omitempty" toml:"status,omitempty" mapstructure:"status,omitempty"`
}

func (x *ControlIndexMigrationStatus) Reset() {
	*x = ControlIndexMigrationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_compliance_status_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlIndexMigrationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlIndexMigrationStatus) ProtoMessage() {}

func (x *ControlIndexMigrationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_compliance_status_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlIndexMigrationStatus.ProtoReflect.Descriptor instead.
func (*ControlIndexMigrationStatus) Descriptor() ([]byte, []int) {
	return file_interservice_compliance_status_status_proto_rawDescGZIP(), []int{1}
}

func (x *ControlIndexMigrationStatus) GetStatus() ControlIndexMigrationStatus_Status {
	if x != nil {
		return x.Status
	}
	return ControlIndexMigrationStatus_NOTCONFIGURED
}

// label is used to differentiate between multiple migrations done by one service
type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label     string                 `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty" toml:"label,omitempty" mapstructure:"label,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty" toml:"timestamp,omitempty" mapstructure:"timestamp,omitempty"`
	Text      string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty" toml:"text,omitempty" mapstructure:"text,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_compliance_status_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_compliance_status_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_interservice_compliance_status_status_proto_rawDescGZIP(), []int{2}
}

func (x *LogEntry) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *LogEntry) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogEntry) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_interservice_compliance_status_status_proto protoreflect.FileDescriptor

var file_interservice_compliance_status_status_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xae, 0x02, 0x0a, 0x0f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6c, 0x6f, 0x67,
	0x73, 0x12, 0x56, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x3e, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x49, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x4b, 0x49, 0x50, 0x50,
	0x45, 0x44, 0x10, 0x04, 0x22, 0xcd, 0x01, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x62, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x4a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55,
	0x52, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x54, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x50, 0x52, 0x4f, 0x47, 0x52,
	0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54,
	0x45, 0x44, 0x10, 0x03, 0x22, 0x6e, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x32, 0x83, 0x02, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x67, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x37,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x1e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x43, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_interservice_compliance_status_status_proto_rawDescOnce sync.Once
	file_interservice_compliance_status_status_proto_rawDescData = file_interservice_compliance_status_status_proto_rawDesc
)

func file_interservice_compliance_status_status_proto_rawDescGZIP() []byte {
	file_interservice_compliance_status_status_proto_rawDescOnce.Do(func() {
		file_interservice_compliance_status_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_compliance_status_status_proto_rawDescData)
	})
	return file_interservice_compliance_status_status_proto_rawDescData
}

var file_interservice_compliance_status_status_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_interservice_compliance_status_status_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_interservice_compliance_status_status_proto_goTypes = []interface{}{
	(MigrationStatus_Status)(0),             // 0: chef.automate.domain.compliance.status.MigrationStatus.Status
	(ControlIndexMigrationStatus_Status)(0), // 1: chef.automate.domain.compliance.status.ControlIndexMigrationStatus.Status
	(*MigrationStatus)(nil),                 // 2: chef.automate.domain.compliance.status.MigrationStatus
	(*ControlIndexMigrationStatus)(nil),     // 3: chef.automate.domain.compliance.status.ControlIndexMigrationStatus
	(*LogEntry)(nil),                        // 4: chef.automate.domain.compliance.status.LogEntry
	(*timestamppb.Timestamp)(nil),           // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                   // 6: google.protobuf.Empty
}
var file_interservice_compliance_status_status_proto_depIdxs = []int32{
	4, // 0: chef.automate.domain.compliance.status.MigrationStatus.logs:type_name -> chef.automate.domain.compliance.status.LogEntry
	0, // 1: chef.automate.domain.compliance.status.MigrationStatus.status:type_name -> chef.automate.domain.compliance.status.MigrationStatus.Status
	1, // 2: chef.automate.domain.compliance.status.ControlIndexMigrationStatus.status:type_name -> chef.automate.domain.compliance.status.ControlIndexMigrationStatus.Status
	5, // 3: chef.automate.domain.compliance.status.LogEntry.timestamp:type_name -> google.protobuf.Timestamp
	6, // 4: chef.automate.domain.compliance.status.ComplianceStatusService.GetMigrationStatus:input_type -> google.protobuf.Empty
	6, // 5: chef.automate.domain.compliance.status.ComplianceStatusService.GetControlIndexMigrationStatus:input_type -> google.protobuf.Empty
	2, // 6: chef.automate.domain.compliance.status.ComplianceStatusService.GetMigrationStatus:output_type -> chef.automate.domain.compliance.status.MigrationStatus
	3, // 7: chef.automate.domain.compliance.status.ComplianceStatusService.GetControlIndexMigrationStatus:output_type -> chef.automate.domain.compliance.status.ControlIndexMigrationStatus
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_interservice_compliance_status_status_proto_init() }
func file_interservice_compliance_status_status_proto_init() {
	if File_interservice_compliance_status_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_compliance_status_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrationStatus); i {
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
		file_interservice_compliance_status_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlIndexMigrationStatus); i {
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
		file_interservice_compliance_status_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
			RawDescriptor: file_interservice_compliance_status_status_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interservice_compliance_status_status_proto_goTypes,
		DependencyIndexes: file_interservice_compliance_status_status_proto_depIdxs,
		EnumInfos:         file_interservice_compliance_status_status_proto_enumTypes,
		MessageInfos:      file_interservice_compliance_status_status_proto_msgTypes,
	}.Build()
	File_interservice_compliance_status_status_proto = out.File
	file_interservice_compliance_status_status_proto_rawDesc = nil
	file_interservice_compliance_status_status_proto_goTypes = nil
	file_interservice_compliance_status_status_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ComplianceStatusServiceClient is the client API for ComplianceStatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComplianceStatusServiceClient interface {
	GetMigrationStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MigrationStatus, error)
	GetControlIndexMigrationStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ControlIndexMigrationStatus, error)
}

type complianceStatusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComplianceStatusServiceClient(cc grpc.ClientConnInterface) ComplianceStatusServiceClient {
	return &complianceStatusServiceClient{cc}
}

func (c *complianceStatusServiceClient) GetMigrationStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MigrationStatus, error) {
	out := new(MigrationStatus)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.status.ComplianceStatusService/GetMigrationStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complianceStatusServiceClient) GetControlIndexMigrationStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ControlIndexMigrationStatus, error) {
	out := new(ControlIndexMigrationStatus)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.status.ComplianceStatusService/GetControlIndexMigrationStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComplianceStatusServiceServer is the server API for ComplianceStatusService service.
type ComplianceStatusServiceServer interface {
	GetMigrationStatus(context.Context, *emptypb.Empty) (*MigrationStatus, error)
	GetControlIndexMigrationStatus(context.Context, *emptypb.Empty) (*ControlIndexMigrationStatus, error)
}

// UnimplementedComplianceStatusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedComplianceStatusServiceServer struct {
}

func (*UnimplementedComplianceStatusServiceServer) GetMigrationStatus(context.Context, *emptypb.Empty) (*MigrationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMigrationStatus not implemented")
}
func (*UnimplementedComplianceStatusServiceServer) GetControlIndexMigrationStatus(context.Context, *emptypb.Empty) (*ControlIndexMigrationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetControlIndexMigrationStatus not implemented")
}

func RegisterComplianceStatusServiceServer(s *grpc.Server, srv ComplianceStatusServiceServer) {
	s.RegisterService(&_ComplianceStatusService_serviceDesc, srv)
}

func _ComplianceStatusService_GetMigrationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceStatusServiceServer).GetMigrationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.status.ComplianceStatusService/GetMigrationStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceStatusServiceServer).GetMigrationStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComplianceStatusService_GetControlIndexMigrationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplianceStatusServiceServer).GetControlIndexMigrationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.status.ComplianceStatusService/GetControlIndexMigrationStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplianceStatusServiceServer).GetControlIndexMigrationStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ComplianceStatusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.compliance.status.ComplianceStatusService",
	HandlerType: (*ComplianceStatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMigrationStatus",
			Handler:    _ComplianceStatusService_GetMigrationStatus_Handler,
		},
		{
			MethodName: "GetControlIndexMigrationStatus",
			Handler:    _ComplianceStatusService_GetControlIndexMigrationStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interservice/compliance/status/status.proto",
}
