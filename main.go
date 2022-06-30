package main

import (
	"fmt"
	"os"

	chefclient "github.com/bdwyertech/packer-plugin-chef/provisioner/chef-client"
	chefsolo "github.com/bdwyertech/packer-plugin-chef/provisioner/chef-solo"

	"github.com/bdwyertech/packer-plugin-chef/version"
	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterProvisioner("client", new(chefclient.Provisioner))
	pps.RegisterProvisioner("solo", new(chefsolo.Provisioner))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
