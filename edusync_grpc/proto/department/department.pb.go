// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: department.proto

package department

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

type Department struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Model         *model.Model           `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber   string                 `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	AddressId     *string                `protobuf:"bytes,5,opt,name=address_id,json=addressId,proto3,oneof" json:"address_id,omitempty"` // Use ID instead of importing Address
	TeacherIds    []string               `protobuf:"bytes,6,rep,name=teacher_ids,json=teacherIds,proto3" json:"teacher_ids,omitempty"`    // Use IDs instead of importing Teacher
	StudentIds    []string               `protobuf:"bytes,7,rep,name=student_ids,json=studentIds,proto3" json:"student_ids,omitempty"`    // Use IDs instead of importing Student
	CourseIds     []string               `protobuf:"bytes,8,rep,name=course_ids,json=courseIds,proto3" json:"course_ids,omitempty"`       // Use IDs instead of importing Course
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Department) Reset() {
	*x = Department{}
	mi := &file_department_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Department) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Department) ProtoMessage() {}

func (x *Department) ProtoReflect() protoreflect.Message {
	mi := &file_department_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Department.ProtoReflect.Descriptor instead.
func (*Department) Descriptor() ([]byte, []int) {
	return file_department_proto_rawDescGZIP(), []int{0}
}

func (x *Department) GetModel() *model.Model {
	if x != nil {
		return x.Model
	}
	return nil
}

func (x *Department) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Department) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Department) GetAddressId() string {
	if x != nil && x.AddressId != nil {
		return *x.AddressId
	}
	return ""
}

func (x *Department) GetTeacherIds() []string {
	if x != nil {
		return x.TeacherIds
	}
	return nil
}

func (x *Department) GetStudentIds() []string {
	if x != nil {
		return x.StudentIds
	}
	return nil
}

func (x *Department) GetCourseIds() []string {
	if x != nil {
		return x.CourseIds
	}
	return nil
}

type CreateDepartmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Department    *Department            `protobuf:"bytes,1,opt,name=department,proto3" json:"department,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDepartmentRequest) Reset() {
	*x = CreateDepartmentRequest{}
	mi := &file_department_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDepartmentRequest) ProtoMessage() {}

func (x *CreateDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_department_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDepartmentRequest.ProtoReflect.Descriptor instead.
func (*CreateDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_department_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDepartmentRequest) GetDepartment() *Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type CreateDepartmentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Department    *Department            `protobuf:"bytes,1,opt,name=department,proto3" json:"department,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDepartmentResponse) Reset() {
	*x = CreateDepartmentResponse{}
	mi := &file_department_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDepartmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDepartmentResponse) ProtoMessage() {}

func (x *CreateDepartmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_department_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDepartmentResponse.ProtoReflect.Descriptor instead.
func (*CreateDepartmentResponse) Descriptor() ([]byte, []int) {
	return file_department_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDepartmentResponse) GetDepartment() *Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type GetAllDepartmentsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllDepartmentsRequest) Reset() {
	*x = GetAllDepartmentsRequest{}
	mi := &file_department_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllDepartmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDepartmentsRequest) ProtoMessage() {}

