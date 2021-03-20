package main

import (
	"context"
	"embed"
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/webview/webview"
	"github.com/zerotier/go-ztcentral"
)

//go:embed frontend/build
//go:embed frontend/public
var fs embed.FS

func getNetworkName(id string) (string, error) {
	client := ztcentral.NewClient(os.Getenv("ZEROTIER_CENTRAL_TOKEN"))
	network, err := client.GetNetwork(context.Background(), id)
	if err != nil {
		return "", err
	}

	return network.Config.Name, nil
}

func main() {
	debug := true

	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal vue example")

	w.Bind("getNetworkName", getNetworkName)

	data, err := fs.ReadFile("frontend/build/bundle.js")
	if err != nil {
		panic(err)
	}

	w.Init("window.onload = () => {" + string(data) + "};")
	data, err = fs.ReadFile("frontend/public/index.html")
	if err != nil {
		panic(err)
	}

	w.Dispatch(func() {
		systray.Register(func() {
			systray.SetIcon(iconData)
			systray.SetTooltip("Minimal vue example")
		}, func() {
		})
	})

	w.Navigate("data:text/html," + string(data))
	w.Run()

	fmt.Println("Exiting")
}
