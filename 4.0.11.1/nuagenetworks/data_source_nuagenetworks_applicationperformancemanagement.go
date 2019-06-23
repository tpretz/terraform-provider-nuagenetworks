package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceApplicationperformancemanagement() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceApplicationperformancemanagementRead,
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
            "read_only": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "app_group_unique_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_performance_monitor_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_performance_monitor": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise", "parent_vport"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_performance_monitor", "parent_vport"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_performance_monitor", "parent_enterprise"},
            },
        },
    }
}


func dataSourceApplicationperformancemanagementRead(d *schema.ResourceData, m interface{}) error {
    filteredApplicationperformancemanagements := vspk.ApplicationperformancemanagementsList{}
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
    if attr, ok := d.GetOk("parent_performance_monitor"); ok {
        parent := &vspk.PerformanceMonitor{ID: attr.(string)}
        filteredApplicationperformancemanagements, err = parent.Applicationperformancemanagements(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredApplicationperformancemanagements, err = parent.Applicationperformancemanagements(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        filteredApplicationperformancemanagements, err = parent.Applicationperformancemanagements(fetchFilter)
        if err != nil {
            return err
        }
    }

    Applicationperformancemanagement := &vspk.Applicationperformancemanagement{}

    if len(filteredApplicationperformancemanagements) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredApplicationperformancemanagements) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Applicationperformancemanagement = filteredApplicationperformancemanagements[0]

    d.Set("name", Applicationperformancemanagement.Name)
    d.Set("read_only", Applicationperformancemanagement.ReadOnly)
    d.Set("description", Applicationperformancemanagement.Description)
    d.Set("app_group_unique_id", Applicationperformancemanagement.AppGroupUniqueId)
    d.Set("associated_performance_monitor_id", Applicationperformancemanagement.AssociatedPerformanceMonitorID)
    
    d.Set("id", Applicationperformancemanagement.Identifier())
    d.Set("parent_id", Applicationperformancemanagement.ParentID)
    d.Set("parent_type", Applicationperformancemanagement.ParentType)
    d.Set("owner", Applicationperformancemanagement.Owner)

    d.SetId(Applicationperformancemanagement.Identifier())
    
    return nil
}