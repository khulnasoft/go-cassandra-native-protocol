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

	"github.com/khulnasoft/go-cassandra-native-protocol/primitive"
)

func TestListType(t *testing.T) {
	ListType := NewList(Varchar)
	assert.Equal(t, primitive.DataTypeCodeList, ListType.Code())
	assert.Equal(t, Varchar, ListType.ElementType)
}

func TestListTypeDeepCopy(t *testing.T) {
	lt := NewList(Varchar)
	clonedObj := lt.DeepCopy()
	assert.Equal(t, lt, clonedObj)
	clonedObj.ElementType = Int
	assert.Equal(t, primitive.DataTypeCodeList, lt.Code())
	assert.Equal(t, Varchar, lt.ElementType)
	assert.Equal(t, primitive.DataTypeCodeList, clonedObj.Code())
	assert.Equal(t, Int, clonedObj.ElementType)
}

func TestWriteListType(t *testing.T) {
	for _, version := range primitive.SupportedProtocolVersions() {
		t.Run(version.String(), func(t *testing.T) {
			tests := []struct {
				name     string
				input    DataType
				expected []byte
				err      error
			}{
				{
					"simple list",
					NewList(Varchar),
					[]byte{0, byte(primitive.DataTypeCodeVarchar & 0xFF)},
					nil,
				},
				{
					"complex list",
					NewList(NewList(Varchar)),
					[]byte{
						0, byte(primitive.DataTypeCodeList & 0xFF),
						0, byte(primitive.DataTypeCodeVarchar & 0xFF)},
					nil,
				},
				{"nil list", nil, nil, errors.New("expected *List, got <nil>")},
			}
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					var dest = &bytes.Buffer{}
					var err error
					err = writeListType(test.input, dest, version)
					actual := dest.Bytes()
					assert.Equal(t, test.expected, actual)
					assert.Equal(t, test.err, err)
				})
			}
		})
	}
}

func TestLengthOfListType(t *testing.T) {
	for _, version := range primitive.SupportedProtocolVersions() {
		t.Run(version.String(), func(t *testing.T) {
			tests := []struct {
				name     string
				input    DataType
				expected int
				err      error
			}{
				{"simple list", NewList(Varchar), primitive.LengthOfShort, nil},
				{"complex list", NewList(NewList(Varchar)), primitive.LengthOfShort + primitive.LengthOfShort, nil},
				{"nil list", nil, -1, errors.New("expected *List, got <nil>")},
			}
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					var actual int
					var err error
					actual, err = lengthOfListType(test.input, version)
					assert.Equal(t, test.expected, actual)
					assert.Equal(t, test.err, err)
				})
			}
		})
	}
}

func TestReadListType(t *testing.T) {
	for _, version := range primitive.SupportedProtocolVersions() {
		t.Run(version.String(), func(t *testing.T) {
			tests := []struct {
				name     string
				input    []byte
				expected DataType
				err      error
			}{
				{
					"simple list",
					[]byte{0, byte(primitive.DataTypeCodeVarchar & 0xff)},
					NewList(Varchar),
					nil,
				},
				{
					"complex list",
					[]byte{
						0, byte(primitive.DataTypeCodeList & 0xff),
						0, byte(primitive.DataTypeCodeVarchar & 0xff)},
					NewList(NewList(Varchar)),
					nil,
				},
				{
					"cannot read list",
					[]byte{},
					nil,
					fmt.Errorf("cannot read list element type: %w",
						fmt.Errorf("cannot read data type code: %w",
							fmt.Errorf("cannot read [short]: %w",
								errors.New("EOF")))),
				},
			}
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					var source = bytes.NewBuffer(test.input)
					var actual DataType
					var err error
					actual, err = readListType(source, version)
					assert.Equal(t, test.expected, actual)
					assert.Equal(t, test.err, err)
				})
			}
		})
	}
}
