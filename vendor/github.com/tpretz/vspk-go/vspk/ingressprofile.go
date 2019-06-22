/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// IngressProfileIdentity represents the Identity of the object
var IngressProfileIdentity = bambou.Identity{
	Name:     "ingressprofile",
	Category: "ingressprofiles",
}

// IngressProfilesList represents a list of IngressProfiles
type IngressProfilesList []*IngressProfile

// IngressProfilesAncestor is the interface that an ancestor of a IngressProfile must implement.
// An Ancestor is defined as an entity that has IngressProfile as a descendant.
// An Ancestor can get a list of its child IngressProfiles, but not necessarily create one.
type IngressProfilesAncestor interface {
	IngressProfiles(*bambou.FetchingInfo) (IngressProfilesList, *bambou.Error)
}

// IngressProfilesParent is the interface that a parent of a IngressProfile must implement.
// A Parent is defined as an entity that has IngressProfile as a child.
// A Parent is an Ancestor which can create a IngressProfile.
type IngressProfilesParent interface {
	IngressProfilesAncestor
	CreateIngressProfile(*IngressProfile) *bambou.Error
}

// IngressProfile represents the model of a ingressprofile
type IngressProfile struct {
	ID                                 string `json:"ID,omitempty"`
	ParentID                           string `json:"parentID,omitempty"`
	ParentType                         string `json:"parentType,omitempty"`
	Owner                              string `json:"owner,omitempty"`
	Name                               string `json:"name,omitempty"`
	Description                        string `json:"description,omitempty"`
	AssociatedIPFilterProfileID        string `json:"associatedIPFilterProfileID,omitempty"`
	AssociatedIPFilterProfileName      string `json:"associatedIPFilterProfileName,omitempty"`
	AssociatedIPv6FilterProfileID      string `json:"associatedIPv6FilterProfileID,omitempty"`
	AssociatedIPv6FilterProfileName    string `json:"associatedIPv6FilterProfileName,omitempty"`
	AssociatedMACFilterProfileID       string `json:"associatedMACFilterProfileID,omitempty"`
	AssociatedMACFilterProfileName     string `json:"associatedMACFilterProfileName,omitempty"`
	AssociatedSAPIngressQoSProfileID   string `json:"associatedSAPIngressQoSProfileID,omitempty"`
	AssociatedSAPIngressQoSProfileName string `json:"associatedSAPIngressQoSProfileName,omitempty"`
}

// NewIngressProfile returns a new *IngressProfile
func NewIngressProfile() *IngressProfile {

	return &IngressProfile{}
}

// Identity returns the Identity of the object.
func (o *IngressProfile) Identity() bambou.Identity {

	return IngressProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *IngressProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *IngressProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the IngressProfile from the server
func (o *IngressProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the IngressProfile into the server
func (o *IngressProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the IngressProfile from the server
func (o *IngressProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// VPorts retrieves the list of child VPorts of the IngressProfile
func (o *IngressProfile) VPorts(info *bambou.FetchingInfo) (VPortsList, *bambou.Error) {

	var list VPortsList
	err := bambou.CurrentSession().FetchChildren(o, VPortIdentity, &list, info)
	return list, err
}
