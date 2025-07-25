````markdown
# socket-lab 🔌🧪

A lab project to practice and understand different socket types in Go:

- ✅ TCP (client/server)
- ✅ UDP (connectionless communication)
- ✅ UNIX domain sockets
- ✅ HTTP over TCP
- ✅ TLS-secured sockets
- 🔄 WebSocket (WS/WSS)

## Structure

Each folder contains one protocol's server and client. Run both to observe behavior.

---

## Run TCP (VM Docker + macOS Client)

### 🐳 Server (in OrbStack VM using Docker)

```bash
cd /mac/Users/YOUR_NAME/projects/socket-lab
docker run --rm -it \
  -v $(pwd):/app \
  -w /app/server/tcp \
  -p 9000:9000 \  
  golang:1.22-alpine \
  go run tcp_server.go
````

> Server must use `0.0.0.0:9000`

### 💻 Client (on macOS host)

1. Find VM IP:

   ```bash
   ip a | grep inet
   ```
2. Update client:

   ```go
   net.Dial("tcp", "192.168.64.X:9000")
   ```
3. Run:

   ```bash
   go run client/tcp/tcp_client.go
   ```

✅ Message exchange should succeed.

---

Instructions for UDP, Unix, TLS, and WebSocket coming soon.

```
