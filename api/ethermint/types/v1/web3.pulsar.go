// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package typesv1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_ExtensionOptionsWeb3Tx                     protoreflect.MessageDescriptor
	fd_ExtensionOptionsWeb3Tx_typed_data_chain_id protoreflect.FieldDescriptor
	fd_ExtensionOptionsWeb3Tx_fee_payer           protoreflect.FieldDescriptor
	fd_ExtensionOptionsWeb3Tx_fee_payer_sig       protoreflect.FieldDescriptor
)

func init() {
	file_ethermint_types_v1_web3_proto_init()
	md_ExtensionOptionsWeb3Tx = File_ethermint_types_v1_web3_proto.Messages().ByName("ExtensionOptionsWeb3Tx")
	fd_ExtensionOptionsWeb3Tx_typed_data_chain_id = md_ExtensionOptionsWeb3Tx.Fields().ByName("typed_data_chain_id")
	fd_ExtensionOptionsWeb3Tx_fee_payer = md_ExtensionOptionsWeb3Tx.Fields().ByName("fee_payer")
	fd_ExtensionOptionsWeb3Tx_fee_payer_sig = md_ExtensionOptionsWeb3Tx.Fields().ByName("fee_payer_sig")
}

var _ protoreflect.Message = (*fastReflection_ExtensionOptionsWeb3Tx)(nil)

type fastReflection_ExtensionOptionsWeb3Tx ExtensionOptionsWeb3Tx

func (x *ExtensionOptionsWeb3Tx) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ExtensionOptionsWeb3Tx)(x)
}

