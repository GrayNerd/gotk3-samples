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
	win.SetDefaultSize(350, 200)
	win.SetTitle("GtkTable")
	win.SetBorderWidth(10)

	// GtkTable has been deprecated, GtkBox is probably the easiest to use in this application
	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	
	hBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 3)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}

	okBtn, err := gtk.ButtonNewWithLabel("OK")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	okBtn.SetSizeRequest(70,30)

	closeBtn, err := gtk.ButtonNewWithLabel("Close")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	closeBtn.SetSizeRequest(70,30)

	hBox.PackEnd(closeBtn, false, false, 0)
	hBox.PackEnd(okBtn, false, false, 0)
	vBox.PackEnd(hBox, false, false, 0)
	win.Add(vBox)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.ShowAll()
	gtk.Main()
}