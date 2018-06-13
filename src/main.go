package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/redisgo/intenal/proto"
	"log"
	"net"
	"strings"
	"time"
)

const (
    //loop address
	address = "127.0.0.1"
	port    = "6379"
)

type ReConn struct {
	conn net.TCPConn
}

type DialOptions struct {
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ReadBufferSize  int           //采用系统默认
	WriteBufferSize int           //采用系统默认
	KeepAlive       bool          //是否设置keepalive
	KeepAlivePeriod time.Duration //保存时间
	NoDelay         bool          //是否开启nodelay算法ß
}

//create tcp connection

//exec command

//parse The Output

//
func main() {
	raddr, err := net.ResolveTCPAddr("tcp4", address+":"+port)
	if err != nil {
		log.Print(err)
	}

	tcpconn, err := net.DialTCP("tcp4", nil, raddr)
	if err != nil {
		log.Print(err)
	}

	defer tcpconn.Close()

	//tcpconn.SetKeepAlive(true)
	//tcpconn.SetKeepAlivePeriod(time.Second)
	//tcpconn.SetReadDeadline(time.Now().Add(time.Second))
	//tcpconn.SetWriteDeadline(time.Now().Add(time.Second))
	//tcpconn.SetReadBuffer(128)
	//tcpconn.SetWriteBuffer(128)
	//tcpconn.SetNoDelay(true)

	br := bufio.NewReader(tcpconn)
	bw := bufio.NewWriter(tcpconn)

	//command := getCommand("mget", "test", "test2", "test3")
	//command := getCommand("set", "test", "2")
	//command := getCommand("get", "htest")
	command := getCommand("incr", "test")
	//len, err := tcpconn.Write([]byte(command))
	fmt.Println("command : ", command)
	len, err := bw.WriteString(command)
	if err != nil {
		log.Print("write err :", err.Error())
	}
	log.Println("write len :", len, err, command)
	bw.Flush()

	//buf := make([]byte, 1024)
	//rlen, err := br.Read(buf)
	parser := proto.NewReplyParser(br)
	err = parser.ParseReply()
	if err != nil {
		log.Print("read failed : ", err.Error())
	}

	//parseReplay(buf)

	//log.Print("read result : ", rlen)
	//log.Print("read result : ", string(buf))
}

func getCommand(com string, params ...string) string {
	buffer := bytes.NewBufferString(com)
	if len(params) > 0 {
		subcoms := strings.Join(params, " ")
		buffer.WriteString(" " + subcoms)
	}

	buffer.WriteString("\r\n")

	return buffer.String()
}

func parseReplay(buf []byte) interface{} {
	reader := bytes.NewReader(buf)
	reader.ReadRune()

	return nil
}
