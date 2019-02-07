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
	win.SetTitle("GtkEntry")
	win.SetResizable(false)
	win.SetBorderWidth(10)

	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetRowSpacing(20)
	grid.SetColumnSpacing(10)

	label1, _ := gtk.LabelNew("Name")
	label2, _ := gtk.LabelNew("Age")
	label3, _ := gtk.LabelNew("Occupation")

	grid.Attach(label1, 0,0,1,1)
	grid.Attach(label2, 0,1,1,1)
	grid.Attach(label3, 0,2,1,1)

	entry1, _ := gtk.EntryNew()
	entry2, _ := gtk.EntryNew()
	entry3, _ := gtk.EntryNew()

	grid.Attach(entry1, 1,0,1,1)
	grid.Attach(entry2, 1,1,1,1)
	grid.Attach(entry3, 1,2,1,1)

	win.Add(grid)

	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()
	gtk.Main()
}