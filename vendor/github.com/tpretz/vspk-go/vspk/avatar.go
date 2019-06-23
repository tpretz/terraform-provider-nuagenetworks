/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// AvatarIdentity represents the Identity of the object
var AvatarIdentity = bambou.Identity{
	Name:     "avatar",
	Category: "avatars",
}

// AvatarsList represents a list of Avatars
type AvatarsList []*Avatar

// AvatarsAncestor is the interface that an ancestor of a Avatar must implement.
// An Ancestor is defined as an entity that has Avatar as a descendant.
// An Ancestor can get a list of its child Avatars, but not necessarily create one.
type AvatarsAncestor interface {
	Avatars(*bambou.FetchingInfo) (AvatarsList, *bambou.Error)
}

// AvatarsParent is the interface that a parent of a Avatar must implement.
// A Parent is defined as an entity that has Avatar as a child.
// A Parent is an Ancestor which can create a Avatar.
type AvatarsParent interface {
	AvatarsAncestor
	CreateAvatar(*Avatar) *bambou.Error
}

// Avatar represents the model of a avatar
type Avatar struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
	Type          string `json:"type,omitempty"`
}

// NewAvatar returns a new *Avatar
func NewAvatar() *Avatar {

	return &Avatar{}
}

// Identity returns the Identity of the object.
func (o *Avatar) Identity() bambou.Identity {

	return AvatarIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Avatar) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Avatar) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Avatar from the server
func (o *Avatar) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Avatar into the server
func (o *Avatar) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Avatar from the server
func (o *Avatar) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Avatar
func (o *Avatar) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Avatar
func (o *Avatar) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Avatar
func (o *Avatar) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Avatar
func (o *Avatar) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
