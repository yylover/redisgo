// Copyright 2018 quickyang
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

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
	case PROTO_PREFIX_ERROR:
		err = this.parseErrorReply()
	case PROTO_PREFIX_INTEGER:
		err = this.parseIntegerReply()
	case PROTO_PREFIX_STRING:
	    err = this.parseStringReply()
	case PROTO_PREFIX_Array:
	    err = this.parseArrayReply()
	}
	if err != nil {
	    return err
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
