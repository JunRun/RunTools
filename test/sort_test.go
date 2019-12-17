/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-03 11:42
 */
package test

import (
	"fmt"
	"github.com/JunRun/RunTools/rdata"
	"github.com/JunRun/RunTools/rgo"
	"github.com/emacsist/go-common/helper/number"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	data := number.GenerateInt(100000, 100000)
	l := rdata.MergeSort(data)
	fmt.Println(l)
}

//func BenchmarkMergeSort(b *testing.B){
//	data := number.GenerateInt(100000, 100000)
//	l := rdata.MergeSort(data)
//	fmt.Println(l)
//}

//func BenchmarkTes(b *testing.B){
//	fmt.Println("sdsdsdsd")
//}

func TestQuickSort(t *testing.T) {
	data := number.GenerateInt(100000, 100000)
	//data:=[]int{2,8,4,7,12,34,2,45,12,5,67,34}
	rdata.QuickSort(data)
	fmt.Println(data)
}

func BenchmarkTestQuickSort(b *testing.B) {
	data := number.GenerateInt(100000, 100000)
	//data:=[]int{2,8,4,7,12,34,2,45,12,5,67,34}
	rdata.QuickSort(data)
	fmt.Println(data)
}

func TestTime(t *testing.T) {

	rgo.Group.Add(1)
	rgo.Timer()

	rgo.Group.Wait()
}

func Hello() {
	time.Sleep(time.Second)
	fmt.Println("hello")
}
