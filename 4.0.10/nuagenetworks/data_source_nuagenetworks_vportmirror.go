package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.10"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVPortMirror() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVPortMirrorRead,
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
            "vport_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mirror_destination_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mirror_destination_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mirror_direction": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enterpise_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "attached_network_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_mirror_destination": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vport"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_mirror_destination"},
            },
        },
    }
}


func dataSourceVPortMirrorRead(d *schema.ResourceData, m interface{}) error {
    filteredVPortMirrors := vspk.VPortMirrorsList{}
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
    if attr, ok := d.GetOk("parent_mirror_destination"); ok {
        parent := &vspk.MirrorDestination{ID: attr.(string)}
        filteredVPortMirrors, err = parent.VPortMirrors(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        filteredVPortMirrors, err = parent.VPortMirrors(fetchFilter)
        if err != nil {
            return err
        }
    }

    VPortMirror := &vspk.VPortMirror{}

    if len(filteredVPortMirrors) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVPortMirrors) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VPortMirror = filteredVPortMirrors[0]

    d.Set("vport_name", VPortMirror.VPortName)
    d.Set("last_updated_by", VPortMirror.LastUpdatedBy)
    d.Set("network_name", VPortMirror.NetworkName)
    d.Set("mirror_destination_id", VPortMirror.MirrorDestinationID)
    d.Set("mirror_destination_name", VPortMirror.MirrorDestinationName)
    d.Set("mirror_direction", VPortMirror.MirrorDirection)
    d.Set("enterpise_name", VPortMirror.EnterpiseName)
    d.Set("entity_scope", VPortMirror.EntityScope)
    d.Set("domain_name", VPortMirror.DomainName)
    d.Set("vport_id", VPortMirror.VportId)
    d.Set("attached_network_type", VPortMirror.AttachedNetworkType)
    d.Set("external_id", VPortMirror.ExternalID)
    
    d.Set("id", VPortMirror.Identifier())
    d.Set("parent_id", VPortMirror.ParentID)
    d.Set("parent_type", VPortMirror.ParentType)
    d.Set("owner", VPortMirror.Owner)

    d.SetId(VPortMirror.Identifier())
    
    return nil
}