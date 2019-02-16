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
	win.SetTitle("GtkFixed")
	win.SetDefaultSize(300, 200)
	win.SetPosition(gtk.WIN_POS_CENTER)

	fixed, err := gtk.FixedNew()
	if err != nil {
		log.Fatal("Unable to create fixed layout:", err)
	}
	win.Add(fixed)

	btn1, err := gtk.ButtonNewWithLabel("Button")
	if err != nil {
		log.Fatal("Unable to create btn1:", err)
	}
	fixed.Put(btn1, 150, 50)
	btn1.SetSizeRequest(80, 30)

	btn2, err := gtk.ButtonNewWithLabel("Button")
	if err != nil {
		log.Fatal("Unable to create btn2:", err)
	}
	fixed.Put(btn2, 15, 15)
	btn2.SetSizeRequest(80, 30)

	btn3, err := gtk.ButtonNewWithLabel("Button")
	if err != nil {
		log.Fatal("Unable to create btn3:", err)
	}
	fixed.Put(btn3, 100, 100)
	btn3.SetSizeRequest(80, 30)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.ShowAll()
	gtk.Main()
}
