/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// BGPProfileIdentity represents the Identity of the object
var BGPProfileIdentity = bambou.Identity{
	Name:     "bgpprofile",
	Category: "bgpprofiles",
}

// BGPProfilesList represents a list of BGPProfiles
type BGPProfilesList []*BGPProfile

// BGPProfilesAncestor is the interface that an ancestor of a BGPProfile must implement.
// An Ancestor is defined as an entity that has BGPProfile as a descendant.
// An Ancestor can get a list of its child BGPProfiles, but not necessarily create one.
type BGPProfilesAncestor interface {
	BGPProfiles(*bambou.FetchingInfo) (BGPProfilesList, *bambou.Error)
}

// BGPProfilesParent is the interface that a parent of a BGPProfile must implement.
// A Parent is defined as an entity that has BGPProfile as a child.
// A Parent is an Ancestor which can create a BGPProfile.
type BGPProfilesParent interface {
	BGPProfilesAncestor
	CreateBGPProfile(*BGPProfile) *bambou.Error
}

// BGPProfile represents the model of a bgpprofile
type BGPProfile struct {
	ID                              string `json:"ID,omitempty"`
	ParentID                        string `json:"parentID,omitempty"`
	ParentType                      string `json:"parentType,omitempty"`
	Owner                           string `json:"owner,omitempty"`
	Name                            string `json:"name,omitempty"`
	DampeningHalfLife               int    `json:"dampeningHalfLife,omitempty"`
	DampeningMaxSuppress            int    `json:"dampeningMaxSuppress,omitempty"`
	DampeningName                   string `json:"dampeningName,omitempty"`
	DampeningReuse                  int    `json:"dampeningReuse,omitempty"`
	DampeningSuppress               int    `json:"dampeningSuppress,omitempty"`
	Description                     string `json:"description,omitempty"`
	EntityScope                     string `json:"entityScope,omitempty"`
	AssociatedExportRoutingPolicyID string `json:"associatedExportRoutingPolicyID,omitempty"`
	AssociatedImportRoutingPolicyID string `json:"associatedImportRoutingPolicyID,omitempty"`
	ExternalID                      string `json:"externalID,omitempty"`
}

// NewBGPProfile returns a new *BGPProfile
func NewBGPProfile() *BGPProfile {

	return &BGPProfile{
		DampeningHalfLife:    15,
		DampeningMaxSuppress: 60,
		DampeningReuse:       750,
		DampeningSuppress:    3000,
	}
}

// Identity returns the Identity of the object.
func (o *BGPProfile) Identity() bambou.Identity {

	return BGPProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *BGPProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *BGPProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the BGPProfile from the server
func (o *BGPProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the BGPProfile into the server
func (o *BGPProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the BGPProfile from the server
func (o *BGPProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the BGPProfile
func (o *BGPProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the BGPProfile
func (o *BGPProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the BGPProfile
func (o *BGPProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the BGPProfile
func (o *BGPProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
