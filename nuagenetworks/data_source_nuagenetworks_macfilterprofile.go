package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceMACFilterProfile() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceMACFilterProfileRead,
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


func dataSourceMACFilterProfileRead(d *schema.ResourceData, m interface{}) (err error) {
    filteredMACFilterProfiles := vspk.MACFilterProfilesList{}
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
        filteredMACFilterProfiles, err = parent.MACFilterProfiles(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_gateway"); ok {
        parent := &vspk.Gateway{ID: attr.(string)}
        filteredMACFilterProfiles, err = parent.MACFilterProfiles(fetchFilter)
        if err != nil {
            return
        }
    }

    MACFilterProfile := &vspk.MACFilterProfile{}

    if len(filteredMACFilterProfiles) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredMACFilterProfiles) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    MACFilterProfile = filteredMACFilterProfiles[0]

    d.Set("name", MACFilterProfile.Name)
    d.Set("description", MACFilterProfile.Description)
    
    d.Set("id", MACFilterProfile.Identifier())
    d.Set("parent_id", MACFilterProfile.ParentID)
    d.Set("parent_type", MACFilterProfile.ParentType)
    d.Set("owner", MACFilterProfile.Owner)

    d.SetId(MACFilterProfile.Identifier())
    
    return
}