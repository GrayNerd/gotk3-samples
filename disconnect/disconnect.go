package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
)

var handleID glib.SignalHandle

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(300, 200)
	win.SetTitle("Disconnect")
	win.SetBorderWidth(15)

	hBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 15)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
		
	btn, err := gtk.ButtonNewWithLabel("Click")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	btn.SetSizeRequest(70,30)
	hBox.PackStart(btn, false, false, 0)

	cb, err := gtk.CheckButtonNewWithLabel("Connect")
	if err != nil {
		log.Fatal("Unable to create checkbox:", err)
	}
	cb.SetActive(true)
	hBox.PackStart(cb, false, false, 0)

	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 5)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	vBox.PackStart(hBox, false, false, 0)
	win.Add(vBox)

	h,err := btn.Connect("clicked", buttonClicked)
	if err != nil {
		log.Fatal("Unable to create btn.Connect:", err)
	}
	handleID = h

	cb.Connect("clicked", func() {
		toggleSignal(cb, btn)
	})

	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()

	gtk.Main()
}

func buttonClicked() {
	fmt.Println("clicked")
}

func toggleSignal(cb *gtk.CheckButton, btn *gtk.Button) {
	if cb.GetActive() == true {

		h, err := btn.Connect("clicked", buttonClicked)
		if err != nil {
			log.Fatal("Unable to reconnect btn.Connect:", err)
		}
		handleID = h
	} else {
		btn.HandlerDisconnect(handleID)
	}
}
