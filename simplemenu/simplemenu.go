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
	win.SetTitle("Simple Menu")

	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("Unable to create box1", err)
	}

	win.Add(vBox)

	//https://stackoverflow.com/questions/28132798/creating-a-menu-in-gtk3-c
	menuBar, err := gtk.MenuBarNew()
	if err != nil {
		log.Fatal("Unable to create menuBar", err)
	}
	fileMenu, err := gtk.MenuItemNewWithLabel("File")
	if err != nil {
		log.Fatal("Unable to create fileMenu:", err)
	}
	subMenu, err := gtk.MenuNew()
	if err != nil {
		log.Fatal("Unable to create subMenu:", err)
	}
	itemQuit, err := gtk.MenuItemNewWithLabel("Quit")
	if err != nil {
		log.Fatal("Unable to create itemQuit:", err)
	}

	subMenu.Append(itemQuit)
	fileMenu.SetSubmenu(subMenu)
	menuBar.Append(fileMenu)
	vBox.PackStart(menuBar, false, false, 0)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	itemQuit.Connect("activate", func() {
		gtk.MainQuit()
	})

	win.ShowAll()
	gtk.Main()
}

