/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// ApplicationBindingIdentity represents the Identity of the object
var ApplicationBindingIdentity = bambou.Identity{
	Name:     "applicationbinding",
	Category: "applicationbindings",
}

// ApplicationBindingsList represents a list of ApplicationBindings
type ApplicationBindingsList []*ApplicationBinding

// ApplicationBindingsAncestor is the interface that an ancestor of a ApplicationBinding must implement.
// An Ancestor is defined as an entity that has ApplicationBinding as a descendant.
// An Ancestor can get a list of its child ApplicationBindings, but not necessarily create one.
type ApplicationBindingsAncestor interface {
	ApplicationBindings(*bambou.FetchingInfo) (ApplicationBindingsList, *bambou.Error)
}

// ApplicationBindingsParent is the interface that a parent of a ApplicationBinding must implement.
// A Parent is defined as an entity that has ApplicationBinding as a child.
// A Parent is an Ancestor which can create a ApplicationBinding.
type ApplicationBindingsParent interface {
	ApplicationBindingsAncestor
	CreateApplicationBinding(*ApplicationBinding) *bambou.Error
}

// ApplicationBinding represents the model of a applicationbinding
type ApplicationBinding struct {
	ID                      string `json:"ID,omitempty"`
	ParentID                string `json:"parentID,omitempty"`
	ParentType              string `json:"parentType,omitempty"`
	Owner                   string `json:"owner,omitempty"`
	LastUpdatedBy           string `json:"lastUpdatedBy,omitempty"`
	ReadOnly                bool   `json:"readOnly"`
	EntityScope             string `json:"entityScope,omitempty"`
	Priority                int    `json:"priority,omitempty"`
	AssociatedApplicationID string `json:"associatedApplicationID,omitempty"`
	ExternalID              string `json:"externalID,omitempty"`
}

// NewApplicationBinding returns a new *ApplicationBinding
func NewApplicationBinding() *ApplicationBinding {

	return &ApplicationBinding{
		ReadOnly: false,
	}
}

// Identity returns the Identity of the object.
func (o *ApplicationBinding) Identity() bambou.Identity {

	return ApplicationBindingIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ApplicationBinding) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ApplicationBinding) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the ApplicationBinding from the server
func (o *ApplicationBinding) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ApplicationBinding into the server
func (o *ApplicationBinding) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ApplicationBinding from the server
func (o *ApplicationBinding) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the ApplicationBinding
func (o *ApplicationBinding) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the ApplicationBinding
func (o *ApplicationBinding) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the ApplicationBinding
func (o *ApplicationBinding) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the ApplicationBinding
func (o *ApplicationBinding) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
