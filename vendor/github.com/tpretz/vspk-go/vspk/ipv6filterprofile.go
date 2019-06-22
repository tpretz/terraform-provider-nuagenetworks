/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IPv6FilterProfileIdentity represents the Identity of the object
var IPv6FilterProfileIdentity = bambou.Identity{
	Name:     "ipv6filterprofile",
	Category: "ipv6filterprofiles",
}

// IPv6FilterProfilesList represents a list of IPv6FilterProfiles
type IPv6FilterProfilesList []*IPv6FilterProfile

// IPv6FilterProfilesAncestor is the interface that an ancestor of a IPv6FilterProfile must implement.
// An Ancestor is defined as an entity that has IPv6FilterProfile as a descendant.
// An Ancestor can get a list of its child IPv6FilterProfiles, but not necessarily create one.
type IPv6FilterProfilesAncestor interface {
	IPv6FilterProfiles(*bambou.FetchingInfo) (IPv6FilterProfilesList, *bambou.Error)
}

// IPv6FilterProfilesParent is the interface that a parent of a IPv6FilterProfile must implement.
// A Parent is defined as an entity that has IPv6FilterProfile as a child.
// A Parent is an Ancestor which can create a IPv6FilterProfile.
type IPv6FilterProfilesParent interface {
	IPv6FilterProfilesAncestor
	CreateIPv6FilterProfile(*IPv6FilterProfile) *bambou.Error
}

// IPv6FilterProfile represents the model of a ipv6filterprofile
type IPv6FilterProfile struct {
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

// NewIPv6FilterProfile returns a new *IPv6FilterProfile
func NewIPv6FilterProfile() *IPv6FilterProfile {

	return &IPv6FilterProfile{}
}

// Identity returns the Identity of the object.
func (o *IPv6FilterProfile) Identity() bambou.Identity {

	return IPv6FilterProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IPv6FilterProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IPv6FilterProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IPv6FilterProfile from the server
func (o *IPv6FilterProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IPv6FilterProfile into the server
func (o *IPv6FilterProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IPv6FilterProfile from the server
func (o *IPv6FilterProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IPv6FilterProfile
func (o *IPv6FilterProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IPv6FilterProfile
func (o *IPv6FilterProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IPv6FilterProfile
func (o *IPv6FilterProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IPv6FilterProfile
func (o *IPv6FilterProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
