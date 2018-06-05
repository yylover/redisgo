package proto

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
)

const (
	PROTO_PREFIX_STATUS  = rune('+')
	PROTO_PREFIX_ERROR   = rune('-')
	PROTO_PREFIX_INTEGER = rune(':')
	PROTO_PREFIX_STRING  = rune('$')
	PROTO_PREFIX_Array   = rune('*')
)

var (
	NilError = errors.New("redis: nil return")
)

type ReplyParser struct {
	reader *bufio.Reader
}

func NewReplyParser(reader *bufio.Reader) *ReplyParser {
	return &ReplyParser{
		reader: reader,
	}
}

func (this *ReplyParser) ParseReply() error {
	r, size, err := this.reader.ReadRune()
	if err != nil {
		return err
	}

	switch r {
	case PROTO_PREFIX_STATUS:
		err = this.parseStatusReply()
		if err != nil {
			return err
		}
	case PROTO_PREFIX_ERROR:
		err = this.parseErrorReply()
		if err != nil {
			return err
		}
	case PROTO_PREFIX_INTEGER:
		err = this.parseIntegerReply()
		if err != nil {
			return err
		}
	case PROTO_PREFIX_STRING:
	case PROTO_PREFIX_Array:
	}

	fmt.Println(string(r))
	fmt.Println(size)

	return nil
}
func (this *ReplyParser) parseStatusReply() error {
	status, isprefix, err := this.reader.ReadLine()
	if err != nil {
		return err
	}

	fmt.Println("status   :", string(status))
	fmt.Println("isprefix :", isprefix)

	return nil
}

func (this *ReplyParser) parseArrayReply() error {
	return nil
}

func (this *ReplyParser) parseErrorReply() error {
	status, isprefix, err := this.reader.ReadLine()
	if err != nil {
		return err
	}
	fmt.Println("status   :", string(status))
	fmt.Println("isprefix :", isprefix)
	return nil
}

func (this *ReplyParser) parseStringReply() error {
	status, isprefix, err := this.reader.ReadLine()
	if err != nil {
		return err
	}
	fmt.Println("status   :", string(status))
	fmt.Println("isprefix :", isprefix)
	return nil
}
func (this *ReplyParser) parseIntegerReply() error {
	status, isprefix, err := this.reader.ReadLine()
	if err != nil {
		return err
	}

	iret, err := strconv.ParseInt(string(status), 10, 32)
	if err != nil {
		return err
	}
	fmt.Println("status   :", iret)
	fmt.Println("isprefix :", isprefix)
	return nil
}

func (this *ReplyParser) isNilReply() error {
	return nil
}
