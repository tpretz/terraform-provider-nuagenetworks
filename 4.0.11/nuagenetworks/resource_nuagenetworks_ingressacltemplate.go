package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.11"
)

func resourceIngressACLTemplate() *schema.Resource {
    return &schema.Resource{
        Create: resourceIngressACLTemplateCreate,
        Read:   resourceIngressACLTemplateRead,
        Update: resourceIngressACLTemplateUpdate,
        Delete: resourceIngressACLTemplateDelete,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "active": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "default_allow_ip": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "default_allow_non_ip": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "allow_address_spoof": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "allow_l2_address_spoof": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "policy_state": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
                Computed: true,
            },
            "priority_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "assoc_acl_template_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_live_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "auto_generate_priority": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_l2_domain_template", "parent_domain_template"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain_template", "parent_domain_template"},
            },
            "parent_l2_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_domain_template"},
            },
            "parent_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_l2_domain_template"},
            },
        },
    }
}

func resourceIngressACLTemplateCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize IngressACLTemplate object
    o := &vspk.IngressACLTemplate{
        Name: d.Get("name").(string),
    }
    if attr, ok := d.GetOk("active"); ok {
        o.Active = attr.(bool)
    }
    if attr, ok := d.GetOk("default_allow_ip"); ok {
        o.DefaultAllowIP = attr.(bool)
    }
    if attr, ok := d.GetOk("default_allow_non_ip"); ok {
        o.DefaultAllowNonIP = attr.(bool)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("allow_address_spoof"); ok {
        o.AllowAddressSpoof = attr.(bool)
    }
    if attr, ok := d.GetOk("allow_l2_address_spoof"); ok {
        o.AllowL2AddressSpoof = attr.(bool)
    }
    if attr, ok := d.GetOk("policy_state"); ok {
        o.PolicyState = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("priority_type"); ok {
        o.PriorityType = attr.(string)
    }
    if attr, ok := d.GetOk("assoc_acl_template_id"); ok {
        o.AssocAclTemplateId = attr.(string)
    }
    if attr, ok := d.GetOk("associated_live_entity_id"); ok {
        o.AssociatedLiveEntityID = attr.(string)
    }
    if attr, ok := d.GetOk("auto_generate_priority"); ok {
        o.AutoGeneratePriority = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        err := parent.CreateIngressACLTemplate(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        err := parent.CreateIngressACLTemplate(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
        parent := &vspk.L2DomainTemplate{ID: attr.(string)}
        err := parent.CreateIngressACLTemplate(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_domain_template"); ok {
        parent := &vspk.DomainTemplate{ID: attr.(string)}
        err := parent.CreateIngressACLTemplate(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceIngressACLTemplateRead(d, m)
}

func resourceIngressACLTemplateRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IngressACLTemplate{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("active", o.Active)
    d.Set("default_allow_ip", o.DefaultAllowIP)
    d.Set("default_allow_non_ip", o.DefaultAllowNonIP)
    d.Set("description", o.Description)
    d.Set("allow_address_spoof", o.AllowAddressSpoof)
    d.Set("allow_l2_address_spoof", o.AllowL2AddressSpoof)
    d.Set("entity_scope", o.EntityScope)
    d.Set("policy_state", o.PolicyState)
    d.Set("priority", o.Priority)
    d.Set("priority_type", o.PriorityType)
    d.Set("assoc_acl_template_id", o.AssocAclTemplateId)
    d.Set("associated_live_entity_id", o.AssociatedLiveEntityID)
    d.Set("auto_generate_priority", o.AutoGeneratePriority)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceIngressACLTemplateUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IngressACLTemplate{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    
    if attr, ok := d.GetOk("active"); ok {
        o.Active = attr.(bool)
    }
    if attr, ok := d.GetOk("default_allow_ip"); ok {
        o.DefaultAllowIP = attr.(bool)
    }
    if attr, ok := d.GetOk("default_allow_non_ip"); ok {
        o.DefaultAllowNonIP = attr.(bool)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("allow_address_spoof"); ok {
        o.AllowAddressSpoof = attr.(bool)
    }
    if attr, ok := d.GetOk("allow_l2_address_spoof"); ok {
        o.AllowL2AddressSpoof = attr.(bool)
    }
    if attr, ok := d.GetOk("policy_state"); ok {
        o.PolicyState = attr.(string)
    }
    if attr, ok := d.GetOk("priority"); ok {
        o.Priority = attr.(int)
    }
    if attr, ok := d.GetOk("priority_type"); ok {
        o.PriorityType = attr.(string)
    }
    if attr, ok := d.GetOk("assoc_acl_template_id"); ok {
        o.AssocAclTemplateId = attr.(string)
    }
    if attr, ok := d.GetOk("associated_live_entity_id"); ok {
        o.AssociatedLiveEntityID = attr.(string)
    }
    if attr, ok := d.GetOk("auto_generate_priority"); ok {
        o.AutoGeneratePriority = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceIngressACLTemplateDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.IngressACLTemplate{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}