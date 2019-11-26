package logAccess

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Reader interface {
	Read(data chan []byte)
}
type Writer interface {
	Write(data chan string)
}
type LogProcess struct {
	Rc    chan []byte
	Wc    chan string
	Read  Reader
	Write Writer
}

type ReadModel struct {
	Path string
}
type WriteModel struct {
	Db string
}

//读取模块
func (r *ReadModel) Read(data chan []byte) {
	f, err := os.Open(r.Path)
	if err != nil {
		log.Fatalln("File open failed.", err)
	}
	//移动到文件末尾一行
	_, _ = f.Seek(0, 2)
	reader := bufio.NewReader(f)
	for {
		msg, er := reader.ReadBytes('\n')
		if er == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if er != nil {
			fmt.Println("false")
			panic(err)
		}
		data <- msg
	}
}

//写入模块
func (w WriteModel) Write(data chan string) {
	for ch := range data {
		log.Print(ch)
	}
}

//分析模块
func (l *LogProcess) LogWrite() {
	for v := range l.Rc {
		l.Wc <- strings.ToUpper(string(v))
	}
}
func (l *LogProcess) WriteInflux() {

}
