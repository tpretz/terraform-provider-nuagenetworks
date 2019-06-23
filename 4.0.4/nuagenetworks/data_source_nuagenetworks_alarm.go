package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    vspk "github.com/tpretz/vspk-go/vspk/4.0.4"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceAlarm() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceAlarmRead,
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
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "target_object": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "acknowledged": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "reason": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "severity": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "timestamp": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "error_condition": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "number_of_occurances": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceAlarmRead(d *schema.ResourceData, m interface{}) error {
    filteredAlarms := vspk.AlarmsList{}
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
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    filteredAlarms, err = parent.Alarms(fetchFilter)
    if err != nil {
        return err
    }

    Alarm := &vspk.Alarm{}

    if len(filteredAlarms) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredAlarms) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Alarm = filteredAlarms[0]

    d.Set("name", Alarm.Name)
    d.Set("target_object", Alarm.TargetObject)
    d.Set("last_updated_by", Alarm.LastUpdatedBy)
    d.Set("acknowledged", Alarm.Acknowledged)
    d.Set("reason", Alarm.Reason)
    d.Set("description", Alarm.Description)
    d.Set("severity", Alarm.Severity)
    d.Set("timestamp", Alarm.Timestamp)
    d.Set("enterprise_id", Alarm.EnterpriseID)
    d.Set("entity_scope", Alarm.EntityScope)
    d.Set("error_condition", Alarm.ErrorCondition)
    d.Set("number_of_occurances", Alarm.NumberOfOccurances)
    d.Set("external_id", Alarm.ExternalID)
    
    d.Set("id", Alarm.Identifier())
    d.Set("parent_id", Alarm.ParentID)
    d.Set("parent_type", Alarm.ParentType)
    d.Set("owner", Alarm.Owner)

    d.SetId(Alarm.Identifier())
    
    return nil
}