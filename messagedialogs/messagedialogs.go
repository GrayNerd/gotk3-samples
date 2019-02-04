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
	win.SetDefaultSize(220, 250)
	win.SetTitle("Message dialogs")
	win.SetBorderWidth(15)

	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}

	grid.SetRowSpacing(2)
	grid.SetColumnSpacing(2)

	info, _ := gtk.ButtonNewWithLabel("Info")
	warn, _ := gtk.ButtonNewWithLabel("Warning")
	qu, _ := gtk.ButtonNewWithLabel("Question")
	er, _ := gtk.ButtonNewWithLabel("Error")

	//gtk.Table is deprecated, using recommended replacement gtk.Grid
	grid.Attach(info, 0, 0, 1, 1)
	grid.AttachNextTo(warn, info, gtk.POS_RIGHT, 1, 1)
	grid.AttachNextTo(qu, info, gtk.POS_BOTTOM, 1, 1)
	grid.AttachNextTo(er, qu, gtk.POS_RIGHT, 1, 1)

	info.Connect("clicked", func() {
		showInfo(win)
	})
	warn.Connect("clicked", func() {
		showWarn(win)
	})
	qu.Connect("clicked", func() {
		showQuest(win)
	})
	er.Connect("clicked", func() {
		showError(win)
	})

	win.Add(grid)
	win.Connect("destory", gtk.MainQuit)
	win.ShowAll()
	gtk.Main()
}

func showInfo(win *gtk.Window) {
	dialog := gtk.MessageDialogNew(win, gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, "Download Completed")
	dialog.SetTitle("Information")
	dialog.Run()
}

func showWarn(win *gtk.Window) {
	dialog := gtk.MessageDialogNew(win, gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, "Unallowed operation")
	dialog.SetTitle("Warning")
	dialog.Run()
}

func showQuest(win *gtk.Window) {
	dialog := gtk.MessageDialogNew(win, gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_INFO, gtk.BUTTONS_OK_CANCEL, "Question")
	dialog.SetTitle("Are you sure you want to quit?")
	dialog.Run()
}

func showError(win *gtk.Window) {
	dialog := gtk.MessageDialogNew(win, gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, "Error loading file")
	dialog.SetTitle("Error")
	dialog.Run()
}
