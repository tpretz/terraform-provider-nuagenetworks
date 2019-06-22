package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceIngressAdvFwdTemplate() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceIngressAdvFwdTemplateRead,
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
            "active": &schema.Schema{
                Type:     schema.TypeBool,
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
            "policy_state": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "priority_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_live_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "auto_generate_priority": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_l2_domain_template", "parent_domain_template"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain_template", "parent_domain_template"},
            },
            "parent_l2_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_domain_template"},
            },
            "parent_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_l2_domain_template"},
            },
        },
    }
}


func dataSourceIngressAdvFwdTemplateRead(d *schema.ResourceData, m interface{}) (err error) {
    filteredIngressAdvFwdTemplates := vspk.IngressAdvFwdTemplatesList{}
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
        filteredIngressAdvFwdTemplates, err = parent.IngressAdvFwdTemplates(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredIngressAdvFwdTemplates, err = parent.IngressAdvFwdTemplates(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
        parent := &vspk.L2DomainTemplate{ID: attr.(string)}
        filteredIngressAdvFwdTemplates, err = parent.IngressAdvFwdTemplates(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_domain_template"); ok {
        parent := &vspk.DomainTemplate{ID: attr.(string)}
        filteredIngressAdvFwdTemplates, err = parent.IngressAdvFwdTemplates(fetchFilter)
        if err != nil {
            return
        }
    }

    IngressAdvFwdTemplate := &vspk.IngressAdvFwdTemplate{}

    if len(filteredIngressAdvFwdTemplates) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredIngressAdvFwdTemplates) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    IngressAdvFwdTemplate = filteredIngressAdvFwdTemplates[0]

    d.Set("name", IngressAdvFwdTemplate.Name)
    d.Set("last_updated_by", IngressAdvFwdTemplate.LastUpdatedBy)
    d.Set("active", IngressAdvFwdTemplate.Active)
    d.Set("description", IngressAdvFwdTemplate.Description)
    d.Set("entity_scope", IngressAdvFwdTemplate.EntityScope)
    d.Set("policy_state", IngressAdvFwdTemplate.PolicyState)
    d.Set("priority", IngressAdvFwdTemplate.Priority)
    d.Set("priority_type", IngressAdvFwdTemplate.PriorityType)
    d.Set("associated_live_entity_id", IngressAdvFwdTemplate.AssociatedLiveEntityID)
    d.Set("auto_generate_priority", IngressAdvFwdTemplate.AutoGeneratePriority)
    d.Set("external_id", IngressAdvFwdTemplate.ExternalID)
    
    d.Set("id", IngressAdvFwdTemplate.Identifier())
    d.Set("parent_id", IngressAdvFwdTemplate.ParentID)
    d.Set("parent_type", IngressAdvFwdTemplate.ParentType)
    d.Set("owner", IngressAdvFwdTemplate.Owner)

    d.SetId(IngressAdvFwdTemplate.Identifier())
    
    return
}