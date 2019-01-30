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
	win.SetTitle("Toolbar")

	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	if err != nil {
		log.Fatal("Unable to create vbox:", err)
	}
	win.Add(vBox)

	toolbar, err := gtk.ToolbarNew()
	if err != nil {
		log.Fatal("Unable to create toolbar:", err)
	}
	toolbar.SetStyle(gtk.TOOLBAR_ICONS)

	img, err := gtk.ImageNew()
	if err != nil {
		log.Fatal("Unable to create image:", err)
	}
	img.SetFromIconName("document-new", gtk.ICON_SIZE_BUTTON)
	newTb, err := gtk.ToolButtonNew(img, "")
	if err != nil {
		log.Fatal("Unable to create newTb:", err)
	}
	toolbar.Insert(newTb, -1)

	img, err = gtk.ImageNew()
	if err != nil {
		log.Fatal("Unable to create image:", err)
	}
	img.SetFromIconName("document-open", gtk.ICON_SIZE_BUTTON)
	openTb, err := gtk.ToolButtonNew(img, "")
	if err != nil {
		log.Fatal("Unable to create openTb:", err)
	}
	toolbar.Insert(openTb, -1)

	img, err = gtk.ImageNew()
	if err != nil {
		log.Fatal("Unable to create image:", err)
	}
	img.SetFromIconName("document-save", gtk.ICON_SIZE_BUTTON)
	saveTb, err := gtk.ToolButtonNew(img, "")
	if err != nil {
		log.Fatal("Unable to create saveTb:", err)
	}
	toolbar.Insert(saveTb, -1)

	sep, err := gtk.SeparatorToolItemNew()
	if err != nil {
		log.Fatal("Unable to create separator", err)
	}
	toolbar.Insert(sep, -1)

	img, err = gtk.ImageNew()
	if err != nil {
		log.Fatal("Unable to create image:", err)
	}
	img.SetFromIconName("application-exit", gtk.ICON_SIZE_BUTTON)
	exitTb, err := gtk.ToolButtonNew(img, "")
	if err != nil {
		log.Fatal("Unable to create exitTb:", err)
	}
	toolbar.Insert(exitTb, -1)

	vBox.PackStart(toolbar, false, false, 5)

	exitTb.Connect("clicked", func() {
		gtk.MainQuit()
	})

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.ShowAll()
	gtk.Main()
}
