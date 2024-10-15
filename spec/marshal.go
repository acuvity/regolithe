// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spec

import (
	"bytes"
	"encoding/json"
)

// JSONMarshal does what a normal json.Marshal does but with EscapeHTML set to false.
func JSONMarshal(v any) ([]byte, error) {

	buffer := &bytes.Buffer{}

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	return bytes.TrimSuffix(buffer.Bytes(), []byte("\n")), nil
}

// JSONMarshal does what a normal json.MarshalIndent does but with EscapeHTML set to false.
func JSONMarshalIndent(v any, prefix string, indent string) ([]byte, error) {

	buffer := &bytes.Buffer{}

	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(prefix, indent)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	return bytes.TrimSuffix(buffer.Bytes(), []byte("\n")), nil
}
