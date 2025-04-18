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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpecification_NewSpecification(t *testing.T) {

	Convey("Given I create a new specification", t, func() {

		spec := NewSpecification()

		Convey("Then the spec should be correctly initialized", func() {
			So(spec, ShouldNotBeNil)
		})
	})
}

func TestSpecification_Validate(t *testing.T) {

	Convey("Given I have a specification with no validation error", t, func() {

		s := &specification{
			RawModel: &Model{
				ResourceName: "things",
				RestName:     "thing",
				Description:  "desc.",
				EntityName:   "toto",
				Package:      "package",
				Group:        "core",
			},
			RawAttributes: map[string][]*Attribute{
				"v1": {
					{
						Name:        "attr1",
						Description: "desc.",
						Type:        "string",
					},
				},
			},
		}

		Convey("When I call validate", func() {

			err := s.Validate()

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a specification with validation error", t, func() {

		s := &specification{
			RawModel: &Model{
				ResourceName: "things",
				RestName:     "thing",
				Description:  "desc.",
				EntityName:   "toto",
				Group:        "core",
			},
			RawAttributes: map[string][]*Attribute{
				"v1": {
					{
						Identifier: false,
						Name:       "not-id",
					},
					{
						Name:        "id",
						Type:        "coucou",
						Description: "wala",
					},
				},
			},
		}

		Convey("When I call validate", func() {

			errs := s.Validate()

			Convey("Then err should not be nil", func() {
				So(errs, ShouldNotBeNil)
				So(formatValidationErrors(errs).Error(), ShouldEqual, `thing.spec: schema error: attributes.v1.0: description is required
thing.spec: schema error: attributes.v1.0: type is required
thing.spec: schema error: attributes.v1.1.type: attributes.v1.1.type must be one of the following: "string", "integer", "float", "boolean", "enum", "list", "object", "time", "external", "ref", "refList", "refMap"
thing.spec: schema error: model: package is required`)
			})
		})
	})

	Convey("Given I have an abstract with no validation error", t, func() {

		s := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": {
					{
						Name:        "attr1",
						Description: "desc.",
						Type:        "string",
					},
				},
			},
		}

		Convey("When I call validate", func() {

			err := s.Validate()

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given I have an abstract with validation error", t, func() {

		s := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": {
					{
						Name:        "attr1",
						Description: "desc.",
						Type:        "string",
					},
				},
			},
			RawRelations: []*Relation{
				{
					RestName: "a",
				},
			},
		}

		Convey("When I call validate", func() {

			err := s.Validate()

			Convey("Then err should not be nil", func() {
				So(len(err), ShouldEqual, 1)
				So(err[0].Error(), ShouldEqual, ".: schema error: (root): Additional property relations is not allowed")
			})
		})
	})
}

func TestSpecification_Getters(t *testing.T) {

	Convey("Given I have a new API", t, func() {

		s := &specification{
			RawModel: &Model{
				EntityName:   "Test",
				ResourceName: "tests",
				RestName:     "test",
				Delete:       &RelationAction{},
				Get:          &RelationAction{},
				Update:       &RelationAction{},
			},
			RawAttributes: map[string][]*Attribute{
				"v1": {
					{
						Identifier: false,
						Name:       "not-id",
					},
					{
						Identifier: true,
						Name:       "id",
					},
				},
			},
		}

		s.buildAttributesMapping() // nolint: errcheck

		Convey("Then the getters should work", func() {
			So(s.Identifier().Name, ShouldEqual, "id")
		})
	})
}

func TestSpecification_TypeProviders(t *testing.T) {

	Convey("Given I have a new API", t, func() {

		s := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": {
					{
						Name:          "not-id",
						ConvertedName: "not-id",
						TypeProvider:  "toto",
					},
					{
						ConvertedName: "id",
						TypeProvider:  "titi",
					},
					{
						ConvertedName: "id2",
						TypeProvider:  "titi",
					},
					{},
				},
			},
		}

		Convey("When I call TypeProviders", func() {

			providers := s.TypeProviders()

			Convey("Then the providers should be correct", func() {
				So(providers, ShouldResemble, []string{"toto", "titi"})
			})
		})
	})
}

