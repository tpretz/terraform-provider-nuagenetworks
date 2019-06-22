package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
)

func resourceContainerInterface() *schema.Resource {
    return &schema.Resource{
        Create: resourceContainerInterfaceCreate,
        Read:   resourceContainerInterfaceRead,
        Update: resourceContainerInterfaceUpdate,
        Delete: resourceContainerInterfaceDelete,
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
            "mac": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "vport_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "ipv6_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "ipv6_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "network_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "network_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "tier_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "endpoint_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "policy_decision_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "zone_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "zone_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "container_uuid": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_floating_ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "attached_network_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "attached_network_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "multi_nic_vport_name": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_container": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceContainerInterfaceCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize ContainerInterface object
    o := &vspk.ContainerInterface{
    }
    if attr, ok := d.GetOk("mac"); ok {
        o.MAC = attr.(string)
    }
    if attr, ok := d.GetOk("ip_address"); ok {
        o.IPAddress = attr.(string)
    }
    if attr, ok := d.GetOk("vport_id"); ok {
        o.VPortID = attr.(string)
    }
    if attr, ok := d.GetOk("vport_name"); ok {
        o.VPortName = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address"); ok {
        o.IPv6Address = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_gateway"); ok {
        o.IPv6Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("network_id"); ok {
        o.NetworkID = attr.(string)
    }
    if attr, ok := d.GetOk("network_name"); ok {
        o.NetworkName = attr.(string)
    }
    if attr, ok := d.GetOk("tier_id"); ok {
        o.TierID = attr.(string)
    }
    if attr, ok := d.GetOk("endpoint_id"); ok {
        o.EndpointID = attr.(string)
    }
    if attr, ok := d.GetOk("policy_decision_id"); ok {
        o.PolicyDecisionID = attr.(string)
    }
    if attr, ok := d.GetOk("domain_id"); ok {
        o.DomainID = attr.(string)
    }
    if attr, ok := d.GetOk("domain_name"); ok {
        o.DomainName = attr.(string)
    }
    if attr, ok := d.GetOk("zone_id"); ok {
        o.ZoneID = attr.(string)
    }
    if attr, ok := d.GetOk("zone_name"); ok {
        o.ZoneName = attr.(string)
    }
    if attr, ok := d.GetOk("container_uuid"); ok {
        o.ContainerUUID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_floating_ip_address"); ok {
        o.AssociatedFloatingIPAddress = attr.(string)
    }
    if attr, ok := d.GetOk("attached_network_id"); ok {
        o.AttachedNetworkID = attr.(string)
    }
    if attr, ok := d.GetOk("attached_network_type"); ok {
        o.AttachedNetworkType = attr.(string)
    }
    if attr, ok := d.GetOk("multi_nic_vport_name"); ok {
        o.MultiNICVPortName = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.Container{ID: d.Get("parent_container").(string)}
    err := parent.CreateContainerInterface(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceContainerInterfaceRead(d, m)
}

func resourceContainerInterfaceRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ContainerInterface{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("mac", o.MAC)
    d.Set("ip_address", o.IPAddress)
    d.Set("vport_id", o.VPortID)
    d.Set("vport_name", o.VPortName)
    d.Set("ipv6_address", o.IPv6Address)
    d.Set("ipv6_gateway", o.IPv6Gateway)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("gateway", o.Gateway)
    d.Set("netmask", o.Netmask)
    d.Set("network_id", o.NetworkID)
    d.Set("network_name", o.NetworkName)
    d.Set("tier_id", o.TierID)
    d.Set("endpoint_id", o.EndpointID)
    d.Set("entity_scope", o.EntityScope)
    d.Set("policy_decision_id", o.PolicyDecisionID)
    d.Set("domain_id", o.DomainID)
    d.Set("domain_name", o.DomainName)
    d.Set("zone_id", o.ZoneID)
    d.Set("zone_name", o.ZoneName)
    d.Set("container_uuid", o.ContainerUUID)
    d.Set("associated_floating_ip_address", o.AssociatedFloatingIPAddress)
    d.Set("attached_network_id", o.AttachedNetworkID)
    d.Set("attached_network_type", o.AttachedNetworkType)
    d.Set("multi_nic_vport_name", o.MultiNICVPortName)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceContainerInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ContainerInterface{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("mac"); ok {
        o.MAC = attr.(string)
    }
    if attr, ok := d.GetOk("ip_address"); ok {
        o.IPAddress = attr.(string)
    }
    if attr, ok := d.GetOk("vport_id"); ok {
        o.VPortID = attr.(string)
    }
    if attr, ok := d.GetOk("vport_name"); ok {
        o.VPortName = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address"); ok {
        o.IPv6Address = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_gateway"); ok {
        o.IPv6Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("gateway"); ok {
        o.Gateway = attr.(string)
    }
    if attr, ok := d.GetOk("netmask"); ok {
        o.Netmask = attr.(string)
    }
    if attr, ok := d.GetOk("network_id"); ok {
        o.NetworkID = attr.(string)
    }
    if attr, ok := d.GetOk("network_name"); ok {
        o.NetworkName = attr.(string)
    }
    if attr, ok := d.GetOk("tier_id"); ok {
        o.TierID = attr.(string)
    }
    if attr, ok := d.GetOk("endpoint_id"); ok {
        o.EndpointID = attr.(string)
    }
    if attr, ok := d.GetOk("policy_decision_id"); ok {
        o.PolicyDecisionID = attr.(string)
    }
    if attr, ok := d.GetOk("domain_id"); ok {
        o.DomainID = attr.(string)
    }
    if attr, ok := d.GetOk("domain_name"); ok {
        o.DomainName = attr.(string)
    }
    if attr, ok := d.GetOk("zone_id"); ok {
        o.ZoneID = attr.(string)
    }
    if attr, ok := d.GetOk("zone_name"); ok {
        o.ZoneName = attr.(string)
    }
    if attr, ok := d.GetOk("container_uuid"); ok {
        o.ContainerUUID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_floating_ip_address"); ok {
        o.AssociatedFloatingIPAddress = attr.(string)
    }
    if attr, ok := d.GetOk("attached_network_id"); ok {
        o.AttachedNetworkID = attr.(string)
    }
    if attr, ok := d.GetOk("attached_network_type"); ok {
        o.AttachedNetworkType = attr.(string)
    }
    if attr, ok := d.GetOk("multi_nic_vport_name"); ok {
        o.MultiNICVPortName = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceContainerInterfaceDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.ContainerInterface{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}