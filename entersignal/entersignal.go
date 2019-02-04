package main

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
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
	win.SetTitle("Enter signal")
	win.SetBorderWidth(15)

	btn, err := gtk.ButtonNewWithLabel("Button")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	btn.SetSizeRequest(70, 30)
	btn.SetVAlign(gtk.ALIGN_START)

	//name the button for the css coloring that is used in GTK+3
	btn.SetName("button1")

	vBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	hBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}
	hBox.Add(btn)
	vBox.Add(hBox)
	win.Add(vBox)

	win.Connect("destroy", gtk.MainQuit)
	btn.Connect("enter", enterButton)

	// added this feature
	// watch for the leave event and change the color back
	btn.Connect("leave", leaveButton)
	
	win.ShowAll()
	gtk.Main()
}

// the modify_bg has been deprecated and css is now the recommended
// way of changing the color
func enterButton(btn *gtk.Button) {
	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		log.Fatal("Unable to create display:", err)
	}
	provider, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal("Unable to create display:", err)
	}

	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	provider.LoadFromData("#button1 { background-color: red; color: white; }")
}

// not in the original zetcode, but this will change the button back when the cursor
// exits the button
func leaveButton(btn *gtk.Button) {
	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		log.Fatal("Unable to create display:", err)
	}
	provider, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal("Unable to create display:", err)
	}
	btn.SetName("button1")

	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	provider.LoadFromData("#button1 { background-color: white; color: black; }")
}
