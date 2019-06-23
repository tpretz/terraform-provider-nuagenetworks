/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// PublicNetworkMacroIdentity represents the Identity of the object
var PublicNetworkMacroIdentity = bambou.Identity{
	Name:     "publicnetwork",
	Category: "publicnetworks",
}

// PublicNetworkMacrosList represents a list of PublicNetworkMacros
type PublicNetworkMacrosList []*PublicNetworkMacro

// PublicNetworkMacrosAncestor is the interface that an ancestor of a PublicNetworkMacro must implement.
// An Ancestor is defined as an entity that has PublicNetworkMacro as a descendant.
// An Ancestor can get a list of its child PublicNetworkMacros, but not necessarily create one.
type PublicNetworkMacrosAncestor interface {
	PublicNetworkMacros(*bambou.FetchingInfo) (PublicNetworkMacrosList, *bambou.Error)
}

// PublicNetworkMacrosParent is the interface that a parent of a PublicNetworkMacro must implement.
// A Parent is defined as an entity that has PublicNetworkMacro as a child.
// A Parent is an Ancestor which can create a PublicNetworkMacro.
type PublicNetworkMacrosParent interface {
	PublicNetworkMacrosAncestor
	CreatePublicNetworkMacro(*PublicNetworkMacro) *bambou.Error
}

// PublicNetworkMacro represents the model of a publicnetwork
type PublicNetworkMacro struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	IPType        string `json:"IPType,omitempty"`
	IPv6Address   string `json:"IPv6Address,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Address       string `json:"address,omitempty"`
	Netmask       string `json:"netmask,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewPublicNetworkMacro returns a new *PublicNetworkMacro
func NewPublicNetworkMacro() *PublicNetworkMacro {

	return &PublicNetworkMacro{}
}

// Identity returns the Identity of the object.
func (o *PublicNetworkMacro) Identity() bambou.Identity {

	return PublicNetworkMacroIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PublicNetworkMacro) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PublicNetworkMacro) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PublicNetworkMacro from the server
func (o *PublicNetworkMacro) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PublicNetworkMacro into the server
func (o *PublicNetworkMacro) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PublicNetworkMacro from the server
func (o *PublicNetworkMacro) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the PublicNetworkMacro
func (o *PublicNetworkMacro) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the PublicNetworkMacro
func (o *PublicNetworkMacro) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the PublicNetworkMacro
func (o *PublicNetworkMacro) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the PublicNetworkMacro
func (o *PublicNetworkMacro) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the PublicNetworkMacro
func (o *PublicNetworkMacro) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
