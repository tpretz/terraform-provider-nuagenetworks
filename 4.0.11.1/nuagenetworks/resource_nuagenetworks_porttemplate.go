package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11.1"
)

func resourcePortTemplate() *schema.Resource {
    return &schema.Resource{
        Create: resourcePortTemplateCreate,
        Read:   resourcePortTemplateRead,
        Update: resourcePortTemplateUpdate,
        Delete: resourcePortTemplateDelete,
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
            "vlan_range": &schema.Schema{
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
                Computed: true,
            },
            "physical_name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "port_type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "associated_egress_qos_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_gateway_template": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourcePortTemplateCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize PortTemplate object
    o := &vspk.PortTemplate{
        Name: d.Get("name").(string),
        PhysicalName: d.Get("physical_name").(string),
        PortType: d.Get("port_type").(string),
    }
    if attr, ok := d.GetOk("vlan_range"); ok {
        o.VLANRange = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
        o.AssociatedEgressQOSPolicyID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.GatewayTemplate{ID: d.Get("parent_gateway_template").(string)}
    err := parent.CreatePortTemplate(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourcePortTemplateRead(d, m)
}

func resourcePortTemplateRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PortTemplate{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("vlan_range", o.VLANRange)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("description", o.Description)
    d.Set("physical_name", o.PhysicalName)
    d.Set("entity_scope", o.EntityScope)
    d.Set("port_type", o.PortType)
    d.Set("associated_egress_qos_policy_id", o.AssociatedEgressQOSPolicyID)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourcePortTemplateUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PortTemplate{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.PhysicalName = d.Get("physical_name").(string)
    o.PortType = d.Get("port_type").(string)
    
    if attr, ok := d.GetOk("vlan_range"); ok {
        o.VLANRange = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
        o.AssociatedEgressQOSPolicyID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourcePortTemplateDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PortTemplate{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}