package pool

import "github.com/upbound/upjet/pkg/config"

// Configure adds configurations for the proxmox_pool resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("proxmox_pool", func(r *config.Resource) {
		r.ShortGroup = "Pool"
	})
}
