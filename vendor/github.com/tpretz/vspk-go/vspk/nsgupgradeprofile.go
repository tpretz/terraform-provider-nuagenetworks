/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// NSGUpgradeProfileIdentity represents the Identity of the object
var NSGUpgradeProfileIdentity = bambou.Identity{
	Name:     "nsgupgradeprofile",
	Category: "nsgupgradeprofiles",
}

// NSGUpgradeProfilesList represents a list of NSGUpgradeProfiles
type NSGUpgradeProfilesList []*NSGUpgradeProfile

// NSGUpgradeProfilesAncestor is the interface that an ancestor of a NSGUpgradeProfile must implement.
// An Ancestor is defined as an entity that has NSGUpgradeProfile as a descendant.
// An Ancestor can get a list of its child NSGUpgradeProfiles, but not necessarily create one.
type NSGUpgradeProfilesAncestor interface {
	NSGUpgradeProfiles(*bambou.FetchingInfo) (NSGUpgradeProfilesList, *bambou.Error)
}

// NSGUpgradeProfilesParent is the interface that a parent of a NSGUpgradeProfile must implement.
// A Parent is defined as an entity that has NSGUpgradeProfile as a child.
// A Parent is an Ancestor which can create a NSGUpgradeProfile.
type NSGUpgradeProfilesParent interface {
	NSGUpgradeProfilesAncestor
	CreateNSGUpgradeProfile(*NSGUpgradeProfile) *bambou.Error
}

// NSGUpgradeProfile represents the model of a nsgupgradeprofile
type NSGUpgradeProfile struct {
	ID                  string `json:"ID,omitempty"`
	ParentID            string `json:"parentID,omitempty"`
	ParentType          string `json:"parentType,omitempty"`
	Owner               string `json:"owner,omitempty"`
	Name                string `json:"name,omitempty"`
	LastUpdatedBy       string `json:"lastUpdatedBy,omitempty"`
	Description         string `json:"description,omitempty"`
	MetadataUpgradePath string `json:"metadataUpgradePath,omitempty"`
	EnterpriseID        string `json:"enterpriseID,omitempty"`
	EntityScope         string `json:"entityScope,omitempty"`
	ExternalID          string `json:"externalID,omitempty"`
}

// NewNSGUpgradeProfile returns a new *NSGUpgradeProfile
func NewNSGUpgradeProfile() *NSGUpgradeProfile {

	return &NSGUpgradeProfile{}
}

// Identity returns the Identity of the object.
func (o *NSGUpgradeProfile) Identity() bambou.Identity {

	return NSGUpgradeProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NSGUpgradeProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NSGUpgradeProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NSGUpgradeProfile from the server
func (o *NSGUpgradeProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NSGUpgradeProfile into the server
func (o *NSGUpgradeProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NSGUpgradeProfile from the server
func (o *NSGUpgradeProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
