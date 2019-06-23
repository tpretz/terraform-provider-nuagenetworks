/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// VRSMetricsIdentity represents the Identity of the object
var VRSMetricsIdentity = bambou.Identity{
	Name:     "vrsmetrics",
	Category: "vrsmetrics",
}

// VRSMetricsList represents a list of VRSMetrics
type VRSMetricsList []*VRSMetrics

// VRSMetricsAncestor is the interface that an ancestor of a VRSMetrics must implement.
// An Ancestor is defined as an entity that has VRSMetrics as a descendant.
// An Ancestor can get a list of its child VRSMetrics, but not necessarily create one.
type VRSMetricsAncestor interface {
	VRSMetrics(*bambou.FetchingInfo) (VRSMetricsList, *bambou.Error)
}

// VRSMetricsParent is the interface that a parent of a VRSMetrics must implement.
// A Parent is defined as an entity that has VRSMetrics as a child.
// A Parent is an Ancestor which can create a VRSMetrics.
type VRSMetricsParent interface {
	VRSMetricsAncestor
	CreateVRSMetrics(*VRSMetrics) *bambou.Error
}

// VRSMetrics represents the model of a vrsmetrics
type VRSMetrics struct {
	ID                            string  `json:"ID,omitempty"`
	ParentID                      string  `json:"parentID,omitempty"`
	ParentType                    string  `json:"parentType,omitempty"`
	Owner                         string  `json:"owner,omitempty"`
	ALUbr0Status                  bool    `json:"ALUbr0Status"`
	CPUUtilization                float64 `json:"CPUUtilization,omitempty"`
	VRSProcess                    bool    `json:"VRSProcess"`
	VRSVSCStatus                  bool    `json:"VRSVSCStatus"`
	LastUpdatedBy                 string  `json:"lastUpdatedBy,omitempty"`
	ReDeploy                      bool    `json:"reDeploy"`
	ReceivingMetrics              bool    `json:"receivingMetrics"`
	MemoryUtilization             float64 `json:"memoryUtilization,omitempty"`
	JesxmonProcess                bool    `json:"jesxmonProcess"`
	EntityScope                   string  `json:"entityScope,omitempty"`
	LogDiskPartitionUtilization   float64 `json:"logDiskPartitionUtilization,omitempty"`
	RootDiskPartitionUtilization  float64 `json:"rootDiskPartitionUtilization,omitempty"`
	AppliedMetricsPushInterval    int     `json:"appliedMetricsPushInterval,omitempty"`
	AssociatedVCenterHypervisorID string  `json:"associatedVCenterHypervisorID,omitempty"`
	CurrentVersion                string  `json:"currentVersion,omitempty"`
	ExternalID                    string  `json:"externalID,omitempty"`
}

// NewVRSMetrics returns a new *VRSMetrics
func NewVRSMetrics() *VRSMetrics {

	return &VRSMetrics{
		AppliedMetricsPushInterval: 60,
	}
}

// Identity returns the Identity of the object.
func (o *VRSMetrics) Identity() bambou.Identity {

	return VRSMetricsIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *VRSMetrics) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *VRSMetrics) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the VRSMetrics from the server
func (o *VRSMetrics) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the VRSMetrics into the server
func (o *VRSMetrics) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the VRSMetrics from the server
func (o *VRSMetrics) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
