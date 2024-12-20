// Copyright 2020 KhulnaSoft
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

package datatype

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/khulnasoft/go-cassandra-native-protocol/primitive"
)

func TestUserDefinedType(t *testing.T) {
	fieldNames := []string{"f1", "f2"}
	fieldTypes := []DataType{Varchar, Int}
	udtType, err := NewUserDefined("ks1", "udt1", fieldNames, fieldTypes)
	assert.Nil(t, err)
	assert.Equal(t, primitive.DataTypeCodeUdt, udtType.Code())
	assert.Equal(t, fieldTypes, udtType.FieldTypes)
	udtType2, err2 := NewUserDefined("ks1", "udt1", fieldNames, []DataType{Varchar, Int, Boolean})
	assert.Nil(t, udtType2)
	assert.Errorf(t, err2, "field names and field types length mismatch: 2 != 3")
}

func TestUserDefinedTypeDeepCopy(t *testing.T) {
	fieldNames := []string{"f1", "f2"}
	fieldTypes := []DataType{Varchar, Int}
	udtType, err := NewUserDefined("ks1", "udt1", fieldNames, fieldTypes)
	assert.Nil(t, err)

	cloned := udtType.DeepCopy()
	assert.Equal(t, udtType, cloned)
	cloned.Name = "udt2"
	cloned.Keyspace = "ks2"
	cloned.FieldNames = []string{"f5", "field6", "f7"}
	cloned.FieldTypes = []DataType{Uuid, Float, Varchar}
	assert.NotEqual(t, udtType, cloned)

	assert.Equal(t, primitive.DataTypeCodeUdt, udtType.Code())
	assert.Equal(t, []DataType{Varchar, Int}, udtType.FieldTypes)
	assert.Equal(t, []string{"f1", "f2"}, udtType.FieldNames)
	assert.Equal(t, "ks1", udtType.Keyspace)
	assert.Equal(t, "udt1", udtType.Name)

	assert.Equal(t, primitive.DataTypeCodeUdt, cloned.Code())
	assert.Equal(t, []DataType{Uuid, Float, Varchar}, cloned.FieldTypes)
	assert.Equal(t, []string{"f5", "field6", "f7"}, cloned.FieldNames)
	assert.Equal(t, "ks2", cloned.Keyspace)
	assert.Equal(t, "udt2", cloned.Name)
}

func TestUserDefinedTypeDeepCopy_NilFieldTypesSlice(t *testing.T) {
	fieldNames := []string{"f1", "f2", "f3"}
	fieldTypes := []DataType{Int, Uuid, Float}
	udtType, err := NewUserDefined("ks1", "udt1", fieldNames, fieldTypes)
	assert.Nil(t, err)
	udtType.FieldTypes = nil

	cloned := udtType.DeepCopy()
	assert.Equal(t, udtType, cloned)
	cloned.FieldTypes = []DataType{Uuid, Float, Varchar}
	assert.NotEqual(t, udtType, cloned)

	assert.Nil(t, udtType.FieldTypes)
	assert.Equal(t, []DataType{Uuid, Float, Varchar}, cloned.FieldTypes)
}

func TestUserDefinedTypeDeepCopy_NilFieldType(t *testing.T) {
	fieldNames := []string{"f1", "f2", "f3"}
	fieldTypes := []DataType{nil, Uuid, Float}
	udtType, err := NewUserDefined("ks1", "udt1", fieldNames, fieldTypes)
	assert.Nil(t, err)

	cloned := udtType.DeepCopy()
	assert.Equal(t, udtType, cloned)
	cloned.FieldTypes = []DataType{Uuid, Float, Varchar}
	assert.NotEqual(t, udtType, cloned)

	assert.Equal(t, []DataType{nil, Uuid, Float}, udtType.FieldTypes)
	assert.Equal(t, []DataType{Uuid, Float, Varchar}, cloned.FieldTypes)
}

