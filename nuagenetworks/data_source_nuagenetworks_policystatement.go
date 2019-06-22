package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourcePolicyStatement() *schema.Resource {
    return &schema.Resource{
        Read: dataSourcePolicyStatementRead,
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
            "description": &schema.Schema{
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
            "parent_link": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourcePolicyStatementRead(d *schema.ResourceData, m interface{}) error {
    filteredPolicyStatements := vspk.PolicyStatementsList{}
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
    parent := &vspk.Link{ID: d.Get("parent_link").(string)}
    filteredPolicyStatements, err = parent.PolicyStatements(fetchFilter)
    if err != nil {
        return err
    }

    PolicyStatement := &vspk.PolicyStatement{}

    if len(filteredPolicyStatements) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredPolicyStatements) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    PolicyStatement = filteredPolicyStatements[0]

    d.Set("name", PolicyStatement.Name)
    d.Set("last_updated_by", PolicyStatement.LastUpdatedBy)
    d.Set("description", PolicyStatement.Description)
    d.Set("entity_scope", PolicyStatement.EntityScope)
    d.Set("external_id", PolicyStatement.ExternalID)
    
    d.Set("id", PolicyStatement.Identifier())
    d.Set("parent_id", PolicyStatement.ParentID)
    d.Set("parent_type", PolicyStatement.ParentType)
    d.Set("owner", PolicyStatement.Owner)

    d.SetId(PolicyStatement.Identifier())
    
    return nil
}