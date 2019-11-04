package main

import "github.com/solo-io/autopilot/codegen"

func main() {
	if err := codegen.Run("examples/quarantine/", true); err != nil {
		panic(err)
	}
}