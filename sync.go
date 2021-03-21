package main

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/erikh/gui-expo.vue/pkg/pollsync"
	one "github.com/erikh/gui-expo.vue/pkg/zt-one"
	"github.com/webview/webview"
)

func syncInit(ctx context.Context, w webview.WebView) {
	c := one.NewClient(os.Getenv("ZEROTIER_ONE_TOKEN"))
	ps := pollsync.New(ctx)

	ps.Register(500*time.Millisecond, "networks", func(ctx context.Context) (interface{}, error) {
		return c.Networks()
	})

	w.Bind("getNetworks", func() (string, error) {
		content, err := json.Marshal(ps.Data("networks"))
		return string(content), err
	})

	w.Bind("getNetwork", func(id string) (string, error) {
		network, err := c.Network(id)
		if err != nil {
			return "", err
		}

		content, err := json.Marshal(network)
		return string(content), err
	})
}
