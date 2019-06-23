package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.10"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNSPortTemplate() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNSPortTemplateRead,
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
            "vlan_range": &schema.Schema{
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "physical_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "infrastructure_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "port_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "speed": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_egress_qos_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mtu": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_ns_gateway_template": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceNSPortTemplateRead(d *schema.ResourceData, m interface{}) error {
    filteredNSPortTemplates := vspk.NSPortTemplatesList{}
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
    parent := &vspk.NSGatewayTemplate{ID: d.Get("parent_ns_gateway_template").(string)}
    filteredNSPortTemplates, err = parent.NSPortTemplates(fetchFilter)
    if err != nil {
        return err
    }

    NSPortTemplate := &vspk.NSPortTemplate{}

    if len(filteredNSPortTemplates) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNSPortTemplates) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NSPortTemplate = filteredNSPortTemplates[0]

    d.Set("vlan_range", NSPortTemplate.VLANRange)
    d.Set("name", NSPortTemplate.Name)
    d.Set("last_updated_by", NSPortTemplate.LastUpdatedBy)
    d.Set("description", NSPortTemplate.Description)
    d.Set("physical_name", NSPortTemplate.PhysicalName)
    d.Set("infrastructure_profile_id", NSPortTemplate.InfrastructureProfileID)
    d.Set("entity_scope", NSPortTemplate.EntityScope)
    d.Set("port_type", NSPortTemplate.PortType)
    d.Set("speed", NSPortTemplate.Speed)
    d.Set("associated_egress_qos_policy_id", NSPortTemplate.AssociatedEgressQOSPolicyID)
    d.Set("mtu", NSPortTemplate.Mtu)
    d.Set("external_id", NSPortTemplate.ExternalID)
    
    d.Set("id", NSPortTemplate.Identifier())
    d.Set("parent_id", NSPortTemplate.ParentID)
    d.Set("parent_type", NSPortTemplate.ParentType)
    d.Set("owner", NSPortTemplate.Owner)

    d.SetId(NSPortTemplate.Identifier())
    
    return nil
}