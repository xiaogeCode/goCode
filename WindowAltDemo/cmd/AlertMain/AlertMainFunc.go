package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)
import (
	//"github.com/lxn/walk"
	"syscall"
	"strconv"
	"strings"
)

//func main()  {
//	//var hwnd,winapi.hwnd
//	mainWindow.Run()
//	fmt.Println("a")
//}

func _text(str string) *uint16  {
	return syscall.StringToUTF16Ptr(str)
}
func _toString(_u int32) string {
	return strconv.Itoa(int(_u))
}


func main() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}

