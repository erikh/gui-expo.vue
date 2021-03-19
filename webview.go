package main

import (
	"embed"
	"fmt"

	"github.com/getlantern/systray"
	"github.com/webview/webview"
)

//go:embed frontend/build
//go:embed frontend/public
var fs embed.FS

func main() {
	debug := true

	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal vue example")

	w.Bind("clicked", func() error {
		fmt.Println("clicked")
		return nil
	})

	data, err := fs.ReadFile("frontend/build/bundle.js")
	if err != nil {
		panic(err)
	}

	w.Init("window.onload = () => {" + string(data) + "};")
	data, err = fs.ReadFile("frontend/public/index.html")
	if err != nil {
		panic(err)
	}

	w.Navigate("data:text/html," + string(data))
	w.Dispatch(func() {
		systray.Run(func() {
			systray.SetIcon(iconData)
			systray.SetTooltip("Minimal vue example")
		}, func() {})
	})
	w.Run()

	// Enter the host system's event loop

	// This is only reached once the user chooses the Exit menu item
	fmt.Println("Exiting")
}
