package main

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/graynerd/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal(err)
	}

	win.SetTitle("Search & highlight")
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetDefaultSize(350, 300)

	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal("Unable to create vertical box: ", err)
	}

	view, err := gtk.TextViewNew()
	if err != nil {
		log.Fatal("Unable to create toolbar: ", err)
	}
	view.AddEvents(gdk.GDK_CONTROL_MASK)
	vBox.PackStart(view, true, true, 0)

	buffer, err := view.GetBuffer()
	if err != nil {
		log.Fatal("Unable to get buffer: ", err)
	}

	table, err := buffer.GetTagTable()
	if err != nil {
		log.Fatal("Unable to get tag table", err)
	}

	tag, err := gtk.TextTagNew("gray_bg")
	if err != nil {
		log.Fatal("Unable to create text tag: ", err)
	}
	tag.SetProperty("background", "lightgray")
	table.Add(tag)

	win.Add(vBox)

	win.Connect("destroy", gtk.MainQuit)
	win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		event := gdk.EventKeyNewFromEvent(ev)
		modifiers := gtk.AcceleratorGetDefaultModMask()

		if (event.State() & uint(modifiers)) == uint(gdk.GDK_CONTROL_MASK) {
			switch event.KeyVal() {
			case gdk.KEY_m:
				start := buffer.GetIterAtMark(buffer.GetMark("insert"))
				end := buffer.GetIterAtMark(buffer.GetMark("selection_bound"))
				buf := start.GetBuffer()
				if start != end {
					t, err := table.Lookup("gray_bg", start, end)
					if err != nil {
						log.Fatal("Unable to get tag: ", err)
					}
					iter.
					buffer.RemoveTag(t, start, end)
					
					text, err := buffer.GetText(start, end, false)
					if err != nil {
						log.Fatal("Unable to get text: ", err)
					}
					buffer.Delete(start, end)
				}
			}
		}
	})

	win.ShowAll()

	gtk.Main()
}
