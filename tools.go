package main

import (
	"bytes"
	"strconv"
)

func reload_string_into_byte(data string) []byte {
	var buf bytes.Buffer
	for a := 0; a < len(data)/2; a++ {
		sep := Substr(data, a*2, 2)
		sep_int, _ := strconv.ParseInt(sep, 16, 16)
		buf.WriteByte(byte(sep_int))
	}
	return buf.Bytes()

}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func crc_16(data []byte) []byte {
	var buf bytes.Buffer
	i := 0
	j := 0
	var crc uint16 = 0xffff
	for i = 0; i < len(data); i++ {
		crc = crc ^ uint16(data[i])
		for j = 0; j < 8; j++ {
			if crc&1 == 1 {
				crc = (crc >> 1) ^ 0x8408
			} else {

				crc = (crc >> 1)
			}
		}
	}
	crc2 := crc >> 8
	last := crc << 8
	crc1 := last >> 8
	buf.WriteByte(byte(crc1))
	buf.WriteByte(byte(crc2))
	crc_byte := buf.Bytes()
	return crc_byte

}
