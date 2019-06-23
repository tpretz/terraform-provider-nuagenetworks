package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
    "github.com/tpretz/go-bambou/bambou"
)

func dataSourceNATMapEntry() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNATMapEntryRead,
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
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "private_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_patnat_pool_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "public_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_patnat_pool": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceNATMapEntryRead(d *schema.ResourceData, m interface{}) error {
    filteredNATMapEntries := vspk.NATMapEntriesList{}
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
    parent := &vspk.PATNATPool{ID: d.Get("parent_patnat_pool").(string)}
    filteredNATMapEntries, err = parent.NATMapEntries(fetchFilter)
    if err != nil {
        return err
    }

    NATMapEntry := &vspk.NATMapEntry{}

    if len(filteredNATMapEntries) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNATMapEntries) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NATMapEntry = filteredNATMapEntries[0]

    d.Set("last_updated_by", NATMapEntry.LastUpdatedBy)
    d.Set("entity_scope", NATMapEntry.EntityScope)
    d.Set("private_ip", NATMapEntry.PrivateIP)
    d.Set("associated_patnat_pool_id", NATMapEntry.AssociatedPATNATPoolID)
    d.Set("public_ip", NATMapEntry.PublicIP)
    d.Set("external_id", NATMapEntry.ExternalID)
    d.Set("type", NATMapEntry.Type)
    
    d.Set("id", NATMapEntry.Identifier())
    d.Set("parent_id", NATMapEntry.ParentID)
    d.Set("parent_type", NATMapEntry.ParentType)
    d.Set("owner", NATMapEntry.Owner)

    d.SetId(NATMapEntry.Identifier())
    
    return nil
}