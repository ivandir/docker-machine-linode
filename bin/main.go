package main

import (
	"github.com/docker/machine/libmachine/drivers/plugin"
	linode "github.com/ivandir/docker-machine-linode"
)

func main() {
	plugin.RegisterDriver(linode.NewDriver("", ""))
}
