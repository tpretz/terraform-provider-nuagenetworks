package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
)

func resourceVPort() *schema.Resource {
    return &schema.Resource{
        Create: resourceVPortCreate,
        Read:   resourceVPortRead,
        Update: resourceVPortUpdate,
        Delete: resourceVPortDelete,
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
            "fip_ignore_default_route": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vlan": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "vlanid": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "dpi": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "INHERITED",
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "has_attached_interfaces": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_mac_move_role": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_port_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "active": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "address_spoofing": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "peer_operational_state": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "segmentation_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "segmentation_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "service_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "zone_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "operational_state": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "trunk_role": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "assoc_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_egress_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_floating_ip_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_gateway_personality": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_gateway_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_ingress_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_multicast_channel_map_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_ssid": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_send_multicast_channel_map_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_trunk_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "sub_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "NONE",
            },
            "multi_nic_vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "multicast": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gw_eligible": &schema.Schema{
                Type:     schema.TypeBool,
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
            "system_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_subnet"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain"},
            },
        },
    }
}

func resourceVPortCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VPort object
    o := &vspk.VPort{
        Name: d.Get("name").(string),
        AddressSpoofing: d.Get("address_spoofing").(string),
        Type: d.Get("type").(string),
    }
    if attr, ok := d.GetOk("fip_ignore_default_route"); ok {
        o.FIPIgnoreDefaultRoute = attr.(string)
    }
    if attr, ok := d.GetOk("vlanid"); ok {
        o.VLANID = attr.(string)
    }
    if attr, ok := d.GetOk("dpi"); ok {
        o.DPI = attr.(string)
    }
    if attr, ok := d.GetOk("has_attached_interfaces"); ok {
        o.HasAttachedInterfaces = attr.(bool)
    }
    if attr, ok := d.GetOk("gateway_mac_move_role"); ok {
        o.GatewayMACMoveRole = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_port_name"); ok {
        o.GatewayPortName = attr.(string)
    }
    if attr, ok := d.GetOk("active"); ok {
        o.Active = attr.(bool)
    }
    if attr, ok := d.GetOk("peer_operational_state"); ok {
        o.PeerOperationalState = attr.(string)
    }
    if attr, ok := d.GetOk("segmentation_id"); ok {
        o.SegmentationID = attr.(int)
    }
    if attr, ok := d.GetOk("segmentation_type"); ok {
        o.SegmentationType = attr.(string)
    }
    if attr, ok := d.GetOk("service_id"); ok {
        o.ServiceID = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("domain_id"); ok {
        o.DomainID = attr.(string)
    }
    if attr, ok := d.GetOk("zone_id"); ok {
        o.ZoneID = attr.(string)
    }
    if attr, ok := d.GetOk("operational_state"); ok {
        o.OperationalState = attr.(string)
    }
    if attr, ok := d.GetOk("trunk_role"); ok {
        o.TrunkRole = attr.(string)
    }
    if attr, ok := d.GetOk("assoc_entity_id"); ok {
        o.AssocEntityID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_egress_profile_id"); ok {
        o.AssociatedEgressProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_floating_ip_id"); ok {
        o.AssociatedFloatingIPID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_ingress_profile_id"); ok {
        o.AssociatedIngressProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
        o.AssociatedMulticastChannelMapID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_ssid"); ok {
        o.AssociatedSSID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_send_multicast_channel_map_id"); ok {
        o.AssociatedSendMulticastChannelMapID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_trunk_id"); ok {
        o.AssociatedTrunkID = attr.(string)
    }
    if attr, ok := d.GetOk("sub_type"); ok {
        o.SubType = attr.(string)
    }
    if attr, ok := d.GetOk("multi_nic_vport_id"); ok {
        o.MultiNICVPortID = attr.(string)
    }
    if attr, ok := d.GetOk("multicast"); ok {
        o.Multicast = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("system_type"); ok {
        o.SystemType = attr.(string)
    }
    if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        err := parent.CreateVPort(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        err := parent.CreateVPort(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    if attr, ok := d.GetOk("redirectiontargets"); ok {
        o.AssignRedirectionTargets(attr.(vspk.RedirectionTargetsList))
    }
    if attr, ok := d.GetOk("policygroups"); ok {
        o.AssignPolicyGroups(attr.(vspk.PolicyGroupsList))
    }
    if attr, ok := d.GetOk("applicationperformancemanagements"); ok {
        o.AssignApplicationperformancemanagements(attr.(vspk.ApplicationperformancemanagementsList))
    }
    return resourceVPortRead(d, m)
}

func resourceVPortRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VPort{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("fip_ignore_default_route", o.FIPIgnoreDefaultRoute)
    d.Set("vlan", o.VLAN)
    d.Set("vlanid", o.VLANID)
    d.Set("dpi", o.DPI)
    d.Set("name", o.Name)
    d.Set("has_attached_interfaces", o.HasAttachedInterfaces)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway_mac_move_role", o.GatewayMACMoveRole)
    d.Set("gateway_port_name", o.GatewayPortName)
    d.Set("active", o.Active)
    d.Set("address_spoofing", o.AddressSpoofing)
    d.Set("peer_operational_state", o.PeerOperationalState)
    d.Set("segmentation_id", o.SegmentationID)
    d.Set("segmentation_type", o.SegmentationType)
    d.Set("service_id", o.ServiceID)
    d.Set("description", o.Description)
    d.Set("entity_scope", o.EntityScope)
    d.Set("domain_id", o.DomainID)
    d.Set("zone_id", o.ZoneID)
    d.Set("operational_state", o.OperationalState)
    d.Set("trunk_role", o.TrunkRole)
    d.Set("assoc_entity_id", o.AssocEntityID)
    d.Set("associated_egress_profile_id", o.AssociatedEgressProfileID)
    d.Set("associated_floating_ip_id", o.AssociatedFloatingIPID)
    d.Set("associated_gateway_id", o.AssociatedGatewayID)
    d.Set("associated_gateway_personality", o.AssociatedGatewayPersonality)
    d.Set("associated_gateway_type", o.AssociatedGatewayType)
    d.Set("associated_ingress_profile_id", o.AssociatedIngressProfileID)
    d.Set("associated_multicast_channel_map_id", o.AssociatedMulticastChannelMapID)
    d.Set("associated_ssid", o.AssociatedSSID)
    d.Set("associated_send_multicast_channel_map_id", o.AssociatedSendMulticastChannelMapID)
    d.Set("associated_trunk_id", o.AssociatedTrunkID)
    d.Set("sub_type", o.SubType)
    d.Set("multi_nic_vport_id", o.MultiNICVPortID)
    d.Set("multicast", o.Multicast)
    d.Set("gw_eligible", o.GwEligible)
    d.Set("external_id", o.ExternalID)
    d.Set("type", o.Type)
    d.Set("system_type", o.SystemType)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVPortUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VPort{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.AddressSpoofing = d.Get("address_spoofing").(string)
    o.Type = d.Get("type").(string)
    
    if attr, ok := d.GetOk("fip_ignore_default_route"); ok {
        o.FIPIgnoreDefaultRoute = attr.(string)
    }
    if attr, ok := d.GetOk("vlanid"); ok {
        o.VLANID = attr.(string)
    }
    if attr, ok := d.GetOk("dpi"); ok {
        o.DPI = attr.(string)
    }
    if attr, ok := d.GetOk("has_attached_interfaces"); ok {
        o.HasAttachedInterfaces = attr.(bool)
    }
    if attr, ok := d.GetOk("gateway_mac_move_role"); ok {
        o.GatewayMACMoveRole = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_port_name"); ok {
        o.GatewayPortName = attr.(string)
    }
    if attr, ok := d.GetOk("active"); ok {
        o.Active = attr.(bool)
    }
    if attr, ok := d.GetOk("peer_operational_state"); ok {
        o.PeerOperationalState = attr.(string)
    }
    if attr, ok := d.GetOk("segmentation_id"); ok {
        o.SegmentationID = attr.(int)
    }
    if attr, ok := d.GetOk("segmentation_type"); ok {
        o.SegmentationType = attr.(string)
    }
    if attr, ok := d.GetOk("service_id"); ok {
        o.ServiceID = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("domain_id"); ok {
        o.DomainID = attr.(string)
    }
    if attr, ok := d.GetOk("zone_id"); ok {
        o.ZoneID = attr.(string)
    }
    if attr, ok := d.GetOk("operational_state"); ok {
        o.OperationalState = attr.(string)
    }
    if attr, ok := d.GetOk("trunk_role"); ok {
        o.TrunkRole = attr.(string)
    }
    if attr, ok := d.GetOk("assoc_entity_id"); ok {
        o.AssocEntityID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_egress_profile_id"); ok {
        o.AssociatedEgressProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_floating_ip_id"); ok {
        o.AssociatedFloatingIPID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_ingress_profile_id"); ok {
        o.AssociatedIngressProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
        o.AssociatedMulticastChannelMapID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_ssid"); ok {
        o.AssociatedSSID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_send_multicast_channel_map_id"); ok {
        o.AssociatedSendMulticastChannelMapID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_trunk_id"); ok {
        o.AssociatedTrunkID = attr.(string)
    }
    if attr, ok := d.GetOk("sub_type"); ok {
        o.SubType = attr.(string)
    }
    if attr, ok := d.GetOk("multi_nic_vport_id"); ok {
        o.MultiNICVPortID = attr.(string)
    }
    if attr, ok := d.GetOk("multicast"); ok {
        o.Multicast = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("system_type"); ok {
        o.SystemType = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVPortDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VPort{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}