func TestSpecification_AttributesProviders(t *testing.T) {

	Convey("Given I have a new API", t, func() {

		s := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": {
					{
						ValidationProviders: map[string]*ValidationMap{
							"a": {
								Import: "a",
							},
							"b": {
								Import: "b",
							},
						},
					},
					{
						ValidationProviders: map[string]*ValidationMap{
							"c": {
								Import: "c",
							},
							"b": {
								Import: "b",
							},
						},
					},
					{
						ValidationProviders: map[string]*ValidationMap{
							"d": {
								Import: "d",
							},
							"b": {
								Import: "b",
							},
						},
					},
					{},
				},
			},
		}

		Convey("When I call ValidationProviders", func() {

			providers := s.ValidationProviders()

			Convey("Then the providers should be correct", func() {
				So(providers, ShouldResemble, []string{"a", "b", "c", "d"})
			})
		})
	})
}

func TestSpecification_AttributeMap(t *testing.T) {

	Convey("Given I load the task specification file", t, func() {

		spec, err := LoadSpecification("./tests/task.spec", true)

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Then the attribute map should be correctly built", func() {
			So(len(spec.Attributes("v1")), ShouldEqual, 4)
			So(spec.Attribute("name", "v1").Name, ShouldEqual, "name")
			So(spec.Attribute("description", "v1").Name, ShouldEqual, "description")
			So(spec.Attribute("status", "v1").Name, ShouldEqual, "status")
		})
	})
}

func TestSpecification_buildAttributesMapping(t *testing.T) {

	Convey("Given I create a specification with the same attribute twice.", t, func() {

		spec := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": {
					{
						Name: "a",
					},
					{
						Name: "a",
					},
				},
			},
		}

		Convey("When I call buildAttributesMapping", func() {

			err := spec.buildAttributesMapping()

			Convey("Then err Should Not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestSpecification_APIMap(t *testing.T) {

	Convey("Given I load the root specification file", t, func() {

		spec, err := LoadSpecification("./tests/root.spec", true)

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Then the relation map should be correctly built", func() {
			So(len(spec.Relations()), ShouldEqual, 2)
			So(spec.Relation("list").RestName, ShouldEqual, "list")
			So(spec.Relation("user").RestName, ShouldEqual, "user")
		})
	})
}

