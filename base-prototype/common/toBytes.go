package common

import (
	bytes2 "bytes"
	"encoding/binary"
	"log"
	"strconv"
)

func Int2Bytes(data int64) (bytes []byte) {
	buffer := new(bytes2.Buffer)
	if err := binary.Write(buffer, binary.BigEndian, data); err != nil {
		log.Panic(err)
	}

	bytes = buffer.Bytes()
	return
}

func Timestamp2Bytes(data int64) (bytes []byte) {

	// 第二个参数是 转成什么进制，如：2 8 10 16
	formatInt := strconv.FormatInt(data, 2)
	bytes = []byte(formatInt)

	return
}
