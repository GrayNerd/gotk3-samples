package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window: ", err)
	}
	win.SetTitle("Tooltip")
	win.SetDefaultSize(300, 200)
	win.SetBorderWidth(15)

	button, err := gtk.ButtonNewWithLabel("Button")
	button.SetTooltipText("Button widget")

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 15)	
	box.SetHAlign(gtk.ALIGN_START)
	if err != nil {
		log.Fatal("Unable to create box: ", err)
	}
	
	box.Add(button)
	win.Add(box)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.ShowAll()

	gtk.Main()
}
