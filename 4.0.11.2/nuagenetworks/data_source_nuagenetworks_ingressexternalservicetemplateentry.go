package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11.2"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceIngressExternalServiceTemplateEntry() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceIngressExternalServiceTemplateEntryRead,
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
            "acl_template_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "icmp_code": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "icmp_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ipv6_address_override": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dscp": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "action": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address_override": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "redirect_external_service_end_point_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mirror_destination_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "flow_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "enterprise_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "location_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "location_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "policy_state": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "protocol": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_application_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_application_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_application_object_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_live_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "stats_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "stats_logging_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "ether_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_ingress_external_service_template": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceIngressExternalServiceTemplateEntryRead(d *schema.ResourceData, m interface{}) error {
    filteredIngressExternalServiceTemplateEntries := vspk.IngressExternalServiceTemplateEntriesList{}
    err := &bambou.Error{}
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
    parent := &vspk.IngressExternalServiceTemplate{ID: d.Get("parent_ingress_external_service_template").(string)}
    filteredIngressExternalServiceTemplateEntries, err = parent.IngressExternalServiceTemplateEntries(fetchFilter)
    if err != nil {
        return err
    }

    IngressExternalServiceTemplateEntry := &vspk.IngressExternalServiceTemplateEntry{}

    if len(filteredIngressExternalServiceTemplateEntries) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredIngressExternalServiceTemplateEntries) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    IngressExternalServiceTemplateEntry = filteredIngressExternalServiceTemplateEntries[0]

    d.Set("acl_template_name", IngressExternalServiceTemplateEntry.ACLTemplateName)
    d.Set("icmp_code", IngressExternalServiceTemplateEntry.ICMPCode)
    d.Set("icmp_type", IngressExternalServiceTemplateEntry.ICMPType)
    d.Set("ipv6_address_override", IngressExternalServiceTemplateEntry.IPv6AddressOverride)
    d.Set("dscp", IngressExternalServiceTemplateEntry.DSCP)
    d.Set("name", IngressExternalServiceTemplateEntry.Name)
    d.Set("last_updated_by", IngressExternalServiceTemplateEntry.LastUpdatedBy)
    d.Set("action", IngressExternalServiceTemplateEntry.Action)
    d.Set("address_override", IngressExternalServiceTemplateEntry.AddressOverride)
    d.Set("redirect_external_service_end_point_id", IngressExternalServiceTemplateEntry.RedirectExternalServiceEndPointID)
    d.Set("description", IngressExternalServiceTemplateEntry.Description)
    d.Set("destination_port", IngressExternalServiceTemplateEntry.DestinationPort)
    d.Set("network_id", IngressExternalServiceTemplateEntry.NetworkID)
    d.Set("network_type", IngressExternalServiceTemplateEntry.NetworkType)
    d.Set("mirror_destination_id", IngressExternalServiceTemplateEntry.MirrorDestinationID)
    d.Set("flow_logging_enabled", IngressExternalServiceTemplateEntry.FlowLoggingEnabled)
    d.Set("enterprise_name", IngressExternalServiceTemplateEntry.EnterpriseName)
    d.Set("entity_scope", IngressExternalServiceTemplateEntry.EntityScope)
    d.Set("location_id", IngressExternalServiceTemplateEntry.LocationID)
    d.Set("location_type", IngressExternalServiceTemplateEntry.LocationType)
    d.Set("policy_state", IngressExternalServiceTemplateEntry.PolicyState)
    d.Set("domain_name", IngressExternalServiceTemplateEntry.DomainName)
    d.Set("source_port", IngressExternalServiceTemplateEntry.SourcePort)
    d.Set("priority", IngressExternalServiceTemplateEntry.Priority)
    d.Set("protocol", IngressExternalServiceTemplateEntry.Protocol)
    d.Set("associated_application_id", IngressExternalServiceTemplateEntry.AssociatedApplicationID)
    d.Set("associated_application_object_id", IngressExternalServiceTemplateEntry.AssociatedApplicationObjectID)
    d.Set("associated_application_object_type", IngressExternalServiceTemplateEntry.AssociatedApplicationObjectType)
    d.Set("associated_live_entity_id", IngressExternalServiceTemplateEntry.AssociatedLiveEntityID)
    d.Set("stats_id", IngressExternalServiceTemplateEntry.StatsID)
    d.Set("stats_logging_enabled", IngressExternalServiceTemplateEntry.StatsLoggingEnabled)
    d.Set("ether_type", IngressExternalServiceTemplateEntry.EtherType)
    d.Set("external_id", IngressExternalServiceTemplateEntry.ExternalID)
    
    d.Set("id", IngressExternalServiceTemplateEntry.Identifier())
    d.Set("parent_id", IngressExternalServiceTemplateEntry.ParentID)
    d.Set("parent_type", IngressExternalServiceTemplateEntry.ParentType)
    d.Set("owner", IngressExternalServiceTemplateEntry.Owner)

    d.SetId(IngressExternalServiceTemplateEntry.Identifier())
    
    return nil
}