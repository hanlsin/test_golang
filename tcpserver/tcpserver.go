package tcpserver

import (
	"fmt"
	"io"
	"net"
	"time"
)

const (
	// DefaultHost default IP
	DefaultHost = "localhost"
	// DefaultPort default port
	DefaultPort = 4567
)

// TCPServer struct for TCP server
type TCPServer struct {
	ip          string
	port        int
	lstnr       *net.TCPListener
	stopSrvrSig chan bool
}

// Session accepted client sesson from connection
type Session struct {
	srvr *TCPServer
	conn net.Conn
	num  int
}

// NewTCPServer create TCPServer struct
func NewTCPServer() *TCPServer {
	return &TCPServer{
		ip:   string(DefaultHost),
		port: int(DefaultPort),
	}
}

// Start start TCP server
func (ts *TCPServer) Start() error {
	lstnr, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(ts.ip), Port: ts.port})
	if err != nil {
		fmt.Println("ERROR: Listen:", err)
		return err
	}
	defer lstnr.Close()

	ts.lstnr = lstnr
	ts.stopSrvrSig = make(chan bool, 1)
	defer close(ts.stopSrvrSig)

	sessionNum := 1
	for {
		stopServer := false
		select {
		case v := <-ts.stopSrvrSig:
			if v {
				fmt.Println("Got Server Stop Signal")
				stopServer = true
			}
		default:
		}
		if stopServer {
			break
		}

		nowTime := time.Now()
		fmt.Println("Start Accept:")
		fmt.Println(nowTime)

		// wake up everty 10 seconds
		err := lstnr.SetDeadline(nowTime.Add(10 * time.Second))
		if err != nil {
			fmt.Println("ERROR: Set timeout:", err)
			return err
		}

		conn, err := lstnr.Accept()
		if err != nil {
			fmt.Println("ERROR: Accept:", err)
			fmt.Println(time.Now())
			continue
		}

		session := &Session{
			srvr: ts,
			conn: conn,
			num:  sessionNum,
		}
		sessionNum++

		go session.handleRequest()
	}

	fmt.Println("Server Stopped!!")

	return nil
}

func (s *Session) handleRequest() {
	fmt.Println("Begin Session ", s.num)
	buf := make([]byte, 1024)
	for {
		len, err := s.conn.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("ERROR: Read:", err)
			break
		}

		fmt.Println("Request len =", len)
		if len == 0 {
			if err == io.EOF {
				break
			}
			continue
		}

		str := string(buf[:len])
		//str = strings.TrimSpace(str)
		fmt.Println("Msg:", str)

		if str == "STOP SERVER" {
			fmt.Println("Trying to Stop Server from", s.num)
			s.srvr.stopSrvrSig <- true
			break
		} else if str == "QUIT" || err == io.EOF {
			break
		}
	}
	s.conn.Close()
	fmt.Println("End Session ", s.num)
}