func (x *GetAllDepartmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_department_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDepartmentsRequest.ProtoReflect.Descriptor instead.
func (*GetAllDepartmentsRequest) Descriptor() ([]byte, []int) {
	return file_department_proto_rawDescGZIP(), []int{3}
}

type GetAllDepartmentsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Departments   []*Department          `protobuf:"bytes,1,rep,name=departments,proto3" json:"departments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllDepartmentsResponse) Reset() {
	*x = GetAllDepartmentsResponse{}
	mi := &file_department_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllDepartmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDepartmentsResponse) ProtoMessage() {}

func (x *GetAllDepartmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_department_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDepartmentsResponse.ProtoReflect.Descriptor instead.
func (*GetAllDepartmentsResponse) Descriptor() ([]byte, []int) {
	return file_department_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllDepartmentsResponse) GetDepartments() []*Department {
	if x != nil {
		return x.Departments
	}
	return nil
}

type GetDepartmentByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDepartmentByIdRequest) Reset() {
	*x = GetDepartmentByIdRequest{}
	mi := &file_department_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDepartmentByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDepartmentByIdRequest) ProtoMessage() {}

func (x *GetDepartmentByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_department_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDepartmentByIdRequest.ProtoReflect.Descriptor instead.
func (*GetDepartmentByIdRequest) Descriptor() ([]byte, []int) {
	return file_department_proto_rawDescGZIP(), []int{5}
}

func (x *GetDepartmentByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDepartmentByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Department    *Department            `protobuf:"bytes,1,opt,name=department,proto3" json:"department,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDepartmentByIdResponse) Reset() {
	*x = GetDepartmentByIdResponse{}
	mi := &file_department_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDepartmentByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDepartmentByIdResponse) ProtoMessage() {}

func (x *GetDepartmentByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_department_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDepartmentByIdResponse.ProtoReflect.Descriptor instead.
func (*GetDepartmentByIdResponse) Descriptor() ([]byte, []int) {
	return file_department_proto_rawDescGZIP(), []int{6}
}

func (x *GetDepartmentByIdResponse) GetDepartment() *Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type UpdateDepartmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Department    *Department            `protobuf:"bytes,1,opt,name=department,proto3" json:"department,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDepartmentRequest) Reset() {
	*x = UpdateDepartmentRequest{}
	mi := &file_department_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDepartmentRequest) ProtoMessage() {}

func (x *UpdateDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_department_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDepartmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_department_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateDepartmentRequest) GetDepartment() *Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type UpdateDepartmentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Department    *Department            `protobuf:"bytes,1,opt,name=department,proto3" json:"department,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDepartmentResponse) Reset() {
	*x = UpdateDepartmentResponse{}
	mi := &file_department_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDepartmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDepartmentResponse) ProtoMessage() {}

func (x *UpdateDepartmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_department_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDepartmentResponse.ProtoReflect.Descriptor instead.
func (*UpdateDepartmentResponse) Descriptor() ([]byte, []int) {
	return file_department_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateDepartmentResponse) GetDepartment() *Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type DeleteDepartmentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDepartmentRequest) Reset() {
	*x = DeleteDepartmentRequest{}
	mi := &file_department_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDepartmentRequest) ProtoMessage() {}

func (x *DeleteDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_department_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDepartmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_department_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteDepartmentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteDepartmentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDepartmentResponse) Reset() {
	*x = DeleteDepartmentResponse{}
	mi := &file_department_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDepartmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDepartmentResponse) ProtoMessage() {}

func (x *DeleteDepartmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_department_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDepartmentResponse.ProtoReflect.Descriptor instead.
func (*DeleteDepartmentResponse) Descriptor() ([]byte, []int) {
	return file_department_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteDepartmentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_department_proto protoreflect.FileDescriptor

var file_department_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x02,
	0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x64,
	0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x22, 0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x49, 0x64, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x38, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x54, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x64, 0x75, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x57, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x55, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x54, 0x0a, 0x18,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x65,
	0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a,
	0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0x88, 0x04, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e,
	0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x26, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x64, 0x75, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x64, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x26, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x65,
	0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x25, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x65, 0x64, 0x75, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x6c,
	0x65, 0x73, 0x2d, 0x7a, 0x2f, 0x65, 0x64, 0x75, 0x73, 0x79, 0x73, 0x6e, 0x63, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_department_proto_rawDescOnce sync.Once
	file_department_proto_rawDescData []byte
)

func file_department_proto_rawDescGZIP() []byte {
	file_department_proto_rawDescOnce.Do(func() {
		file_department_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_department_proto_rawDesc), len(file_department_proto_rawDesc)))
	})
	return file_department_proto_rawDescData
}

