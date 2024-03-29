package progress_bar

import (
	"fmt"
	"testing"
	"time"
)

func TestBar_Play(t *testing.T) {
	var bar Bar
	bar.NewOption(0, 100)
	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
}

func TestSt(t *testing.T) {
	num := 65
	str := string(num)
	fmt.Printf("%v, %T\n", str, str)
}
