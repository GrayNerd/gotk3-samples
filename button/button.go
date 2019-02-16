package main

import (
	"log"
	
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)
	
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(230, 200)
	win.SetTitle("GtkButton")
	win.SetBorderWidth(15)

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}

	btn, err := gtk.ButtonNewWithLabel("Quit")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	box.PackStart(btn, false, false, 0)
	win.Add(box)

	btn.Connect("clicked", gtk.MainQuit)
	win.Connect("destroy", gtk.MainQuit)

	win.ShowAll()

	gtk.Main()
}