package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.1.2"
)

func resourceVMResync() *schema.Resource {
    return &schema.Resource{
        Create: resourceVMResyncCreate,
        Read:   resourceVMResyncRead,
        Update: resourceVMResyncUpdate,
        Delete: resourceVMResyncDelete,
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
            "last_request_timestamp": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "last_time_resync_initiated": &schema.Schema{
                Type:     schema.TypeInt,
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
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vm"},
            },
            "parent_vm": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_subnet"},
            },
        },
    }
}

func resourceVMResyncCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VMResync object
    o := &vspk.VMResync{
    }
    if attr, ok := d.GetOk("last_request_timestamp"); ok {
        o.LastRequestTimestamp = attr.(int)
    }
    if attr, ok := d.GetOk("last_time_resync_initiated"); ok {
        o.LastTimeResyncInitiated = attr.(int)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        err := parent.CreateVMResync(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vm"); ok {
        parent := &vspk.VM{ID: attr.(string)}
        err := parent.CreateVMResync(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceVMResyncRead(d, m)
}

func resourceVMResyncRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VMResync{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("last_request_timestamp", o.LastRequestTimestamp)
    d.Set("last_time_resync_initiated", o.LastTimeResyncInitiated)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("entity_scope", o.EntityScope)
    d.Set("status", o.Status)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVMResyncUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VMResync{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("last_request_timestamp"); ok {
        o.LastRequestTimestamp = attr.(int)
    }
    if attr, ok := d.GetOk("last_time_resync_initiated"); ok {
        o.LastTimeResyncInitiated = attr.(int)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVMResyncDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VMResync{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}