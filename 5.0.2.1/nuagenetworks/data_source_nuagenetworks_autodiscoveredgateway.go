package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.0.2.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceAutoDiscoveredGateway() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceAutoDiscoveredGatewayRead,
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
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "peer": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "personality": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "controllers": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "use_gateway_vlanvnid": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "vtep": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "system_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceAutoDiscoveredGatewayRead(d *schema.ResourceData, m interface{}) error {
    filteredAutoDiscoveredGateways := vspk.AutoDiscoveredGatewaysList{}
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
    parent := m.(*vspk.Me)
    filteredAutoDiscoveredGateways, err = parent.AutoDiscoveredGateways(fetchFilter)
    if err != nil {
        return err
    }

    AutoDiscoveredGateway := &vspk.AutoDiscoveredGateway{}

    if len(filteredAutoDiscoveredGateways) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredAutoDiscoveredGateways) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    AutoDiscoveredGateway = filteredAutoDiscoveredGateways[0]

    d.Set("name", AutoDiscoveredGateway.Name)
    d.Set("last_updated_by", AutoDiscoveredGateway.LastUpdatedBy)
    d.Set("gateway_id", AutoDiscoveredGateway.GatewayID)
    d.Set("peer", AutoDiscoveredGateway.Peer)
    d.Set("personality", AutoDiscoveredGateway.Personality)
    d.Set("description", AutoDiscoveredGateway.Description)
    d.Set("entity_scope", AutoDiscoveredGateway.EntityScope)
    d.Set("controllers", AutoDiscoveredGateway.Controllers)
    d.Set("use_gateway_vlanvnid", AutoDiscoveredGateway.UseGatewayVLANVNID)
    d.Set("vtep", AutoDiscoveredGateway.Vtep)
    d.Set("external_id", AutoDiscoveredGateway.ExternalID)
    d.Set("system_id", AutoDiscoveredGateway.SystemID)
    
    d.Set("id", AutoDiscoveredGateway.Identifier())
    d.Set("parent_id", AutoDiscoveredGateway.ParentID)
    d.Set("parent_type", AutoDiscoveredGateway.ParentType)
    d.Set("owner", AutoDiscoveredGateway.Owner)

    d.SetId(AutoDiscoveredGateway.Identifier())
    
    return nil
}