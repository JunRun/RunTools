/**
 *
 * @Description: 读取文件
 * @Version: 1.0.0
 * @Date: 2020-03-17 13:06
 */
package rfile

import (
	"io/ioutil"
	"sync"
)

var fileDir chan int

type Video struct {
	Name string
	Url  string
}

var VideoList []Video
var sys sync.WaitGroup

func FileRead(file string, name string) {

	//fileDir:=make(chan int,20)
	files, _ := ioutil.ReadDir(file + "/" + name)
	for _, f := range files {
		if f.IsDir() {
			FileRead(file, f.Name())
		} else {
			VideoList = append(VideoList, Video{
				Name: name,
				Url:  f.Name(),
			})
		}
	}

}
