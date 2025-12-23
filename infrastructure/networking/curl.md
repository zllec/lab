# cURL

- `curl -X POST https://example.com/api -H "Content-Type: application/json" -d '{"key":"value"}'` : Send a POST request with JSON data.

```bash
curl -X POST https://api.boot.dev/v1/courses_rest_api/learn-http/users -H "Content-Type: application/json" -d '{
  "role": "QA Job Safety",
  "experience": 2,
  "remote": true,
  "user": {
    "name": "Dan",
    "location": "NOR",
    "age": 29
  }
}' > /tmp/user.json
```

```bash
curl -X POST localhost:8080/api/validate_chirp -H "Content-Type: application/json" -d '{
"body": "this is my first chirp!",
}'
```
