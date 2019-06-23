/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// AutoDiscoverClusterIdentity represents the Identity of the object
var AutoDiscoverClusterIdentity = bambou.Identity{
	Name:     "autodiscoveredcluster",
	Category: "autodiscoveredclusters",
}

// AutoDiscoverClustersList represents a list of AutoDiscoverClusters
type AutoDiscoverClustersList []*AutoDiscoverCluster

// AutoDiscoverClustersAncestor is the interface that an ancestor of a AutoDiscoverCluster must implement.
// An Ancestor is defined as an entity that has AutoDiscoverCluster as a descendant.
// An Ancestor can get a list of its child AutoDiscoverClusters, but not necessarily create one.
type AutoDiscoverClustersAncestor interface {
	AutoDiscoverClusters(*bambou.FetchingInfo) (AutoDiscoverClustersList, *bambou.Error)
}

// AutoDiscoverClustersParent is the interface that a parent of a AutoDiscoverCluster must implement.
// A Parent is defined as an entity that has AutoDiscoverCluster as a child.
// A Parent is an Ancestor which can create a AutoDiscoverCluster.
type AutoDiscoverClustersParent interface {
	AutoDiscoverClustersAncestor
	CreateAutoDiscoverCluster(*AutoDiscoverCluster) *bambou.Error
}

// AutoDiscoverCluster represents the model of a autodiscoveredcluster
type AutoDiscoverCluster struct {
	ID                       string `json:"ID,omitempty"`
	ParentID                 string `json:"parentID,omitempty"`
	ParentType               string `json:"parentType,omitempty"`
	Owner                    string `json:"owner,omitempty"`
	Name                     string `json:"name,omitempty"`
	ManagedObjectID          string `json:"managedObjectID,omitempty"`
	LastUpdatedBy            string `json:"lastUpdatedBy,omitempty"`
	EntityScope              string `json:"entityScope,omitempty"`
	AssocVCenterDataCenterID string `json:"assocVCenterDataCenterID,omitempty"`
	ExternalID               string `json:"externalID,omitempty"`
}

// NewAutoDiscoverCluster returns a new *AutoDiscoverCluster
func NewAutoDiscoverCluster() *AutoDiscoverCluster {

	return &AutoDiscoverCluster{}
}

// Identity returns the Identity of the object.
func (o *AutoDiscoverCluster) Identity() bambou.Identity {

	return AutoDiscoverClusterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AutoDiscoverCluster) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AutoDiscoverCluster) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the AutoDiscoverCluster from the server
func (o *AutoDiscoverCluster) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the AutoDiscoverCluster into the server
func (o *AutoDiscoverCluster) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the AutoDiscoverCluster from the server
func (o *AutoDiscoverCluster) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
