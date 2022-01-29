// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: test.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Testing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *Testing) Reset() {
	*x = Testing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Testing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Testing) ProtoMessage() {}

func (x *Testing) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Testing.ProtoReflect.Descriptor instead.
func (*Testing) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *Testing) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Testing) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type TestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Testing *Testing `protobuf:"bytes,1,opt,name=testing,proto3" json:"testing,omitempty"`
}

func (x *TestRequest) Reset() {
	*x = TestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRequest) ProtoMessage() {}

func (x *TestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRequest.ProtoReflect.Descriptor instead.
func (*TestRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestRequest) GetTesting() *Testing {
	if x != nil {
		return x.Testing
	}
	return nil
}

type TestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *TestResponse) Reset() {
	*x = TestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResponse) ProtoMessage() {}

func (x *TestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResponse.ProtoReflect.Descriptor instead.
func (*TestResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type TestManyTimesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Testing *Testing `protobuf:"bytes,1,opt,name=testing,proto3" json:"testing,omitempty"`
}

func (x *TestManyTimesRequest) Reset() {
	*x = TestManyTimesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestManyTimesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestManyTimesRequest) ProtoMessage() {}

func (x *TestManyTimesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestManyTimesRequest.ProtoReflect.Descriptor instead.
func (*TestManyTimesRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3}
}

func (x *TestManyTimesRequest) GetTesting() *Testing {
	if x != nil {
		return x.Testing
	}
	return nil
}

type TestManyTimesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *TestManyTimesResponse) Reset() {
	*x = TestManyTimesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestManyTimesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestManyTimesResponse) ProtoMessage() {}

func (x *TestManyTimesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestManyTimesResponse.ProtoReflect.Descriptor instead.
func (*TestManyTimesResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{4}
}

func (x *TestManyTimesResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type LongGreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Testing *Testing `protobuf:"bytes,1,opt,name=testing,proto3" json:"testing,omitempty"`
}

func (x *LongGreetRequest) Reset() {
	*x = LongGreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongGreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongGreetRequest) ProtoMessage() {}

func (x *LongGreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongGreetRequest.ProtoReflect.Descriptor instead.
func (*LongGreetRequest) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{5}
}

func (x *LongGreetRequest) GetTesting() *Testing {
	if x != nil {
		return x.Testing
	}
	return nil
}

type LongGreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LongGreetResponse) Reset() {
	*x = LongGreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongGreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongGreetResponse) ProtoMessage() {}

func (x *LongGreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongGreetResponse.ProtoReflect.Descriptor instead.
func (*LongGreetResponse) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{6}
}

func (x *LongGreetResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x22, 0x45, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x0b, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x26, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3f, 0x0a, 0x14, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x2f, 0x0a, 0x15, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3b, 0x0a, 0x10, 0x4c,
	0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x07, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x2b, 0x0a, 0x11, 0x4c, 0x6f, 0x6e, 0x67,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xce, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x11, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61,
	0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d,
	0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x12, 0x16, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_test_proto_goTypes = []interface{}{
	(*Testing)(nil),               // 0: test.Testing
	(*TestRequest)(nil),           // 1: test.TestRequest
	(*TestResponse)(nil),          // 2: test.TestResponse
	(*TestManyTimesRequest)(nil),  // 3: test.TestManyTimesRequest
	(*TestManyTimesResponse)(nil), // 4: test.TestManyTimesResponse
	(*LongGreetRequest)(nil),      // 5: test.LongGreetRequest
	(*LongGreetResponse)(nil),     // 6: test.LongGreetResponse
}
var file_test_proto_depIdxs = []int32{
	0, // 0: test.TestRequest.testing:type_name -> test.Testing
	0, // 1: test.TestManyTimesRequest.testing:type_name -> test.Testing
	0, // 2: test.LongGreetRequest.testing:type_name -> test.Testing
	1, // 3: test.TestService.Test:input_type -> test.TestRequest
	3, // 4: test.TestService.TestManyTimes:input_type -> test.TestManyTimesRequest
	5, // 5: test.TestService.LongGreet:input_type -> test.LongGreetRequest
	2, // 6: test.TestService.Test:output_type -> test.TestResponse
	4, // 7: test.TestService.TestManyTimes:output_type -> test.TestManyTimesResponse
	6, // 8: test.TestService.LongGreet:output_type -> test.LongGreetResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Testing); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRequest); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResponse); i {
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
		file_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestManyTimesRequest); i {
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
		file_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestManyTimesResponse); i {
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
		file_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongGreetRequest); i {
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
		file_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongGreetResponse); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	// unary
	Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
	// Server Streaming
	TestManyTimes(ctx context.Context, in *TestManyTimesRequest, opts ...grpc.CallOption) (TestService_TestManyTimesClient, error)
	// Client Streaming
	LongGreet(ctx context.Context, opts ...grpc.CallOption) (TestService_LongGreetClient, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/test.TestService/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) TestManyTimes(ctx context.Context, in *TestManyTimesRequest, opts ...grpc.CallOption) (TestService_TestManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[0], "/test.TestService/TestManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceTestManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_TestManyTimesClient interface {
	Recv() (*TestManyTimesResponse, error)
	grpc.ClientStream
}

type testServiceTestManyTimesClient struct {
	grpc.ClientStream
}

func (x *testServiceTestManyTimesClient) Recv() (*TestManyTimesResponse, error) {
	m := new(TestManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) LongGreet(ctx context.Context, opts ...grpc.CallOption) (TestService_LongGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestService_serviceDesc.Streams[1], "/test.TestService/LongGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceLongGreetClient{stream}
	return x, nil
}

type TestService_LongGreetClient interface {
	Send(*LongGreetRequest) error
	CloseAndRecv() (*LongGreetResponse, error)
	grpc.ClientStream
}

type testServiceLongGreetClient struct {
	grpc.ClientStream
}

func (x *testServiceLongGreetClient) Send(m *LongGreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceLongGreetClient) CloseAndRecv() (*LongGreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LongGreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	// unary
	Test(context.Context, *TestRequest) (*TestResponse, error)
	// Server Streaming
	TestManyTimes(*TestManyTimesRequest, TestService_TestManyTimesServer) error
	// Client Streaming
	LongGreet(TestService_LongGreetServer) error
}

// UnimplementedTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (*UnimplementedTestServiceServer) Test(context.Context, *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (*UnimplementedTestServiceServer) TestManyTimes(*TestManyTimesRequest, TestService_TestManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method TestManyTimes not implemented")
}
func (*UnimplementedTestServiceServer) LongGreet(TestService_LongGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method LongGreet not implemented")
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestService/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Test(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_TestManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TestManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).TestManyTimes(m, &testServiceTestManyTimesServer{stream})
}

type TestService_TestManyTimesServer interface {
	Send(*TestManyTimesResponse) error
	grpc.ServerStream
}

type testServiceTestManyTimesServer struct {
	grpc.ServerStream
}

func (x *testServiceTestManyTimesServer) Send(m *TestManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_LongGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).LongGreet(&testServiceLongGreetServer{stream})
}

type TestService_LongGreetServer interface {
	SendAndClose(*LongGreetResponse) error
	Recv() (*LongGreetRequest, error)
	grpc.ServerStream
}

type testServiceLongGreetServer struct {
	grpc.ServerStream
}

func (x *testServiceLongGreetServer) SendAndClose(m *LongGreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceLongGreetServer) Recv() (*LongGreetRequest, error) {
	m := new(LongGreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _TestService_Test_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestManyTimes",
			Handler:       _TestService_TestManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LongGreet",
			Handler:       _TestService_LongGreet_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}