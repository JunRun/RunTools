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
	"io"
	"log"
	"os"
	"strings"
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

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.Seek(0, 2)
	s := bufio.NewReader(f)
	by, err := s.ReadString('\n')
	if err != nil {

		fmt.Println(err)
		return
	}
	fmt.Println(by)
}

func TestLogFile(t *testing.T) {

	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	i := 0
	go func() {
		for {
			time.Sleep(time.Second * 1)
			log.Println("tsd", i)
			i++
		}
	}()

	go listenFile("D:\\Users\\Admin\\go\\RunTools\\test\\log.txt")
	select {}
}

func listenFile(path string) {

}

func TestLPL(t *testing.T) {
	var (
		s   = "D:\\test-1"
		err error
	)
	//os.Mkdir(s,0666)
	for i := 1; i <= 20000; i++ {
		path := fmt.Sprintf("%s/%d", s, i)
		err = os.Mkdir(path, 0777)
		if err != nil {
			panic(err)
		}
		tmp := fmt.Sprintf("%s/%d", path, i)
		file, _ := os.Create(tmp)
		_, err = io.Copy(file, strings.NewReader(fmt.Sprintf("%s", path)))
	}

}

func TestMk(t *testing.T) {
	f := func() { fmt.Print("A") }
	defer f()
	f = func() { fmt.Print("B") }
	defer f()
}

type integer int

func (i integer) String() string {
	return "hello"
}

func TestM(t *testing.T) {
	fmt.Println(integer(5))
}
