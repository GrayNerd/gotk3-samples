package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

var count = 2

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
	img.SetFromIconName("edit-undo", gtk.ICON_SIZE_BUTTON)
	undoTb, err := gtk.ToolButtonNew(img, "")
	if err != nil {
		log.Fatal("Unable to create newTb:", err)
	}
	undoTb.SetName("undo")
	toolbar.Insert(undoTb, -1)

	img, err = gtk.ImageNew()
	if err != nil {
		log.Fatal("Unable to create image:", err)
	}
	img.SetFromIconName("edit-redo", gtk.ICON_SIZE_BUTTON)
	redoTb, err := gtk.ToolButtonNew(img, "")
	if err != nil {
		log.Fatal("Unable to create redoTb:", err)
	}
	redoTb.SetName("redo")
	toolbar.Insert(redoTb, -1)

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

	undoTb.Connect("clicked", func() {
		undoRedo(undoTb, redoTb)
	})
	redoTb.Connect("clicked", func() {
		undoRedo(redoTb, undoTb)
	})

	exitTb.Connect("clicked", func() {
		gtk.MainQuit()
	})

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.ShowAll()
	gtk.Main()
}

func undoRedo(btn *gtk.ToolButton, item *gtk.ToolButton) {
	name, err := btn.GetName()
	if err != nil {
		log.Fatal("Unable to get button name:", err)
	}
	if name == "undo" {
		count++
	} else {
		count--
	}

	if count < 0 || count > 5 {
		btn.SetSensitive(false)
		item.SetSensitive(true)
	}
}
