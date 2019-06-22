/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// GroupIdentity represents the Identity of the object
var GroupIdentity = bambou.Identity{
	Name:     "group",
	Category: "groups",
}

// GroupsList represents a list of Groups
type GroupsList []*Group

// GroupsAncestor is the interface that an ancestor of a Group must implement.
// An Ancestor is defined as an entity that has Group as a descendant.
// An Ancestor can get a list of its child Groups, but not necessarily create one.
type GroupsAncestor interface {
	Groups(*bambou.FetchingInfo) (GroupsList, *bambou.Error)
}

// GroupsParent is the interface that a parent of a Group must implement.
// A Parent is defined as an entity that has Group as a child.
// A Parent is an Ancestor which can create a Group.
type GroupsParent interface {
	GroupsAncestor
	CreateGroup(*Group) *bambou.Error
}

// Group represents the model of a group
type Group struct {
	ID                  string  `json:"ID,omitempty"`
	ParentID            string  `json:"parentID,omitempty"`
	ParentType          string  `json:"parentType,omitempty"`
	Owner               string  `json:"owner,omitempty"`
	LDAPGroupDN         string  `json:"LDAPGroupDN,omitempty"`
	Name                string  `json:"name,omitempty"`
	ManagementMode      string  `json:"managementMode,omitempty"`
	LastUpdatedBy       string  `json:"lastUpdatedBy,omitempty"`
	AccountRestrictions bool    `json:"accountRestrictions"`
	Description         string  `json:"description,omitempty"`
	RestrictionDate     float64 `json:"restrictionDate,omitempty"`
	EntityScope         string  `json:"entityScope,omitempty"`
	Role                string  `json:"role,omitempty"`
	Private             bool    `json:"private"`
	ExternalID          string  `json:"externalID,omitempty"`
}

// NewGroup returns a new *Group
func NewGroup() *Group {

	return &Group{}
}

// Identity returns the Identity of the object.
func (o *Group) Identity() bambou.Identity {

	return GroupIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Group) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Group) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Group from the server
func (o *Group) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Group into the server
func (o *Group) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Group from the server
func (o *Group) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the Group
func (o *Group) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Group
func (o *Group) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Group
func (o *Group) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Group
func (o *Group) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Users retrieves the list of child Users of the Group
func (o *Group) Users(info *bambou.FetchingInfo) (UsersList, *bambou.Error) {

	var list UsersList
	err := bambou.CurrentSession().FetchChildren(o, UserIdentity, &list, info)
	return list, err
}

// AssignUsers assigns the list of Users to the Group
func (o *Group) AssignUsers(children UsersList) *bambou.Error {

	list := []bambou.Identifiable{}
	for _, c := range children {
		list = append(list, c)
	}

	return bambou.CurrentSession().AssignChildren(o, list, UserIdentity)
}

// EventLogs retrieves the list of child EventLogs of the Group
func (o *Group) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
