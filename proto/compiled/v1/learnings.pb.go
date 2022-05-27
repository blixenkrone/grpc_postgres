// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: v1/learnings.proto

package learningsv1

import (
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

type Level int32

const (
	Level_LEVEL_UNSPECIFIED  Level = 0
	Level_LEVEL_BEGINNER     Level = 1
	Level_LEVEL_INTERMEDIATE Level = 2
	Level_LEVEL_PRO          Level = 3
	Level_LEVEL_MASTER       Level = 4
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "LEVEL_UNSPECIFIED",
		1: "LEVEL_BEGINNER",
		2: "LEVEL_INTERMEDIATE",
		3: "LEVEL_PRO",
		4: "LEVEL_MASTER",
	}
	Level_value = map[string]int32{
		"LEVEL_UNSPECIFIED":  0,
		"LEVEL_BEGINNER":     1,
		"LEVEL_INTERMEDIATE": 2,
		"LEVEL_PRO":          3,
		"LEVEL_MASTER":       4,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_learnings_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_v1_learnings_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_v1_learnings_proto_rawDescGZIP(), []int{0}
}

type Material struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url []string `protobuf:"bytes,2,rep,name=url,proto3" json:"url,omitempty"`
}

func (x *Material) Reset() {
	*x = Material{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_learnings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Material) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Material) ProtoMessage() {}

func (x *Material) ProtoReflect() protoreflect.Message {
	mi := &file_v1_learnings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Material.ProtoReflect.Descriptor instead.
func (*Material) Descriptor() ([]byte, []int) {
	return file_v1_learnings_proto_rawDescGZIP(), []int{0}
}

func (x *Material) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Material) GetUrl() []string {
	if x != nil {
		return x.Url
	}
	return nil
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label       string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_learnings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_v1_learnings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_v1_learnings_proto_rawDescGZIP(), []int{1}
}

func (x *Category) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Category) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Category) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CourseId string      `protobuf:"bytes,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Level    Level       `protobuf:"varint,3,opt,name=level,proto3,enum=proto.learnings.v1.Level" json:"level,omitempty"`
	Material []*Material `protobuf:"bytes,4,rep,name=material,proto3" json:"material,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_learnings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_v1_learnings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_v1_learnings_proto_rawDescGZIP(), []int{2}
}

func (x *Module) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Module) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *Module) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_LEVEL_UNSPECIFIED
}

func (x *Module) GetMaterial() []*Material {
	if x != nil {
		return x.Material
	}
	return nil
}

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IsActive  bool                   `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Name      string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	ModuleIds []string               `protobuf:"bytes,5,rep,name=module_ids,json=moduleIds,proto3" json:"module_ids,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_learnings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_v1_learnings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_v1_learnings_proto_rawDescGZIP(), []int{3}
}

func (x *Course) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Course) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetModuleIds() []string {
	if x != nil {
		return x.ModuleIds
	}
	return nil
}

