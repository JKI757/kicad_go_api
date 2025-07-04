//
// This program source code file is part of KiCad, a free EDA CAD application.
//
// Copyright (C) 2024 KiCad Developers, see AUTHORS.txt for contributors.
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: common/commands/project_commands.proto

package commands

import (
	types "github.com/kicad/proto/common/types"
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

type GetNetClasses struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetNetClasses) Reset() {
	*x = GetNetClasses{}
	mi := &file_common_commands_project_commands_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetNetClasses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetClasses) ProtoMessage() {}

func (x *GetNetClasses) ProtoReflect() protoreflect.Message {
	mi := &file_common_commands_project_commands_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetClasses.ProtoReflect.Descriptor instead.
func (*GetNetClasses) Descriptor() ([]byte, []int) {
	return file_common_commands_project_commands_proto_rawDescGZIP(), []int{0}
}

type NetClassesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NetClasses    []*types.NetClass      `protobuf:"bytes,1,rep,name=net_classes,json=netClasses,proto3" json:"net_classes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NetClassesResponse) Reset() {
	*x = NetClassesResponse{}
	mi := &file_common_commands_project_commands_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetClassesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetClassesResponse) ProtoMessage() {}

func (x *NetClassesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_commands_project_commands_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetClassesResponse.ProtoReflect.Descriptor instead.
func (*NetClassesResponse) Descriptor() ([]byte, []int) {
	return file_common_commands_project_commands_proto_rawDescGZIP(), []int{1}
}

func (x *NetClassesResponse) GetNetClasses() []*types.NetClass {
	if x != nil {
		return x.NetClasses
	}
	return nil
}

type SetNetClasses struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	NetClasses []*types.NetClass      `protobuf:"bytes,1,rep,name=net_classes,json=netClasses,proto3" json:"net_classes,omitempty"`
	// Whether to merge or replace the existing netclasses with the contents of this message
	// Note that this only happens at the level of netclass name: for example, if merge_mode is set to
	// MMM_MERGE, the design has netclasses ["Default", "HV"], and this message has netclasses
	// ["Default", "LV"], the resulting set will be ["Default", "HV", "LV"] -- the Default netclass
	// will have its properties replaced with those in this message, the "LV" netclass will be added,
	// and the "HV" netclass will be left alone.  If merge_mode is set to MMM_REPLACE, the "HV" class
	// will be erased.  Note that there must always be a "Default" netclass, so it will not be erased
	// even if merge_mode is MMM_REPLACE and there is no "Default" class specified in this message.
	MergeMode     types.MapMergeMode `protobuf:"varint,3,opt,name=merge_mode,json=mergeMode,proto3,enum=kiapi.common.types.MapMergeMode" json:"merge_mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetNetClasses) Reset() {
	*x = SetNetClasses{}
	mi := &file_common_commands_project_commands_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetNetClasses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNetClasses) ProtoMessage() {}

func (x *SetNetClasses) ProtoReflect() protoreflect.Message {
	mi := &file_common_commands_project_commands_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNetClasses.ProtoReflect.Descriptor instead.
func (*SetNetClasses) Descriptor() ([]byte, []int) {
	return file_common_commands_project_commands_proto_rawDescGZIP(), []int{2}
}

func (x *SetNetClasses) GetNetClasses() []*types.NetClass {
	if x != nil {
		return x.NetClasses
	}
	return nil
}

func (x *SetNetClasses) GetMergeMode() types.MapMergeMode {
	if x != nil {
		return x.MergeMode
	}
	return types.MapMergeMode(0)
}

type ExpandTextVariables struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Document      *types.DocumentSpecifier `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	Text          []string                 `protobuf:"bytes,2,rep,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpandTextVariables) Reset() {
	*x = ExpandTextVariables{}
	mi := &file_common_commands_project_commands_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpandTextVariables) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandTextVariables) ProtoMessage() {}

