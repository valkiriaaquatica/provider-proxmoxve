package disk

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for the proxmox_lxc_disk resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("proxmox_lxc_disk", func(r *config.Resource) {
		r.ShortGroup = "Disk"
	})
}
