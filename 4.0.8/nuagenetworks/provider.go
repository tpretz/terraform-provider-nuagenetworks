package nuagenetworks

import (
    "errors"
    "log"
    "crypto/tls"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.8"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func Provider() terraform.ResourceProvider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "enterprise": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                DefaultFunc: schema.EnvDefaultFunc("VSD_ORGANIZATION", "csp"),
            },
            "vsd_endpoint": &schema.Schema{
                Type:        schema.TypeString,
                Required:    true,
                DefaultFunc: schema.EnvDefaultFunc("VSD_ENDPOINT", nil),
            },
            "username": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                DefaultFunc: schema.EnvDefaultFunc("VSD_USERNAME", "csproot"),
            },
            "password": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                DefaultFunc: schema.EnvDefaultFunc("VSD_PASSWORD", "csproot"),
            },
            "certificate_path": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                DefaultFunc: schema.EnvDefaultFunc("VSD_CERTIFICATE_PATH", nil),
            },
            "key_path": &schema.Schema{
                Type:        schema.TypeString,
                Optional:    true,
                DefaultFunc: schema.EnvDefaultFunc("VSD_KEY_PATH", nil),
            },
        },
        ConfigureFunc: providerConfigure,
        DataSourcesMap: map[string]*schema.Resource{
            "nuagenetworks_performance_monitor": dataSourcePerformanceMonitor(),
            "nuagenetworks_vm_resync": dataSourceVMResync(),
            "nuagenetworks_shared_network_resource": dataSourceSharedNetworkResource(),
            "nuagenetworks_enterprise_secured_data": dataSourceEnterpriseSecuredData(),
            "nuagenetworks_statistics": dataSourceStatistics(),
            "nuagenetworks_ike_gateway": dataSourceIKEGateway(),
            "nuagenetworks_domain_fip_acl_template_entry": dataSourceDomainFIPAclTemplateEntry(),
            "nuagenetworks_dscp_forwarding_class_table": dataSourceDSCPForwardingClassTable(),
            "nuagenetworks_bootstrap": dataSourceBootstrap(),
            "nuagenetworks_auto_discover_cluster": dataSourceAutoDiscoverCluster(),
            "nuagenetworks_domain": dataSourceDomain(),
            "nuagenetworks_tca": dataSourceTCA(),
            "nuagenetworks_enterprise_profile": dataSourceEnterpriseProfile(),
            "nuagenetworks_rate_limiter": dataSourceRateLimiter(),
            "nuagenetworks_license_status": dataSourceLicenseStatus(),
            "nuagenetworks_cloud_mgmt_system": dataSourceCloudMgmtSystem(),
            "nuagenetworks_overlay_address_pool": dataSourceOverlayAddressPool(),
            "nuagenetworks_vm_interface": dataSourceVMInterface(),
            "nuagenetworks_auto_discovered_gateway": dataSourceAutoDiscoveredGateway(),
            "nuagenetworks_host_interface": dataSourceHostInterface(),
            "nuagenetworks_ike_gateway_profile": dataSourceIKEGatewayProfile(),
            "nuagenetworks_vcenter_hypervisor": dataSourceVCenterHypervisor(),
            "nuagenetworks_ike_subnet": dataSourceIKESubnet(),
            "nuagenetworks_qos": dataSourceQOS(),
            "nuagenetworks_global_metadata": dataSourceGlobalMetadata(),
            "nuagenetworks_external_app_service": dataSourceExternalAppService(),
            "nuagenetworks_vrs_address_range": dataSourceVRSAddressRange(),
            "nuagenetworks_ingress_external_service_template_entry": dataSourceIngressExternalServiceTemplateEntry(),
            "nuagenetworks_nsg_info": dataSourceNSGInfo(),
            "nuagenetworks_permission": dataSourcePermission(),
            "nuagenetworks_container_interface": dataSourceContainerInterface(),
            "nuagenetworks_event_log": dataSourceEventLog(),
            "nuagenetworks_vsp": dataSourceVSP(),
            "nuagenetworks_vlan_template": dataSourceVLANTemplate(),
            "nuagenetworks_pat_mapper": dataSourcePATMapper(),
            "nuagenetworks_mirror_destination": dataSourceMirrorDestination(),
            "nuagenetworks_group_key_encryption_profile": dataSourceGroupKeyEncryptionProfile(),
            "nuagenetworks_l2_domain": dataSourceL2Domain(),
            "nuagenetworks_network_layout": dataSourceNetworkLayout(),
            "nuagenetworks_application_service": dataSourceApplicationService(),
            "nuagenetworks_ns_port_template": dataSourceNSPortTemplate(),
            "nuagenetworks_alarm": dataSourceAlarm(),
            "nuagenetworks_vpn_connection": dataSourceVPNConnection(),
            "nuagenetworks_key_server_monitor_seed": dataSourceKeyServerMonitorSeed(),
            "nuagenetworks_vcenter_data_center": dataSourceVCenterDataCenter(),
            "nuagenetworks_firewall_rule": dataSourceFirewallRule(),
            "nuagenetworks_ike_gateway_config": dataSourceIKEGatewayConfig(),
            "nuagenetworks_static_route": dataSourceStaticRoute(),
            "nuagenetworks_bridge_interface": dataSourceBridgeInterface(),
            "nuagenetworks_redirection_target_template": dataSourceRedirectionTargetTemplate(),
            "nuagenetworks_dscp_forwarding_class_mapping": dataSourceDSCPForwardingClassMapping(),
            "nuagenetworks_aggregate_metadata": dataSourceAggregateMetadata(),
            "nuagenetworks_l2_domain_template": dataSourceL2DomainTemplate(),
            "nuagenetworks_subnet": dataSourceSubnet(),
            "nuagenetworks_ns_port": dataSourceNSPort(),
            "nuagenetworks_infrastructure_vsc_profile": dataSourceInfrastructureVscProfile(),
            "nuagenetworks_port_mapping": dataSourcePortMapping(),
            "nuagenetworks_bulk_statistics": dataSourceBulkStatistics(),
            "nuagenetworks_infrastructure_config": dataSourceInfrastructureConfig(),
            "nuagenetworks_policy_decision": dataSourcePolicyDecision(),
            "nuagenetworks_user": dataSourceUser(),
            "nuagenetworks_duc_group": dataSourceDUCGroup(),
            "nuagenetworks_license": dataSourceLicense(),
            "nuagenetworks_vsd_component": dataSourceVSDComponent(),
            "nuagenetworks_ltestatistics": dataSourceLtestatistics(),
            "nuagenetworks_via": dataSourceVia(),
            "nuagenetworks_vrs_metrics": dataSourceVRSMetrics(),
            "nuagenetworks_dhcp_option": dataSourceDHCPOption(),
            "nuagenetworks_enterprise_network": dataSourceEnterpriseNetwork(),
            "nuagenetworks_ns_redundant_gateway_group": dataSourceNSRedundantGatewayGroup(),
            "nuagenetworks_zone_template": dataSourceZoneTemplate(),
            "nuagenetworks_key_server_member": dataSourceKeyServerMember(),
            "nuagenetworks_key_server_monitor_sek": dataSourceKeyServerMonitorSEK(),
            "nuagenetworks_next_hop_address": dataSourceNextHopAddress(),
            "nuagenetworks_network_performance_binding": dataSourceNetworkPerformanceBinding(),
            "nuagenetworks_connectionendpoint": dataSourceConnectionendpoint(),
            "nuagenetworks_gateway_security": dataSourceGatewaySecurity(),
            "nuagenetworks_custom_property": dataSourceCustomProperty(),
            "nuagenetworks_routing_policy": dataSourceRoutingPolicy(),
            "nuagenetworks_application_binding": dataSourceApplicationBinding(),
            "nuagenetworks_gateway_template": dataSourceGatewayTemplate(),
            "nuagenetworks_underlay": dataSourceUnderlay(),
            "nuagenetworks_redirection_target": dataSourceRedirectionTarget(),
            "nuagenetworks_ldap_configuration": dataSourceLDAPConfiguration(),
            "nuagenetworks_enterprise": dataSourceEnterprise(),
            "nuagenetworks_wan_service": dataSourceWANService(),
            "nuagenetworks_site_info": dataSourceSiteInfo(),
            "nuagenetworks_avatar": dataSourceAvatar(),
            "nuagenetworks_key_server_monitor_encrypted_seed": dataSourceKeyServerMonitorEncryptedSeed(),
            "nuagenetworks_vsg_redundant_port": dataSourceVsgRedundantPort(),
            "nuagenetworks_policy_group_template": dataSourcePolicyGroupTemplate(),
            "nuagenetworks_application": dataSourceApplication(),
            "nuagenetworks_egress_acl_template": dataSourceEgressACLTemplate(),
            "nuagenetworks_metadata": dataSourceMetadata(),
            "nuagenetworks_ingress_adv_fwd_entry_template": dataSourceIngressAdvFwdEntryTemplate(),
            "nuagenetworks_redundant_port": dataSourceRedundantPort(),
            "nuagenetworks_link": dataSourceLink(),
            "nuagenetworks_ssh_key": dataSourceSSHKey(),
            "nuagenetworks_demarcation_service": dataSourceDemarcationService(),
            "nuagenetworks_public_network_macro": dataSourcePublicNetworkMacro(),
            "nuagenetworks_ingress_external_service_template": dataSourceIngressExternalServiceTemplate(),
            "nuagenetworks_bgp_peer": dataSourceBGPPeer(),
            "nuagenetworks_vport_mirror": dataSourceVPortMirror(),
            "nuagenetworks_floating_ipacl_template_entry": dataSourceFloatingIPACLTemplateEntry(),
            "nuagenetworks_flow_security_policy": dataSourceFlowSecurityPolicy(),
            "nuagenetworks_network_performance_measurement": dataSourceNetworkPerformanceMeasurement(),
            "nuagenetworks_external_service": dataSourceExternalService(),
            "nuagenetworks_ingress_acl_entry_template": dataSourceIngressACLEntryTemplate(),
            "nuagenetworks_policy_group": dataSourcePolicyGroup(),
            "nuagenetworks_vcenter_cluster": dataSourceVCenterCluster(),
            "nuagenetworks_enterprise_security": dataSourceEnterpriseSecurity(),
            "nuagenetworks_duc_group_binding": dataSourceDUCGroupBinding(),
            "nuagenetworks_zone": dataSourceZone(),
            "nuagenetworks_multi_cast_list": dataSourceMultiCastList(),
            "nuagenetworks_container_resync": dataSourceContainerResync(),
            "nuagenetworks_system_config": dataSourceSystemConfig(),
            "nuagenetworks_multi_cast_channel_map": dataSourceMultiCastChannelMap(),
            "nuagenetworks_bgp_neighbor": dataSourceBGPNeighbor(),
            "nuagenetworks_vlan": dataSourceVLAN(),
            "nuagenetworks_auto_discover_hypervisor_from_cluster": dataSourceAutoDiscoverHypervisorFromCluster(),
            "nuagenetworks_statistics_policy": dataSourceStatisticsPolicy(),
            "nuagenetworks_zfb_request": dataSourceZFBRequest(),
            "nuagenetworks_job": dataSourceJob(),
            "nuagenetworks_network_macro_group": dataSourceNetworkMacroGroup(),
            "nuagenetworks_ns_gateway_template": dataSourceNSGatewayTemplate(),
            "nuagenetworks_patnat_pool": dataSourcePATNATPool(),
            "nuagenetworks_vport": dataSourceVPort(),
            "nuagenetworks_all_alarm": dataSourceAllAlarm(),
            "nuagenetworks_applicationperformancemanagementbinding": dataSourceApplicationperformancemanagementbinding(),
            "nuagenetworks_vsc": dataSourceVSC(),
            "nuagenetworks_l7applicationsignature": dataSourceL7applicationsignature(),
            "nuagenetworks_hsc": dataSourceHSC(),
            "nuagenetworks_vrs": dataSourceVRS(),
            "nuagenetworks_egress_qos_policy": dataSourceEgressQOSPolicy(),
            "nuagenetworks_group": dataSourceGroup(),
            "nuagenetworks_address_range": dataSourceAddressRange(),
            "nuagenetworks_bgp_profile": dataSourceBGPProfile(),
            "nuagenetworks_ingress_adv_fwd_template": dataSourceIngressAdvFwdTemplate(),
            "nuagenetworks_location": dataSourceLocation(),
            "nuagenetworks_egress_acl_entry_template": dataSourceEgressACLEntryTemplate(),
            "nuagenetworks_tier": dataSourceTier(),
            "nuagenetworks_ikepsk": dataSourceIKEPSK(),
            "nuagenetworks_port_template": dataSourcePortTemplate(),
            "nuagenetworks_multi_nic_vport": dataSourceMultiNICVPort(),
            "nuagenetworks_ike_encryptionprofile": dataSourceIKEEncryptionprofile(),
            "nuagenetworks_floating_ip": dataSourceFloatingIp(),
            "nuagenetworks_flow_forwarding_policy": dataSourceFlowForwardingPolicy(),
            "nuagenetworks_vcenter_eam_config": dataSourceVCenterEAMConfig(),
            "nuagenetworks_ns_gateway": dataSourceNSGateway(),
            "nuagenetworks_enterprise_permission": dataSourceEnterprisePermission(),
            "nuagenetworks_multi_cast_range": dataSourceMultiCastRange(),
            "nuagenetworks_domain_template": dataSourceDomainTemplate(),
            "nuagenetworks_monitorscope": dataSourceMonitorscope(),
            "nuagenetworks_nsg_group": dataSourceNSGGroup(),
            "nuagenetworks_end_point": dataSourceEndPoint(),
            "nuagenetworks_vcenter": dataSourceVCenter(),
            "nuagenetworks_autodiscovereddatacenter": dataSourceAutodiscovereddatacenter(),
            "nuagenetworks_overlay_patnat_entry": dataSourceOverlayPATNATEntry(),
            "nuagenetworks_infrastructure_gateway_profile": dataSourceInfrastructureGatewayProfile(),
            "nuagenetworks_ingress_acl_template": dataSourceIngressACLTemplate(),
            "nuagenetworks_vrs_redeploymentpolicy": dataSourceVRSRedeploymentpolicy(),
            "nuagenetworks_vcenter_vrs_config": dataSourceVCenterVRSConfig(),
            "nuagenetworks_flow": dataSourceFlow(),
            "nuagenetworks_key_server_monitor": dataSourceKeyServerMonitor(),
            "nuagenetworks_monitoring_port": dataSourceMonitoringPort(),
            "nuagenetworks_uplink_rd": dataSourceUplinkRD(),
            "nuagenetworks_floating_ipacl_template": dataSourceFloatingIPACLTemplate(),
            "nuagenetworks_uplink_connection": dataSourceUplinkConnection(),
            "nuagenetworks_zfb_auto_assignment": dataSourceZFBAutoAssignment(),
            "nuagenetworks_port": dataSourcePort(),
            "nuagenetworks_redundancy_group": dataSourceRedundancyGroup(),
            "nuagenetworks_patip_entry": dataSourcePATIPEntry(),
            "nuagenetworks_metadata_tag": dataSourceMetadataTag(),
            "nuagenetworks_vm": dataSourceVM(),
            "nuagenetworks_container": dataSourceContainer(),
            "nuagenetworks_virtual_ip": dataSourceVirtualIP(),
            "nuagenetworks_applicationperformancemanagement": dataSourceApplicationperformancemanagement(),
            "nuagenetworks_infrastructure_access_profile": dataSourceInfrastructureAccessProfile(),
            "nuagenetworks_subnet_template": dataSourceSubnetTemplate(),
            "nuagenetworks_br_connection": dataSourceBRConnection(),
            "nuagenetworks_ip_reservation": dataSourceIPReservation(),
            "nuagenetworks_firewall_acl": dataSourceFirewallAcl(),
            "nuagenetworks_stats_collector_info": dataSourceStatsCollectorInfo(),
            "nuagenetworks_domain_fip_acl_template": dataSourceDomainFIPAclTemplate(),
            "nuagenetworks_address_map": dataSourceAddressMap(),
            "nuagenetworks_gateway": dataSourceGateway(),
            "nuagenetworks_ike_gateway_connection": dataSourceIKEGatewayConnection(),
            "nuagenetworks_gateway_secured_data": dataSourceGatewaySecuredData(),
            "nuagenetworks_ike_certificate": dataSourceIKECertificate(),
            "nuagenetworks_vsd": dataSourceVSD(),
            "nuagenetworks_nat_map_entry": dataSourceNATMapEntry(),
        },
        ResourcesMap: map[string]*schema.Resource{
            "nuagenetworks_performance_monitor": resourcePerformanceMonitor(),
            "nuagenetworks_vm_resync": resourceVMResync(),
            "nuagenetworks_shared_network_resource": resourceSharedNetworkResource(),
            "nuagenetworks_enterprise_secured_data": resourceEnterpriseSecuredData(),
            "nuagenetworks_ike_gateway": resourceIKEGateway(),
            "nuagenetworks_domain_fip_acl_template_entry": resourceDomainFIPAclTemplateEntry(),
            "nuagenetworks_dscp_forwarding_class_table": resourceDSCPForwardingClassTable(),
            "nuagenetworks_domain": resourceDomain(),
            "nuagenetworks_tca": resourceTCA(),
            "nuagenetworks_enterprise_profile": resourceEnterpriseProfile(),
            "nuagenetworks_rate_limiter": resourceRateLimiter(),
            "nuagenetworks_cloud_mgmt_system": resourceCloudMgmtSystem(),
            "nuagenetworks_overlay_address_pool": resourceOverlayAddressPool(),
            "nuagenetworks_vm_interface": resourceVMInterface(),
            "nuagenetworks_host_interface": resourceHostInterface(),
            "nuagenetworks_ike_gateway_profile": resourceIKEGatewayProfile(),
            "nuagenetworks_vcenter_hypervisor": resourceVCenterHypervisor(),
            "nuagenetworks_ike_subnet": resourceIKESubnet(),
            "nuagenetworks_qos": resourceQOS(),
            "nuagenetworks_global_metadata": resourceGlobalMetadata(),
            "nuagenetworks_external_app_service": resourceExternalAppService(),
            "nuagenetworks_vrs_address_range": resourceVRSAddressRange(),
            "nuagenetworks_ingress_external_service_template_entry": resourceIngressExternalServiceTemplateEntry(),
            "nuagenetworks_permission": resourcePermission(),
            "nuagenetworks_container_interface": resourceContainerInterface(),
            "nuagenetworks_bootstrap_activation": resourceBootstrapActivation(),
            "nuagenetworks_vlan_template": resourceVLANTemplate(),
            "nuagenetworks_pat_mapper": resourcePATMapper(),
            "nuagenetworks_mirror_destination": resourceMirrorDestination(),
            "nuagenetworks_l2_domain": resourceL2Domain(),
            "nuagenetworks_application_service": resourceApplicationService(),
            "nuagenetworks_ns_port_template": resourceNSPortTemplate(),
            "nuagenetworks_alarm": resourceAlarm(),
            "nuagenetworks_vpn_connection": resourceVPNConnection(),
            "nuagenetworks_key_server_monitor_seed": resourceKeyServerMonitorSeed(),
            "nuagenetworks_vcenter_data_center": resourceVCenterDataCenter(),
            "nuagenetworks_firewall_rule": resourceFirewallRule(),
            "nuagenetworks_static_route": resourceStaticRoute(),
            "nuagenetworks_bridge_interface": resourceBridgeInterface(),
            "nuagenetworks_redirection_target_template": resourceRedirectionTargetTemplate(),
            "nuagenetworks_dscp_forwarding_class_mapping": resourceDSCPForwardingClassMapping(),
            "nuagenetworks_l2_domain_template": resourceL2DomainTemplate(),
            "nuagenetworks_subnet": resourceSubnet(),
            "nuagenetworks_ns_port": resourceNSPort(),
            "nuagenetworks_infrastructure_vsc_profile": resourceInfrastructureVscProfile(),
            "nuagenetworks_user": resourceUser(),
            "nuagenetworks_duc_group": resourceDUCGroup(),
            "nuagenetworks_license": resourceLicense(),
            "nuagenetworks_ltestatistics": resourceLtestatistics(),
            "nuagenetworks_dhcp_option": resourceDHCPOption(),
            "nuagenetworks_enterprise_network": resourceEnterpriseNetwork(),
            "nuagenetworks_ns_redundant_gateway_group": resourceNSRedundantGatewayGroup(),
            "nuagenetworks_zone_template": resourceZoneTemplate(),
            "nuagenetworks_key_server_member": resourceKeyServerMember(),
            "nuagenetworks_key_server_monitor_sek": resourceKeyServerMonitorSEK(),
            "nuagenetworks_next_hop_address": resourceNextHopAddress(),
            "nuagenetworks_network_performance_binding": resourceNetworkPerformanceBinding(),
            "nuagenetworks_custom_property": resourceCustomProperty(),
            "nuagenetworks_routing_policy": resourceRoutingPolicy(),
            "nuagenetworks_application_binding": resourceApplicationBinding(),
            "nuagenetworks_gateway_template": resourceGatewayTemplate(),
            "nuagenetworks_underlay": resourceUnderlay(),
            "nuagenetworks_redirection_target": resourceRedirectionTarget(),
            "nuagenetworks_enterprise": resourceEnterprise(),
            "nuagenetworks_wan_service": resourceWANService(),
            "nuagenetworks_site_info": resourceSiteInfo(),
            "nuagenetworks_avatar": resourceAvatar(),
            "nuagenetworks_key_server_monitor_encrypted_seed": resourceKeyServerMonitorEncryptedSeed(),
            "nuagenetworks_vsg_redundant_port": resourceVsgRedundantPort(),
            "nuagenetworks_policy_group_template": resourcePolicyGroupTemplate(),
            "nuagenetworks_application": resourceApplication(),
            "nuagenetworks_egress_acl_template": resourceEgressACLTemplate(),
            "nuagenetworks_metadata": resourceMetadata(),
            "nuagenetworks_ingress_adv_fwd_entry_template": resourceIngressAdvFwdEntryTemplate(),
            "nuagenetworks_redundant_port": resourceRedundantPort(),
            "nuagenetworks_link": resourceLink(),
            "nuagenetworks_demarcation_service": resourceDemarcationService(),
            "nuagenetworks_public_network_macro": resourcePublicNetworkMacro(),
            "nuagenetworks_ingress_external_service_template": resourceIngressExternalServiceTemplate(),
            "nuagenetworks_vport_mirror": resourceVPortMirror(),
            "nuagenetworks_floating_ipacl_template_entry": resourceFloatingIPACLTemplateEntry(),
            "nuagenetworks_flow_security_policy": resourceFlowSecurityPolicy(),
            "nuagenetworks_network_performance_measurement": resourceNetworkPerformanceMeasurement(),
            "nuagenetworks_external_service": resourceExternalService(),
            "nuagenetworks_ingress_acl_entry_template": resourceIngressACLEntryTemplate(),
            "nuagenetworks_policy_group": resourcePolicyGroup(),
            "nuagenetworks_vcenter_cluster": resourceVCenterCluster(),
            "nuagenetworks_duc_group_binding": resourceDUCGroupBinding(),
            "nuagenetworks_zone": resourceZone(),
            "nuagenetworks_container_resync": resourceContainerResync(),
            "nuagenetworks_multi_cast_channel_map": resourceMultiCastChannelMap(),
            "nuagenetworks_bgp_neighbor": resourceBGPNeighbor(),
            "nuagenetworks_vlan": resourceVLAN(),
            "nuagenetworks_statistics_policy": resourceStatisticsPolicy(),
            "nuagenetworks_zfb_request": resourceZFBRequest(),
            "nuagenetworks_job": resourceJob(),
            "nuagenetworks_network_macro_group": resourceNetworkMacroGroup(),
            "nuagenetworks_ns_gateway_template": resourceNSGatewayTemplate(),
            "nuagenetworks_patnat_pool": resourcePATNATPool(),
            "nuagenetworks_vport": resourceVPort(),
            "nuagenetworks_applicationperformancemanagementbinding": resourceApplicationperformancemanagementbinding(),
            "nuagenetworks_egress_qos_policy": resourceEgressQOSPolicy(),
            "nuagenetworks_group": resourceGroup(),
            "nuagenetworks_address_range": resourceAddressRange(),
            "nuagenetworks_bgp_profile": resourceBGPProfile(),
            "nuagenetworks_ingress_adv_fwd_template": resourceIngressAdvFwdTemplate(),
            "nuagenetworks_egress_acl_entry_template": resourceEgressACLEntryTemplate(),
            "nuagenetworks_ikepsk": resourceIKEPSK(),
            "nuagenetworks_port_template": resourcePortTemplate(),
            "nuagenetworks_ike_encryptionprofile": resourceIKEEncryptionprofile(),
            "nuagenetworks_floating_ip": resourceFloatingIp(),
            "nuagenetworks_flow_forwarding_policy": resourceFlowForwardingPolicy(),
            "nuagenetworks_ns_gateway": resourceNSGateway(),
            "nuagenetworks_enterprise_permission": resourceEnterprisePermission(),
            "nuagenetworks_multi_cast_range": resourceMultiCastRange(),
            "nuagenetworks_domain_template": resourceDomainTemplate(),
            "nuagenetworks_monitorscope": resourceMonitorscope(),
            "nuagenetworks_nsg_group": resourceNSGGroup(),
            "nuagenetworks_vcenter": resourceVCenter(),
            "nuagenetworks_overlay_patnat_entry": resourceOverlayPATNATEntry(),
            "nuagenetworks_infrastructure_gateway_profile": resourceInfrastructureGatewayProfile(),
            "nuagenetworks_ingress_acl_template": resourceIngressACLTemplate(),
            "nuagenetworks_vrs_redeploymentpolicy": resourceVRSRedeploymentpolicy(),
            "nuagenetworks_floating_ipacl_template": resourceFloatingIPACLTemplate(),
            "nuagenetworks_uplink_connection": resourceUplinkConnection(),
            "nuagenetworks_zfb_auto_assignment": resourceZFBAutoAssignment(),
            "nuagenetworks_port": resourcePort(),
            "nuagenetworks_redundancy_group": resourceRedundancyGroup(),
            "nuagenetworks_patip_entry": resourcePATIPEntry(),
            "nuagenetworks_metadata_tag": resourceMetadataTag(),
            "nuagenetworks_vm": resourceVM(),
            "nuagenetworks_container": resourceContainer(),
            "nuagenetworks_virtual_ip": resourceVirtualIP(),
            "nuagenetworks_applicationperformancemanagement": resourceApplicationperformancemanagement(),
            "nuagenetworks_infrastructure_access_profile": resourceInfrastructureAccessProfile(),
            "nuagenetworks_subnet_template": resourceSubnetTemplate(),
            "nuagenetworks_br_connection": resourceBRConnection(),
            "nuagenetworks_ip_reservation": resourceIPReservation(),
            "nuagenetworks_firewall_acl": resourceFirewallAcl(),
            "nuagenetworks_certificate": resourceCertificate(),
            "nuagenetworks_domain_fip_acl_template": resourceDomainFIPAclTemplate(),
            "nuagenetworks_address_map": resourceAddressMap(),
            "nuagenetworks_gateway": resourceGateway(),
            "nuagenetworks_ike_gateway_connection": resourceIKEGatewayConnection(),
            "nuagenetworks_gateway_secured_data": resourceGatewaySecuredData(),
            "nuagenetworks_ike_certificate": resourceIKECertificate(),
            "nuagenetworks_nat_map_entry": resourceNATMapEntry(),
        },
    }
}

func providerConfigure(d *schema.ResourceData) (root interface{}, err error) {
    // if we have a certificate path, we use cert auth
    log.Println("[INFO] Initializing Nuage Networks VSD client")

    var s *bambou.Session

    if certPathRaw, certPathOk := d.GetOk("certificate_path"); certPathOk {
      cert, tlsErr := tls.LoadX509KeyPair(certPathRaw.(string), d.Get("key_path").(string))
      if tlsErr != nil {
          return nil, errors.New("Error loading VSD generated certificates to authenticate with VSD: " + tlsErr.Error())
      }
      s, root = vspk.NewX509Session(&cert, d.Get("vsd_endpoint").(string))
    } else {
      s, root = vspk.NewSession(d.Get("username").(string), d.Get("password").(string), d.Get("enterprise").(string), d.Get("vsd_endpoint").(string))
    }

    berr := s.Start()

    if berr != nil {
        err = errors.New("Unable to connect to Nuage VSD: " + berr.Description)
        return
    }

    log.Println("[INFO] Nuage Networks VSD client initialized")

    return
}