package main

import(
  "fmt"
  "net"
  "bufio"
  "os"
  "flag"
)

var ifiFlag string
var ifi *net.Interface

func init() {
  flag.StringVar(&ifiFlag, "ifi", "eth0", "Interfase")
}

func main() {
  flag.Parse()

  //ifi, _ = net.InterfaceByName(ifiFlag)
 
  //addrs, err := ifi.Addrs()
  //if err!=nil { fmt.Println("Ifi Addrs Error: ", err); os.Exit(1)}

  //addr, err := net.ResolveUDPAddr("udp4", addrs[0].String() + ":4040") 
  addr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:4040") 
  if err!=nil { fmt.Println("Resolve Error: ", err); os.Exit(1)}
   
  console_ch := make(chan string, 1)

  connection, err := net.ListenUDP("udp4", addr)
  if err!=nil { fmt.Println("Listen Error: ", err); os.Exit(1)}

  go ChatReader(connection)
  go ChatWriter(console_ch, connection)
  ConsoleReader(console_ch)

}

func ConsoleReader(channel chan string){
  fmt.Println("Enter command")
  reader := bufio.NewReader(os.Stdin)
  for {
    fmt.Print("> ")
    line, _ := reader.ReadString('\n')
    channel <- line
  }
}

func ChatReader(conn *net.UDPConn){
  reader := bufio.NewReader(conn)
  for {
    line, _ := reader.ReadString('\n')
    fmt.Print("Incomming message: ", line)
    fmt.Print("> ")
  }

}

func ChatWriter(channel chan string, conn *net.UDPConn){

    //addrs, _ := ifi.MulticastAddrs()

    addr, err := net.ResolveUDPAddr("udp4", "192.168.1.255:4040") 
    if err!=nil { fmt.Println("Resolve Error: ", err); os.Exit(1)}

    for{
      select {
      case line := <-channel:
        
        _, err := conn.WriteTo([]byte(line), addr)
        if err!=nil { fmt.Println("Write error: ", err); os.Exit(1)}

      } 
    }
}
