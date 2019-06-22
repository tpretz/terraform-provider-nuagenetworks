/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// NSGPatchProfileIdentity represents the Identity of the object
var NSGPatchProfileIdentity = bambou.Identity{
	Name:     "nsgpatchprofile",
	Category: "nsgpatchprofiles",
}

// NSGPatchProfilesList represents a list of NSGPatchProfiles
type NSGPatchProfilesList []*NSGPatchProfile

// NSGPatchProfilesAncestor is the interface that an ancestor of a NSGPatchProfile must implement.
// An Ancestor is defined as an entity that has NSGPatchProfile as a descendant.
// An Ancestor can get a list of its child NSGPatchProfiles, but not necessarily create one.
type NSGPatchProfilesAncestor interface {
	NSGPatchProfiles(*bambou.FetchingInfo) (NSGPatchProfilesList, *bambou.Error)
}

// NSGPatchProfilesParent is the interface that a parent of a NSGPatchProfile must implement.
// A Parent is defined as an entity that has NSGPatchProfile as a child.
// A Parent is an Ancestor which can create a NSGPatchProfile.
type NSGPatchProfilesParent interface {
	NSGPatchProfilesAncestor
	CreateNSGPatchProfile(*NSGPatchProfile) *bambou.Error
}

// NSGPatchProfile represents the model of a nsgpatchprofile
type NSGPatchProfile struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	PatchTag      string `json:"patchTag,omitempty"`
	PatchURL      string `json:"patchURL,omitempty"`
	Description   string `json:"description,omitempty"`
	EnterpriseID  string `json:"enterpriseID,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewNSGPatchProfile returns a new *NSGPatchProfile
func NewNSGPatchProfile() *NSGPatchProfile {

	return &NSGPatchProfile{}
}

// Identity returns the Identity of the object.
func (o *NSGPatchProfile) Identity() bambou.Identity {

	return NSGPatchProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSGPatchProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSGPatchProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSGPatchProfile from the server
func (o *NSGPatchProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSGPatchProfile into the server
func (o *NSGPatchProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSGPatchProfile from the server
func (o *NSGPatchProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
