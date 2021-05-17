// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: cep.proto

package pb

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

type CepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cep string `protobuf:"bytes,1,opt,name=cep,proto3" json:"cep,omitempty"`
}

func (x *CepRequest) Reset() {
	*x = CepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cep_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CepRequest) ProtoMessage() {}

func (x *CepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cep_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CepRequest.ProtoReflect.Descriptor instead.
func (*CepRequest) Descriptor() ([]byte, []int) {
	return file_cep_proto_rawDescGZIP(), []int{0}
}

func (x *CepRequest) GetCep() string {
	if x != nil {
		return x.Cep
	}
	return ""
}

type CepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cep          string            `protobuf:"bytes,1,opt,name=cep,proto3" json:"cep,omitempty"`
	Street       string            `protobuf:"bytes,2,opt,name=street,proto3" json:"street,omitempty"`
	Complement   string            `protobuf:"bytes,3,opt,name=complement,proto3" json:"complement,omitempty"`
	Neighborhood string            `protobuf:"bytes,4,opt,name=neighborhood,proto3" json:"neighborhood,omitempty"`
	City         string            `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	Uf           string            `protobuf:"bytes,6,opt,name=uf,proto3" json:"uf,omitempty"`
	Location     *LocationResponse `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CepResponse) Reset() {
	*x = CepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cep_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CepResponse) ProtoMessage() {}

func (x *CepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cep_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CepResponse.ProtoReflect.Descriptor instead.
func (*CepResponse) Descriptor() ([]byte, []int) {
	return file_cep_proto_rawDescGZIP(), []int{1}
}

func (x *CepResponse) GetCep() string {
	if x != nil {
		return x.Cep
	}
	return ""
}

func (x *CepResponse) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *CepResponse) GetComplement() string {
	if x != nil {
		return x.Complement
	}
	return ""
}

func (x *CepResponse) GetNeighborhood() string {
	if x != nil {
		return x.Neighborhood
	}
	return ""
}

func (x *CepResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CepResponse) GetUf() string {
	if x != nil {
		return x.Uf
	}
	return ""
}

func (x *CepResponse) GetLocation() *LocationResponse {
	if x != nil {
		return x.Location
	}
	return nil
}

type LocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float32 `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng float32 `protobuf:"fixed32,2,opt,name=lng,proto3" json:"lng,omitempty"`
}

func (x *LocationResponse) Reset() {
	*x = LocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cep_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationResponse) ProtoMessage() {}

func (x *LocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cep_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationResponse.ProtoReflect.Descriptor instead.
func (*LocationResponse) Descriptor() ([]byte, []int) {
	return file_cep_proto_rawDescGZIP(), []int{2}
}

func (x *LocationResponse) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *LocationResponse) GetLng() float32 {
	if x != nil {
		return x.Lng
	}
	return 0
}

var File_cep_proto protoreflect.FileDescriptor

var file_cep_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x61, 0x69, 0x6c, 0x74, 0x6f, 0x6e, 0x6a,
	0x75, 0x6e, 0x69, 0x6f, 0x72, 0x39, 0x34, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x22, 0x1e, 0x0a, 0x0a, 0x43, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x63, 0x65, 0x70, 0x22, 0xf6, 0x01, 0x0a, 0x0b, 0x43, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x63, 0x65, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f, 0x6f, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x68, 0x6f,
	0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x75, 0x66, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x75, 0x66, 0x12, 0x55, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x61, 0x69, 0x6c, 0x74, 0x6f, 0x6e, 0x6a, 0x75, 0x6e,
	0x69, 0x6f, 0x72, 0x39, 0x34, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x36, 0x0a,
	0x10, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03,
	0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x6c, 0x6e, 0x67, 0x32, 0x83, 0x01, 0x0a, 0x0a, 0x43, 0x65, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x75, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x43, 0x65, 0x70, 0x12, 0x33,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x61, 0x69, 0x6c,
	0x74, 0x6f, 0x6e, 0x6a, 0x75, 0x6e, 0x69, 0x6f, 0x72, 0x39, 0x34, 0x2e, 0x67, 0x6f, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x6a, 0x61, 0x69, 0x6c, 0x74, 0x6f, 0x6e, 0x6a, 0x75, 0x6e, 0x69, 0x6f, 0x72, 0x39, 0x34,
	0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x65,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cep_proto_rawDescOnce sync.Once
	file_cep_proto_rawDescData = file_cep_proto_rawDesc
)

func file_cep_proto_rawDescGZIP() []byte {
	file_cep_proto_rawDescOnce.Do(func() {
		file_cep_proto_rawDescData = protoimpl.X.CompressGZIP(file_cep_proto_rawDescData)
	})
	return file_cep_proto_rawDescData
}

var file_cep_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cep_proto_goTypes = []interface{}{
	(*CepRequest)(nil),       // 0: github.com.jailtonjunior94.go.challenge.CepRequest
	(*CepResponse)(nil),      // 1: github.com.jailtonjunior94.go.challenge.CepResponse
	(*LocationResponse)(nil), // 2: github.com.jailtonjunior94.go.challenge.LocationResponse
}
var file_cep_proto_depIdxs = []int32{
	2, // 0: github.com.jailtonjunior94.go.challenge.CepResponse.location:type_name -> github.com.jailtonjunior94.go.challenge.LocationResponse
	0, // 1: github.com.jailtonjunior94.go.challenge.CepService.GetCep:input_type -> github.com.jailtonjunior94.go.challenge.CepRequest
	1, // 2: github.com.jailtonjunior94.go.challenge.CepService.GetCep:output_type -> github.com.jailtonjunior94.go.challenge.CepResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cep_proto_init() }
func file_cep_proto_init() {
	if File_cep_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cep_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CepRequest); i {
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
		file_cep_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CepResponse); i {
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
		file_cep_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationResponse); i {
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
			RawDescriptor: file_cep_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cep_proto_goTypes,
		DependencyIndexes: file_cep_proto_depIdxs,
		MessageInfos:      file_cep_proto_msgTypes,
	}.Build()
	File_cep_proto = out.File
	file_cep_proto_rawDesc = nil
	file_cep_proto_goTypes = nil
	file_cep_proto_depIdxs = nil
}
