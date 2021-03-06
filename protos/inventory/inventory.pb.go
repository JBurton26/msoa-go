// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: inventory.proto

package inventory

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

type LevelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=Location,proto3" json:"Location,omitempty"`
}

func (x *LevelRequest) Reset() {
	*x = LevelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelRequest) ProtoMessage() {}

func (x *LevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelRequest.ProtoReflect.Descriptor instead.
func (*LevelRequest) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *LevelRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *LevelRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type ShortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=Location,proto3" json:"Location,omitempty"`
}

func (x *ShortRequest) Reset() {
	*x = ShortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortRequest) ProtoMessage() {}

func (x *ShortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortRequest.ProtoReflect.Descriptor instead.
func (*ShortRequest) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *ShortRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type AmendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Location string  `protobuf:"bytes,2,opt,name=Location,proto3" json:"Location,omitempty"`
	Amount   int64   `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Name     string  `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Price    float64 `protobuf:"fixed64,5,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *AmendRequest) Reset() {
	*x = AmendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmendRequest) ProtoMessage() {}

func (x *AmendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmendRequest.ProtoReflect.Descriptor instead.
func (*AmendRequest) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *AmendRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *AmendRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *AmendRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *AmendRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AmendRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type AmendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *AmendResponse) Reset() {
	*x = AmendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmendResponse) ProtoMessage() {}

func (x *AmendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmendResponse.ProtoReflect.Descriptor instead.
func (*AmendResponse) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *AmendResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type StockItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Location   string `protobuf:"bytes,2,opt,name=Location,proto3" json:"Location,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	StockCount int64  `protobuf:"varint,4,opt,name=StockCount,proto3" json:"StockCount,omitempty"`
}

func (x *StockItem) Reset() {
	*x = StockItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockItem) ProtoMessage() {}

func (x *StockItem) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockItem.ProtoReflect.Descriptor instead.
func (*StockItem) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *StockItem) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *StockItem) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *StockItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StockItem) GetStockCount() int64 {
	if x != nil {
		return x.StockCount
	}
	return 0
}

type ShortList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SList []*StockItem `protobuf:"bytes,1,rep,name=SList,proto3" json:"SList,omitempty"`
}

func (x *ShortList) Reset() {
	*x = ShortList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortList) ProtoMessage() {}

func (x *ShortList) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortList.ProtoReflect.Descriptor instead.
func (*ShortList) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{5}
}

func (x *ShortList) GetSList() []*StockItem {
	if x != nil {
		return x.SList
	}
	return nil
}

var File_inventory_proto protoreflect.FileDescriptor

var file_inventory_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3a, 0x0a, 0x0c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a,
	0x0c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x0c, 0x41, 0x6d, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x41, 0x6d, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x2d, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x05, 0x53, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x53, 0x4c, 0x69, 0x73, 0x74,
	0x32, 0xb0, 0x01, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x25,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0d, 0x2e, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2c, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0d, 0x2e, 0x41, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x41, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x12, 0x0d, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0a, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x0d, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4a, 0x42, 0x75, 0x72, 0x74, 0x6f, 0x6e, 0x32, 0x36, 0x2f, 0x6d, 0x73, 0x6f, 0x61,
	0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_proto_rawDescOnce sync.Once
	file_inventory_proto_rawDescData = file_inventory_proto_rawDesc
)

func file_inventory_proto_rawDescGZIP() []byte {
	file_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_proto_rawDescData)
	})
	return file_inventory_proto_rawDescData
}

var file_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_inventory_proto_goTypes = []interface{}{
	(*LevelRequest)(nil),  // 0: LevelRequest
	(*ShortRequest)(nil),  // 1: ShortRequest
	(*AmendRequest)(nil),  // 2: AmendRequest
	(*AmendResponse)(nil), // 3: AmendResponse
	(*StockItem)(nil),     // 4: StockItem
	(*ShortList)(nil),     // 5: ShortList
}
var file_inventory_proto_depIdxs = []int32{
	4, // 0: ShortList.SList:type_name -> StockItem
	0, // 1: Inventory.GetStock:input_type -> LevelRequest
	2, // 2: Inventory.ChangeStock:input_type -> AmendRequest
	1, // 3: Inventory.CheckShort:input_type -> ShortRequest
	1, // 4: Inventory.GetStore:input_type -> ShortRequest
	4, // 5: Inventory.GetStock:output_type -> StockItem
	3, // 6: Inventory.ChangeStock:output_type -> AmendResponse
	5, // 7: Inventory.CheckShort:output_type -> ShortList
	5, // 8: Inventory.GetStore:output_type -> ShortList
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_inventory_proto_init() }
func file_inventory_proto_init() {
	if File_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelRequest); i {
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
		file_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortRequest); i {
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
		file_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmendRequest); i {
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
		file_inventory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmendResponse); i {
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
		file_inventory_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockItem); i {
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
		file_inventory_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortList); i {
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
			RawDescriptor: file_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_proto_msgTypes,
	}.Build()
	File_inventory_proto = out.File
	file_inventory_proto_rawDesc = nil
	file_inventory_proto_goTypes = nil
	file_inventory_proto_depIdxs = nil
}
