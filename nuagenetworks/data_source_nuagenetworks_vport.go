package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVPort() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVPortRead,
        Schema: map[string]*schema.Schema{
            "filter": dataSourceFiltersSchema(),
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "fip_ignore_default_route": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vlan": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vlanid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dpi": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "has_attached_interfaces": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_mac_move_role": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_port_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "active": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "address_spoofing": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "peer_operational_state": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "segmentation_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "segmentation_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "service_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zone_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "operational_state": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "trunk_role": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "assoc_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_egress_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_floating_ip_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_gateway_personality": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_gateway_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ingress_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_multicast_channel_map_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ssid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_send_multicast_channel_map_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_trunk_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "sub_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multi_nic_vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gw_eligible": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "system_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_egress_profile", "parent_zone", "parent_vrs", "parent_multi_nic_vport", "parent_floating_ip", "parent_trunk", "parent_ingress_profile", "parent_overlay_mirror_destination"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_egress_profile", "parent_zone", "parent_vrs", "parent_multi_nic_vport", "parent_floating_ip", "parent_trunk", "parent_ingress_profile", "parent_overlay_mirror_destination"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_redirection_target", "parent_policy_group", "parent_egress_profile", "parent_zone", "parent_vrs", "parent_multi_nic_vport", "parent_floating_ip", "parent_trunk", "parent_ingress_profile", "parent_overlay_mirror_destination"},
            },
            "parent_redirection_target": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_policy_group", "parent_egress_profile", "parent_zone", "parent_vrs", "parent_multi_nic_vport", "parent_floating_ip", "parent_trunk", "parent_ingress_profile", "parent_overlay_mirror_destination"},
            },
            "parent_policy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_egress_profile", "parent_zone", "parent_vrs", "parent_multi_nic_vport", "parent_floating_ip", "parent_trunk", "parent_ingress_profile", "parent_overlay_mirror_destination"},
            },
            "parent_egress_profile": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_zone", "parent_vrs", "parent_multi_nic_vport", "parent_floating_ip", "parent_trunk", "parent_ingress_profile", "parent_overlay_mirror_destination"},
            },
            "parent_zone": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_egress_profile", "parent_vrs", "parent_multi_nic_vport", "parent_floating_ip", "parent_trunk", "parent_ingress_profile", "parent_overlay_mirror_destination"},
            },
            "parent_vrs": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_egress_profile", "parent_zone", "parent_multi_nic_vport", "parent_floating_ip", "parent_trunk", "parent_ingress_profile", "parent_overlay_mirror_destination"},
            },
            "parent_multi_nic_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_egress_profile", "parent_zone", "parent_vrs", "parent_floating_ip", "parent_trunk", "parent_ingress_profile", "parent_overlay_mirror_destination"},
            },
            "parent_floating_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_egress_profile", "parent_zone", "parent_vrs", "parent_multi_nic_vport", "parent_trunk", "parent_ingress_profile", "parent_overlay_mirror_destination"},
            },
            "parent_trunk": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_egress_profile", "parent_zone", "parent_vrs", "parent_multi_nic_vport", "parent_floating_ip", "parent_ingress_profile", "parent_overlay_mirror_destination"},
            },
            "parent_ingress_profile": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_egress_profile", "parent_zone", "parent_vrs", "parent_multi_nic_vport", "parent_floating_ip", "parent_trunk", "parent_overlay_mirror_destination"},
            },
            "parent_overlay_mirror_destination": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_redirection_target", "parent_policy_group", "parent_egress_profile", "parent_zone", "parent_vrs", "parent_multi_nic_vport", "parent_floating_ip", "parent_trunk", "parent_ingress_profile"},
            },
        },
    }
}


