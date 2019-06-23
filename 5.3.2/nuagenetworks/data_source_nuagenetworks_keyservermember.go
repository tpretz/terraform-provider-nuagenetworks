package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceKeyServerMember() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceKeyServerMemberRead,
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
            "pem_encoded": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "certificate_serial_number": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "fqdn": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "issuer_dn": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "subject_dn": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "public_key": &schema.Schema{
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


func dataSourceKeyServerMemberRead(d *schema.ResourceData, m interface{}) error {
    filteredKeyServerMembers := vspk.KeyServerMembersList{}
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
    filteredKeyServerMembers, err = parent.KeyServerMembers(fetchFilter)
    if err != nil {
        return err
    }

    KeyServerMember := &vspk.KeyServerMember{}

    if len(filteredKeyServerMembers) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredKeyServerMembers) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    KeyServerMember = filteredKeyServerMembers[0]

    d.Set("last_updated_by", KeyServerMember.LastUpdatedBy)
    d.Set("pem_encoded", KeyServerMember.PemEncoded)
    d.Set("certificate_serial_number", KeyServerMember.CertificateSerialNumber)
    d.Set("entity_scope", KeyServerMember.EntityScope)
    d.Set("fqdn", KeyServerMember.Fqdn)
    d.Set("issuer_dn", KeyServerMember.IssuerDN)
    d.Set("subject_dn", KeyServerMember.SubjectDN)
    d.Set("public_key", KeyServerMember.PublicKey)
    d.Set("external_id", KeyServerMember.ExternalID)
    
    d.Set("id", KeyServerMember.Identifier())
    d.Set("parent_id", KeyServerMember.ParentID)
    d.Set("parent_type", KeyServerMember.ParentType)
    d.Set("owner", KeyServerMember.Owner)

    d.SetId(KeyServerMember.Identifier())
    
    return nil
}