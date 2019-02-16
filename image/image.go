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
	win.SetDefaultSize(300, 200)
	win.SetTitle("Dog")

	image, err := gtk.ImageNewFromFile("dog.jpg")
	if err != nil {
		log.Fatal("Unable to load image from file:", err)
	}

	win.Add(image)
	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()
	gtk.Main()
}