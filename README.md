#  Go API with Gin

---

##  Installation Steps

### Step 1: Install Required Packages

#### Requirements

1. **Install `sudo` (if not installed):**
   ```bash
   apt-get update
   apt-get install sudo
   ```

2. **Install Go:**
   ```bash
   sudo apt install golang-go
   ```

---

##  Build and Run the API Locally

### Step 2: Project Folder Structure

I maintained the following folder structure:

```
go_project/
│
├── build.sh
├── go.mod
├── go.sum
├── bin/
│   └── (compiled binary after build)
└── src/
    └── main.go
```

---

### Step 3: `main.go` Content

In `main.go` I wrote the following code:

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {ID: "1", Title: "Book One", Author: "Author One"},
    {ID: "2", Title: "Book Two", Author: "Author Two"},
}

func getBooks(c *gin.Context) {
    c.JSON(http.StatusOK, books)
}

func main() {
    r := gin.Default()
    r.GET("/books", getBooks)
    r.Run(":8081") // Run server on port 8081
}
```

---

### Step 4: Build Script

`build.sh`:
```bash
#!/bin/bash

set -e

echo " Creating bin directory if it doesn't exist..."
mkdir -p bin

echo " Building the Go application from src/main.go..."
go build -o bin/bookapi ./src/main.go

echo " Build completed! Executable is at bin/bookapi"
```

Make it executable and run:
```bash
chmod +x build.sh
./build.sh
```

---

### Step 5: Initialize Go Module

```bash
go mod init bookapi
```

---

##  Troubleshooting: SSL Certificate Errors

I encountered the following error when fetching dependencies:

```
go: module github.com/gin-gonic/gin: Get "https://proxy.golang.org/github.com/gin-gonic/gin/@v/list": tls: failed to verify certificate: x509: certificate signed by unknown authority
```

### Cause:
- System lacks CA certificates
- Proxy or network blocking access

### Fix:
```bash
sudo apt update
sudo apt install ca-certificates
sudo update-ca-certificates
go get -u github.com/gin-gonic/gin
```

---

##  Running the Application

### Step 6: Run the Application

```bash
./bin/bookapi
```

The API runs at:  
`http://localhost:8081/books`

---

### Step 7: Run in Docker

When I run the app inside Docker, I map the port like this:

```bash
docker run -p 8081:8081 bookapi
```

Now I can access the API at:  
`http://localhost:8081/books`

---

##  Testing the API

### Using Browser:
```
http://localhost:8081/books
```

###  Using Terminal:
```bash
curl http://localhost:8081/books
```

---

##  Common Issues & Fixes

###  `localhost refused to connect`

Since I ran the app inside Docker, I needed to map the port properly:

```bash
docker run -p 8081:8081 bookapi
```

---
