package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.2.1"
)

func resourceIPReservation() *schema.Resource {
    return &schema.Resource{
        Create: resourceIPReservationCreate,
        Read:   resourceIPReservationRead,
        Update: resourceIPReservationUpdate,
        Delete: resourceIPReservationDelete,
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
            "mac": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
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
            "dynamic_allocation_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceIPReservationCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize IPReservation object
    o := &vspk.IPReservation{
        MAC: d.Get("mac").(string),
        IPAddress: d.Get("ip_address").(string),
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("dynamic_allocation_enabled"); ok {
        o.DynamicAllocationEnabled = attr.(bool)
    }
    parent := &vspk.Subnet{ID: d.Get("parent_subnet").(string)}
    err := parent.CreateIPReservation(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceIPReservationRead(d, m)
}

func resourceIPReservationRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IPReservation{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("mac", o.MAC)
    d.Set("ip_address", o.IPAddress)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("entity_scope", o.EntityScope)
    d.Set("external_id", o.ExternalID)
    d.Set("dynamic_allocation_enabled", o.DynamicAllocationEnabled)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceIPReservationUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IPReservation{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.MAC = d.Get("mac").(string)
    o.IPAddress = d.Get("ip_address").(string)
    
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("dynamic_allocation_enabled"); ok {
        o.DynamicAllocationEnabled = attr.(bool)
    }

    o.Save()

    return nil
}

func resourceIPReservationDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IPReservation{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}