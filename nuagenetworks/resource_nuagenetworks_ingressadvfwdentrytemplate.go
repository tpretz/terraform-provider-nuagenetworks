package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
)

func resourceIngressAdvFwdEntryTemplate() *schema.Resource {
    return &schema.Resource{
        Create: resourceIngressAdvFwdEntryTemplateCreate,
        Read:   resourceIngressAdvFwdEntryTemplateRead,
        Update: resourceIngressAdvFwdEntryTemplateUpdate,
        Delete: resourceIngressAdvFwdEntryTemplateDelete,
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
            "acl_template_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "icmp_code": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "icmp_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "fc_override": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ipv6_address_override": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "dscp": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "dscp_remarking": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "failsafe_datapath": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "FAIL_TO_BLOCK",
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "action": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "address_override": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "redirect_rewrite_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "redirect_rewrite_value": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "redirect_vport_tag_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "remote_uplink_preference": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "DEFAULT",
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "destination_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "network_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "network_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "mirror_destination_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vlan_range": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "flow_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "enterprise_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "location_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "location_type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "policy_state": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "uplink_preference": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "app_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "NONE",
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "protocol": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "is_sla_aware": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "associated_application_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_forwarding_path_list_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_live_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_traffic_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_traffic_type_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "stats_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "stats_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "ether_type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_ingress_adv_fwd_template": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceIngressAdvFwdEntryTemplateCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize IngressAdvFwdEntryTemplate object
    o := &vspk.IngressAdvFwdEntryTemplate{
        DSCP: d.Get("dscp").(string),
        Action: d.Get("action").(string),
        RedirectVPortTagID: d.Get("redirect_vport_tag_id").(string),
        LocationType: d.Get("location_type").(string),
        EtherType: d.Get("ether_type").(string),
    }
    if attr, ok := d.GetOk("acl_template_name"); ok {
        o.ACLTemplateName = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_code"); ok {
        o.ICMPCode = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_type"); ok {
        o.ICMPType = attr.(string)
    }
    if attr, ok := d.GetOk("fc_override"); ok {
        o.FCOverride = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address_override"); ok {
        o.IPv6AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("dscp_remarking"); ok {
        o.DSCPRemarking = attr.(string)
    }
    if attr, ok := d.GetOk("failsafe_datapath"); ok {
        o.FailsafeDatapath = attr.(string)
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("address_override"); ok {
        o.AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("redirect_rewrite_type"); ok {
        o.RedirectRewriteType = attr.(string)
    }
    if attr, ok := d.GetOk("redirect_rewrite_value"); ok {
        o.RedirectRewriteValue = attr.(string)
    }
    if attr, ok := d.GetOk("remote_uplink_preference"); ok {
        o.RemoteUplinkPreference = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_port"); ok {
        o.DestinationPort = attr.(string)
    }
    if attr, ok := d.GetOk("network_id"); ok {
        o.NetworkID = attr.(string)
    }
    if attr, ok := d.GetOk("network_type"); ok {
        o.NetworkType = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_destination_id"); ok {
        o.MirrorDestinationID = attr.(string)
    }
    if attr, ok := d.GetOk("vlan_range"); ok {
        o.VlanRange = attr.(string)
    }
    if attr, ok := d.GetOk("flow_logging_enabled"); ok {
        o.FlowLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("location_id"); ok {
        o.LocationID = attr.(string)
    }
    if attr, ok := d.GetOk("policy_state"); ok {
        o.PolicyState = attr.(string)
    }
    if attr, ok := d.GetOk("source_port"); ok {
        o.SourcePort = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_preference"); ok {
        o.UplinkPreference = attr.(string)
    }
    if attr, ok := d.GetOk("app_type"); ok {
        o.AppType = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("protocol"); ok {
        o.Protocol = attr.(string)
    }
    if attr, ok := d.GetOk("is_sla_aware"); ok {
        o.IsSLAAware = attr.(bool)
    }
    if attr, ok := d.GetOk("associated_application_id"); ok {
        o.AssociatedApplicationID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_forwarding_path_list_id"); ok {
        o.AssociatedForwardingPathListID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_live_entity_id"); ok {
        o.AssociatedLiveEntityID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_traffic_type"); ok {
        o.AssociatedTrafficType = attr.(string)
    }
    if attr, ok := d.GetOk("associated_traffic_type_id"); ok {
        o.AssociatedTrafficTypeID = attr.(string)
    }
    if attr, ok := d.GetOk("stats_logging_enabled"); ok {
        o.StatsLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.IngressAdvFwdTemplate{ID: d.Get("parent_ingress_adv_fwd_template").(string)}
    err := parent.CreateIngressAdvFwdEntryTemplate(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceIngressAdvFwdEntryTemplateRead(d, m)
}

func resourceIngressAdvFwdEntryTemplateRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IngressAdvFwdEntryTemplate{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("acl_template_name", o.ACLTemplateName)
    d.Set("icmp_code", o.ICMPCode)
    d.Set("icmp_type", o.ICMPType)
    d.Set("fc_override", o.FCOverride)
    d.Set("ipv6_address_override", o.IPv6AddressOverride)
    d.Set("dscp", o.DSCP)
    d.Set("dscp_remarking", o.DSCPRemarking)
    d.Set("failsafe_datapath", o.FailsafeDatapath)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("action", o.Action)
    d.Set("address_override", o.AddressOverride)
    d.Set("redirect_rewrite_type", o.RedirectRewriteType)
    d.Set("redirect_rewrite_value", o.RedirectRewriteValue)
    d.Set("redirect_vport_tag_id", o.RedirectVPortTagID)
    d.Set("remote_uplink_preference", o.RemoteUplinkPreference)
    d.Set("description", o.Description)
    d.Set("destination_port", o.DestinationPort)
    d.Set("network_id", o.NetworkID)
    d.Set("network_type", o.NetworkType)
    d.Set("mirror_destination_id", o.MirrorDestinationID)
    d.Set("vlan_range", o.VlanRange)
    d.Set("flow_logging_enabled", o.FlowLoggingEnabled)
    d.Set("enterprise_name", o.EnterpriseName)
    d.Set("entity_scope", o.EntityScope)
    d.Set("location_id", o.LocationID)
    d.Set("location_type", o.LocationType)
    d.Set("policy_state", o.PolicyState)
    d.Set("domain_name", o.DomainName)
    d.Set("source_port", o.SourcePort)
    d.Set("uplink_preference", o.UplinkPreference)
    d.Set("app_type", o.AppType)
    d.Set("priority", o.Priority)
    d.Set("protocol", o.Protocol)
    d.Set("is_sla_aware", o.IsSLAAware)
    d.Set("associated_application_id", o.AssociatedApplicationID)
    d.Set("associated_forwarding_path_list_id", o.AssociatedForwardingPathListID)
    d.Set("associated_live_entity_id", o.AssociatedLiveEntityID)
    d.Set("associated_traffic_type", o.AssociatedTrafficType)
    d.Set("associated_traffic_type_id", o.AssociatedTrafficTypeID)
    d.Set("stats_id", o.StatsID)
    d.Set("stats_logging_enabled", o.StatsLoggingEnabled)
    d.Set("ether_type", o.EtherType)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceIngressAdvFwdEntryTemplateUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IngressAdvFwdEntryTemplate{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.DSCP = d.Get("dscp").(string)
    o.Action = d.Get("action").(string)
    o.RedirectVPortTagID = d.Get("redirect_vport_tag_id").(string)
    o.LocationType = d.Get("location_type").(string)
    o.EtherType = d.Get("ether_type").(string)
    
    if attr, ok := d.GetOk("acl_template_name"); ok {
        o.ACLTemplateName = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_code"); ok {
        o.ICMPCode = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_type"); ok {
        o.ICMPType = attr.(string)
    }
    if attr, ok := d.GetOk("fc_override"); ok {
        o.FCOverride = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address_override"); ok {
        o.IPv6AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("dscp_remarking"); ok {
        o.DSCPRemarking = attr.(string)
    }
    if attr, ok := d.GetOk("failsafe_datapath"); ok {
        o.FailsafeDatapath = attr.(string)
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("address_override"); ok {
        o.AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("redirect_rewrite_type"); ok {
        o.RedirectRewriteType = attr.(string)
    }
    if attr, ok := d.GetOk("redirect_rewrite_value"); ok {
        o.RedirectRewriteValue = attr.(string)
    }
    if attr, ok := d.GetOk("remote_uplink_preference"); ok {
        o.RemoteUplinkPreference = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("destination_port"); ok {
        o.DestinationPort = attr.(string)
    }
    if attr, ok := d.GetOk("network_id"); ok {
        o.NetworkID = attr.(string)
    }
    if attr, ok := d.GetOk("network_type"); ok {
        o.NetworkType = attr.(string)
    }
    if attr, ok := d.GetOk("mirror_destination_id"); ok {
        o.MirrorDestinationID = attr.(string)
    }
    if attr, ok := d.GetOk("vlan_range"); ok {
        o.VlanRange = attr.(string)
    }
    if attr, ok := d.GetOk("flow_logging_enabled"); ok {
        o.FlowLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("location_id"); ok {
        o.LocationID = attr.(string)
    }
    if attr, ok := d.GetOk("policy_state"); ok {
        o.PolicyState = attr.(string)
    }
    if attr, ok := d.GetOk("source_port"); ok {
        o.SourcePort = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_preference"); ok {
        o.UplinkPreference = attr.(string)
    }
    if attr, ok := d.GetOk("app_type"); ok {
        o.AppType = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("protocol"); ok {
        o.Protocol = attr.(string)
    }
    if attr, ok := d.GetOk("is_sla_aware"); ok {
        o.IsSLAAware = attr.(bool)
    }
    if attr, ok := d.GetOk("associated_application_id"); ok {
        o.AssociatedApplicationID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_forwarding_path_list_id"); ok {
        o.AssociatedForwardingPathListID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_live_entity_id"); ok {
        o.AssociatedLiveEntityID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_traffic_type"); ok {
        o.AssociatedTrafficType = attr.(string)
    }
    if attr, ok := d.GetOk("associated_traffic_type_id"); ok {
        o.AssociatedTrafficTypeID = attr.(string)
    }
    if attr, ok := d.GetOk("stats_logging_enabled"); ok {
        o.StatsLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceIngressAdvFwdEntryTemplateDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IngressAdvFwdEntryTemplate{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}