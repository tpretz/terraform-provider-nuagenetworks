package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.6.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceIKECertificate() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceIKECertificateRead,
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
            "pem_encoded": &schema.Schema{
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
            "serial_number": &schema.Schema{
                Type:     schema.TypeInt,
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
            "not_after": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "not_before": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "associated_enterprise_id": &schema.Schema{
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
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceIKECertificateRead(d *schema.ResourceData, m interface{}) error {
    filteredIKECertificates := vspk.IKECertificatesList{}
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
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    filteredIKECertificates, err = parent.IKECertificates(fetchFilter)
    if err != nil {
        return err
    }

    IKECertificate := &vspk.IKECertificate{}

    if len(filteredIKECertificates) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredIKECertificates) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    IKECertificate = filteredIKECertificates[0]

    d.Set("pem_encoded", IKECertificate.PEMEncoded)
    d.Set("name", IKECertificate.Name)
    d.Set("last_updated_by", IKECertificate.LastUpdatedBy)
    d.Set("serial_number", IKECertificate.SerialNumber)
    d.Set("description", IKECertificate.Description)
    d.Set("entity_scope", IKECertificate.EntityScope)
    d.Set("not_after", IKECertificate.NotAfter)
    d.Set("not_before", IKECertificate.NotBefore)
    d.Set("associated_enterprise_id", IKECertificate.AssociatedEnterpriseID)
    d.Set("issuer_dn", IKECertificate.IssuerDN)
    d.Set("subject_dn", IKECertificate.SubjectDN)
    d.Set("external_id", IKECertificate.ExternalID)
    
    d.Set("id", IKECertificate.Identifier())
    d.Set("parent_id", IKECertificate.ParentID)
    d.Set("parent_type", IKECertificate.ParentType)
    d.Set("owner", IKECertificate.Owner)

    d.SetId(IKECertificate.Identifier())
    
    return nil
}