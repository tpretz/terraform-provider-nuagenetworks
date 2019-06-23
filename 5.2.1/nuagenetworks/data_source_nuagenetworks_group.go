package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/5.2.1"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceGroup() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceGroupRead,
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
            "ldap_group_dn": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "management_mode": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "account_restrictions": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "restriction_date": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "role": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "private": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_l2_domain_template", "parent_user", "parent_enterprise", "parent_zone", "parent_domain_template"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain_template", "parent_user", "parent_enterprise", "parent_zone", "parent_domain_template"},
            },
            "parent_l2_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_user", "parent_enterprise", "parent_zone", "parent_domain_template"},
            },
            "parent_user": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_l2_domain_template", "parent_enterprise", "parent_zone", "parent_domain_template"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_l2_domain_template", "parent_user", "parent_zone", "parent_domain_template"},
            },
            "parent_zone": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_l2_domain_template", "parent_user", "parent_enterprise", "parent_domain_template"},
            },
            "parent_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_l2_domain_template", "parent_user", "parent_enterprise", "parent_zone"},
            },
        },
    }
}


func dataSourceGroupRead(d *schema.ResourceData, m interface{}) error {
    filteredGroups := vspk.GroupsList{}
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
        filteredGroups, err = parent.Groups(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredGroups, err = parent.Groups(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
        parent := &vspk.L2DomainTemplate{ID: attr.(string)}
        filteredGroups, err = parent.Groups(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_user"); ok {
        parent := &vspk.User{ID: attr.(string)}
        filteredGroups, err = parent.Groups(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredGroups, err = parent.Groups(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_zone"); ok {
        parent := &vspk.Zone{ID: attr.(string)}
        filteredGroups, err = parent.Groups(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_domain_template"); ok {
        parent := &vspk.DomainTemplate{ID: attr.(string)}
        filteredGroups, err = parent.Groups(fetchFilter)
        if err != nil {
            return err
        }
    }

    Group := &vspk.Group{}

    if len(filteredGroups) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredGroups) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Group = filteredGroups[0]

    d.Set("ldap_group_dn", Group.LDAPGroupDN)
    d.Set("name", Group.Name)
    d.Set("management_mode", Group.ManagementMode)
    d.Set("last_updated_by", Group.LastUpdatedBy)
    d.Set("account_restrictions", Group.AccountRestrictions)
    d.Set("description", Group.Description)
    d.Set("restriction_date", Group.RestrictionDate)
    d.Set("entity_scope", Group.EntityScope)
    d.Set("role", Group.Role)
    d.Set("private", Group.Private)
    d.Set("external_id", Group.ExternalID)
    
    d.Set("id", Group.Identifier())
    d.Set("parent_id", Group.ParentID)
    d.Set("parent_type", Group.ParentType)
    d.Set("owner", Group.Owner)

    d.SetId(Group.Identifier())
    
    return nil
}