func (x *ExtensionOptionsWeb3Tx) slowProtoReflect() protoreflect.Message {
	mi := &file_ethermint_types_v1_web3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ExtensionOptionsWeb3Tx_messageType fastReflection_ExtensionOptionsWeb3Tx_messageType
var _ protoreflect.MessageType = fastReflection_ExtensionOptionsWeb3Tx_messageType{}

type fastReflection_ExtensionOptionsWeb3Tx_messageType struct{}

func (x fastReflection_ExtensionOptionsWeb3Tx_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ExtensionOptionsWeb3Tx)(nil)
}
func (x fastReflection_ExtensionOptionsWeb3Tx_messageType) New() protoreflect.Message {
	return new(fastReflection_ExtensionOptionsWeb3Tx)
}
func (x fastReflection_ExtensionOptionsWeb3Tx_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ExtensionOptionsWeb3Tx
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ExtensionOptionsWeb3Tx) Descriptor() protoreflect.MessageDescriptor {
	return md_ExtensionOptionsWeb3Tx
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ExtensionOptionsWeb3Tx) Type() protoreflect.MessageType {
	return _fastReflection_ExtensionOptionsWeb3Tx_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ExtensionOptionsWeb3Tx) New() protoreflect.Message {
	return new(fastReflection_ExtensionOptionsWeb3Tx)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ExtensionOptionsWeb3Tx) Interface() protoreflect.ProtoMessage {
	return (*ExtensionOptionsWeb3Tx)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ExtensionOptionsWeb3Tx) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.TypedDataChainId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.TypedDataChainId)
		if !f(fd_ExtensionOptionsWeb3Tx_typed_data_chain_id, value) {
			return
		}
	}
	if x.FeePayer != "" {
		value := protoreflect.ValueOfString(x.FeePayer)
		if !f(fd_ExtensionOptionsWeb3Tx_fee_payer, value) {
			return
		}
	}
	if len(x.FeePayerSig) != 0 {
		value := protoreflect.ValueOfBytes(x.FeePayerSig)
		if !f(fd_ExtensionOptionsWeb3Tx_fee_payer_sig, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ExtensionOptionsWeb3Tx) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.typed_data_chain_id":
		return x.TypedDataChainId != uint64(0)
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.fee_payer":
		return x.FeePayer != ""
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.fee_payer_sig":
		return len(x.FeePayerSig) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ethermint.types.v1.ExtensionOptionsWeb3Tx"))
		}
		panic(fmt.Errorf("message ethermint.types.v1.ExtensionOptionsWeb3Tx does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ExtensionOptionsWeb3Tx) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.typed_data_chain_id":
		x.TypedDataChainId = uint64(0)
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.fee_payer":
		x.FeePayer = ""
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.fee_payer_sig":
		x.FeePayerSig = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ethermint.types.v1.ExtensionOptionsWeb3Tx"))
		}
		panic(fmt.Errorf("message ethermint.types.v1.ExtensionOptionsWeb3Tx does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ExtensionOptionsWeb3Tx) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.typed_data_chain_id":
		value := x.TypedDataChainId
		return protoreflect.ValueOfUint64(value)
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.fee_payer":
		value := x.FeePayer
		return protoreflect.ValueOfString(value)
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.fee_payer_sig":
		value := x.FeePayerSig
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ethermint.types.v1.ExtensionOptionsWeb3Tx"))
		}
		panic(fmt.Errorf("message ethermint.types.v1.ExtensionOptionsWeb3Tx does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ExtensionOptionsWeb3Tx) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.typed_data_chain_id":
		x.TypedDataChainId = value.Uint()
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.fee_payer":
		x.FeePayer = value.Interface().(string)
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.fee_payer_sig":
		x.FeePayerSig = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ethermint.types.v1.ExtensionOptionsWeb3Tx"))
		}
		panic(fmt.Errorf("message ethermint.types.v1.ExtensionOptionsWeb3Tx does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ExtensionOptionsWeb3Tx) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.typed_data_chain_id":
		panic(fmt.Errorf("field typed_data_chain_id of message ethermint.types.v1.ExtensionOptionsWeb3Tx is not mutable"))
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.fee_payer":
		panic(fmt.Errorf("field fee_payer of message ethermint.types.v1.ExtensionOptionsWeb3Tx is not mutable"))
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.fee_payer_sig":
		panic(fmt.Errorf("field fee_payer_sig of message ethermint.types.v1.ExtensionOptionsWeb3Tx is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ethermint.types.v1.ExtensionOptionsWeb3Tx"))
		}
		panic(fmt.Errorf("message ethermint.types.v1.ExtensionOptionsWeb3Tx does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ExtensionOptionsWeb3Tx) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.typed_data_chain_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.fee_payer":
		return protoreflect.ValueOfString("")
	case "ethermint.types.v1.ExtensionOptionsWeb3Tx.fee_payer_sig":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: ethermint.types.v1.ExtensionOptionsWeb3Tx"))
		}
		panic(fmt.Errorf("message ethermint.types.v1.ExtensionOptionsWeb3Tx does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ExtensionOptionsWeb3Tx) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in ethermint.types.v1.ExtensionOptionsWeb3Tx", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ExtensionOptionsWeb3Tx) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ExtensionOptionsWeb3Tx) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ExtensionOptionsWeb3Tx) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ExtensionOptionsWeb3Tx) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ExtensionOptionsWeb3Tx)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.TypedDataChainId != 0 {
			n += 1 + runtime.Sov(uint64(x.TypedDataChainId))
		}
		l = len(x.FeePayer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.FeePayerSig)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ExtensionOptionsWeb3Tx)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.FeePayerSig) > 0 {
			i -= len(x.FeePayerSig)
			copy(dAtA[i:], x.FeePayerSig)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.FeePayerSig)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.FeePayer) > 0 {
			i -= len(x.FeePayer)
			copy(dAtA[i:], x.FeePayer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.FeePayer)))
			i--
			dAtA[i] = 0x12
		}
		if x.TypedDataChainId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.TypedDataChainId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ExtensionOptionsWeb3Tx)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ExtensionOptionsWeb3Tx: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ExtensionOptionsWeb3Tx: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TypedDataChainId", wireType)
				}
				x.TypedDataChainId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.TypedDataChainId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FeePayer", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.FeePayer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FeePayerSig", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.FeePayerSig = append(x.FeePayerSig[:0], dAtA[iNdEx:postIndex]...)
				if x.FeePayerSig == nil {
					x.FeePayerSig = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Copyright Tharsis Labs Ltd.(Aizel)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AizelNetwork/evmos/blob/main/LICENSE)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: ethermint/types/v1/web3.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ExtensionOptionsWeb3Tx is an extension option that specifies the typed chain id,
// the fee payer as well as its signature data.
type ExtensionOptionsWeb3Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// typed_data_chain_id is used only in EIP712 Domain and should match
	// Ethereum network ID in a Web3 provider (e.g. Metamask).
	TypedDataChainId uint64 `protobuf:"varint,1,opt,name=typed_data_chain_id,json=typedDataChainId,proto3" json:"typed_data_chain_id,omitempty"`
	// fee_payer is an account address for the fee payer. It will be validated
	// during EIP712 signature checking.
	FeePayer string `protobuf:"bytes,2,opt,name=fee_payer,json=feePayer,proto3" json:"fee_payer,omitempty"`
	// fee_payer_sig is a signature data from the fee paying account,
	// allows to perform fee delegation when using EIP712 Domain.
	FeePayerSig []byte `protobuf:"bytes,3,opt,name=fee_payer_sig,json=feePayerSig,proto3" json:"fee_payer_sig,omitempty"`
}

