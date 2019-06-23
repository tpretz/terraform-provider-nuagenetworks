package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.7"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVCenterCluster() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVCenterClusterRead,
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
            "vrs_configuration_time_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "v_require_nuage_metadata": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "managed_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "data_dns1": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "data_dns2": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "data_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "data_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "datapath_sync_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "scope": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "secondary_nuage_controller": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "deleted_from_vcenter_data_center": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "generic_split_activation": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "separate_data_network": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "personality": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_mirror_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "metadata_server_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "metadata_server_listen_port": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "metadata_server_port": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "metadata_service_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "network_uplink_interface": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_uplink_interface_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_uplink_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_uplink_interface_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nfs_log_server": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nfs_mount_path": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mgmt_dns1": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mgmt_dns2": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mgmt_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mgmt_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dhcp_relay_server": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mirror_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "site_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "allow_data_dhcp": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "allow_mgmt_dhcp": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "flow_eviction_threshold": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vm_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "portgroup_metadata": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "nova_client_version": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "nova_metadata_service_auth_url": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_metadata_service_endpoint": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_metadata_service_password": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_metadata_service_tenant": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_metadata_service_username": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_metadata_shared_secret": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_region_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "upgrade_package_password": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "upgrade_package_url": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "upgrade_package_username": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "upgrade_script_time_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "primary_nuage_controller": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vrs_password": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vrs_user_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "assoc_vcenter_data_center_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "assoc_vcenter_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "static_route": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "static_route_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "static_route_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ntp_server1": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ntp_server2": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mtu": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "multi_vmssupport": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "multicast_receive_interface": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_receive_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_receive_interface_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_receive_range": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_send_interface": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_send_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_send_interface_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_source_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "customized_script_url": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ovf_url": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_vcenter_data_center": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceVCenterClusterRead(d *schema.ResourceData, m interface{}) error {
    filteredVCenterClusters := vspk.VCenterClustersList{}
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
    parent := &vspk.VCenterDataCenter{ID: d.Get("parent_vcenter_data_center").(string)}
    filteredVCenterClusters, err = parent.VCenterClusters(fetchFilter)
    if err != nil {
        return err
    }

    VCenterCluster := &vspk.VCenterCluster{}

    if len(filteredVCenterClusters) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVCenterClusters) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VCenterCluster = filteredVCenterClusters[0]

    d.Set("vrs_configuration_time_limit", VCenterCluster.VRSConfigurationTimeLimit)
    d.Set("v_require_nuage_metadata", VCenterCluster.VRequireNuageMetadata)
    d.Set("name", VCenterCluster.Name)
    d.Set("managed_object_id", VCenterCluster.ManagedObjectID)
    d.Set("last_updated_by", VCenterCluster.LastUpdatedBy)
    d.Set("data_dns1", VCenterCluster.DataDNS1)
    d.Set("data_dns2", VCenterCluster.DataDNS2)
    d.Set("data_gateway", VCenterCluster.DataGateway)
    d.Set("data_network_portgroup", VCenterCluster.DataNetworkPortgroup)
    d.Set("datapath_sync_timeout", VCenterCluster.DatapathSyncTimeout)
    d.Set("scope", VCenterCluster.Scope)
    d.Set("secondary_nuage_controller", VCenterCluster.SecondaryNuageController)
    d.Set("deleted_from_vcenter_data_center", VCenterCluster.DeletedFromVCenterDataCenter)
    d.Set("generic_split_activation", VCenterCluster.GenericSplitActivation)
    d.Set("separate_data_network", VCenterCluster.SeparateDataNetwork)
    d.Set("personality", VCenterCluster.Personality)
    d.Set("description", VCenterCluster.Description)
    d.Set("destination_mirror_port", VCenterCluster.DestinationMirrorPort)
    d.Set("metadata_server_ip", VCenterCluster.MetadataServerIP)
    d.Set("metadata_server_listen_port", VCenterCluster.MetadataServerListenPort)
    d.Set("metadata_server_port", VCenterCluster.MetadataServerPort)
    d.Set("metadata_service_enabled", VCenterCluster.MetadataServiceEnabled)
    d.Set("network_uplink_interface", VCenterCluster.NetworkUplinkInterface)
    d.Set("network_uplink_interface_gateway", VCenterCluster.NetworkUplinkInterfaceGateway)
    d.Set("network_uplink_interface_ip", VCenterCluster.NetworkUplinkInterfaceIp)
    d.Set("network_uplink_interface_netmask", VCenterCluster.NetworkUplinkInterfaceNetmask)
    d.Set("nfs_log_server", VCenterCluster.NfsLogServer)
    d.Set("nfs_mount_path", VCenterCluster.NfsMountPath)
    d.Set("mgmt_dns1", VCenterCluster.MgmtDNS1)
    d.Set("mgmt_dns2", VCenterCluster.MgmtDNS2)
    d.Set("mgmt_gateway", VCenterCluster.MgmtGateway)
    d.Set("mgmt_network_portgroup", VCenterCluster.MgmtNetworkPortgroup)
    d.Set("dhcp_relay_server", VCenterCluster.DhcpRelayServer)
    d.Set("mirror_network_portgroup", VCenterCluster.MirrorNetworkPortgroup)
    d.Set("site_id", VCenterCluster.SiteId)
    d.Set("allow_data_dhcp", VCenterCluster.AllowDataDHCP)
    d.Set("allow_mgmt_dhcp", VCenterCluster.AllowMgmtDHCP)
    d.Set("flow_eviction_threshold", VCenterCluster.FlowEvictionThreshold)
    d.Set("vm_network_portgroup", VCenterCluster.VmNetworkPortgroup)
    d.Set("entity_scope", VCenterCluster.EntityScope)
    d.Set("portgroup_metadata", VCenterCluster.PortgroupMetadata)
    d.Set("nova_client_version", VCenterCluster.NovaClientVersion)
    d.Set("nova_metadata_service_auth_url", VCenterCluster.NovaMetadataServiceAuthUrl)
    d.Set("nova_metadata_service_endpoint", VCenterCluster.NovaMetadataServiceEndpoint)
    d.Set("nova_metadata_service_password", VCenterCluster.NovaMetadataServicePassword)
    d.Set("nova_metadata_service_tenant", VCenterCluster.NovaMetadataServiceTenant)
    d.Set("nova_metadata_service_username", VCenterCluster.NovaMetadataServiceUsername)
    d.Set("nova_metadata_shared_secret", VCenterCluster.NovaMetadataSharedSecret)
    d.Set("nova_region_name", VCenterCluster.NovaRegionName)
    d.Set("upgrade_package_password", VCenterCluster.UpgradePackagePassword)
    d.Set("upgrade_package_url", VCenterCluster.UpgradePackageURL)
    d.Set("upgrade_package_username", VCenterCluster.UpgradePackageUsername)
    d.Set("upgrade_script_time_limit", VCenterCluster.UpgradeScriptTimeLimit)
    d.Set("primary_nuage_controller", VCenterCluster.PrimaryNuageController)
    d.Set("vrs_password", VCenterCluster.VrsPassword)
    d.Set("vrs_user_name", VCenterCluster.VrsUserName)
    d.Set("assoc_vcenter_data_center_id", VCenterCluster.AssocVCenterDataCenterID)
    d.Set("assoc_vcenter_id", VCenterCluster.AssocVCenterID)
    d.Set("static_route", VCenterCluster.StaticRoute)
    d.Set("static_route_gateway", VCenterCluster.StaticRouteGateway)
    d.Set("static_route_netmask", VCenterCluster.StaticRouteNetmask)
    d.Set("ntp_server1", VCenterCluster.NtpServer1)
    d.Set("ntp_server2", VCenterCluster.NtpServer2)
    d.Set("mtu", VCenterCluster.Mtu)
    d.Set("multi_vmssupport", VCenterCluster.MultiVMSsupport)
    d.Set("multicast_receive_interface", VCenterCluster.MulticastReceiveInterface)
    d.Set("multicast_receive_interface_ip", VCenterCluster.MulticastReceiveInterfaceIP)
    d.Set("multicast_receive_interface_netmask", VCenterCluster.MulticastReceiveInterfaceNetmask)
    d.Set("multicast_receive_range", VCenterCluster.MulticastReceiveRange)
    d.Set("multicast_send_interface", VCenterCluster.MulticastSendInterface)
    d.Set("multicast_send_interface_ip", VCenterCluster.MulticastSendInterfaceIP)
    d.Set("multicast_send_interface_netmask", VCenterCluster.MulticastSendInterfaceNetmask)
    d.Set("multicast_source_portgroup", VCenterCluster.MulticastSourcePortgroup)
    d.Set("customized_script_url", VCenterCluster.CustomizedScriptURL)
    d.Set("ovf_url", VCenterCluster.OvfURL)
    d.Set("external_id", VCenterCluster.ExternalID)
    
    d.Set("id", VCenterCluster.Identifier())
    d.Set("parent_id", VCenterCluster.ParentID)
    d.Set("parent_type", VCenterCluster.ParentType)
    d.Set("owner", VCenterCluster.Owner)

    d.SetId(VCenterCluster.Identifier())
    
    return nil
}