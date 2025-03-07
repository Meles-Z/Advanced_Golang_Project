// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: teacher.proto

package teacher

import (
	address "github.com/meles-z/edusysnc_grpc/proto/address"
	model "github.com/meles-z/edusysnc_grpc/proto/model"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Teacher struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Model         *model.Model           `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber   string                 `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	CourseIds     []string               `protobuf:"bytes,5,rep,name=course_ids,json=courseIds,proto3" json:"course_ids,omitempty"`                // Use IDs instead of importing Course
	DepartmentId  *string                `protobuf:"bytes,6,opt,name=department_id,json=departmentId,proto3,oneof" json:"department_id,omitempty"` // Use ID instead of importing Department
	Address       *address.Address       `protobuf:"bytes,7,opt,name=address,proto3,oneof" json:"address,omitempty"`                               // Reference to the address
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Teacher) Reset() {
	*x = Teacher{}
	mi := &file_teacher_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Teacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teacher) ProtoMessage() {}

func (x *Teacher) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teacher.ProtoReflect.Descriptor instead.
func (*Teacher) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{0}
}

func (x *Teacher) GetModel() *model.Model {
	if x != nil {
		return x.Model
	}
	return nil
}

func (x *Teacher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Teacher) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Teacher) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Teacher) GetCourseIds() []string {
	if x != nil {
		return x.CourseIds
	}
	return nil
}

func (x *Teacher) GetDepartmentId() string {
	if x != nil && x.DepartmentId != nil {
		return *x.DepartmentId
	}
	return ""
}

func (x *Teacher) GetAddress() *address.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type CreateTeacherRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Teacher       *Teacher               `protobuf:"bytes,1,opt,name=teacher,proto3" json:"teacher,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTeacherRequest) Reset() {
	*x = CreateTeacherRequest{}
	mi := &file_teacher_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeacherRequest) ProtoMessage() {}

func (x *CreateTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeacherRequest.ProtoReflect.Descriptor instead.
func (*CreateTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTeacherRequest) GetTeacher() *Teacher {
	if x != nil {
		return x.Teacher
	}
	return nil
}

type CreateTeacherResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Teacher       *Teacher               `protobuf:"bytes,1,opt,name=teacher,proto3" json:"teacher,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTeacherResponse) Reset() {
	*x = CreateTeacherResponse{}
	mi := &file_teacher_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTeacherResponse) ProtoMessage() {}

func (x *CreateTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTeacherResponse.ProtoReflect.Descriptor instead.
func (*CreateTeacherResponse) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTeacherResponse) GetTeacher() *Teacher {
	if x != nil {
		return x.Teacher
	}
	return nil
}

type GetAllTeachersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllTeachersRequest) Reset() {
	*x = GetAllTeachersRequest{}
	mi := &file_teacher_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllTeachersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTeachersRequest) ProtoMessage() {}

func (x *GetAllTeachersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTeachersRequest.ProtoReflect.Descriptor instead.
func (*GetAllTeachersRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{3}
}

type GetAllTeachersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Teachers      []*Teacher             `protobuf:"bytes,1,rep,name=teachers,proto3" json:"teachers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllTeachersResponse) Reset() {
	*x = GetAllTeachersResponse{}
	mi := &file_teacher_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllTeachersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTeachersResponse) ProtoMessage() {}

func (x *GetAllTeachersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTeachersResponse.ProtoReflect.Descriptor instead.
func (*GetAllTeachersResponse) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllTeachersResponse) GetTeachers() []*Teacher {
	if x != nil {
		return x.Teachers
	}
	return nil
}

type GetTeacherByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTeacherByIdRequest) Reset() {
	*x = GetTeacherByIdRequest{}
	mi := &file_teacher_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTeacherByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeacherByIdRequest) ProtoMessage() {}

func (x *GetTeacherByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeacherByIdRequest.ProtoReflect.Descriptor instead.
func (*GetTeacherByIdRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{5}
}

func (x *GetTeacherByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTeacherByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Teacher       *Teacher               `protobuf:"bytes,1,opt,name=teacher,proto3" json:"teacher,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTeacherByIdResponse) Reset() {
	*x = GetTeacherByIdResponse{}
	mi := &file_teacher_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTeacherByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeacherByIdResponse) ProtoMessage() {}

