/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// SAPIngressQoSProfileIdentity represents the Identity of the object
var SAPIngressQoSProfileIdentity = bambou.Identity{
	Name:     "sapingressqosprofile",
	Category: "sapingressqosprofiles",
}

// SAPIngressQoSProfilesList represents a list of SAPIngressQoSProfiles
type SAPIngressQoSProfilesList []*SAPIngressQoSProfile

// SAPIngressQoSProfilesAncestor is the interface that an ancestor of a SAPIngressQoSProfile must implement.
// An Ancestor is defined as an entity that has SAPIngressQoSProfile as a descendant.
// An Ancestor can get a list of its child SAPIngressQoSProfiles, but not necessarily create one.
type SAPIngressQoSProfilesAncestor interface {
	SAPIngressQoSProfiles(*bambou.FetchingInfo) (SAPIngressQoSProfilesList, *bambou.Error)
}

// SAPIngressQoSProfilesParent is the interface that a parent of a SAPIngressQoSProfile must implement.
// A Parent is defined as an entity that has SAPIngressQoSProfile as a child.
// A Parent is an Ancestor which can create a SAPIngressQoSProfile.
type SAPIngressQoSProfilesParent interface {
	SAPIngressQoSProfilesAncestor
	CreateSAPIngressQoSProfile(*SAPIngressQoSProfile) *bambou.Error
}

// SAPIngressQoSProfile represents the model of a sapingressqosprofile
type SAPIngressQoSProfile struct {
	ID              string `json:"ID,omitempty"`
	ParentID        string `json:"parentID,omitempty"`
	ParentType      string `json:"parentType,omitempty"`
	Owner           string `json:"owner,omitempty"`
	Name            string `json:"name,omitempty"`
	LastUpdatedBy   string `json:"lastUpdatedBy,omitempty"`
	Description     string `json:"description,omitempty"`
	EntityScope     string `json:"entityScope,omitempty"`
	AssocEntityType string `json:"assocEntityType,omitempty"`
	ExternalID      string `json:"externalID,omitempty"`
}

// NewSAPIngressQoSProfile returns a new *SAPIngressQoSProfile
func NewSAPIngressQoSProfile() *SAPIngressQoSProfile {

	return &SAPIngressQoSProfile{}
}

// Identity returns the Identity of the object.
func (o *SAPIngressQoSProfile) Identity() bambou.Identity {

	return SAPIngressQoSProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SAPIngressQoSProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SAPIngressQoSProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the SAPIngressQoSProfile from the server
func (o *SAPIngressQoSProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the SAPIngressQoSProfile into the server
func (o *SAPIngressQoSProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the SAPIngressQoSProfile from the server
func (o *SAPIngressQoSProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the SAPIngressQoSProfile
func (o *SAPIngressQoSProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the SAPIngressQoSProfile
func (o *SAPIngressQoSProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the SAPIngressQoSProfile
func (o *SAPIngressQoSProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the SAPIngressQoSProfile
func (o *SAPIngressQoSProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
