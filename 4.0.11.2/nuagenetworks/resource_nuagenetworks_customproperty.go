package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11.2"
)

func resourceCustomProperty() *schema.Resource {
    return &schema.Resource{
        Create: resourceCustomPropertyCreate,
        Read:   resourceCustomPropertyRead,
        Update: resourceCustomPropertyUpdate,
        Delete: resourceCustomPropertyDelete,
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
            "attribute_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "attribute_value": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_uplink_connection": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceCustomPropertyCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize CustomProperty object
    o := &vspk.CustomProperty{
    }
    if attr, ok := d.GetOk("attribute_name"); ok {
        o.AttributeName = attr.(string)
    }
    if attr, ok := d.GetOk("attribute_value"); ok {
        o.AttributeValue = attr.(string)
    }
    parent := &vspk.UplinkConnection{ID: d.Get("parent_uplink_connection").(string)}
    err := parent.CreateCustomProperty(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceCustomPropertyRead(d, m)
}

func resourceCustomPropertyRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.CustomProperty{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("attribute_name", o.AttributeName)
    d.Set("attribute_value", o.AttributeValue)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceCustomPropertyUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.CustomProperty{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("attribute_name"); ok {
        o.AttributeName = attr.(string)
    }
    if attr, ok := d.GetOk("attribute_value"); ok {
        o.AttributeValue = attr.(string)
    }

    o.Save()

    return nil
}

func resourceCustomPropertyDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.CustomProperty{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}