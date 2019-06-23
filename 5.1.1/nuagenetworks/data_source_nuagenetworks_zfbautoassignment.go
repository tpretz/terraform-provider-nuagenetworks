package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.1.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceZFBAutoAssignment() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceZFBAutoAssignmentRead,
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
            "zfb_match_attribute": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zfb_match_attribute_values": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
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
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "associated_enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_enterprise_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceZFBAutoAssignmentRead(d *schema.ResourceData, m interface{}) error {
    filteredZFBAutoAssignments := vspk.ZFBAutoAssignmentsList{}
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
    parent := m.(*vspk.Me)
    filteredZFBAutoAssignments, err = parent.ZFBAutoAssignments(fetchFilter)
    if err != nil {
        return err
    }

    ZFBAutoAssignment := &vspk.ZFBAutoAssignment{}

    if len(filteredZFBAutoAssignments) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredZFBAutoAssignments) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    ZFBAutoAssignment = filteredZFBAutoAssignments[0]

    d.Set("zfb_match_attribute", ZFBAutoAssignment.ZFBMatchAttribute)
    d.Set("zfb_match_attribute_values", ZFBAutoAssignment.ZFBMatchAttributeValues)
    d.Set("name", ZFBAutoAssignment.Name)
    d.Set("last_updated_by", ZFBAutoAssignment.LastUpdatedBy)
    d.Set("description", ZFBAutoAssignment.Description)
    d.Set("entity_scope", ZFBAutoAssignment.EntityScope)
    d.Set("priority", ZFBAutoAssignment.Priority)
    d.Set("associated_enterprise_id", ZFBAutoAssignment.AssociatedEnterpriseID)
    d.Set("associated_enterprise_name", ZFBAutoAssignment.AssociatedEnterpriseName)
    d.Set("external_id", ZFBAutoAssignment.ExternalID)
    
    d.Set("id", ZFBAutoAssignment.Identifier())
    d.Set("parent_id", ZFBAutoAssignment.ParentID)
    d.Set("parent_type", ZFBAutoAssignment.ParentType)
    d.Set("owner", ZFBAutoAssignment.Owner)

    d.SetId(ZFBAutoAssignment.Identifier())
    
    return nil
}