func TestSpecification_buildRelationsMapping(t *testing.T) {

	Convey("Given I create a specification with the same relation twice.", t, func() {

		spec := &specification{
			RawRelations: []*Relation{
				{
					RestName: "a",
				},
				{
					RestName: "a",
				},
			},
		}

		Convey("When I call buildRelationsMapping", func() {

			err := spec.buildRelationsMapping()

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestSpecification_LoadSpecification(t *testing.T) {

	Convey("Given I load a non existing specification file", t, func() {

		_, err := LoadSpecification("./tests/not.spec", true)

		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Given I load a bad formatted specification file", t, func() {

		_, err := LoadSpecification("./tests/task.spec.bad", true)

		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Given I load the root specification file", t, func() {

		spec, err := LoadSpecification("./tests/root.spec", true)
		rels := spec.Relations()

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Then the spec should be correctly initialized", func() {

			model := spec.Model()

			So(model.Get, ShouldNotBeNil)
			So(model.Delete, ShouldBeNil)
			So(model.Update, ShouldBeNil)
			So(model.Description, ShouldEqual, "Root object of the API.")
			So(model.EntityName, ShouldEqual, "Root")
			So(model.Package, ShouldEqual, "todo-list")
			So(model.ResourceName, ShouldEqual, "root")
			So(model.RestName, ShouldEqual, "root")
			So(model.Extends, ShouldBeNil)
			So(model.IsRoot, ShouldBeTrue)
			So(model.Aliases, ShouldBeNil)
		})

		Convey("Then the number of relations should be correct", func() {
			So(len(rels), ShouldEqual, 2)
		})

		Convey("Then the list of relations should be correct", func() {
			So(rels[0].Create, ShouldNotBeNil)
			So(rels[0].Delete, ShouldBeNil)
			So(rels[0].Update, ShouldBeNil)
			So(rels[0].Get, ShouldNotBeNil)
			So(rels[0].RestName, ShouldEqual, "list")
		})

		Convey("Then the user relation should be correct", func() {
			So(rels[1].Create, ShouldNotBeNil)
			So(rels[1].Delete, ShouldBeNil)
			So(rels[1].Update, ShouldBeNil)
			So(rels[1].Get, ShouldNotBeNil)
			So(rels[1].RestName, ShouldEqual, "user")
		})

	})

	Convey("Given I load the task specification file", t, func() {

		spec, err := LoadSpecification("./tests/task.spec", true)
		attrs := spec.Attributes("v1")
		model := spec.Model()

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Then the spec should be correctly initialized", func() {
			So(model.Get, ShouldNotBeNil)
			So(model.Delete, ShouldNotBeNil)
			So(model.Update, ShouldNotBeNil)
			So(model.Description, ShouldEqual, "Represent a task to do in a listd.")
			So(model.EntityName, ShouldEqual, "Task")
			So(model.Package, ShouldEqual, "todo-list")
			So(model.ResourceName, ShouldEqual, "tasks")
			So(model.RestName, ShouldEqual, "task")
			So(model.Extends, ShouldResemble, []string{"@base"})
			So(model.IsRoot, ShouldBeFalse)
			So(model.Aliases, ShouldResemble, []string{"tsk"})
		})

		Convey("Then the number of attributes should be correct", func() {
			So(len(spec.Attributes("v1")), ShouldEqual, 4)
		})

		Convey("Then the spec attribute description be correctly initialized", func() {

			So(attrs[0].AllowedChars, ShouldBeEmpty)
			So(attrs[0].AllowedChoices, ShouldBeEmpty)
			So(attrs[0].Autogenerated, ShouldBeFalse)
			So(attrs[0].CreationOnly, ShouldBeFalse)
			So(attrs[0].DefaultValue, ShouldBeNil)
			So(attrs[0].Deprecated, ShouldBeFalse)
			So(attrs[0].Description, ShouldEqual, "The description.")
			So(attrs[0].Exposed, ShouldBeTrue)
			So(attrs[0].Filterable, ShouldBeTrue)
			So(attrs[0].ForeignKey, ShouldBeFalse)
			So(attrs[0].Getter, ShouldBeFalse)
			So(attrs[0].Identifier, ShouldBeFalse)
			So(attrs[0].MaxLength, ShouldEqual, 0)
			So(attrs[0].MaxValue, ShouldEqual, 0.0)
			So(attrs[0].MinLength, ShouldEqual, 0)
			So(attrs[0].MinValue, ShouldEqual, 0.0)
			So(attrs[0].Name, ShouldEqual, "description")
			So(attrs[0].Orderable, ShouldBeTrue)
			So(attrs[0].PrimaryKey, ShouldBeFalse)
			So(attrs[0].ReadOnly, ShouldBeFalse)
			So(attrs[0].Required, ShouldBeFalse)
			So(attrs[0].Secret, ShouldBeFalse)
			So(attrs[0].Setter, ShouldBeFalse)
			So(attrs[0].Stored, ShouldBeTrue)
			So(attrs[0].SubType, ShouldEqual, "")
			So(attrs[0].Transient, ShouldBeFalse)
			So(attrs[0].Type, ShouldEqual, AttributeTypeString)
		})

		Convey("Then the spec attribute name be correctly initialized", func() {

			So(attrs[1].AllowedChars, ShouldBeEmpty)
			So(attrs[1].AllowedChoices, ShouldBeEmpty)
			So(attrs[1].Autogenerated, ShouldBeFalse)
			So(attrs[1].CreationOnly, ShouldBeFalse)
			So(attrs[1].DefaultValue, ShouldBeNil)
			So(attrs[1].Deprecated, ShouldBeFalse)
			So(attrs[1].Description, ShouldEqual, "The name.")
			So(attrs[1].Exposed, ShouldBeTrue)
			So(attrs[1].Filterable, ShouldBeTrue)
			So(attrs[1].ForeignKey, ShouldBeFalse)
			So(attrs[1].Getter, ShouldBeTrue)
			So(attrs[1].Identifier, ShouldBeFalse)
			So(attrs[1].MaxLength, ShouldEqual, 0)
			So(attrs[1].MaxValue, ShouldEqual, 0.0)
			So(attrs[1].MinLength, ShouldEqual, 0)
			So(attrs[1].MinValue, ShouldEqual, 0.0)
			So(attrs[1].Name, ShouldEqual, "name")
			So(attrs[1].Orderable, ShouldBeTrue)
			So(attrs[1].PrimaryKey, ShouldBeFalse)
			So(attrs[1].ReadOnly, ShouldBeFalse)
			So(attrs[1].Required, ShouldBeTrue)
			So(attrs[1].Secret, ShouldBeFalse)
			So(attrs[1].Setter, ShouldBeTrue)
			So(attrs[1].Stored, ShouldBeTrue)
			So(attrs[1].SubType, ShouldEqual, "")
			So(attrs[1].Transient, ShouldBeFalse)
			So(attrs[1].Type, ShouldEqual, AttributeTypeString)
		})

		Convey("Then the spec attribbute status be correctly initialized", func() {

			So(attrs[3].AllowedChars, ShouldBeEmpty)
			So(attrs[3].AllowedChoices, ShouldResemble, []string{"DONE", "PROGRESS", "TODO"})
			So(attrs[3].Autogenerated, ShouldBeFalse)
			So(attrs[3].CreationOnly, ShouldBeFalse)
			So(attrs[3].DefaultValue, ShouldResemble, "TODO")
			So(attrs[3].Deprecated, ShouldBeFalse)
			So(attrs[3].Description, ShouldEqual, "The status of the task.")
			So(attrs[3].Exposed, ShouldBeTrue)
			So(attrs[3].Filterable, ShouldBeTrue)
			So(attrs[3].ForeignKey, ShouldBeFalse)
			So(attrs[3].Getter, ShouldBeFalse)
			So(attrs[3].Identifier, ShouldBeFalse)
			So(attrs[3].MaxLength, ShouldEqual, 0)
			So(attrs[3].MaxValue, ShouldEqual, 0.0)
			So(attrs[3].MinLength, ShouldEqual, 0)
			So(attrs[3].MinValue, ShouldEqual, 0.0)
			So(attrs[3].Name, ShouldEqual, "status")
			So(attrs[3].Orderable, ShouldBeTrue)
			So(attrs[3].PrimaryKey, ShouldBeFalse)
			So(attrs[3].ReadOnly, ShouldBeFalse)
			So(attrs[3].Required, ShouldBeFalse)
			So(attrs[3].Secret, ShouldBeFalse)
			So(attrs[3].Setter, ShouldBeFalse)
			So(attrs[3].Stored, ShouldBeTrue)
			So(attrs[3].SubType, ShouldEqual, "")
			So(attrs[3].Transient, ShouldBeFalse)
			So(attrs[3].Type, ShouldEqual, AttributeTypeEnum)
		})

		Convey("When I apply the base specification", func() {

			base, err := LoadSpecification("./tests/@base.abs", true)
			spec.ApplyBaseSpecifications(base) // nolint: errcheck

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the number of attributes should be correct", func() {
				So(len(spec.Attributes("v1")), ShouldEqual, 7)
			})

			Convey("Then the additional attributes should have been applied", func() {
				So(spec.Attribute("ID", "v1").Name, ShouldEqual, "ID")
				So(spec.Attribute("ID", "v1").Autogenerated, ShouldBeTrue)
				So(spec.Attribute("ID", "v1").Description, ShouldEqual, "The identifier.")
				So(spec.Attribute("ID", "v1").Identifier, ShouldBeTrue)
				So(spec.Attribute("ID", "v1").PrimaryKey, ShouldBeTrue)
				So(spec.Attribute("ID", "v1").ReadOnly, ShouldBeTrue)

				So(spec.Attribute("parentID", "v1").Name, ShouldEqual, "parentID")
				So(spec.Attribute("parentID", "v1").Autogenerated, ShouldBeTrue)
				So(spec.Attribute("parentID", "v1").Description, ShouldEqual, "The identifier of the parent of the object.")
				So(spec.Attribute("parentID", "v1").ForeignKey, ShouldBeTrue)
				So(spec.Attribute("parentID", "v1").Identifier, ShouldBeFalse)
				So(spec.Attribute("parentID", "v1").PrimaryKey, ShouldBeFalse)
				So(spec.Attribute("parentID", "v1").ReadOnly, ShouldBeTrue)
			})
		})
	})
}

