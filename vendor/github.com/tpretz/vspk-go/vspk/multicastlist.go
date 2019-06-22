/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// MultiCastListIdentity represents the Identity of the object
var MultiCastListIdentity = bambou.Identity{
	Name:     "multicastlist",
	Category: "multicastlists",
}

// MultiCastListsList represents a list of MultiCastLists
type MultiCastListsList []*MultiCastList

// MultiCastListsAncestor is the interface that an ancestor of a MultiCastList must implement.
// An Ancestor is defined as an entity that has MultiCastList as a descendant.
// An Ancestor can get a list of its child MultiCastLists, but not necessarily create one.
type MultiCastListsAncestor interface {
	MultiCastLists(*bambou.FetchingInfo) (MultiCastListsList, *bambou.Error)
}

// MultiCastListsParent is the interface that a parent of a MultiCastList must implement.
// A Parent is defined as an entity that has MultiCastList as a child.
// A Parent is an Ancestor which can create a MultiCastList.
type MultiCastListsParent interface {
	MultiCastListsAncestor
	CreateMultiCastList(*MultiCastList) *bambou.Error
}

// MultiCastList represents the model of a multicastlist
type MultiCastList struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	McastType     string `json:"mcastType,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewMultiCastList returns a new *MultiCastList
func NewMultiCastList() *MultiCastList {

	return &MultiCastList{}
}

// Identity returns the Identity of the object.
func (o *MultiCastList) Identity() bambou.Identity {

	return MultiCastListIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MultiCastList) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MultiCastList) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the MultiCastList from the server
func (o *MultiCastList) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the MultiCastList into the server
func (o *MultiCastList) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the MultiCastList from the server
func (o *MultiCastList) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the MultiCastList
func (o *MultiCastList) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the MultiCastList
func (o *MultiCastList) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the MultiCastList
func (o *MultiCastList) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the MultiCastList
func (o *MultiCastList) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// MultiCastChannelMaps retrieves the list of child MultiCastChannelMaps of the MultiCastList
func (o *MultiCastList) MultiCastChannelMaps(info *bambou.FetchingInfo) (MultiCastChannelMapsList, *bambou.Error) {

	var list MultiCastChannelMapsList
	err := bambou.CurrentSession().FetchChildren(o, MultiCastChannelMapIdentity, &list, info)
	return list, err
}

// AssignMultiCastChannelMaps assigns the list of MultiCastChannelMaps to the MultiCastList
func (o *MultiCastList) AssignMultiCastChannelMaps(children MultiCastChannelMapsList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, MultiCastChannelMapIdentity)
}
