package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourceIKECertificate() *schema.Resource {
    return &schema.Resource{
        Create: resourceIKECertificateCreate,
        Read:   resourceIKECertificateRead,
        Update: resourceIKECertificateUpdate,
        Delete: resourceIKECertificateDelete,
        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },
        Schema: map[string]*schema.Schema{
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "pem_encoded": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "serial_number": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "not_after": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
                Computed: true,
            },
            "not_before": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
                Computed: true,
            },
            "associated_enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "issuer_dn": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "subject_dn": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceIKECertificateCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize IKECertificate object
    o := &vspk.IKECertificate{
    }
    if attr, ok := d.GetOk("pem_encoded"); ok {
        o.PEMEncoded = attr.(string)
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_enterprise_id"); ok {
        o.AssociatedEnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateIKECertificate(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceIKECertificateRead(d, m)
}

func resourceIKECertificateRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IKECertificate{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("pem_encoded", o.PEMEncoded)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("serial_number", o.SerialNumber)
    d.Set("description", o.Description)
    d.Set("entity_scope", o.EntityScope)
    d.Set("not_after", o.NotAfter)
    d.Set("not_before", o.NotBefore)
    d.Set("associated_enterprise_id", o.AssociatedEnterpriseID)
    d.Set("issuer_dn", o.IssuerDN)
    d.Set("subject_dn", o.SubjectDN)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceIKECertificateUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IKECertificate{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("pem_encoded"); ok {
        o.PEMEncoded = attr.(string)
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_enterprise_id"); ok {
        o.AssociatedEnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceIKECertificateDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IKECertificate{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}