func TestSpecification_ApplyBaseSpecifications(t *testing.T) {

	Convey("Given I have a spec and an abstract", t, func() {

		s1 := &specification{
			RawModel: &Model{
				ResourceName: "things",
				RestName:     "thing",
				Description:  "desc.",
				EntityName:   "toto",
				Package:      "package",
			},
			RawAttributes: map[string][]*Attribute{
				"v1": {
					{
						Name:        "attr1",
						Description: "desc.",
						Type:        "string",
					},
				},
			},
		}

		abs := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": {
					{
						Name:        "attr1",
						Description: "desc from abs.",
						Type:        "string",
					},
					{
						Name:        "attr2",
						Description: "desc2",
						Type:        "string",
					},
				},
			},
		}

		Convey("When I call ApplyBaseSpecifications", func() {

			err := s1.ApplyBaseSpecifications(abs)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the extends should be applied", func() {
				So(len(s1.Attributes("v1")), ShouldEqual, 2)
				So(s1.Attribute("attr1", "v1").Name, ShouldEqual, "attr1")
				So(s1.Attribute("attr2", "v1").Name, ShouldEqual, "attr2")
			})
		})
	})

	Convey("Given I have a spec with no attributes and an abstract", t, func() {

		s1 := &specification{
			RawModel: &Model{
				ResourceName: "things",
				RestName:     "thing",
				Description:  "desc.",
				EntityName:   "toto",
				Package:      "package",
			},
		}

		abs := &specification{
			RawAttributes: map[string][]*Attribute{
				"v1": {
					{
						Name:        "attr2",
						Description: "desc2",
						Type:        "string",
					},
				},
			},
		}

		Convey("When I call ApplyBaseSpecifications", func() {

			err := s1.ApplyBaseSpecifications(abs)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the extends should be applied", func() {
				So(len(s1.Attributes("v1")), ShouldEqual, 1)
				So(s1.Attribute("attr2", "v1").Name, ShouldEqual, "attr2")
			})
		})
	})
}

