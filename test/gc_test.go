/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/27 2:30 下午
 */
package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"runtime/trace"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestGc(t *testing.T) {
	go func() {
		for {

		}
	}()

	time.Sleep(time.Millisecond)
	runtime.GC()
	fmt.Println("ok")
}

// go trace 的使用 和go1.14 的异步抢占式调用
func TestTrace(t *testing.T) {
	runtime.GOMAXPROCS(1)
	create, _ := os.Create("trace.output")
	defer create.Close()
	_ = trace.Start(create)
	defer trace.Stop()
	var w sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calcSum(&w, i)
	}
	wg.Wait()
}
func calcSum(w *sync.WaitGroup, idx int) {
	defer w.Done()
	var sum, n int64
	for ; n < 1000000000; n++ {
		sum += n
	}
	fmt.Println(idx, sum)
}

func TestPutFile(t *testing.T) {
	cl := http.Client{}
	re, err := http.NewRequest("PUT", "http://222.186.132.205:7480/njnbus3test03/test", strings.NewReader("sdsads"))
	if err != nil {
		fmt.Println(err)
		return
	}
	re.Header.Set("x-amz-meta-backuptime", "1626821861")
	re.Header.Set("x-amz-meta-client", "itaotest")
	re.Header.Set("x-amz-meta-copynum", "1")
	re.Header.Set("x-amz-meta-dpaid", "NetBackup")
	re.Header.Set("x-amz-meta-dpaversion", "NBU_75")
	re.Header.Set("x-amz-meta-flags", "1")
	re.Header.Set("x-amz-meta-fragmentnum", "1")
	re.Header.Set("x-amz-meta-fulldate", "604800")
	re.Header.Set("x-amz-meta-imagetype", "2")
	re.Header.Set("x-amz-meta-masterserver", "nbunjtest")
	re.Header.Set("x-amz-meta-ostdate", "1626821861")
	re.Header.Set("x-amz-meta-policy", "NJTEST_itaotest_file_s3_bak")
	re.Header.Set("x-amz-meta-prepareforreadflag", "0")
	re.Header.Set("x-amz-meta-saveas", "132")
	re.Header.Set("x-amz-meta-sizebytes", "0")
	re.Header.Set("x-amz-meta-status", "33")
	re.Header.Set("x-amz-meta-streamnum", "1")
	re.Header.Set("x-amz-meta-version", "11")

	resp, err := cl.Do(re)
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		if string(body) == "" {
			fmt.Println(resp.Status)
		} else {
			fmt.Println(string(body))
		}
	}
}
