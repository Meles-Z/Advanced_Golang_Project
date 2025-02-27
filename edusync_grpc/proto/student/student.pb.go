// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: student.proto

package student

import (
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

type Student struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Model         *model.Model           `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber   string                 `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	CourseIds     []string               `protobuf:"bytes,5,rep,name=course_ids,json=courseIds,proto3" json:"course_ids,omitempty"`                // Use IDs instead of importing Course
	DepartmentId  *string                `protobuf:"bytes,6,opt,name=department_id,json=departmentId,proto3,oneof" json:"department_id,omitempty"` // Use ID instead of importing Department
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Student) Reset() {
	*x = Student{}
	mi := &file_student_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetModel() *model.Model {
	if x != nil {
		return x.Model
	}
	return nil
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Student) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Student) GetCourseIds() []string {
	if x != nil {
		return x.CourseIds
	}
	return nil
}

func (x *Student) GetDepartmentId() string {
	if x != nil && x.DepartmentId != nil {
		return *x.DepartmentId
	}
	return ""
}

type CreateStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Student       *Student               `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateStudentRequest) Reset() {
	*x = CreateStudentRequest{}
	mi := &file_student_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentRequest) ProtoMessage() {}

func (x *CreateStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentRequest.ProtoReflect.Descriptor instead.
func (*CreateStudentRequest) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStudentRequest) GetStudent() *Student {
	if x != nil {
		return x.Student
	}
	return nil
}

type CreateStudentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Student       *Student               `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateStudentResponse) Reset() {
	*x = CreateStudentResponse{}
	mi := &file_student_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentResponse) ProtoMessage() {}

func (x *CreateStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentResponse.ProtoReflect.Descriptor instead.
func (*CreateStudentResponse) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStudentResponse) GetStudent() *Student {
	if x != nil {
		return x.Student
	}
	return nil
}

type GetAllStudentsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllStudentsRequest) Reset() {
	*x = GetAllStudentsRequest{}
	mi := &file_student_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllStudentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStudentsRequest) ProtoMessage() {}

func (x *GetAllStudentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStudentsRequest.ProtoReflect.Descriptor instead.
func (*GetAllStudentsRequest) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{3}
}

type GetAllStudentsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Students      []*Student             `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllStudentsResponse) Reset() {
	*x = GetAllStudentsResponse{}
	mi := &file_student_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllStudentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllStudentsResponse) ProtoMessage() {}

func (x *GetAllStudentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllStudentsResponse.ProtoReflect.Descriptor instead.
func (*GetAllStudentsResponse) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllStudentsResponse) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

type GetStudentByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStudentByIdRequest) Reset() {
	*x = GetStudentByIdRequest{}
	mi := &file_student_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStudentByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentByIdRequest) ProtoMessage() {}

func (x *GetStudentByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentByIdRequest.ProtoReflect.Descriptor instead.
func (*GetStudentByIdRequest) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{5}
}

func (x *GetStudentByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetStudentByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Student       *Student               `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStudentByIdResponse) Reset() {
	*x = GetStudentByIdResponse{}
	mi := &file_student_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStudentByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentByIdResponse) ProtoMessage() {}

func (x *GetStudentByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentByIdResponse.ProtoReflect.Descriptor instead.
func (*GetStudentByIdResponse) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{6}
}

func (x *GetStudentByIdResponse) GetStudent() *Student {
	if x != nil {
		return x.Student
	}
	return nil
}

type UpdateStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Student       *Student               `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStudentRequest) Reset() {
	*x = UpdateStudentRequest{}
	mi := &file_student_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentRequest) ProtoMessage() {}

func (x *UpdateStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentRequest.ProtoReflect.Descriptor instead.
func (*UpdateStudentRequest) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateStudentRequest) GetStudent() *Student {
	if x != nil {
		return x.Student
	}
	return nil
}

type UpdateStudentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Student       *Student               `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStudentResponse) Reset() {
	*x = UpdateStudentResponse{}
	mi := &file_student_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentResponse) ProtoMessage() {}

func (x *UpdateStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentResponse.ProtoReflect.Descriptor instead.
func (*UpdateStudentResponse) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateStudentResponse) GetStudent() *Student {
	if x != nil {
		return x.Student
	}
	return nil
}

type DeleteStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteStudentRequest) Reset() {
	*x = DeleteStudentRequest{}
	mi := &file_student_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentRequest) ProtoMessage() {}

func (x *DeleteStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentRequest.ProtoReflect.Descriptor instead.
func (*DeleteStudentRequest) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteStudentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteStudentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteStudentResponse) Reset() {
	*x = DeleteStudentResponse{}
	mi := &file_student_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentResponse) ProtoMessage() {}

