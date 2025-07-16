# Task Manager API Documentation

This document describes the endpoints for the Task Manager REST API.

---

## 🔹 GET /tasks

**Description:**
Returns all tasks in the system.

**Method:** GET
**URL:** `/tasks`

**Response:**

```json
[
  {
    "id": "1",
    "title": "Learn Go",
    "status": "not_started",
    "due_date": "2025-07-23T00:00:00Z"
  }
]
```

---

## 🔹 GET /tasks/\:id

**Description:**
Returns a task by its ID.

**Method:** GET
**URL:** `/tasks/{id}`

**Responses:**

✅ **200 OK**

```json
{
  "id": "1",
  "title": "Learn Go",
  "status": "not_started",
  "due_date": "2025-07-23T00:00:00Z"
}
```

❌ **404 Not Found**

```json
{
  "message": "Task not found"
}
```

---

## 🔹 POST /tasks

**Description:**
Creates a new task.

**Method:** POST
**URL:** `/tasks`
**Request Body (JSON):**

```json
{
  "id": "5",
  "title": "Write Documentation",
  "status": "not_started",
  "due_date": "2025-08-01T00:00:00Z"
}
```

**Responses:**

✅ **201 Created**

```json
{
  "id": "5",
  "title": "Write Documentation",
  "status": "not_started",
  "due_date": "2025-08-01T00:00:00Z"
}
```

❌ **400 Bad Request**

```json
{
  "error": "Invalid request body"
}
```

---

## 🔹 PUT /tasks/\:id

**Description:**
Updates an existing task. Only send the fields you want to update.

**Method:** PUT
**URL:** `/tasks/{id}`
**Request Body (JSON):**

```json
{
  "title": "Update Docs",
  "status": "in_progress"
}
```

**Responses:**

✅ **200 OK**

```json
{
  "message": "Task updated"
}
```

❌ **404 Not Found**

```json
{
  "message": "Task not found"
}
```

---

## 🔹 DELETE /tasks/\:id

**Description:**
Deletes a task by ID.

**Method:** DELETE
**URL:** `/tasks/{id}`

**Responses:**

✅ **200 OK**

```json
{
  "message": "Deleted successfully"
}
```

❌ **404 Not Found**

```json
{
  "message": "Task not found"
}
```


 ## The postman documentation 
 

 [Link here](https://web.postman.co/workspace/My-Workspace~06c69e86-ca8d-4725-815b-5fedfd883ba2/collection/46778945-d905b1fa-5b0c-4636-a4d3-f82cfc3654a2?action=share&source=copy-link&creator=46778945)