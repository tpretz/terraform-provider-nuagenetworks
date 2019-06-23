/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// PATIPEntryIdentity represents the Identity of the object
var PATIPEntryIdentity = bambou.Identity{
	Name:     "patipentry",
	Category: "patipentries",
}

// PATIPEntriesList represents a list of PATIPEntries
type PATIPEntriesList []*PATIPEntry

// PATIPEntriesAncestor is the interface that an ancestor of a PATIPEntry must implement.
// An Ancestor is defined as an entity that has PATIPEntry as a descendant.
// An Ancestor can get a list of its child PATIPEntries, but not necessarily create one.
type PATIPEntriesAncestor interface {
	PATIPEntries(*bambou.FetchingInfo) (PATIPEntriesList, *bambou.Error)
}

// PATIPEntriesParent is the interface that a parent of a PATIPEntry must implement.
// A Parent is defined as an entity that has PATIPEntry as a child.
// A Parent is an Ancestor which can create a PATIPEntry.
type PATIPEntriesParent interface {
	PATIPEntriesAncestor
	CreatePATIPEntry(*PATIPEntry) *bambou.Error
}

// PATIPEntry represents the model of a patipentry
type PATIPEntry struct {
	ID                 string `json:"ID,omitempty"`
	ParentID           string `json:"parentID,omitempty"`
	ParentType         string `json:"parentType,omitempty"`
	Owner              string `json:"owner,omitempty"`
	PATCentralized     bool   `json:"PATCentralized"`
	IPAddress          string `json:"IPAddress,omitempty"`
	IPType             string `json:"IPType,omitempty"`
	LastUpdatedBy      string `json:"lastUpdatedBy,omitempty"`
	EntityScope        string `json:"entityScope,omitempty"`
	AssociatedDomainID string `json:"associatedDomainID,omitempty"`
	ExternalID         string `json:"externalID,omitempty"`
	HypervisorID       string `json:"hypervisorID,omitempty"`
}

// NewPATIPEntry returns a new *PATIPEntry
func NewPATIPEntry() *PATIPEntry {

	return &PATIPEntry{
		IPType: "IPV4",
	}
}

// Identity returns the Identity of the object.
func (o *PATIPEntry) Identity() bambou.Identity {

	return PATIPEntryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PATIPEntry) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PATIPEntry) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the PATIPEntry from the server
func (o *PATIPEntry) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the PATIPEntry into the server
func (o *PATIPEntry) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the PATIPEntry from the server
func (o *PATIPEntry) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
