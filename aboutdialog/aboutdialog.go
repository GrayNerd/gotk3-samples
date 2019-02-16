package main

import (
	"log"
	
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/gdk"
)

func main() {
	gtk.Init(nil)

	win, err:= gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(220, 200)
	win.SetTitle("Battery")
	win.SetBorderWidth(15)

	win.Connect("button-press-event", func(win *gtk.Window, event *gdk.Event) {
		showAbout(win)
	})

	win.Connect("destroy", gtk.MainQuit)

	win.ShowAll()
	gtk.Main()
}

func showAbout(win *gtk.Window) {
	pixbuf, err := gdk.PixbufNewFromFile("battery.png")
	if err != nil {
		log.Fatal("Unable to load pixbuf from file:", err)
	}
	dialog, err := gtk.AboutDialogNew()
	if err != nil {
		log.Fatal("Unable to create dialog:", err)
	}
	
	dialog.SetTitle("Battery")
	dialog.SetVersion("0.9")
	dialog.SetCopyright("n/a")
	dialog.SetComments("Battery is a simple tool for battery checking.")
	dialog.SetWebsite("http://www.batteryhq.net")
	dialog.SetLogo(pixbuf)

	dialog.SetTransientFor(win)
	dialog.Run()
}