func (x *ExtensionOptionsWeb3Tx) Reset() {
	*x = ExtensionOptionsWeb3Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethermint_types_v1_web3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionOptionsWeb3Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionOptionsWeb3Tx) ProtoMessage() {}

// Deprecated: Use ExtensionOptionsWeb3Tx.ProtoReflect.Descriptor instead.
func (*ExtensionOptionsWeb3Tx) Descriptor() ([]byte, []int) {
	return file_ethermint_types_v1_web3_proto_rawDescGZIP(), []int{0}
}

func (x *ExtensionOptionsWeb3Tx) GetTypedDataChainId() uint64 {
	if x != nil {
		return x.TypedDataChainId
	}
	return 0
}

func (x *ExtensionOptionsWeb3Tx) GetFeePayer() string {
	if x != nil {
		return x.FeePayer
	}
	return ""
}

func (x *ExtensionOptionsWeb3Tx) GetFeePayerSig() []byte {
	if x != nil {
		return x.FeePayerSig
	}
	return nil
}

var File_ethermint_types_v1_web3_proto protoreflect.FileDescriptor

var file_ethermint_types_v1_web3_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x16, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57, 0x65,
	0x62, 0x33, 0x54, 0x78, 0x12, 0x61, 0x0a, 0x13, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x32, 0xe2, 0xde, 0x1f, 0x10, 0x54, 0x79, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0xea, 0xde, 0x1f, 0x1a, 0x74, 0x79, 0x70, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x2c, 0x6f, 0x6d, 0x69, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x10, 0x74, 0x79, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x66, 0x65, 0x65, 0x5f, 0x70,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xea, 0xde, 0x1f, 0x12,
	0x66, 0x65, 0x65, 0x50, 0x61, 0x79, 0x65, 0x72, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x08, 0x66, 0x65, 0x65, 0x50, 0x61, 0x79, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0d,
	0x66, 0x65, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x19, 0xea, 0xde, 0x1f, 0x15, 0x66, 0x65, 0x65, 0x50, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x69, 0x67, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x0b,
	0x66, 0x65, 0x65, 0x50, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69, 0x67, 0x3a, 0x04, 0x88, 0xa0, 0x1f,
	0x00, 0x42, 0xba, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x57, 0x65,
	0x62, 0x33, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x45, 0x54, 0x58, 0xaa, 0x02, 0x12, 0x45,
	0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x12, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x5c, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ethermint_types_v1_web3_proto_rawDescOnce sync.Once
	file_ethermint_types_v1_web3_proto_rawDescData = file_ethermint_types_v1_web3_proto_rawDesc
)

func file_ethermint_types_v1_web3_proto_rawDescGZIP() []byte {
	file_ethermint_types_v1_web3_proto_rawDescOnce.Do(func() {
		file_ethermint_types_v1_web3_proto_rawDescData = protoimpl.X.CompressGZIP(file_ethermint_types_v1_web3_proto_rawDescData)
	})
	return file_ethermint_types_v1_web3_proto_rawDescData
}

var file_ethermint_types_v1_web3_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ethermint_types_v1_web3_proto_goTypes = []interface{}{
	(*ExtensionOptionsWeb3Tx)(nil), // 0: ethermint.types.v1.ExtensionOptionsWeb3Tx
}
var file_ethermint_types_v1_web3_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ethermint_types_v1_web3_proto_init() }
func file_ethermint_types_v1_web3_proto_init() {
	if File_ethermint_types_v1_web3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ethermint_types_v1_web3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionOptionsWeb3Tx); i {
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
			RawDescriptor: file_ethermint_types_v1_web3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ethermint_types_v1_web3_proto_goTypes,
		DependencyIndexes: file_ethermint_types_v1_web3_proto_depIdxs,
		MessageInfos:      file_ethermint_types_v1_web3_proto_msgTypes,
	}.Build()
	File_ethermint_types_v1_web3_proto = out.File
	file_ethermint_types_v1_web3_proto_rawDesc = nil
	file_ethermint_types_v1_web3_proto_goTypes = nil
	file_ethermint_types_v1_web3_proto_depIdxs = nil
}
