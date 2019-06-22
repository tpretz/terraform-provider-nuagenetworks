/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// LTEInformationIdentity represents the Identity of the object
var LTEInformationIdentity = bambou.Identity{
	Name:     "lteinformation",
	Category: "lteinformations",
}

// LTEInformationsList represents a list of LTEInformations
type LTEInformationsList []*LTEInformation

// LTEInformationsAncestor is the interface that an ancestor of a LTEInformation must implement.
// An Ancestor is defined as an entity that has LTEInformation as a descendant.
// An Ancestor can get a list of its child LTEInformations, but not necessarily create one.
type LTEInformationsAncestor interface {
	LTEInformations(*bambou.FetchingInfo) (LTEInformationsList, *bambou.Error)
}

// LTEInformationsParent is the interface that a parent of a LTEInformation must implement.
// A Parent is defined as an entity that has LTEInformation as a child.
// A Parent is an Ancestor which can create a LTEInformation.
type LTEInformationsParent interface {
	LTEInformationsAncestor
	CreateLTEInformation(*LTEInformation) *bambou.Error
}

// LTEInformation represents the model of a lteinformation
type LTEInformation struct {
	ID                string `json:"ID,omitempty"`
	ParentID          string `json:"parentID,omitempty"`
	ParentType        string `json:"parentType,omitempty"`
	Owner             string `json:"owner,omitempty"`
	LTEConnectionInfo string `json:"LTEConnectionInfo,omitempty"`
}

// NewLTEInformation returns a new *LTEInformation
func NewLTEInformation() *LTEInformation {

	return &LTEInformation{}
}

// Identity returns the Identity of the object.
func (o *LTEInformation) Identity() bambou.Identity {

	return LTEInformationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LTEInformation) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LTEInformation) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the LTEInformation from the server
func (o *LTEInformation) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the LTEInformation into the server
func (o *LTEInformation) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the LTEInformation from the server
func (o *LTEInformation) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
