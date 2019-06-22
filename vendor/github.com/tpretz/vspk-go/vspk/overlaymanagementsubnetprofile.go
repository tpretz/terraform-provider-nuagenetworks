/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// OverlayManagementSubnetProfileIdentity represents the Identity of the object
var OverlayManagementSubnetProfileIdentity = bambou.Identity{
	Name:     "overlaymanagementsubnetprofile",
	Category: "overlaymanagementsubnetprofiles",
}

// OverlayManagementSubnetProfilesList represents a list of OverlayManagementSubnetProfiles
type OverlayManagementSubnetProfilesList []*OverlayManagementSubnetProfile

// OverlayManagementSubnetProfilesAncestor is the interface that an ancestor of a OverlayManagementSubnetProfile must implement.
// An Ancestor is defined as an entity that has OverlayManagementSubnetProfile as a descendant.
// An Ancestor can get a list of its child OverlayManagementSubnetProfiles, but not necessarily create one.
type OverlayManagementSubnetProfilesAncestor interface {
	OverlayManagementSubnetProfiles(*bambou.FetchingInfo) (OverlayManagementSubnetProfilesList, *bambou.Error)
}

// OverlayManagementSubnetProfilesParent is the interface that a parent of a OverlayManagementSubnetProfile must implement.
// A Parent is defined as an entity that has OverlayManagementSubnetProfile as a child.
// A Parent is an Ancestor which can create a OverlayManagementSubnetProfile.
type OverlayManagementSubnetProfilesParent interface {
	OverlayManagementSubnetProfilesAncestor
	CreateOverlayManagementSubnetProfile(*OverlayManagementSubnetProfile) *bambou.Error
}

// OverlayManagementSubnetProfile represents the model of a overlaymanagementsubnetprofile
type OverlayManagementSubnetProfile struct {
	ID                    string        `json:"ID,omitempty"`
	ParentID              string        `json:"parentID,omitempty"`
	ParentType            string        `json:"parentType,omitempty"`
	Owner                 string        `json:"owner,omitempty"`
	Name                  string        `json:"name,omitempty"`
	Description           string        `json:"description,omitempty"`
	AssociatedDNASubnetID string        `json:"associatedDNASubnetID,omitempty"`
	SyslogDestinationIDs  []interface{} `json:"syslogDestinationIDs,omitempty"`
}

// NewOverlayManagementSubnetProfile returns a new *OverlayManagementSubnetProfile
func NewOverlayManagementSubnetProfile() *OverlayManagementSubnetProfile {

	return &OverlayManagementSubnetProfile{}
}

// Identity returns the Identity of the object.
func (o *OverlayManagementSubnetProfile) Identity() bambou.Identity {

	return OverlayManagementSubnetProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *OverlayManagementSubnetProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *OverlayManagementSubnetProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the OverlayManagementSubnetProfile from the server
func (o *OverlayManagementSubnetProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the OverlayManagementSubnetProfile into the server
func (o *OverlayManagementSubnetProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the OverlayManagementSubnetProfile from the server
func (o *OverlayManagementSubnetProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
