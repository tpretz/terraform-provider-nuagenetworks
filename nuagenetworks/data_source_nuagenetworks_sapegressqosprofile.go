package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceSAPEgressQoSProfile() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceSAPEgressQoSProfileRead,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_redundancy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_gateway"},
            },
            "parent_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_redundancy_group"},
            },
        },
    }
}


func dataSourceSAPEgressQoSProfileRead(d *schema.ResourceData, m interface{}) (err error) {
    filteredSAPEgressQoSProfiles := vspk.SAPEgressQoSProfilesList{}
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
    if attr, ok := d.GetOk("parent_redundancy_group"); ok {
        parent := &vspk.RedundancyGroup{ID: attr.(string)}
        filteredSAPEgressQoSProfiles, err = parent.SAPEgressQoSProfiles(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_gateway"); ok {
        parent := &vspk.Gateway{ID: attr.(string)}
        filteredSAPEgressQoSProfiles, err = parent.SAPEgressQoSProfiles(fetchFilter)
        if err != nil {
            return
        }
    }

    SAPEgressQoSProfile := &vspk.SAPEgressQoSProfile{}

    if len(filteredSAPEgressQoSProfiles) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredSAPEgressQoSProfiles) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    SAPEgressQoSProfile = filteredSAPEgressQoSProfiles[0]

    d.Set("name", SAPEgressQoSProfile.Name)
    d.Set("description", SAPEgressQoSProfile.Description)
    
    d.Set("id", SAPEgressQoSProfile.Identifier())
    d.Set("parent_id", SAPEgressQoSProfile.ParentID)
    d.Set("parent_type", SAPEgressQoSProfile.ParentType)
    d.Set("owner", SAPEgressQoSProfile.Owner)

    d.SetId(SAPEgressQoSProfile.Identifier())
    
    return
}