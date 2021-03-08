//
// These are messages and services relating to mental health tracking data, allowing for
// the logging of user mental health data and tracking the quality of their mental health state
// from day to day.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: proto/health.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/golang/protobuf/ptypes/timestamp"
	common "github.com/kic/media/pkg/proto/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//
//These are errors used to inform the client requesting health data what the issue is.
//The variable names denote the issue.
type HealthDataError int32

const (
	// USER_NOT_FOUND denotes if user is not found.
	HealthDataError_USER_NOT_FOUND HealthDataError = 0
	// DATE_NOT_FOUND denotes if date is not found.
	HealthDataError_DATE_NOT_FOUND HealthDataError = 1
)

// Enum value maps for HealthDataError.
var (
	HealthDataError_name = map[int32]string{
		0: "USER_NOT_FOUND",
		1: "DATE_NOT_FOUND",
	}
	HealthDataError_value = map[string]int32{
		"USER_NOT_FOUND": 0,
		"DATE_NOT_FOUND": 1,
	}
)

func (x HealthDataError) Enum() *HealthDataError {
	p := new(HealthDataError)
	*p = x
	return p
}

func (x HealthDataError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthDataError) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_health_proto_enumTypes[0].Descriptor()
}

func (HealthDataError) Type() protoreflect.EnumType {
	return &file_proto_health_proto_enumTypes[0]
}

func (x HealthDataError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthDataError.Descriptor instead.
func (HealthDataError) EnumDescriptor() ([]byte, []int) {
	return file_proto_health_proto_rawDescGZIP(), []int{0}
}

//
//Response to a user when there is a mental health data error.
type HealthDataErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Error denotes if error occurred with health data.
	Error HealthDataError `protobuf:"varint,1,opt,name=error,proto3,enum=kic.health.HealthDataError" json:"error,omitempty"`
}

func (x *HealthDataErrorResponse) Reset() {
	*x = HealthDataErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthDataErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthDataErrorResponse) ProtoMessage() {}

func (x *HealthDataErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthDataErrorResponse.ProtoReflect.Descriptor instead.
func (*HealthDataErrorResponse) Descriptor() ([]byte, []int) {
	return file_proto_health_proto_rawDescGZIP(), []int{0}
}

func (x *HealthDataErrorResponse) GetError() HealthDataError {
	if x != nil {
		return x.Error
	}
	return HealthDataError_USER_NOT_FOUND
}

//
//Request from a user to get their mental health tracking data.
type GetHealthDataForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the user in the user database, used globally for identification.
	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetHealthDataForUserRequest) Reset() {
	*x = GetHealthDataForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHealthDataForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHealthDataForUserRequest) ProtoMessage() {}

