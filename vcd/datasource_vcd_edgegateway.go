package vcd

import "github.com/hashicorp/terraform/helper/schema"

func datasourceVcdEdgeGateway() *schema.Resource {
	return &schema.Resource{
		Read: datasourceVcdEdgeGatewayRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"org": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vdc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"advanced": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if the gateway uses advanced networking. (Enabled by default)",
			},
			"configuration": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Configuration of the vShield edge VM for this gateway. One of: compact, full ("Large"), full4 ("Quad Large"), x-large`,
			},
			"ha_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable high availability on this edge gateway",
			},
			"external_networks": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of external networks to be used by the edge gateway",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"external_networks_ip": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of IP addresses for each external network interface.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"external_networks_netmask": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of netmasks for each external network interface.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"external_networks_gateway": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of gateways for each external network interface.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"default_gateway_network": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External network to be used as default gateway. Its name must be included in 'external_networks'. An empty value will skip the default gateway",
			},
			"distributed_routing": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If advanced networking enabled, also enable distributed routing",
			},
			"lb_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable load balancing. (Disabled by default)",
			},
			"lb_acceleration_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable load balancer acceleration. (Disabled by default)",
			},
			"lb_logging_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable load balancer logging. (Disabled by default)",
			},
			"lb_loglevel": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Description: "Log level. One of 'emergency', 'alert', 'critical', 'error', " +
					"'warning', 'notice', 'info', 'debug'. ('info' by default)",
			},
		},
	}
}

func datasourceVcdEdgeGatewayRead(d *schema.ResourceData, meta interface{}) error {
	return genericVcdEdgeGatewayRead(d, meta, "datasource")
}
