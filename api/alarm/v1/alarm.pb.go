// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: alarm/v1/alarm.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	api "prometheus-manager/api"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatisticsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatisticsRequest) Reset() {
	*x = StatisticsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticsRequest) ProtoMessage() {}

func (x *StatisticsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticsRequest.ProtoReflect.Descriptor instead.
func (*StatisticsRequest) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{0}
}

type StatisticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总报警数量
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 各个报警页面各自的报警数量
	Pages map[string]int64 `protobuf:"bytes,2,rep,name=pages,proto3" json:"pages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// 各个报警等级的报警数量
	Levels map[string]int64 `protobuf:"bytes,3,rep,name=levels,proto3" json:"levels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// 各个业务属性的报警数量
	Business map[string]int64 `protobuf:"bytes,4,rep,name=business,proto3" json:"business,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// 各个规则组的报警数量
	Groups map[string]int64 `protobuf:"bytes,5,rep,name=groups,proto3" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *StatisticsResponse) Reset() {
	*x = StatisticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticsResponse) ProtoMessage() {}

func (x *StatisticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticsResponse.ProtoReflect.Descriptor instead.
func (*StatisticsResponse) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{1}
}

func (x *StatisticsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *StatisticsResponse) GetPages() map[string]int64 {
	if x != nil {
		return x.Pages
	}
	return nil
}

func (x *StatisticsResponse) GetLevels() map[string]int64 {
	if x != nil {
		return x.Levels
	}
	return nil
}

func (x *StatisticsResponse) GetBusiness() map[string]int64 {
	if x != nil {
		return x.Business
	}
	return nil
}

func (x *StatisticsResponse) GetGroups() map[string]int64 {
	if x != nil {
		return x.Groups
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 报警页面
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	// 报警等级
	Level int32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	// 业务属性
	Business int32 `protobuf:"varint,3,opt,name=business,proto3" json:"business,omitempty"`
	// 规则组
	Group int32 `protobuf:"varint,4,opt,name=group,proto3" json:"group,omitempty"`
	// 报警状态
	Status int32 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	// 报警开始时间
	StartAt int64 `protobuf:"varint,6,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	// 报警结束时间
	EndAt int64 `protobuf:"varint,7,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{2}
}

func (x *Filter) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Filter) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Filter) GetBusiness() int32 {
	if x != nil {
		return x.Business
	}
	return 0
}

func (x *Filter) GetGroup() int32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *Filter) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Filter) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *Filter) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	Page *api.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	// 过滤参数
	Filter *Filter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{3}
}

func (x *ListRequest) GetPage() *api.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListRequest) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type AlarmItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AlarmItem) Reset() {
	*x = AlarmItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlarmItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlarmItem) ProtoMessage() {}

func (x *AlarmItem) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlarmItem.ProtoReflect.Descriptor instead.
func (*AlarmItem) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{4}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 报警列表
	Alarms []*AlarmItem `protobuf:"bytes,1,rep,name=alarms,proto3" json:"alarms,omitempty"`
	// 分页参数
	Page     *api.PageReply `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Response *api.Response  `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alarm_v1_alarm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alarm_v1_alarm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_alarm_v1_alarm_proto_rawDescGZIP(), []int{5}
}

func (x *ListResponse) GetAlarms() []*AlarmItem {
	if x != nil {
		return x.Alarms
	}
	return nil
}

func (x *ListResponse) GetPage() *api.PageReply {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListResponse) GetResponse() *api.Response {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_alarm_v1_alarm_proto protoreflect.FileDescriptor

var file_alarm_v1_alarm_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x61, 0x72, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13,
	0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xb2, 0x04, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x41, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x12, 0x4a, 0x0a, 0x08, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x44, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x3b, 0x0a, 0x0d, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a,
	0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xae, 0x01, 0x0a, 0x06, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x22, 0x61, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x2c,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x0b, 0x0a, 0x09,
	0x41, 0x6c, 0x61, 0x72, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x8e, 0x01, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x61, 0x6c,
	0x61, 0x72, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x06, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x29, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xcd, 0x01, 0x0a, 0x05, 0x41,
	0x6c, 0x61, 0x72, 0x6d, 0x12, 0x6d, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f,
	0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x12, 0x55, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x61,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x6c, 0x61,
	0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x34, 0x0a, 0x0c, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x22, 0x70, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alarm_v1_alarm_proto_rawDescOnce sync.Once
	file_alarm_v1_alarm_proto_rawDescData = file_alarm_v1_alarm_proto_rawDesc
)

