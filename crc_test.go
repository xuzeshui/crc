/*
 * Copyrignt (c) dingdong.top. All Rights Reserved.
 * Author: xuzeshui@dingdong.top
 * Created Time: 2017-05-17 16:43:22
 * Last Modified: 2017-07-17 19:54:31
 * File Name: crc/crc_test.go
 * Description:
 */

//
package crc

import (
	"encoding/hex"
	"strconv"
	"testing"
)

func Test_1(t *testing.T) {
	hexData := "aa800001010026010f313233343536373839303132333435146162636465666768696a6b6c6d6e6f70717273744521"
	data, _ := hex.DecodeString(hexData)
	h := NewHash(CCITT)
	val := h.CalculateCRC(data[:len(data)-2])
	t.Logf("%04X\n", val)

}

func Test_2(t *testing.T) {
	for i := 0; i < 100; i++ {
		v := strconv.Itoa(i)
		h := NewHash(CCITT)
		val := h.CalculateCRC([]byte(v))
		t.Logf("crc16[%d] = %04X\n", i, val)
	}
}

func Test_3(t *testing.T) {
	h := NewHash(CCITT)
	val := h.CalculateCRC([]byte("HelloWorld"))
	t.Logf("%04X\n", val)
}