var file_department_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_department_proto_goTypes = []any{
	(*Department)(nil),                // 0: edusync_grpc.Department
	(*CreateDepartmentRequest)(nil),   // 1: edusync_grpc.CreateDepartmentRequest
	(*CreateDepartmentResponse)(nil),  // 2: edusync_grpc.CreateDepartmentResponse
	(*GetAllDepartmentsRequest)(nil),  // 3: edusync_grpc.GetAllDepartmentsRequest
	(*GetAllDepartmentsResponse)(nil), // 4: edusync_grpc.GetAllDepartmentsResponse
	(*GetDepartmentByIdRequest)(nil),  // 5: edusync_grpc.GetDepartmentByIdRequest
	(*GetDepartmentByIdResponse)(nil), // 6: edusync_grpc.GetDepartmentByIdResponse
	(*UpdateDepartmentRequest)(nil),   // 7: edusync_grpc.UpdateDepartmentRequest
	(*UpdateDepartmentResponse)(nil),  // 8: edusync_grpc.UpdateDepartmentResponse
	(*DeleteDepartmentRequest)(nil),   // 9: edusync_grpc.DeleteDepartmentRequest
	(*DeleteDepartmentResponse)(nil),  // 10: edusync_grpc.DeleteDepartmentResponse
	(*model.Model)(nil),               // 11: edusync_grpc.Model
}
var file_department_proto_depIdxs = []int32{
	11, // 0: edusync_grpc.Department.model:type_name -> edusync_grpc.Model
	0,  // 1: edusync_grpc.CreateDepartmentRequest.department:type_name -> edusync_grpc.Department
	0,  // 2: edusync_grpc.CreateDepartmentResponse.department:type_name -> edusync_grpc.Department
	0,  // 3: edusync_grpc.GetAllDepartmentsResponse.departments:type_name -> edusync_grpc.Department
	0,  // 4: edusync_grpc.GetDepartmentByIdResponse.department:type_name -> edusync_grpc.Department
	0,  // 5: edusync_grpc.UpdateDepartmentRequest.department:type_name -> edusync_grpc.Department
	0,  // 6: edusync_grpc.UpdateDepartmentResponse.department:type_name -> edusync_grpc.Department
	1,  // 7: edusync_grpc.DepartmentService.CreateDepartment:input_type -> edusync_grpc.CreateDepartmentRequest
	3,  // 8: edusync_grpc.DepartmentService.GetAllDepartments:input_type -> edusync_grpc.GetAllDepartmentsRequest
	5,  // 9: edusync_grpc.DepartmentService.GetDepartmentById:input_type -> edusync_grpc.GetDepartmentByIdRequest
	7,  // 10: edusync_grpc.DepartmentService.UpdateDepartment:input_type -> edusync_grpc.UpdateDepartmentRequest
	9,  // 11: edusync_grpc.DepartmentService.DeleteDepartment:input_type -> edusync_grpc.DeleteDepartmentRequest
	2,  // 12: edusync_grpc.DepartmentService.CreateDepartment:output_type -> edusync_grpc.CreateDepartmentResponse
	4,  // 13: edusync_grpc.DepartmentService.GetAllDepartments:output_type -> edusync_grpc.GetAllDepartmentsResponse
	6,  // 14: edusync_grpc.DepartmentService.GetDepartmentById:output_type -> edusync_grpc.GetDepartmentByIdResponse
	8,  // 15: edusync_grpc.DepartmentService.UpdateDepartment:output_type -> edusync_grpc.UpdateDepartmentResponse
	10, // 16: edusync_grpc.DepartmentService.DeleteDepartment:output_type -> edusync_grpc.DeleteDepartmentResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_department_proto_init() }
func file_department_proto_init() {
	if File_department_proto != nil {
		return
	}
	file_department_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_department_proto_rawDesc), len(file_department_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_department_proto_goTypes,
		DependencyIndexes: file_department_proto_depIdxs,
		MessageInfos:      file_department_proto_msgTypes,
	}.Build()
	File_department_proto = out.File
	file_department_proto_goTypes = nil
	file_department_proto_depIdxs = nil
}
