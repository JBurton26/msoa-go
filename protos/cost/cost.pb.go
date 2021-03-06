// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: cost.proto

package cost

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

type Basket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Basket_Item `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *Basket) Reset() {
	*x = Basket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Basket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Basket) ProtoMessage() {}

func (x *Basket) ProtoReflect() protoreflect.Message {
	mi := &file_cost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Basket.ProtoReflect.Descriptor instead.
func (*Basket) Descriptor() ([]byte, []int) {
	return file_cost_proto_rawDescGZIP(), []int{0}
}

func (x *Basket) GetItems() []*Basket_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type CostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CostRequest) Reset() {
	*x = CostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cost_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostRequest) ProtoMessage() {}

func (x *CostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cost_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostRequest.ProtoReflect.Descriptor instead.
func (*CostRequest) Descriptor() ([]byte, []int) {
	return file_cost_proto_rawDescGZIP(), []int{1}
}

func (x *CostRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Price float64 `protobuf:"fixed64,2,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cost_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cost_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_cost_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price float64 `protobuf:"fixed64,1,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *CostResponse) Reset() {
	*x = CostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cost_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostResponse) ProtoMessage() {}

func (x *CostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cost_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostResponse.ProtoReflect.Descriptor instead.
func (*CostResponse) Descriptor() ([]byte, []int) {
	return file_cost_proto_rawDescGZIP(), []int{3}
}

func (x *CostResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cost_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cost_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_cost_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Basket_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *Basket_Item) Reset() {
	*x = Basket_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cost_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Basket_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Basket_Item) ProtoMessage() {}

func (x *Basket_Item) ProtoReflect() protoreflect.Message {
	mi := &file_cost_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Basket_Item.ProtoReflect.Descriptor instead.
func (*Basket_Item) Descriptor() ([]byte, []int) {
	return file_cost_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Basket_Item) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Basket_Item) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_cost_proto protoreflect.FileDescriptor

var file_cost_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x06,
	0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x2c, 0x0a, 0x04, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1d, 0x0a, 0x0b, 0x43, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x35, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x24,
	0x0a, 0x0c, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0xbc, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x55, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x6e, 0x69, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x55,
	0x6e, 0x69, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x12, 0x07, 0x2e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x1a, 0x0d, 0x2e, 0x43, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x42,
	0x75, 0x72, 0x74, 0x6f, 0x6e, 0x32, 0x36, 0x2f, 0x6d, 0x73, 0x6f, 0x61, 0x2d, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cost_proto_rawDescOnce sync.Once
	file_cost_proto_rawDescData = file_cost_proto_rawDesc
)

func file_cost_proto_rawDescGZIP() []byte {
	file_cost_proto_rawDescOnce.Do(func() {
		file_cost_proto_rawDescData = protoimpl.X.CompressGZIP(file_cost_proto_rawDescData)
	})
	return file_cost_proto_rawDescData
}

var file_cost_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cost_proto_goTypes = []interface{}{
	(*Basket)(nil),         // 0: Basket
	(*CostRequest)(nil),    // 1: CostRequest
	(*UpdateRequest)(nil),  // 2: UpdateRequest
	(*CostResponse)(nil),   // 3: CostResponse
	(*UpdateResponse)(nil), // 4: UpdateResponse
	(*Basket_Item)(nil),    // 5: Basket.Item
}
var file_cost_proto_depIdxs = []int32{
	5, // 0: Basket.Items:type_name -> Basket.Item
	1, // 1: Cost.GetUnitCost:input_type -> CostRequest
	2, // 2: Cost.UpdateUnitCost:input_type -> UpdateRequest
	2, // 3: Cost.AddUnitCost:input_type -> UpdateRequest
	0, // 4: Cost.TotalBasket:input_type -> Basket
	3, // 5: Cost.GetUnitCost:output_type -> CostResponse
	4, // 6: Cost.UpdateUnitCost:output_type -> UpdateResponse
	4, // 7: Cost.AddUnitCost:output_type -> UpdateResponse
	3, // 8: Cost.TotalBasket:output_type -> CostResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cost_proto_init() }
func file_cost_proto_init() {
	if File_cost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Basket); i {
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
		file_cost_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostRequest); i {
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
		file_cost_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_cost_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostResponse); i {
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
		file_cost_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_cost_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Basket_Item); i {
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
			RawDescriptor: file_cost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cost_proto_goTypes,
		DependencyIndexes: file_cost_proto_depIdxs,
		MessageInfos:      file_cost_proto_msgTypes,
	}.Build()
	File_cost_proto = out.File
	file_cost_proto_rawDesc = nil
	file_cost_proto_goTypes = nil
	file_cost_proto_depIdxs = nil
}
