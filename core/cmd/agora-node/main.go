package main

import "github.com/Jaime2003z/Agora/core/node"

func main() {
	node, err := node.NewNode()
	if err != nil {
		panic(err)
	}

	node.RegisterProtocols()
	node.SetupDiscovery()

	select {}
}
