package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
)

func resourceBGPNeighbor() *schema.Resource {
    return &schema.Resource{
        Create: resourceBGPNeighborCreate,
        Read:   resourceBGPNeighborRead,
        Update: resourceBGPNeighborUpdate,
        Delete: resourceBGPNeighborDelete,
        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },
        Schema: map[string]*schema.Schema{
            "id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
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
            "bfd_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Default: false,
            },
            "ip_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "IPV4",
            },
            "ipv6_address": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "dampening_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "peer_as": &schema.Schema{
                Type:     schema.TypeInt,
                Required: true,
            },
            "peer_configuration": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "peer_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "session": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "domain_service_label": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_export_routing_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "associated_import_routing_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan", "parent_vport"},
            },
            "parent_vlan": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_subnet", "parent_vport"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_subnet", "parent_vlan"},
            },
        },
    }
}

func resourceBGPNeighborCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize BGPNeighbor object
    o := &vspk.BGPNeighbor{
        Name: d.Get("name").(string),
        PeerAS: d.Get("peer_as").(int),
    }
    if attr, ok := d.GetOk("bfd_enabled"); ok {
        o.BFDEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address"); ok {
        o.IPv6Address = attr.(string)
    }
    if attr, ok := d.GetOk("dampening_enabled"); ok {
        o.DampeningEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("peer_configuration"); ok {
        o.PeerConfiguration = attr.(string)
    }
    if attr, ok := d.GetOk("peer_ip"); ok {
        o.PeerIP = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("session"); ok {
        o.Session = attr.(string)
    }
    if attr, ok := d.GetOk("domain_service_label"); ok {
        o.DomainServiceLabel = attr.(string)
    }
    if attr, ok := d.GetOk("associated_export_routing_policy_id"); ok {
        o.AssociatedExportRoutingPolicyID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_import_routing_policy_id"); ok {
        o.AssociatedImportRoutingPolicyID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        err := parent.CreateBGPNeighbor(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vlan"); ok {
        parent := &vspk.VLAN{ID: attr.(string)}
        err := parent.CreateBGPNeighbor(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        err := parent.CreateBGPNeighbor(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceBGPNeighborRead(d, m)
}

func resourceBGPNeighborRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.BGPNeighbor{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("bfd_enabled", o.BFDEnabled)
    d.Set("ip_type", o.IPType)
    d.Set("ipv6_address", o.IPv6Address)
    d.Set("name", o.Name)
    d.Set("dampening_enabled", o.DampeningEnabled)
    d.Set("peer_as", o.PeerAS)
    d.Set("peer_configuration", o.PeerConfiguration)
    d.Set("peer_ip", o.PeerIP)
    d.Set("description", o.Description)
    d.Set("session", o.Session)
    d.Set("entity_scope", o.EntityScope)
    d.Set("domain_service_label", o.DomainServiceLabel)
    d.Set("associated_export_routing_policy_id", o.AssociatedExportRoutingPolicyID)
    d.Set("associated_import_routing_policy_id", o.AssociatedImportRoutingPolicyID)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceBGPNeighborUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.BGPNeighbor{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.PeerAS = d.Get("peer_as").(int)
    
    if attr, ok := d.GetOk("bfd_enabled"); ok {
        o.BFDEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("ip_type"); ok {
        o.IPType = attr.(string)
    }
    if attr, ok := d.GetOk("ipv6_address"); ok {
        o.IPv6Address = attr.(string)
    }
    if attr, ok := d.GetOk("dampening_enabled"); ok {
        o.DampeningEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("peer_configuration"); ok {
        o.PeerConfiguration = attr.(string)
    }
    if attr, ok := d.GetOk("peer_ip"); ok {
        o.PeerIP = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("session"); ok {
        o.Session = attr.(string)
    }
    if attr, ok := d.GetOk("domain_service_label"); ok {
        o.DomainServiceLabel = attr.(string)
    }
    if attr, ok := d.GetOk("associated_export_routing_policy_id"); ok {
        o.AssociatedExportRoutingPolicyID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_import_routing_policy_id"); ok {
        o.AssociatedImportRoutingPolicyID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceBGPNeighborDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.BGPNeighbor{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}