func (x *ExpandTextVariables) ProtoReflect() protoreflect.Message {
	mi := &file_common_commands_project_commands_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandTextVariables.ProtoReflect.Descriptor instead.
func (*ExpandTextVariables) Descriptor() ([]byte, []int) {
	return file_common_commands_project_commands_proto_rawDescGZIP(), []int{3}
}

func (x *ExpandTextVariables) GetDocument() *types.DocumentSpecifier {
	if x != nil {
		return x.Document
	}
	return nil
}

func (x *ExpandTextVariables) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

type ExpandTextVariablesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          []string               `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExpandTextVariablesResponse) Reset() {
	*x = ExpandTextVariablesResponse{}
	mi := &file_common_commands_project_commands_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExpandTextVariablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandTextVariablesResponse) ProtoMessage() {}

func (x *ExpandTextVariablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_commands_project_commands_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandTextVariablesResponse.ProtoReflect.Descriptor instead.
func (*ExpandTextVariablesResponse) Descriptor() ([]byte, []int) {
	return file_common_commands_project_commands_proto_rawDescGZIP(), []int{4}
}

func (x *ExpandTextVariablesResponse) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

// returns kiapi.common.project.TextVariables
type GetTextVariables struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Document      *types.DocumentSpecifier `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTextVariables) Reset() {
	*x = GetTextVariables{}
	mi := &file_common_commands_project_commands_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTextVariables) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextVariables) ProtoMessage() {}

func (x *GetTextVariables) ProtoReflect() protoreflect.Message {
	mi := &file_common_commands_project_commands_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextVariables.ProtoReflect.Descriptor instead.
func (*GetTextVariables) Descriptor() ([]byte, []int) {
	return file_common_commands_project_commands_proto_rawDescGZIP(), []int{5}
}

func (x *GetTextVariables) GetDocument() *types.DocumentSpecifier {
	if x != nil {
		return x.Document
	}
	return nil
}

type SetTextVariables struct {
	state     protoimpl.MessageState   `protogen:"open.v1"`
	Document  *types.DocumentSpecifier `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	Variables *types.TextVariables     `protobuf:"bytes,2,opt,name=variables,proto3" json:"variables,omitempty"`
	// Whether to merge or replace the existing text variables map with the contents of this message
	MergeMode     types.MapMergeMode `protobuf:"varint,3,opt,name=merge_mode,json=mergeMode,proto3,enum=kiapi.common.types.MapMergeMode" json:"merge_mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetTextVariables) Reset() {
	*x = SetTextVariables{}
	mi := &file_common_commands_project_commands_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetTextVariables) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTextVariables) ProtoMessage() {}

func (x *SetTextVariables) ProtoReflect() protoreflect.Message {
	mi := &file_common_commands_project_commands_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTextVariables.ProtoReflect.Descriptor instead.
func (*SetTextVariables) Descriptor() ([]byte, []int) {
	return file_common_commands_project_commands_proto_rawDescGZIP(), []int{6}
}

func (x *SetTextVariables) GetDocument() *types.DocumentSpecifier {
	if x != nil {
		return x.Document
	}
	return nil
}

func (x *SetTextVariables) GetVariables() *types.TextVariables {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *SetTextVariables) GetMergeMode() types.MapMergeMode {
	if x != nil {
		return x.MergeMode
	}
	return types.MapMergeMode(0)
}

var File_common_commands_project_commands_proto protoreflect.FileDescriptor

const file_common_commands_project_commands_proto_rawDesc = "" +
	"\n" +
	"&common/commands/project_commands.proto\x12\x15kiapi.common.commands\x1a\x1dcommon/types/base_types.proto\x1a#common/types/project_settings.proto\"\x0f\n" +
	"\rGetNetClasses\"U\n" +
	"\x12NetClassesResponse\x12?\n" +
	"\vnet_classes\x18\x01 \x03(\v2\x1e.kiapi.common.project.NetClassR\n" +
	"netClasses\"\x91\x01\n" +
	"\rSetNetClasses\x12?\n" +
	"\vnet_classes\x18\x01 \x03(\v2\x1e.kiapi.common.project.NetClassR\n" +
	"netClasses\x12?\n" +
	"\n" +
	"merge_mode\x18\x03 \x01(\x0e2 .kiapi.common.types.MapMergeModeR\tmergeMode\"l\n" +
	"\x13ExpandTextVariables\x12A\n" +
	"\bdocument\x18\x01 \x01(\v2%.kiapi.common.types.DocumentSpecifierR\bdocument\x12\x12\n" +
	"\x04text\x18\x02 \x03(\tR\x04text\"1\n" +
	"\x1bExpandTextVariablesResponse\x12\x12\n" +
	"\x04text\x18\x01 \x03(\tR\x04text\"U\n" +
	"\x10GetTextVariables\x12A\n" +
	"\bdocument\x18\x01 \x01(\v2%.kiapi.common.types.DocumentSpecifierR\bdocument\"\xd9\x01\n" +
	"\x10SetTextVariables\x12A\n" +
	"\bdocument\x18\x01 \x01(\v2%.kiapi.common.types.DocumentSpecifierR\bdocument\x12A\n" +
	"\tvariables\x18\x02 \x01(\v2#.kiapi.common.project.TextVariablesR\tvariables\x12?\n" +
	"\n" +
	"merge_mode\x18\x03 \x01(\x0e2 .kiapi.common.types.MapMergeModeR\tmergeModeB(Z&github.com/kicad/proto/common/commandsb\x06proto3"

var (
	file_common_commands_project_commands_proto_rawDescOnce sync.Once
	file_common_commands_project_commands_proto_rawDescData []byte
)

func file_common_commands_project_commands_proto_rawDescGZIP() []byte {
	file_common_commands_project_commands_proto_rawDescOnce.Do(func() {
		file_common_commands_project_commands_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_commands_project_commands_proto_rawDesc), len(file_common_commands_project_commands_proto_rawDesc)))
	})
	return file_common_commands_project_commands_proto_rawDescData
}

