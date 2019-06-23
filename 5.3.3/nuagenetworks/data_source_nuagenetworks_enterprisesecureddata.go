package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceEnterpriseSecuredData() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceEnterpriseSecuredDataRead,
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
            "hash": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "data": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "seed_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "sek_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "keyserver_cert_serial_number": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "signed_hash": &schema.Schema{
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
            "parent_enterprise_security": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceEnterpriseSecuredDataRead(d *schema.ResourceData, m interface{}) error {
    filteredEnterpriseSecuredDatas := vspk.EnterpriseSecuredDatasList{}
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
    parent := &vspk.EnterpriseSecurity{ID: d.Get("parent_enterprise_security").(string)}
    filteredEnterpriseSecuredDatas, err = parent.EnterpriseSecuredDatas(fetchFilter)
    if err != nil {
        return err
    }

    EnterpriseSecuredData := &vspk.EnterpriseSecuredData{}

    if len(filteredEnterpriseSecuredDatas) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredEnterpriseSecuredDatas) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    EnterpriseSecuredData = filteredEnterpriseSecuredDatas[0]

    d.Set("hash", EnterpriseSecuredData.Hash)
    d.Set("last_updated_by", EnterpriseSecuredData.LastUpdatedBy)
    d.Set("data", EnterpriseSecuredData.Data)
    d.Set("seed_type", EnterpriseSecuredData.SeedType)
    d.Set("sek_id", EnterpriseSecuredData.SekId)
    d.Set("keyserver_cert_serial_number", EnterpriseSecuredData.KeyserverCertSerialNumber)
    d.Set("signed_hash", EnterpriseSecuredData.SignedHash)
    d.Set("entity_scope", EnterpriseSecuredData.EntityScope)
    d.Set("external_id", EnterpriseSecuredData.ExternalID)
    
    d.Set("id", EnterpriseSecuredData.Identifier())
    d.Set("parent_id", EnterpriseSecuredData.ParentID)
    d.Set("parent_type", EnterpriseSecuredData.ParentType)
    d.Set("owner", EnterpriseSecuredData.Owner)

    d.SetId(EnterpriseSecuredData.Identifier())
    
    return nil
}