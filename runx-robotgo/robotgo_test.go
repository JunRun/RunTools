package runx_robotgo

import (
	"github.com/go-vgo/robotgo"
	"testing"
	"time"
)

//func TestGetWindow(t *testing.T){
//	x, y := robotgo.GetMousePos()
//	fmt.Println("pos: ", x, y)
//
//	color := robotgo.GetPixelColor(100, 200)
//	fmt.Println("color---- ", color)
//}
//
//func TestGetWindow2(t *testing.T){
//	fpid, err := robotgo.FindIds("Google")
//	if err == nil {
//		fmt.Println("pids... ", fpid)
//
//		if len(fpid) > 0 {
//			robotgo.ActivePID(fpid[0])
//
//			robotgo.Kill(fpid[0])
//		}
//	}
//
//
//	robotgo.ActiveName("Chrome")
//
//	isExist, err := robotgo.PidExists(100)
//	if err == nil && isExist {
//		fmt.Println("pid exists is", isExist)
//
//		robotgo.Kill(100)
//	}
//
//	abool := robotgo.ShowAlert("test", "robotgo")
//	if abool {
//		fmt.Println("ok@@@ ", "ok")
//	}
//
//	title := robotgo.GetTitle()
//	fmt.Println("title@@@ ", title)
//
//}

func TestKeyboard(t *testing.T) {
	robotgo.GetActive()
	robotgo.PasteStr("hello")

	robotgo.KeyTap("enter")

	robotgo.PasteStr("hello")
	time.Sleep(time.Second)
}