var file_common_commands_project_commands_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_common_commands_project_commands_proto_goTypes = []any{
	(*GetNetClasses)(nil),               // 0: kiapi.common.commands.GetNetClasses
	(*NetClassesResponse)(nil),          // 1: kiapi.common.commands.NetClassesResponse
	(*SetNetClasses)(nil),               // 2: kiapi.common.commands.SetNetClasses
	(*ExpandTextVariables)(nil),         // 3: kiapi.common.commands.ExpandTextVariables
	(*ExpandTextVariablesResponse)(nil), // 4: kiapi.common.commands.ExpandTextVariablesResponse
	(*GetTextVariables)(nil),            // 5: kiapi.common.commands.GetTextVariables
	(*SetTextVariables)(nil),            // 6: kiapi.common.commands.SetTextVariables
	(*types.NetClass)(nil),              // 7: kiapi.common.project.NetClass
	(types.MapMergeMode)(0),             // 8: kiapi.common.types.MapMergeMode
	(*types.DocumentSpecifier)(nil),     // 9: kiapi.common.types.DocumentSpecifier
	(*types.TextVariables)(nil),         // 10: kiapi.common.project.TextVariables
}
var file_common_commands_project_commands_proto_depIdxs = []int32{
	7,  // 0: kiapi.common.commands.NetClassesResponse.net_classes:type_name -> kiapi.common.project.NetClass
	7,  // 1: kiapi.common.commands.SetNetClasses.net_classes:type_name -> kiapi.common.project.NetClass
	8,  // 2: kiapi.common.commands.SetNetClasses.merge_mode:type_name -> kiapi.common.types.MapMergeMode
	9,  // 3: kiapi.common.commands.ExpandTextVariables.document:type_name -> kiapi.common.types.DocumentSpecifier
	9,  // 4: kiapi.common.commands.GetTextVariables.document:type_name -> kiapi.common.types.DocumentSpecifier
	9,  // 5: kiapi.common.commands.SetTextVariables.document:type_name -> kiapi.common.types.DocumentSpecifier
	10, // 6: kiapi.common.commands.SetTextVariables.variables:type_name -> kiapi.common.project.TextVariables
	8,  // 7: kiapi.common.commands.SetTextVariables.merge_mode:type_name -> kiapi.common.types.MapMergeMode
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_common_commands_project_commands_proto_init() }
func file_common_commands_project_commands_proto_init() {
	if File_common_commands_project_commands_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_commands_project_commands_proto_rawDesc), len(file_common_commands_project_commands_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_commands_project_commands_proto_goTypes,
		DependencyIndexes: file_common_commands_project_commands_proto_depIdxs,
		MessageInfos:      file_common_commands_project_commands_proto_msgTypes,
	}.Build()
	File_common_commands_project_commands_proto = out.File
	file_common_commands_project_commands_proto_goTypes = nil
	file_common_commands_project_commands_proto_depIdxs = nil
}
