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

package segment

import (
	"io"
)

type PayloadCompressor interface {

	// Compress compresses the source, reading it fully, and writes the compressed result to dest.
	Compress(source io.Reader, dest io.Writer) error

	// Decompress decompresses the source, reading it fully, and writes the decompressed result to dest.
	Decompress(source io.Reader, dest io.Writer) error
}
