/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// VSDComponentIdentity represents the Identity of the object
var VSDComponentIdentity = bambou.Identity{
	Name:     "component",
	Category: "components",
}

// VSDComponentsList represents a list of VSDComponents
type VSDComponentsList []*VSDComponent

// VSDComponentsAncestor is the interface that an ancestor of a VSDComponent must implement.
// An Ancestor is defined as an entity that has VSDComponent as a descendant.
// An Ancestor can get a list of its child VSDComponents, but not necessarily create one.
type VSDComponentsAncestor interface {
	VSDComponents(*bambou.FetchingInfo) (VSDComponentsList, *bambou.Error)
}

// VSDComponentsParent is the interface that a parent of a VSDComponent must implement.
// A Parent is defined as an entity that has VSDComponent as a child.
// A Parent is an Ancestor which can create a VSDComponent.
type VSDComponentsParent interface {
	VSDComponentsAncestor
	CreateVSDComponent(*VSDComponent) *bambou.Error
}

// VSDComponent represents the model of a component
type VSDComponent struct {
	ID             string `json:"ID,omitempty"`
	ParentID       string `json:"parentID,omitempty"`
	ParentType     string `json:"parentType,omitempty"`
	Owner          string `json:"owner,omitempty"`
	Name           string `json:"name,omitempty"`
	ManagementIP   string `json:"managementIP,omitempty"`
	Address        string `json:"address,omitempty"`
	Description    string `json:"description,omitempty"`
	EntityScope    string `json:"entityScope,omitempty"`
	Location       string `json:"location,omitempty"`
	ProductVersion string `json:"productVersion,omitempty"`
	Status         string `json:"status,omitempty"`
	ExternalID     string `json:"externalID,omitempty"`
	Type           string `json:"type,omitempty"`
}

// NewVSDComponent returns a new *VSDComponent
func NewVSDComponent() *VSDComponent {

	return &VSDComponent{}
}

// Identity returns the Identity of the object.
func (o *VSDComponent) Identity() bambou.Identity {

	return VSDComponentIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VSDComponent) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VSDComponent) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VSDComponent from the server
func (o *VSDComponent) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VSDComponent into the server
func (o *VSDComponent) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VSDComponent from the server
func (o *VSDComponent) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the VSDComponent
func (o *VSDComponent) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the VSDComponent
func (o *VSDComponent) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the VSDComponent
func (o *VSDComponent) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the VSDComponent
func (o *VSDComponent) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