func (x *Course) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Course) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type AddCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsActive bool   `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Level    Level  `protobuf:"varint,3,opt,name=level,proto3,enum=proto.learnings.v1.Level" json:"level,omitempty"`
}

func (x *AddCourseRequest) Reset() {
	*x = AddCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_learnings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCourseRequest) ProtoMessage() {}

func (x *AddCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_learnings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCourseRequest.ProtoReflect.Descriptor instead.
func (*AddCourseRequest) Descriptor() ([]byte, []int) {
	return file_v1_learnings_proto_rawDescGZIP(), []int{4}
}

func (x *AddCourseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddCourseRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *AddCourseRequest) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_LEVEL_UNSPECIFIED
}

type ListCoursesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCoursesRequest) Reset() {
	*x = ListCoursesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_learnings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCoursesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoursesRequest) ProtoMessage() {}

func (x *ListCoursesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_learnings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoursesRequest.ProtoReflect.Descriptor instead.
func (*ListCoursesRequest) Descriptor() ([]byte, []int) {
	return file_v1_learnings_proto_rawDescGZIP(), []int{5}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_learnings_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_learnings_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_v1_learnings_proto_rawDescGZIP(), []int{6}
}

func (x *PingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_v1_learnings_proto protoreflect.FileDescriptor

var file_v1_learnings_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0x52, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa0, 0x01, 0x0a, 0x06, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64,
	0x12, 0x2f, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x22, 0xde, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x74, 0x0a, 0x10,
	0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x2f, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2a, 0x6b, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x15, 0x0a, 0x11, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x42, 0x45, 0x47, 0x49,
	0x4e, 0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x50, 0x52, 0x4f, 0x10, 0x03, 0x12, 0x10, 0x0a,
	0x0c, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x10, 0x04, 0x32,
	0xfe, 0x01, 0x0a, 0x10, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x42, 0xc3, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c,
	0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x4c, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x69, 0x78, 0x65, 0x6e,
	0x6b, 0x72, 0x6f, 0x6e, 0x65, 0x2f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x50, 0x4c, 0x58, 0xaa, 0x02, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x65, 0x61,
	0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5c, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_learnings_proto_rawDescOnce sync.Once
	file_v1_learnings_proto_rawDescData = file_v1_learnings_proto_rawDesc
)

func file_v1_learnings_proto_rawDescGZIP() []byte {
	file_v1_learnings_proto_rawDescOnce.Do(func() {
		file_v1_learnings_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_learnings_proto_rawDescData)
	})
	return file_v1_learnings_proto_rawDescData
}

var file_v1_learnings_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_learnings_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v1_learnings_proto_goTypes = []interface{}{
	(Level)(0),                    // 0: proto.learnings.v1.Level
	(*Material)(nil),              // 1: proto.learnings.v1.Material
	(*Category)(nil),              // 2: proto.learnings.v1.Category
	(*Module)(nil),                // 3: proto.learnings.v1.Module
	(*Course)(nil),                // 4: proto.learnings.v1.Course
	(*AddCourseRequest)(nil),      // 5: proto.learnings.v1.AddCourseRequest
	(*ListCoursesRequest)(nil),    // 6: proto.learnings.v1.ListCoursesRequest
	(*PingResponse)(nil),          // 7: proto.learnings.v1.PingResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_v1_learnings_proto_depIdxs = []int32{
	0, // 0: proto.learnings.v1.Module.level:type_name -> proto.learnings.v1.Level
	1, // 1: proto.learnings.v1.Module.material:type_name -> proto.learnings.v1.Material
	8, // 2: proto.learnings.v1.Course.created_at:type_name -> google.protobuf.Timestamp
	8, // 3: proto.learnings.v1.Course.updated_at:type_name -> google.protobuf.Timestamp
	0, // 4: proto.learnings.v1.AddCourseRequest.level:type_name -> proto.learnings.v1.Level
	9, // 5: proto.learnings.v1.LearningsService.Ping:input_type -> google.protobuf.Empty
	5, // 6: proto.learnings.v1.LearningsService.AddCourse:input_type -> proto.learnings.v1.AddCourseRequest
	6, // 7: proto.learnings.v1.LearningsService.ListCourses:input_type -> proto.learnings.v1.ListCoursesRequest
	7, // 8: proto.learnings.v1.LearningsService.Ping:output_type -> proto.learnings.v1.PingResponse
	4, // 9: proto.learnings.v1.LearningsService.AddCourse:output_type -> proto.learnings.v1.Course
	4, // 10: proto.learnings.v1.LearningsService.ListCourses:output_type -> proto.learnings.v1.Course
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_v1_learnings_proto_init() }
func file_v1_learnings_proto_init() {
	if File_v1_learnings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_learnings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Material); i {
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
		file_v1_learnings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_v1_learnings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
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
		file_v1_learnings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
		file_v1_learnings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCourseRequest); i {
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
		file_v1_learnings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCoursesRequest); i {
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
		file_v1_learnings_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
			RawDescriptor: file_v1_learnings_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_learnings_proto_goTypes,
		DependencyIndexes: file_v1_learnings_proto_depIdxs,
		EnumInfos:         file_v1_learnings_proto_enumTypes,
		MessageInfos:      file_v1_learnings_proto_msgTypes,
	}.Build()
	File_v1_learnings_proto = out.File
	file_v1_learnings_proto_rawDesc = nil
	file_v1_learnings_proto_goTypes = nil
	file_v1_learnings_proto_depIdxs = nil
}
