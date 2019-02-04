package main

import (
	"fmt"
	"log"
	"time"
	
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
)

var buf string

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	da, err := gtk.DrawingAreaNew()
	if err != nil {
		log.Fatal("Unable to create drawing area:", err)
	}
	win.Add(da)
	da.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context){
		cr.SetSourceRGB(0,0,0)
		cr.MoveTo(30,30)
		cr.SetFontSize(15)
		cr.ShowText(buf)
	})

	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(300, 200)
	win.SetTitle("Timer")
	win.SetBorderWidth(15)

	glib.TimeoutAdd(1000, timeHandler, win)

	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()
	timeHandler(win)
	gtk.Main()
}

func timeHandler(win *gtk.Window) bool {
	if win == nil {
		return false
	}

	t := time.Now()
	
	format := "15:04:05"
	buf = fmt.Sprint(t.Format(format))
	win.QueueDraw()

	return true
}
