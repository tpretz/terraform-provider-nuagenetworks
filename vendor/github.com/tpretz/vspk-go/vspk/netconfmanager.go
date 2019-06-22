/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NetconfManagerIdentity represents the Identity of the object
var NetconfManagerIdentity = bambou.Identity{
	Name:     "netconfmanager",
	Category: "netconfmanagers",
}

// NetconfManagersList represents a list of NetconfManagers
type NetconfManagersList []*NetconfManager

// NetconfManagersAncestor is the interface that an ancestor of a NetconfManager must implement.
// An Ancestor is defined as an entity that has NetconfManager as a descendant.
// An Ancestor can get a list of its child NetconfManagers, but not necessarily create one.
type NetconfManagersAncestor interface {
	NetconfManagers(*bambou.FetchingInfo) (NetconfManagersList, *bambou.Error)
}

// NetconfManagersParent is the interface that a parent of a NetconfManager must implement.
// A Parent is defined as an entity that has NetconfManager as a child.
// A Parent is an Ancestor which can create a NetconfManager.
type NetconfManagersParent interface {
	NetconfManagersAncestor
	CreateNetconfManager(*NetconfManager) *bambou.Error
}

// NetconfManager represents the model of a netconfmanager
type NetconfManager struct {
	ID              string `json:"ID,omitempty"`
	ParentID        string `json:"parentID,omitempty"`
	ParentType      string `json:"parentType,omitempty"`
	Owner           string `json:"owner,omitempty"`
	Name            string `json:"name,omitempty"`
	LastUpdatedBy   string `json:"lastUpdatedBy,omitempty"`
	Release         string `json:"release,omitempty"`
	EntityScope     string `json:"entityScope,omitempty"`
	AssocEntityType string `json:"assocEntityType,omitempty"`
	Status          string `json:"status,omitempty"`
	ExternalID      string `json:"externalID,omitempty"`
}

// NewNetconfManager returns a new *NetconfManager
func NewNetconfManager() *NetconfManager {

	return &NetconfManager{}
}

// Identity returns the Identity of the object.
func (o *NetconfManager) Identity() bambou.Identity {

	return NetconfManagerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NetconfManager) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NetconfManager) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NetconfManager from the server
func (o *NetconfManager) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NetconfManager into the server
func (o *NetconfManager) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NetconfManager from the server
func (o *NetconfManager) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the NetconfManager
func (o *NetconfManager) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the NetconfManager
func (o *NetconfManager) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// NetconfSessions retrieves the list of child NetconfSessions of the NetconfManager
func (o *NetconfManager) NetconfSessions(info *bambou.FetchingInfo) (NetconfSessionsList, *bambou.Error) {

	var list NetconfSessionsList
	err := bambou.CurrentSession().FetchChildren(o, NetconfSessionIdentity, &list, info)
	return list, err
}

// CreateNetconfSession creates a new child NetconfSession under the NetconfManager
func (o *NetconfManager) CreateNetconfSession(child *NetconfSession) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Alarms retrieves the list of child Alarms of the NetconfManager
func (o *NetconfManager) Alarms(info *bambou.FetchingInfo) (AlarmsList, *bambou.Error) {

	var list AlarmsList
	err := bambou.CurrentSession().FetchChildren(o, AlarmIdentity, &list, info)
	return list, err
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the NetconfManager
func (o *NetconfManager) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the NetconfManager
func (o *NetconfManager) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
