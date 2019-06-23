package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.2.1.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceTier() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceTierRead,
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
            "packet_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "tier_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "down_threshold_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "probe_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_performance_monitor": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceTierRead(d *schema.ResourceData, m interface{}) error {
    filteredTiers := vspk.TiersList{}
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
    parent := &vspk.PerformanceMonitor{ID: d.Get("parent_performance_monitor").(string)}
    filteredTiers, err = parent.Tiers(fetchFilter)
    if err != nil {
        return err
    }

    Tier := &vspk.Tier{}

    if len(filteredTiers) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredTiers) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Tier = filteredTiers[0]

    d.Set("packet_count", Tier.PacketCount)
    d.Set("last_updated_by", Tier.LastUpdatedBy)
    d.Set("tier_type", Tier.TierType)
    d.Set("timeout", Tier.Timeout)
    d.Set("entity_scope", Tier.EntityScope)
    d.Set("down_threshold_count", Tier.DownThresholdCount)
    d.Set("probe_interval", Tier.ProbeInterval)
    d.Set("external_id", Tier.ExternalID)
    
    d.Set("id", Tier.Identifier())
    d.Set("parent_id", Tier.ParentID)
    d.Set("parent_type", Tier.ParentType)
    d.Set("owner", Tier.Owner)

    d.SetId(Tier.Identifier())
    
    return nil
}