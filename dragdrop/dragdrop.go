package main

// the zet-code version could not be implemented as the gtk_window_begin_move_drag function
// is not currently in gotk3
// this version is a normal drag and drop, just drop anything onto the label in the window.
// the label text will change to what was dropped (only tested with files and webpage text)

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
	win.SetDefaultSize(250, 200)
	win.SetTitle("Drag & Drop")
	win.SetBorderWidth(15)

	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	if err != nil {
		log.Fatal("Unable to create box:", err)
	}

	lbl, err := gtk.LabelNew("Drop here!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	box.PackStart(lbl, true, true, 0)
	win.Add(box)
	entry, err := gtk.TargetEntryNew("text/plain", gtk.TARGET_OTHER_APP, 0)
	if err != nil {
		log.Fatal("Unable to create Target Entry:", err)
	}
	target := []gtk.TargetEntry{*entry}
	
	win.DragDestSet(gtk.DEST_DEFAULT_ALL, target, gdk.ACTION_COPY)

	win.Connect("drag-data-received", func(win *gtk.Window, context *gdk.DragContext, row, col int, data uintptr, m int, t uint) {
		lbl.SetLabel(string(gtk.GetData(data)))
		log.Printf("received: %s", string(gtk.GetData(data)))
	})

	// next to consume the drap-drop signal or drag-data-received does not work
	win.Connect("drag-drop", func() {
		log.Printf("dropped")
	})

	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()
	gtk.Main()
}
