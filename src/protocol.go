package main

import (
	"bytes"
	"encoding/binary"
)
/**
	#lf#000001{"cmd":"connect","uuid":"123456"}
	#lf#000001{"cmd":"msg","uuid":"123456","content":"消息内容","receiveUuid":"888888"}
	#lf#000001{"cmd":"disconnect","uuid":"123456"}
	#lf#000001{"cmd":"heart","uuid":"123456"}
 */
const (
	Header = "#lf#"
	HeaderLength = 4
	DataLength = 4
)


func unPack(buffer []byte,readerChannel chan []byte)[]byte{
	length := len(buffer)
	var i int
	for i = 0;i < length;i++{
		if length < i + HeaderLength + DataLength{
			break
		}
		if string(buffer[i:i+HeaderLength]) == Header{
			messageLength := BytesToInt(buffer[i+HeaderLength : i+HeaderLength+DataLength])
			if length < i+HeaderLength+DataLength+messageLength {
				break
			}
			data := buffer[i+HeaderLength+DataLength : i+HeaderLength+DataLength+messageLength]
			readerChannel <- data

			i += HeaderLength + DataLength + messageLength - 1
		}
	}
	if i == length{
		return make([]byte,0)
	}
	return buffer[i:]
}


//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}