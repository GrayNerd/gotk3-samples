package main

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)
//#cgo pkg-config: gtk+-3.0
import "C"

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(300, 200)
	win.SetTitle("GtkButton")
	
	// TODO: fix this
	//gtk_widget_add_events(GTK_WIDGET(window), GDK_CONFIGURE);
	//win.AddEvents(gdk.EXPOSURE_MASK)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.Connect("configure-event", func(win *gtk.Window, event *gdk.Event){
		configureCallback(win, event)
	})

	win.ShowAll()
	gtk.Main()
}

func configureCallback(win *gtk.Window, event *gdk.Event) {
	// TODO: fix this
	// there does not seem to be an gdk.EventConfigure in gotk3
	z := *C.GdkEventConfigure(unsafe.Pointer(event.native()))
	// z := gdk.EventButtonNewFromEvent(event)
	x := z.X
	y := z.Y
	
	win.SetTitle(fmt.Sprintf("here %d %d", x, y))
}
