		package main

import (
	"fmt"
	"log"
	_"reflect"
	"github.com/gotk3/gotk3/gdk"
	"unsafe"
	"github.com/gotk3/gotk3/gtk"
)

// #cgo pkg-config: gtk+-3.0
import "C"
	

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

	//https://stackoverflow.com/questions/28132798/creating-a-menu-in-gtk3-c
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

	// https://developer.gnome.org/gdk3/stable/gdk3-Event-Structures.html#GdkEventButton
	eventBox.Connect("button-press-event", func(event *gtk.EventBox) {
		
		ev := (*gdk.Event)(unsafe.Pointer(event))
		
		popupMenu.PopupAtPointer(ev)  
		
		fmt.Println("here")
	})

	win.ShowAll()
	gtk.Main()
}
