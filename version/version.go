package main

import "fmt"

// #cgo pkg-config: gtk+-3.0
// #include <glib.h>
import "C"

import (
	"github.com/gotk3/gotk3/gtk"
)


func main() {
	gtk.Init(nil)

	fmt.Printf("GTK+ version: %d.%d.%d\n", gtk.GetMajorVersion(), gtk.GetMinorVersion(), gtk.GetMicroVersion())
	fmt.Printf("Glib version: %d.%d.%d\n", C.GLIB_MAJOR_VERSION, C.GLIB_MINOR_VERSION, C.GLIB_MICRO_VERSION)
}