package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.1.1"
)

func resourceEgressAdvFwdTemplate() *schema.Resource {
    return &schema.Resource{
        Create: resourceEgressAdvFwdTemplateCreate,
        Read:   resourceEgressAdvFwdTemplateRead,
        Update: resourceEgressAdvFwdTemplateUpdate,
        Delete: resourceEgressAdvFwdTemplateDelete,
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
                Optional: true,
                Computed: true,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
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

func resourceEgressAdvFwdTemplateCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize EgressAdvFwdTemplate object
    o := &vspk.EgressAdvFwdTemplate{
    }
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("active"); ok {
        o.Active = attr.(bool)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
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
        err := parent.CreateEgressAdvFwdTemplate(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        err := parent.CreateEgressAdvFwdTemplate(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
        parent := &vspk.L2DomainTemplate{ID: attr.(string)}
        err := parent.CreateEgressAdvFwdTemplate(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_domain_template"); ok {
        parent := &vspk.DomainTemplate{ID: attr.(string)}
        err := parent.CreateEgressAdvFwdTemplate(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    return resourceEgressAdvFwdTemplateRead(d, m)
}

func resourceEgressAdvFwdTemplateRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.EgressAdvFwdTemplate{
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
    d.Set("description", o.Description)
    d.Set("entity_scope", o.EntityScope)
    d.Set("policy_state", o.PolicyState)
    d.Set("priority", o.Priority)
    d.Set("priority_type", o.PriorityType)
    d.Set("associated_live_entity_id", o.AssociatedLiveEntityID)
    d.Set("auto_generate_priority", o.AutoGeneratePriority)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceEgressAdvFwdTemplateUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.EgressAdvFwdTemplate{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    
    if attr, ok := d.GetOk("name"); ok {
        o.Name = attr.(string)
    }
    if attr, ok := d.GetOk("active"); ok {
        o.Active = attr.(bool)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
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

func resourceEgressAdvFwdTemplateDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.EgressAdvFwdTemplate{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}