/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// FloatingIpIdentity represents the Identity of the object
var FloatingIpIdentity = bambou.Identity{
	Name:     "floatingip",
	Category: "floatingips",
}

// FloatingIpsList represents a list of FloatingIps
type FloatingIpsList []*FloatingIp

// FloatingIpsAncestor is the interface that an ancestor of a FloatingIp must implement.
// An Ancestor is defined as an entity that has FloatingIp as a descendant.
// An Ancestor can get a list of its child FloatingIps, but not necessarily create one.
type FloatingIpsAncestor interface {
	FloatingIps(*bambou.FetchingInfo) (FloatingIpsList, *bambou.Error)
}

// FloatingIpsParent is the interface that a parent of a FloatingIp must implement.
// A Parent is defined as an entity that has FloatingIp as a child.
// A Parent is an Ancestor which can create a FloatingIp.
type FloatingIpsParent interface {
	FloatingIpsAncestor
	CreateFloatingIp(*FloatingIp) *bambou.Error
}

// FloatingIp represents the model of a floatingip
type FloatingIp struct {
	ID                                string `json:"ID,omitempty"`
	ParentID                          string `json:"parentID,omitempty"`
	ParentType                        string `json:"parentType,omitempty"`
	Owner                             string `json:"owner,omitempty"`
	LastUpdatedBy                     string `json:"lastUpdatedBy,omitempty"`
	AccessControl                     bool   `json:"accessControl"`
	Address                           string `json:"address,omitempty"`
	EntityScope                       string `json:"entityScope,omitempty"`
	Assigned                          bool   `json:"assigned"`
	AssignedToObjectType              string `json:"assignedToObjectType,omitempty"`
	AssociatedSharedNetworkResourceID string `json:"associatedSharedNetworkResourceID,omitempty"`
	ExternalID                        string `json:"externalID,omitempty"`
}

// NewFloatingIp returns a new *FloatingIp
func NewFloatingIp() *FloatingIp {

	return &FloatingIp{}
}

// Identity returns the Identity of the object.
func (o *FloatingIp) Identity() bambou.Identity {

	return FloatingIpIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FloatingIp) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FloatingIp) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the FloatingIp from the server
func (o *FloatingIp) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the FloatingIp into the server
func (o *FloatingIp) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the FloatingIp from the server
func (o *FloatingIp) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the FloatingIp
func (o *FloatingIp) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the FloatingIp
func (o *FloatingIp) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the FloatingIp
func (o *FloatingIp) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the FloatingIp
func (o *FloatingIp) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VPorts retrieves the list of child VPorts of the FloatingIp
func (o *FloatingIp) VPorts(info *bambou.FetchingInfo) (VPortsList, *bambou.Error) {

	var list VPortsList
	err := bambou.CurrentSession().FetchChildren(o, VPortIdentity, &list, info)
	return list, err
}

// EventLogs retrieves the list of child EventLogs of the FloatingIp
func (o *FloatingIp) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
