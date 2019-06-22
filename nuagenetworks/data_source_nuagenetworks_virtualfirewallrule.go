package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVirtualFirewallRule() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVirtualFirewallRuleRead,
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
            "last_updated_by": &schema.Schema{
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
            "web_filter_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "web_filter_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_port": &schema.Schema{
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
            "entity_scope": &schema.Schema{
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
            "policy_state": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "protocol": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_egress_entry_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ingress_entry_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_l7_application_signature_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_live_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_live_template_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_traffic_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_traffic_type_id": &schema.Schema{
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
            "overlay_mirror_destination_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_virtual_firewall_policy"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_virtual_firewall_policy"},
            },
            "parent_virtual_firewall_policy": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain"},
            },
        },
    }
}


func dataSourceVirtualFirewallRuleRead(d *schema.ResourceData, m interface{}) error {
    filteredVirtualFirewallRules := vspk.VirtualFirewallRulesList{}
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
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        filteredVirtualFirewallRules, err = parent.VirtualFirewallRules(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredVirtualFirewallRules, err = parent.VirtualFirewallRules(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_virtual_firewall_policy"); ok {
        parent := &vspk.VirtualFirewallPolicy{ID: attr.(string)}
        filteredVirtualFirewallRules, err = parent.VirtualFirewallRules(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredVirtualFirewallRules, err = parent.VirtualFirewallRules(fetchFilter)
        if err != nil {
            return err
        }
    }

    VirtualFirewallRule := &vspk.VirtualFirewallRule{}

    if len(filteredVirtualFirewallRules) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVirtualFirewallRules) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VirtualFirewallRule = filteredVirtualFirewallRules[0]

    d.Set("acl_template_name", VirtualFirewallRule.ACLTemplateName)
    d.Set("icmp_code", VirtualFirewallRule.ICMPCode)
    d.Set("icmp_type", VirtualFirewallRule.ICMPType)
    d.Set("ipv6_address_override", VirtualFirewallRule.IPv6AddressOverride)
    d.Set("dscp", VirtualFirewallRule.DSCP)
    d.Set("last_updated_by", VirtualFirewallRule.LastUpdatedBy)
    d.Set("action", VirtualFirewallRule.Action)
    d.Set("address_override", VirtualFirewallRule.AddressOverride)
    d.Set("web_filter_id", VirtualFirewallRule.WebFilterID)
    d.Set("web_filter_type", VirtualFirewallRule.WebFilterType)
    d.Set("description", VirtualFirewallRule.Description)
    d.Set("destination_port", VirtualFirewallRule.DestinationPort)
    d.Set("network_id", VirtualFirewallRule.NetworkID)
    d.Set("network_type", VirtualFirewallRule.NetworkType)
    d.Set("mirror_destination_id", VirtualFirewallRule.MirrorDestinationID)
    d.Set("flow_logging_enabled", VirtualFirewallRule.FlowLoggingEnabled)
    d.Set("enterprise_name", VirtualFirewallRule.EnterpriseName)
    d.Set("entity_scope", VirtualFirewallRule.EntityScope)
    d.Set("location_id", VirtualFirewallRule.LocationID)
    d.Set("location_type", VirtualFirewallRule.LocationType)
    d.Set("policy_state", VirtualFirewallRule.PolicyState)
    d.Set("domain_name", VirtualFirewallRule.DomainName)
    d.Set("source_port", VirtualFirewallRule.SourcePort)
    d.Set("priority", VirtualFirewallRule.Priority)
    d.Set("protocol", VirtualFirewallRule.Protocol)
    d.Set("associated_egress_entry_id", VirtualFirewallRule.AssociatedEgressEntryID)
    d.Set("associated_ingress_entry_id", VirtualFirewallRule.AssociatedIngressEntryID)
    d.Set("associated_l7_application_signature_id", VirtualFirewallRule.AssociatedL7ApplicationSignatureID)
    d.Set("associated_live_entity_id", VirtualFirewallRule.AssociatedLiveEntityID)
    d.Set("associated_live_template_id", VirtualFirewallRule.AssociatedLiveTemplateID)
    d.Set("associated_traffic_type", VirtualFirewallRule.AssociatedTrafficType)
    d.Set("associated_traffic_type_id", VirtualFirewallRule.AssociatedTrafficTypeID)
    d.Set("stateful", VirtualFirewallRule.Stateful)
    d.Set("stats_id", VirtualFirewallRule.StatsID)
    d.Set("stats_logging_enabled", VirtualFirewallRule.StatsLoggingEnabled)
    d.Set("ether_type", VirtualFirewallRule.EtherType)
    d.Set("overlay_mirror_destination_id", VirtualFirewallRule.OverlayMirrorDestinationID)
    d.Set("external_id", VirtualFirewallRule.ExternalID)
    
    d.Set("id", VirtualFirewallRule.Identifier())
    d.Set("parent_id", VirtualFirewallRule.ParentID)
    d.Set("parent_type", VirtualFirewallRule.ParentType)
    d.Set("owner", VirtualFirewallRule.Owner)

    d.SetId(VirtualFirewallRule.Identifier())
    
    return nil
}