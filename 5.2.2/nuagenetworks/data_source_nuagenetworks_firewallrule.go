package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.2.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceFirewallRule() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceFirewallRuleRead,
        Schema: map[string]*schema.Schema{
            "filter": dataSourceFiltersSchema(),
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "acl_template_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "icmp_code": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "icmp_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ipv6_address_override": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dscp": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "action": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address_override": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dest_network": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dest_pg_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dest_pg_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_ipv6_value": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_value": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mirror_destination_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "flow_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "enterprise_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "location_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "location_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_ipv6_value": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_network": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_pg_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_pg_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_value": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_application_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_application_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associatedfirewall_aclid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "stateful": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "stats_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "stats_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "ether_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_firewall_acl"},
            },
            "parent_firewall_acl": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
        },
    }
}


func dataSourceFirewallRuleRead(d *schema.ResourceData, m interface{}) error {
    filteredFirewallRules := vspk.FirewallRulesList{}
    err := &bambou.Error{}
    fetchFilter := &bambou.FetchingInfo{}
    
    filters, filtersOk := d.GetOk("filter")
    if filtersOk {
        fetchFilter = bambou.NewFetchingInfo()
        for _, v := range filters.(*schema.Set).List() {
            m := v.(map[string]interface{})
            if fetchFilter.Filter != "" {
                fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string),  m["operator"].(string),  m["value"].(string))
            } else {
                fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
            }
           
        }
    }
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredFirewallRules, err = parent.FirewallRules(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_firewall_acl"); ok {
        parent := &vspk.FirewallAcl{ID: attr.(string)}
        filteredFirewallRules, err = parent.FirewallRules(fetchFilter)
        if err != nil {
            return err
        }
    }

    FirewallRule := &vspk.FirewallRule{}

    if len(filteredFirewallRules) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredFirewallRules) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    FirewallRule = filteredFirewallRules[0]

    d.Set("acl_template_name", FirewallRule.ACLTemplateName)
    d.Set("icmp_code", FirewallRule.ICMPCode)
    d.Set("icmp_type", FirewallRule.ICMPType)
    d.Set("ipv6_address_override", FirewallRule.IPv6AddressOverride)
    d.Set("dscp", FirewallRule.DSCP)
    d.Set("action", FirewallRule.Action)
    d.Set("address_override", FirewallRule.AddressOverride)
    d.Set("description", FirewallRule.Description)
    d.Set("dest_network", FirewallRule.DestNetwork)
    d.Set("dest_pg_id", FirewallRule.DestPgId)
    d.Set("dest_pg_type", FirewallRule.DestPgType)
    d.Set("destination_ipv6_value", FirewallRule.DestinationIpv6Value)
    d.Set("destination_port", FirewallRule.DestinationPort)
    d.Set("destination_type", FirewallRule.DestinationType)
    d.Set("destination_value", FirewallRule.DestinationValue)
    d.Set("network_id", FirewallRule.NetworkID)
    d.Set("network_type", FirewallRule.NetworkType)
    d.Set("mirror_destination_id", FirewallRule.MirrorDestinationID)
    d.Set("flow_logging_enabled", FirewallRule.FlowLoggingEnabled)
    d.Set("enterprise_name", FirewallRule.EnterpriseName)
    d.Set("location_id", FirewallRule.LocationID)
    d.Set("location_type", FirewallRule.LocationType)
    d.Set("domain_name", FirewallRule.DomainName)
    d.Set("source_ipv6_value", FirewallRule.SourceIpv6Value)
    d.Set("source_network", FirewallRule.SourceNetwork)
    d.Set("source_pg_id", FirewallRule.SourcePgId)
    d.Set("source_pg_type", FirewallRule.SourcePgType)
    d.Set("source_port", FirewallRule.SourcePort)
    d.Set("source_type", FirewallRule.SourceType)
    d.Set("source_value", FirewallRule.SourceValue)
    d.Set("priority", FirewallRule.Priority)
    d.Set("associated_application_id", FirewallRule.AssociatedApplicationID)
    d.Set("associated_application_object_id", FirewallRule.AssociatedApplicationObjectID)
    d.Set("associatedfirewall_aclid", FirewallRule.AssociatedfirewallACLID)
    d.Set("stateful", FirewallRule.Stateful)
    d.Set("stats_id", FirewallRule.StatsID)
    d.Set("stats_logging_enabled", FirewallRule.StatsLoggingEnabled)
    d.Set("ether_type", FirewallRule.EtherType)
    
    d.Set("id", FirewallRule.Identifier())
    d.Set("parent_id", FirewallRule.ParentID)
    d.Set("parent_type", FirewallRule.ParentType)
    d.Set("owner", FirewallRule.Owner)

    d.SetId(FirewallRule.Identifier())
    
    return nil
}