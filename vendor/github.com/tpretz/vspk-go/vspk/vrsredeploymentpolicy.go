/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// VRSRedeploymentpolicyIdentity represents the Identity of the object
var VRSRedeploymentpolicyIdentity = bambou.Identity{
	Name:     "vrsredeploymentpolicy",
	Category: "vrsredeploymentpolicies",
}

// VRSRedeploymentpoliciesList represents a list of VRSRedeploymentpolicies
type VRSRedeploymentpoliciesList []*VRSRedeploymentpolicy

// VRSRedeploymentpoliciesAncestor is the interface that an ancestor of a VRSRedeploymentpolicy must implement.
// An Ancestor is defined as an entity that has VRSRedeploymentpolicy as a descendant.
// An Ancestor can get a list of its child VRSRedeploymentpolicies, but not necessarily create one.
type VRSRedeploymentpoliciesAncestor interface {
	VRSRedeploymentpolicies(*bambou.FetchingInfo) (VRSRedeploymentpoliciesList, *bambou.Error)
}

// VRSRedeploymentpoliciesParent is the interface that a parent of a VRSRedeploymentpolicy must implement.
// A Parent is defined as an entity that has VRSRedeploymentpolicy as a child.
// A Parent is an Ancestor which can create a VRSRedeploymentpolicy.
type VRSRedeploymentpoliciesParent interface {
	VRSRedeploymentpoliciesAncestor
	CreateVRSRedeploymentpolicy(*VRSRedeploymentpolicy) *bambou.Error
}

// VRSRedeploymentpolicy represents the model of a vrsredeploymentpolicy
type VRSRedeploymentpolicy struct {
	ID                                     string  `json:"ID,omitempty"`
	ParentID                               string  `json:"parentID,omitempty"`
	ParentType                             string  `json:"parentType,omitempty"`
	Owner                                  string  `json:"owner,omitempty"`
	ALUbr0StatusRedeploymentEnabled        bool    `json:"ALUbr0StatusRedeploymentEnabled"`
	CPUUtilizationRedeploymentEnabled      bool    `json:"CPUUtilizationRedeploymentEnabled"`
	CPUUtilizationThreshold                float64 `json:"CPUUtilizationThreshold,omitempty"`
	VRSCorrectiveActionDelay               int     `json:"VRSCorrectiveActionDelay,omitempty"`
	VRSProcessRedeploymentEnabled          bool    `json:"VRSProcessRedeploymentEnabled"`
	VRSVSCStatusRedeploymentEnabled        bool    `json:"VRSVSCStatusRedeploymentEnabled"`
	LastUpdatedBy                          string  `json:"lastUpdatedBy,omitempty"`
	RedeploymentDelay                      int     `json:"redeploymentDelay,omitempty"`
	MemoryUtilizationRedeploymentEnabled   bool    `json:"memoryUtilizationRedeploymentEnabled"`
	MemoryUtilizationThreshold             float64 `json:"memoryUtilizationThreshold,omitempty"`
	DeploymentCountThreshold               int     `json:"deploymentCountThreshold,omitempty"`
	JesxmonProcessRedeploymentEnabled      bool    `json:"jesxmonProcessRedeploymentEnabled"`
	EntityScope                            string  `json:"entityScope,omitempty"`
	LogDiskUtilizationRedeploymentEnabled  bool    `json:"logDiskUtilizationRedeploymentEnabled"`
	LogDiskUtilizationThreshold            float64 `json:"logDiskUtilizationThreshold,omitempty"`
	RootDiskUtilizationRedeploymentEnabled bool    `json:"rootDiskUtilizationRedeploymentEnabled"`
	RootDiskUtilizationThreshold           float64 `json:"rootDiskUtilizationThreshold,omitempty"`
	ExternalID                             string  `json:"externalID,omitempty"`
}

// NewVRSRedeploymentpolicy returns a new *VRSRedeploymentpolicy
func NewVRSRedeploymentpolicy() *VRSRedeploymentpolicy {

	return &VRSRedeploymentpolicy{}
}

// Identity returns the Identity of the object.
func (o *VRSRedeploymentpolicy) Identity() bambou.Identity {

	return VRSRedeploymentpolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VRSRedeploymentpolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VRSRedeploymentpolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VRSRedeploymentpolicy from the server
func (o *VRSRedeploymentpolicy) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VRSRedeploymentpolicy into the server
func (o *VRSRedeploymentpolicy) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VRSRedeploymentpolicy from the server
func (o *VRSRedeploymentpolicy) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
