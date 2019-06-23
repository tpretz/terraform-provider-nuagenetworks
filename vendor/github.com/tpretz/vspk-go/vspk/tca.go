/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// TCAIdentity represents the Identity of the object
var TCAIdentity = bambou.Identity{
	Name:     "tca",
	Category: "tcas",
}

// TCAsList represents a list of TCAs
type TCAsList []*TCA

// TCAsAncestor is the interface that an ancestor of a TCA must implement.
// An Ancestor is defined as an entity that has TCA as a descendant.
// An Ancestor can get a list of its child TCAs, but not necessarily create one.
type TCAsAncestor interface {
	TCAs(*bambou.FetchingInfo) (TCAsList, *bambou.Error)
}

// TCAsParent is the interface that a parent of a TCA must implement.
// A Parent is defined as an entity that has TCA as a child.
// A Parent is an Ancestor which can create a TCA.
type TCAsParent interface {
	TCAsAncestor
	CreateTCA(*TCA) *bambou.Error
}

// TCA represents the model of a tca
type TCA struct {
	ID                  string `json:"ID,omitempty"`
	ParentID            string `json:"parentID,omitempty"`
	ParentType          string `json:"parentType,omitempty"`
	Owner               string `json:"owner,omitempty"`
	URLEndPoint         string `json:"URLEndPoint,omitempty"`
	Name                string `json:"name,omitempty"`
	TargetPolicyGroupID string `json:"targetPolicyGroupID,omitempty"`
	LastUpdatedBy       string `json:"lastUpdatedBy,omitempty"`
	Action              string `json:"action,omitempty"`
	Period              int    `json:"period,omitempty"`
	Description         string `json:"description,omitempty"`
	Metric              string `json:"metric,omitempty"`
	Threshold           int    `json:"threshold,omitempty"`
	ThrottleTime        int    `json:"throttleTime,omitempty"`
	Disable             bool   `json:"disable"`
	DisplayStatus       string `json:"displayStatus,omitempty"`
	EntityScope         string `json:"entityScope,omitempty"`
	Count               int    `json:"count,omitempty"`
	Status              bool   `json:"status"`
	ExternalID          string `json:"externalID,omitempty"`
	Type                string `json:"type,omitempty"`
}

// NewTCA returns a new *TCA
func NewTCA() *TCA {

	return &TCA{
		ThrottleTime: 0,
		Disable:      false,
		Count:        0,
		Status:       false,
	}
}

// Identity returns the Identity of the object.
func (o *TCA) Identity() bambou.Identity {

	return TCAIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TCA) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TCA) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the TCA from the server
func (o *TCA) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the TCA into the server
func (o *TCA) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the TCA from the server
func (o *TCA) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the TCA
func (o *TCA) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the TCA
func (o *TCA) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the TCA
func (o *TCA) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// CreateAlarm creates a new child Alarm under the TCA
func (o *TCA) CreateAlarm(child *Alarm) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the TCA
func (o *TCA) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the TCA
func (o *TCA) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the TCA
func (o *TCA) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
