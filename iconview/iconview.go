package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(300, 250)
	win.SetTitle("IconView")

	iconView, err := gtk.IconViewNew()
	if err != nil {
		log.Fatal("Unable to create iconview:", err)
	}
	iconView.SetColumns(2)
	iconView.SetItemWidth(100)

	listStore, err := gtk.ListStoreNew(createPixbufType(), glib.TYPE_STRING)
	if err != nil {
		log.Fatal("Unable to create liststore:", err)
	}
	iconView.SetModel(listStore)

	iconView.SetPixbufColumn(0)
	iconView.SetTextColumn(1)

	win.Add(iconView)

	// Icons made by freepik from www.flaticon.com
	homePath := os.Getenv("PWD") + "/iconview/"
	appendRow(listStore, homePath, "abacus.png", "Abacus")
	appendRow(listStore, homePath, "lightbulb.png", "Light Bulb")
	appendRow(listStore, homePath, "medal.png", "Medal")
	appendRow(listStore, homePath, "owl.png", "Owl")

	iconView.SetSelectionMode(gtk.SELECTION_MULTIPLE)

	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()
	gtk.Main()
}

func createPixbufType() glib.Type {
	pixbuf, err := gdk.PixbufNew(gdk.COLORSPACE_RGB, true, 8, 32, 32)
	if err != nil {
		log.Fatal("Unable to create pixbuf:", err)
	}
	return pixbuf.TypeFromInstance()
}

func appendRow(ls *gtk.ListStore, homepath, pix, text string) {
	iter := ls.Append()
	fpath := fmt.Sprintf("%s%s", homepath, pix)
	icon, err := gdk.PixbufNewFromFile(fpath)
	if err != nil {
		log.Fatal("Unable to get icon:", err)
	}

	ls.SetValue(iter, 0, icon)
	ls.SetValue(iter, 1, text)
}
