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
	win.SetDefaultSize(400, 300)
	win.SetTitle("GtkFrame")
	win.SetBorderWidth(10)

	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetRowSpacing(20)
	grid.SetColumnSpacing(20)
	grid.SetRowHomogeneous(true)
	grid.SetColumnHomogeneous(true)

	frame1, _ := gtk.FrameNew("Shadow")
	frame1.SetShadowType(gtk.SHADOW_NONE)
	frame1.SetName("Shadow")
	frame2, _ := gtk.FrameNew("Rounded Shadow")
	frame2.SetShadowType(gtk.SHADOW_NONE)
	frame2.SetName("ShadowRounded")
	frame3, _ := gtk.FrameNew("Etched")
	frame3.SetShadowType(gtk.SHADOW_NONE)
	frame3.SetName("Etched")
	frame4, _ := gtk.FrameNew("Hover")
	frame4.SetShadowType(gtk.SHADOW_NONE)
	frame4.SetName("Hover")

	grid.Attach(frame1, 0, 0, 1, 1)
	grid.AttachNextTo(frame2, frame1, gtk.POS_RIGHT, 1, 1)
	grid.Attach(frame3, 0, 1, 1, 1)
	grid.Attach(frame4, 1, 1, 1, 1)

	win.Add(grid)
	win.ShowAll()
	loadCSS()
	win.Connect("destroy", gtk.MainQuit)
	gtk.Main()
}

func loadCSS() {
	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		log.Fatal("Unable to create display:", err)
	}
	provider, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal("Unable to create display:", err)
	}

	gtk.AddProviderForScreen(screen, provider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
	err = provider.LoadFromData(`
		#Shadow { border-style: none; box-shadow: 1px 2px 4px rgba(0, 0, 0, .5); padding: 10px; background: white; }
		#ShadowRounded { border-radius: 10px; box-shadow: 1px 2px 4px rgba(0, 0, 0, .5);
			padding: 10px; background: white; }
		#Etched{ border-style: inset; border-width: 5px; border-color: blue;
			padding: 10px; background: white;  }
	 	#Hover { box-shadow: 0px 10px 13px -7px #000000, 5px 5px 15px 5px rgba(0,0,0,0) }
	 `)
	if err != nil {
		log.Println(err)
	}
}
