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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_JSONMarshal(t *testing.T) {

	Convey("Given I have a body with HTML content", t, func() {

		body := `{"error": "<< .Messages | join ", " >>"}`

		Convey("Then JSONMarshal should work", func() {
			data, err := JSONMarshal(body)

			So(err, ShouldBeNil)
			So(string(data), ShouldEqual, `"{\"error\": \"<< .Messages | join \", \" >>\"}"`)
		})
	})

	Convey("Given I have a malformed body of content", t, func() {

		Convey("Then JSONMarshal should not work", func() {
			_, err := JSONMarshal(make(chan int))

			So(err, ShouldNotBeNil)
		})
	})
}

func Test_JSONMarshalIndent(t *testing.T) {

	Convey("Given I have a functional body with HTML content", t, func() {

		body := `{"error": "<< .Messages | join ", " >>"}`

		Convey("Then JSONMarshalIndent should work", func() {
			data, err := JSONMarshalIndent(body, "", "  ")

			So(err, ShouldBeNil)
			So(string(data), ShouldEqual, `"{\"error\": \"<< .Messages | join \", \" >>\"}"`)
		})
	})

	Convey("Given I have a malformed body of content", t, func() {

		Convey("Then JSONMarshalIndent should not work", func() {
			_, err := JSONMarshalIndent(make(chan int), "", "  ")

			So(err, ShouldNotBeNil)
		})
	})
}
