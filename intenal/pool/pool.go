package pool

import (
	"sync"
	"time"
)

//get redis client pool
type Pool struct {
	//config
	MaxIdle         uint32        //最大空闲
	MaxActive       uint32        //最大活跃连接个数
	IdleTimeout     time.Duration //空闲时间
	MaxConnLifeTime time.Duration //连接最大使用时间
	TestConnOnGet   func(c Conn, t time.Time) error
	Dial            func() (Conn, error)

	active   uint32 //当前活跃连接数
	idle     uint32
	mu       sync.RWMutex
	connlist ConnList
	//idle list
}

func NewPool() *Pool {
	return &Pool{}
}

// get a new Conn
func (p *Pool) Get() *Conn {
	if p.connlist.IsEmpty() {

	}

	conn := p.connlist.PopHead()
	return conn
}
