// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: automate-gateway/api/event_feed/event_feed.proto

package event_feed

import (
	context "context"
	request "github.com/chef/automate/components/automate-gateway/api/event_feed/request"
	response "github.com/chef/automate/components/automate-gateway/api/event_feed/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_automate_gateway_api_event_feed_event_feed_proto protoreflect.FileDescriptor

var file_automate_gateway_api_event_feed_event_feed_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65,
	0x64, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64,
	0x1a, 0x33, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65,
	0x64, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x34, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65,
	0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x35, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xac, 0x06, 0x0a, 0x09, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x46, 0x65, 0x65, 0x64, 0x12, 0xb4, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x46, 0x65, 0x65, 0x64, 0x12, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x2d, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x42, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x66, 0x65, 0x65, 0x64, 0x8a, 0xb5, 0x18, 0x0e, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x8a, 0xb5, 0x18, 0x13, 0x12, 0x11, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xcd,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x32,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x30, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x8a, 0xb5, 0x18, 0x0e, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x8a, 0xb5, 0x18, 0x13, 0x12, 0x11, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xcd,
	0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x32,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x30, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x8a, 0xb5, 0x18, 0x0e, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x8a, 0xb5, 0x18, 0x13, 0x12, 0x11, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xc7,
	0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x32, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x33, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0x45, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x30, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x8a,
	0xb5, 0x18, 0x0e, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x8a, 0xb5, 0x18, 0x13, 0x12, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_automate_gateway_api_event_feed_event_feed_proto_goTypes = []interface{}{
	(*request.EventFilter)(nil),       // 0: chef.automate.api.event_feed.request.EventFilter
	(*request.EventCountsFilter)(nil), // 1: chef.automate.api.event_feed.request.EventCountsFilter
	(*request.EventStrings)(nil),      // 2: chef.automate.api.event_feed.request.EventStrings
	(*response.Events)(nil),           // 3: chef.automate.api.event_feed.response.Events
	(*response.EventCounts)(nil),      // 4: chef.automate.api.event_feed.response.EventCounts
	(*response.EventStrings)(nil),     // 5: chef.automate.api.event_feed.response.EventStrings
}
var file_automate_gateway_api_event_feed_event_feed_proto_depIdxs = []int32{
	0, // 0: chef.automate.api.event_feed.EventFeed.GetEventFeed:input_type -> chef.automate.api.event_feed.request.EventFilter
	1, // 1: chef.automate.api.event_feed.EventFeed.GetEventTypeCounts:input_type -> chef.automate.api.event_feed.request.EventCountsFilter
	1, // 2: chef.automate.api.event_feed.EventFeed.GetEventTaskCounts:input_type -> chef.automate.api.event_feed.request.EventCountsFilter
	2, // 3: chef.automate.api.event_feed.EventFeed.GetEventStringBuckets:input_type -> chef.automate.api.event_feed.request.EventStrings
	3, // 4: chef.automate.api.event_feed.EventFeed.GetEventFeed:output_type -> chef.automate.api.event_feed.response.Events
	4, // 5: chef.automate.api.event_feed.EventFeed.GetEventTypeCounts:output_type -> chef.automate.api.event_feed.response.EventCounts
	4, // 6: chef.automate.api.event_feed.EventFeed.GetEventTaskCounts:output_type -> chef.automate.api.event_feed.response.EventCounts
	5, // 7: chef.automate.api.event_feed.EventFeed.GetEventStringBuckets:output_type -> chef.automate.api.event_feed.response.EventStrings
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_automate_gateway_api_event_feed_event_feed_proto_init() }
func file_automate_gateway_api_event_feed_event_feed_proto_init() {
	if File_automate_gateway_api_event_feed_event_feed_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_automate_gateway_api_event_feed_event_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_automate_gateway_api_event_feed_event_feed_proto_goTypes,
		DependencyIndexes: file_automate_gateway_api_event_feed_event_feed_proto_depIdxs,
	}.Build()
	File_automate_gateway_api_event_feed_event_feed_proto = out.File
	file_automate_gateway_api_event_feed_event_feed_proto_rawDesc = nil
	file_automate_gateway_api_event_feed_event_feed_proto_goTypes = nil
	file_automate_gateway_api_event_feed_event_feed_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventFeedClient is the client API for EventFeed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventFeedClient interface {
	//
	//List Events
	//
	//Returns a list of recorded events in Chef Automate, such as Infra Server events (cookbook creation, policyfile updates, and node creation) and Chef Automate internal events (profile installs and scan job creation).
	//Adding a filter makes a list of all events that meet the filter criteria.
	//
	//Example:
	//```
	//eventfeed?collapse=true&filter=organization:The%2520Watchmen&page_size=100&start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventFeed(ctx context.Context, in *request.EventFilter, opts ...grpc.CallOption) (*response.Events, error)
	//
	//List Count of Event Type Occurrence
	//
	//Returns totals for role, cookbook, and organization event types.
	//
	//Example:
	//```
	//event_type_counts?start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventTypeCounts(ctx context.Context, in *request.EventCountsFilter, opts ...grpc.CallOption) (*response.EventCounts, error)
	//
	//List Counts Per Event Task Occurrence
	//
	//Returns totals for update, create, and delete event tasks, which are the actions taken on the event.
	//
	//Example:
	//```
	//event_task_counts?start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventTaskCounts(ctx context.Context, in *request.EventCountsFilter, opts ...grpc.CallOption) (*response.EventCounts, error)
	//
	//List Summary Stats About Events
	//
	//Returns data that populates the guitar strings visual on the top of the event feed.
	//
	//Example:
	//```
	//eventstrings?timezone=America/Denver&hours_between=1&start=2020-06-19&end=2020-06-25
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventStringBuckets(ctx context.Context, in *request.EventStrings, opts ...grpc.CallOption) (*response.EventStrings, error)
}

type eventFeedClient struct {
	cc grpc.ClientConnInterface
}

func NewEventFeedClient(cc grpc.ClientConnInterface) EventFeedClient {
	return &eventFeedClient{cc}
}

func (c *eventFeedClient) GetEventFeed(ctx context.Context, in *request.EventFilter, opts ...grpc.CallOption) (*response.Events, error) {
	out := new(response.Events)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeed/GetEventFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedClient) GetEventTypeCounts(ctx context.Context, in *request.EventCountsFilter, opts ...grpc.CallOption) (*response.EventCounts, error) {
	out := new(response.EventCounts)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeed/GetEventTypeCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedClient) GetEventTaskCounts(ctx context.Context, in *request.EventCountsFilter, opts ...grpc.CallOption) (*response.EventCounts, error) {
	out := new(response.EventCounts)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeed/GetEventTaskCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedClient) GetEventStringBuckets(ctx context.Context, in *request.EventStrings, opts ...grpc.CallOption) (*response.EventStrings, error) {
	out := new(response.EventStrings)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeed/GetEventStringBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventFeedServer is the server API for EventFeed service.
type EventFeedServer interface {
	//
	//List Events
	//
	//Returns a list of recorded events in Chef Automate, such as Infra Server events (cookbook creation, policyfile updates, and node creation) and Chef Automate internal events (profile installs and scan job creation).
	//Adding a filter makes a list of all events that meet the filter criteria.
	//
	//Example:
	//```
	//eventfeed?collapse=true&filter=organization:The%2520Watchmen&page_size=100&start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventFeed(context.Context, *request.EventFilter) (*response.Events, error)
	//
	//List Count of Event Type Occurrence
	//
	//Returns totals for role, cookbook, and organization event types.
	//
	//Example:
	//```
	//event_type_counts?start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventTypeCounts(context.Context, *request.EventCountsFilter) (*response.EventCounts, error)
	//
	//List Counts Per Event Task Occurrence
	//
	//Returns totals for update, create, and delete event tasks, which are the actions taken on the event.
	//
	//Example:
	//```
	//event_task_counts?start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventTaskCounts(context.Context, *request.EventCountsFilter) (*response.EventCounts, error)
	//
	//List Summary Stats About Events
	//
	//Returns data that populates the guitar strings visual on the top of the event feed.
	//
	//Example:
	//```
	//eventstrings?timezone=America/Denver&hours_between=1&start=2020-06-19&end=2020-06-25
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventStringBuckets(context.Context, *request.EventStrings) (*response.EventStrings, error)
}

// UnimplementedEventFeedServer can be embedded to have forward compatible implementations.
type UnimplementedEventFeedServer struct {
}

func (*UnimplementedEventFeedServer) GetEventFeed(context.Context, *request.EventFilter) (*response.Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventFeed not implemented")
}
func (*UnimplementedEventFeedServer) GetEventTypeCounts(context.Context, *request.EventCountsFilter) (*response.EventCounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventTypeCounts not implemented")
}
func (*UnimplementedEventFeedServer) GetEventTaskCounts(context.Context, *request.EventCountsFilter) (*response.EventCounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventTaskCounts not implemented")
}
func (*UnimplementedEventFeedServer) GetEventStringBuckets(context.Context, *request.EventStrings) (*response.EventStrings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventStringBuckets not implemented")
}

func RegisterEventFeedServer(s *grpc.Server, srv EventFeedServer) {
	s.RegisterService(&_EventFeed_serviceDesc, srv)
}

func _EventFeed_GetEventFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EventFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServer).GetEventFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeed/GetEventFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServer).GetEventFeed(ctx, req.(*request.EventFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeed_GetEventTypeCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EventCountsFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServer).GetEventTypeCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeed/GetEventTypeCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServer).GetEventTypeCounts(ctx, req.(*request.EventCountsFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeed_GetEventTaskCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EventCountsFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServer).GetEventTaskCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeed/GetEventTaskCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServer).GetEventTaskCounts(ctx, req.(*request.EventCountsFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeed_GetEventStringBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EventStrings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServer).GetEventStringBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeed/GetEventStringBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServer).GetEventStringBuckets(ctx, req.(*request.EventStrings))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventFeed_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.event_feed.EventFeed",
	HandlerType: (*EventFeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEventFeed",
			Handler:    _EventFeed_GetEventFeed_Handler,
		},
		{
			MethodName: "GetEventTypeCounts",
			Handler:    _EventFeed_GetEventTypeCounts_Handler,
		},
		{
			MethodName: "GetEventTaskCounts",
			Handler:    _EventFeed_GetEventTaskCounts_Handler,
		},
		{
			MethodName: "GetEventStringBuckets",
			Handler:    _EventFeed_GetEventStringBuckets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "automate-gateway/api/event_feed/event_feed.proto",
}