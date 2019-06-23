package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceIngressExternalServiceTemplate() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceIngressExternalServiceTemplateRead,
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


func dataSourceIngressExternalServiceTemplateRead(d *schema.ResourceData, m interface{}) error {
    filteredIngressExternalServiceTemplates := vspk.IngressExternalServiceTemplatesList{}
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
        filteredIngressExternalServiceTemplates, err = parent.IngressExternalServiceTemplates(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredIngressExternalServiceTemplates, err = parent.IngressExternalServiceTemplates(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
        parent := &vspk.L2DomainTemplate{ID: attr.(string)}
        filteredIngressExternalServiceTemplates, err = parent.IngressExternalServiceTemplates(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_domain_template"); ok {
        parent := &vspk.DomainTemplate{ID: attr.(string)}
        filteredIngressExternalServiceTemplates, err = parent.IngressExternalServiceTemplates(fetchFilter)
        if err != nil {
            return err
        }
    }

    IngressExternalServiceTemplate := &vspk.IngressExternalServiceTemplate{}

    if len(filteredIngressExternalServiceTemplates) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredIngressExternalServiceTemplates) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    IngressExternalServiceTemplate = filteredIngressExternalServiceTemplates[0]

    d.Set("name", IngressExternalServiceTemplate.Name)
    d.Set("active", IngressExternalServiceTemplate.Active)
    d.Set("description", IngressExternalServiceTemplate.Description)
    d.Set("entity_scope", IngressExternalServiceTemplate.EntityScope)
    d.Set("policy_state", IngressExternalServiceTemplate.PolicyState)
    d.Set("priority", IngressExternalServiceTemplate.Priority)
    d.Set("priority_type", IngressExternalServiceTemplate.PriorityType)
    d.Set("associated_live_entity_id", IngressExternalServiceTemplate.AssociatedLiveEntityID)
    d.Set("external_id", IngressExternalServiceTemplate.ExternalID)
    
    d.Set("id", IngressExternalServiceTemplate.Identifier())
    d.Set("parent_id", IngressExternalServiceTemplate.ParentID)
    d.Set("parent_type", IngressExternalServiceTemplate.ParentType)
    d.Set("owner", IngressExternalServiceTemplate.Owner)

    d.SetId(IngressExternalServiceTemplate.Identifier())
    
    return nil
}