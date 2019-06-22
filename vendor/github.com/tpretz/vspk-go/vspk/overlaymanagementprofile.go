/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// OverlayManagementProfileIdentity represents the Identity of the object
var OverlayManagementProfileIdentity = bambou.Identity{
	Name:     "overlaymanagementprofile",
	Category: "overlaymanagementprofiles",
}

// OverlayManagementProfilesList represents a list of OverlayManagementProfiles
type OverlayManagementProfilesList []*OverlayManagementProfile

// OverlayManagementProfilesAncestor is the interface that an ancestor of a OverlayManagementProfile must implement.
// An Ancestor is defined as an entity that has OverlayManagementProfile as a descendant.
// An Ancestor can get a list of its child OverlayManagementProfiles, but not necessarily create one.
type OverlayManagementProfilesAncestor interface {
	OverlayManagementProfiles(*bambou.FetchingInfo) (OverlayManagementProfilesList, *bambou.Error)
}

// OverlayManagementProfilesParent is the interface that a parent of a OverlayManagementProfile must implement.
// A Parent is defined as an entity that has OverlayManagementProfile as a child.
// A Parent is an Ancestor which can create a OverlayManagementProfile.
type OverlayManagementProfilesParent interface {
	OverlayManagementProfilesAncestor
	CreateOverlayManagementProfile(*OverlayManagementProfile) *bambou.Error
}

// OverlayManagementProfile represents the model of a overlaymanagementprofile
type OverlayManagementProfile struct {
	ID          string `json:"ID,omitempty"`
	ParentID    string `json:"parentID,omitempty"`
	ParentType  string `json:"parentType,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// NewOverlayManagementProfile returns a new *OverlayManagementProfile
func NewOverlayManagementProfile() *OverlayManagementProfile {

	return &OverlayManagementProfile{}
}

// Identity returns the Identity of the object.
func (o *OverlayManagementProfile) Identity() bambou.Identity {

	return OverlayManagementProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *OverlayManagementProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *OverlayManagementProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the OverlayManagementProfile from the server
func (o *OverlayManagementProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the OverlayManagementProfile into the server
func (o *OverlayManagementProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the OverlayManagementProfile from the server
func (o *OverlayManagementProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// OverlayManagementSubnetProfiles retrieves the list of child OverlayManagementSubnetProfiles of the OverlayManagementProfile
func (o *OverlayManagementProfile) OverlayManagementSubnetProfiles(info *bambou.FetchingInfo) (OverlayManagementSubnetProfilesList, *bambou.Error) {

	var list OverlayManagementSubnetProfilesList
	err := bambou.CurrentSession().FetchChildren(o, OverlayManagementSubnetProfileIdentity, &list, info)
	return list, err
}

// CreateOverlayManagementSubnetProfile creates a new child OverlayManagementSubnetProfile under the OverlayManagementProfile
func (o *OverlayManagementProfile) CreateOverlayManagementSubnetProfile(child *OverlayManagementSubnetProfile) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
