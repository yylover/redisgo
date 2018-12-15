package pool

import (
	"net"
	"time"
	"bufio"
)

/**
1. tcp的连接参数
2. 命令的取消
3. 超时控制的实现原理
4.

GO:
1. context超时控制
2. 命令取消


close connection, tcp层的表现是直接发送TCP断开流程，然后直接断开，后续的所有包，直接反馈RST
*/

type Conn struct {
	iNumConsumedu int32 //使用次数

	conn net.Conn


	pending int
	err error
	readTimeout time.Duration
	br *bufio.Reader


	writeTimeout time.Duration
	bw *bufio.Writer

	lenScratch [32]byte
	numScratch [40]byte
}

func (this *Conn) Do() (interface{}, error) {

	return nil, nil
}

//Close 释放资源
func (this *Conn) Close() {
	if this.conn != nil {
		this.conn.Close()
	}

}
