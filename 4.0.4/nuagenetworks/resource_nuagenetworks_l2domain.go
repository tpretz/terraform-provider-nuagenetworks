package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.4"
)

func resourceL2Domain() *schema.Resource {
    return &schema.Resource{
        Create: resourceL2DomainCreate,
        Read:   resourceL2DomainRead,
        Update: resourceL2DomainUpdate,
        Delete: resourceL2DomainDelete,
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
            "dhcp_managed": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "dpi": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "DISABLED",
            },
            "ip_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "maintenance_mode": &schema.Schema{
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
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway_mac_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "template_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "service_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "vn_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "encryption": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "policy_change_status": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "route_distinguisher": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "route_target": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "uplink_preference": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_multicast_channel_map_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_shared_network_resource_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "stretched": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "multicast": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
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

func resourceL2DomainCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize L2Domain object
    o := &vspk.L2Domain{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("dhcp_managed"); ok {
        o.DHCPManaged = attr.(bool)
    }
    if attr, ok := d.GetOk("dpi"); ok {
        o.DPI = attr.(string)
    }
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("maintenance_mode"); ok {
        o.MaintenanceMode = attr.(string)
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_mac_address"); ok {
        o.GatewayMACAddress = attr.(string)
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("template_id"); ok {
        o.TemplateID = attr.(string)
    }
    if attr, ok := d.GetOk("service_id"); ok {
        o.ServiceID = attr.(int)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("vn_id"); ok {
        o.VnId = attr.(int)
    }
    if attr, ok := d.GetOk("encryption"); ok {
        o.Encryption = attr.(string)
    }
    if attr, ok := d.GetOk("policy_change_status"); ok {
        o.PolicyChangeStatus = attr.(string)
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("route_target"); ok {
        o.RouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_preference"); ok {
        o.UplinkPreference = attr.(string)
    }
    if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
        o.AssociatedMulticastChannelMapID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_shared_network_resource_id"); ok {
        o.AssociatedSharedNetworkResourceID = attr.(string)
    }
    if attr, ok := d.GetOk("stretched"); ok {
        o.Stretched = attr.(bool)
    }
    if attr, ok := d.GetOk("multicast"); ok {
        o.Multicast = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    err := parent.CreateL2Domain(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceL2DomainRead(d, m)
}

func resourceL2DomainRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.L2Domain{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("dhcp_managed", o.DHCPManaged)
    d.Set("dpi", o.DPI)
    d.Set("ip_type", o.IPType)
    d.Set("maintenance_mode", o.MaintenanceMode)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway", o.Gateway)
    d.Set("gateway_mac_address", o.GatewayMACAddress)
    d.Set("address", o.Address)
    d.Set("template_id", o.TemplateID)
    d.Set("service_id", o.ServiceID)
    d.Set("description", o.Description)
    d.Set("netmask", o.Netmask)
    d.Set("vn_id", o.VnId)
    d.Set("encryption", o.Encryption)
    d.Set("entity_scope", o.EntityScope)
    d.Set("policy_change_status", o.PolicyChangeStatus)
    d.Set("route_distinguisher", o.RouteDistinguisher)
    d.Set("route_target", o.RouteTarget)
    d.Set("uplink_preference", o.UplinkPreference)
    d.Set("associated_multicast_channel_map_id", o.AssociatedMulticastChannelMapID)
    d.Set("associated_shared_network_resource_id", o.AssociatedSharedNetworkResourceID)
    d.Set("stretched", o.Stretched)
    d.Set("multicast", o.Multicast)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceL2DomainUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.L2Domain{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("dhcp_managed"); ok {
        o.DHCPManaged = attr.(bool)
    }
    if attr, ok := d.GetOk("dpi"); ok {
        o.DPI = attr.(string)
    }
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("maintenance_mode"); ok {
        o.MaintenanceMode = attr.(string)
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("gateway_mac_address"); ok {
        o.GatewayMACAddress = attr.(string)
    }
    if attr, ok := d.GetOk("address"); ok {
        o.Address = attr.(string)
    }
    if attr, ok := d.GetOk("template_id"); ok {
        o.TemplateID = attr.(string)
    }
    if attr, ok := d.GetOk("service_id"); ok {
        o.ServiceID = attr.(int)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("vn_id"); ok {
        o.VnId = attr.(int)
    }
    if attr, ok := d.GetOk("encryption"); ok {
        o.Encryption = attr.(string)
    }
    if attr, ok := d.GetOk("policy_change_status"); ok {
        o.PolicyChangeStatus = attr.(string)
    }
    if attr, ok := d.GetOk("route_distinguisher"); ok {
        o.RouteDistinguisher = attr.(string)
    }
    if attr, ok := d.GetOk("route_target"); ok {
        o.RouteTarget = attr.(string)
    }
    if attr, ok := d.GetOk("uplink_preference"); ok {
        o.UplinkPreference = attr.(string)
    }
    if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
        o.AssociatedMulticastChannelMapID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_shared_network_resource_id"); ok {
        o.AssociatedSharedNetworkResourceID = attr.(string)
    }
    if attr, ok := d.GetOk("stretched"); ok {
        o.Stretched = attr.(bool)
    }
    if attr, ok := d.GetOk("multicast"); ok {
        o.Multicast = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceL2DomainDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.L2Domain{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}