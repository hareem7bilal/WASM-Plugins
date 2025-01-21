# WASM Plugins

This repository contains three simple WebAssembly (WASM) plugins that can be executed using Extism. Each plugin demonstrates different functionality:

- **httpget**: Performs an HTTP GET request.
- **api**: Makes a customizable HTTP request.
- **greet**: Greets a user with a given name.

---

## Prerequisites

- [Extism CLI](https://extism.io/) installed.
- [Go programming language](https://go.dev/) installed (for running Go SDK examples).

---

## Running the Plugins with Extism CLI

### 1. **httpget Plugin**
This plugin performs a GET request to a specified URL.

```bash
extism call --wasi --allow-host jsonplaceholder.typicode.com plugin.wasm httpGet
```

### 2. **api Plugin**
This plugin allows you to make customizable HTTP requests.

```bash
extism call plugin.wasm makeRequest --wasi --allow-host="*.typicode.com" --input '{"url": "https://jsonplaceholder.typicode.com/posts", "method": "POST", "headers": {"Content-Type": "application/json"}, "body": "{\"title\": \"foo\", \"body\": \"bar\", \"userId\": 1}"}'
```

### 3. **greet Plugin**
This plugin greets a user with a given name.

```bash
extism call plugin.wasm greet --input "Hareem" --wasi
```

---

## Testing with Go SDK
Each plugin folder (`api`, `greet`, `httpget`) contains a `main.go` file that demonstrates how to use the plugin with the Extism Go SDK.

### Steps:
1. Navigate to the plugin directory:
   ```bash
   cd api
   ```
2. Run the Go program:
   ```bash
   go run main.go
   ```
   This will compile the plugin to WASM and execute it using the Go SDK.

---

## Plugin Details

### **httpget**
- **Location**: `/httpget`
- **Function**: Performs a GET request to `https://jsonplaceholder.typicode.com/todos/1`.
- **Output**: JSON response from the API.

### **api**
- **Location**: `/api`
- **Function**: Makes a customizable HTTP request.
- **Input**: JSON object with `url`, `method`, `headers`, and `body`.
- **Output**: Response from the API.

### **greet**
- **Location**: `/greet`
- **Function**: Greets a user.
- **Input**: User's name.
- **Output**: Greeting message.

---

## Example Output

### **httpget Plugin**
Response:
```json
{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
```

### **api Plugin**
Response:
```json
{
  "id": 101,
  "title": "foo",
  "body": "bar",
  "userId": 1
}
```

### **greet Plugin**
Response:
```
Hello, Hareem!
```

