# Model
model:
  rest_name: subtask
  resource_name: subtasks
  entity_name: Subtask
  friendly_name: Subtask
  package: todo-list
  group: core
  description: Represent a subtask object.
  get:
    description: Retrieve the subtask object with the given ID.
  update:
    description: Updates the subtask object with the given ID.
  delete:
    description: Deletes the subtask object with the given ID.
  extends:
  - '@base'

# Attributes
attributes:
  v1:
  - name: Name
    friendly_name: Name
    description: The Name.
    type: string
    exposed: true
    stored: true
    filterable: true
    orderable: true