func TestSpecifications_Versionning(t *testing.T) {

	Convey("Given I have a specification", t, func() {

		s := &specification{
			RawModel: &Model{
				ResourceName: "things",
				RestName:     "thing",
				Description:  "desc.",
				EntityName:   "toto",
				Package:      "package",
			},
			RawAttributes: map[string][]*Attribute{
				"v1": {
					{
						Name:        "1.1",
						Description: "desc.",
						Type:        "string",
					},
					{
						Name:        "1.2",
						Description: "desc.",
						Type:        "string",
					},
				},
				"v2": {
					{
						Name:        "2.1",
						Description: "desc.",
						Type:        "string",
					},
				},
				"v3": {
					{
						Name:        "3.1",
						Description: "desc.",
						Type:        "string",
					},
					{
						Name:        "3.2",
						Description: "desc.",
						Type:        "string",
					},
					{
						Name:        "1.1",
						Description: "desc.",
						Type:        "string",
					},
				},
			},
		}

		Convey("When I versionsFrom with v1", func() {

			versions := s.versionsFrom("v1")

			Convey("Then versions should be correct", func() {
				So(versions, ShouldResemble, []string{"v1"})
			})
		})

		Convey("When I versionsFrom with v2", func() {

			versions := s.versionsFrom("v2")

			Convey("Then versions should be correct", func() {
				So(versions, ShouldResemble, []string{"v1", "v2"})
			})
		})

		Convey("When I versionsFrom with v3", func() {

			versions := s.versionsFrom("v3")

			Convey("Then versions should be correct", func() {
				So(versions, ShouldResemble, []string{"v1", "v2", "v3"})
			})
		})

		Convey("When I versionsFrom with vNope", func() {

			Convey("Then it should panic", func() {
				So(func() { s.versionsFrom("vNope") }, ShouldPanicWith, "Invalid version 'vNope'")
			})
		})

		Convey("When I call Attributes on v1", func() {

			attrs := s.Attributes("v1")

			Convey("Then attributes should be correct", func() {
				So(len(attrs), ShouldEqual, 2)
				So(attrs[0].Name, ShouldEqual, "1.1")
				So(attrs[1].Name, ShouldEqual, "1.2")
			})
		})

		Convey("When I call Attributes on v2", func() {

			attrs := s.Attributes("v2")

			Convey("Then attributes should be correct", func() {
				So(len(attrs), ShouldEqual, 3)
				So(attrs[0].Name, ShouldEqual, "1.1")
				So(attrs[1].Name, ShouldEqual, "1.2")
				So(attrs[2].Name, ShouldEqual, "2.1")
			})
		})

		Convey("When I call Attributes on v3", func() {

			attrs := s.Attributes("v3")

			Convey("Then attributes should be correct", func() {
				So(len(attrs), ShouldEqual, 5)
				So(attrs[0].Name, ShouldEqual, "1.1")
				So(attrs[1].Name, ShouldEqual, "1.2")
				So(attrs[2].Name, ShouldEqual, "2.1")
				So(attrs[3].Name, ShouldEqual, "3.1")
				So(attrs[4].Name, ShouldEqual, "3.2")
			})
		})
	})
}

