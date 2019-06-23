package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.1.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceOverlayMirrorDestinationTemplate() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceOverlayMirrorDestinationTemplateRead,
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
            "redundancy_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "end_point_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "trigger_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_l2_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceOverlayMirrorDestinationTemplateRead(d *schema.ResourceData, m interface{}) error {
    filteredOverlayMirrorDestinationTemplates := vspk.OverlayMirrorDestinationTemplatesList{}
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
    parent := &vspk.L2DomainTemplate{ID: d.Get("parent_l2_domain_template").(string)}
    filteredOverlayMirrorDestinationTemplates, err = parent.OverlayMirrorDestinationTemplates(fetchFilter)
    if err != nil {
        return err
    }

    OverlayMirrorDestinationTemplate := &vspk.OverlayMirrorDestinationTemplate{}

    if len(filteredOverlayMirrorDestinationTemplates) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredOverlayMirrorDestinationTemplates) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    OverlayMirrorDestinationTemplate = filteredOverlayMirrorDestinationTemplates[0]

    d.Set("name", OverlayMirrorDestinationTemplate.Name)
    d.Set("last_updated_by", OverlayMirrorDestinationTemplate.LastUpdatedBy)
    d.Set("redundancy_enabled", OverlayMirrorDestinationTemplate.RedundancyEnabled)
    d.Set("description", OverlayMirrorDestinationTemplate.Description)
    d.Set("end_point_type", OverlayMirrorDestinationTemplate.EndPointType)
    d.Set("entity_scope", OverlayMirrorDestinationTemplate.EntityScope)
    d.Set("trigger_type", OverlayMirrorDestinationTemplate.TriggerType)
    d.Set("external_id", OverlayMirrorDestinationTemplate.ExternalID)
    
    d.Set("id", OverlayMirrorDestinationTemplate.Identifier())
    d.Set("parent_id", OverlayMirrorDestinationTemplate.ParentID)
    d.Set("parent_type", OverlayMirrorDestinationTemplate.ParentType)
    d.Set("owner", OverlayMirrorDestinationTemplate.Owner)

    d.SetId(OverlayMirrorDestinationTemplate.Identifier())
    
    return nil
}