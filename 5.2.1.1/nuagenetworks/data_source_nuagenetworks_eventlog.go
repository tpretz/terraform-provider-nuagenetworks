package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.2.1.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceEventLog() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceEventLogRead,
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
            "request_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "diff": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entities": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "user": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "event_received_time": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_tca": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_enterprise_profile": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_vm_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_auto_discovered_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_host_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_qos": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_permission": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_container_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_vsp": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_static_route": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_bridge_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_redirection_target_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_l2_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_ns_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_user": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_license": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_dhcp_option": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_enterprise_network": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_ns_redundant_gateway_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_zone_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_redirection_target": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_wan_service": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_policy_group_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_egress_acl_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_metadata": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_wireless_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_proxy_arp_filter": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_public_network_macro": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_policy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_zone": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_multi_cast_channel_map": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_ssid_connection": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_vlan": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_vsc": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_hsc": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_vrs": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_address_range": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_floating_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_multi_cast_range": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_ingress_acl_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_redundancy_group": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_vm": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_container": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_virtual_ip": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_subnet_template", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_subnet_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_ip_reservation", "parent_gateway", "parent_vsd"},
            },
            "parent_ip_reservation": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_gateway", "parent_vsd"},
            },
            "parent_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_vsd"},
            },
            "parent_vsd": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_tca", "parent_enterprise_profile", "parent_vm_interface", "parent_auto_discovered_gateway", "parent_host_interface", "parent_qos", "parent_permission", "parent_container_interface", "parent_vsp", "parent_l2_domain", "parent_static_route", "parent_bridge_interface", "parent_redirection_target_template", "parent_l2_domain_template", "parent_subnet", "parent_ns_port", "parent_user", "parent_license", "parent_dhcp_option", "parent_enterprise_network", "parent_ns_redundant_gateway_group", "parent_zone_template", "parent_redirection_target", "parent_enterprise", "parent_wan_service", "parent_policy_group_template", "parent_egress_acl_template", "parent_metadata", "parent_wireless_port", "parent_proxy_arp_filter", "parent_public_network_macro", "parent_policy_group", "parent_zone", "parent_multi_cast_channel_map", "parent_ssid_connection", "parent_vlan", "parent_vport", "parent_vsc", "parent_hsc", "parent_vrs", "parent_group", "parent_address_range", "parent_floating_ip", "parent_ns_gateway", "parent_multi_cast_range", "parent_domain_template", "parent_ingress_acl_template", "parent_port", "parent_redundancy_group", "parent_vm", "parent_container", "parent_virtual_ip", "parent_subnet_template", "parent_ip_reservation", "parent_gateway"},
            },
        },
    }
}


