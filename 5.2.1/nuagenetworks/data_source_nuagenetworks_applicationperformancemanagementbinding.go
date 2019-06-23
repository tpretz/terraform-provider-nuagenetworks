package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.2.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceApplicationperformancemanagementbinding() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceApplicationperformancemanagementbindingRead,
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
            "read_only": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "associated_application_performance_management_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceApplicationperformancemanagementbindingRead(d *schema.ResourceData, m interface{}) error {
    filteredApplicationperformancemanagementbindings := vspk.ApplicationperformancemanagementbindingsList{}
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
    parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
    filteredApplicationperformancemanagementbindings, err = parent.Applicationperformancemanagementbindings(fetchFilter)
    if err != nil {
        return err
    }

    Applicationperformancemanagementbinding := &vspk.Applicationperformancemanagementbinding{}

    if len(filteredApplicationperformancemanagementbindings) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredApplicationperformancemanagementbindings) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Applicationperformancemanagementbinding = filteredApplicationperformancemanagementbindings[0]

    d.Set("read_only", Applicationperformancemanagementbinding.ReadOnly)
    d.Set("priority", Applicationperformancemanagementbinding.Priority)
    d.Set("associated_application_performance_management_id", Applicationperformancemanagementbinding.AssociatedApplicationPerformanceManagementID)
    
    d.Set("id", Applicationperformancemanagementbinding.Identifier())
    d.Set("parent_id", Applicationperformancemanagementbinding.ParentID)
    d.Set("parent_type", Applicationperformancemanagementbinding.ParentType)
    d.Set("owner", Applicationperformancemanagementbinding.Owner)

    d.SetId(Applicationperformancemanagementbinding.Identifier())
    
    return nil
}