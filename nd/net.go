package nd

import (
	"fmt"
	"log"
	"net"
	"net-debugger/util"
	"strconv"
)

type Net interface {
	Connect() net.Conn
	Listen() net.Conn
}

type tcpConnector struct {
	Host string
	Port string
}

func (c tcpConnector) addr() string {
	return c.Host + ":" + c.Port
}

func (c tcpConnector) Listen() net.Conn {
	address := c.addr()
	log.Printf("try listen on %s", address)
	listen, err := net.Listen("tcp", address)
	util.CheckFatalError(err, "failed to listen", address)
	log.Printf("wait for client...")
	conn, err := listen.Accept()
	util.CheckFatalError(err, "failed to accept connection")
	log.Printf("client %s connected", conn.RemoteAddr().String())
	return conn
}

func (c tcpConnector) Connect() net.Conn {
	address := c.addr()
	log.Printf("try to connect tcp->%s", address)
	dial, err := net.Dial("tcp", address)
	util.CheckFatalError(err, "failed to connect", address)
	log.Printf("connected, remote addres: %s", dial.RemoteAddr())
	return dial
}

func Tcp(host string, port string) Net {
	return tcpConnector{Host: host, Port: port}
}

type udpConnector struct {
	Host string
	Port int
}

func Udp(host string, port string) *udpConnector {
	p, err := strconv.Atoi(port)
	util.CheckFatalError(err, "address must be port number when protocol is udp")
	return &udpConnector{Host: host, Port: p}
}

func (u udpConnector) Connect() net.Conn {
	address := fmt.Sprintf("%s:%v", u.Host, u.Port)
	log.Printf("try to connect udp->%s", address)
	dial, err := net.Dial("udp", address)
	util.CheckFatalError(err, "failed to connect", address)
	log.Printf("connected, remote addres: %s", dial.RemoteAddr())
	return dial
}

func (u udpConnector) Listen() net.Conn {
	log.Printf("try listen on... udp->0.0.0.0:%v", u.Port)
	udp, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: u.Port})
	util.CheckFatalError(err, "failed to listen")
	log.Printf("wait for message...")
	return udp
}
