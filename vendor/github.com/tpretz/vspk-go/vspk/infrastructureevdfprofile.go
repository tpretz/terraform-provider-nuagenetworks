/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// InfrastructureEVDFProfileIdentity represents the Identity of the object
var InfrastructureEVDFProfileIdentity = bambou.Identity{
	Name:     "infrastructureevdfprofile",
	Category: "infrastructureevdfprofiles",
}

// InfrastructureEVDFProfilesList represents a list of InfrastructureEVDFProfiles
type InfrastructureEVDFProfilesList []*InfrastructureEVDFProfile

// InfrastructureEVDFProfilesAncestor is the interface that an ancestor of a InfrastructureEVDFProfile must implement.
// An Ancestor is defined as an entity that has InfrastructureEVDFProfile as a descendant.
// An Ancestor can get a list of its child InfrastructureEVDFProfiles, but not necessarily create one.
type InfrastructureEVDFProfilesAncestor interface {
	InfrastructureEVDFProfiles(*bambou.FetchingInfo) (InfrastructureEVDFProfilesList, *bambou.Error)
}

// InfrastructureEVDFProfilesParent is the interface that a parent of a InfrastructureEVDFProfile must implement.
// A Parent is defined as an entity that has InfrastructureEVDFProfile as a child.
// A Parent is an Ancestor which can create a InfrastructureEVDFProfile.
type InfrastructureEVDFProfilesParent interface {
	InfrastructureEVDFProfilesAncestor
	CreateInfrastructureEVDFProfile(*InfrastructureEVDFProfile) *bambou.Error
}

// InfrastructureEVDFProfile represents the model of a infrastructureevdfprofile
type InfrastructureEVDFProfile struct {
	ID                string `json:"ID,omitempty"`
	ParentID          string `json:"parentID,omitempty"`
	ParentType        string `json:"parentType,omitempty"`
	Owner             string `json:"owner,omitempty"`
	NTPServerKey      string `json:"NTPServerKey,omitempty"`
	NTPServerKeyID    int    `json:"NTPServerKeyID,omitempty"`
	Name              string `json:"name,omitempty"`
	LastUpdatedBy     string `json:"lastUpdatedBy,omitempty"`
	ActiveController  string `json:"activeController,omitempty"`
	ServiceIPv4Subnet string `json:"serviceIPv4Subnet,omitempty"`
	Description       string `json:"description,omitempty"`
	EntityScope       string `json:"entityScope,omitempty"`
	ProxyDNSName      string `json:"proxyDNSName,omitempty"`
	UseTwoFactor      bool   `json:"useTwoFactor"`
	StandbyController string `json:"standbyController,omitempty"`
	NuagePlatform     string `json:"nuagePlatform,omitempty"`
	ExternalID        string `json:"externalID,omitempty"`
}

// NewInfrastructureEVDFProfile returns a new *InfrastructureEVDFProfile
func NewInfrastructureEVDFProfile() *InfrastructureEVDFProfile {

	return &InfrastructureEVDFProfile{
		ServiceIPv4Subnet: "0.0.0.0/8",
		UseTwoFactor:      false,
		NuagePlatform:     "KVM",
	}
}

// Identity returns the Identity of the object.
func (o *InfrastructureEVDFProfile) Identity() bambou.Identity {

	return InfrastructureEVDFProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *InfrastructureEVDFProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *InfrastructureEVDFProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the InfrastructureEVDFProfile from the server
func (o *InfrastructureEVDFProfile) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the InfrastructureEVDFProfile into the server
func (o *InfrastructureEVDFProfile) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the InfrastructureEVDFProfile from the server
func (o *InfrastructureEVDFProfile) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the InfrastructureEVDFProfile
func (o *InfrastructureEVDFProfile) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the InfrastructureEVDFProfile
func (o *InfrastructureEVDFProfile) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the InfrastructureEVDFProfile
func (o *InfrastructureEVDFProfile) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the InfrastructureEVDFProfile
func (o *InfrastructureEVDFProfile) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
