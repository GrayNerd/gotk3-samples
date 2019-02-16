package main

import (
	"log"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/gdk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(300, 200)
	win.SetTitle("GtkCheckMenuItem")

	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal("Unable to create vbox:", err)
	}
	win.Add(vBox)

	//https://stackoverflow.com/questions/28132798/creating-a-menu-in-gtk3-c
	menuBar, err := gtk.MenuBarNew()
	if err != nil {
		log.Fatal("Unable to create menuBar:", err)
	}
	viewMenuItem, err := gtk.MenuItemNewWithLabel("View")
	if err != nil {
		log.Fatal("Unable to create viewMenuItem:", err)
	}

	toggleMenuItem, err := gtk.CheckMenuItemNewWithLabel("View statusbar")
	if err != nil {
		log.Fatal("Unable to create newMenuItem:", err)
	}

	subMenu, err := gtk.MenuNew()
	if err != nil {
		log.Fatal("Unable to create subMenu:", err)
	}
	subMenu.Append(toggleMenuItem)


	viewMenuItem.SetSubmenu(subMenu)
	menuBar.Append(viewMenuItem)
	vBox.PackStart(menuBar, false, false, 0)

	// statusbar will not appear unless something is added to it
	// in this case a label
	statusbar,err := gtk.StatusbarNew()
	if err != nil {
		log.Fatal("Unable to create statusbar:", err)
	}
	
	label, err := gtk.LabelNew("Status Label")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	
	statusbar.Add(label)
	vBox.PackEnd(statusbar, false, true, 0)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	toggleMenuItem.Connect("activate", func(e gdk.Event) {
		fmt.Println(e)
		toggleStatusbar(toggleMenuItem, statusbar) 
	})

	win.ShowAll()
	gtk.Main()
}

func toggleStatusbar(item *gtk.CheckMenuItem, statusbar *gtk.Statusbar) {
	if item.GetActive() == true {
		statusbar.Show()
	} else {
		statusbar.Hide()
	}
}