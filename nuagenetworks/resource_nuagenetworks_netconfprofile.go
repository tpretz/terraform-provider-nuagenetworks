package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
)

func resourceNetconfProfile() *schema.Resource {
    return &schema.Resource{
        Create: resourceNetconfProfileCreate,
        Read:   resourceNetconfProfileRead,
        Update: resourceNetconfProfileUpdate,
        Delete: resourceNetconfProfileDelete,
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
            },
            "password": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "port": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Default: 830,
            },
            "user_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}

func resourceNetconfProfileCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize NetconfProfile object
    o := &vspk.NetconfProfile{
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("password"); ok {
        o.Password = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("port"); ok {
        o.Port = attr.(int)
    }
    if attr, ok := d.GetOk("user_name"); ok {
        o.UserName = attr.(string)
    }
    if attr, ok := d.GetOk("parent_me"); ok {
        parent := &vspk.Me{ID: attr.(string)}
        err := parent.CreateNetconfProfile(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        err := parent.CreateNetconfProfile(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceNetconfProfileRead(d, m)
}

func resourceNetconfProfileRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetconfProfile{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("password", o.Password)
    d.Set("description", o.Description)
    d.Set("port", o.Port)
    d.Set("user_name", o.UserName)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceNetconfProfileUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetconfProfile{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("password"); ok {
        o.Password = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("port"); ok {
        o.Port = attr.(int)
    }
    if attr, ok := d.GetOk("user_name"); ok {
        o.UserName = attr.(string)
    }

    o.Save()

    return nil
}

func resourceNetconfProfileDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetconfProfile{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}