func TestUserDefinedTypeDeepCopy_ComplexFieldTypes(t *testing.T) {
	fieldNames := []string{"f1", "f2", "f3"}
	fieldTypes := []DataType{NewList(NewTuple(Varchar)), Uuid, Float}
	udtType, err := NewUserDefined("ks1", "udt1", fieldNames, fieldTypes)
	assert.Nil(t, err)

	cloned := udtType.DeepCopy()
	assert.Equal(t, udtType, cloned)
	cloned.FieldTypes[0].(*List).ElementType = NewTuple(Int)
	assert.NotEqual(t, udtType, cloned)

	assert.Equal(t, []DataType{NewList(NewTuple(Varchar)), Uuid, Float}, udtType.FieldTypes)
	assert.Equal(t, []DataType{NewList(NewTuple(Int)), Uuid, Float}, cloned.FieldTypes)
}

var udt1, _ = NewUserDefined("ks1", "udt1", []string{"f1", "f2"}, []DataType{Varchar, Int})
var udt2, _ = NewUserDefined("ks1", "udt2", []string{"f1"}, []DataType{udt1})

func TestWriteUserDefinedType(t *testing.T) {
	tests := []struct {
		name     string
		input    DataType
		expected []byte
		err      error
	}{
		{
			"simple udt",
			udt1,
			[]byte{
				0, byte(primitive.DataTypeCodeUdt & 0xFF),
				0, 3, byte('k'), byte('s'), byte('1'),
				0, 4, byte('u'), byte('d'), byte('t'), byte('1'),
				0, 2, // field count
				0, 2, byte('f'), byte('1'),
				0, byte(primitive.DataTypeCodeVarchar & 0xFF),
				0, 2, byte('f'), byte('2'),
				0, byte(primitive.DataTypeCodeInt & 0xFF),
			},
			nil,
		},
		{
			"complex udt",
			udt2,
			[]byte{
				0, byte(primitive.DataTypeCodeUdt & 0xFF),
				0, 3, byte('k'), byte('s'), byte('1'),
				0, 4, byte('u'), byte('d'), byte('t'), byte('2'),
				0, 1, // field count
				0, 2, byte('f'), byte('1'),
				0, byte(primitive.DataTypeCodeUdt & 0xFF),
				0, 3, byte('k'), byte('s'), byte('1'),
				0, 4, byte('u'), byte('d'), byte('t'), byte('1'),
				0, 2, // field count
				0, 2, byte('f'), byte('1'),
				0, byte(primitive.DataTypeCodeVarchar & 0xFF),
				0, 2, byte('f'), byte('2'),
				0, byte(primitive.DataTypeCodeInt & 0xFF),
			},
			nil,
		},
		{"nil udt", nil, nil, errors.New("DataType can not be nil")},
	}

	t.Run("versions_with_udt_support", func(t *testing.T) {
		for _, version := range primitive.SupportedProtocolVersionsGreaterThanOrEqualTo(primitive.ProtocolVersion3) {
			t.Run(version.String(), func(t *testing.T) {
				for _, test := range tests {
					t.Run(test.name, func(t *testing.T) {
						var dest = &bytes.Buffer{}
						var err error
						err = WriteDataType(test.input, dest, version)
						assert.Equal(t, test.err, err)
						actual := dest.Bytes()
						assert.Equal(t, test.expected, actual)
					})
				}
			})
		}
	})
}

func TestLengthOfUserDefinedType(t *testing.T) {
	for _, version := range primitive.SupportedProtocolVersions() {
		t.Run(version.String(), func(t *testing.T) {
			tests := []struct {
				name     string
				input    DataType
				expected int
				err      error
			}{
				{
					"simple udt",
					udt1,
					primitive.LengthOfString("ks1") +
						primitive.LengthOfString("udt1") +
						primitive.LengthOfShort + // field count
						primitive.LengthOfString("f1") +
						primitive.LengthOfShort + // varchar
						primitive.LengthOfString("f2") +
						primitive.LengthOfShort, // int
					nil,
				},
				{
					"complex udt",
					udt2,
					primitive.LengthOfString("ks1") +
						primitive.LengthOfString("udt2") +
						primitive.LengthOfShort + // field count
						primitive.LengthOfString("f1") +
						primitive.LengthOfShort + // UDT
						primitive.LengthOfString("ks1") +
						primitive.LengthOfString("udt1") +
						primitive.LengthOfShort + // field count
						primitive.LengthOfString("f1") +
						primitive.LengthOfShort + // varchar
						primitive.LengthOfString("f2") +
						primitive.LengthOfShort, // int
					nil,
				},
				{"nil udt", nil, -1, errors.New("expected *UserDefined, got <nil>")},
			}
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					var actual int
					var err error
					actual, err = lengthOfUserDefinedType(test.input, version)
					assert.Equal(t, test.expected, actual)
					assert.Equal(t, test.err, err)
				})
			}
		})
	}
}

