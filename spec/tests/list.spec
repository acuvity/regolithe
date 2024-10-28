# Model
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
