/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// QosPolicerIdentity represents the Identity of the object
var QosPolicerIdentity = bambou.Identity{
	Name:     "qospolicer",
	Category: "qospolicers",
}

// QosPolicersList represents a list of QosPolicers
type QosPolicersList []*QosPolicer

// QosPolicersAncestor is the interface that an ancestor of a QosPolicer must implement.
// An Ancestor is defined as an entity that has QosPolicer as a descendant.
// An Ancestor can get a list of its child QosPolicers, but not necessarily create one.
type QosPolicersAncestor interface {
	QosPolicers(*bambou.FetchingInfo) (QosPolicersList, *bambou.Error)
}

// QosPolicersParent is the interface that a parent of a QosPolicer must implement.
// A Parent is defined as an entity that has QosPolicer as a child.
// A Parent is an Ancestor which can create a QosPolicer.
type QosPolicersParent interface {
	QosPolicersAncestor
	CreateQosPolicer(*QosPolicer) *bambou.Error
}

// QosPolicer represents the model of a qospolicer
type QosPolicer struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Rate          int    `json:"rate,omitempty"`
	Description   string `json:"description,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	Burst         int    `json:"burst,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewQosPolicer returns a new *QosPolicer
func NewQosPolicer() *QosPolicer {

	return &QosPolicer{
		Rate:  1,
		Burst: 1,
	}
}

// Identity returns the Identity of the object.
func (o *QosPolicer) Identity() bambou.Identity {

	return QosPolicerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *QosPolicer) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *QosPolicer) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the QosPolicer from the server
func (o *QosPolicer) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the QosPolicer into the server
func (o *QosPolicer) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the QosPolicer from the server
func (o *QosPolicer) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the QosPolicer
func (o *QosPolicer) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the QosPolicer
func (o *QosPolicer) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the QosPolicer
func (o *QosPolicer) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the QosPolicer
func (o *QosPolicer) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
