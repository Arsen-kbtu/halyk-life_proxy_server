# Proxy Server API

This project is a Proxy Server API built with Go. It provides a simple proxy server with defined endpoints and handlers.

### Installation

1. Clone the repository:
git clone https://github.com/Arsen-kbtu/halyk-life_proxy_server.git
2. ```bash
   cd halyk-life_proxy_server
   ```
3. Build and run the application using Makefile:
4. ```bash
   make build
   ```
   ```bash
   make up
   ```

5. Once the application is running, you can access it at http://localhost:8080

### API Documentation
The API documentation can be found at http://localhost:8080/swagger/index.html once the application is running. It provides detailed information about the available endpoints and how to use them.

#### Sending request 
```http
  POST /proxy
```
#### Check history 
```http
  GET /history
```

#### Health check
```http
  GET /health
```


You can also access this project on https://halyk-life-proxy-server.onrender.com
