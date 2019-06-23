package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.2.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceGatewayTemplate() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceGatewayTemplateRead,
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
            "personality": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}


func dataSourceGatewayTemplateRead(d *schema.ResourceData, m interface{}) error {
    filteredGatewayTemplates := vspk.GatewayTemplatesList{}
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
        filteredGatewayTemplates, err = parent.GatewayTemplates(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredGatewayTemplates, err = parent.GatewayTemplates(fetchFilter)
        if err != nil {
            return err
        }
    }

    GatewayTemplate := &vspk.GatewayTemplate{}

    if len(filteredGatewayTemplates) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredGatewayTemplates) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    GatewayTemplate = filteredGatewayTemplates[0]

    d.Set("name", GatewayTemplate.Name)
    d.Set("last_updated_by", GatewayTemplate.LastUpdatedBy)
    d.Set("personality", GatewayTemplate.Personality)
    d.Set("description", GatewayTemplate.Description)
    d.Set("enterprise_id", GatewayTemplate.EnterpriseID)
    d.Set("entity_scope", GatewayTemplate.EntityScope)
    d.Set("external_id", GatewayTemplate.ExternalID)
    
    d.Set("id", GatewayTemplate.Identifier())
    d.Set("parent_id", GatewayTemplate.ParentID)
    d.Set("parent_type", GatewayTemplate.ParentType)
    d.Set("owner", GatewayTemplate.Owner)

    d.SetId(GatewayTemplate.Identifier())
    
    return nil
}