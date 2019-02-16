package main

import (
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal(err)
	}

	win.SetTitle("List view")
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.SetBorderWidth(10)
	win.SetDefaultSize(370, 270)

	adj, err := gtk.AdjustmentNew(10, 0, 0, 0, 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	sw, err := gtk.ScrolledWindowNew(adj, adj)
	if err != nil {
		log.Fatal(err)
	}

	list, err := gtk.TreeViewNew() 
	if err != nil {
		log.Fatal(err)
	}

	sw.Add(list)
	sw.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	sw.SetShadowType(gtk.SHADOW_ETCHED_IN)
	list.SetHeadersVisible(false)

	vBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	if err != nil {
		log.Fatal(err)
	}
	vBox.PackStart(sw, true, true, 5)

	hBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if err != nil {
		log.Fatal(err)
	}

	addBtn, err := gtk.ButtonNewWithLabel("Add")
	if err != nil {
		log.Fatal(err)
	}

	remBtn, err := gtk.ButtonNewWithLabel("Remove")
	if err != nil {
		log.Fatal(err)
	}
	remAllBtn, err := gtk.ButtonNewWithLabel("Remove All")
	if err != nil {
		log.Fatal(err)
	}

	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal(err)
	}

	hBox.PackStart(addBtn, false, true, 3)
	hBox.PackStart(entry, false, true, 3)
	hBox.PackStart(remBtn, false, true, 3)
	hBox.PackStart(remAllBtn, false, true, 3)

	vBox.PackStart(hBox, false, true, 3)
	win.Add(vBox)

	store, err := initList(list)
	if err != nil {
		log.Fatal(err)
	}

	selection, err := list.GetSelection()
	if err != nil {
		log.Fatal(err)
	}

	addBtn.Connect("clicked", func() { appendItem(store, entry) })
	remBtn.Connect("clicked", func() { removeItem(store, selection) })
	remAllBtn.Connect("clicked", func() { removeAll(store, selection) })
	win.Connect("destroy", gtk.MainQuit)

	win.ShowAll()
	gtk.Main()
}

func initList(list *gtk.TreeView) (*gtk.ListStore, error) {
	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		return nil, err
	}

	column, err := gtk.TreeViewColumnNewWithAttribute("List Items", cellRenderer, "text", 0)
	if err != nil {
		return nil, err
	}
	list.AppendColumn(column)

	store, err := gtk.ListStoreNew(glib.TYPE_STRING)
	if err != nil {
		return nil, err
	}
	list.SetModel(store)

	return store, nil
}
func appendItem(store *gtk.ListStore, entry *gtk.Entry) {
	str, err := entry.GetText()
	if err != nil {
		log.Fatal(err)
	}
	if len(str) > 0 {
		iter := store.Append()
		store.SetValue(iter, 0, str)
		entry.SetText("")
	}
}

func removeItem(store *gtk.ListStore, selection *gtk.TreeSelection) {
	_, iter, ok := selection.GetSelected()
	if ok {
		store.Remove(iter)
	}
}

func removeAll(store *gtk.ListStore, selection *gtk.TreeSelection) {
	_, ok := store.GetIterFirst()
	if ok {
		store.Clear()
	}
}
