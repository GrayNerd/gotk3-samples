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
	win.SetDefaultSize(280, 200)
	win.SetTitle("Color Selection Dialog")
	win.SetBorderWidth(15)

	toolBar, err := gtk.ToolbarNew()
	if err != nil {
		log.Fatal("Unable to create toolbar:", err)
	}
	toolBar.SetBorderWidth(2)
	
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	win.Add(box)

	img, err := gtk.ImageNew()
	if err != nil {
		log.Fatal("Unable to create image:", err)
	}
	img.SetFromIconName("applications-graphics", gtk.ICON_SIZE_BUTTON)
	colorTb, err := gtk.ToolButtonNew(img, "")
	if err != nil {
		log.Fatal("Unable to create newTb:", err)
	}
	colorTb.SetName("undo")
	toolBar.Insert(colorTb, -1)

	box.Add(toolBar)

	label, err := gtk.LabelNew("Graynerd")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	label.SetJustify(gtk.JUSTIFY_LEFT)
	box.PackStart(label, true, false, 5)

	colorTb.Connect("clicked", func() {
		dialog, err := gtk.ColorChooserDialogNew("Choose Color", win)
		if err != nil {
			log.Fatal("Unable to create color chooser dialog:", err)
		}
		dialog.Run()
	})

	win.Connect("destroy", gtk.MainQuit)

	win.ShowAll()
	gtk.Main()
}