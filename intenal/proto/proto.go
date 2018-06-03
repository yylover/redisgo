package proto

import "errors"

const (
	PROTO_PREFIX_STATUS  = "+"
	PROTO_PREFIX_ERROR   = "-"
	PROTO_PREFIX_INTEGER = ":"
	PROTO_PREFIX_STRING  = "$"
	PROTO_PREFIX_Array   = "*"
)

var (
	NilError = errors.New("redis: nil return")
)


//const NilError = errors.New("redis: nil return")


func ParseReply() {

}

func parseStatusReply() {

}

func parseArrayReply() {

}

func isNilReply() {

}


