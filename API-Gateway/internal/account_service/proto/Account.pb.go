// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: Account.proto

package proto

import (
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

type AccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IdAccount     string                 `protobuf:"bytes,1,opt,name=IdAccount,proto3" json:"IdAccount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccountRequest) Reset() {
	*x = AccountRequest{}
	mi := &file_Account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountRequest) ProtoMessage() {}

func (x *AccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountRequest.ProtoReflect.Descriptor instead.
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{0}
}

func (x *AccountRequest) GetIdAccount() string {
	if x != nil {
		return x.IdAccount
	}
	return ""
}

type AccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IdAccount     string                 `protobuf:"bytes,1,opt,name=IdAccount,proto3" json:"IdAccount,omitempty"`
	Credits       int64                  `protobuf:"varint,2,opt,name=credits,proto3" json:"credits,omitempty"`
	CreationDate  string                 `protobuf:"bytes,3,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccountResponse) Reset() {
	*x = AccountResponse{}
	mi := &file_Account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountResponse) ProtoMessage() {}

func (x *AccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountResponse.ProtoReflect.Descriptor instead.
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountResponse) GetIdAccount() string {
	if x != nil {
		return x.IdAccount
	}
	return ""
}

func (x *AccountResponse) GetCredits() int64 {
	if x != nil {
		return x.Credits
	}
	return 0
}

func (x *AccountResponse) GetCreationDate() string {
	if x != nil {
		return x.CreationDate
	}
	return ""
}

type UpdateAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IdAccount     string                 `protobuf:"bytes,1,opt,name=IdAccount,proto3" json:"IdAccount,omitempty"`
	Credits       int64                  `protobuf:"varint,2,opt,name=credits,proto3" json:"credits,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAccountRequest) Reset() {
	*x = UpdateAccountRequest{}
	mi := &file_Account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountRequest) ProtoMessage() {}

func (x *UpdateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAccountRequest) GetIdAccount() string {
	if x != nil {
		return x.IdAccount
	}
	return ""
}

func (x *UpdateAccountRequest) GetCredits() int64 {
	if x != nil {
		return x.Credits
	}
	return 0
}

type UpdateAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAccountResponse) Reset() {
	*x = UpdateAccountResponse{}
	mi := &file_Account_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountResponse) ProtoMessage() {}

func (x *UpdateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountResponse.ProtoReflect.Descriptor instead.
func (*UpdateAccountResponse) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAccountResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateAccountResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IdAccount     string                 `protobuf:"bytes,1,opt,name=IdAccount,proto3" json:"IdAccount,omitempty"`
	Credits       float64                `protobuf:"fixed64,2,opt,name=credits,proto3" json:"credits,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	mi := &file_Account_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{4}
}

func (x *CreateAccountRequest) GetIdAccount() string {
	if x != nil {
		return x.IdAccount
	}
	return ""
}

func (x *CreateAccountRequest) GetCredits() float64 {
	if x != nil {
		return x.Credits
	}
	return 0
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	mi := &file_Account_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Account_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_Account_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAccountResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateAccountResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_Account_proto protoreflect.FileDescriptor

const file_Account_proto_rawDesc = "" +
	"\n" +
	"\rAccount.proto\x12\aAccount\".\n" +
	"\x0eAccountRequest\x12\x1c\n" +
	"\tIdAccount\x18\x01 \x01(\tR\tIdAccount\"m\n" +
	"\x0fAccountResponse\x12\x1c\n" +
	"\tIdAccount\x18\x01 \x01(\tR\tIdAccount\x12\x18\n" +
	"\acredits\x18\x02 \x01(\x03R\acredits\x12\"\n" +
	"\fcreationDate\x18\x03 \x01(\tR\fcreationDate\"N\n" +
	"\x14UpdateAccountRequest\x12\x1c\n" +
	"\tIdAccount\x18\x01 \x01(\tR\tIdAccount\x12\x18\n" +
	"\acredits\x18\x02 \x01(\x03R\acredits\"K\n" +
	"\x15UpdateAccountResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"N\n" +
	"\x14CreateAccountRequest\x12\x1c\n" +
	"\tIdAccount\x18\x01 \x01(\tR\tIdAccount\x12\x18\n" +
	"\acredits\x18\x02 \x01(\x01R\acredits\"K\n" +
	"\x15CreateAccountResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2\xf8\x01\n" +
	"\aAccount\x12F\n" +
	"\x11getAccountBalance\x12\x17.Account.AccountRequest\x1a\x18.Account.AccountResponse\x12U\n" +
	"\x14updateAccountBalance\x12\x1d.Account.UpdateAccountRequest\x1a\x1e.Account.UpdateAccountResponse\x12N\n" +
	"\rcreateAccount\x12\x1d.Account.CreateAccountRequest\x1a\x1e.Account.CreateAccountResponseBq\n" +
	"\x10com.account.grpcB\x0eExchangeSystemP\x01ZKgithub.com/JeroZp/gRPC-MOM/API-Gateway/internal/account_service/proto;protob\x06proto3"

var (
	file_Account_proto_rawDescOnce sync.Once
	file_Account_proto_rawDescData []byte
)

func file_Account_proto_rawDescGZIP() []byte {
	file_Account_proto_rawDescOnce.Do(func() {
		file_Account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_Account_proto_rawDesc), len(file_Account_proto_rawDesc)))
	})
	return file_Account_proto_rawDescData
}

var file_Account_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Account_proto_goTypes = []any{
	(*AccountRequest)(nil),        // 0: Account.AccountRequest
	(*AccountResponse)(nil),       // 1: Account.AccountResponse
	(*UpdateAccountRequest)(nil),  // 2: Account.UpdateAccountRequest
	(*UpdateAccountResponse)(nil), // 3: Account.UpdateAccountResponse
	(*CreateAccountRequest)(nil),  // 4: Account.CreateAccountRequest
	(*CreateAccountResponse)(nil), // 5: Account.CreateAccountResponse
}
var file_Account_proto_depIdxs = []int32{
	0, // 0: Account.Account.getAccountBalance:input_type -> Account.AccountRequest
	2, // 1: Account.Account.updateAccountBalance:input_type -> Account.UpdateAccountRequest
	4, // 2: Account.Account.createAccount:input_type -> Account.CreateAccountRequest
	1, // 3: Account.Account.getAccountBalance:output_type -> Account.AccountResponse
	3, // 4: Account.Account.updateAccountBalance:output_type -> Account.UpdateAccountResponse
	5, // 5: Account.Account.createAccount:output_type -> Account.CreateAccountResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Account_proto_init() }
func file_Account_proto_init() {
	if File_Account_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_Account_proto_rawDesc), len(file_Account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Account_proto_goTypes,
		DependencyIndexes: file_Account_proto_depIdxs,
		MessageInfos:      file_Account_proto_msgTypes,
	}.Build()
	File_Account_proto = out.File
	file_Account_proto_goTypes = nil
	file_Account_proto_depIdxs = nil
}
