/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// AggregateMetadataIdentity represents the Identity of the object
var AggregateMetadataIdentity = bambou.Identity{
	Name:     "aggregatemetadata",
	Category: "aggregatemetadatas",
}

// AggregateMetadatasList represents a list of AggregateMetadatas
type AggregateMetadatasList []*AggregateMetadata

// AggregateMetadatasAncestor is the interface that an ancestor of a AggregateMetadata must implement.
// An Ancestor is defined as an entity that has AggregateMetadata as a descendant.
// An Ancestor can get a list of its child AggregateMetadatas, but not necessarily create one.
type AggregateMetadatasAncestor interface {
	AggregateMetadatas(*bambou.FetchingInfo) (AggregateMetadatasList, *bambou.Error)
}

// AggregateMetadatasParent is the interface that a parent of a AggregateMetadata must implement.
// A Parent is defined as an entity that has AggregateMetadata as a child.
// A Parent is an Ancestor which can create a AggregateMetadata.
type AggregateMetadatasParent interface {
	AggregateMetadatasAncestor
	CreateAggregateMetadata(*AggregateMetadata) *bambou.Error
}

// AggregateMetadata represents the model of a aggregatemetadata
type AggregateMetadata struct {
	ID                          string        `json:"ID,omitempty"`
	ParentID                    string        `json:"parentID,omitempty"`
	ParentType                  string        `json:"parentType,omitempty"`
	Owner                       string        `json:"owner,omitempty"`
	Name                        string        `json:"name,omitempty"`
	Description                 string        `json:"description,omitempty"`
	MetadataTagIDs              []interface{} `json:"metadataTagIDs,omitempty"`
	NetworkNotificationDisabled bool          `json:"networkNotificationDisabled"`
	Blob                        string        `json:"blob,omitempty"`
	Global                      bool          `json:"global"`
	EntityScope                 string        `json:"entityScope,omitempty"`
	ExternalID                  string        `json:"externalID,omitempty"`
}

// NewAggregateMetadata returns a new *AggregateMetadata
func NewAggregateMetadata() *AggregateMetadata {

	return &AggregateMetadata{}
}

// Identity returns the Identity of the object.
func (o *AggregateMetadata) Identity() bambou.Identity {

	return AggregateMetadataIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AggregateMetadata) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AggregateMetadata) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the AggregateMetadata from the server
func (o *AggregateMetadata) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the AggregateMetadata into the server
func (o *AggregateMetadata) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the AggregateMetadata from the server
func (o *AggregateMetadata) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
