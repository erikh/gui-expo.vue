package main

import (
	"fmt"
	"os"

	one "github.com/erikh/gui-expo.vue/pkg/zt-one"
)

func main() {
	c := one.NewClient(os.Getenv("ZEROTIER_ONE_TOKEN"))
	networks, err := c.Networks()
	if err != nil {
		panic(err)
	}

	peers, err := c.Peers()
	if err != nil {
		panic(err)
	}

	fmt.Println("Networks w/ MAC:")
	for _, network := range networks {
		fmt.Println(network.ID, network.MAC)
	}

	fmt.Println("Peers w/ Latency:")
	for _, peer := range peers {
		fmt.Println(peer.Address, peer.Latency)
	}
}
