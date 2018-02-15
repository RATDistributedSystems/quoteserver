package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/textproto"
	"strconv"
	"strings"
	"time"

	"github.com/RATDistributedSystems/utilities"
)

var confs = utilities.GetConfigurationFile("config.json")

func main() {
	addr, protocol := confs.GetServerDetails("quote")
	l, err := net.Listen(protocol, addr)
	if err != nil {
		log.Fatalln("Error listening:", err.Error())
	}
	defer l.Close()
	log.Printf("Simulating Quoteserver on %s...", addr)

	// Handle requests till infinity
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Error accepting connections: %s", err.Error())
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	defer conn.Close()
	tp := textproto.NewReader(bufio.NewReader(conn))
	msg, _ := tp.ReadLine()
	log.Printf("Recieved request: %s", msg)
	result := strings.Split(string(msg), ",")
	//if not enough arguments, or incorrect format
	//send NA and close connection
	if len(result) != 2 || len(result[0]) > 3 {
		log.Println("Invalid request")
		conn.Write([]byte("NA"))
		return
	}

	username := strings.TrimSpace(result[1])
	stock_sym := strings.TrimSpace(result[0])
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand_price := r.Float64() * float64(rand.Intn(1000))
	stock_price := strconv.FormatFloat(rand_price, 'f', 2, 64)
	t := time.Now().UTC().UnixNano()
	time := strconv.Itoa(int(t))
	crypto := "7777777777"
	response := fmt.Sprintf("%s,%s,%s,%s,%s", stock_price, stock_sym, username, time, crypto)
	log.Printf("Response: %s\n", response)
	fmt.Fprintf(conn, response)
}
