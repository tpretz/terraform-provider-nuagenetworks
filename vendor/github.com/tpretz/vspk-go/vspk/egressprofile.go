/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// EgressProfileIdentity represents the Identity of the object
var EgressProfileIdentity = bambou.Identity{
	Name:     "egressprofile",
	Category: "egressprofiles",
}

// EgressProfilesList represents a list of EgressProfiles
type EgressProfilesList []*EgressProfile

// EgressProfilesAncestor is the interface that an ancestor of a EgressProfile must implement.
// An Ancestor is defined as an entity that has EgressProfile as a descendant.
// An Ancestor can get a list of its child EgressProfiles, but not necessarily create one.
type EgressProfilesAncestor interface {
	EgressProfiles(*bambou.FetchingInfo) (EgressProfilesList, *bambou.Error)
}

// EgressProfilesParent is the interface that a parent of a EgressProfile must implement.
// A Parent is defined as an entity that has EgressProfile as a child.
// A Parent is an Ancestor which can create a EgressProfile.
type EgressProfilesParent interface {
	EgressProfilesAncestor
	CreateEgressProfile(*EgressProfile) *bambou.Error
}

// EgressProfile represents the model of a egressprofile
type EgressProfile struct {
	ID                                string `json:"ID,omitempty"`
	ParentID                          string `json:"parentID,omitempty"`
	ParentType                        string `json:"parentType,omitempty"`
	Owner                             string `json:"owner,omitempty"`
	Name                              string `json:"name,omitempty"`
	LastUpdatedBy                     string `json:"lastUpdatedBy,omitempty"`
	Description                       string `json:"description,omitempty"`
	EntityScope                       string `json:"entityScope,omitempty"`
	AssocEntityType                   string `json:"assocEntityType,omitempty"`
	AssociatedIPFilterProfileID       string `json:"associatedIPFilterProfileID,omitempty"`
	AssociatedIPFilterProfileName     string `json:"associatedIPFilterProfileName,omitempty"`
	AssociatedIPv6FilterProfileID     string `json:"associatedIPv6FilterProfileID,omitempty"`
	AssociatedIPv6FilterProfileName   string `json:"associatedIPv6FilterProfileName,omitempty"`
	AssociatedMACFilterProfileID      string `json:"associatedMACFilterProfileID,omitempty"`
	AssociatedMACFilterProfileName    string `json:"associatedMACFilterProfileName,omitempty"`
	AssociatedSAPEgressQoSProfileID   string `json:"associatedSAPEgressQoSProfileID,omitempty"`
	AssociatedSAPEgressQoSProfileName string `json:"associatedSAPEgressQoSProfileName,omitempty"`
	ExternalID                        string `json:"externalID,omitempty"`
}

// NewEgressProfile returns a new *EgressProfile
func NewEgressProfile() *EgressProfile {

	return &EgressProfile{}
}

// Identity returns the Identity of the object.
func (o *EgressProfile) Identity() bambou.Identity {

	return EgressProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EgressProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EgressProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EgressProfile from the server
func (o *EgressProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EgressProfile into the server
func (o *EgressProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EgressProfile from the server
func (o *EgressProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// DeploymentFailures retrieves the list of child DeploymentFailures of the EgressProfile
func (o *EgressProfile) DeploymentFailures(info *bambou.FetchingInfo) (DeploymentFailuresList, *bambou.Error) {

	var list DeploymentFailuresList
	err := bambou.CurrentSession().FetchChildren(o, DeploymentFailureIdentity, &list, info)
	return list, err
}

// CreateDeploymentFailure creates a new child DeploymentFailure under the EgressProfile
func (o *EgressProfile) CreateDeploymentFailure(child *DeploymentFailure) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the EgressProfile
func (o *EgressProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EgressProfile
func (o *EgressProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EgressProfile
func (o *EgressProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EgressProfile
func (o *EgressProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VPorts retrieves the list of child VPorts of the EgressProfile
func (o *EgressProfile) VPorts(info *bambou.FetchingInfo) (VPortsList, *bambou.Error) {

	var list VPortsList
	err := bambou.CurrentSession().FetchChildren(o, VPortIdentity, &list, info)
	return list, err
}
