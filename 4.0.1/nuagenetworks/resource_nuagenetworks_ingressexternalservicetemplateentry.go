package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.1"
)

func resourceIngressExternalServiceTemplateEntry() *schema.Resource {
    return &schema.Resource{
        Create: resourceIngressExternalServiceTemplateEntryCreate,
        Read:   resourceIngressExternalServiceTemplateEntryRead,
        Update: resourceIngressExternalServiceTemplateEntryUpdate,
        Delete: resourceIngressExternalServiceTemplateEntryDelete,
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
            "dscp": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
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
            "redirect_external_service_end_point_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
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
            "flow_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
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
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
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
            "associated_application_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_application_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_application_object_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_live_entity_id": &schema.Schema{
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
            "parent_ingress_external_service_template": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceIngressExternalServiceTemplateEntryCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize IngressExternalServiceTemplateEntry object
    o := &vspk.IngressExternalServiceTemplateEntry{
        DSCP: d.Get("dscp").(string),
        Action: d.Get("action").(string),
        LocationType: d.Get("location_type").(string),
        EtherType: d.Get("ether_type").(string),
    }
    if attr, ok := d.GetOk("icmp_code"); ok {
        o.ICMPCode = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_type"); ok {
        o.ICMPType = attr.(string)
    }
    if attr, ok := d.GetOk("address_override"); ok {
        o.AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("redirect_external_service_end_point_id"); ok {
        o.RedirectExternalServiceEndPointID = attr.(string)
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
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("protocol"); ok {
        o.Protocol = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_id"); ok {
        o.AssociatedApplicationID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_object_id"); ok {
        o.AssociatedApplicationObjectID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_object_type"); ok {
        o.AssociatedApplicationObjectType = attr.(string)
    }
    if attr, ok := d.GetOk("associated_live_entity_id"); ok {
        o.AssociatedLiveEntityID = attr.(string)
    }
    if attr, ok := d.GetOk("stats_logging_enabled"); ok {
        o.StatsLoggingEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.IngressExternalServiceTemplate{ID: d.Get("parent_ingress_external_service_template").(string)}
    err := parent.CreateIngressExternalServiceTemplateEntry(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceIngressExternalServiceTemplateEntryRead(d, m)
}

func resourceIngressExternalServiceTemplateEntryRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IngressExternalServiceTemplateEntry{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("icmp_code", o.ICMPCode)
    d.Set("icmp_type", o.ICMPType)
    d.Set("dscp", o.DSCP)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("action", o.Action)
    d.Set("address_override", o.AddressOverride)
    d.Set("redirect_external_service_end_point_id", o.RedirectExternalServiceEndPointID)
    d.Set("description", o.Description)
    d.Set("destination_port", o.DestinationPort)
    d.Set("network_id", o.NetworkID)
    d.Set("network_type", o.NetworkType)
    d.Set("mirror_destination_id", o.MirrorDestinationID)
    d.Set("flow_logging_enabled", o.FlowLoggingEnabled)
    d.Set("entity_scope", o.EntityScope)
    d.Set("location_id", o.LocationID)
    d.Set("location_type", o.LocationType)
    d.Set("policy_state", o.PolicyState)
    d.Set("source_port", o.SourcePort)
    d.Set("priority", o.Priority)
    d.Set("protocol", o.Protocol)
    d.Set("associated_application_id", o.AssociatedApplicationID)
    d.Set("associated_application_object_id", o.AssociatedApplicationObjectID)
    d.Set("associated_application_object_type", o.AssociatedApplicationObjectType)
    d.Set("associated_live_entity_id", o.AssociatedLiveEntityID)
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

func resourceIngressExternalServiceTemplateEntryUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IngressExternalServiceTemplateEntry{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.DSCP = d.Get("dscp").(string)
    o.Action = d.Get("action").(string)
    o.LocationType = d.Get("location_type").(string)
    o.EtherType = d.Get("ether_type").(string)
    
    if attr, ok := d.GetOk("icmp_code"); ok {
        o.ICMPCode = attr.(string)
    }
    if attr, ok := d.GetOk("icmp_type"); ok {
        o.ICMPType = attr.(string)
    }
    if attr, ok := d.GetOk("address_override"); ok {
        o.AddressOverride = attr.(string)
    }
    if attr, ok := d.GetOk("redirect_external_service_end_point_id"); ok {
        o.RedirectExternalServiceEndPointID = attr.(string)
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
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("protocol"); ok {
        o.Protocol = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_id"); ok {
        o.AssociatedApplicationID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_object_id"); ok {
        o.AssociatedApplicationObjectID = attr.(string)
    }
    if attr, ok := d.GetOk("associated_application_object_type"); ok {
        o.AssociatedApplicationObjectType = attr.(string)
    }
    if attr, ok := d.GetOk("associated_live_entity_id"); ok {
        o.AssociatedLiveEntityID = attr.(string)
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

func resourceIngressExternalServiceTemplateEntryDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IngressExternalServiceTemplateEntry{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}