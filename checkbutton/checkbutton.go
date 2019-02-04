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
	win.SetTitle("GtkCheckButton")
	win.SetBorderWidth(15)

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}

	cb, err := gtk.CheckButtonNewWithLabel("Show title")
	cb.SetActive(true)

	cb.SetCanFocus(false)
	box.PackStart(cb, false, false, 0)

	win.Connect("destroy", gtk.MainQuit)
	cb.Connect("clicked", func() {
		toggleTitle(cb, win)
	})

	win.Add(box)
	win.ShowAll()

	gtk.Main()
}

func toggleTitle(cb *gtk.CheckButton, win *gtk.Window) {
	if cb.GetActive() == true {
		win.SetTitle("GtkCheckButton")
	} else {
		win.SetTitle("")
	}
}
