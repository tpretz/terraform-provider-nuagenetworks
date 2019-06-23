package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.1.1"
)

func resourceCommand() *schema.Resource {
    return &schema.Resource{
        Create: resourceCommandCreate,
        Read:   resourceCommandRead,
        Update: resourceCommandUpdate,
        Delete: resourceCommandDelete,
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
            "detailed_status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "detailed_status_code": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "command": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "command_information": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "summary": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceCommandCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize Command object
    o := &vspk.Command{
        Command: d.Get("command").(string),
        Summary: d.Get("summary").(string),
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.NSGateway{ID: d.Get("parent_ns_gateway").(string)}
    err := parent.CreateCommand(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceCommandRead(d, m)
}

func resourceCommandRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Command{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("detailed_status", o.DetailedStatus)
    d.Set("detailed_status_code", o.DetailedStatusCode)
    d.Set("entity_scope", o.EntityScope)
    d.Set("command", o.Command)
    d.Set("command_information", o.CommandInformation)
    d.Set("status", o.Status)
    d.Set("summary", o.Summary)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceCommandUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Command{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Command = d.Get("command").(string)
    o.Summary = d.Get("summary").(string)
    
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceCommandDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.Command{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}