/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// CaptivePortalProfileIdentity represents the Identity of the object
var CaptivePortalProfileIdentity = bambou.Identity{
	Name:     "captiveportalprofile",
	Category: "captiveportalprofiles",
}

// CaptivePortalProfilesList represents a list of CaptivePortalProfiles
type CaptivePortalProfilesList []*CaptivePortalProfile

// CaptivePortalProfilesAncestor is the interface that an ancestor of a CaptivePortalProfile must implement.
// An Ancestor is defined as an entity that has CaptivePortalProfile as a descendant.
// An Ancestor can get a list of its child CaptivePortalProfiles, but not necessarily create one.
type CaptivePortalProfilesAncestor interface {
	CaptivePortalProfiles(*bambou.FetchingInfo) (CaptivePortalProfilesList, *bambou.Error)
}

// CaptivePortalProfilesParent is the interface that a parent of a CaptivePortalProfile must implement.
// A Parent is defined as an entity that has CaptivePortalProfile as a child.
// A Parent is an Ancestor which can create a CaptivePortalProfile.
type CaptivePortalProfilesParent interface {
	CaptivePortalProfilesAncestor
	CreateCaptivePortalProfile(*CaptivePortalProfile) *bambou.Error
}

// CaptivePortalProfile represents the model of a captiveportalprofile
type CaptivePortalProfile struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	Name          string `json:"name,omitempty"`
	CaptivePage   string `json:"captivePage,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Description   string `json:"description,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	PortalType    string `json:"portalType,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewCaptivePortalProfile returns a new *CaptivePortalProfile
func NewCaptivePortalProfile() *CaptivePortalProfile {

	return &CaptivePortalProfile{
		PortalType: "CLICK_THROUGH",
	}
}

// Identity returns the Identity of the object.
func (o *CaptivePortalProfile) Identity() bambou.Identity {

	return CaptivePortalProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CaptivePortalProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CaptivePortalProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the CaptivePortalProfile from the server
func (o *CaptivePortalProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the CaptivePortalProfile into the server
func (o *CaptivePortalProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the CaptivePortalProfile from the server
func (o *CaptivePortalProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
