/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NSGatewayMonitorIdentity represents the Identity of the object
var NSGatewayMonitorIdentity = bambou.Identity{
	Name:     "nsgatewaysmonitor",
	Category: "nsgatewaysmonitors",
}

// NSGatewayMonitorsList represents a list of NSGatewayMonitors
type NSGatewayMonitorsList []*NSGatewayMonitor

// NSGatewayMonitorsAncestor is the interface that an ancestor of a NSGatewayMonitor must implement.
// An Ancestor is defined as an entity that has NSGatewayMonitor as a descendant.
// An Ancestor can get a list of its child NSGatewayMonitors, but not necessarily create one.
type NSGatewayMonitorsAncestor interface {
	NSGatewayMonitors(*bambou.FetchingInfo) (NSGatewayMonitorsList, *bambou.Error)
}

// NSGatewayMonitorsParent is the interface that a parent of a NSGatewayMonitor must implement.
// A Parent is defined as an entity that has NSGatewayMonitor as a child.
// A Parent is an Ancestor which can create a NSGatewayMonitor.
type NSGatewayMonitorsParent interface {
	NSGatewayMonitorsAncestor
	CreateNSGatewayMonitor(*NSGatewayMonitor) *bambou.Error
}

// NSGatewayMonitor represents the model of a nsgatewaysmonitor
type NSGatewayMonitor struct {
	ID         string        `json:"ID,omitempty"`
	ParentID   string        `json:"parentID,omitempty"`
	ParentType string        `json:"parentType,omitempty"`
	Owner      string        `json:"owner,omitempty"`
	Vrsinfo    interface{}   `json:"vrsinfo,omitempty"`
	Vscs       []interface{} `json:"vscs,omitempty"`
	Nsginfo    interface{}   `json:"nsginfo,omitempty"`
	Nsgstate   interface{}   `json:"nsgstate,omitempty"`
	Nsgsummary interface{}   `json:"nsgsummary,omitempty"`
}

// NewNSGatewayMonitor returns a new *NSGatewayMonitor
func NewNSGatewayMonitor() *NSGatewayMonitor {

	return &NSGatewayMonitor{}
}

// Identity returns the Identity of the object.
func (o *NSGatewayMonitor) Identity() bambou.Identity {

	return NSGatewayMonitorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSGatewayMonitor) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSGatewayMonitor) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSGatewayMonitor from the server
func (o *NSGatewayMonitor) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSGatewayMonitor into the server
func (o *NSGatewayMonitor) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSGatewayMonitor from the server
func (o *NSGatewayMonitor) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
