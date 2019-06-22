/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NSPortInfoIdentity represents the Identity of the object
var NSPortInfoIdentity = bambou.Identity{
	Name:     "portinfo",
	Category: "portinfos",
}

// NSPortInfosList represents a list of NSPortInfos
type NSPortInfosList []*NSPortInfo

// NSPortInfosAncestor is the interface that an ancestor of a NSPortInfo must implement.
// An Ancestor is defined as an entity that has NSPortInfo as a descendant.
// An Ancestor can get a list of its child NSPortInfos, but not necessarily create one.
type NSPortInfosAncestor interface {
	NSPortInfos(*bambou.FetchingInfo) (NSPortInfosList, *bambou.Error)
}

// NSPortInfosParent is the interface that a parent of a NSPortInfo must implement.
// A Parent is defined as an entity that has NSPortInfo as a child.
// A Parent is an Ancestor which can create a NSPortInfo.
type NSPortInfosParent interface {
	NSPortInfosAncestor
	CreateNSPortInfo(*NSPortInfo) *bambou.Error
}

// NSPortInfo represents the model of a portinfo
type NSPortInfo struct {
	ID              string        `json:"ID,omitempty"`
	ParentID        string        `json:"parentID,omitempty"`
	ParentType      string        `json:"parentType,omitempty"`
	Owner           string        `json:"owner,omitempty"`
	WirelessPorts   []interface{} `json:"wirelessPorts,omitempty"`
	MonitoringPorts []interface{} `json:"monitoringPorts,omitempty"`
	Ports           []interface{} `json:"ports,omitempty"`
}

// NewNSPortInfo returns a new *NSPortInfo
func NewNSPortInfo() *NSPortInfo {

	return &NSPortInfo{}
}

// Identity returns the Identity of the object.
func (o *NSPortInfo) Identity() bambou.Identity {

	return NSPortInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSPortInfo) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSPortInfo) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSPortInfo from the server
func (o *NSPortInfo) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSPortInfo into the server
func (o *NSPortInfo) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSPortInfo from the server
func (o *NSPortInfo) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
