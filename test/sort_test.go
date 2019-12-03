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
	"testing"
)

func TestMergeSort(t *testing.T) {
	s := []int{14, 5, 3, 23, 6, 8, 4, 8}
	l := rdata.MergeSort(s)
	fmt.Println(l)
}
