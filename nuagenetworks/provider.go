package nuagenetworks

import (
	"crypto/tls"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/nuagenetworks/vspk-go/vspk"
	"log"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			// Only supporting cert based authentication for now
			// "username": &schema.Schema{
			//     Type:        schema.TypeString,
			//     Required:    true,
			//     DefaultFunc: schema.EnvDefaultFunc("VSD_USERNAME", "csproot"),
			// },
			// "password": &schema.Schema{
			//     Type:        schema.TypeString,
			//     Required:    true,
			//     Sensitive:   true,
			//     DefaultFunc: schema.EnvDefaultFunc("VSD_PASSWORD", "csproot"),
			// },
			// "enterprise": &schema.Schema{
			//     Type:        schema.TypeString,
			//     Required:    true,
			//     DefaultFunc: schema.EnvDefaultFunc("VSD_ORGANIZATION", "csp"),
			// },
			"certificate_path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSD_CERTIFICATE_PATH", nil),
			},
			"key_path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSD_KEY_PATH", nil),
			},

			"vsd_endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSD_ENDPOINT", nil),
			},
		},
		ConfigureFunc: providerConfigure,
		DataSourcesMap: map[string]*schema.Resource{
			"nuagenetworks_container_interface":                     dataSourceContainerInterface(),
			"nuagenetworks_group_key_encryption_profile":            dataSourceGroupKeyEncryptionProfile(),
			"nuagenetworks_qos":                                     dataSourceQOS(),
			"nuagenetworks_ingress_adv_fwd_entry_template":          dataSourceIngressAdvFwdEntryTemplate(),
			"nuagenetworks_bgp_peer":                                dataSourceBGPPeer(),
			"nuagenetworks_shared_network_resource":                 dataSourceSharedNetworkResource(),
			"nuagenetworks_overlay_mirror_destination":              dataSourceOverlayMirrorDestination(),
			"nuagenetworks_virtual_ip":                              dataSourceVirtualIP(),
			"nuagenetworks_pspat_map":                               dataSourcePSPATMap(),
			"nuagenetworks_egress_adv_fwd_entry_template":           dataSourceEgressAdvFwdEntryTemplate(),
			"nuagenetworks_dscp_forwarding_class_table":             dataSourceDSCPForwardingClassTable(),
			"nuagenetworks_multi_cast_channel_map":                  dataSourceMultiCastChannelMap(),
			"nuagenetworks_redundancy_group":                        dataSourceRedundancyGroup(),
			"nuagenetworks_tca":                                     dataSourceTCA(),
			"nuagenetworks_group":                                   dataSourceGroup(),
			"nuagenetworks_vsg_redundant_port":                      dataSourceVsgRedundantPort(),
			"nuagenetworks_zone":                                    dataSourceZone(),
			"nuagenetworks_ike_gateway_profile":                     dataSourceIKEGatewayProfile(),
			"nuagenetworks_ike_subnet":                              dataSourceIKESubnet(),
			"nuagenetworks_infrastructure_gateway_profile":          dataSourceInfrastructureGatewayProfile(),
			"nuagenetworks_vnf_interface":                           dataSourceVNFInterface(),
			"nuagenetworks_vcenter_eam_config":                      dataSourceVCenterEAMConfig(),
			"nuagenetworks_location":                                dataSourceLocation(),
			"nuagenetworks_enterprise_security":                     dataSourceEnterpriseSecurity(),
			"nuagenetworks_vpn_connection":                          dataSourceVPNConnection(),
			"nuagenetworks_duc_group":                               dataSourceDUCGroup(),
			"nuagenetworks_enterprise_network":                      dataSourceEnterpriseNetwork(),
			"nuagenetworks_infrastructure_vsc_profile":              dataSourceInfrastructureVscProfile(),
			"nuagenetworks_permission":                              dataSourcePermission(),
			"nuagenetworks_captive_portal_profile":                  dataSourceCaptivePortalProfile(),
			"nuagenetworks_cos_remarking_policy":                    dataSourceCOSRemarkingPolicy(),
			"nuagenetworks_l4_service_group":                        dataSourceL4ServiceGroup(),
			"nuagenetworks_ip_reservation":                          dataSourceIPReservation(),
			"nuagenetworks_redirection_target_template":             dataSourceRedirectionTargetTemplate(),
			"nuagenetworks_subnet_template":                         dataSourceSubnetTemplate(),
			"nuagenetworks_vnf_domain_mapping":                      dataSourceVNFDomainMapping(),
			"nuagenetworks_floating_ipacl_template":                 dataSourceFloatingIPACLTemplate(),
			"nuagenetworks_ike_gateway_connection":                  dataSourceIKEGatewayConnection(),
			"nuagenetworks_ingress_adv_fwd_template":                dataSourceIngressAdvFwdTemplate(),
			"nuagenetworks_redirection_target":                      dataSourceRedirectionTarget(),
			"nuagenetworks_firewall_acl":                            dataSourceFirewallAcl(),
			"nuagenetworks_patip_entry":                             dataSourcePATIPEntry(),
			"nuagenetworks_aggregate_metadata":                      dataSourceAggregateMetadata(),
			"nuagenetworks_egress_acl_entry_template":               dataSourceEgressACLEntryTemplate(),
			"nuagenetworks_ingress_external_service_template_entry": dataSourceIngressExternalServiceTemplateEntry(),
			"nuagenetworks_avatar":                                  dataSourceAvatar(),
			"nuagenetworks_overlay_address_pool":                    dataSourceOverlayAddressPool(),
			"nuagenetworks_license_status":                          dataSourceLicenseStatus(),
			"nuagenetworks_lte_information":                         dataSourceLTEInformation(),
			"nuagenetworks_cloud_mgmt_system":                       dataSourceCloudMgmtSystem(),
			"nuagenetworks_vsp":                                     dataSourceVSP(),
			"nuagenetworks_domain":                                  dataSourceDomain(),
			"nuagenetworks_redundant_port":                          dataSourceRedundantPort(),
			"nuagenetworks_proxy_arp_filter":                        dataSourceProxyARPFilter(),
			"nuagenetworks_nsg_patch_profile":                       dataSourceNSGPatchProfile(),
			"nuagenetworks_dscp_forwarding_class_mapping":           dataSourceDSCPForwardingClassMapping(),
			"nuagenetworks_policy_object_group":                     dataSourcePolicyObjectGroup(),
			"nuagenetworks_network_layout":                          dataSourceNetworkLayout(),
			"nuagenetworks_l4_service":                              dataSourceL4Service(),
			"nuagenetworks_overlay_patnat_entry":                    dataSourceOverlayPATNATEntry(),
			"nuagenetworks_wan_service":                             dataSourceWANService(),
			"nuagenetworks_vsd":                                     dataSourceVSD(),
			"nuagenetworks_vnf_interface_descriptor":                dataSourceVNFInterfaceDescriptor(),
			"nuagenetworks_zfb_auto_assignment":                     dataSourceZFBAutoAssignment(),
			"nuagenetworks_vport":                                   dataSourceVPort(),
			"nuagenetworks_tier":                                    dataSourceTier(),
			"nuagenetworks_port":                                    dataSourcePort(),
			"nuagenetworks_policy_statement":                        dataSourcePolicyStatement(),
			"nuagenetworks_statistics_policy":                       dataSourceStatisticsPolicy(),
			"nuagenetworks_subnet":                                  dataSourceSubnet(),
			"nuagenetworks_container":                               dataSourceContainer(),
			"nuagenetworks_rate_limiter":                            dataSourceRateLimiter(),
			"nuagenetworks_key_server_monitor_encrypted_seed":       dataSourceKeyServerMonitorEncryptedSeed(),
			"nuagenetworks_policy_entry":                            dataSourcePolicyEntry(),
			"nuagenetworks_metadata":                                dataSourceMetadata(),
			"nuagenetworks_zone_template":                           dataSourceZoneTemplate(),
			"nuagenetworks_multi_nic_vport":                         dataSourceMultiNICVPort(),
			"nuagenetworks_key_server_monitor_seed":                 dataSourceKeyServerMonitorSeed(),
			"nuagenetworks_gateway_template":                        dataSourceGatewayTemplate(),
			"nuagenetworks_vrs":                                     dataSourceVRS(),
			"nuagenetworks_hsc":                                     dataSourceHSC(),
			"nuagenetworks_l2_domain_template":                      dataSourceL2DomainTemplate(),
			"nuagenetworks_vlan":                                    dataSourceVLAN(),
			"nuagenetworks_overlay_mirror_destination_template":     dataSourceOverlayMirrorDestinationTemplate(),
			"nuagenetworks_ldap_configuration":                      dataSourceLDAPConfiguration(),
			"nuagenetworks_vsd_component":                           dataSourceVSDComponent(),
			"nuagenetworks_zfb_request":                             dataSourceZFBRequest(),
			"nuagenetworks_enterprise_permission":                   dataSourceEnterprisePermission(),
			"nuagenetworks_static_route":                            dataSourceStaticRoute(),
			"nuagenetworks_connectionendpoint":                      dataSourceConnectionendpoint(),
			"nuagenetworks_job":                                     dataSourceJob(),
			"nuagenetworks_vm_interface":                            dataSourceVMInterface(),
			"nuagenetworks_gateway_secured_data":                    dataSourceGatewaySecuredData(),
			"nuagenetworks_trunk":                                   dataSourceTrunk(),
			"nuagenetworks_vnf_descriptor":                          dataSourceVNFDescriptor(),
			"nuagenetworks_key_server_monitor":                      dataSourceKeyServerMonitor(),
			"nuagenetworks_vnf_catalog":                             dataSourceVNFCatalog(),
			"nuagenetworks_duc_group_binding":                       dataSourceDUCGroupBinding(),
			"nuagenetworks_vrs_metrics":                             dataSourceVRSMetrics(),
			"nuagenetworks_event_log":                               dataSourceEventLog(),
			"nuagenetworks_license":                                 dataSourceLicense(),
			"nuagenetworks_enterprise_profile":                      dataSourceEnterpriseProfile(),
			"nuagenetworks_performance_monitor":                     dataSourcePerformanceMonitor(),
			"nuagenetworks_policy_group_template":                   dataSourcePolicyGroupTemplate(),
			"nuagenetworks_bridge_interface":                        dataSourceBridgeInterface(),
			"nuagenetworks_vcenter_cluster":                         dataSourceVCenterCluster(),
			"nuagenetworks_pg_expression":                           dataSourcePGExpression(),
			"nuagenetworks_dscp_remarking_policy":                   dataSourceDSCPRemarkingPolicy(),
			"nuagenetworks_ssid_connection":                         dataSourceSSIDConnection(),
			"nuagenetworks_network_macro_group":                     dataSourceNetworkMacroGroup(),
			"nuagenetworks_infrastructure_access_profile":           dataSourceInfrastructureAccessProfile(),
			"nuagenetworks_application_binding":                     dataSourceApplicationBinding(),
			"nuagenetworks_destinationurl":                          dataSourceDestinationurl(),
			"nuagenetworks_auto_discovered_gateway":                 dataSourceAutoDiscoveredGateway(),
			"nuagenetworks_multi_cast_list":                         dataSourceMultiCastList(),
			"nuagenetworks_next_hop":                                dataSourceNextHop(),
			"nuagenetworks_mirror_destination":                      dataSourceMirrorDestination(),
			"nuagenetworks_nat_map_entry":                           dataSourceNATMapEntry(),
			"nuagenetworks_domain_fip_acl_template":                 dataSourceDomainFIPAclTemplate(),
			"nuagenetworks_ospf_interface":                          dataSourceOSPFInterface(),
			"nuagenetworks_address_map":                             dataSourceAddressMap(),
			"nuagenetworks_underlay":                                dataSourceUnderlay(),
			"nuagenetworks_gateway":                                 dataSourceGateway(),
			"nuagenetworks_egress_qos_policy":                       dataSourceEgressQOSPolicy(),
			"nuagenetworks_vm":                                      dataSourceVM(),
			"nuagenetworks_bfd_session":                             dataSourceBFDSession(),
			"nuagenetworks_statistics":                              dataSourceStatistics(),
			"nuagenetworks_ns_port_template":                        dataSourceNSPortTemplate(),
			"nuagenetworks_ssh_key":                                 dataSourceSSHKey(),
			"nuagenetworks_vcenter_data_center":                     dataSourceVCenterDataCenter(),
			"nuagenetworks_custom_property":                         dataSourceCustomProperty(),
			"nuagenetworks_ltestatistics":                           dataSourceLtestatistics(),
			"nuagenetworks_virtual_firewall_rule":                   dataSourceVirtualFirewallRule(),
			"nuagenetworks_dscp_remarking_policy_table":             dataSourceDSCPRemarkingPolicyTable(),
			"nuagenetworks_vrs_redeploymentpolicy":                  dataSourceVRSRedeploymentpolicy(),
			"nuagenetworks_p_translation_map":                       dataSourcePTranslationMap(),
			"nuagenetworks_ike_gateway":                             dataSourceIKEGateway(),
			"nuagenetworks_csnat_pool":                              dataSourceCSNATPool(),
			"nuagenetworks_stats_collector_info":                    dataSourceStatsCollectorInfo(),
			"nuagenetworks_vcenter":                                 dataSourceVCenter(),
			"nuagenetworks_bulk_statistics":                         dataSourceBulkStatistics(),
			"nuagenetworks_ingress_acl_entry_template":              dataSourceIngressACLEntryTemplate(),
			"nuagenetworks_ingress_qos_policy":                      dataSourceIngressQOSPolicy(),
			"nuagenetworks_routing_policy":                          dataSourceRoutingPolicy(),
			"nuagenetworks_network_performance_binding":             dataSourceNetworkPerformanceBinding(),
			"nuagenetworks_vnf_threshold_policy":                    dataSourceVNFThresholdPolicy(),
			"nuagenetworks_default_gateway":                         dataSourceDefaultGateway(),
			"nuagenetworks_l2_domain":                               dataSourceL2Domain(),
			"nuagenetworks_ike_gateway_config":                      dataSourceIKEGatewayConfig(),
			"nuagenetworks_host_interface":                          dataSourceHostInterface(),
			"nuagenetworks_enterprise_secured_data":                 dataSourceEnterpriseSecuredData(),
			"nuagenetworks_applicationperformancemanagement":        dataSourceApplicationperformancemanagement(),
			"nuagenetworks_qos_policer":                             dataSourceQosPolicer(),
			"nuagenetworks_ike_certificate":                         dataSourceIKECertificate(),
			"nuagenetworks_ingress_external_service_template":       dataSourceIngressExternalServiceTemplate(),
			"nuagenetworks_port_template":                           dataSourcePortTemplate(),
			"nuagenetworks_egress_adv_fwd_template":                 dataSourceEgressAdvFwdTemplate(),
			"nuagenetworks_user_context":                            dataSourceUserContext(),
			"nuagenetworks_l7applicationsignature":                  dataSourceL7applicationsignature(),
			"nuagenetworks_dhcp_option":                             dataSourceDHCPOption(),
			"nuagenetworks_key_server_member":                       dataSourceKeyServerMember(),
			"nuagenetworks_ns_gateway":                              dataSourceNSGateway(),
			"nuagenetworks_ns_gateway_template":                     dataSourceNSGatewayTemplate(),
			"nuagenetworks_vsc":                                     dataSourceVSC(),
			"nuagenetworks_uplink_rd":                               dataSourceUplinkRD(),
			"nuagenetworks_nsg_group":                               dataSourceNSGGroup(),
			"nuagenetworks_site_info":                               dataSourceSiteInfo(),
			"nuagenetworks_ns_port":                                 dataSourceNSPort(),
			"nuagenetworks_nsg_routing_policy_binding":              dataSourceNSGRoutingPolicyBinding(),
			"nuagenetworks_vcenter_vrs_config":                      dataSourceVCenterVRSConfig(),
			"nuagenetworks_vrs_address_range":                       dataSourceVRSAddressRange(),
			"nuagenetworks_bgp_profile":                             dataSourceBGPProfile(),
			"nuagenetworks_c_translation_map":                       dataSourceCTranslationMap(),
			"nuagenetworks_public_network_macro":                    dataSourcePublicNetworkMacro(),
			"nuagenetworks_domain_fip_acl_template_entry":           dataSourceDomainFIPAclTemplateEntry(),
			"nuagenetworks_auto_discover_cluster":                   dataSourceAutoDiscoverCluster(),
			"nuagenetworks_address_range":                           dataSourceAddressRange(),
			"nuagenetworks_domain_template":                         dataSourceDomainTemplate(),
			"nuagenetworks_ospf_instance":                           dataSourceOSPFInstance(),
			"nuagenetworks_virtual_firewall_policy":                 dataSourceVirtualFirewallPolicy(),
			"nuagenetworks_vm_resync":                               dataSourceVMResync(),
			"nuagenetworks_port_mapping":                            dataSourcePortMapping(),
			"nuagenetworks_uplink_connection":                       dataSourceUplinkConnection(),
			"nuagenetworks_policy_decision":                         dataSourcePolicyDecision(),
			"nuagenetworks_pg_expression_template":                  dataSourcePGExpressionTemplate(),
			"nuagenetworks_applicationperformancemanagementbinding": dataSourceApplicationperformancemanagementbinding(),
			"nuagenetworks_autodiscovereddatacenter":                dataSourceAutodiscovereddatacenter(),
			"nuagenetworks_spat_sources_pool":                       dataSourceSPATSourcesPool(),
			"nuagenetworks_floating_ip":                             dataSourceFloatingIp(),
			"nuagenetworks_egress_acl_template":                     dataSourceEgressACLTemplate(),
			"nuagenetworks_monitoring_port":                         dataSourceMonitoringPort(),
			"nuagenetworks_monitorscope":                            dataSourceMonitorscope(),
			"nuagenetworks_auto_discover_hypervisor_from_cluster":   dataSourceAutoDiscoverHypervisorFromCluster(),
			"nuagenetworks_multi_cast_range":                        dataSourceMultiCastRange(),
			"nuagenetworks_ns_redundant_gateway_group":              dataSourceNSRedundantGatewayGroup(),
			"nuagenetworks_application":                             dataSourceApplication(),
			"nuagenetworks_key_server_monitor_sek":                  dataSourceKeyServerMonitorSEK(),
			"nuagenetworks_cos_remarking_policy_table":              dataSourceCOSRemarkingPolicyTable(),
			"nuagenetworks_vport_mirror":                            dataSourceVPortMirror(),
			"nuagenetworks_patnat_pool":                             dataSourcePATNATPool(),
			"nuagenetworks_pat_mapper":                              dataSourcePATMapper(),
			"nuagenetworks_psnat_pool":                              dataSourcePSNATPool(),
			"nuagenetworks_bgp_neighbor":                            dataSourceBGPNeighbor(),
			"nuagenetworks_container_resync":                        dataSourceContainerResync(),
			"nuagenetworks_vnf":                                     dataSourceVNF(),
			"nuagenetworks_all_alarm":                               dataSourceAllAlarm(),
			"nuagenetworks_ikepsk":                                  dataSourceIKEPSK(),
			"nuagenetworks_wireless_port":                           dataSourceWirelessPort(),
			"nuagenetworks_system_config":                           dataSourceSystemConfig(),
			"nuagenetworks_ike_encryptionprofile":                   dataSourceIKEEncryptionprofile(),
			"nuagenetworks_floating_ipacl_template_entry":           dataSourceFloatingIPACLTemplateEntry(),
			"nuagenetworks_infrastructure_config":                   dataSourceInfrastructureConfig(),
			"nuagenetworks_user":                                    dataSourceUser(),
			"nuagenetworks_br_connection":                           dataSourceBRConnection(),
			"nuagenetworks_policy_group":                            dataSourcePolicyGroup(),
			"nuagenetworks_demarcation_service":                     dataSourceDemarcationService(),
			"nuagenetworks_nsg_info":                                dataSourceNSGInfo(),
			"nuagenetworks_firewall_rule":                           dataSourceFirewallRule(),
			"nuagenetworks_alarm":                                   dataSourceAlarm(),
			"nuagenetworks_bootstrap":                               dataSourceBootstrap(),
			"nuagenetworks_vlan_template":                           dataSourceVLANTemplate(),
			"nuagenetworks_nsg_upgrade_profile":                     dataSourceNSGUpgradeProfile(),
			"nuagenetworks_gateway_security":                        dataSourceGatewaySecurity(),
			"nuagenetworks_global_metadata":                         dataSourceGlobalMetadata(),
			"nuagenetworks_network_performance_measurement":         dataSourceNetworkPerformanceMeasurement(),
			"nuagenetworks_ospf_area":                               dataSourceOSPFArea(),
			"nuagenetworks_command":                                 dataSourceCommand(),
			"nuagenetworks_enterprise":                              dataSourceEnterprise(),
			"nuagenetworks_link":                                    dataSourceLink(),
			"nuagenetworks_ingress_acl_template":                    dataSourceIngressACLTemplate(),
			"nuagenetworks_vnf_metadata":                            dataSourceVNFMetadata(),
			"nuagenetworks_vcenter_hypervisor":                      dataSourceVCenterHypervisor(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"nuagenetworks_container_interface":                     resourceContainerInterface(),
			"nuagenetworks_qos":                                     resourceQOS(),
			"nuagenetworks_ingress_adv_fwd_entry_template":          resourceIngressAdvFwdEntryTemplate(),
			"nuagenetworks_shared_network_resource":                 resourceSharedNetworkResource(),
			"nuagenetworks_overlay_mirror_destination":              resourceOverlayMirrorDestination(),
			"nuagenetworks_virtual_ip":                              resourceVirtualIP(),
			"nuagenetworks_pspat_map":                               resourcePSPATMap(),
			"nuagenetworks_egress_adv_fwd_entry_template":           resourceEgressAdvFwdEntryTemplate(),
			"nuagenetworks_dscp_forwarding_class_table":             resourceDSCPForwardingClassTable(),
			"nuagenetworks_multi_cast_channel_map":                  resourceMultiCastChannelMap(),
			"nuagenetworks_redundancy_group":                        resourceRedundancyGroup(),
			"nuagenetworks_tca":                                     resourceTCA(),
			"nuagenetworks_group":                                   resourceGroup(),
			"nuagenetworks_vsg_redundant_port":                      resourceVsgRedundantPort(),
			"nuagenetworks_zone":                                    resourceZone(),
			"nuagenetworks_ike_gateway_profile":                     resourceIKEGatewayProfile(),
			"nuagenetworks_ike_subnet":                              resourceIKESubnet(),
			"nuagenetworks_infrastructure_gateway_profile":          resourceInfrastructureGatewayProfile(),
			"nuagenetworks_vpn_connection":                          resourceVPNConnection(),
			"nuagenetworks_duc_group":                               resourceDUCGroup(),
			"nuagenetworks_enterprise_network":                      resourceEnterpriseNetwork(),
			"nuagenetworks_infrastructure_vsc_profile":              resourceInfrastructureVscProfile(),
			"nuagenetworks_permission":                              resourcePermission(),
			"nuagenetworks_captive_portal_profile":                  resourceCaptivePortalProfile(),
			"nuagenetworks_cos_remarking_policy":                    resourceCOSRemarkingPolicy(),
			"nuagenetworks_l4_service_group":                        resourceL4ServiceGroup(),
			"nuagenetworks_ip_reservation":                          resourceIPReservation(),
			"nuagenetworks_redirection_target_template":             resourceRedirectionTargetTemplate(),
			"nuagenetworks_subnet_template":                         resourceSubnetTemplate(),
			"nuagenetworks_vnf_domain_mapping":                      resourceVNFDomainMapping(),
			"nuagenetworks_floating_ipacl_template":                 resourceFloatingIPACLTemplate(),
			"nuagenetworks_ike_gateway_connection":                  resourceIKEGatewayConnection(),
			"nuagenetworks_ingress_adv_fwd_template":                resourceIngressAdvFwdTemplate(),
			"nuagenetworks_redirection_target":                      resourceRedirectionTarget(),
			"nuagenetworks_firewall_acl":                            resourceFirewallAcl(),
			"nuagenetworks_patip_entry":                             resourcePATIPEntry(),
			"nuagenetworks_egress_acl_entry_template":               resourceEgressACLEntryTemplate(),
			"nuagenetworks_ingress_external_service_template_entry": resourceIngressExternalServiceTemplateEntry(),
			"nuagenetworks_avatar":                                  resourceAvatar(),
			"nuagenetworks_overlay_address_pool":                    resourceOverlayAddressPool(),
			"nuagenetworks_bootstrap_activation":                    resourceBootstrapActivation(),
			"nuagenetworks_cloud_mgmt_system":                       resourceCloudMgmtSystem(),
			"nuagenetworks_domain":                                  resourceDomain(),
			"nuagenetworks_redundant_port":                          resourceRedundantPort(),
			"nuagenetworks_proxy_arp_filter":                        resourceProxyARPFilter(),
			"nuagenetworks_nsg_patch_profile":                       resourceNSGPatchProfile(),
			"nuagenetworks_dscp_forwarding_class_mapping":           resourceDSCPForwardingClassMapping(),
			"nuagenetworks_policy_object_group":                     resourcePolicyObjectGroup(),
			"nuagenetworks_l4_service":                              resourceL4Service(),
			"nuagenetworks_overlay_patnat_entry":                    resourceOverlayPATNATEntry(),
			"nuagenetworks_wan_service":                             resourceWANService(),
			"nuagenetworks_vnf_interface_descriptor":                resourceVNFInterfaceDescriptor(),
			"nuagenetworks_zfb_auto_assignment":                     resourceZFBAutoAssignment(),
			"nuagenetworks_vport":                                   resourceVPort(),
			"nuagenetworks_port":                                    resourcePort(),
			"nuagenetworks_policy_statement":                        resourcePolicyStatement(),
			"nuagenetworks_statistics_policy":                       resourceStatisticsPolicy(),
			"nuagenetworks_subnet":                                  resourceSubnet(),
			"nuagenetworks_container":                               resourceContainer(),
			"nuagenetworks_rate_limiter":                            resourceRateLimiter(),
			"nuagenetworks_key_server_monitor_encrypted_seed":       resourceKeyServerMonitorEncryptedSeed(),
			"nuagenetworks_policy_entry":                            resourcePolicyEntry(),
			"nuagenetworks_metadata":                                resourceMetadata(),
			"nuagenetworks_zone_template":                           resourceZoneTemplate(),
			"nuagenetworks_key_server_monitor_seed":                 resourceKeyServerMonitorSeed(),
			"nuagenetworks_gateway_template":                        resourceGatewayTemplate(),
			"nuagenetworks_l2_domain_template":                      resourceL2DomainTemplate(),
			"nuagenetworks_vlan":                                    resourceVLAN(),
			"nuagenetworks_overlay_mirror_destination_template":     resourceOverlayMirrorDestinationTemplate(),
			"nuagenetworks_zfb_request":                             resourceZFBRequest(),
			"nuagenetworks_enterprise_permission":                   resourceEnterprisePermission(),
			"nuagenetworks_static_route":                            resourceStaticRoute(),
			"nuagenetworks_connectionendpoint":                      resourceConnectionendpoint(),
			"nuagenetworks_job":                                     resourceJob(),
			"nuagenetworks_vm_interface":                            resourceVMInterface(),
			"nuagenetworks_gateway_secured_data":                    resourceGatewaySecuredData(),
			"nuagenetworks_trunk":                                   resourceTrunk(),
			"nuagenetworks_vnf_descriptor":                          resourceVNFDescriptor(),
			"nuagenetworks_duc_group_binding":                       resourceDUCGroupBinding(),
			"nuagenetworks_license":                                 resourceLicense(),
			"nuagenetworks_enterprise_profile":                      resourceEnterpriseProfile(),
			"nuagenetworks_performance_monitor":                     resourcePerformanceMonitor(),
			"nuagenetworks_policy_group_template":                   resourcePolicyGroupTemplate(),
			"nuagenetworks_bridge_interface":                        resourceBridgeInterface(),
			"nuagenetworks_vcenter_cluster":                         resourceVCenterCluster(),
			"nuagenetworks_pg_expression":                           resourcePGExpression(),
			"nuagenetworks_dscp_remarking_policy":                   resourceDSCPRemarkingPolicy(),
			"nuagenetworks_ssid_connection":                         resourceSSIDConnection(),
			"nuagenetworks_network_macro_group":                     resourceNetworkMacroGroup(),
			"nuagenetworks_infrastructure_access_profile":           resourceInfrastructureAccessProfile(),
			"nuagenetworks_application_binding":                     resourceApplicationBinding(),
			"nuagenetworks_destinationurl":                          resourceDestinationurl(),
			"nuagenetworks_next_hop":                                resourceNextHop(),
			"nuagenetworks_mirror_destination":                      resourceMirrorDestination(),
			"nuagenetworks_nat_map_entry":                           resourceNATMapEntry(),
			"nuagenetworks_domain_fip_acl_template":                 resourceDomainFIPAclTemplate(),
			"nuagenetworks_ospf_interface":                          resourceOSPFInterface(),
			"nuagenetworks_address_map":                             resourceAddressMap(),
			"nuagenetworks_underlay":                                resourceUnderlay(),
			"nuagenetworks_gateway":                                 resourceGateway(),
			"nuagenetworks_egress_qos_policy":                       resourceEgressQOSPolicy(),
			"nuagenetworks_vm":                                      resourceVM(),
			"nuagenetworks_bfd_session":                             resourceBFDSession(),
			"nuagenetworks_ns_port_template":                        resourceNSPortTemplate(),
			"nuagenetworks_ssh_key":                                 resourceSSHKey(),
			"nuagenetworks_certificate":                             resourceCertificate(),
			"nuagenetworks_vcenter_data_center":                     resourceVCenterDataCenter(),
			"nuagenetworks_custom_property":                         resourceCustomProperty(),
			"nuagenetworks_ltestatistics":                           resourceLtestatistics(),
			"nuagenetworks_virtual_firewall_rule":                   resourceVirtualFirewallRule(),
			"nuagenetworks_dscp_remarking_policy_table":             resourceDSCPRemarkingPolicyTable(),
			"nuagenetworks_vrs_redeploymentpolicy":                  resourceVRSRedeploymentpolicy(),
			"nuagenetworks_p_translation_map":                       resourcePTranslationMap(),
			"nuagenetworks_ike_gateway":                             resourceIKEGateway(),
			"nuagenetworks_csnat_pool":                              resourceCSNATPool(),
			"nuagenetworks_vcenter":                                 resourceVCenter(),
			"nuagenetworks_ingress_acl_entry_template":              resourceIngressACLEntryTemplate(),
			"nuagenetworks_ingress_qos_policy":                      resourceIngressQOSPolicy(),
			"nuagenetworks_routing_policy":                          resourceRoutingPolicy(),
			"nuagenetworks_network_performance_binding":             resourceNetworkPerformanceBinding(),
			"nuagenetworks_vnf_threshold_policy":                    resourceVNFThresholdPolicy(),
			"nuagenetworks_l2_domain":                               resourceL2Domain(),
			"nuagenetworks_host_interface":                          resourceHostInterface(),
			"nuagenetworks_enterprise_secured_data":                 resourceEnterpriseSecuredData(),
			"nuagenetworks_applicationperformancemanagement":        resourceApplicationperformancemanagement(),
			"nuagenetworks_qos_policer":                             resourceQosPolicer(),
			"nuagenetworks_ike_certificate":                         resourceIKECertificate(),
			"nuagenetworks_ingress_external_service_template":       resourceIngressExternalServiceTemplate(),
			"nuagenetworks_port_template":                           resourcePortTemplate(),
			"nuagenetworks_egress_adv_fwd_template":                 resourceEgressAdvFwdTemplate(),
			"nuagenetworks_dhcp_option":                             resourceDHCPOption(),
			"nuagenetworks_key_server_member":                       resourceKeyServerMember(),
			"nuagenetworks_ns_gateway":                              resourceNSGateway(),
			"nuagenetworks_ns_gateway_template":                     resourceNSGatewayTemplate(),
			"nuagenetworks_nsg_group":                               resourceNSGGroup(),
			"nuagenetworks_site_info":                               resourceSiteInfo(),
			"nuagenetworks_ns_port":                                 resourceNSPort(),
			"nuagenetworks_nsg_routing_policy_binding":              resourceNSGRoutingPolicyBinding(),
			"nuagenetworks_vrs_address_range":                       resourceVRSAddressRange(),
			"nuagenetworks_bgp_profile":                             resourceBGPProfile(),
			"nuagenetworks_c_translation_map":                       resourceCTranslationMap(),
			"nuagenetworks_public_network_macro":                    resourcePublicNetworkMacro(),
			"nuagenetworks_domain_fip_acl_template_entry":           resourceDomainFIPAclTemplateEntry(),
			"nuagenetworks_address_range":                           resourceAddressRange(),
			"nuagenetworks_domain_template":                         resourceDomainTemplate(),
			"nuagenetworks_ospf_instance":                           resourceOSPFInstance(),
			"nuagenetworks_virtual_firewall_policy":                 resourceVirtualFirewallPolicy(),
			"nuagenetworks_vm_resync":                               resourceVMResync(),
			"nuagenetworks_uplink_connection":                       resourceUplinkConnection(),
			"nuagenetworks_pg_expression_template":                  resourcePGExpressionTemplate(),
			"nuagenetworks_applicationperformancemanagementbinding": resourceApplicationperformancemanagementbinding(),
			"nuagenetworks_spat_sources_pool":                       resourceSPATSourcesPool(),
			"nuagenetworks_floating_ip":                             resourceFloatingIp(),
			"nuagenetworks_egress_acl_template":                     resourceEgressACLTemplate(),
			"nuagenetworks_monitorscope":                            resourceMonitorscope(),
			"nuagenetworks_multi_cast_range":                        resourceMultiCastRange(),
			"nuagenetworks_ns_redundant_gateway_group":              resourceNSRedundantGatewayGroup(),
			"nuagenetworks_application":                             resourceApplication(),
			"nuagenetworks_key_server_monitor_sek":                  resourceKeyServerMonitorSEK(),
			"nuagenetworks_cos_remarking_policy_table":              resourceCOSRemarkingPolicyTable(),
			"nuagenetworks_vport_mirror":                            resourceVPortMirror(),
			"nuagenetworks_patnat_pool":                             resourcePATNATPool(),
			"nuagenetworks_pat_mapper":                              resourcePATMapper(),
			"nuagenetworks_psnat_pool":                              resourcePSNATPool(),
			"nuagenetworks_bgp_neighbor":                            resourceBGPNeighbor(),
			"nuagenetworks_container_resync":                        resourceContainerResync(),
			"nuagenetworks_vnf":                                     resourceVNF(),
			"nuagenetworks_ikepsk":                                  resourceIKEPSK(),
			"nuagenetworks_wireless_port":                           resourceWirelessPort(),
			"nuagenetworks_ike_encryptionprofile":                   resourceIKEEncryptionprofile(),
			"nuagenetworks_floating_ipacl_template_entry":           resourceFloatingIPACLTemplateEntry(),
			"nuagenetworks_user":                                    resourceUser(),
			"nuagenetworks_br_connection":                           resourceBRConnection(),
			"nuagenetworks_policy_group":                            resourcePolicyGroup(),
			"nuagenetworks_demarcation_service":                     resourceDemarcationService(),
			"nuagenetworks_firewall_rule":                           resourceFirewallRule(),
			"nuagenetworks_alarm":                                   resourceAlarm(),
			"nuagenetworks_vlan_template":                           resourceVLANTemplate(),
			"nuagenetworks_nsg_upgrade_profile":                     resourceNSGUpgradeProfile(),
			"nuagenetworks_global_metadata":                         resourceGlobalMetadata(),
			"nuagenetworks_network_performance_measurement":         resourceNetworkPerformanceMeasurement(),
			"nuagenetworks_ospf_area":                               resourceOSPFArea(),
			"nuagenetworks_command":                                 resourceCommand(),
			"nuagenetworks_enterprise":                              resourceEnterprise(),
			"nuagenetworks_link":                                    resourceLink(),
			"nuagenetworks_ingress_acl_template":                    resourceIngressACLTemplate(),
			"nuagenetworks_vnf_metadata":                            resourceVNFMetadata(),
			"nuagenetworks_vcenter_hypervisor":                      resourceVCenterHypervisor(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	// s, root := vspk.NewSession(d.Get("username").(string), d.Get("password").(string), d.Get("enterprise").(string), d.Get("vsd_endpoint").(string))

	cert, tlsErr := tls.LoadX509KeyPair(d.Get("certificate_path").(string), d.Get("key_path").(string))
	if tlsErr != nil {
		return nil, errors.New("Error loading VSD generated certificates to authenticate with VSD: " + tlsErr.Error())
	}
	s, root := vspk.NewX509Session(&cert, d.Get("vsd_endpoint").(string))

	log.Println("[INFO] Initializing Nuage Networks VSD client")
	err := s.Start()
	if err != nil {
		return nil, errors.New("Unable to connect to Nuage VSD: " + err.Description)
	}
	log.Println("[INFO] Nuage Networks VSD client initialized")

	return root, nil
}