func (x *GetTeacherByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeacherByIdResponse.ProtoReflect.Descriptor instead.
func (*GetTeacherByIdResponse) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{6}
}

func (x *GetTeacherByIdResponse) GetTeacher() *Teacher {
	if x != nil {
		return x.Teacher
	}
	return nil
}

type UpdateTeacherRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Teacher       *Teacher               `protobuf:"bytes,1,opt,name=teacher,proto3" json:"teacher,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTeacherRequest) Reset() {
	*x = UpdateTeacherRequest{}
	mi := &file_teacher_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeacherRequest) ProtoMessage() {}

func (x *UpdateTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeacherRequest.ProtoReflect.Descriptor instead.
func (*UpdateTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTeacherRequest) GetTeacher() *Teacher {
	if x != nil {
		return x.Teacher
	}
	return nil
}

type UpdateTeacherResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Teacher       *Teacher               `protobuf:"bytes,1,opt,name=teacher,proto3" json:"teacher,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTeacherResponse) Reset() {
	*x = UpdateTeacherResponse{}
	mi := &file_teacher_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTeacherResponse) ProtoMessage() {}

func (x *UpdateTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTeacherResponse.ProtoReflect.Descriptor instead.
func (*UpdateTeacherResponse) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateTeacherResponse) GetTeacher() *Teacher {
	if x != nil {
		return x.Teacher
	}
	return nil
}

type DeleteTeacherRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTeacherRequest) Reset() {
	*x = DeleteTeacherRequest{}
	mi := &file_teacher_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTeacherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeacherRequest) ProtoMessage() {}

func (x *DeleteTeacherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeacherRequest.ProtoReflect.Descriptor instead.
func (*DeleteTeacherRequest) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteTeacherRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTeacherResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTeacherResponse) Reset() {
	*x = DeleteTeacherResponse{}
	mi := &file_teacher_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTeacherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTeacherResponse) ProtoMessage() {}

func (x *DeleteTeacherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teacher_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTeacherResponse.ProtoReflect.Descriptor instead.
func (*DeleteTeacherResponse) Descriptor() ([]byte, []int) {
	return file_teacher_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteTeacherResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_teacher_proto protoreflect.FileDescriptor

var file_teacher_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0d, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a, 0x07, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x0d,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x01,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e,
	0x5f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x47, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0x17, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0x47, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2f, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x22, 0x48, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x74, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x64, 0x75,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x52, 0x07, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xd8, 0x03, 0x0a, 0x0e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x65, 0x64, 0x75, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x64, 0x75, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x23, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x22, 0x2e,
	0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x64,
	0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x65, 0x6c, 0x65, 0x73, 0x2d, 0x7a, 0x2f, 0x65, 0x64, 0x75, 0x73, 0x79, 0x73, 0x6e, 0x63, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_teacher_proto_rawDescOnce sync.Once
	file_teacher_proto_rawDescData []byte
)

func file_teacher_proto_rawDescGZIP() []byte {
	file_teacher_proto_rawDescOnce.Do(func() {
		file_teacher_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teacher_proto_rawDesc), len(file_teacher_proto_rawDesc)))
	})
	return file_teacher_proto_rawDescData
}

