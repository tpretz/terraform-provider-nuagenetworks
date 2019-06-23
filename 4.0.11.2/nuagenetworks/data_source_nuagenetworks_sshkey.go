package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceSSHKey() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceSSHKeyRead,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "key_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "public_key": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceSSHKeyRead(d *schema.ResourceData, m interface{}) error {
    filteredSSHKeys := vspk.SSHKeysList{}
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

    SSHKey := &vspk.SSHKey{}

    if len(filteredSSHKeys) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredSSHKeys) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    SSHKey = filteredSSHKeys[0]

    d.Set("name", SSHKey.Name)
    d.Set("description", SSHKey.Description)
    d.Set("key_type", SSHKey.KeyType)
    d.Set("public_key", SSHKey.PublicKey)
    
    d.Set("id", SSHKey.Identifier())
    d.Set("parent_id", SSHKey.ParentID)
    d.Set("parent_type", SSHKey.ParentType)
    d.Set("owner", SSHKey.Owner)

    d.SetId(SSHKey.Identifier())
    
    return nil
}