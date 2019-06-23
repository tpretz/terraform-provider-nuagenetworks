package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
)

func resourceEnterpriseSecuredData() *schema.Resource {
    return &schema.Resource{
        Create: resourceEnterpriseSecuredDataCreate,
        Read:   resourceEnterpriseSecuredDataRead,
        Update: resourceEnterpriseSecuredDataUpdate,
        Delete: resourceEnterpriseSecuredDataDelete,
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
            "hash": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "data": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "sek_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "keyserver_cert_serial_number": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "signed_hash": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_enterprise_security": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceEnterpriseSecuredDataCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize EnterpriseSecuredData object
    o := &vspk.EnterpriseSecuredData{
    }
    if attr, ok := d.GetOk("hash"); ok {
        o.Hash = attr.(string)
    }
    if attr, ok := d.GetOk("data"); ok {
        o.Data = attr.(string)
    }
    if attr, ok := d.GetOk("sek_id"); ok {
        o.SekId = attr.(int)
    }
    if attr, ok := d.GetOk("keyserver_cert_serial_number"); ok {
        o.KeyserverCertSerialNumber = attr.(string)
    }
    if attr, ok := d.GetOk("signed_hash"); ok {
        o.SignedHash = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.EnterpriseSecurity{ID: d.Get("parent_enterprise_security").(string)}
    err := parent.CreateEnterpriseSecuredData(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceEnterpriseSecuredDataRead(d, m)
}

func resourceEnterpriseSecuredDataRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.EnterpriseSecuredData{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("hash", o.Hash)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("data", o.Data)
    d.Set("sek_id", o.SekId)
    d.Set("keyserver_cert_serial_number", o.KeyserverCertSerialNumber)
    d.Set("signed_hash", o.SignedHash)
    d.Set("entity_scope", o.EntityScope)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceEnterpriseSecuredDataUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.EnterpriseSecuredData{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("hash"); ok {
        o.Hash = attr.(string)
    }
    if attr, ok := d.GetOk("data"); ok {
        o.Data = attr.(string)
    }
    if attr, ok := d.GetOk("sek_id"); ok {
        o.SekId = attr.(int)
    }
    if attr, ok := d.GetOk("keyserver_cert_serial_number"); ok {
        o.KeyserverCertSerialNumber = attr.(string)
    }
    if attr, ok := d.GetOk("signed_hash"); ok {
        o.SignedHash = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceEnterpriseSecuredDataDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.EnterpriseSecuredData{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}