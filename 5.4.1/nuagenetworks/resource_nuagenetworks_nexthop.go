package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.4.1"
)

func resourceNextHop() *schema.Resource {
    return &schema.Resource{
        Create: resourceNextHopCreate,
        Read:   resourceNextHopRead,
        Update: resourceNextHopUpdate,
        Delete: resourceNextHopDelete,
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
            "ip_type": &schema.Schema{
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
            "route_distinguisher": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ip": &schema.Schema{
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
            "parent_link": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceNextHopCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize NextHop object
    o := &vspk.NextHop{
    }
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("ip"); ok {
        o.Ip = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("type"); ok {
        o.Type = attr.(string)
    }
    parent := &vspk.Link{ID: d.Get("parent_link").(string)}
    err := parent.CreateNextHop(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceNextHopRead(d, m)
}

func resourceNextHopRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NextHop{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("ip_type", o.IPType)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("entity_scope", o.EntityScope)
    d.Set("route_distinguisher", o.RouteDistinguisher)
    d.Set("ip", o.Ip)
    d.Set("external_id", o.ExternalID)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceNextHopUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NextHop{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("ip"); ok {
        o.Ip = attr.(string)
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

func resourceNextHopDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NextHop{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}