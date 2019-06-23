package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11"
)

func resourceApplicationBinding() *schema.Resource {
    return &schema.Resource{
        Create: resourceApplicationBindingCreate,
        Read:   resourceApplicationBindingRead,
        Update: resourceApplicationBindingUpdate,
        Delete: resourceApplicationBindingDelete,
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
            "read_only": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "associated_application_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "parent_applicationperformancemanagement": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceApplicationBindingCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize ApplicationBinding object
    o := &vspk.ApplicationBinding{
        AssociatedApplicationID: d.Get("associated_application_id").(string),
    }
    if attr, ok := d.GetOk("read_only"); ok {
        o.ReadOnly = attr.(bool)
    }
    parent := &vspk.Applicationperformancemanagement{ID: d.Get("parent_applicationperformancemanagement").(string)}
    err := parent.CreateApplicationBinding(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceApplicationBindingRead(d, m)
}

func resourceApplicationBindingRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ApplicationBinding{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("read_only", o.ReadOnly)
    d.Set("priority", o.Priority)
    d.Set("associated_application_id", o.AssociatedApplicationID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceApplicationBindingUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ApplicationBinding{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.AssociatedApplicationID = d.Get("associated_application_id").(string)
    
    if attr, ok := d.GetOk("read_only"); ok {
        o.ReadOnly = attr.(bool)
    }

    o.Save()

    return nil
}

func resourceApplicationBindingDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ApplicationBinding{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}