var file_teacher_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_teacher_proto_goTypes = []any{
	(*Teacher)(nil),                // 0: edusync_grpc.Teacher
	(*CreateTeacherRequest)(nil),   // 1: edusync_grpc.CreateTeacherRequest
	(*CreateTeacherResponse)(nil),  // 2: edusync_grpc.CreateTeacherResponse
	(*GetAllTeachersRequest)(nil),  // 3: edusync_grpc.GetAllTeachersRequest
	(*GetAllTeachersResponse)(nil), // 4: edusync_grpc.GetAllTeachersResponse
	(*GetTeacherByIdRequest)(nil),  // 5: edusync_grpc.GetTeacherByIdRequest
	(*GetTeacherByIdResponse)(nil), // 6: edusync_grpc.GetTeacherByIdResponse
	(*UpdateTeacherRequest)(nil),   // 7: edusync_grpc.UpdateTeacherRequest
	(*UpdateTeacherResponse)(nil),  // 8: edusync_grpc.UpdateTeacherResponse
	(*DeleteTeacherRequest)(nil),   // 9: edusync_grpc.DeleteTeacherRequest
	(*DeleteTeacherResponse)(nil),  // 10: edusync_grpc.DeleteTeacherResponse
	(*model.Model)(nil),            // 11: edusync_grpc.Model
	(*address.Address)(nil),        // 12: edusync_grpc.Address
}
var file_teacher_proto_depIdxs = []int32{
	11, // 0: edusync_grpc.Teacher.model:type_name -> edusync_grpc.Model
	12, // 1: edusync_grpc.Teacher.address:type_name -> edusync_grpc.Address
	0,  // 2: edusync_grpc.CreateTeacherRequest.teacher:type_name -> edusync_grpc.Teacher
	0,  // 3: edusync_grpc.CreateTeacherResponse.teacher:type_name -> edusync_grpc.Teacher
	0,  // 4: edusync_grpc.GetAllTeachersResponse.teachers:type_name -> edusync_grpc.Teacher
	0,  // 5: edusync_grpc.GetTeacherByIdResponse.teacher:type_name -> edusync_grpc.Teacher
	0,  // 6: edusync_grpc.UpdateTeacherRequest.teacher:type_name -> edusync_grpc.Teacher
	0,  // 7: edusync_grpc.UpdateTeacherResponse.teacher:type_name -> edusync_grpc.Teacher
	1,  // 8: edusync_grpc.TeacherService.CreateTeacher:input_type -> edusync_grpc.CreateTeacherRequest
	3,  // 9: edusync_grpc.TeacherService.GetAllTeachers:input_type -> edusync_grpc.GetAllTeachersRequest
	5,  // 10: edusync_grpc.TeacherService.GetTeacherById:input_type -> edusync_grpc.GetTeacherByIdRequest
	7,  // 11: edusync_grpc.TeacherService.UpdateTeacher:input_type -> edusync_grpc.UpdateTeacherRequest
	9,  // 12: edusync_grpc.TeacherService.DeleteTeacher:input_type -> edusync_grpc.DeleteTeacherRequest
	2,  // 13: edusync_grpc.TeacherService.CreateTeacher:output_type -> edusync_grpc.CreateTeacherResponse
	4,  // 14: edusync_grpc.TeacherService.GetAllTeachers:output_type -> edusync_grpc.GetAllTeachersResponse
	6,  // 15: edusync_grpc.TeacherService.GetTeacherById:output_type -> edusync_grpc.GetTeacherByIdResponse
	8,  // 16: edusync_grpc.TeacherService.UpdateTeacher:output_type -> edusync_grpc.UpdateTeacherResponse
	10, // 17: edusync_grpc.TeacherService.DeleteTeacher:output_type -> edusync_grpc.DeleteTeacherResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_teacher_proto_init() }
func file_teacher_proto_init() {
	if File_teacher_proto != nil {
		return
	}
	file_teacher_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teacher_proto_rawDesc), len(file_teacher_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teacher_proto_goTypes,
		DependencyIndexes: file_teacher_proto_depIdxs,
		MessageInfos:      file_teacher_proto_msgTypes,
	}.Build()
	File_teacher_proto = out.File
	file_teacher_proto_goTypes = nil
	file_teacher_proto_depIdxs = nil
}
