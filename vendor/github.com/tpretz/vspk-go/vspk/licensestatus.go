/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// LicenseStatusIdentity represents the Identity of the object
var LicenseStatusIdentity = bambou.Identity{
	Name:     "licensestatus",
	Category: "licensestatus",
}

// LicenseStatusList represents a list of LicenseStatus
type LicenseStatusList []*LicenseStatus

// LicenseStatusAncestor is the interface that an ancestor of a LicenseStatus must implement.
// An Ancestor is defined as an entity that has LicenseStatus as a descendant.
// An Ancestor can get a list of its child LicenseStatus, but not necessarily create one.
type LicenseStatusAncestor interface {
	LicenseStatus(*bambou.FetchingInfo) (LicenseStatusList, *bambou.Error)
}

// LicenseStatusParent is the interface that a parent of a LicenseStatus must implement.
// A Parent is defined as an entity that has LicenseStatus as a child.
// A Parent is an Ancestor which can create a LicenseStatus.
type LicenseStatusParent interface {
	LicenseStatusAncestor
	CreateLicenseStatus(*LicenseStatus) *bambou.Error
}

// LicenseStatus represents the model of a licensestatus
type LicenseStatus struct {
	ID                           string `json:"ID,omitempty"`
	ParentID                     string `json:"parentID,omitempty"`
	ParentType                   string `json:"parentType,omitempty"`
	Owner                        string `json:"owner,omitempty"`
	AccumulateLicensesEnabled    bool   `json:"accumulateLicensesEnabled"`
	TotalLicensedAVRSGsCount     int    `json:"totalLicensedAVRSGsCount,omitempty"`
	TotalLicensedAVRSsCount      int    `json:"totalLicensedAVRSsCount,omitempty"`
	TotalLicensedGatewaysCount   int    `json:"totalLicensedGatewaysCount,omitempty"`
	TotalLicensedNICsCount       int    `json:"totalLicensedNICsCount,omitempty"`
	TotalLicensedNSGsCount       int    `json:"totalLicensedNSGsCount,omitempty"`
	TotalLicensedUsedAVRSGsCount int    `json:"totalLicensedUsedAVRSGsCount,omitempty"`
	TotalLicensedUsedAVRSsCount  int    `json:"totalLicensedUsedAVRSsCount,omitempty"`
	TotalLicensedUsedNICsCount   int    `json:"totalLicensedUsedNICsCount,omitempty"`
	TotalLicensedUsedNSGsCount   int    `json:"totalLicensedUsedNSGsCount,omitempty"`
	TotalLicensedUsedVDFGsCount  int    `json:"totalLicensedUsedVDFGsCount,omitempty"`
	TotalLicensedUsedVDFsCount   int    `json:"totalLicensedUsedVDFsCount,omitempty"`
	TotalLicensedUsedVMsCount    int    `json:"totalLicensedUsedVMsCount,omitempty"`
	TotalLicensedUsedVRSGsCount  int    `json:"totalLicensedUsedVRSGsCount,omitempty"`
	TotalLicensedUsedVRSsCount   int    `json:"totalLicensedUsedVRSsCount,omitempty"`
	TotalLicensedVDFGsCount      int    `json:"totalLicensedVDFGsCount,omitempty"`
	TotalLicensedVDFsCount       int    `json:"totalLicensedVDFsCount,omitempty"`
	TotalLicensedVMsCount        int    `json:"totalLicensedVMsCount,omitempty"`
	TotalLicensedVRSGsCount      int    `json:"totalLicensedVRSGsCount,omitempty"`
	TotalLicensedVRSsCount       int    `json:"totalLicensedVRSsCount,omitempty"`
	TotalUsedGatewaysCount       int    `json:"totalUsedGatewaysCount,omitempty"`
}

// NewLicenseStatus returns a new *LicenseStatus
func NewLicenseStatus() *LicenseStatus {

	return &LicenseStatus{
		AccumulateLicensesEnabled: false,
	}
}

// Identity returns the Identity of the object.
func (o *LicenseStatus) Identity() bambou.Identity {

	return LicenseStatusIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *LicenseStatus) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *LicenseStatus) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the LicenseStatus from the server
func (o *LicenseStatus) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the LicenseStatus into the server
func (o *LicenseStatus) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the LicenseStatus from the server
func (o *LicenseStatus) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
