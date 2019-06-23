package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.1.1"
)

func resourceAvatar() *schema.Resource {
    return &schema.Resource{
        Create: resourceAvatarCreate,
        Read:   resourceAvatarRead,
        Update: resourceAvatarUpdate,
        Delete: resourceAvatarDelete,
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
            "parent_user": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_user"},
            },
        },
    }
}

func resourceAvatarCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Avatar object
    o := &vspk.Avatar{
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }
    if attr, ok := d.GetOk("parent_user"); ok {
        parent := &vspk.User{ID: attr.(string)}
        err := parent.CreateAvatar(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        err := parent.CreateAvatar(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceAvatarRead(d, m)
}

func resourceAvatarRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Avatar{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("entity_scope", o.EntityScope)
    d.Set("external_id", o.ExternalID)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceAvatarUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Avatar{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
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

func resourceAvatarDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Avatar{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}