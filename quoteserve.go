package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
    "strings"
    "math/rand"
    "time"
    "strconv"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  // will listen for message to process ending in newline (\n)
    print("received request")
    message, _ := bufio.NewReader(conn).ReadString('\n')
    //remove new line character and any spaces received
    message = strings.TrimSuffix(message, "\n")
    message = strings.TrimSpace(message)
    //try to split the message into tokens for processing
    result := strings.Split(string(message), ",")
    //if not enough arguments, or incorrect format
    //send NA and close connection
    if len(result) != 2 || len(result[0]) > 3 {
      conn.Write([]byte("NA"))
      conn.Close()
    //correct input
    } else {
      username := result[1]
      username = strings.TrimSpace(username)
      stock_sym := result[0]
      stock_sym = strings.TrimSpace(stock_sym)
      r := rand.New(rand.NewSource(time.Now().UnixNano()))
      //generate random price
      rand_price := r.Float64() * float64(rand.Intn(1000))
      stock_price := strconv.FormatFloat(rand_price, 'f', 2, 64)
      //get curret timestamp in UTC
      t := time.Now().UTC().UnixNano()
      time := strconv.Itoa(int(t))
      //generate cryptokey
      crypto := "7777777777"
      /*
      print(stock_price + ",")
      print(result[0] +",")
      print(result[1] + ",")
      print(time + ",")
      print(crypto)
      */
      fmt.Fprintf(conn, stock_price + "," + stock_sym + "," + username + "," + time + "," + crypto)

    }
  // Close the connection when you're done with it.
  conn.Close()
}
