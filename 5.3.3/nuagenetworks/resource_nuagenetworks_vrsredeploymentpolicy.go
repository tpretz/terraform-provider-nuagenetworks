package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.3.3"
)

func resourceVRSRedeploymentpolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceVRSRedeploymentpolicyCreate,
        Read:   resourceVRSRedeploymentpolicyRead,
        Update: resourceVRSRedeploymentpolicyUpdate,
        Delete: resourceVRSRedeploymentpolicyDelete,
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
            "al_ubr0_status_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "cpu_utilization_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "cpu_utilization_threshold": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
                Computed: true,
            },
            "vrs_corrective_action_delay": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "vrs_process_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "vrsvsc_status_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "redeployment_delay": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "memory_utilization_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "memory_utilization_threshold": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
                Computed: true,
            },
            "deployment_count_threshold": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "jesxmon_process_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "log_disk_utilization_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "log_disk_utilization_threshold": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
                Computed: true,
            },
            "root_disk_utilization_redeployment_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "root_disk_utilization_threshold": &schema.Schema{
                Type:     schema.TypeFloat,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_vcenter_hypervisor": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_data_center", "parent_vcenter_cluster", "parent_vcenter", "parent_vcenter_vrs_config"},
            },
            "parent_vcenter_data_center": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_hypervisor", "parent_vcenter_cluster", "parent_vcenter", "parent_vcenter_vrs_config"},
            },
            "parent_vcenter_cluster": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_hypervisor", "parent_vcenter_data_center", "parent_vcenter", "parent_vcenter_vrs_config"},
            },
            "parent_vcenter": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_hypervisor", "parent_vcenter_data_center", "parent_vcenter_cluster", "parent_vcenter_vrs_config"},
            },
            "parent_vcenter_vrs_config": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vcenter_hypervisor", "parent_vcenter_data_center", "parent_vcenter_cluster", "parent_vcenter"},
            },
        },
    }
}

func resourceVRSRedeploymentpolicyCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize VRSRedeploymentpolicy object
    o := &vspk.VRSRedeploymentpolicy{
    }
    if attr, ok := d.GetOk("al_ubr0_status_redeployment_enabled"); ok {
        o.ALUbr0StatusRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("cpu_utilization_redeployment_enabled"); ok {
        o.CPUUtilizationRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("cpu_utilization_threshold"); ok {
        o.CPUUtilizationThreshold = attr.(float64)
    }
    if attr, ok := d.GetOk("vrs_corrective_action_delay"); ok {
        o.VRSCorrectiveActionDelay = attr.(int)
    }
    if attr, ok := d.GetOk("vrs_process_redeployment_enabled"); ok {
        o.VRSProcessRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("vrsvsc_status_redeployment_enabled"); ok {
        o.VRSVSCStatusRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("redeployment_delay"); ok {
        o.RedeploymentDelay = attr.(int)
    }
    if attr, ok := d.GetOk("memory_utilization_redeployment_enabled"); ok {
        o.MemoryUtilizationRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("memory_utilization_threshold"); ok {
        o.MemoryUtilizationThreshold = attr.(float64)
    }
    if attr, ok := d.GetOk("deployment_count_threshold"); ok {
        o.DeploymentCountThreshold = attr.(int)
    }
    if attr, ok := d.GetOk("jesxmon_process_redeployment_enabled"); ok {
        o.JesxmonProcessRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("log_disk_utilization_redeployment_enabled"); ok {
        o.LogDiskUtilizationRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("log_disk_utilization_threshold"); ok {
        o.LogDiskUtilizationThreshold = attr.(float64)
    }
    if attr, ok := d.GetOk("root_disk_utilization_redeployment_enabled"); ok {
        o.RootDiskUtilizationRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("root_disk_utilization_threshold"); ok {
        o.RootDiskUtilizationThreshold = attr.(float64)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_vcenter_hypervisor"); ok {
        parent := &vspk.VCenterHypervisor{ID: attr.(string)}
        err := parent.CreateVRSRedeploymentpolicy(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vcenter_data_center"); ok {
        parent := &vspk.VCenterDataCenter{ID: attr.(string)}
        err := parent.CreateVRSRedeploymentpolicy(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vcenter_cluster"); ok {
        parent := &vspk.VCenterCluster{ID: attr.(string)}
        err := parent.CreateVRSRedeploymentpolicy(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vcenter"); ok {
        parent := &vspk.VCenter{ID: attr.(string)}
        err := parent.CreateVRSRedeploymentpolicy(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_vcenter_vrs_config"); ok {
        parent := &vspk.VCenterVRSConfig{ID: attr.(string)}
        err := parent.CreateVRSRedeploymentpolicy(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceVRSRedeploymentpolicyRead(d, m)
}

func resourceVRSRedeploymentpolicyRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VRSRedeploymentpolicy{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("al_ubr0_status_redeployment_enabled", o.ALUbr0StatusRedeploymentEnabled)
    d.Set("cpu_utilization_redeployment_enabled", o.CPUUtilizationRedeploymentEnabled)
    d.Set("cpu_utilization_threshold", o.CPUUtilizationThreshold)
    d.Set("vrs_corrective_action_delay", o.VRSCorrectiveActionDelay)
    d.Set("vrs_process_redeployment_enabled", o.VRSProcessRedeploymentEnabled)
    d.Set("vrsvsc_status_redeployment_enabled", o.VRSVSCStatusRedeploymentEnabled)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("redeployment_delay", o.RedeploymentDelay)
    d.Set("memory_utilization_redeployment_enabled", o.MemoryUtilizationRedeploymentEnabled)
    d.Set("memory_utilization_threshold", o.MemoryUtilizationThreshold)
    d.Set("deployment_count_threshold", o.DeploymentCountThreshold)
    d.Set("jesxmon_process_redeployment_enabled", o.JesxmonProcessRedeploymentEnabled)
    d.Set("entity_scope", o.EntityScope)
    d.Set("log_disk_utilization_redeployment_enabled", o.LogDiskUtilizationRedeploymentEnabled)
    d.Set("log_disk_utilization_threshold", o.LogDiskUtilizationThreshold)
    d.Set("root_disk_utilization_redeployment_enabled", o.RootDiskUtilizationRedeploymentEnabled)
    d.Set("root_disk_utilization_threshold", o.RootDiskUtilizationThreshold)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceVRSRedeploymentpolicyUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VRSRedeploymentpolicy{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("al_ubr0_status_redeployment_enabled"); ok {
        o.ALUbr0StatusRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("cpu_utilization_redeployment_enabled"); ok {
        o.CPUUtilizationRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("cpu_utilization_threshold"); ok {
        o.CPUUtilizationThreshold = attr.(float64)
    }
    if attr, ok := d.GetOk("vrs_corrective_action_delay"); ok {
        o.VRSCorrectiveActionDelay = attr.(int)
    }
    if attr, ok := d.GetOk("vrs_process_redeployment_enabled"); ok {
        o.VRSProcessRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("vrsvsc_status_redeployment_enabled"); ok {
        o.VRSVSCStatusRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("redeployment_delay"); ok {
        o.RedeploymentDelay = attr.(int)
    }
    if attr, ok := d.GetOk("memory_utilization_redeployment_enabled"); ok {
        o.MemoryUtilizationRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("memory_utilization_threshold"); ok {
        o.MemoryUtilizationThreshold = attr.(float64)
    }
    if attr, ok := d.GetOk("deployment_count_threshold"); ok {
        o.DeploymentCountThreshold = attr.(int)
    }
    if attr, ok := d.GetOk("jesxmon_process_redeployment_enabled"); ok {
        o.JesxmonProcessRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("log_disk_utilization_redeployment_enabled"); ok {
        o.LogDiskUtilizationRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("log_disk_utilization_threshold"); ok {
        o.LogDiskUtilizationThreshold = attr.(float64)
    }
    if attr, ok := d.GetOk("root_disk_utilization_redeployment_enabled"); ok {
        o.RootDiskUtilizationRedeploymentEnabled = attr.(bool)
    }
    if attr, ok := d.GetOk("root_disk_utilization_threshold"); ok {
        o.RootDiskUtilizationThreshold = attr.(float64)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceVRSRedeploymentpolicyDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.VRSRedeploymentpolicy{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}