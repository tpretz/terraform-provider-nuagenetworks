package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.0.2.1"
)

func resourceVLAN() *schema.Resource {
    return &schema.Resource{
        Create: resourceVLANCreate,
        Read:   resourceVLANRead,
        Update: resourceVLANUpdate,
        Delete: resourceVLANDelete,
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
                Type:     schema.TypeInt,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "readonly": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "template_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "restricted": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "use_user_mnemonic": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "user_mnemonic": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_bgp_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_egress_qos_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_uplink_connection_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_vsc_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "duc_vlan": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
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
            "parent_ns_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vsg_redundant_port", "parent_redundant_port", "parent_port"},
            },
            "parent_vsg_redundant_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_ns_port", "parent_redundant_port", "parent_port"},
            },
            "parent_redundant_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_ns_port", "parent_vsg_redundant_port", "parent_port"},
            },
            "parent_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_ns_port", "parent_vsg_redundant_port", "parent_redundant_port"},
            },
        },
    }
}

func resourceVLANCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VLAN object
    o := &vspk.VLAN{
        Value: d.Get("value").(int),
    }
    if attr, ok := d.GetOk("gateway_id"); ok {
        o.GatewayID = attr.(string)
    }
    if attr, ok := d.GetOk("readonly"); ok {
        o.Readonly = attr.(bool)
    }
    if attr, ok := d.GetOk("template_id"); ok {
        o.TemplateID = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("restricted"); ok {
        o.Restricted = attr.(bool)
    }
    if attr, ok := d.GetOk("vport_id"); ok {
        o.VportID = attr.(string)
    }
    if attr, ok := d.GetOk("use_user_mnemonic"); ok {
        o.UseUserMnemonic = attr.(bool)
    }
    if attr, ok := d.GetOk("user_mnemonic"); ok {
        o.UserMnemonic = attr.(string)
    }
    if attr, ok := d.GetOk("associated_bgp_profile_id"); ok {
        o.AssociatedBGPProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
        o.AssociatedEgressQOSPolicyID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_uplink_connection_id"); ok {
        o.AssociatedUplinkConnectionID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_vsc_profile_id"); ok {
        o.AssociatedVSCProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("duc_vlan"); ok {
        o.DucVlan = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_ns_port"); ok {
        parent := &vspk.NSPort{ID: attr.(string)}
        err := parent.CreateVLAN(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vsg_redundant_port"); ok {
        parent := &vspk.VsgRedundantPort{ID: attr.(string)}
        err := parent.CreateVLAN(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_redundant_port"); ok {
        parent := &vspk.RedundantPort{ID: attr.(string)}
        err := parent.CreateVLAN(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_port"); ok {
        parent := &vspk.Port{ID: attr.(string)}
        err := parent.CreateVLAN(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    if attr, ok := d.GetOk("patnatpools"); ok {
        o.AssignPATNATPools(attr.(vspk.PATNATPoolsList))
    }
    return resourceVLANRead(d, m)
}

func resourceVLANRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VLAN{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("value", o.Value)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway_id", o.GatewayID)
    d.Set("readonly", o.Readonly)
    d.Set("template_id", o.TemplateID)
    d.Set("permitted_action", o.PermittedAction)
    d.Set("description", o.Description)
    d.Set("restricted", o.Restricted)
    d.Set("entity_scope", o.EntityScope)
    d.Set("vport_id", o.VportID)
    d.Set("use_user_mnemonic", o.UseUserMnemonic)
    d.Set("user_mnemonic", o.UserMnemonic)
    d.Set("associated_bgp_profile_id", o.AssociatedBGPProfileID)
    d.Set("associated_egress_qos_policy_id", o.AssociatedEgressQOSPolicyID)
    d.Set("associated_uplink_connection_id", o.AssociatedUplinkConnectionID)
    d.Set("associated_vsc_profile_id", o.AssociatedVSCProfileID)
    d.Set("status", o.Status)
    d.Set("duc_vlan", o.DucVlan)
    d.Set("external_id", o.ExternalID)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVLANUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VLAN{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Value = d.Get("value").(int)
    
    if attr, ok := d.GetOk("gateway_id"); ok {
        o.GatewayID = attr.(string)
    }
    if attr, ok := d.GetOk("readonly"); ok {
        o.Readonly = attr.(bool)
    }
    if attr, ok := d.GetOk("template_id"); ok {
        o.TemplateID = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("restricted"); ok {
        o.Restricted = attr.(bool)
    }
    if attr, ok := d.GetOk("vport_id"); ok {
        o.VportID = attr.(string)
    }
    if attr, ok := d.GetOk("use_user_mnemonic"); ok {
        o.UseUserMnemonic = attr.(bool)
    }
    if attr, ok := d.GetOk("user_mnemonic"); ok {
        o.UserMnemonic = attr.(string)
    }
    if attr, ok := d.GetOk("associated_bgp_profile_id"); ok {
        o.AssociatedBGPProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
        o.AssociatedEgressQOSPolicyID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_uplink_connection_id"); ok {
        o.AssociatedUplinkConnectionID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_vsc_profile_id"); ok {
        o.AssociatedVSCProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("duc_vlan"); ok {
        o.DucVlan = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVLANDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VLAN{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}