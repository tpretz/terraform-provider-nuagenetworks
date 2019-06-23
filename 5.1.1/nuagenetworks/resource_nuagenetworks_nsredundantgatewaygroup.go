package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.1.1"
)

func resourceNSRedundantGatewayGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourceNSRedundantGatewayGroupCreate,
        Read:   resourceNSRedundantGatewayGroupRead,
        Update: resourceNSRedundantGatewayGroupUpdate,
        Delete: resourceNSRedundantGatewayGroupDelete,
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
            "gateway_peer1_autodiscovered_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_peer1_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_peer1_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_peer2_autodiscovered_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_peer2_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_peer2_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "heartbeat_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Default: 500,
            },
            "heartbeat_vlanid": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Default: 4094,
            },
            "redundancy_port_ids": &schema.Schema{
                Type:     schema.TypeList,
                Optional: true,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "redundant_gateway_status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "personality": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "consecutive_failures_count": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Default: 3,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceNSRedundantGatewayGroupCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize NSRedundantGatewayGroup object
    o := &vspk.NSRedundantGatewayGroup{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("gateway_peer1_autodiscovered_gateway_id"); ok {
        o.GatewayPeer1AutodiscoveredGatewayID = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_peer1_id"); ok {
        o.GatewayPeer1ID = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_peer1_name"); ok {
        o.GatewayPeer1Name = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_peer2_autodiscovered_gateway_id"); ok {
        o.GatewayPeer2AutodiscoveredGatewayID = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_peer2_id"); ok {
        o.GatewayPeer2ID = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_peer2_name"); ok {
        o.GatewayPeer2Name = attr.(string)
    }
    if attr, ok := d.GetOk("heartbeat_interval"); ok {
        o.HeartbeatInterval = attr.(int)
    }
    if attr, ok := d.GetOk("heartbeat_vlanid"); ok {
        o.HeartbeatVLANID = attr.(int)
    }
    if attr, ok := d.GetOk("redundancy_port_ids"); ok {
        o.RedundancyPortIDs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("redundant_gateway_status"); ok {
        o.RedundantGatewayStatus = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("personality"); ok {
        o.Personality = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("enterprise_id"); ok {
        o.EnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("consecutive_failures_count"); ok {
        o.ConsecutiveFailuresCount = attr.(int)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateNSRedundantGatewayGroup(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceNSRedundantGatewayGroupRead(d, m)
}

func resourceNSRedundantGatewayGroupRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NSRedundantGatewayGroup{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway_peer1_autodiscovered_gateway_id", o.GatewayPeer1AutodiscoveredGatewayID)
    d.Set("gateway_peer1_id", o.GatewayPeer1ID)
    d.Set("gateway_peer1_name", o.GatewayPeer1Name)
    d.Set("gateway_peer2_autodiscovered_gateway_id", o.GatewayPeer2AutodiscoveredGatewayID)
    d.Set("gateway_peer2_id", o.GatewayPeer2ID)
    d.Set("gateway_peer2_name", o.GatewayPeer2Name)
    d.Set("heartbeat_interval", o.HeartbeatInterval)
    d.Set("heartbeat_vlanid", o.HeartbeatVLANID)
    d.Set("redundancy_port_ids", o.RedundancyPortIDs)
    d.Set("redundant_gateway_status", o.RedundantGatewayStatus)
    d.Set("permitted_action", o.PermittedAction)
    d.Set("personality", o.Personality)
    d.Set("description", o.Description)
    d.Set("enterprise_id", o.EnterpriseID)
    d.Set("entity_scope", o.EntityScope)
    d.Set("consecutive_failures_count", o.ConsecutiveFailuresCount)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceNSRedundantGatewayGroupUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NSRedundantGatewayGroup{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("gateway_peer1_autodiscovered_gateway_id"); ok {
        o.GatewayPeer1AutodiscoveredGatewayID = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_peer1_id"); ok {
        o.GatewayPeer1ID = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_peer1_name"); ok {
        o.GatewayPeer1Name = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_peer2_autodiscovered_gateway_id"); ok {
        o.GatewayPeer2AutodiscoveredGatewayID = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_peer2_id"); ok {
        o.GatewayPeer2ID = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_peer2_name"); ok {
        o.GatewayPeer2Name = attr.(string)
    }
    if attr, ok := d.GetOk("heartbeat_interval"); ok {
        o.HeartbeatInterval = attr.(int)
    }
    if attr, ok := d.GetOk("heartbeat_vlanid"); ok {
        o.HeartbeatVLANID = attr.(int)
    }
    if attr, ok := d.GetOk("redundancy_port_ids"); ok {
        o.RedundancyPortIDs = attr.([]interface{})
    }
    if attr, ok := d.GetOk("redundant_gateway_status"); ok {
        o.RedundantGatewayStatus = attr.(string)
    }
    if attr, ok := d.GetOk("permitted_action"); ok {
        o.PermittedAction = attr.(string)
    }
    if attr, ok := d.GetOk("personality"); ok {
        o.Personality = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("enterprise_id"); ok {
        o.EnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("consecutive_failures_count"); ok {
        o.ConsecutiveFailuresCount = attr.(int)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceNSRedundantGatewayGroupDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NSRedundantGatewayGroup{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}