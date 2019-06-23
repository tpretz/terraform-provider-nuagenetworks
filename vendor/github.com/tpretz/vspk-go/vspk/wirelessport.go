/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// WirelessPortIdentity represents the Identity of the object
var WirelessPortIdentity = bambou.Identity{
	Name:     "wirelessport",
	Category: "wirelessports",
}

// WirelessPortsList represents a list of WirelessPorts
type WirelessPortsList []*WirelessPort

// WirelessPortsAncestor is the interface that an ancestor of a WirelessPort must implement.
// An Ancestor is defined as an entity that has WirelessPort as a descendant.
// An Ancestor can get a list of its child WirelessPorts, but not necessarily create one.
type WirelessPortsAncestor interface {
	WirelessPorts(*bambou.FetchingInfo) (WirelessPortsList, *bambou.Error)
}

// WirelessPortsParent is the interface that a parent of a WirelessPort must implement.
// A Parent is defined as an entity that has WirelessPort as a child.
// A Parent is an Ancestor which can create a WirelessPort.
type WirelessPortsParent interface {
	WirelessPortsAncestor
	CreateWirelessPort(*WirelessPort) *bambou.Error
}

// WirelessPort represents the model of a wirelessport
type WirelessPort struct {
	ID                string `json:"ID,omitempty"`
	ParentID          string `json:"parentID,omitempty"`
	ParentType        string `json:"parentType,omitempty"`
	Owner             string `json:"owner,omitempty"`
	Name              string `json:"name,omitempty"`
	GenericConfig     string `json:"genericConfig,omitempty"`
	Description       string `json:"description,omitempty"`
	PhysicalName      string `json:"physicalName,omitempty"`
	WifiFrequencyBand string `json:"wifiFrequencyBand,omitempty"`
	WifiMode          string `json:"wifiMode,omitempty"`
	PortType          string `json:"portType,omitempty"`
	CountryCode       string `json:"countryCode,omitempty"`
	FrequencyChannel  string `json:"frequencyChannel,omitempty"`
}

// NewWirelessPort returns a new *WirelessPort
func NewWirelessPort() *WirelessPort {

	return &WirelessPort{
		WifiFrequencyBand: "FREQ_2_4_GHZ",
		WifiMode:          "WIFI_B_G_N",
		PortType:          "ACCESS",
		FrequencyChannel:  "CH_0",
	}
}

// Identity returns the Identity of the object.
func (o *WirelessPort) Identity() bambou.Identity {

	return WirelessPortIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *WirelessPort) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *WirelessPort) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the WirelessPort from the server
func (o *WirelessPort) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the WirelessPort into the server
func (o *WirelessPort) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the WirelessPort from the server
func (o *WirelessPort) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Alarms retrieves the list of child Alarms of the WirelessPort
func (o *WirelessPort) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// SSIDConnections retrieves the list of child SSIDConnections of the WirelessPort
func (o *WirelessPort) SSIDConnections(info *bambou.FetchingInfo) (SSIDConnectionsList, *bambou.Error) {

	var list SSIDConnectionsList
	err := bambou.CurrentSession().FetchChildren(o, SSIDConnectionIdentity, &list, info)
	return list, err
}

// CreateSSIDConnection creates a new child SSIDConnection under the WirelessPort
func (o *WirelessPort) CreateSSIDConnection(child *SSIDConnection) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Statistics retrieves the list of child Statistics of the WirelessPort
func (o *WirelessPort) Statistics(info *bambou.FetchingInfo) (StatisticsList, *bambou.Error) {

	var list StatisticsList
	err := bambou.CurrentSession().FetchChildren(o, StatisticsIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the WirelessPort
func (o *WirelessPort) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
