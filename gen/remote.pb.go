// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: remote.proto

package remote

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

type ProviderInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ModelName     string                 `protobuf:"bytes,1,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	ApiKey        string                 `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ProviderName  string                 `protobuf:"bytes,3,opt,name=provider_name,json=providerName,proto3" json:"provider_name,omitempty"`
	Env           map[string]string      `protobuf:"bytes,4,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProviderInfo) Reset() {
	*x = ProviderInfo{}
	mi := &file_remote_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProviderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderInfo) ProtoMessage() {}

func (x *ProviderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderInfo.ProtoReflect.Descriptor instead.
func (*ProviderInfo) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{0}
}

func (x *ProviderInfo) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *ProviderInfo) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *ProviderInfo) GetProviderName() string {
	if x != nil {
		return x.ProviderName
	}
	return ""
}

func (x *ProviderInfo) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

type GitHubInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ApiToken      string                 `protobuf:"bytes,1,opt,name=api_token,json=apiToken,proto3" json:"api_token,omitempty"`
	ApiUrl        string                 `protobuf:"bytes,2,opt,name=api_url,json=apiUrl,proto3" json:"api_url,omitempty"`
	Repo          string                 `protobuf:"bytes,3,opt,name=repo,proto3" json:"repo,omitempty"`
	FullRepoUrl   string                 `protobuf:"bytes,4,opt,name=full_repo_url,json=fullRepoUrl,proto3" json:"full_repo_url,omitempty"`
	PrNumber      int32                  `protobuf:"varint,5,opt,name=pr_number,json=prNumber,proto3" json:"pr_number,omitempty"`
	IssueNumber   int32                  `protobuf:"varint,6,opt,name=issue_number,json=issueNumber,proto3" json:"issue_number,omitempty"`
	BranchName    string                 `protobuf:"bytes,7,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GitHubInfo) Reset() {
	*x = GitHubInfo{}
	mi := &file_remote_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GitHubInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitHubInfo) ProtoMessage() {}

func (x *GitHubInfo) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitHubInfo.ProtoReflect.Descriptor instead.
func (*GitHubInfo) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{1}
}

func (x *GitHubInfo) GetApiToken() string {
	if x != nil {
		return x.ApiToken
	}
	return ""
}

func (x *GitHubInfo) GetApiUrl() string {
	if x != nil {
		return x.ApiUrl
	}
	return ""
}

func (x *GitHubInfo) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *GitHubInfo) GetFullRepoUrl() string {
	if x != nil {
		return x.FullRepoUrl
	}
	return ""
}

func (x *GitHubInfo) GetPrNumber() int32 {
	if x != nil {
		return x.PrNumber
	}
	return 0
}

func (x *GitHubInfo) GetIssueNumber() int32 {
	if x != nil {
		return x.IssueNumber
	}
	return 0
}

func (x *GitHubInfo) GetBranchName() string {
	if x != nil {
		return x.BranchName
	}
	return ""
}

type SlackInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChannelId     string                 `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SlackInfo) Reset() {
	*x = SlackInfo{}
	mi := &file_remote_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SlackInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlackInfo) ProtoMessage() {}

func (x *SlackInfo) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlackInfo.ProtoReflect.Descriptor instead.
func (*SlackInfo) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{2}
}

func (x *SlackInfo) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type ExecuteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SessionId     string                 `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Provider      *ProviderInfo          `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Github        *GitHubInfo            `protobuf:"bytes,3,opt,name=github,proto3,oneof" json:"github,omitempty"`
	Slack         *SlackInfo             `protobuf:"bytes,4,opt,name=slack,proto3,oneof" json:"slack,omitempty"`
	Instruction   string                 `protobuf:"bytes,5,opt,name=instruction,proto3" json:"instruction,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecuteTaskRequest) Reset() {
	*x = ExecuteTaskRequest{}
	mi := &file_remote_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteTaskRequest) ProtoMessage() {}

func (x *ExecuteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteTaskRequest.ProtoReflect.Descriptor instead.
func (*ExecuteTaskRequest) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{3}
}

func (x *ExecuteTaskRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ExecuteTaskRequest) GetProvider() *ProviderInfo {
	if x != nil {
		return x.Provider
	}
	return nil
}

func (x *ExecuteTaskRequest) GetGithub() *GitHubInfo {
	if x != nil {
		return x.Github
	}
	return nil
}

func (x *ExecuteTaskRequest) GetSlack() *SlackInfo {
	if x != nil {
		return x.Slack
	}
	return nil
}

func (x *ExecuteTaskRequest) GetInstruction() string {
	if x != nil {
		return x.Instruction
	}
	return ""
}

type ExecuteTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SessionId     string                 `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Stdout        string                 `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr        string                 `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Success       bool                   `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExecuteTaskResponse) Reset() {
	*x = ExecuteTaskResponse{}
	mi := &file_remote_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteTaskResponse) ProtoMessage() {}

func (x *ExecuteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteTaskResponse.ProtoReflect.Descriptor instead.
func (*ExecuteTaskResponse) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{4}
}

func (x *ExecuteTaskResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ExecuteTaskResponse) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *ExecuteTaskResponse) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *ExecuteTaskResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	mi := &file_remote_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{5}
}

type PingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	mi := &file_remote_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[6]
	if x != nil {
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
	return file_remote_proto_rawDescGZIP(), []int{6}
}

