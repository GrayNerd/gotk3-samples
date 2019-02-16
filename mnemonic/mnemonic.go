package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window: ", err)
	}
	win.SetTitle("Mnemonic")
	win.SetDefaultSize(300, 200)
	win.SetBorderWidth(15)

	button, err := gtk.ButtonNewWithMnemonic("_Button")
	button.Connect("clicked", func() {
		fmt.Println("Button Clicked")
	})

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