func (x *GetHealthDataForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHealthDataForUserRequest.ProtoReflect.Descriptor instead.
func (*GetHealthDataForUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_health_proto_rawDescGZIP(), []int{1}
}

func (x *GetHealthDataForUserRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

//
//Response to a user with complete mental health log
type MentalHealthLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Date of Mental Health Log Entry
	LogDate *common.Date `protobuf:"bytes,1,opt,name=logDate,proto3" json:"logDate,omitempty"`
	// Score denotes the mental health tracking score from logDate.
	Score uint32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *MentalHealthLog) Reset() {
	*x = MentalHealthLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MentalHealthLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MentalHealthLog) ProtoMessage() {}

func (x *MentalHealthLog) ProtoReflect() protoreflect.Message {
	mi := &file_proto_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MentalHealthLog.ProtoReflect.Descriptor instead.
func (*MentalHealthLog) Descriptor() ([]byte, []int) {
	return file_proto_health_proto_rawDescGZIP(), []int{2}
}

func (x *MentalHealthLog) GetLogDate() *common.Date {
	if x != nil {
		return x.LogDate
	}
	return nil
}

func (x *MentalHealthLog) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

//
//Response to a user when user asks for health data.
type GetHealthDataForUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Error denotes if error occurred when obtaining health data and what the error was.
	Error HealthDataError `protobuf:"varint,1,opt,name=error,proto3,enum=kic.health.HealthDataError" json:"error,omitempty"`
	// healthData denotes the data that was requested by user from mental health log
	HealthData []*MentalHealthLog `protobuf:"bytes,2,rep,name=healthData,proto3" json:"healthData,omitempty"`
}

func (x *GetHealthDataForUserResponse) Reset() {
	*x = GetHealthDataForUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_health_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHealthDataForUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHealthDataForUserResponse) ProtoMessage() {}

func (x *GetHealthDataForUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_health_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHealthDataForUserResponse.ProtoReflect.Descriptor instead.
func (*GetHealthDataForUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_health_proto_rawDescGZIP(), []int{3}
}

func (x *GetHealthDataForUserResponse) GetError() HealthDataError {
	if x != nil {
		return x.Error
	}
	return HealthDataError_USER_NOT_FOUND
}

func (x *GetHealthDataForUserResponse) GetHealthData() []*MentalHealthLog {
	if x != nil {
		return x.HealthData
	}
	return nil
}

//
//Request from a user to add their mental health data to MentalHealthLog.
type AddHealthDataForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the user in the user database, used globally for identification.
	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// newEntry denotes the ID of the new entry that is requested to be made.
	NewEntry *MentalHealthLog `protobuf:"bytes,2,opt,name=newEntry,proto3" json:"newEntry,omitempty"`
}

func (x *AddHealthDataForUserRequest) Reset() {
	*x = AddHealthDataForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_health_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHealthDataForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHealthDataForUserRequest) ProtoMessage() {}

func (x *AddHealthDataForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_health_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHealthDataForUserRequest.ProtoReflect.Descriptor instead.
func (*AddHealthDataForUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_health_proto_rawDescGZIP(), []int{4}
}

func (x *AddHealthDataForUserRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *AddHealthDataForUserRequest) GetNewEntry() *MentalHealthLog {
	if x != nil {
		return x.NewEntry
	}
	return nil
}

//
//Request from a user to delete their mental health data from MentalHealthLog.
type DeleteHealthDataForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the user in the user database, used globally for identification.
	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	//Either delete all of data or delete specific date
	//
	// Types that are assignable to Data:
	//	*DeleteHealthDataForUserRequest_All
	//	*DeleteHealthDataForUserRequest_DateToRemove
	Data isDeleteHealthDataForUserRequest_Data `protobuf_oneof:"data"`
}

func (x *DeleteHealthDataForUserRequest) Reset() {
	*x = DeleteHealthDataForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_health_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHealthDataForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHealthDataForUserRequest) ProtoMessage() {}

func (x *DeleteHealthDataForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_health_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHealthDataForUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteHealthDataForUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_health_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteHealthDataForUserRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (m *DeleteHealthDataForUserRequest) GetData() isDeleteHealthDataForUserRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *DeleteHealthDataForUserRequest) GetAll() bool {
	if x, ok := x.GetData().(*DeleteHealthDataForUserRequest_All); ok {
		return x.All
	}
	return false
}

func (x *DeleteHealthDataForUserRequest) GetDateToRemove() *common.Date {
	if x, ok := x.GetData().(*DeleteHealthDataForUserRequest_DateToRemove); ok {
		return x.DateToRemove
	}
	return nil
}

type isDeleteHealthDataForUserRequest_Data interface {
	isDeleteHealthDataForUserRequest_Data()
}

type DeleteHealthDataForUserRequest_All struct {
	// all denotes if all of the health data should be removed or not.
	All bool `protobuf:"varint,2,opt,name=all,proto3,oneof"`
}

type DeleteHealthDataForUserRequest_DateToRemove struct {
	// dateToRemove denotes the date of the mental health log data to remove.
	DateToRemove *common.Date `protobuf:"bytes,3,opt,name=dateToRemove,proto3,oneof"`
}

func (*DeleteHealthDataForUserRequest_All) isDeleteHealthDataForUserRequest_Data() {}

func (*DeleteHealthDataForUserRequest_DateToRemove) isDeleteHealthDataForUserRequest_Data() {}

//
//Response to a user when user asks to delete health data.
type DeleteHealthDataForUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Error denotes if error occurred when deleting health data and the ID of the error it was.
	Error HealthDataError `protobuf:"varint,1,opt,name=error,proto3,enum=kic.health.HealthDataError" json:"error,omitempty"`
	// entriesDeleted denotes the mental health log entries that was successfully deleted for the user
	EntriesDeleted uint32 `protobuf:"varint,2,opt,name=entriesDeleted,proto3" json:"entriesDeleted,omitempty"`
}

