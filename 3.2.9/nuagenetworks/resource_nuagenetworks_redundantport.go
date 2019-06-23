package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/3.2.9"
)

func resourceRedundantPort() *schema.Resource {
    return &schema.Resource{
        Create: resourceRedundantPortCreate,
        Read:   resourceRedundantPortRead,
        Update: resourceRedundantPortUpdate,
        Delete: resourceRedundantPortDelete,
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
            "physical_name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "infrastructure_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "port_peer1_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "port_peer2_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "port_type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "use_untagged_heartbeat_vlan": &schema.Schema{
                Type:     schema.TypeBool,
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
            "associated_egress_qos_policy_id": &schema.Schema{
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
            "parent_ns_redundant_gateway_group": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceRedundantPortCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize RedundantPort object
    o := &vspk.RedundantPort{
        Name: d.Get("name").(string),
        PhysicalName: d.Get("physical_name").(string),
        PortType: d.Get("port_type").(string),
    }
    if attr, ok := d.GetOk("vlan_range"); ok {
        o.VLANRange = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("infrastructure_profile_id"); ok {
        o.InfrastructureProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("port_peer1_id"); ok {
        o.PortPeer1ID = attr.(string)
    }
    if attr, ok := d.GetOk("port_peer2_id"); ok {
        o.PortPeer2ID = attr.(string)
    }
    if attr, ok := d.GetOk("use_untagged_heartbeat_vlan"); ok {
        o.UseUntaggedHeartbeatVlan = attr.(bool)
    }
    if attr, ok := d.GetOk("use_user_mnemonic"); ok {
        o.UseUserMnemonic = attr.(bool)
    }
    if attr, ok := d.GetOk("user_mnemonic"); ok {
        o.UserMnemonic = attr.(string)
    }
    if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
        o.AssociatedEgressQOSPolicyID = attr.(string)
    }
    if attr, ok := d.GetOk("status"); ok {
        o.Status = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.NSRedundantGatewayGroup{ID: d.Get("parent_ns_redundant_gateway_group").(string)}
    err := parent.CreateRedundantPort(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceRedundantPortRead(d, m)
}

func resourceRedundantPortRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.RedundantPort{
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
    d.Set("permitted_action", o.PermittedAction)
    d.Set("description", o.Description)
    d.Set("physical_name", o.PhysicalName)
    d.Set("infrastructure_profile_id", o.InfrastructureProfileID)
    d.Set("entity_scope", o.EntityScope)
    d.Set("port_peer1_id", o.PortPeer1ID)
    d.Set("port_peer2_id", o.PortPeer2ID)
    d.Set("port_type", o.PortType)
    d.Set("use_untagged_heartbeat_vlan", o.UseUntaggedHeartbeatVlan)
    d.Set("use_user_mnemonic", o.UseUserMnemonic)
    d.Set("user_mnemonic", o.UserMnemonic)
    d.Set("associated_egress_qos_policy_id", o.AssociatedEgressQOSPolicyID)
    d.Set("status", o.Status)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceRedundantPortUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.RedundantPort{
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
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("infrastructure_profile_id"); ok {
        o.InfrastructureProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("port_peer1_id"); ok {
        o.PortPeer1ID = attr.(string)
    }
    if attr, ok := d.GetOk("port_peer2_id"); ok {
        o.PortPeer2ID = attr.(string)
    }
    if attr, ok := d.GetOk("use_untagged_heartbeat_vlan"); ok {
        o.UseUntaggedHeartbeatVlan = attr.(bool)
    }
    if attr, ok := d.GetOk("use_user_mnemonic"); ok {
        o.UseUserMnemonic = attr.(bool)
    }
    if attr, ok := d.GetOk("user_mnemonic"); ok {
        o.UserMnemonic = attr.(string)
    }
    if attr, ok := d.GetOk("associated_egress_qos_policy_id"); ok {
        o.AssociatedEgressQOSPolicyID = attr.(string)
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

func resourceRedundantPortDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.RedundantPort{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}