func dataSourceEventLogRead(d *schema.ResourceData, m interface{}) error {
    filteredEventLogs := vspk.EventLogsList{}
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
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_tca"); ok {
        parent := &vspk.TCA{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise_profile"); ok {
        parent := &vspk.EnterpriseProfile{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vm_interface"); ok {
        parent := &vspk.VMInterface{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_auto_discovered_gateway"); ok {
        parent := &vspk.AutoDiscoveredGateway{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_host_interface"); ok {
        parent := &vspk.HostInterface{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_qos"); ok {
        parent := &vspk.QOS{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_permission"); ok {
        parent := &vspk.Permission{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_container_interface"); ok {
        parent := &vspk.ContainerInterface{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vsp"); ok {
        parent := &vspk.VSP{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_static_route"); ok {
        parent := &vspk.StaticRoute{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_bridge_interface"); ok {
        parent := &vspk.BridgeInterface{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_redirection_target_template"); ok {
        parent := &vspk.RedirectionTargetTemplate{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
        parent := &vspk.L2DomainTemplate{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ns_port"); ok {
        parent := &vspk.NSPort{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_user"); ok {
        parent := &vspk.User{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_license"); ok {
        parent := &vspk.License{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_dhcp_option"); ok {
        parent := &vspk.DHCPOption{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise_network"); ok {
        parent := &vspk.EnterpriseNetwork{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ns_redundant_gateway_group"); ok {
        parent := &vspk.NSRedundantGatewayGroup{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_zone_template"); ok {
        parent := &vspk.ZoneTemplate{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_redirection_target"); ok {
        parent := &vspk.RedirectionTarget{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_wan_service"); ok {
        parent := &vspk.WANService{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_policy_group_template"); ok {
        parent := &vspk.PolicyGroupTemplate{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_egress_acl_template"); ok {
        parent := &vspk.EgressACLTemplate{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_metadata"); ok {
        parent := &vspk.Metadata{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_wireless_port"); ok {
        parent := &vspk.WirelessPort{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_proxy_arp_filter"); ok {
        parent := &vspk.ProxyARPFilter{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_public_network_macro"); ok {
        parent := &vspk.PublicNetworkMacro{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_policy_group"); ok {
        parent := &vspk.PolicyGroup{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_zone"); ok {
        parent := &vspk.Zone{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_multi_cast_channel_map"); ok {
        parent := &vspk.MultiCastChannelMap{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ssid_connection"); ok {
        parent := &vspk.SSIDConnection{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vlan"); ok {
        parent := &vspk.VLAN{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vsc"); ok {
        parent := &vspk.VSC{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_hsc"); ok {
        parent := &vspk.HSC{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vrs"); ok {
        parent := &vspk.VRS{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_group"); ok {
        parent := &vspk.Group{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_address_range"); ok {
        parent := &vspk.AddressRange{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_floating_ip"); ok {
        parent := &vspk.FloatingIp{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ns_gateway"); ok {
        parent := &vspk.NSGateway{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_multi_cast_range"); ok {
        parent := &vspk.MultiCastRange{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_domain_template"); ok {
        parent := &vspk.DomainTemplate{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ingress_acl_template"); ok {
        parent := &vspk.IngressACLTemplate{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_port"); ok {
        parent := &vspk.Port{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_redundancy_group"); ok {
        parent := &vspk.RedundancyGroup{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vm"); ok {
        parent := &vspk.VM{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_container"); ok {
        parent := &vspk.Container{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_virtual_ip"); ok {
        parent := &vspk.VirtualIP{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_subnet_template"); ok {
        parent := &vspk.SubnetTemplate{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ip_reservation"); ok {
        parent := &vspk.IPReservation{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_gateway"); ok {
        parent := &vspk.Gateway{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vsd"); ok {
        parent := &vspk.VSD{ID: attr.(string)}
        filteredEventLogs, err = parent.EventLogs(fetchFilter)
        if err != nil {
            return err
        }
    }

    EventLog := &vspk.EventLog{}

    if len(filteredEventLogs) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredEventLogs) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    EventLog = filteredEventLogs[0]

    d.Set("request_id", EventLog.RequestID)
    d.Set("diff", EventLog.Diff)
    d.Set("enterprise", EventLog.Enterprise)
    d.Set("entities", EventLog.Entities)
    d.Set("entity_id", EventLog.EntityID)
    d.Set("entity_parent_id", EventLog.EntityParentID)
    d.Set("entity_parent_type", EventLog.EntityParentType)
    d.Set("entity_scope", EventLog.EntityScope)
    d.Set("entity_type", EventLog.EntityType)
    d.Set("user", EventLog.User)
    d.Set("event_received_time", EventLog.EventReceivedTime)
    d.Set("external_id", EventLog.ExternalID)
    d.Set("type", EventLog.Type)
    
    d.Set("id", EventLog.Identifier())
    d.Set("parent_id", EventLog.ParentID)
    d.Set("parent_type", EventLog.ParentType)
    d.Set("owner", EventLog.Owner)

    d.SetId(EventLog.Identifier())
    
    return nil
}