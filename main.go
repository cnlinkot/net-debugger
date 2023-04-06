package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net"
	"net-debugger/nd"
	"net-debugger/util"
	"os"
	"sync"
)

var (
	host     string
	port     string
	protocol string
	body     string
	wg       = sync.WaitGroup{}
	n        nd.Net
	encoder  nd.Encoder
	conn     net.Conn
)

func main() {
	flag.StringVar(&host, "h", "127.0.0.1", "host ip")
	flag.StringVar(&port, "p", "8080", "port number")
	flag.StringVar(&protocol, "pr", "tcp", "tcp or udp")
	flag.StringVar(&body, "d", "text", "data type, text or hex")
	b := flag.Bool("s", false, "server mode")
	flag.Parse()
	println("net debug tool by sa@linkot.cn")
	initNet()
	initEncoder()
	initConn(*b)
	wg.Add(2)
	go input(conn)
	go output(conn)
	wg.Wait()
}

func initConn(b bool) {
	if b {
		conn = n.Listen()
	} else {
		conn = n.Connect()
	}
}

func initEncoder() {
	switch body {
	case "text":
		encoder = nd.PlainEncoder{}
		break
	case "hex":
		encoder = nd.HexEncoder{}
		break
	default:
		log.Fatal("no such encoder: ", body)
	}
}

func initNet() {
	switch protocol {
	case "tcp":
		n = nd.Tcp(host, port)
		break
	case "udp":
		n = nd.Udp(host, port)
		break
	default:
		log.Fatal("unknown protocol: ", protocol)
	}
}

func input(w io.Writer) {
	in := os.Stdin
	reader := bufio.NewReader(in)
	for true {
		bytes, err := reader.ReadBytes('\n')
		util.CheckError(err)
		_, err = w.Write(encoder.Decode(bytes))
		util.CheckError(err)
	}
	wg.Done()
}

func output(r io.Reader) {
	buffer := make([]byte, 4096)
	for true {
		i, err := r.Read(buffer)
		util.CheckFatalError(err, "failed to receive message")
		log.Printf("received message: %s", string(encoder.Encode(buffer[:i])))
	}
	wg.Done()
}
