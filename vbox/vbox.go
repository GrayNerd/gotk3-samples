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
	win.SetDefaultSize(230, 250)
	win.SetTitle("GtkVBox")
	win.SetBorderWidth(5)

	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	if err != nil {
		log.Fatal("Unable to create box container:", err)
	}
	win.Add(vBox)

	settings, err := gtk.ButtonNewWithLabel("Settings")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	accounts, err := gtk.ButtonNewWithLabel("Accounts")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	loans, err := gtk.ButtonNewWithLabel("Loans")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	cash, err := gtk.ButtonNewWithLabel("Cash")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	debts, err := gtk.ButtonNewWithLabel("Debts")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	vBox.PackStart(settings, true, true, 0)
	vBox.PackStart(accounts, true, true, 0)
	vBox.PackStart(loans, true, true, 0)
	vBox.PackStart(cash, true, true, 0)
	vBox.PackStart(debts, true, true, 0)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.ShowAll()
	gtk.Main()
}
