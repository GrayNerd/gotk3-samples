package main

import (
	"log"
	"strconv"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(250, 180)
	win.SetTitle("GtkTable")
	win.SetBorderWidth(5)

	//gtk.Table is deprecated, gtk.Grid is the recommended replacement
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetColumnHomogeneous(true)
	grid.SetRowHomogeneous(true)
	for c := 1; c < 5; c++ {
		grid.InsertRow(c)
		grid.InsertColumn(c)
	}
	pos := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			btn, err := gtk.ButtonNewWithLabel(strconv.Itoa(pos))
			if err != nil {
				log.Fatal("Unable to create button:", err)
			}
			grid.Attach(btn, i, j, 1, 1)
			pos++
		}
	}
	win.Add(grid)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.ShowAll()
	gtk.Main()
}
