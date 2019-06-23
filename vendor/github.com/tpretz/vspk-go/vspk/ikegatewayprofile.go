/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// IKEGatewayProfileIdentity represents the Identity of the object
var IKEGatewayProfileIdentity = bambou.Identity{
	Name:     "ikegatewayprofile",
	Category: "ikegatewayprofiles",
}

// IKEGatewayProfilesList represents a list of IKEGatewayProfiles
type IKEGatewayProfilesList []*IKEGatewayProfile

// IKEGatewayProfilesAncestor is the interface that an ancestor of a IKEGatewayProfile must implement.
// An Ancestor is defined as an entity that has IKEGatewayProfile as a descendant.
// An Ancestor can get a list of its child IKEGatewayProfiles, but not necessarily create one.
type IKEGatewayProfilesAncestor interface {
	IKEGatewayProfiles(*bambou.FetchingInfo) (IKEGatewayProfilesList, *bambou.Error)
}

// IKEGatewayProfilesParent is the interface that a parent of a IKEGatewayProfile must implement.
// A Parent is defined as an entity that has IKEGatewayProfile as a child.
// A Parent is an Ancestor which can create a IKEGatewayProfile.
type IKEGatewayProfilesParent interface {
	IKEGatewayProfilesAncestor
	CreateIKEGatewayProfile(*IKEGatewayProfile) *bambou.Error
}

// IKEGatewayProfile represents the model of a ikegatewayprofile
type IKEGatewayProfile struct {
	ID                               string `json:"ID,omitempty"`
	ParentID                         string `json:"parentID,omitempty"`
	ParentType                       string `json:"parentType,omitempty"`
	Owner                            string `json:"owner,omitempty"`
	IKEGatewayIdentifier             string `json:"IKEGatewayIdentifier,omitempty"`
	IKEGatewayIdentifierType         string `json:"IKEGatewayIdentifierType,omitempty"`
	Name                             string `json:"name,omitempty"`
	LastUpdatedBy                    string `json:"lastUpdatedBy,omitempty"`
	ServiceClass                     string `json:"serviceClass,omitempty"`
	Description                      string `json:"description,omitempty"`
	AntiReplayCheck                  bool   `json:"antiReplayCheck"`
	EntityScope                      string `json:"entityScope,omitempty"`
	AssociatedEnterpriseID           string `json:"associatedEnterpriseID,omitempty"`
	AssociatedIKEAuthenticationID    string `json:"associatedIKEAuthenticationID,omitempty"`
	AssociatedIKEAuthenticationType  string `json:"associatedIKEAuthenticationType,omitempty"`
	AssociatedIKEEncryptionProfileID string `json:"associatedIKEEncryptionProfileID,omitempty"`
	AssociatedIKEGatewayID           string `json:"associatedIKEGatewayID,omitempty"`
	ExternalID                       string `json:"externalID,omitempty"`
}

// NewIKEGatewayProfile returns a new *IKEGatewayProfile
func NewIKEGatewayProfile() *IKEGatewayProfile {

	return &IKEGatewayProfile{
		IKEGatewayIdentifierType: "ID_IPV4_ADDR",
	}
}

// Identity returns the Identity of the object.
func (o *IKEGatewayProfile) Identity() bambou.Identity {

	return IKEGatewayProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IKEGatewayProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IKEGatewayProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IKEGatewayProfile from the server
func (o *IKEGatewayProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IKEGatewayProfile into the server
func (o *IKEGatewayProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IKEGatewayProfile from the server
func (o *IKEGatewayProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IKEGatewayProfile
func (o *IKEGatewayProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IKEGatewayProfile
func (o *IKEGatewayProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IKEGatewayProfile
func (o *IKEGatewayProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IKEGatewayProfile
func (o *IKEGatewayProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
