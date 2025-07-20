# Task Management API with MongoDB

## Base URL
`http://localhost:8080`

## Endpoints

### GET /tasks
**Description**: Retrieve all tasks  
**Response**: 200 OK  
```
[
  {
    "id": "string",
    "title": "string",
    "due_date": "timestamp",
    "status": "pending"
  }
]
```

### GET /tasks/:id
**Description**: Retrieve a task by ID  
**Response**: 200 OK / 404 Not Found

### POST /tasks
**Description**: Create a new task  
**Body**:
```json
{
  "title": "New Task",
  "due_date": "2025-08-01T00:00:00Z",
  "status": "pending"
}
```
**Response**: 201 Created

### PUT /tasks/:id
**Description**: Update an existing task  
**Body**:
```json
{
  "title": "Updated Task",
  "status": "in_progress"
}
```
**Response**: 200 OK

### DELETE /tasks/:id
**Description**: Delete a task  
**Response**: 200 OK

## Status Enum
- `pending`
- `in_progress`
- `completed`