func (x *DeleteHealthDataForUserResponse) Reset() {
	*x = DeleteHealthDataForUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_health_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHealthDataForUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHealthDataForUserResponse) ProtoMessage() {}

func (x *DeleteHealthDataForUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_health_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHealthDataForUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteHealthDataForUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_health_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteHealthDataForUserResponse) GetError() HealthDataError {
	if x != nil {
		return x.Error
	}
	return HealthDataError_USER_NOT_FOUND
}

func (x *DeleteHealthDataForUserResponse) GetEntriesDeleted() uint32 {
	if x != nil {
		return x.EntriesDeleted
	}
	return 0
}

//
//Request from a user to update their mental health tracking data for a particular date.
type UpdateHealthDataForDateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the user in the user database, used globally for identification.
	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// The desiredLogInfo denotes the log info that the user would like to update.
	DesiredLogInfo *MentalHealthLog `protobuf:"bytes,2,opt,name=desiredLogInfo,proto3" json:"desiredLogInfo,omitempty"`
}

func (x *UpdateHealthDataForDateRequest) Reset() {
	*x = UpdateHealthDataForDateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_health_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHealthDataForDateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHealthDataForDateRequest) ProtoMessage() {}

func (x *UpdateHealthDataForDateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_health_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHealthDataForDateRequest.ProtoReflect.Descriptor instead.
func (*UpdateHealthDataForDateRequest) Descriptor() ([]byte, []int) {
	return file_proto_health_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateHealthDataForDateRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UpdateHealthDataForDateRequest) GetDesiredLogInfo() *MentalHealthLog {
	if x != nil {
		return x.DesiredLogInfo
	}
	return nil
}

var File_proto_health_proto protoreflect.FileDescriptor

var file_proto_health_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x69, 0x63, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x17, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x6b, 0x69, 0x63, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x35, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x53, 0x0a, 0x0f, 0x4d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x12, 0x2a, 0x0a,
	0x07, 0x6c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6b, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x07, 0x6c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22,
	0x8e, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74,
	0x61, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x6b, 0x69, 0x63, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0a, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x69, 0x63, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2e, 0x4d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x4c, 0x6f, 0x67, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x6e, 0x0a, 0x1b, 0x41, 0x64, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74,
	0x61, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x69, 0x63, 0x2e,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x4d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x22, 0x8c, 0x01, 0x0a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x03, 0x61,
	0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x12,
	0x36, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x7c, 0x0a, 0x1f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x6b, 0x69, 0x63, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x7d, 0x0a,
	0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74,
	0x61, 0x46, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x43, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x69, 0x72,
	0x65, 0x64, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6b, 0x69, 0x63, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x4d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x52, 0x0e, 0x64, 0x65,
	0x73, 0x69, 0x72, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x39, 0x0a, 0x0f,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x32, 0xc1, 0x03, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x69, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x27, 0x2e, 0x6b, 0x69, 0x63, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6b, 0x69,
	0x63, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x27, 0x2e,
	0x6b, 0x69, 0x63, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6b, 0x69, 0x63, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x17, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46,
	0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x6b, 0x69, 0x63, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6b, 0x69, 0x63, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61,
	0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6a, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x6b, 0x69, 0x63,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6b, 0x69, 0x63, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_health_proto_rawDescOnce sync.Once
	file_proto_health_proto_rawDescData = file_proto_health_proto_rawDesc
)

func file_proto_health_proto_rawDescGZIP() []byte {
	file_proto_health_proto_rawDescOnce.Do(func() {
		file_proto_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_health_proto_rawDescData)
	})
	return file_proto_health_proto_rawDescData
}