func file_alarm_v1_alarm_proto_rawDescGZIP() []byte {
	file_alarm_v1_alarm_proto_rawDescOnce.Do(func() {
		file_alarm_v1_alarm_proto_rawDescData = protoimpl.X.CompressGZIP(file_alarm_v1_alarm_proto_rawDescData)
	})
	return file_alarm_v1_alarm_proto_rawDescData
}

var file_alarm_v1_alarm_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_alarm_v1_alarm_proto_goTypes = []interface{}{
	(*StatisticsRequest)(nil),  // 0: api.alarm.v1.StatisticsRequest
	(*StatisticsResponse)(nil), // 1: api.alarm.v1.StatisticsResponse
	(*Filter)(nil),             // 2: api.alarm.v1.Filter
	(*ListRequest)(nil),        // 3: api.alarm.v1.ListRequest
	(*AlarmItem)(nil),          // 4: api.alarm.v1.AlarmItem
	(*ListResponse)(nil),       // 5: api.alarm.v1.ListResponse
	nil,                        // 6: api.alarm.v1.StatisticsResponse.PagesEntry
	nil,                        // 7: api.alarm.v1.StatisticsResponse.LevelsEntry
	nil,                        // 8: api.alarm.v1.StatisticsResponse.BusinessEntry
	nil,                        // 9: api.alarm.v1.StatisticsResponse.GroupsEntry
	(*api.PageRequest)(nil),    // 10: api.PageRequest
	(*api.PageReply)(nil),      // 11: api.PageReply
	(*api.Response)(nil),       // 12: api.Response
}
var file_alarm_v1_alarm_proto_depIdxs = []int32{
	6,  // 0: api.alarm.v1.StatisticsResponse.pages:type_name -> api.alarm.v1.StatisticsResponse.PagesEntry
	7,  // 1: api.alarm.v1.StatisticsResponse.levels:type_name -> api.alarm.v1.StatisticsResponse.LevelsEntry
	8,  // 2: api.alarm.v1.StatisticsResponse.business:type_name -> api.alarm.v1.StatisticsResponse.BusinessEntry
	9,  // 3: api.alarm.v1.StatisticsResponse.groups:type_name -> api.alarm.v1.StatisticsResponse.GroupsEntry
	10, // 4: api.alarm.v1.ListRequest.page:type_name -> api.PageRequest
	2,  // 5: api.alarm.v1.ListRequest.filter:type_name -> api.alarm.v1.Filter
	4,  // 6: api.alarm.v1.ListResponse.alarms:type_name -> api.alarm.v1.AlarmItem
	11, // 7: api.alarm.v1.ListResponse.page:type_name -> api.PageReply
	12, // 8: api.alarm.v1.ListResponse.response:type_name -> api.Response
	0,  // 9: api.alarm.v1.Alarm.Statistics:input_type -> api.alarm.v1.StatisticsRequest
	3,  // 10: api.alarm.v1.Alarm.List:input_type -> api.alarm.v1.ListRequest
	1,  // 11: api.alarm.v1.Alarm.Statistics:output_type -> api.alarm.v1.StatisticsResponse
	5,  // 12: api.alarm.v1.Alarm.List:output_type -> api.alarm.v1.ListResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_alarm_v1_alarm_proto_init() }
func file_alarm_v1_alarm_proto_init() {
	if File_alarm_v1_alarm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alarm_v1_alarm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticsRequest); i {
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
		file_alarm_v1_alarm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticsResponse); i {
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
		file_alarm_v1_alarm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_alarm_v1_alarm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_alarm_v1_alarm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlarmItem); i {
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
		file_alarm_v1_alarm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_alarm_v1_alarm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alarm_v1_alarm_proto_goTypes,
		DependencyIndexes: file_alarm_v1_alarm_proto_depIdxs,
		MessageInfos:      file_alarm_v1_alarm_proto_msgTypes,
	}.Build()
	File_alarm_v1_alarm_proto = out.File
	file_alarm_v1_alarm_proto_rawDesc = nil
	file_alarm_v1_alarm_proto_goTypes = nil
	file_alarm_v1_alarm_proto_depIdxs = nil
}