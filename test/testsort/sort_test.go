/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-03 11:42
 */
package testsort

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"testing"
	"time"

	"github.com/emacsist/go-common/helper/number"

	"github.com/JunRun/RunTools/rdata"
)

func TestMergeSort(t *testing.T) {
	data := number.GenerateInt(100000, 100000)
	_ = rdata.MergeSort(data)
}

func BenchmarkMergeSort(b *testing.B) {
	data := number.GenerateInt(100000, 100000)
	_ = rdata.MergeSort(data)
}

//func BenchmarkTes(b *testing.B){
//	fmt.Println("sdsdsdsd")
//}

func TestQuickSort(t *testing.T) {
	cpu_file := "cpu_file.prof"
	f, err := os.Create(cpu_file)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	data := number.GenerateInt(100000, 100000)
	//data:=[]int{2,8,4,7,12,34,2,45,12,5,67,34}
	rdata.QuickSort(data)
}

func BenchmarkTestQuickSort(b *testing.B) {
	data := number.GenerateInt(100000, 100000)
	//data:=[]int{2,8,4,7,12,34,2,45,12,5,67,34}
	rdata.QuickSort(data)
}

func Hello() {
	time.Sleep(time.Second)
	fmt.Println("hello")
}

func TestABF(t *testing.T) {
	fmt.Println(foo(), bar())
	fmt.Println(foo(), bar())
}

func fo(p *int) int {
	*p = 123
	return *p
}

func foo() int {
	var x int
	y, _ := x, fo(&x)
	return y
}

func bar() int {
	var x int
	var y, _ = x, fo(&x)
	return y
}
