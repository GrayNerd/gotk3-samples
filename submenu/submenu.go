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
	win.SetTitle("SubMenu")

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
	fileMenuItem, err := gtk.MenuItemNewWithLabel("File")
	if err != nil {
		log.Fatal("Unable to create fileMenu:", err)
	}

	imprMenuItem, err := gtk.MenuItemNewWithLabel("Import")
	imprMI, err := gtk.MenuItemNewWithLabel("Import news feed...")
	feedMI, err := gtk.MenuItemNewWithLabel("Import mail...")
	bookMI, err := gtk.MenuItemNewWithLabel("Import bookmarks...")

	importSubMenu, err := gtk.MenuNew()
	if err != nil {
		log.Fatal("Unable to create subMenu:", err)
	}

	importSubMenu.Append(imprMI)
	importSubMenu.Append(feedMI)
	importSubMenu.Append(bookMI)
	imprMenuItem.SetSubmenu(importSubMenu)

	subMenu, err := gtk.MenuNew()
	subMenu.Append(imprMenuItem)
	
	sep, err := gtk.SeparatorMenuItemNew()
	if err != nil {
		log.Fatal("Unable to create Separator", err)
	}
	subMenu.Append(sep)
	
	itemQuit, err := gtk.MenuItemNewWithLabel("Quit")
	if err != nil {
		log.Fatal("Unable to create itemQuit:", err)
	}
	subMenu.Append(itemQuit)
	
	fileMenuItem.SetSubmenu(subMenu)
	menuBar.Append(fileMenuItem)
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
