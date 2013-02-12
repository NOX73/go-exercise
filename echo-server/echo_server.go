package main

import(
  "fmt"
  "net"
  //"strings"
  //"io"
  "bufio"
)

type Message struct{
  Text string
}

func main() {
  
  psock, _ := net.Listen("tcp", "127.0.0.1:4000")
  for{
    conn, _ :=  psock.Accept()

    go func(){
      messages_channel := make(chan Message, 5)

      go func() {
        for{
          select {
          case msg := <-messages_channel:
            fmt.Println("Echo to client: ", msg.Text)
            conn.Write([]byte("Echo: " + msg.Text + "\n"))
          }
        }
      }();

      var message Message
      //var line string
      reader := bufio.NewReader(conn)

      for {
        line, _ := reader.ReadString('\n')
        
        s_line := line[:len(line)-1] 
        fmt.Println("Message recieve: ", s_line)
        message = Message{s_line}
        messages_channel <- message
      
      }
    }()
  }

}
