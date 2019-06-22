package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/nuagenetworks/vspk-go/vspk"
)

func resourceNSGUpgradeProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceNSGUpgradeProfileCreate,
        Read:   resourceNSGUpgradeProfileRead,
        Update: resourceNSGUpgradeProfileUpdate,
        Delete: resourceNSGUpgradeProfileDelete,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "metadata_upgrade_path": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "download_rate_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Default: 10000,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}

func resourceNSGUpgradeProfileCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize NSGUpgradeProfile object
    o := &vspk.NSGUpgradeProfile{
        Name: d.Get("name").(string),
        MetadataUpgradePath: d.Get("metadata_upgrade_path").(string),
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("enterprise_id"); ok {
        o.EnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("download_rate_limit"); ok {
        o.DownloadRateLimit = attr.(int)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := m.(*vspk.Me)
    err := parent.CreateNSGUpgradeProfile(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceNSGUpgradeProfileRead(d, m)
}

func resourceNSGUpgradeProfileRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NSGUpgradeProfile{
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
    d.Set("metadata_upgrade_path", o.MetadataUpgradePath)
    d.Set("enterprise_id", o.EnterpriseID)
    d.Set("entity_scope", o.EntityScope)
    d.Set("download_rate_limit", o.DownloadRateLimit)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceNSGUpgradeProfileUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NSGUpgradeProfile{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.MetadataUpgradePath = d.Get("metadata_upgrade_path").(string)
    
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("enterprise_id"); ok {
        o.EnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("download_rate_limit"); ok {
        o.DownloadRateLimit = attr.(int)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceNSGUpgradeProfileDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NSGUpgradeProfile{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}