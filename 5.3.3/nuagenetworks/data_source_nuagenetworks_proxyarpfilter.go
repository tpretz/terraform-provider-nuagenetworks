package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceProxyARPFilter() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceProxyARPFilterRead,
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
            "ip_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "max_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "min_address": &schema.Schema{
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
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_subnet"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain"},
            },
        },
    }
}


func dataSourceProxyARPFilterRead(d *schema.ResourceData, m interface{}) error {
    filteredProxyARPFilters := vspk.ProxyARPFiltersList{}
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
    if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredProxyARPFilters, err = parent.ProxyARPFilters(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        filteredProxyARPFilters, err = parent.ProxyARPFilters(fetchFilter)
        if err != nil {
            return err
        }
    }

    ProxyARPFilter := &vspk.ProxyARPFilter{}

    if len(filteredProxyARPFilters) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredProxyARPFilters) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    ProxyARPFilter = filteredProxyARPFilters[0]

    d.Set("ip_type", ProxyARPFilter.IPType)
    d.Set("last_updated_by", ProxyARPFilter.LastUpdatedBy)
    d.Set("max_address", ProxyARPFilter.MaxAddress)
    d.Set("min_address", ProxyARPFilter.MinAddress)
    d.Set("entity_scope", ProxyARPFilter.EntityScope)
    d.Set("external_id", ProxyARPFilter.ExternalID)
    
    d.Set("id", ProxyARPFilter.Identifier())
    d.Set("parent_id", ProxyARPFilter.ParentID)
    d.Set("parent_type", ProxyARPFilter.ParentType)
    d.Set("owner", ProxyARPFilter.Owner)

    d.SetId(ProxyARPFilter.Identifier())
    
    return nil
}