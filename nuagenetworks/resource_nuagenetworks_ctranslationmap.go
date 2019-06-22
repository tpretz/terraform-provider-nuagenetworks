package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/tpretz/vspk-go/vspk"
)

func resourceCTranslationMap() *schema.Resource {
    return &schema.Resource{
        Create: resourceCTranslationMapCreate,
        Read:   resourceCTranslationMapRead,
        Update: resourceCTranslationMapUpdate,
        Delete: resourceCTranslationMapDelete,
        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },
        Schema: map[string]*schema.Schema{
            "id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
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
            "mapping_type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "associated_domain_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "customer_alias_ip": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "customer_ip": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "parent_csnat_pool": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceCTranslationMapCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize CTranslationMap object
    o := &vspk.CTranslationMap{
        MappingType: d.Get("mapping_type").(string),
        CustomerAliasIP: d.Get("customer_alias_ip").(string),
        CustomerIP: d.Get("customer_ip").(string),
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := &vspk.CSNATPool{ID: d.Get("parent_csnat_pool").(string)}
    err := parent.CreateCTranslationMap(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceCTranslationMapRead(d, m)
}

func resourceCTranslationMapRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.CTranslationMap{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("mapping_type", o.MappingType)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("entity_scope", o.EntityScope)
    d.Set("associated_domain_id", o.AssociatedDomainID)
    d.Set("customer_alias_ip", o.CustomerAliasIP)
    d.Set("customer_ip", o.CustomerIP)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceCTranslationMapUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.CTranslationMap{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.MappingType = d.Get("mapping_type").(string)
    o.CustomerAliasIP = d.Get("customer_alias_ip").(string)
    o.CustomerIP = d.Get("customer_ip").(string)
    
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceCTranslationMapDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.CTranslationMap{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}