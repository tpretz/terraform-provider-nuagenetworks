package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.9"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceAutoDiscoverHypervisorFromCluster() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceAutoDiscoverHypervisorFromClusterRead,
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
            "network_list": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "assoc_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "hypervisor_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_vcenter_data_center": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_cluster"},
            },
            "parent_vcenter_cluster": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_data_center"},
            },
        },
    }
}


func dataSourceAutoDiscoverHypervisorFromClusterRead(d *schema.ResourceData, m interface{}) error {
    filteredAutoDiscoverHypervisorFromClusters := vspk.AutoDiscoverHypervisorFromClustersList{}
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
    if attr, ok := d.GetOk("parent_vcenter_data_center"); ok {
        parent := &vspk.VCenterDataCenter{ID: attr.(string)}
        filteredAutoDiscoverHypervisorFromClusters, err = parent.AutoDiscoverHypervisorFromClusters(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vcenter_cluster"); ok {
        parent := &vspk.VCenterCluster{ID: attr.(string)}
        filteredAutoDiscoverHypervisorFromClusters, err = parent.AutoDiscoverHypervisorFromClusters(fetchFilter)
        if err != nil {
            return err
        }
    }

    AutoDiscoverHypervisorFromCluster := &vspk.AutoDiscoverHypervisorFromCluster{}

    if len(filteredAutoDiscoverHypervisorFromClusters) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredAutoDiscoverHypervisorFromClusters) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    AutoDiscoverHypervisorFromCluster = filteredAutoDiscoverHypervisorFromClusters[0]

    d.Set("last_updated_by", AutoDiscoverHypervisorFromCluster.LastUpdatedBy)
    d.Set("network_list", AutoDiscoverHypervisorFromCluster.NetworkList)
    d.Set("entity_scope", AutoDiscoverHypervisorFromCluster.EntityScope)
    d.Set("assoc_entity_id", AutoDiscoverHypervisorFromCluster.AssocEntityID)
    d.Set("external_id", AutoDiscoverHypervisorFromCluster.ExternalID)
    d.Set("hypervisor_ip", AutoDiscoverHypervisorFromCluster.HypervisorIP)
    
    d.Set("id", AutoDiscoverHypervisorFromCluster.Identifier())
    d.Set("parent_id", AutoDiscoverHypervisorFromCluster.ParentID)
    d.Set("parent_type", AutoDiscoverHypervisorFromCluster.ParentType)
    d.Set("owner", AutoDiscoverHypervisorFromCluster.Owner)

    d.SetId(AutoDiscoverHypervisorFromCluster.Identifier())
    
    return nil
}