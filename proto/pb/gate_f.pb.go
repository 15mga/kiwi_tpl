// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.1
// source: fail/gate_f.proto

package pb

import (
	_ "github.com/15mga/kiwi_tool"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_fail_gate_f_proto protoreflect.FileDescriptor

var file_fail_gate_f_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x61, 0x69, 0x6c, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x31, 0x35, 0x6d, 0x67, 0x61, 0x2f, 0x6b, 0x69, 0x77, 0x69, 0x5f, 0x74,
	0x6f, 0x6f, 0x6c, 0x2f, 0x6b, 0x69, 0x77, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x64,
	0xca, 0xdd, 0x49, 0x56, 0x12, 0x04, 0x67, 0x61, 0x74, 0x65, 0x2a, 0x23, 0x08, 0xe8, 0x07, 0x12,
	0x11, 0x67, 0x61, 0x74, 0x65, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x20,
	0x69, 0x64, 0x1a, 0x0b, 0x69, 0x64, 0xe4, 0xb8, 0x8d, 0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0x2a,
	0x29, 0x08, 0xe9, 0x07, 0x12, 0x13, 0x67, 0x61, 0x74, 0x65, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x20, 0x61, 0x64, 0x64, 0x72, 0x1a, 0x0f, 0xe5, 0x9c, 0xb0, 0xe5, 0x9d,
	0x80, 0xe4, 0xb8, 0x8d, 0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0xaa,
	0x02, 0x02, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_fail_gate_f_proto_goTypes = []interface{}{}
var file_fail_gate_f_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fail_gate_f_proto_init() }
func file_fail_gate_f_proto_init() {
	if File_fail_gate_f_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fail_gate_f_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fail_gate_f_proto_goTypes,
		DependencyIndexes: file_fail_gate_f_proto_depIdxs,
	}.Build()
	File_fail_gate_f_proto = out.File
	file_fail_gate_f_proto_rawDesc = nil
	file_fail_gate_f_proto_goTypes = nil
	file_fail_gate_f_proto_depIdxs = nil
}
