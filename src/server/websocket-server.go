package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    
    defer conn.Close()

    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("read: ", err)
            break
        }
        log.Print("recv: %s", message)

        // Echo the message back to the client
        if err := conn.WriteMessage(messageType, message); err != nil {
            log.Println("write: ", err)
            break
        }
    }
}

func main() {
    http.HandleFunc("/ws", wsHandler)
    fmt.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}


