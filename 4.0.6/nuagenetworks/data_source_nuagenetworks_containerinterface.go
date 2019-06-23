package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.6"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceContainerInterface() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceContainerInterfaceRead,
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
            "mac": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vport_name": &schema.Schema{
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
            "gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "tier_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "endpoint_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "policy_decision_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zone_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zone_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "container_uuid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_floating_ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "attached_network_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "attached_network_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multi_nic_vport_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_subnet", "parent_zone", "parent_vport", "parent_container"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_subnet", "parent_zone", "parent_vport", "parent_container"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_zone", "parent_vport", "parent_container"},
            },
            "parent_zone": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_vport", "parent_container"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_zone", "parent_container"},
            },
            "parent_container": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_subnet", "parent_zone", "parent_vport"},
            },
        },
    }
}


func dataSourceContainerInterfaceRead(d *schema.ResourceData, m interface{}) error {
    filteredContainerInterfaces := vspk.ContainerInterfacesList{}
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
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_zone"); ok {
        parent := &vspk.Zone{ID: attr.(string)}
        filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_container"); ok {
        parent := &vspk.Container{ID: attr.(string)}
        filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
        if err != nil {
            return err
        }
    }

    ContainerInterface := &vspk.ContainerInterface{}

    if len(filteredContainerInterfaces) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredContainerInterfaces) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    ContainerInterface = filteredContainerInterfaces[0]

    d.Set("mac", ContainerInterface.MAC)
    d.Set("ip_address", ContainerInterface.IPAddress)
    d.Set("vport_id", ContainerInterface.VPortID)
    d.Set("vport_name", ContainerInterface.VPortName)
    d.Set("name", ContainerInterface.Name)
    d.Set("last_updated_by", ContainerInterface.LastUpdatedBy)
    d.Set("gateway", ContainerInterface.Gateway)
    d.Set("netmask", ContainerInterface.Netmask)
    d.Set("network_id", ContainerInterface.NetworkID)
    d.Set("network_name", ContainerInterface.NetworkName)
    d.Set("tier_id", ContainerInterface.TierID)
    d.Set("endpoint_id", ContainerInterface.EndpointID)
    d.Set("entity_scope", ContainerInterface.EntityScope)
    d.Set("policy_decision_id", ContainerInterface.PolicyDecisionID)
    d.Set("domain_id", ContainerInterface.DomainID)
    d.Set("domain_name", ContainerInterface.DomainName)
    d.Set("zone_id", ContainerInterface.ZoneID)
    d.Set("zone_name", ContainerInterface.ZoneName)
    d.Set("container_uuid", ContainerInterface.ContainerUUID)
    d.Set("associated_floating_ip_address", ContainerInterface.AssociatedFloatingIPAddress)
    d.Set("attached_network_id", ContainerInterface.AttachedNetworkID)
    d.Set("attached_network_type", ContainerInterface.AttachedNetworkType)
    d.Set("multi_nic_vport_name", ContainerInterface.MultiNICVPortName)
    d.Set("external_id", ContainerInterface.ExternalID)
    
    d.Set("id", ContainerInterface.Identifier())
    d.Set("parent_id", ContainerInterface.ParentID)
    d.Set("parent_type", ContainerInterface.ParentType)
    d.Set("owner", ContainerInterface.Owner)

    d.SetId(ContainerInterface.Identifier())
    
    return nil
}