func TestReadUserDefinedType(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected DataType
		err      error
	}{
		{
			"simple udt",
			[]byte{
				0, byte(primitive.DataTypeCodeUdt & 0xFF),
				0, 3, byte('k'), byte('s'), byte('1'),
				0, 4, byte('u'), byte('d'), byte('t'), byte('1'),
				0, 2, // field count
				0, 2, byte('f'), byte('1'),
				0, byte(primitive.DataTypeCodeVarchar & 0xFF),
				0, 2, byte('f'), byte('2'),
				0, byte(primitive.DataTypeCodeInt & 0xFF),
			},
			udt1,
			nil,
		},
		{
			"complex udt",
			[]byte{
				0, byte(primitive.DataTypeCodeUdt & 0xFF),
				0, 3, byte('k'), byte('s'), byte('1'),
				0, 4, byte('u'), byte('d'), byte('t'), byte('2'),
				0, 1, // field count
				0, 2, byte('f'), byte('1'),
				0, byte(primitive.DataTypeCodeUdt & 0xFF),
				0, 3, byte('k'), byte('s'), byte('1'),
				0, 4, byte('u'), byte('d'), byte('t'), byte('1'),
				0, 2, // field count
				0, 2, byte('f'), byte('1'),
				0, byte(primitive.DataTypeCodeVarchar & 0xFF),
				0, 2, byte('f'), byte('2'),
				0, byte(primitive.DataTypeCodeInt & 0xFF),
			},
			udt2,
			nil,
		},
		{
			"cannot read udt",
			[]byte{0, byte(primitive.DataTypeCodeUdt & 0xFF)},
			nil,
			fmt.Errorf("cannot read udt keyspace: %w",
				fmt.Errorf("cannot read [string] length: %w",
					fmt.Errorf("cannot read [short]: %w",
						errors.New("EOF")))),
		},
	}

	t.Run("versions_with_udt_support", func(t *testing.T) {
		for _, version := range primitive.SupportedProtocolVersionsGreaterThanOrEqualTo(primitive.ProtocolVersion3) {
			t.Run(version.String(), func(t *testing.T) {
				for _, test := range tests {
					t.Run(test.name, func(t *testing.T) {
						var source = bytes.NewBuffer(test.input)
						var actual DataType
						var err error
						actual, err = ReadDataType(source, version)
						assert.Equal(t, test.err, err)
						assert.Equal(t, test.expected, actual)
					})
				}
			})
		}
	})
}

func Test_userDefinedType_String(t1 *testing.T) {
	tests := []struct {
		name       string
		keyspace   string
		udtName    string
		fieldNames []string
		fieldTypes []DataType
		want       string
	}{
		{"empty", "ks1", "type1", []string{}, []DataType{}, "ks1.type1<>"},
		{"simple", "ks1", "type1", []string{"f1", "f2"}, []DataType{Int, Varchar}, "ks1.type1<f1:int,f2:varchar>"},
		{
			"complex",
			"ks1",
			"type1",
			[]string{"f1", "f2"},
			[]DataType{Int, func() DataType {
				udt2, _ := NewUserDefined("ks1", "type2", []string{"f2a", "f2b"}, []DataType{Varchar, Boolean})
				return udt2
			}()},
			"ks1.type1<f1:int,f2:ks1.type2<f2a:varchar,f2b:boolean>>",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t *testing.T) {
			udt, err := NewUserDefined(tt.keyspace, tt.udtName, tt.fieldNames, tt.fieldTypes)
			require.NoError(t, err)
			got := udt.AsCql()
			assert.Equal(t, tt.want, got)
		})
	}
}
