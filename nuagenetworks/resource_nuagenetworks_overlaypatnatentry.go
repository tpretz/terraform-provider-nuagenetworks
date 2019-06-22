package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
)

func resourceOverlayPATNATEntry() *schema.Resource {
    return &schema.Resource{
        Create: resourceOverlayPATNATEntryCreate,
        Read:   resourceOverlayPATNATEntryRead,
        Update: resourceOverlayPATNATEntryUpdate,
        Delete: resourceOverlayPATNATEntryDelete,
        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },
        Schema: map[string]*schema.Schema{
            "id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
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
            "nat_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: true,
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
                Optional: true,
            },
            "associated_domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_link_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "public_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_overlay_address_pool": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceOverlayPATNATEntryCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize OverlayPATNATEntry object
    o := &vspk.OverlayPATNATEntry{
    }
    if attr, ok := d.GetOk("nat_enabled"); ok {
        o.NATEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("private_ip"); ok {
        o.PrivateIP = attr.(string)
    }
    if attr, ok := d.GetOk("associated_domain_id"); ok {
        o.AssociatedDomainID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_link_id"); ok {
        o.AssociatedLinkID = attr.(string)
    }
    if attr, ok := d.GetOk("public_ip"); ok {
        o.PublicIP = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.OverlayAddressPool{ID: d.Get("parent_overlay_address_pool").(string)}
    err := parent.CreateOverlayPATNATEntry(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceOverlayPATNATEntryRead(d, m)
}

func resourceOverlayPATNATEntryRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayPATNATEntry{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("nat_enabled", o.NATEnabled)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("entity_scope", o.EntityScope)
    d.Set("private_ip", o.PrivateIP)
    d.Set("associated_domain_id", o.AssociatedDomainID)
    d.Set("associated_link_id", o.AssociatedLinkID)
    d.Set("public_ip", o.PublicIP)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceOverlayPATNATEntryUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayPATNATEntry{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("nat_enabled"); ok {
        o.NATEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("private_ip"); ok {
        o.PrivateIP = attr.(string)
    }
    if attr, ok := d.GetOk("associated_domain_id"); ok {
        o.AssociatedDomainID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_link_id"); ok {
        o.AssociatedLinkID = attr.(string)
    }
    if attr, ok := d.GetOk("public_ip"); ok {
        o.PublicIP = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceOverlayPATNATEntryDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.OverlayPATNATEntry{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}