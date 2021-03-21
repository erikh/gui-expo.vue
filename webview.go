package main

import (
	"context"
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

	data, err := fs.ReadFile("frontend/build/bundle.js")
	if err != nil {
		panic(err)
	}

	w.Init("window.onload = () => {" + string(data) + "};")
	data, err = fs.ReadFile("frontend/public/index.html")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	w.Dispatch(func() {
		syncInit(ctx, w)
	})

	w.Dispatch(func() {
		systray.Register(func() {
			systray.SetIcon(iconData)
			systray.SetTooltip("Minimal vue example")
		}, func() {
		})
	})

	w.Navigate("data:text/html," + string(data))
	w.Run()

	cancel()
	fmt.Println("Exiting")
}