func dataSourceVPortRead(d *schema.ResourceData, m interface{}) (err error) {
    filteredVPorts := vspk.VPortsList{}
    fetchFilter := &bambou.FetchingInfo{}
    
    filters, filtersOk := d.GetOk("filter")
    if filtersOk {
        fetchFilter = bambou.NewFetchingInfo()
        for _, v := range filters.(*schema.Set).List() {
            m := v.(map[string]interface{})
            if fetchFilter.Filter != "" {
                fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string),  m["operator"].(string),  m["value"].(string))
            } else {
                fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
            }
           
        }
    }
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_redirection_target"); ok {
        parent := &vspk.RedirectionTarget{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_policy_group"); ok {
        parent := &vspk.PolicyGroup{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_egress_profile"); ok {
        parent := &vspk.EgressProfile{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_zone"); ok {
        parent := &vspk.Zone{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_vrs"); ok {
        parent := &vspk.VRS{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_multi_nic_vport"); ok {
        parent := &vspk.MultiNICVPort{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_floating_ip"); ok {
        parent := &vspk.FloatingIp{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_trunk"); ok {
        parent := &vspk.Trunk{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_ingress_profile"); ok {
        parent := &vspk.IngressProfile{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    } else if attr, ok := d.GetOk("parent_overlay_mirror_destination"); ok {
        parent := &vspk.OverlayMirrorDestination{ID: attr.(string)}
        filteredVPorts, err = parent.VPorts(fetchFilter)
        if err != nil {
            return
        }
    }

    VPort := &vspk.VPort{}

    if len(filteredVPorts) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVPorts) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VPort = filteredVPorts[0]

    d.Set("fip_ignore_default_route", VPort.FIPIgnoreDefaultRoute)
    d.Set("vlan", VPort.VLAN)
    d.Set("vlanid", VPort.VLANID)
    d.Set("dpi", VPort.DPI)
    d.Set("name", VPort.Name)
    d.Set("has_attached_interfaces", VPort.HasAttachedInterfaces)
    d.Set("last_updated_by", VPort.LastUpdatedBy)
    d.Set("gateway_mac_move_role", VPort.GatewayMACMoveRole)
    d.Set("gateway_port_name", VPort.GatewayPortName)
    d.Set("active", VPort.Active)
    d.Set("address_spoofing", VPort.AddressSpoofing)
    d.Set("peer_operational_state", VPort.PeerOperationalState)
    d.Set("segmentation_id", VPort.SegmentationID)
    d.Set("segmentation_type", VPort.SegmentationType)
    d.Set("service_id", VPort.ServiceID)
    d.Set("description", VPort.Description)
    d.Set("entity_scope", VPort.EntityScope)
    d.Set("domain_id", VPort.DomainID)
    d.Set("zone_id", VPort.ZoneID)
    d.Set("operational_state", VPort.OperationalState)
    d.Set("trunk_role", VPort.TrunkRole)
    d.Set("assoc_entity_id", VPort.AssocEntityID)
    d.Set("associated_egress_profile_id", VPort.AssociatedEgressProfileID)
    d.Set("associated_floating_ip_id", VPort.AssociatedFloatingIPID)
    d.Set("associated_gateway_id", VPort.AssociatedGatewayID)
    d.Set("associated_gateway_personality", VPort.AssociatedGatewayPersonality)
    d.Set("associated_gateway_type", VPort.AssociatedGatewayType)
    d.Set("associated_ingress_profile_id", VPort.AssociatedIngressProfileID)
    d.Set("associated_multicast_channel_map_id", VPort.AssociatedMulticastChannelMapID)
    d.Set("associated_ssid", VPort.AssociatedSSID)
    d.Set("associated_send_multicast_channel_map_id", VPort.AssociatedSendMulticastChannelMapID)
    d.Set("associated_trunk_id", VPort.AssociatedTrunkID)
    d.Set("sub_type", VPort.SubType)
    d.Set("multi_nic_vport_id", VPort.MultiNICVPortID)
    d.Set("multicast", VPort.Multicast)
    d.Set("gw_eligible", VPort.GwEligible)
    d.Set("external_id", VPort.ExternalID)
    d.Set("type", VPort.Type)
    d.Set("system_type", VPort.SystemType)
    
    d.Set("id", VPort.Identifier())
    d.Set("parent_id", VPort.ParentID)
    d.Set("parent_type", VPort.ParentType)
    d.Set("owner", VPort.Owner)

    d.SetId(VPort.Identifier())
    
    return
}