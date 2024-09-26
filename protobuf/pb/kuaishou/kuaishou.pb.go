// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: kuaishou/kuaishou.proto

package kuaishou

import (
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

// 定义请求参数
type RtaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId     string             `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`             //本次请求的唯一标识 id。
	Channel       string             `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`                                  //渠道标识
	RequestTime   int64              `protobuf:"varint,3,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`      //发起调用时的 unix 时间戳，毫秒。快手侧会与服务器当前时间进行比对，两者差值不能大于 10 分钟
	Sign          string             `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`                                        //验签值，sign = md5(request_id + request_time + 授权码)，三个字段直接拼接，中间没有"+"号
	Device        *RtaRequest_Device `protobuf:"bytes,5,opt,name=device,proto3" json:"device,omitempty"`                                    //设备信息
	PromotionType []string           `protobuf:"bytes,6,rep,name=promotion_type,json=promotionType,proto3" json:"promotion_type,omitempty"` //该设备需要问询的推广类型，如快手拉新、快手极速版拉活。具体值在线下约定快手内部文档请勿外传
}

func (x *RtaRequest) Reset() {
	*x = RtaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuaishou_kuaishou_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RtaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RtaRequest) ProtoMessage() {}

func (x *RtaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kuaishou_kuaishou_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RtaRequest.ProtoReflect.Descriptor instead.
func (*RtaRequest) Descriptor() ([]byte, []int) {
	return file_kuaishou_kuaishou_proto_rawDescGZIP(), []int{0}
}

func (x *RtaRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *RtaRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *RtaRequest) GetRequestTime() int64 {
	if x != nil {
		return x.RequestTime
	}
	return 0
}

func (x *RtaRequest) GetSign() string {
	if x != nil {
		return x.Sign
	}
	return ""
}

func (x *RtaRequest) GetDevice() *RtaRequest_Device {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *RtaRequest) GetPromotionType() []string {
	if x != nil {
		return x.PromotionType
	}
	return nil
}

// 定义响应参数
type RtaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode      int32                          `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 服务状态码，正常:0，异常:非 0
	PromotionResult []*RtaResponse_PromotionResult `protobuf:"bytes,2,rep,name=promotion_result,json=promotionResult,proto3" json:"promotion_result,omitempty"`
}

func (x *RtaResponse) Reset() {
	*x = RtaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuaishou_kuaishou_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RtaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RtaResponse) ProtoMessage() {}

func (x *RtaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kuaishou_kuaishou_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RtaResponse.ProtoReflect.Descriptor instead.
func (*RtaResponse) Descriptor() ([]byte, []int) {
	return file_kuaishou_kuaishou_proto_rawDescGZIP(), []int{1}
}

func (x *RtaResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RtaResponse) GetPromotionResult() []*RtaResponse_PromotionResult {
	if x != nil {
		return x.PromotionResult
	}
	return nil
}

type RtaRequest_Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imei    string `protobuf:"bytes,1,opt,name=imei,proto3" json:"imei,omitempty"`       //imei 原值
	ImeiMd5 string `protobuf:"bytes,2,opt,name=imeiMd5,proto3" json:"imeiMd5,omitempty"` //imeiMd5 = toLowerCase(md5(imei 原值))
	Oaid    string `protobuf:"bytes,3,opt,name=oaid,proto3" json:"oaid,omitempty"`       //oaid 原值
	OaidMd5 string `protobuf:"bytes,4,opt,name=oaidMd5,proto3" json:"oaidMd5,omitempty"` //oaidMd5 = toLowerCase(md5(oaid 原值))
	Idfa    string `protobuf:"bytes,5,opt,name=idfa,proto3" json:"idfa,omitempty"`       //idfa 原值
	IdfaMd5 string `protobuf:"bytes,6,opt,name=idfaMd5,proto3" json:"idfaMd5,omitempty"` //idfaMd5 = toLowerCase(md5(idfa 原值))
}

func (x *RtaRequest_Device) Reset() {
	*x = RtaRequest_Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuaishou_kuaishou_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RtaRequest_Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RtaRequest_Device) ProtoMessage() {}

func (x *RtaRequest_Device) ProtoReflect() protoreflect.Message {
	mi := &file_kuaishou_kuaishou_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RtaRequest_Device.ProtoReflect.Descriptor instead.
func (*RtaRequest_Device) Descriptor() ([]byte, []int) {
	return file_kuaishou_kuaishou_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RtaRequest_Device) GetImei() string {
	if x != nil {
		return x.Imei
	}
	return ""
}

func (x *RtaRequest_Device) GetImeiMd5() string {
	if x != nil {
		return x.ImeiMd5
	}
	return ""
}

func (x *RtaRequest_Device) GetOaid() string {
	if x != nil {
		return x.Oaid
	}
	return ""
}

func (x *RtaRequest_Device) GetOaidMd5() string {
	if x != nil {
		return x.OaidMd5
	}
	return ""
}

func (x *RtaRequest_Device) GetIdfa() string {
	if x != nil {
		return x.Idfa
	}
	return ""
}

func (x *RtaRequest_Device) GetIdfaMd5() string {
	if x != nil {
		return x.IdfaMd5
	}
	return ""
}

type RtaResponse_PromotionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PromotionType string `protobuf:"bytes,1,opt,name=promotion_type,json=promotionType,proto3" json:"promotion_type,omitempty"` //对应 RtaRequest.promotion_type
	Accept        bool   `protobuf:"varint,2,opt,name=accept,proto3" json:"accept,omitempty"`                                   //true:选择该流量，可以参与竞价
}

func (x *RtaResponse_PromotionResult) Reset() {
	*x = RtaResponse_PromotionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kuaishou_kuaishou_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RtaResponse_PromotionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RtaResponse_PromotionResult) ProtoMessage() {}

func (x *RtaResponse_PromotionResult) ProtoReflect() protoreflect.Message {
	mi := &file_kuaishou_kuaishou_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RtaResponse_PromotionResult.ProtoReflect.Descriptor instead.
func (*RtaResponse_PromotionResult) Descriptor() ([]byte, []int) {
	return file_kuaishou_kuaishou_proto_rawDescGZIP(), []int{1, 0}
}

func (x *RtaResponse_PromotionResult) GetPromotionType() string {
	if x != nil {
		return x.PromotionType
	}
	return ""
}

func (x *RtaResponse_PromotionResult) GetAccept() bool {
	if x != nil {
		return x.Accept
	}
	return false
}

var File_kuaishou_kuaishou_proto protoreflect.FileDescriptor

var file_kuaishou_kuaishou_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6b, 0x75, 0x61, 0x69, 0x73, 0x68, 0x6f, 0x75, 0x2f, 0x6b, 0x75, 0x61, 0x69, 0x73,
	0x68, 0x6f, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6b, 0x75, 0x61, 0x69, 0x73,
	0x68, 0x6f, 0x75, 0x22, 0xed, 0x02, 0x0a, 0x0a, 0x52, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69,
	0x67, 0x6e, 0x12, 0x33, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x75, 0x61, 0x69, 0x73, 0x68, 0x6f, 0x75, 0x2e, 0x52, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x92,
	0x01, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x65,
	0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6d, 0x65, 0x69, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x6d, 0x65, 0x69, 0x4d, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x6d, 0x65, 0x69, 0x4d, 0x64, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x61, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x61, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x61, 0x69, 0x64, 0x4d, 0x64, 0x35, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x61,
	0x69, 0x64, 0x4d, 0x64, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x64, 0x66, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x64, 0x66, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x66,
	0x61, 0x4d, 0x64, 0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x66, 0x61,
	0x4d, 0x64, 0x35, 0x22, 0xd2, 0x01, 0x0a, 0x0b, 0x52, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x50, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x6b, 0x75, 0x61, 0x69, 0x73, 0x68, 0x6f, 0x75, 0x2e, 0x52, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x50, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x42, 0x1c, 0x5a, 0x1a, 0x79, 0x64, 0x73, 0x64,
	0x5f, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x6b, 0x75,
	0x61, 0x69, 0x73, 0x68, 0x6f, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kuaishou_kuaishou_proto_rawDescOnce sync.Once
	file_kuaishou_kuaishou_proto_rawDescData = file_kuaishou_kuaishou_proto_rawDesc
)

func file_kuaishou_kuaishou_proto_rawDescGZIP() []byte {
	file_kuaishou_kuaishou_proto_rawDescOnce.Do(func() {
		file_kuaishou_kuaishou_proto_rawDescData = protoimpl.X.CompressGZIP(file_kuaishou_kuaishou_proto_rawDescData)
	})
	return file_kuaishou_kuaishou_proto_rawDescData
}

var file_kuaishou_kuaishou_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kuaishou_kuaishou_proto_goTypes = []any{
	(*RtaRequest)(nil),                  // 0: kuaishou.RtaRequest
	(*RtaResponse)(nil),                 // 1: kuaishou.RtaResponse
	(*RtaRequest_Device)(nil),           // 2: kuaishou.RtaRequest.Device
	(*RtaResponse_PromotionResult)(nil), // 3: kuaishou.RtaResponse.PromotionResult
}
var file_kuaishou_kuaishou_proto_depIdxs = []int32{
	2, // 0: kuaishou.RtaRequest.device:type_name -> kuaishou.RtaRequest.Device
	3, // 1: kuaishou.RtaResponse.promotion_result:type_name -> kuaishou.RtaResponse.PromotionResult
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_kuaishou_kuaishou_proto_init() }
func file_kuaishou_kuaishou_proto_init() {
	if File_kuaishou_kuaishou_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kuaishou_kuaishou_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RtaRequest); i {
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
		file_kuaishou_kuaishou_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RtaResponse); i {
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
		file_kuaishou_kuaishou_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RtaRequest_Device); i {
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
		file_kuaishou_kuaishou_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RtaResponse_PromotionResult); i {
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
			RawDescriptor: file_kuaishou_kuaishou_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kuaishou_kuaishou_proto_goTypes,
		DependencyIndexes: file_kuaishou_kuaishou_proto_depIdxs,
		MessageInfos:      file_kuaishou_kuaishou_proto_msgTypes,
	}.Build()
	File_kuaishou_kuaishou_proto = out.File
	file_kuaishou_kuaishou_proto_rawDesc = nil
	file_kuaishou_kuaishou_proto_goTypes = nil
	file_kuaishou_kuaishou_proto_depIdxs = nil
}