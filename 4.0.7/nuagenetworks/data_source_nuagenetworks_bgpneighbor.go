package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.7"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceBGPNeighbor() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceBGPNeighborRead,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dampening_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "peer_as": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "peer_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "session": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_export_routing_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_import_routing_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan"},
            },
            "parent_vlan": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_subnet"},
            },
        },
    }
}


func dataSourceBGPNeighborRead(d *schema.ResourceData, m interface{}) error {
    filteredBGPNeighbors := vspk.BGPNeighborsList{}
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
    if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        filteredBGPNeighbors, err = parent.BGPNeighbors(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vlan"); ok {
        parent := &vspk.VLAN{ID: attr.(string)}
        filteredBGPNeighbors, err = parent.BGPNeighbors(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredBGPNeighbors, err = parent.BGPNeighbors(fetchFilter)
        if err != nil {
            return err
        }
    }

    BGPNeighbor := &vspk.BGPNeighbor{}

    if len(filteredBGPNeighbors) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredBGPNeighbors) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    BGPNeighbor = filteredBGPNeighbors[0]

    d.Set("name", BGPNeighbor.Name)
    d.Set("dampening_enabled", BGPNeighbor.DampeningEnabled)
    d.Set("peer_as", BGPNeighbor.PeerAS)
    d.Set("peer_ip", BGPNeighbor.PeerIP)
    d.Set("description", BGPNeighbor.Description)
    d.Set("session", BGPNeighbor.Session)
    d.Set("entity_scope", BGPNeighbor.EntityScope)
    d.Set("associated_export_routing_policy_id", BGPNeighbor.AssociatedExportRoutingPolicyID)
    d.Set("associated_import_routing_policy_id", BGPNeighbor.AssociatedImportRoutingPolicyID)
    d.Set("external_id", BGPNeighbor.ExternalID)
    
    d.Set("id", BGPNeighbor.Identifier())
    d.Set("parent_id", BGPNeighbor.ParentID)
    d.Set("parent_type", BGPNeighbor.ParentType)
    d.Set("owner", BGPNeighbor.Owner)

    d.SetId(BGPNeighbor.Identifier())
    
    return nil
}