/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// RedundantPortIdentity represents the Identity of the object
var RedundantPortIdentity = bambou.Identity{
	Name:     "nsredundantport",
	Category: "nsredundantports",
}

// RedundantPortsList represents a list of RedundantPorts
type RedundantPortsList []*RedundantPort

// RedundantPortsAncestor is the interface that an ancestor of a RedundantPort must implement.
// An Ancestor is defined as an entity that has RedundantPort as a descendant.
// An Ancestor can get a list of its child RedundantPorts, but not necessarily create one.
type RedundantPortsAncestor interface {
	RedundantPorts(*bambou.FetchingInfo) (RedundantPortsList, *bambou.Error)
}

// RedundantPortsParent is the interface that a parent of a RedundantPort must implement.
// A Parent is defined as an entity that has RedundantPort as a child.
// A Parent is an Ancestor which can create a RedundantPort.
type RedundantPortsParent interface {
	RedundantPortsAncestor
	CreateRedundantPort(*RedundantPort) *bambou.Error
}

// RedundantPort represents the model of a nsredundantport
type RedundantPort struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	VLANRange                   string `json:"VLANRange,omitempty"`
	MTU                         int    `json:"MTU,omitempty"`
	Name                        string `json:"name,omitempty"`
	LastUpdatedBy               string `json:"lastUpdatedBy,omitempty"`
	PermittedAction             string `json:"permittedAction,omitempty"`
	Description                 string `json:"description,omitempty"`
	PhysicalName                string `json:"physicalName,omitempty"`
	InfrastructureProfileID     string `json:"infrastructureProfileID,omitempty"`
	EntityScope                 string `json:"entityScope,omitempty"`
	PortPeer1ID                 string `json:"portPeer1ID,omitempty"`
	PortPeer2ID                 string `json:"portPeer2ID,omitempty"`
	PortType                    string `json:"portType,omitempty"`
	Speed                       string `json:"speed,omitempty"`
	UseUntaggedHeartbeatVlan    bool   `json:"useUntaggedHeartbeatVlan"`
	UseUserMnemonic             bool   `json:"useUserMnemonic"`
	UserMnemonic                string `json:"userMnemonic,omitempty"`
	AssociatedEgressQOSPolicyID string `json:"associatedEgressQOSPolicyID,omitempty"`
	Status                      string `json:"status,omitempty"`
	ExternalID                  string `json:"externalID,omitempty"`
}

// NewRedundantPort returns a new *RedundantPort
func NewRedundantPort() *RedundantPort {

	return &RedundantPort{
		VLANRange: "0-4094",
		MTU:       1500,
	}
}

// Identity returns the Identity of the object.
func (o *RedundantPort) Identity() bambou.Identity {

	return RedundantPortIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RedundantPort) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RedundantPort) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the RedundantPort from the server
func (o *RedundantPort) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the RedundantPort into the server
func (o *RedundantPort) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the RedundantPort from the server
func (o *RedundantPort) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the RedundantPort
func (o *RedundantPort) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the RedundantPort
func (o *RedundantPort) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VLANs retrieves the list of child VLANs of the RedundantPort
func (o *RedundantPort) VLANs(info *bambou.FetchingInfo) (VLANsList, *bambou.Error) {

	var list VLANsList
	err := bambou.CurrentSession().FetchChildren(o, VLANIdentity, &list, info)
	return list, err
}

// CreateVLAN creates a new child VLAN under the RedundantPort
func (o *RedundantPort) CreateVLAN(child *VLAN) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the RedundantPort
func (o *RedundantPort) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the RedundantPort
func (o *RedundantPort) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// NSPorts retrieves the list of child NSPorts of the RedundantPort
func (o *RedundantPort) NSPorts(info *bambou.FetchingInfo) (NSPortsList, *bambou.Error) {

	var list NSPortsList
	err := bambou.CurrentSession().FetchChildren(o, NSPortIdentity, &list, info)
	return list, err
}
