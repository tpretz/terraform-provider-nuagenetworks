package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceBridgeInterface() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceBridgeInterfaceRead,
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
            "vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vport_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "tier_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "policy_decision_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zone_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zone_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_floating_ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "attached_network_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "attached_network_type": &schema.Schema{
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
                ConflictsWith: []string{"parent_l2_domain", "parent_vport"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vport"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain"},
            },
        },
    }
}


func dataSourceBridgeInterfaceRead(d *schema.ResourceData, m interface{}) error {
    filteredBridgeInterfaces := vspk.BridgeInterfacesList{}
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
        filteredBridgeInterfaces, err = parent.BridgeInterfaces(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredBridgeInterfaces, err = parent.BridgeInterfaces(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        filteredBridgeInterfaces, err = parent.BridgeInterfaces(fetchFilter)
        if err != nil {
            return err
        }
    }

    BridgeInterface := &vspk.BridgeInterface{}

    if len(filteredBridgeInterfaces) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredBridgeInterfaces) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    BridgeInterface = filteredBridgeInterfaces[0]

    d.Set("vport_id", BridgeInterface.VPortID)
    d.Set("vport_name", BridgeInterface.VPortName)
    d.Set("name", BridgeInterface.Name)
    d.Set("last_updated_by", BridgeInterface.LastUpdatedBy)
    d.Set("gateway", BridgeInterface.Gateway)
    d.Set("netmask", BridgeInterface.Netmask)
    d.Set("network_name", BridgeInterface.NetworkName)
    d.Set("tier_id", BridgeInterface.TierID)
    d.Set("entity_scope", BridgeInterface.EntityScope)
    d.Set("policy_decision_id", BridgeInterface.PolicyDecisionID)
    d.Set("domain_id", BridgeInterface.DomainID)
    d.Set("domain_name", BridgeInterface.DomainName)
    d.Set("zone_id", BridgeInterface.ZoneID)
    d.Set("zone_name", BridgeInterface.ZoneName)
    d.Set("associated_floating_ip_address", BridgeInterface.AssociatedFloatingIPAddress)
    d.Set("attached_network_id", BridgeInterface.AttachedNetworkID)
    d.Set("attached_network_type", BridgeInterface.AttachedNetworkType)
    d.Set("external_id", BridgeInterface.ExternalID)
    
    d.Set("id", BridgeInterface.Identifier())
    d.Set("parent_id", BridgeInterface.ParentID)
    d.Set("parent_type", BridgeInterface.ParentType)
    d.Set("owner", BridgeInterface.Owner)

    d.SetId(BridgeInterface.Identifier())
    
    return nil
}