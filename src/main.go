// Example window-names fetches a list of all top-level client windows managed
// by the currently running window manager, and prints the name and size
// of each window.
//
// This example demonstrates how to use some aspects of the ewmh and icccm
// packages. It also shows how to use the xwindow package to find the
// geometry of a client window. In particular, finding the geometry is
// intelligent, as it includes the geometry of the decorations if they exist.
package main

import (
	"log"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/icccm"
)

func main() {
	// Connect to the X server using the DISPLAY environment variable.
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	// Get a list of all client ids.
	clientids, err := ewmh.ClientListGet(X)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through each client, find its name
	for _, clientid := range clientids {
		name, err := ewmh.WmNameGet(X, clientid)

		// If there was a problem getting _NET_WM_NAME or if its empty,
		// try the old-school version.
		if err != nil || len(name) == 0 {
			name, err = icccm.WmNameGet(X, clientid)

			// If we still can't find anything, give up.
			if err != nil || len(name) == 0 {
				name = "N/A"
			}
		}

   println(name)
	}
}
