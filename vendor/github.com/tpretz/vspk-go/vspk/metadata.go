/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// MetadataIdentity represents the Identity of the object
var MetadataIdentity = bambou.Identity{
	Name:     "metadata",
	Category: "metadatas",
}

// MetadatasList represents a list of Metadatas
type MetadatasList []*Metadata

// MetadatasAncestor is the interface that an ancestor of a Metadata must implement.
// An Ancestor is defined as an entity that has Metadata as a descendant.
// An Ancestor can get a list of its child Metadatas, but not necessarily create one.
type MetadatasAncestor interface {
	Metadatas(*bambou.FetchingInfo) (MetadatasList, *bambou.Error)
}

// MetadatasParent is the interface that a parent of a Metadata must implement.
// A Parent is defined as an entity that has Metadata as a child.
// A Parent is an Ancestor which can create a Metadata.
type MetadatasParent interface {
	MetadatasAncestor
	CreateMetadata(*Metadata) *bambou.Error
}

// Metadata represents the model of a metadata
type Metadata struct {
	ID                          string        `json:"ID,omitempty"`
	ParentID                    string        `json:"parentID,omitempty"`
	ParentType                  string        `json:"parentType,omitempty"`
	Owner                       string        `json:"owner,omitempty"`
	Name                        string        `json:"name,omitempty"`
	LastUpdatedBy               string        `json:"lastUpdatedBy,omitempty"`
	Description                 string        `json:"description,omitempty"`
	MetadataTagIDs              []interface{} `json:"metadataTagIDs,omitempty"`
	NetworkNotificationDisabled bool          `json:"networkNotificationDisabled"`
	Blob                        string        `json:"blob,omitempty"`
	Global                      bool          `json:"global"`
	EntityScope                 string        `json:"entityScope,omitempty"`
	AssocEntityID               string        `json:"assocEntityID,omitempty"`
	AssocEntityType             string        `json:"assocEntityType,omitempty"`
	ExternalID                  string        `json:"externalID,omitempty"`
}

// NewMetadata returns a new *Metadata
func NewMetadata() *Metadata {

	return &Metadata{}
}

// Identity returns the Identity of the object.
func (o *Metadata) Identity() bambou.Identity {

	return MetadataIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Metadata) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Metadata) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Metadata from the server
func (o *Metadata) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Metadata into the server
func (o *Metadata) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Metadata from the server
func (o *Metadata) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// EventLogs retrieves the list of child EventLogs of the Metadata
func (o *Metadata) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
