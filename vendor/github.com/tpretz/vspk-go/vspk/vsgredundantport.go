/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VsgRedundantPortIdentity represents the Identity of the object
var VsgRedundantPortIdentity = bambou.Identity{
	Name:     "vsgredundantport",
	Category: "vsgredundantports",
}

// VsgRedundantPortsList represents a list of VsgRedundantPorts
type VsgRedundantPortsList []*VsgRedundantPort

// VsgRedundantPortsAncestor is the interface that an ancestor of a VsgRedundantPort must implement.
// An Ancestor is defined as an entity that has VsgRedundantPort as a descendant.
// An Ancestor can get a list of its child VsgRedundantPorts, but not necessarily create one.
type VsgRedundantPortsAncestor interface {
	VsgRedundantPorts(*bambou.FetchingInfo) (VsgRedundantPortsList, *bambou.Error)
}

// VsgRedundantPortsParent is the interface that a parent of a VsgRedundantPort must implement.
// A Parent is defined as an entity that has VsgRedundantPort as a child.
// A Parent is an Ancestor which can create a VsgRedundantPort.
type VsgRedundantPortsParent interface {
	VsgRedundantPortsAncestor
	CreateVsgRedundantPort(*VsgRedundantPort) *bambou.Error
}

// VsgRedundantPort represents the model of a vsgredundantport
type VsgRedundantPort struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	VLANRange                   string `json:"VLANRange,omitempty"`
	Name                        string `json:"name,omitempty"`
	LastUpdatedBy               string `json:"lastUpdatedBy,omitempty"`
	PeerLink                    bool   `json:"peerLink"`
	PermittedAction             string `json:"permittedAction,omitempty"`
	Description                 string `json:"description,omitempty"`
	PhysicalName                string `json:"physicalName,omitempty"`
	EntityScope                 string `json:"entityScope,omitempty"`
	PortPeer1ID                 string `json:"portPeer1ID,omitempty"`
	PortPeer2ID                 string `json:"portPeer2ID,omitempty"`
	PortType                    string `json:"portType,omitempty"`
	UseUserMnemonic             bool   `json:"useUserMnemonic"`
	UserMnemonic                string `json:"userMnemonic,omitempty"`
	AssociatedEgressQOSPolicyID string `json:"associatedEgressQOSPolicyID,omitempty"`
	Status                      string `json:"status,omitempty"`
	ExternalID                  string `json:"externalID,omitempty"`
}

// NewVsgRedundantPort returns a new *VsgRedundantPort
func NewVsgRedundantPort() *VsgRedundantPort {

	return &VsgRedundantPort{
		PeerLink: false,
	}
}

// Identity returns the Identity of the object.
func (o *VsgRedundantPort) Identity() bambou.Identity {

	return VsgRedundantPortIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VsgRedundantPort) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VsgRedundantPort) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VsgRedundantPort from the server
func (o *VsgRedundantPort) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VsgRedundantPort into the server
func (o *VsgRedundantPort) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VsgRedundantPort from the server
func (o *VsgRedundantPort) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Permissions retrieves the list of child Permissions of the VsgRedundantPort
func (o *VsgRedundantPort) Permissions(info *bambou.FetchingInfo) (PermissionsList, *bambou.Error) {

	var list PermissionsList
	err := bambou.CurrentSession().FetchChildren(o, PermissionIdentity, &list, info)
	return list, err
}

// CreatePermission creates a new child Permission under the VsgRedundantPort
func (o *VsgRedundantPort) CreatePermission(child *Permission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the VsgRedundantPort
func (o *VsgRedundantPort) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VsgRedundantPort
func (o *VsgRedundantPort) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VLANs retrieves the list of child VLANs of the VsgRedundantPort
func (o *VsgRedundantPort) VLANs(info *bambou.FetchingInfo) (VLANsList, *bambou.Error) {

	var list VLANsList
	err := bambou.CurrentSession().FetchChildren(o, VLANIdentity, &list, info)
	return list, err
}

// CreateVLAN creates a new child VLAN under the VsgRedundantPort
func (o *VsgRedundantPort) CreateVLAN(child *VLAN) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the VsgRedundantPort
func (o *VsgRedundantPort) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VsgRedundantPort
func (o *VsgRedundantPort) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VsgRedundantPort
func (o *VsgRedundantPort) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EnterprisePermissions retrieves the list of child EnterprisePermissions of the VsgRedundantPort
func (o *VsgRedundantPort) EnterprisePermissions(info *bambou.FetchingInfo) (EnterprisePermissionsList, *bambou.Error) {

	var list EnterprisePermissionsList
	err := bambou.CurrentSession().FetchChildren(o, EnterprisePermissionIdentity, &list, info)
	return list, err
}

// CreateEnterprisePermission creates a new child EnterprisePermission under the VsgRedundantPort
func (o *VsgRedundantPort) CreateEnterprisePermission(child *EnterprisePermission) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