func TestSpecification_Write(t *testing.T) {

	Convey("Given I load the task specification file", t, func() {

		spec, err := LoadSpecification("./tests/list.spec", true)

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("When I call write", func() {
			buf := bytes.NewBuffer(nil)

			err = spec.Write(buf)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then buff should be correct", func() {
				So(buf.String(), ShouldEqual, `# Model
model:
  rest_name: list
  resource_name: lists
  entity_name: List
  friendly_name: List
  package: todo-list
  group: core
  description: Represent a a list of task to do.
  aliases:
  - lst
  get:
    description: Retrieves the list with the given ID.
    global_parameters:
    - sharedParameterA
    - sharedParameterB
    parameters:
      entries:
      - name: lgp1
        description: this is lgp1.
        type: string
        example_value: lgp1

      - name: lgp2
        description: this is lgp2.
        type: boolean
        example_value: "true"
  update:
    description: Updates the list with the given ID.
    parameters:
      entries:
      - name: lup1
        description: this is lup1.
        type: string
        example_value: lup1

      - name: lup2
        description: this is lup2.
        type: boolean
        example_value: "true"
  delete:
    description: Deletes the list with the given ID.
    parameters:
      entries:
      - name: ldp1
        description: this is ldp1.
        type: string
        example_value: ldp1

      - name: ldp2
        description: this is ldp2.
        type: boolean
        example_value: "true"
  extends:
  - '@base'

# Attributes
attributes:
  v1:
  - name: creationOnly
    friendly_name: CreationOnly
    description: This attribute is creation only.
    type: string
    exposed: true
    stored: true
    creation_only: true
    filterable: true
    orderable: true

  - name: date
    friendly_name: Date
    description: The date.
    type: time
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: description
    friendly_name: Description
    description: The description.
    type: string
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: name
    friendly_name: Name
    description: The name.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: the name
    filterable: true
    getter: true
    setter: true
    orderable: true

  - name: readOnly
    friendly_name: ReadOnly
    description: This attribute is readonly.
    type: string
    exposed: true
    stored: true
    read_only: true
    filterable: true
    orderable: true

  - name: ref
    friendly_name: Ref
    description: This attribute is a ref to a single object.
    type: ref
    subtype: task
    stored: true
    filterable: true
    orderable: true
    extensions:
      refMode: pointer

  - name: refList
    friendly_name: RefList
    description: This attribute is a ref to a objects.
    type: refList
    subtype: task
    stored: true
    filterable: true
    orderable: true

  - name: refMap
    friendly_name: RefMap
    description: This attribute is a ref map to a objects.
    type: refMap
    subtype: task
    stored: true
    filterable: true
    orderable: true
    extensions:
      refMode: pointer

  - name: secret
    friendly_name: Secret
    description: This attribute is secret.
    type: string
    exposed: true
    stored: true
    filterable: true
    orderable: true
    secret: true

  - name: slice
    friendly_name: Slice
    description: this is a slice.
    type: list
    exposed: true
    subtype: string
    stored: true
    filterable: true
    orderable: true

  - name: unexposed
    friendly_name: Unexposed
    description: This attribute is not exposed.
    type: string
    stored: true
    filterable: true
    orderable: true

# Relations
relations:
- rest_name: task
  get:
    description: yeye.
    parameters:
      entries:
      - name: ltgp1
        description: this is ltgp1.
        type: string
        example_value: ltgp1

      - name: ltgp2
        description: this is ltgp2.
        type: boolean
        example_value: "true"
  create:
    description: yoyo.
    parameters:
      entries:
      - name: ltcp1
        description: this is ltcp1.
        type: string
        example_value: ltcp1

      - name: ltcp2
        description: this is ltcp2.
        type: boolean
        example_value: "true"

- rest_name: user
  get:
    description: yeye.
`)
			})
		})
	})
}
