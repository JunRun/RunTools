/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-13 14:52
 */
package test

import (
	"fmt"
	"testing"
)

type Option struct {
	num  int
	name string
}

type ModOption func(option *Option)

func TestAd(t *testing.T) {
	const mask = 1<<32 - 1
	var i uint32
	var s uint64
	s = 3132313123113
	i = uint32((s >> 32) & mask)
	fmt.Printf("%32b,\n %64b", i, s)
}
