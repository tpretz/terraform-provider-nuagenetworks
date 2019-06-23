/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// EventLogIdentity represents the Identity of the object
var EventLogIdentity = bambou.Identity{
	Name:     "eventlog",
	Category: "eventlogs",
}

// EventLogsList represents a list of EventLogs
type EventLogsList []*EventLog

// EventLogsAncestor is the interface that an ancestor of a EventLog must implement.
// An Ancestor is defined as an entity that has EventLog as a descendant.
// An Ancestor can get a list of its child EventLogs, but not necessarily create one.
type EventLogsAncestor interface {
	EventLogs(*bambou.FetchingInfo) (EventLogsList, *bambou.Error)
}

// EventLogsParent is the interface that a parent of a EventLog must implement.
// A Parent is defined as an entity that has EventLog as a child.
// A Parent is an Ancestor which can create a EventLog.
type EventLogsParent interface {
	EventLogsAncestor
	CreateEventLog(*EventLog) *bambou.Error
}

// EventLog represents the model of a eventlog
type EventLog struct {
	ID                string        `json:"ID,omitempty"`
	ParentID          string        `json:"parentID,omitempty"`
	ParentType        string        `json:"parentType,omitempty"`
	Owner             string        `json:"owner,omitempty"`
	RequestID         string        `json:"requestID,omitempty"`
	Diff              interface{}   `json:"diff,omitempty"`
	Enterprise        string        `json:"enterprise,omitempty"`
	Entities          []interface{} `json:"entities,omitempty"`
	EntityID          string        `json:"entityID,omitempty"`
	EntityParentID    string        `json:"entityParentID,omitempty"`
	EntityParentType  string        `json:"entityParentType,omitempty"`
	EntityScope       string        `json:"entityScope,omitempty"`
	EntityType        string        `json:"entityType,omitempty"`
	User              string        `json:"user,omitempty"`
	EventReceivedTime float64       `json:"eventReceivedTime,omitempty"`
	ExternalID        string        `json:"externalID,omitempty"`
	Type              string        `json:"type,omitempty"`
}

// NewEventLog returns a new *EventLog
func NewEventLog() *EventLog {

	return &EventLog{}
}

// Identity returns the Identity of the object.
func (o *EventLog) Identity() bambou.Identity {

	return EventLogIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EventLog) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EventLog) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the EventLog from the server
func (o *EventLog) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the EventLog into the server
func (o *EventLog) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the EventLog from the server
func (o *EventLog) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the EventLog
func (o *EventLog) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the EventLog
func (o *EventLog) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the EventLog
func (o *EventLog) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the EventLog
func (o *EventLog) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
