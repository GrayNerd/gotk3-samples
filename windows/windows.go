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
	win.SetDefaultSize(300, 250)
	win.SetTitle("Windows")
	win.SetBorderWidth(15)

	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetProperty("column-spacing", 5)
	grid.SetProperty("row-spacing", 5)

	grid.SetProperty("column-homogeneous", true)
	grid.SetProperty("row-homogeneous", true)

	win.Add(grid)

	label, err := gtk.LabelNewWithMnemonic("_Windows")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	label.SetHAlign(gtk.ALIGN_START)
	grid.Attach(label, 0, 0, 1, 1)

	tv, err := gtk.TextViewNew()
	if err != nil {
		log.Fatal("Unable to create textview:", err)
	}
	tv.SetEditable(false)
	tv.SetCursorVisible(false)
	tv.SetHExpand(true)
	tv.SetVExpand(true)
	// tv.SetSizeRequest(200, 50)
	grid.Attach(tv, 0, 1, 4, 10)

	actBtn, err := gtk.ButtonNewWithMnemonic("_Activate")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	actBtn.SetSizeRequest(50, 30)
	actBtn.SetVExpand(false)
	grid.AttachNextTo(actBtn, tv, gtk.POS_RIGHT, 1, 1)

	closeBtn, err := gtk.ButtonNewWithMnemonic("_Close")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	closeBtn.SetSizeRequest(50, 30)
	grid.AttachNextTo(closeBtn, actBtn, gtk.POS_BOTTOM, 1, 1)

	helpBtn, err := gtk.ButtonNewWithMnemonic("_Help")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	helpBtn.SetSizeRequest(50, 30)
	// helpBtn.SetHAlign(gtk.ALIGN_START)
	// helpBtn.SetHExpand(true)
	grid.Attach(helpBtn, 0,12,1, 1)

	okBtn, err := gtk.ButtonNewWithLabel("OK")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	okBtn.SetSizeRequest(50, 30)
	okBtn.SetVAlign(gtk.ALIGN_END)
	grid.Attach(okBtn, 6, 12, 1, 1)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.ShowAll()
	gtk.Main()
}
