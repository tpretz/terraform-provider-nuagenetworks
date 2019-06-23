package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.1"
)

func resourceIKEPSK() *schema.Resource {
    return &schema.Resource{
        Create: resourceIKEPSKCreate,
        Read:   resourceIKEPSKRead,
        Update: resourceIKEPSKUpdate,
        Delete: resourceIKEPSKDelete,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "signature": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "signing_certificate_serial_number": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "encrypted_psk": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "encrypting_certificate_serial_number": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "unencrypted_psk": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "auto_created": &schema.Schema{
                Type:     schema.TypeBool,
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

func resourceIKEPSKCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize IKEPSK object
    o := &vspk.IKEPSK{
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("signature"); ok {
        o.Signature = attr.(string)
    }
    if attr, ok := d.GetOk("signing_certificate_serial_number"); ok {
        o.SigningCertificateSerialNumber = attr.(int)
    }
    if attr, ok := d.GetOk("encrypted_psk"); ok {
        o.EncryptedPSK = attr.(string)
    }
    if attr, ok := d.GetOk("encrypting_certificate_serial_number"); ok {
        o.EncryptingCertificateSerialNumber = attr.(int)
    }
    if attr, ok := d.GetOk("unencrypted_psk"); ok {
        o.UnencryptedPSK = attr.(string)
    }
    if attr, ok := d.GetOk("associated_enterprise_id"); ok {
        o.AssociatedEnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("auto_created"); ok {
        o.AutoCreated = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateIKEPSK(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceIKEPSKRead(d, m)
}

func resourceIKEPSKRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IKEPSK{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("description", o.Description)
    d.Set("signature", o.Signature)
    d.Set("signing_certificate_serial_number", o.SigningCertificateSerialNumber)
    d.Set("encrypted_psk", o.EncryptedPSK)
    d.Set("encrypting_certificate_serial_number", o.EncryptingCertificateSerialNumber)
    d.Set("unencrypted_psk", o.UnencryptedPSK)
    d.Set("entity_scope", o.EntityScope)
    d.Set("associated_enterprise_id", o.AssociatedEnterpriseID)
    d.Set("auto_created", o.AutoCreated)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceIKEPSKUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IKEPSK{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("signature"); ok {
        o.Signature = attr.(string)
    }
    if attr, ok := d.GetOk("signing_certificate_serial_number"); ok {
        o.SigningCertificateSerialNumber = attr.(int)
    }
    if attr, ok := d.GetOk("encrypted_psk"); ok {
        o.EncryptedPSK = attr.(string)
    }
    if attr, ok := d.GetOk("encrypting_certificate_serial_number"); ok {
        o.EncryptingCertificateSerialNumber = attr.(int)
    }
    if attr, ok := d.GetOk("unencrypted_psk"); ok {
        o.UnencryptedPSK = attr.(string)
    }
    if attr, ok := d.GetOk("associated_enterprise_id"); ok {
        o.AssociatedEnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("auto_created"); ok {
        o.AutoCreated = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceIKEPSKDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IKEPSK{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}