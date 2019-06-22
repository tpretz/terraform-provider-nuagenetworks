package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourcePolicyGroupCategory() *schema.Resource {
    return &schema.Resource{
        Read: dataSourcePolicyGroupCategoryRead,
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
            "default_tag": &schema.Schema{
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
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourcePolicyGroupCategoryRead(d *schema.ResourceData, m interface{}) (err error) {
    filteredPolicyGroupCategories := vspk.PolicyGroupCategoriesList{}
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
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    filteredPolicyGroupCategories, err = parent.PolicyGroupCategories(fetchFilter)
    if err != nil {
        return
    }

    PolicyGroupCategory := &vspk.PolicyGroupCategory{}

    if len(filteredPolicyGroupCategories) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredPolicyGroupCategories) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    PolicyGroupCategory = filteredPolicyGroupCategories[0]

    d.Set("name", PolicyGroupCategory.Name)
    d.Set("last_updated_by", PolicyGroupCategory.LastUpdatedBy)
    d.Set("default_tag", PolicyGroupCategory.DefaultTag)
    d.Set("description", PolicyGroupCategory.Description)
    d.Set("entity_scope", PolicyGroupCategory.EntityScope)
    d.Set("external_id", PolicyGroupCategory.ExternalID)
    
    d.Set("id", PolicyGroupCategory.Identifier())
    d.Set("parent_id", PolicyGroupCategory.ParentID)
    d.Set("parent_type", PolicyGroupCategory.ParentType)
    d.Set("owner", PolicyGroupCategory.Owner)

    d.SetId(PolicyGroupCategory.Identifier())
    
    return
}