/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// AutoDiscoverHypervisorFromClusterIdentity represents the Identity of the object
var AutoDiscoverHypervisorFromClusterIdentity = bambou.Identity{
	Name:     "autodiscoveredhypervisor",
	Category: "autodiscoveredhypervisors",
}

// AutoDiscoverHypervisorFromClustersList represents a list of AutoDiscoverHypervisorFromClusters
type AutoDiscoverHypervisorFromClustersList []*AutoDiscoverHypervisorFromCluster

// AutoDiscoverHypervisorFromClustersAncestor is the interface that an ancestor of a AutoDiscoverHypervisorFromCluster must implement.
// An Ancestor is defined as an entity that has AutoDiscoverHypervisorFromCluster as a descendant.
// An Ancestor can get a list of its child AutoDiscoverHypervisorFromClusters, but not necessarily create one.
type AutoDiscoverHypervisorFromClustersAncestor interface {
	AutoDiscoverHypervisorFromClusters(*bambou.FetchingInfo) (AutoDiscoverHypervisorFromClustersList, *bambou.Error)
}

// AutoDiscoverHypervisorFromClustersParent is the interface that a parent of a AutoDiscoverHypervisorFromCluster must implement.
// A Parent is defined as an entity that has AutoDiscoverHypervisorFromCluster as a child.
// A Parent is an Ancestor which can create a AutoDiscoverHypervisorFromCluster.
type AutoDiscoverHypervisorFromClustersParent interface {
	AutoDiscoverHypervisorFromClustersAncestor
	CreateAutoDiscoverHypervisorFromCluster(*AutoDiscoverHypervisorFromCluster) *bambou.Error
}

// AutoDiscoverHypervisorFromCluster represents the model of a autodiscoveredhypervisor
type AutoDiscoverHypervisorFromCluster struct {
	ID            string        `json:"ID,omitempty"`
	ParentID      string        `json:"parentID,omitempty"`
	ParentType    string        `json:"parentType,omitempty"`
	Owner         string        `json:"owner,omitempty"`
	LastUpdatedBy string        `json:"lastUpdatedBy,omitempty"`
	NetworkList   []interface{} `json:"networkList,omitempty"`
	EntityScope   string        `json:"entityScope,omitempty"`
	AssocEntityID string        `json:"assocEntityID,omitempty"`
	ExternalID    string        `json:"externalID,omitempty"`
	HypervisorIP  string        `json:"hypervisorIP,omitempty"`
}

// NewAutoDiscoverHypervisorFromCluster returns a new *AutoDiscoverHypervisorFromCluster
func NewAutoDiscoverHypervisorFromCluster() *AutoDiscoverHypervisorFromCluster {

	return &AutoDiscoverHypervisorFromCluster{}
}

// Identity returns the Identity of the object.
func (o *AutoDiscoverHypervisorFromCluster) Identity() bambou.Identity {

	return AutoDiscoverHypervisorFromClusterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AutoDiscoverHypervisorFromCluster) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AutoDiscoverHypervisorFromCluster) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the AutoDiscoverHypervisorFromCluster from the server
func (o *AutoDiscoverHypervisorFromCluster) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the AutoDiscoverHypervisorFromCluster into the server
func (o *AutoDiscoverHypervisorFromCluster) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the AutoDiscoverHypervisorFromCluster from the server
func (o *AutoDiscoverHypervisorFromCluster) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
