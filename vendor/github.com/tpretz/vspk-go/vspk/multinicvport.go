/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// MultiNICVPortIdentity represents the Identity of the object
var MultiNICVPortIdentity = bambou.Identity{
	Name:     "multinicvport",
	Category: "multinicvports",
}

// MultiNICVPortsList represents a list of MultiNICVPorts
type MultiNICVPortsList []*MultiNICVPort

// MultiNICVPortsAncestor is the interface that an ancestor of a MultiNICVPort must implement.
// An Ancestor is defined as an entity that has MultiNICVPort as a descendant.
// An Ancestor can get a list of its child MultiNICVPorts, but not necessarily create one.
type MultiNICVPortsAncestor interface {
	MultiNICVPorts(*bambou.FetchingInfo) (MultiNICVPortsList, *bambou.Error)
}

// MultiNICVPortsParent is the interface that a parent of a MultiNICVPort must implement.
// A Parent is defined as an entity that has MultiNICVPort as a child.
// A Parent is an Ancestor which can create a MultiNICVPort.
type MultiNICVPortsParent interface {
	MultiNICVPortsAncestor
	CreateMultiNICVPort(*MultiNICVPort) *bambou.Error
}

// MultiNICVPort represents the model of a multinicvport
type MultiNICVPort struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewMultiNICVPort returns a new *MultiNICVPort
func NewMultiNICVPort() *MultiNICVPort {

	return &MultiNICVPort{}
}

// Identity returns the Identity of the object.
func (o *MultiNICVPort) Identity() bambou.Identity {

	return MultiNICVPortIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MultiNICVPort) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MultiNICVPort) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the MultiNICVPort from the server
func (o *MultiNICVPort) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the MultiNICVPort into the server
func (o *MultiNICVPort) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the MultiNICVPort from the server
func (o *MultiNICVPort) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the MultiNICVPort
func (o *MultiNICVPort) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the MultiNICVPort
func (o *MultiNICVPort) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the MultiNICVPort
func (o *MultiNICVPort) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the MultiNICVPort
func (o *MultiNICVPort) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// VPorts retrieves the list of child VPorts of the MultiNICVPort
func (o *MultiNICVPort) VPorts(info *bambou.FetchingInfo) (VPortsList, *bambou.Error) {

	var list VPortsList
	err := bambou.CurrentSession().FetchChildren(o, VPortIdentity, &list, info)
	return list, err
}
