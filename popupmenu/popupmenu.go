		package main

import (
	"log"
	"github.com/gotk3/gotk3/gdk"
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
	win.SetTitle("Popup Menu")

	eventBox, err := gtk.EventBoxNew()
	if err != nil {
	log.Fatal("Unable to create eventBox:", err)
	}
	win.Add(eventBox)

	popupMenu, err := gtk.MenuNew()
	if err != nil {
		log.Fatal("Unable to create popupMenu", err)
	}

	hideMenuItem, err := gtk.MenuItemNewWithLabel("Minimize")
	if err != nil {
		log.Fatal("Unable to create hideMenuItem:", err)
	}
	hideMenuItem.Show()
	popupMenu.Append(hideMenuItem)

	quitMenuItem, err := gtk.MenuItemNewWithLabel("Quit")
	if err != nil {
		log.Fatal("Unable to create quitMenuItem:", err)
	}
	quitMenuItem.Show()
	popupMenu.Append(quitMenuItem)

	hideMenuItem.Connect("activate", func() {
		win.Iconify()
	})

	quitMenuItem.Connect("activate", func() {
		gtk.MainQuit()
	})

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	eventBox.Connect("button-press-event", func(win *gtk.Window, event *gdk.Event) {
		popupMenu.PopupAtPointer(event)  
	})

	win.ShowAll()
	gtk.Main()
}
