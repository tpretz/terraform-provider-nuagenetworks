package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
)

func resourceNetconfSession() *schema.Resource {
    return &schema.Resource{
        Create: resourceNetconfSessionCreate,
        Read:   resourceNetconfSessionRead,
        Update: resourceNetconfSessionUpdate,
        Delete: resourceNetconfSessionDelete,
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
            "associated_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_gateway_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_netconf_manager": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceNetconfSessionCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize NetconfSession object
    o := &vspk.NetconfSession{
    }
    if attr, ok := d.GetOk("associated_gateway_id"); ok {
        o.AssociatedGatewayID = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    parent := &vspk.NetconfManager{ID: d.Get("parent_netconf_manager").(string)}
    err := parent.CreateNetconfSession(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceNetconfSessionRead(d, m)
}

func resourceNetconfSessionRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetconfSession{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("associated_gateway_id", o.AssociatedGatewayID)
    d.Set("associated_gateway_name", o.AssociatedGatewayName)
    d.Set("status", o.Status)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceNetconfSessionUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetconfSession{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("associated_gateway_id"); ok {
        o.AssociatedGatewayID = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }

    o.Save()

    return nil
}

func resourceNetconfSessionDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetconfSession{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}