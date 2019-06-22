/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// ZFBAutoAssignmentIdentity represents the Identity of the object
var ZFBAutoAssignmentIdentity = bambou.Identity{
	Name:     "zfbautoassignment",
	Category: "zfbautoassignments",
}

// ZFBAutoAssignmentsList represents a list of ZFBAutoAssignments
type ZFBAutoAssignmentsList []*ZFBAutoAssignment

// ZFBAutoAssignmentsAncestor is the interface that an ancestor of a ZFBAutoAssignment must implement.
// An Ancestor is defined as an entity that has ZFBAutoAssignment as a descendant.
// An Ancestor can get a list of its child ZFBAutoAssignments, but not necessarily create one.
type ZFBAutoAssignmentsAncestor interface {
	ZFBAutoAssignments(*bambou.FetchingInfo) (ZFBAutoAssignmentsList, *bambou.Error)
}

// ZFBAutoAssignmentsParent is the interface that a parent of a ZFBAutoAssignment must implement.
// A Parent is defined as an entity that has ZFBAutoAssignment as a child.
// A Parent is an Ancestor which can create a ZFBAutoAssignment.
type ZFBAutoAssignmentsParent interface {
	ZFBAutoAssignmentsAncestor
	CreateZFBAutoAssignment(*ZFBAutoAssignment) *bambou.Error
}

// ZFBAutoAssignment represents the model of a zfbautoassignment
type ZFBAutoAssignment struct {
	ID                       string        `json:"ID,omitempty"`
	ParentID                 string        `json:"parentID,omitempty"`
	ParentType               string        `json:"parentType,omitempty"`
	Owner                    string        `json:"owner,omitempty"`
	ZFBMatchAttribute        string        `json:"ZFBMatchAttribute,omitempty"`
	ZFBMatchAttributeValues  []interface{} `json:"ZFBMatchAttributeValues,omitempty"`
	Name                     string        `json:"name,omitempty"`
	LastUpdatedBy            string        `json:"lastUpdatedBy,omitempty"`
	Description              string        `json:"description,omitempty"`
	EntityScope              string        `json:"entityScope,omitempty"`
	Priority                 int           `json:"priority,omitempty"`
	AssociatedEnterpriseID   string        `json:"associatedEnterpriseID,omitempty"`
	AssociatedEnterpriseName string        `json:"associatedEnterpriseName,omitempty"`
	ExternalID               string        `json:"externalID,omitempty"`
}

// NewZFBAutoAssignment returns a new *ZFBAutoAssignment
func NewZFBAutoAssignment() *ZFBAutoAssignment {

	return &ZFBAutoAssignment{}
}

// Identity returns the Identity of the object.
func (o *ZFBAutoAssignment) Identity() bambou.Identity {

	return ZFBAutoAssignmentIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ZFBAutoAssignment) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ZFBAutoAssignment) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the ZFBAutoAssignment from the server
func (o *ZFBAutoAssignment) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ZFBAutoAssignment into the server
func (o *ZFBAutoAssignment) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ZFBAutoAssignment from the server
func (o *ZFBAutoAssignment) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
