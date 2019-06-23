package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceRateLimiter() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceRateLimiterRead,
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
            "peak_burst_size": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "peak_information_rate": &schema.Schema{
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
            "committed_information_rate": &schema.Schema{
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


func dataSourceRateLimiterRead(d *schema.ResourceData, m interface{}) error {
    filteredRateLimiters := vspk.RateLimitersList{}
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
        filteredRateLimiters, err = parent.RateLimiters(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredRateLimiters, err = parent.RateLimiters(fetchFilter)
        if err != nil {
            return err
        }
    }

    RateLimiter := &vspk.RateLimiter{}

    if len(filteredRateLimiters) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredRateLimiters) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    RateLimiter = filteredRateLimiters[0]

    d.Set("name", RateLimiter.Name)
    d.Set("last_updated_by", RateLimiter.LastUpdatedBy)
    d.Set("peak_burst_size", RateLimiter.PeakBurstSize)
    d.Set("peak_information_rate", RateLimiter.PeakInformationRate)
    d.Set("description", RateLimiter.Description)
    d.Set("entity_scope", RateLimiter.EntityScope)
    d.Set("committed_information_rate", RateLimiter.CommittedInformationRate)
    d.Set("external_id", RateLimiter.ExternalID)
    
    d.Set("id", RateLimiter.Identifier())
    d.Set("parent_id", RateLimiter.ParentID)
    d.Set("parent_type", RateLimiter.ParentType)
    d.Set("owner", RateLimiter.Owner)

    d.SetId(RateLimiter.Identifier())
    
    return nil
}