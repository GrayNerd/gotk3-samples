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
	win.SetTitle("Font Selection Dialog")
	win.SetBorderWidth(15)
	
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	win.Add(box)
	
	toolBar, err := gtk.ToolbarNew()
	if err != nil {
		log.Fatal("Unable to create toolbar:", err)
	}
	toolBar.SetBorderWidth(2)
	
	btn, err := gtk.FontButtonNew()
	if err != nil {
		log.Fatal("Unable to create font button:", err)
	}
	box.PackStart(btn, false, false, 0)

	newIcon, err := gtk.ImageNewFromIconName("preferences-desktop-font", gtk.ICON_SIZE_LARGE_TOOLBAR)
	if err != nil {
		log.Fatal("Unable to load preferences-desktop-font icon:", err)
	}
	font, err := gtk.ToolButtonNew(newIcon, "Font")
	if err != nil {
		log.Fatal("Unable to create toolbutton:", err)
	}
	toolBar.Add(btn)

	box.PackStart(toolBar, false, false, 5)

	label, err := gtk.LabelNew("Graynerd")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	box.PackStart(label, true, false, 5)



	font.Connect("clicked", selectFont)
	win.Connect("destroy", gtk.MainQuit)

	win.ShowAll()
	gtk.Main()
}

func selectFont() {
}
	
