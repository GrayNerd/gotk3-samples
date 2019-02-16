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
	win.SetTitle("GtkAlignment")
	win.SetDefaultSize(300,200)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetBorderWidth(5)

	// as per the documentation for GtkAlignment: 
	// GtkAlignment has been deprecated in 3.14 and should not be used in newly-written code. 
	// The desired effect can be achieved by using the “halign”, “valign” and “margin” 
	// properties on the child widget.

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}

	lbl, err := gtk.LabelNew("bottom-left")
	if err != nil {
		log.Fatal("Unable to create label", err)
	}
	lbl.SetHAlign(gtk.ALIGN_START)
	lbl.SetVAlign(gtk.ALIGN_END)

	// need this to put it all the way to the bottom, not sure why
	lbl.SetVExpand(true)
	
	box.Add(lbl)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.Add(box)
	win.ShowAll()
	gtk.Main()
}