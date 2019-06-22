/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IPFilterProfileIdentity represents the Identity of the object
var IPFilterProfileIdentity = bambou.Identity{
	Name:     "ipfilterprofile",
	Category: "ipfilterprofiles",
}

// IPFilterProfilesList represents a list of IPFilterProfiles
type IPFilterProfilesList []*IPFilterProfile

// IPFilterProfilesAncestor is the interface that an ancestor of a IPFilterProfile must implement.
// An Ancestor is defined as an entity that has IPFilterProfile as a descendant.
// An Ancestor can get a list of its child IPFilterProfiles, but not necessarily create one.
type IPFilterProfilesAncestor interface {
	IPFilterProfiles(*bambou.FetchingInfo) (IPFilterProfilesList, *bambou.Error)
}

// IPFilterProfilesParent is the interface that a parent of a IPFilterProfile must implement.
// A Parent is defined as an entity that has IPFilterProfile as a child.
// A Parent is an Ancestor which can create a IPFilterProfile.
type IPFilterProfilesParent interface {
	IPFilterProfilesAncestor
	CreateIPFilterProfile(*IPFilterProfile) *bambou.Error
}

// IPFilterProfile represents the model of a ipfilterprofile
type IPFilterProfile struct {
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

// NewIPFilterProfile returns a new *IPFilterProfile
func NewIPFilterProfile() *IPFilterProfile {

	return &IPFilterProfile{}
}

// Identity returns the Identity of the object.
func (o *IPFilterProfile) Identity() bambou.Identity {

	return IPFilterProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IPFilterProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IPFilterProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IPFilterProfile from the server
func (o *IPFilterProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IPFilterProfile into the server
func (o *IPFilterProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IPFilterProfile from the server
func (o *IPFilterProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the IPFilterProfile
func (o *IPFilterProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the IPFilterProfile
func (o *IPFilterProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the IPFilterProfile
func (o *IPFilterProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the IPFilterProfile
func (o *IPFilterProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
