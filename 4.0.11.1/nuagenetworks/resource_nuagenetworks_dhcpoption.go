package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11.1"
)

func resourceDHCPOption() *schema.Resource {
    return &schema.Resource{
        Create: resourceDHCPOptionCreate,
        Read:   resourceDHCPOptionRead,
        Update: resourceDHCPOptionUpdate,
        Delete: resourceDHCPOptionDelete,
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
            "value": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "actual_type": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "actual_values": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "length": &schema.Schema{
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
                Required: true,
            },
            "parent_shared_network_resource": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_zone", "parent_vport"},
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_shared_network_resource", "parent_l2_domain", "parent_subnet", "parent_zone", "parent_vport"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_shared_network_resource", "parent_domain", "parent_subnet", "parent_zone", "parent_vport"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_shared_network_resource", "parent_domain", "parent_l2_domain", "parent_zone", "parent_vport"},
            },
            "parent_zone": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_shared_network_resource", "parent_domain", "parent_l2_domain", "parent_subnet", "parent_vport"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_shared_network_resource", "parent_domain", "parent_l2_domain", "parent_subnet", "parent_zone"},
            },
        },
    }
}

func resourceDHCPOptionCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize DHCPOption object
    o := &vspk.DHCPOption{
        Type: d.Get("type").(string),
    }
    if attr, ok := d.GetOk("value"); ok {
        o.Value = attr.(string)
    }
    if attr, ok := d.GetOk("actual_type"); ok {
        o.ActualType = attr.(int)
    }
    if attr, ok := d.GetOk("actual_values"); ok {
        o.ActualValues = attr.([]interface{})
    }
    if attr, ok := d.GetOk("length"); ok {
        o.Length = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_shared_network_resource"); ok {
        parent := &vspk.SharedNetworkResource{ID: attr.(string)}
        err := parent.CreateDHCPOption(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        err := parent.CreateDHCPOption(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        err := parent.CreateDHCPOption(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        err := parent.CreateDHCPOption(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_zone"); ok {
        parent := &vspk.Zone{ID: attr.(string)}
        err := parent.CreateDHCPOption(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        err := parent.CreateDHCPOption(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceDHCPOptionRead(d, m)
}

func resourceDHCPOptionRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DHCPOption{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("value", o.Value)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("actual_type", o.ActualType)
    d.Set("actual_values", o.ActualValues)
    d.Set("length", o.Length)
    d.Set("entity_scope", o.EntityScope)
    d.Set("external_id", o.ExternalID)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceDHCPOptionUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DHCPOption{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Type = d.Get("type").(string)
    
    if attr, ok := d.GetOk("value"); ok {
        o.Value = attr.(string)
    }
    if attr, ok := d.GetOk("actual_type"); ok {
        o.ActualType = attr.(int)
    }
    if attr, ok := d.GetOk("actual_values"); ok {
        o.ActualValues = attr.([]interface{})
    }
    if attr, ok := d.GetOk("length"); ok {
        o.Length = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceDHCPOptionDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.DHCPOption{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}