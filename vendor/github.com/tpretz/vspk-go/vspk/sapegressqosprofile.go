/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// SAPEgressQoSProfileIdentity represents the Identity of the object
var SAPEgressQoSProfileIdentity = bambou.Identity{
	Name:     "sapegressqosprofile",
	Category: "sapegressqosprofiles",
}

// SAPEgressQoSProfilesList represents a list of SAPEgressQoSProfiles
type SAPEgressQoSProfilesList []*SAPEgressQoSProfile

// SAPEgressQoSProfilesAncestor is the interface that an ancestor of a SAPEgressQoSProfile must implement.
// An Ancestor is defined as an entity that has SAPEgressQoSProfile as a descendant.
// An Ancestor can get a list of its child SAPEgressQoSProfiles, but not necessarily create one.
type SAPEgressQoSProfilesAncestor interface {
	SAPEgressQoSProfiles(*bambou.FetchingInfo) (SAPEgressQoSProfilesList, *bambou.Error)
}

// SAPEgressQoSProfilesParent is the interface that a parent of a SAPEgressQoSProfile must implement.
// A Parent is defined as an entity that has SAPEgressQoSProfile as a child.
// A Parent is an Ancestor which can create a SAPEgressQoSProfile.
type SAPEgressQoSProfilesParent interface {
	SAPEgressQoSProfilesAncestor
	CreateSAPEgressQoSProfile(*SAPEgressQoSProfile) *bambou.Error
}

// SAPEgressQoSProfile represents the model of a sapegressqosprofile
type SAPEgressQoSProfile struct {
	ID          string `json:"ID,omitempty"`
	ParentID    string `json:"parentID,omitempty"`
	ParentType  string `json:"parentType,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// NewSAPEgressQoSProfile returns a new *SAPEgressQoSProfile
func NewSAPEgressQoSProfile() *SAPEgressQoSProfile {

	return &SAPEgressQoSProfile{}
}

// Identity returns the Identity of the object.
func (o *SAPEgressQoSProfile) Identity() bambou.Identity {

	return SAPEgressQoSProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SAPEgressQoSProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SAPEgressQoSProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the SAPEgressQoSProfile from the server
func (o *SAPEgressQoSProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the SAPEgressQoSProfile into the server
func (o *SAPEgressQoSProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the SAPEgressQoSProfile from the server
func (o *SAPEgressQoSProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}