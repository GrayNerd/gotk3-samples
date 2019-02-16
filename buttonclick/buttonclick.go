package main

import (
	"log"
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(300, 200)
	win.SetTitle("GtkButton")
	win.SetBorderWidth(15)

	box, err:=gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	
	btn, err := gtk.ButtonNewWithLabel("Click")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	btn.SetSizeRequest(70,30)

	box.SetHAlign(gtk.ALIGN_START)
	box.Add(btn)
	win.Add(box)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	btn.Connect("clicked", buttonClicked)

	win.ShowAll()
	gtk.Main()
}

func buttonClicked() {
	fmt.Println("clicked")
}