var (
	file_proto_health_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_proto_health_proto_msgTypes  = make([]protoimpl.MessageInfo, 8)
	file_proto_health_proto_goTypes   = []interface{}{
		(HealthDataError)(0),                    // 0: kic.health.HealthDataError
		(*HealthDataErrorResponse)(nil),         // 1: kic.health.HealthDataErrorResponse
		(*GetHealthDataForUserRequest)(nil),     // 2: kic.health.GetHealthDataForUserRequest
		(*MentalHealthLog)(nil),                 // 3: kic.health.MentalHealthLog
		(*GetHealthDataForUserResponse)(nil),    // 4: kic.health.GetHealthDataForUserResponse
		(*AddHealthDataForUserRequest)(nil),     // 5: kic.health.AddHealthDataForUserRequest
		(*DeleteHealthDataForUserRequest)(nil),  // 6: kic.health.DeleteHealthDataForUserRequest
		(*DeleteHealthDataForUserResponse)(nil), // 7: kic.health.DeleteHealthDataForUserResponse
		(*UpdateHealthDataForDateRequest)(nil),  // 8: kic.health.UpdateHealthDataForDateRequest
		(*common.Date)(nil),                     // 9: kic.common.Date
	}
)

var file_proto_health_proto_depIdxs = []int32{
	0,  // 0: kic.health.HealthDataErrorResponse.error:type_name -> kic.health.HealthDataError
	9,  // 1: kic.health.MentalHealthLog.logDate:type_name -> kic.common.Date
	0,  // 2: kic.health.GetHealthDataForUserResponse.error:type_name -> kic.health.HealthDataError
	3,  // 3: kic.health.GetHealthDataForUserResponse.healthData:type_name -> kic.health.MentalHealthLog
	3,  // 4: kic.health.AddHealthDataForUserRequest.newEntry:type_name -> kic.health.MentalHealthLog
	9,  // 5: kic.health.DeleteHealthDataForUserRequest.dateToRemove:type_name -> kic.common.Date
	0,  // 6: kic.health.DeleteHealthDataForUserResponse.error:type_name -> kic.health.HealthDataError
	3,  // 7: kic.health.UpdateHealthDataForDateRequest.desiredLogInfo:type_name -> kic.health.MentalHealthLog
	2,  // 8: kic.health.HealthTracking.GetHealthDataForUser:input_type -> kic.health.GetHealthDataForUserRequest
	5,  // 9: kic.health.HealthTracking.AddHealthDataForUser:input_type -> kic.health.AddHealthDataForUserRequest
	6,  // 10: kic.health.HealthTracking.DeleteHealthDataForUser:input_type -> kic.health.DeleteHealthDataForUserRequest
	8,  // 11: kic.health.HealthTracking.UpdateHealthDataForDate:input_type -> kic.health.UpdateHealthDataForDateRequest
	4,  // 12: kic.health.HealthTracking.GetHealthDataForUser:output_type -> kic.health.GetHealthDataForUserResponse
	1,  // 13: kic.health.HealthTracking.AddHealthDataForUser:output_type -> kic.health.HealthDataErrorResponse
	7,  // 14: kic.health.HealthTracking.DeleteHealthDataForUser:output_type -> kic.health.DeleteHealthDataForUserResponse
	1,  // 15: kic.health.HealthTracking.UpdateHealthDataForDate:output_type -> kic.health.HealthDataErrorResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_health_proto_init() }
func file_proto_health_proto_init() {
	if File_proto_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthDataErrorResponse); i {
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
		file_proto_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHealthDataForUserRequest); i {
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
		file_proto_health_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MentalHealthLog); i {
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
		file_proto_health_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHealthDataForUserResponse); i {
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
		file_proto_health_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddHealthDataForUserRequest); i {
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
		file_proto_health_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHealthDataForUserRequest); i {
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
		file_proto_health_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHealthDataForUserResponse); i {
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
		file_proto_health_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHealthDataForDateRequest); i {
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
	file_proto_health_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*DeleteHealthDataForUserRequest_All)(nil),
		(*DeleteHealthDataForUserRequest_DateToRemove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_health_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_health_proto_goTypes,
		DependencyIndexes: file_proto_health_proto_depIdxs,
		EnumInfos:         file_proto_health_proto_enumTypes,
		MessageInfos:      file_proto_health_proto_msgTypes,
	}.Build()
	File_proto_health_proto = out.File
	file_proto_health_proto_rawDesc = nil
	file_proto_health_proto_goTypes = nil
	file_proto_health_proto_depIdxs = nil
}
