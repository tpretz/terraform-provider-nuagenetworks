package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceUplinkConnection() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceUplinkConnectionRead,
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
            "pat_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "dns_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dns_address_v6": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "password": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_v6": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address_family": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address_v6": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "advertisement_criteria": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "secondary_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vlan": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "underlay_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "underlay_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "inherited": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "installer_managed": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "interface_connection_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mode": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "role": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "role_order": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "port_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "download_rate_limit": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "uplink_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "username": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "assoc_underlay_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_bgp_neighbor_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_underlay_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "auxiliary_link": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_vlan_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan", "parent_ns_gateway"},
            },
            "parent_vlan": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan_template", "parent_ns_gateway"},
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan_template", "parent_vlan"},
            },
        },
    }
}


func dataSourceUplinkConnectionRead(d *schema.ResourceData, m interface{}) error {
    filteredUplinkConnections := vspk.UplinkConnectionsList{}
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
    if attr, ok := d.GetOk("parent_vlan_template"); ok {
        parent := &vspk.VLANTemplate{ID: attr.(string)}
        filteredUplinkConnections, err = parent.UplinkConnections(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vlan"); ok {
        parent := &vspk.VLAN{ID: attr.(string)}
        filteredUplinkConnections, err = parent.UplinkConnections(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ns_gateway"); ok {
        parent := &vspk.NSGateway{ID: attr.(string)}
        filteredUplinkConnections, err = parent.UplinkConnections(fetchFilter)
        if err != nil {
            return err
        }
    }

    UplinkConnection := &vspk.UplinkConnection{}

    if len(filteredUplinkConnections) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredUplinkConnections) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    UplinkConnection = filteredUplinkConnections[0]

    d.Set("pat_enabled", UplinkConnection.PATEnabled)
    d.Set("dns_address", UplinkConnection.DNSAddress)
    d.Set("dns_address_v6", UplinkConnection.DNSAddressV6)
    d.Set("password", UplinkConnection.Password)
    d.Set("last_updated_by", UplinkConnection.LastUpdatedBy)
    d.Set("gateway", UplinkConnection.Gateway)
    d.Set("gateway_v6", UplinkConnection.GatewayV6)
    d.Set("address", UplinkConnection.Address)
    d.Set("address_family", UplinkConnection.AddressFamily)
    d.Set("address_v6", UplinkConnection.AddressV6)
    d.Set("advertisement_criteria", UplinkConnection.AdvertisementCriteria)
    d.Set("secondary_address", UplinkConnection.SecondaryAddress)
    d.Set("netmask", UplinkConnection.Netmask)
    d.Set("vlan", UplinkConnection.Vlan)
    d.Set("underlay_enabled", UplinkConnection.UnderlayEnabled)
    d.Set("underlay_id", UplinkConnection.UnderlayID)
    d.Set("inherited", UplinkConnection.Inherited)
    d.Set("installer_managed", UplinkConnection.InstallerManaged)
    d.Set("interface_connection_type", UplinkConnection.InterfaceConnectionType)
    d.Set("entity_scope", UplinkConnection.EntityScope)
    d.Set("mode", UplinkConnection.Mode)
    d.Set("role", UplinkConnection.Role)
    d.Set("role_order", UplinkConnection.RoleOrder)
    d.Set("port_name", UplinkConnection.PortName)
    d.Set("download_rate_limit", UplinkConnection.DownloadRateLimit)
    d.Set("uplink_id", UplinkConnection.UplinkID)
    d.Set("username", UplinkConnection.Username)
    d.Set("assoc_underlay_id", UplinkConnection.AssocUnderlayID)
    d.Set("associated_bgp_neighbor_id", UplinkConnection.AssociatedBGPNeighborID)
    d.Set("associated_underlay_name", UplinkConnection.AssociatedUnderlayName)
    d.Set("auxiliary_link", UplinkConnection.AuxiliaryLink)
    d.Set("external_id", UplinkConnection.ExternalID)
    
    d.Set("id", UplinkConnection.Identifier())
    d.Set("parent_id", UplinkConnection.ParentID)
    d.Set("parent_type", UplinkConnection.ParentType)
    d.Set("owner", UplinkConnection.Owner)

    d.SetId(UplinkConnection.Identifier())
    
    return nil
}