func (x *DeleteStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_student_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentResponse.ProtoReflect.Descriptor instead.
func (*DeleteStudentResponse) Descriptor() ([]byte, []int) {
	return file_student_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteStudentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_student_proto protoreflect.FileDescriptor

var file_student_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0b, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x07, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x73, 0x12, 0x28, 0x0a,
	0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2f, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x22, 0x48, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65,
	0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x17, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x47, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22,
	0x48, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x64, 0x75, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x32, 0xcb, 0x03, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x65,
	0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x30,
	0x01, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x23, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12,
	0x22, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x65, 0x64, 0x75, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x65, 0x6c, 0x65, 0x73, 0x2d, 0x7a, 0x2f, 0x65, 0x64, 0x75, 0x73, 0x79, 0x73, 0x6e,
	0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_student_proto_rawDescOnce sync.Once
	file_student_proto_rawDescData []byte
)

func file_student_proto_rawDescGZIP() []byte {
	file_student_proto_rawDescOnce.Do(func() {
		file_student_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_student_proto_rawDesc), len(file_student_proto_rawDesc)))
	})
	return file_student_proto_rawDescData
}

var file_student_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_student_proto_goTypes = []any{
	(*Student)(nil),                // 0: edusync_grpc.Student
	(*CreateStudentRequest)(nil),   // 1: edusync_grpc.CreateStudentRequest
	(*CreateStudentResponse)(nil),  // 2: edusync_grpc.CreateStudentResponse
	(*GetAllStudentsRequest)(nil),  // 3: edusync_grpc.GetAllStudentsRequest
	(*GetAllStudentsResponse)(nil), // 4: edusync_grpc.GetAllStudentsResponse
	(*GetStudentByIdRequest)(nil),  // 5: edusync_grpc.GetStudentByIdRequest
	(*GetStudentByIdResponse)(nil), // 6: edusync_grpc.GetStudentByIdResponse
	(*UpdateStudentRequest)(nil),   // 7: edusync_grpc.UpdateStudentRequest
	(*UpdateStudentResponse)(nil),  // 8: edusync_grpc.UpdateStudentResponse
	(*DeleteStudentRequest)(nil),   // 9: edusync_grpc.DeleteStudentRequest
	(*DeleteStudentResponse)(nil),  // 10: edusync_grpc.DeleteStudentResponse
	(*model.Model)(nil),            // 11: edusync_grpc.Model
}
var file_student_proto_depIdxs = []int32{
	11, // 0: edusync_grpc.Student.model:type_name -> edusync_grpc.Model
	0,  // 1: edusync_grpc.CreateStudentRequest.student:type_name -> edusync_grpc.Student
	0,  // 2: edusync_grpc.CreateStudentResponse.student:type_name -> edusync_grpc.Student
	0,  // 3: edusync_grpc.GetAllStudentsResponse.students:type_name -> edusync_grpc.Student
	0,  // 4: edusync_grpc.GetStudentByIdResponse.student:type_name -> edusync_grpc.Student
	0,  // 5: edusync_grpc.UpdateStudentRequest.student:type_name -> edusync_grpc.Student
	0,  // 6: edusync_grpc.UpdateStudentResponse.student:type_name -> edusync_grpc.Student
	1,  // 7: edusync_grpc.StudentService.CreateStudent:input_type -> edusync_grpc.CreateStudentRequest
	3,  // 8: edusync_grpc.StudentService.GetAllStudents:input_type -> edusync_grpc.GetAllStudentsRequest
	5,  // 9: edusync_grpc.StudentService.GetStudentById:input_type -> edusync_grpc.GetStudentByIdRequest
	7,  // 10: edusync_grpc.StudentService.UpdateStudent:input_type -> edusync_grpc.UpdateStudentRequest
	9,  // 11: edusync_grpc.StudentService.DeleteStudent:input_type -> edusync_grpc.DeleteStudentRequest
	2,  // 12: edusync_grpc.StudentService.CreateStudent:output_type -> edusync_grpc.CreateStudentResponse
	0,  // 13: edusync_grpc.StudentService.GetAllStudents:output_type -> edusync_grpc.Student
	6,  // 14: edusync_grpc.StudentService.GetStudentById:output_type -> edusync_grpc.GetStudentByIdResponse
	8,  // 15: edusync_grpc.StudentService.UpdateStudent:output_type -> edusync_grpc.UpdateStudentResponse
	10, // 16: edusync_grpc.StudentService.DeleteStudent:output_type -> edusync_grpc.DeleteStudentResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_student_proto_init() }
func file_student_proto_init() {
	if File_student_proto != nil {
		return
	}
	file_student_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_student_proto_rawDesc), len(file_student_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_student_proto_goTypes,
		DependencyIndexes: file_student_proto_depIdxs,
		MessageInfos:      file_student_proto_msgTypes,
	}.Build()
	File_student_proto = out.File
	file_student_proto_goTypes = nil
	file_student_proto_depIdxs = nil
}
