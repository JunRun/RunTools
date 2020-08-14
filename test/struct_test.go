/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-13 14:52
 */
package test

import (
	"encoding/json"
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

func TestJson(t *testing.T) {
	//st:="{\"LOTUS_MINER_PATH=/mnt/ceph/10.160.0.4/miner\",\"LOTUS_WORKER_PATH=/mnt/cd_sg201/worker/worker-cd-sg201-10-160-0-43-CG-40\"}"
	var ss []string
	ss = append(ss, "LOTUS_MINER_PATH=/mnt/ceph/10.160.0.4/miner")
	ss = append(ss, "LOTUS_WORKER_PATH=/mnt/cd_sg201/worker/worker-cd-sg201-10-160-0-")
	var tt []byte
	if sd, err := json.Marshal(ss); err != nil {
		fmt.Println(err)
	} else {
		tt = sd
		fmt.Println(string(sd))
	}
	var sp []string
	err := json.Unmarshal(tt, &sp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sp[1])
	}

}
