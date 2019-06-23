/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// NetconfProfileIdentity represents the Identity of the object
var NetconfProfileIdentity = bambou.Identity{
	Name:     "netconfprofile",
	Category: "netconfprofiles",
}

// NetconfProfilesList represents a list of NetconfProfiles
type NetconfProfilesList []*NetconfProfile

// NetconfProfilesAncestor is the interface that an ancestor of a NetconfProfile must implement.
// An Ancestor is defined as an entity that has NetconfProfile as a descendant.
// An Ancestor can get a list of its child NetconfProfiles, but not necessarily create one.
type NetconfProfilesAncestor interface {
	NetconfProfiles(*bambou.FetchingInfo) (NetconfProfilesList, *bambou.Error)
}

// NetconfProfilesParent is the interface that a parent of a NetconfProfile must implement.
// A Parent is defined as an entity that has NetconfProfile as a child.
// A Parent is an Ancestor which can create a NetconfProfile.
type NetconfProfilesParent interface {
	NetconfProfilesAncestor
	CreateNetconfProfile(*NetconfProfile) *bambou.Error
}

// NetconfProfile represents the model of a netconfprofile
type NetconfProfile struct {
	ID          string `json:"ID,omitempty"`
	ParentID    string `json:"parentID,omitempty"`
	ParentType  string `json:"parentType,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Name        string `json:"name,omitempty"`
	Password    string `json:"password,omitempty"`
	Description string `json:"description,omitempty"`
	Port        int    `json:"port,omitempty"`
	UserName    string `json:"userName,omitempty"`
}

// NewNetconfProfile returns a new *NetconfProfile
func NewNetconfProfile() *NetconfProfile {

	return &NetconfProfile{
		Port: 830,
	}
}

// Identity returns the Identity of the object.
func (o *NetconfProfile) Identity() bambou.Identity {

	return NetconfProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NetconfProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NetconfProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the NetconfProfile from the server
func (o *NetconfProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the NetconfProfile into the server
func (o *NetconfProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the NetconfProfile from the server
func (o *NetconfProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
