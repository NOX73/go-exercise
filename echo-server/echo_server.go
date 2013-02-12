package main

import(
  "fmt"
  "net"
  //"strings"
  "io"
)

type Message struct{
  Text string
}

func main() {
  
  psock, _ := net.Listen("tcp", "127.0.0.1:4000")
  conn, _ :=  psock.Accept()

  messages_channel := make(chan Message, 5)

  go func() {
    for{
      select {
      case msg := <-messages_channel:
        fmt.Println("Recieve to channel: ", msg)
        conn.Write([]byte("Echo: " + msg.Text + "\n\r"))
      }
    }
  }();

  buffer := make([]byte, 24)
  var message Message;
  for {
    l,e := conn.Read(buffer)
    fmt.Println("\tReaded bytes: ", l)
    fmt.Println("\tError: ", e)
    switch {
    case e == nil:
      str := string(buffer)
      fmt.Println("\tMessage recieve: ", str)
      message = Message{str}
      messages_channel <- message
    case e == io.EOF:
      fmt.Println("EOF")
    }
  
  }
  return
  
}
