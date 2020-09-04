/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020-03-17 13:12
 */
package test

import (
	"bufio"
	"context"
	"fmt"
	"github.com/JunRun/RunTools/rfile"
	"os"
	"testing"
	"time"
)

func TestFile(t *testing.T) {

	rfile.FileRead("/Volumes/videos/crunchyroll_video", "")
	fmt.Println(len(rfile.VideoList))
	for _, v := range rfile.VideoList {
		fmt.Println(v.Name, v.Url)
	}

}

func TestTimeOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
	defer cancel()
	go func() {
		time.Sleep(time.Second * 1)

	}()
	select {
	case <-ctx.Done():
		fmt.Println("call successfully")
	case <-time.After(time.Duration(time.Second * 2)):
		fmt.Println("2")
	}
}
func TestCha(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println("接收值", <-ch)
	close(ch)

}

//测试文件游标
func TestFileSeek(t *testing.T) {
	f, err := os.Open("C:\\Users\\Admin\\Desktop\\sectors(2).log")
	fi, err := os.Stat("C:\\Users\\Admin\\Desktop\\sectors(2).log")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.Seek(fi.Size()-320, 1)
	s := bufio.NewReader(f)
	by, err := s.ReadString('\n')
	if err != nil {

		fmt.Println(err)
		return
	}
	fmt.Println(by)

}
