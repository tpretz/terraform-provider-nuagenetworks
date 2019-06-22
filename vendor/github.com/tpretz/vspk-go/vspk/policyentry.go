/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// PolicyEntryIdentity represents the Identity of the object
var PolicyEntryIdentity = bambou.Identity{
	Name:     "policyentry",
	Category: "policyentries",
}

// PolicyEntriesList represents a list of PolicyEntries
type PolicyEntriesList []*PolicyEntry

// PolicyEntriesAncestor is the interface that an ancestor of a PolicyEntry must implement.
// An Ancestor is defined as an entity that has PolicyEntry as a descendant.
// An Ancestor can get a list of its child PolicyEntries, but not necessarily create one.
type PolicyEntriesAncestor interface {
	PolicyEntries(*bambou.FetchingInfo) (PolicyEntriesList, *bambou.Error)
}

// PolicyEntriesParent is the interface that a parent of a PolicyEntry must implement.
// A Parent is defined as an entity that has PolicyEntry as a child.
// A Parent is an Ancestor which can create a PolicyEntry.
type PolicyEntriesParent interface {
	PolicyEntriesAncestor
	CreatePolicyEntry(*PolicyEntry) *bambou.Error
}

// PolicyEntry represents the model of a policyentry
type PolicyEntry struct {
	ID            string      `json:"ID,omitempty"`
	ParentID      string      `json:"parentID,omitempty"`
	ParentType    string      `json:"parentType,omitempty"`
	Owner         string      `json:"owner,omitempty"`
	Name          string      `json:"name,omitempty"`
	LastUpdatedBy string      `json:"lastUpdatedBy,omitempty"`
	MatchCriteria interface{} `json:"matchCriteria,omitempty"`
	Actions       interface{} `json:"actions,omitempty"`
	Description   string      `json:"description,omitempty"`
	EntityScope   string      `json:"entityScope,omitempty"`
	ExternalID    string      `json:"externalID,omitempty"`
}

// NewPolicyEntry returns a new *PolicyEntry
func NewPolicyEntry() *PolicyEntry {

	return &PolicyEntry{}
}

// Identity returns the Identity of the object.
func (o *PolicyEntry) Identity() bambou.Identity {

	return PolicyEntryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyEntry) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyEntry) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PolicyEntry from the server
func (o *PolicyEntry) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PolicyEntry into the server
func (o *PolicyEntry) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PolicyEntry from the server
func (o *PolicyEntry) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the PolicyEntry
func (o *PolicyEntry) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the PolicyEntry
func (o *PolicyEntry) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the PolicyEntry
func (o *PolicyEntry) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the PolicyEntry
func (o *PolicyEntry) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
