# API Testing Guide

This guide provides instructions on how to test the API endpoints both locally and on the live backend.

## Base URLs

- **Localhost**: `http://localhost:8082` (default)
- **Production**: `https://golang-learning.onrender.com`

---

## 1. Health Check
Verify the server is running.

### Localhost
```bash
curl -i http://localhost:8082/health
```

### Production
```bash
curl -i https://golang-learning.onrender.com/health
```

**Expected Response:**
```json
{"status":"ok"}
```

---

## 2. Create Student
Add a new student record.

### Localhost
```bash
curl -i -X POST http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "age": 20}'
```

### Production
```bash
curl -i -X POST https://golang-learning.onrender.com/api/students \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "age": 20}'
```

**Expected Response:** `201 Created` with the new student ID.

---

## 3. Get Student by ID
Retrieve a student by their ID. Replace `{id}` with the actual ID returned from the creation step.

### Localhost
```bash
curl -i http://localhost:8082/api/students/1
```

### Production
```bash
curl -i https://golang-learning.onrender.com/api/students/1
```

**Expected Response:** `200 OK` with student details.

---

## 4. Update Student
Update an existing student's information.

### Localhost
```bash
curl -i -X PUT http://localhost:8082/api/students/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "email": "jane@example.com", "age": 21}'
```

### Production
```bash
curl -i -X PUT https://golang-learning.onrender.com/api/students/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "email": "jane@example.com", "age": 21}'
```

**Expected Response:** `200 OK` with `{"status": "updated"}`.

---

## 5. Delete Student
Delete a student record.

### Localhost
```bash
curl -i -X DELETE http://localhost:8082/api/students/1
```

### Production
```bash
curl -i -X DELETE https://golang-learning.onrender.com/api/students/1
```

**Expected Response:** `200 OK` with `{"status": "deleted"}`.
