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
	win.SetDefaultSize(300, 200)
	win.SetTitle("GtkStatusbar")
	win.SetBorderWidth(10)

	vBox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 10)
	hBox, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)

	btnOK, _ := gtk.ButtonNewWithLabel("OK")
	btnApply, _ := gtk.ButtonNewWithLabel("Apply")
	hBox.PackStart(btnOK, false, false, 5)
	hBox.PackStart(btnApply, false, false, 0)
	vBox.PackStart(hBox, false, false, 0)

	tv, _ := gtk.TextViewNew()
	vBox.PackStart(tv, true, true, 0)

	sb, _ := gtk.StatusbarNew()
	vBox.PackStart(sb, false, false, 0)
	
	si, _ := gtk.ImageNewFromIconName("face-monkey", gtk.ICON_SIZE_LARGE_TOOLBAR)
	sb.PackStart(si, false, false, 10)
	
	btnOK.Connect("clicked", func() {
		buttonPressed(btnOK, sb)
	})
	btnApply.Connect("clicked", func() {
		buttonPressed(btnApply, sb)
	})

	win.Add(vBox)
	win.Connect("destroy", gtk.MainQuit)

	win.ShowAll()
	gtk.Main()
}

func buttonPressed(btn *gtk.Button, sb *gtk.Statusbar) {
	id := sb.GetContextId("example")
	str, _ := btn.GetLabel()
	sb.Push(id, fmt.Sprintf("%s button pushed", str))
}
