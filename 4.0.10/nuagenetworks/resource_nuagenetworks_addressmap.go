package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.10"
)

func resourceAddressMap() *schema.Resource {
    return &schema.Resource{
        Create: resourceAddressMapCreate,
        Read:   resourceAddressMapRead,
        Update: resourceAddressMapUpdate,
        Delete: resourceAddressMapDelete,
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
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "private_ip": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "private_port": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "associated_patnat_pool_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "public_ip": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "public_port": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_patnat_pool": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceAddressMapCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize AddressMap object
    o := &vspk.AddressMap{
        PrivateIP: d.Get("private_ip").(string),
        AssociatedPATNATPoolID: d.Get("associated_patnat_pool_id").(string),
        PublicIP: d.Get("public_ip").(string),
    }
    if attr, ok := d.GetOk("private_port"); ok {
        o.PrivatePort = attr.(int)
    }
    if attr, ok := d.GetOk("public_port"); ok {
        o.PublicPort = attr.(int)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }
    parent := &vspk.PATNATPool{ID: d.Get("parent_patnat_pool").(string)}
    err := parent.CreateAddressMap(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceAddressMapRead(d, m)
}

func resourceAddressMapRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.AddressMap{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("entity_scope", o.EntityScope)
    d.Set("private_ip", o.PrivateIP)
    d.Set("private_port", o.PrivatePort)
    d.Set("associated_patnat_pool_id", o.AssociatedPATNATPoolID)
    d.Set("public_ip", o.PublicIP)
    d.Set("public_port", o.PublicPort)
    d.Set("external_id", o.ExternalID)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceAddressMapUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.AddressMap{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.PrivateIP = d.Get("private_ip").(string)
    o.AssociatedPATNATPoolID = d.Get("associated_patnat_pool_id").(string)
    o.PublicIP = d.Get("public_ip").(string)
    
    if attr, ok := d.GetOk("private_port"); ok {
        o.PrivatePort = attr.(int)
    }
    if attr, ok := d.GetOk("public_port"); ok {
        o.PublicPort = attr.(int)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }

    o.Save()

    return nil
}

func resourceAddressMapDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.AddressMap{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}