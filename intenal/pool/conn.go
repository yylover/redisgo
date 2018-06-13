package pool

import "net"

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