func (x *PingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_remote_proto protoreflect.FileDescriptor

const file_remote_proto_rawDesc = "" +
	"\n" +
	"\fremote.proto\x12\x06remote\"\xd4\x01\n" +
	"\fProviderInfo\x12\x1d\n" +
	"\n" +
	"model_name\x18\x01 \x01(\tR\tmodelName\x12\x17\n" +
	"\aapi_key\x18\x02 \x01(\tR\x06apiKey\x12#\n" +
	"\rprovider_name\x18\x03 \x01(\tR\fproviderName\x12/\n" +
	"\x03env\x18\x04 \x03(\v2\x1d.remote.ProviderInfo.EnvEntryR\x03env\x1a6\n" +
	"\bEnvEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xdb\x01\n" +
	"\n" +
	"GitHubInfo\x12\x1b\n" +
	"\tapi_token\x18\x01 \x01(\tR\bapiToken\x12\x17\n" +
	"\aapi_url\x18\x02 \x01(\tR\x06apiUrl\x12\x12\n" +
	"\x04repo\x18\x03 \x01(\tR\x04repo\x12\"\n" +
	"\rfull_repo_url\x18\x04 \x01(\tR\vfullRepoUrl\x12\x1b\n" +
	"\tpr_number\x18\x05 \x01(\x05R\bprNumber\x12!\n" +
	"\fissue_number\x18\x06 \x01(\x05R\vissueNumber\x12\x1f\n" +
	"\vbranch_name\x18\a \x01(\tR\n" +
	"branchName\"*\n" +
	"\tSlackInfo\x12\x1d\n" +
	"\n" +
	"channel_id\x18\x01 \x01(\tR\tchannelId\"\xfb\x01\n" +
	"\x12ExecuteTaskRequest\x12\x1d\n" +
	"\n" +
	"session_id\x18\x01 \x01(\tR\tsessionId\x120\n" +
	"\bprovider\x18\x02 \x01(\v2\x14.remote.ProviderInfoR\bprovider\x12/\n" +
	"\x06github\x18\x03 \x01(\v2\x12.remote.GitHubInfoH\x00R\x06github\x88\x01\x01\x12,\n" +
	"\x05slack\x18\x04 \x01(\v2\x11.remote.SlackInfoH\x01R\x05slack\x88\x01\x01\x12 \n" +
	"\vinstruction\x18\x05 \x01(\tR\vinstructionB\t\n" +
	"\a_githubB\b\n" +
	"\x06_slack\"~\n" +
	"\x13ExecuteTaskResponse\x12\x1d\n" +
	"\n" +
	"session_id\x18\x01 \x01(\tR\tsessionId\x12\x16\n" +
	"\x06stdout\x18\x02 \x01(\tR\x06stdout\x12\x16\n" +
	"\x06stderr\x18\x03 \x01(\tR\x06stderr\x12\x18\n" +
	"\asuccess\x18\x04 \x01(\bR\asuccess\"\r\n" +
	"\vPingRequest\"&\n" +
	"\fPingResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2\x93\x01\n" +
	"\x12RemoteAgentService\x12H\n" +
	"\vExecuteTask\x12\x1a.remote.ExecuteTaskRequest\x1a\x1b.remote.ExecuteTaskResponse\"\x00\x123\n" +
	"\x04Ping\x12\x13.remote.PingRequest\x1a\x14.remote.PingResponse\"\x00B\x80\x01\n" +
	"\n" +
	"com.remoteB\vRemoteProtoP\x01Z-github.com/kommon-ai/agent-connect/pkg/remote\xa2\x02\x03RXX\xaa\x02\x06Remote\xca\x02\x06Remote\xe2\x02\x12Remote\\GPBMetadata\xea\x02\x06Remoteb\x06proto3"

var (
	file_remote_proto_rawDescOnce sync.Once
	file_remote_proto_rawDescData []byte
)

func file_remote_proto_rawDescGZIP() []byte {
	file_remote_proto_rawDescOnce.Do(func() {
		file_remote_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_remote_proto_rawDesc), len(file_remote_proto_rawDesc)))
	})
	return file_remote_proto_rawDescData
}

var file_remote_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_remote_proto_goTypes = []any{
	(*ProviderInfo)(nil),        // 0: remote.ProviderInfo
	(*GitHubInfo)(nil),          // 1: remote.GitHubInfo
	(*SlackInfo)(nil),           // 2: remote.SlackInfo
	(*ExecuteTaskRequest)(nil),  // 3: remote.ExecuteTaskRequest
	(*ExecuteTaskResponse)(nil), // 4: remote.ExecuteTaskResponse
	(*PingRequest)(nil),         // 5: remote.PingRequest
	(*PingResponse)(nil),        // 6: remote.PingResponse
	nil,                         // 7: remote.ProviderInfo.EnvEntry
}
var file_remote_proto_depIdxs = []int32{
	7, // 0: remote.ProviderInfo.env:type_name -> remote.ProviderInfo.EnvEntry
	0, // 1: remote.ExecuteTaskRequest.provider:type_name -> remote.ProviderInfo
	1, // 2: remote.ExecuteTaskRequest.github:type_name -> remote.GitHubInfo
	2, // 3: remote.ExecuteTaskRequest.slack:type_name -> remote.SlackInfo
	3, // 4: remote.RemoteAgentService.ExecuteTask:input_type -> remote.ExecuteTaskRequest
	5, // 5: remote.RemoteAgentService.Ping:input_type -> remote.PingRequest
	4, // 6: remote.RemoteAgentService.ExecuteTask:output_type -> remote.ExecuteTaskResponse
	6, // 7: remote.RemoteAgentService.Ping:output_type -> remote.PingResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_remote_proto_init() }
func file_remote_proto_init() {
	if File_remote_proto != nil {
		return
	}
	file_remote_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_remote_proto_rawDesc), len(file_remote_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remote_proto_goTypes,
		DependencyIndexes: file_remote_proto_depIdxs,
		MessageInfos:      file_remote_proto_msgTypes,
	}.Build()
	File_remote_proto = out.File
	file_remote_proto_goTypes = nil
	file_remote_proto_depIdxs = nil
}
