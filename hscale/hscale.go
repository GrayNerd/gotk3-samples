package main

import (
	"fmt"
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
	win.SetDefaultSize(300, 250)
	win.SetTitle("GtkHScale")
	win.SetBorderWidth(10)

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}

	hScale, err := gtk.ScaleNewWithRange(gtk.ORIENTATION_HORIZONTAL, 0, 100, 1.0)
	if err != nil {
		log.Fatal("Unable to create hScale", err)
	}
	hScale.SetProperty("value-pos", gtk.POS_RIGHT)

	box.PackStart(hScale, true, true, 0)
	win.Add(box)

	win.Connect("destroy", gtk.MainQuit)

	win.ShowAll()
	gtk.Main()
}

func valueChanged(h *gtk.Scale, l *gtk.Label) {
	val := h.GetValue()
	l.SetLabel(fmt.Sprintf("%3.1f", val))
}