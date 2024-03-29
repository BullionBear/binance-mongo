// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: protocols/wsdepth.proto

package wsdepth

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

// The depth event message
type WsDepthEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event         string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Time          int64  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Symbol        string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	LastUpdateID  int64  `protobuf:"varint,4,opt,name=lastUpdateID,proto3" json:"lastUpdateID,omitempty"`
	FirstUpdateID int64  `protobuf:"varint,5,opt,name=firstUpdateID,proto3" json:"firstUpdateID,omitempty"`
	Bids          []*Bid `protobuf:"bytes,6,rep,name=bids,proto3" json:"bids,omitempty"`
	Asks          []*Ask `protobuf:"bytes,7,rep,name=asks,proto3" json:"asks,omitempty"`
}

func (x *WsDepthEvent) Reset() {
	*x = WsDepthEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_wsdepth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WsDepthEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WsDepthEvent) ProtoMessage() {}

func (x *WsDepthEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_wsdepth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WsDepthEvent.ProtoReflect.Descriptor instead.
func (*WsDepthEvent) Descriptor() ([]byte, []int) {
	return file_protocols_wsdepth_proto_rawDescGZIP(), []int{0}
}

func (x *WsDepthEvent) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *WsDepthEvent) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *WsDepthEvent) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *WsDepthEvent) GetLastUpdateID() int64 {
	if x != nil {
		return x.LastUpdateID
	}
	return 0
}

func (x *WsDepthEvent) GetFirstUpdateID() int64 {
	if x != nil {
		return x.FirstUpdateID
	}
	return 0
}

func (x *WsDepthEvent) GetBids() []*Bid {
	if x != nil {
		return x.Bids
	}
	return nil
}

func (x *WsDepthEvent) GetAsks() []*Ask {
	if x != nil {
		return x.Asks
	}
	return nil
}

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price    string `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	Quantity string `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_wsdepth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_wsdepth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_protocols_wsdepth_proto_rawDescGZIP(), []int{1}
}

func (x *Bid) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Bid) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

type Ask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price    string `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	Quantity string `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Ask) Reset() {
	*x = Ask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_wsdepth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ask) ProtoMessage() {}

func (x *Ask) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_wsdepth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ask.ProtoReflect.Descriptor instead.
func (*Ask) Descriptor() ([]byte, []int) {
	return file_protocols_wsdepth_proto_rawDescGZIP(), []int{2}
}

func (x *Ask) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Ask) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

type StreamDepthEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StreamDepthEventResponse) Reset() {
	*x = StreamDepthEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocols_wsdepth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamDepthEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamDepthEventResponse) ProtoMessage() {}

func (x *StreamDepthEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocols_wsdepth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamDepthEventResponse.ProtoReflect.Descriptor instead.
func (*StreamDepthEventResponse) Descriptor() ([]byte, []int) {
	return file_protocols_wsdepth_proto_rawDescGZIP(), []int{3}
}

func (x *StreamDepthEventResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protocols_wsdepth_proto protoreflect.FileDescriptor

var file_protocols_wsdepth_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x77, 0x73, 0x64, 0x65,
	0x70, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x65, 0x70, 0x74, 0x68,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xe4, 0x01, 0x0a, 0x0c, 0x57, 0x73, 0x44, 0x65, 0x70, 0x74,
	0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x44, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x64, 0x65, 0x70, 0x74, 0x68, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x69,
	0x64, 0x52, 0x04, 0x62, 0x69, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x04, 0x61, 0x73, 0x6b, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x65, 0x70, 0x74, 0x68, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x41, 0x73, 0x6b, 0x52, 0x04, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x37, 0x0a, 0x03,
	0x42, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x37, 0x0a, 0x03, 0x41, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x34,
	0x0a, 0x18, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x70, 0x74, 0x68, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x69, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x74, 0x68, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x44, 0x65, 0x70, 0x74, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e,
	0x64, 0x65, 0x70, 0x74, 0x68, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x73, 0x44, 0x65, 0x70,
	0x74, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x65, 0x70, 0x74, 0x68, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x65, 0x70, 0x74, 0x68,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42,
	0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x73, 0x64, 0x65, 0x70, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocols_wsdepth_proto_rawDescOnce sync.Once
	file_protocols_wsdepth_proto_rawDescData = file_protocols_wsdepth_proto_rawDesc
)

func file_protocols_wsdepth_proto_rawDescGZIP() []byte {
	file_protocols_wsdepth_proto_rawDescOnce.Do(func() {
		file_protocols_wsdepth_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocols_wsdepth_proto_rawDescData)
	})
	return file_protocols_wsdepth_proto_rawDescData
}

var file_protocols_wsdepth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protocols_wsdepth_proto_goTypes = []interface{}{
	(*WsDepthEvent)(nil),             // 0: depthevent.WsDepthEvent
	(*Bid)(nil),                      // 1: depthevent.Bid
	(*Ask)(nil),                      // 2: depthevent.Ask
	(*StreamDepthEventResponse)(nil), // 3: depthevent.StreamDepthEventResponse
}
var file_protocols_wsdepth_proto_depIdxs = []int32{
	1, // 0: depthevent.WsDepthEvent.bids:type_name -> depthevent.Bid
	2, // 1: depthevent.WsDepthEvent.asks:type_name -> depthevent.Ask
	0, // 2: depthevent.DepthEventService.StreamDepthEvent:input_type -> depthevent.WsDepthEvent
	3, // 3: depthevent.DepthEventService.StreamDepthEvent:output_type -> depthevent.StreamDepthEventResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protocols_wsdepth_proto_init() }
func file_protocols_wsdepth_proto_init() {
	if File_protocols_wsdepth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocols_wsdepth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WsDepthEvent); i {
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
		file_protocols_wsdepth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bid); i {
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
		file_protocols_wsdepth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ask); i {
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
		file_protocols_wsdepth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamDepthEventResponse); i {
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
			RawDescriptor: file_protocols_wsdepth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocols_wsdepth_proto_goTypes,
		DependencyIndexes: file_protocols_wsdepth_proto_depIdxs,
		MessageInfos:      file_protocols_wsdepth_proto_msgTypes,
	}.Build()
	File_protocols_wsdepth_proto = out.File
	file_protocols_wsdepth_proto_rawDesc = nil
	file_protocols_wsdepth_proto_goTypes = nil
	file_protocols_wsdepth_proto_depIdxs = nil
}
