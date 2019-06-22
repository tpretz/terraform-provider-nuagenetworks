package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
)

func resourceNetworkMacroGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceNetworkMacroGroupCreate,
        Read:   resourceNetworkMacroGroupRead,
        Update: resourceNetworkMacroGroupUpdate,
        Delete: resourceNetworkMacroGroupDelete,
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
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "network_macros": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "is_saa_s_type": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceNetworkMacroGroupCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize NetworkMacroGroup object
    o := &vspk.NetworkMacroGroup{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("network_macros"); ok {
        o.NetworkMacros = attr.([]interface{})
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateNetworkMacroGroup(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    if attr, ok := d.GetOk("enterprisenetworks"); ok {
        o.AssignEnterpriseNetworks(attr.(vspk.EnterpriseNetworksList))
    }
    return resourceNetworkMacroGroupRead(d, m)
}

func resourceNetworkMacroGroupRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetworkMacroGroup{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("description", o.Description)
    d.Set("network_macros", o.NetworkMacros)
    d.Set("entity_scope", o.EntityScope)
    d.Set("is_saa_s_type", o.IsSaaSType)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceNetworkMacroGroupUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetworkMacroGroup{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("network_macros"); ok {
        o.NetworkMacros = attr.([]interface{})
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceNetworkMacroGroupDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NetworkMacroGroup{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}