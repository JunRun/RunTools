package test

import (
	"fmt"
	"github.com/JunRun/RunTools/logAccess"
	"github.com/JunRun/RunTools/rlog"
	"os"
	"testing"
	"time"
)

func TestC(t *testing.T) {
	r := &logAccess.ReadModel{
		Path: "./accee.log",
	}
	w := &logAccess.WriteModel{
		Db: "ssdad",
	}

	lp := &logAccess.LogProcess{
		Rc:    make(chan []byte),
		Wc:    make(chan string),
		Read:  r,
		Write: w,
	}
	go lp.Read.Read(lp.Rc)
	go lp.LogWrite()
	go lp.Write.Write(lp.Wc)
	time.Sleep(30 * time.Second)
	fmt.Println("exit")
}

func TestS(t *testing.T) {
	log := rlog.NewRLog("", os.Stderr, rlog.BitStdFlag)
	log.Info("ss")

}

func f1(in chan int) {
	fmt.Println(<-in)
}

func TestO(t *testing.T) {
	out := make(chan int)
	out <- 2
	go f1(out)
}

func TestTime(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Dubai")
	tm := time.Now().In(location)
	startTime := tm.Format("2006-01-02T15:04")
	tm = tm.Add(1)
	s := tm.Format("2006-01-02T15:04")

	fmt.Println(startTime)
	fmt.Printf(s)
}
func AB() {
	fmt.Println("ad" + time.Now().String())
}
func TestAS(t *testing.T) {

	go func() {
		t := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-t.C:
				AB()
			}
		}
	}()
	time.Sleep(20 * time.Second)
}

//func TestListenWorker(t *testing.T){
//	worksList:=make(map[string]Works)
//	worksList["1"] =
//}
type Works struct {
	InfoList map[string]Info
}
type Info struct {
}
