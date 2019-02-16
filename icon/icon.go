package main

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func createPixBuf(filename string) *gdk.Pixbuf {
	pixbuf, err := gdk.PixbufNewFromFile(filename)
	if err != nil {
		log.Fatal("Unable to create pixbuf: ", err)
	}
	return pixbuf
}

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window: ", err)
	}
	win.SetTitle("Icon")
	win.SetDefaultSize(230, 150)
	win.SetPosition(gtk.WIN_POS_CENTER)

	//icon is from https://iconmonstr.com
	err = win.SetIconFromFile("globe.png") 
	if err != nil {
		log.Fatal("Unable to load icon: ", err)
	}

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.Show()

	gtk.Main()
}
