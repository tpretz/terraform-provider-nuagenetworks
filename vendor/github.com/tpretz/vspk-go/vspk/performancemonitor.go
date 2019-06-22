/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PerformanceMonitorIdentity represents the Identity of the object
var PerformanceMonitorIdentity = bambou.Identity{
	Name:     "performancemonitor",
	Category: "performancemonitors",
}

// PerformanceMonitorsList represents a list of PerformanceMonitors
type PerformanceMonitorsList []*PerformanceMonitor

// PerformanceMonitorsAncestor is the interface that an ancestor of a PerformanceMonitor must implement.
// An Ancestor is defined as an entity that has PerformanceMonitor as a descendant.
// An Ancestor can get a list of its child PerformanceMonitors, but not necessarily create one.
type PerformanceMonitorsAncestor interface {
	PerformanceMonitors(*bambou.FetchingInfo) (PerformanceMonitorsList, *bambou.Error)
}

// PerformanceMonitorsParent is the interface that a parent of a PerformanceMonitor must implement.
// A Parent is defined as an entity that has PerformanceMonitor as a child.
// A Parent is an Ancestor which can create a PerformanceMonitor.
type PerformanceMonitorsParent interface {
	PerformanceMonitorsAncestor
	CreatePerformanceMonitor(*PerformanceMonitor) *bambou.Error
}

// PerformanceMonitor represents the model of a performancemonitor
type PerformanceMonitor struct {
	ID                    string        `json:"ID,omitempty"`
	ParentID              string        `json:"parentID,omitempty"`
	ParentType            string        `json:"parentType,omitempty"`
	Owner                 string        `json:"owner,omitempty"`
	Name                  string        `json:"name,omitempty"`
	LastUpdatedBy         string        `json:"lastUpdatedBy,omitempty"`
	PayloadSize           int           `json:"payloadSize,omitempty"`
	ReadOnly              bool          `json:"readOnly"`
	ServiceClass          string        `json:"serviceClass,omitempty"`
	Description           string        `json:"description,omitempty"`
	DestinationTargetList []interface{} `json:"destinationTargetList,omitempty"`
	Timeout               int           `json:"timeout,omitempty"`
	Interval              int           `json:"interval,omitempty"`
	EntityScope           string        `json:"entityScope,omitempty"`
	HoldDownTimer         int           `json:"holdDownTimer,omitempty"`
	ProbeType             string        `json:"probeType,omitempty"`
	NumberOfPackets       int           `json:"numberOfPackets,omitempty"`
	ExternalID            string        `json:"externalID,omitempty"`
}

// NewPerformanceMonitor returns a new *PerformanceMonitor
func NewPerformanceMonitor() *PerformanceMonitor {

	return &PerformanceMonitor{
		PayloadSize:     137,
		ReadOnly:        false,
		ServiceClass:    "H",
		Timeout:         1000,
		Interval:        3000,
		HoldDownTimer:   1000,
		ProbeType:       "ONEWAY",
		NumberOfPackets: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PerformanceMonitor) Identity() bambou.Identity {

	return PerformanceMonitorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PerformanceMonitor) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PerformanceMonitor) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PerformanceMonitor from the server
func (o *PerformanceMonitor) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PerformanceMonitor into the server
func (o *PerformanceMonitor) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PerformanceMonitor from the server
func (o *PerformanceMonitor) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Tiers retrieves the list of child Tiers of the PerformanceMonitor
func (o *PerformanceMonitor) Tiers(info *bambou.FetchingInfo) (TiersList, *bambou.Error) {

	var list TiersList
	err := bambou.CurrentSession().FetchChildren(o, TierIdentity, &list, info)
	return list, err
}

// Applicationperformancemanagements retrieves the list of child Applicationperformancemanagements of the PerformanceMonitor
func (o *PerformanceMonitor) Applicationperformancemanagements(info *bambou.FetchingInfo) (ApplicationperformancemanagementsList, *bambou.Error) {

	var list ApplicationperformancemanagementsList
	err := bambou.CurrentSession().FetchChildren(o, ApplicationperformancemanagementIdentity, &list, info)
	return list, err
}

// NSGateways retrieves the list of child NSGateways of the PerformanceMonitor
func (o *PerformanceMonitor) NSGateways(info *bambou.FetchingInfo) (NSGatewaysList, *bambou.Error) {

	var list NSGatewaysList
	err := bambou.CurrentSession().FetchChildren(o, NSGatewayIdentity, &list, info)
	return list, err
}

// AssignNSGateways assigns the list of NSGateways to the PerformanceMonitor
func (o *PerformanceMonitor) AssignNSGateways(children NSGatewaysList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, NSGatewayIdentity)
}
