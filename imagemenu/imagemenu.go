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

	itemBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)
	if err != nil {
		log.Fatal("Unable to create itemBox:", err)
	}

	itemPix, err := gtk.ImageNewFromIconName("contact-new",gtk.ICON_SIZE_MENU)
	if err != nil {
		log.Fatal("Unable to create itemPix:", err)
	}
	itemBox.PackStart(itemPix, false, false, 0)

	itemLabel, err := gtk.LabelNew("New")
	if err != nil {
		log.Fatal("Unable to create itemLabel:", err)
	}
	itemBox.PackStart(itemLabel, false, false, 0)

	newMenuItem, err := gtk.MenuItemNew()
	if err != nil {
		log.Fatal("Unable to create newMenuItem:", err)
	}
	newMenuItem.Add(itemBox)
		
	subMenu, err := gtk.MenuNew()
	if err != nil {
		log.Fatal("Unable to create subMenu:", err)
	}
	subMenu.Append(newMenuItem)
	
	quitItemBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL,0)
	if err != nil {
		log.Fatal("Unable to create itemBox:", err)
	}

	quitItemPix, err := gtk.ImageNewFromIconName("application-exit", gtk.ICON_SIZE_MENU)
	if err != nil {
		log.Fatal("Unable to create itemPix:", err)
	}
	quitItemBox.PackStart(quitItemPix, false, false, 0)

	quitItemLabel, err := gtk.LabelNew("Quit")
	if err != nil {
		log.Fatal("Unable to create itemLabel:", err)
	}
	quitItemBox.PackStart(quitItemLabel, false, false, 0)

	quitMenuItem, err := gtk.MenuItemNew()
	if err != nil {
		log.Fatal("Unable to create newMenuItem:", err)
	}
	quitMenuItem.Add(quitItemBox)

	subMenu.Append(quitMenuItem)
	
	fileMenuItem.SetSubmenu(subMenu)
	menuBar.Append(fileMenuItem)
	vBox.PackStart(menuBar, false, false, 0)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	quitMenuItem.Connect("activate", func() {
		gtk.MainQuit()
	})

	win.ShowAll()
	gtk.Main()
}
