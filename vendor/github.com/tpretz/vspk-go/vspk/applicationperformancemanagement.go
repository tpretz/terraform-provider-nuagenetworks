/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// ApplicationperformancemanagementIdentity represents the Identity of the object
var ApplicationperformancemanagementIdentity = bambou.Identity{
	Name:     "applicationperformancemanagement",
	Category: "applicationperformancemanagements",
}

// ApplicationperformancemanagementsList represents a list of Applicationperformancemanagements
type ApplicationperformancemanagementsList []*Applicationperformancemanagement

// ApplicationperformancemanagementsAncestor is the interface that an ancestor of a Applicationperformancemanagement must implement.
// An Ancestor is defined as an entity that has Applicationperformancemanagement as a descendant.
// An Ancestor can get a list of its child Applicationperformancemanagements, but not necessarily create one.
type ApplicationperformancemanagementsAncestor interface {
	Applicationperformancemanagements(*bambou.FetchingInfo) (ApplicationperformancemanagementsList, *bambou.Error)
}

// ApplicationperformancemanagementsParent is the interface that a parent of a Applicationperformancemanagement must implement.
// A Parent is defined as an entity that has Applicationperformancemanagement as a child.
// A Parent is an Ancestor which can create a Applicationperformancemanagement.
type ApplicationperformancemanagementsParent interface {
	ApplicationperformancemanagementsAncestor
	CreateApplicationperformancemanagement(*Applicationperformancemanagement) *bambou.Error
}

// Applicationperformancemanagement represents the model of a applicationperformancemanagement
type Applicationperformancemanagement struct {
	ID                             string `json:"ID,omitempty"`
	ParentID                       string `json:"parentID,omitempty"`
	ParentType                     string `json:"parentType,omitempty"`
	Owner                          string `json:"owner,omitempty"`
	Name                           string `json:"name,omitempty"`
	LastUpdatedBy                  string `json:"lastUpdatedBy,omitempty"`
	ReadOnly                       bool   `json:"readOnly"`
	Description                    string `json:"description,omitempty"`
	EntityScope                    string `json:"entityScope,omitempty"`
	AssociatedPerformanceMonitorID string `json:"associatedPerformanceMonitorID,omitempty"`
	ExternalID                     string `json:"externalID,omitempty"`
}

// NewApplicationperformancemanagement returns a new *Applicationperformancemanagement
func NewApplicationperformancemanagement() *Applicationperformancemanagement {

	return &Applicationperformancemanagement{
		ReadOnly: false,
	}
}

// Identity returns the Identity of the object.
func (o *Applicationperformancemanagement) Identity() bambou.Identity {

	return ApplicationperformancemanagementIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Applicationperformancemanagement) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Applicationperformancemanagement) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Applicationperformancemanagement from the server
func (o *Applicationperformancemanagement) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Applicationperformancemanagement into the server
func (o *Applicationperformancemanagement) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Applicationperformancemanagement from the server
func (o *Applicationperformancemanagement) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Applicationperformancemanagement
func (o *Applicationperformancemanagement) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Applicationperformancemanagement
func (o *Applicationperformancemanagement) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Applicationperformancemanagement
func (o *Applicationperformancemanagement) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Applicationperformancemanagement
func (o *Applicationperformancemanagement) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// ApplicationBindings retrieves the list of child ApplicationBindings of the Applicationperformancemanagement
func (o *Applicationperformancemanagement) ApplicationBindings(info *bambou.FetchingInfo) (ApplicationBindingsList, *bambou.Error) {

	var list ApplicationBindingsList
	err := bambou.CurrentSession().FetchChildren(o, ApplicationBindingIdentity, &list, info)
	return list, err
}

// CreateApplicationBinding creates a new child ApplicationBinding under the Applicationperformancemanagement
func (o *Applicationperformancemanagement) CreateApplicationBinding(child *ApplicationBinding) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
