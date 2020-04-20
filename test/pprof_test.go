/**
 *
 * @Description: pprof 学习
 * @Version: 1.0.0
 * @Date: 2020/4/20 10:21 上午
 */
package test

import (
	"encoding/json"
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestArr(t *testing.T) {

	mapData := map[string]string{
		"ss1": "123131233123",
		"ss2": "123131233123",
		"ss3": "123131233123",
		"ss4": "123131233123",
		"ss5": "123131233123",
	}
	cpu_file := "cpu_file.prof"
	f, err := os.Create(cpu_file)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	for i := 0; i < 10000000; i++ {
		_, _ = json.Marshal(mapData)
	}

}
