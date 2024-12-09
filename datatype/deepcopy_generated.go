//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 KhulnaSoft
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package datatype

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Custom) DeepCopyInto(out *Custom) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Custom.
func (in *Custom) DeepCopy() *Custom {
	if in == nil {
		return nil
	}
	out := new(Custom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyDataType is an autogenerated deepcopy function, copying the receiver, creating a new DataType.
func (in *Custom) DeepCopyDataType() DataType {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *List) DeepCopyInto(out *List) {
	*out = *in
	if in.ElementType != nil {
		out.ElementType = in.ElementType.DeepCopyDataType()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new List.
func (in *List) DeepCopy() *List {
	if in == nil {
		return nil
	}
	out := new(List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyDataType is an autogenerated deepcopy function, copying the receiver, creating a new DataType.
func (in *List) DeepCopyDataType() DataType {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Map) DeepCopyInto(out *Map) {
	*out = *in
	if in.KeyType != nil {
		out.KeyType = in.KeyType.DeepCopyDataType()
	}
	if in.ValueType != nil {
		out.ValueType = in.ValueType.DeepCopyDataType()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Map.
func (in *Map) DeepCopy() *Map {
	if in == nil {
		return nil
	}
	out := new(Map)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyDataType is an autogenerated deepcopy function, copying the receiver, creating a new DataType.
func (in *Map) DeepCopyDataType() DataType {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimitiveType) DeepCopyInto(out *PrimitiveType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimitiveType.
func (in *PrimitiveType) DeepCopy() *PrimitiveType {
	if in == nil {
		return nil
	}
	out := new(PrimitiveType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyDataType is an autogenerated deepcopy function, copying the receiver, creating a new DataType.
func (in *PrimitiveType) DeepCopyDataType() DataType {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Set) DeepCopyInto(out *Set) {
	*out = *in
	if in.ElementType != nil {
		out.ElementType = in.ElementType.DeepCopyDataType()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Set.
func (in *Set) DeepCopy() *Set {
	if in == nil {
		return nil
	}
	out := new(Set)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyDataType is an autogenerated deepcopy function, copying the receiver, creating a new DataType.
func (in *Set) DeepCopyDataType() DataType {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tuple) DeepCopyInto(out *Tuple) {
	*out = *in
	if in.FieldTypes != nil {
		in, out := &in.FieldTypes, &out.FieldTypes
		*out = make([]DataType, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				(*out)[i] = (*in)[i].DeepCopyDataType()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tuple.
func (in *Tuple) DeepCopy() *Tuple {
	if in == nil {
		return nil
	}
	out := new(Tuple)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyDataType is an autogenerated deepcopy function, copying the receiver, creating a new DataType.
func (in *Tuple) DeepCopyDataType() DataType {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserDefined) DeepCopyInto(out *UserDefined) {
	*out = *in
	if in.FieldNames != nil {
		in, out := &in.FieldNames, &out.FieldNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FieldTypes != nil {
		in, out := &in.FieldTypes, &out.FieldTypes
		*out = make([]DataType, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				(*out)[i] = (*in)[i].DeepCopyDataType()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserDefined.
func (in *UserDefined) DeepCopy() *UserDefined {
	if in == nil {
		return nil
	}
	out := new(UserDefined)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyDataType is an autogenerated deepcopy function, copying the receiver, creating a new DataType.
func (in *UserDefined) DeepCopyDataType() DataType {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
