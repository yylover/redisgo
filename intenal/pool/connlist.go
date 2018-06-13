package pool

type ConnNode struct {
	conn      Conn
	pre, next *ConnNode
}

type ConnList struct {
	head, tail *ConnNode
	len        uint32
}

func NewConnList() *ConnList {
	return &ConnList{}
}

func (this *ConnList) PushTail(conn Conn) {
	node := &ConnNode{
		conn: conn,
	}
	if this.tail == nil {
		this.head = node
		this.tail = node
	} else {
		node.pre = this.tail
		this.tail.next = node
		this.tail = node
	}
	this.len++
}

func (this *ConnList) PopTail() *Conn {
	if this.tail == nil {
		return nil
	} else {
		node := this.tail
		this.tail = this.tail.pre
		if this.tail != nil {
			this.tail.next = nil
		}
		this.len--
		return &node.conn
	}
}

func (this *ConnList) PushHead(conn Conn) {
	node := &ConnNode{
		conn: conn,
	}
	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		node.next = this.head
		this.head.pre = node
		this.head = node
	}
	this.len++
}

func (this *ConnList) PopHead() *Conn {
	if this.head == nil {
		return nil
	} else {
		node := this.head
		this.head = this.head.next
		if this.head != nil {
			this.head.next = nil
		}
		this.len--
		return &node.conn
	}
}

func (this *ConnList) IsEmpty() bool {
	return this.len == 0
}
