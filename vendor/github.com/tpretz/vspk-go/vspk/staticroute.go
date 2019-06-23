/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// StaticRouteIdentity represents the Identity of the object
var StaticRouteIdentity = bambou.Identity{
	Name:     "staticroute",
	Category: "staticroutes",
}

// StaticRoutesList represents a list of StaticRoutes
type StaticRoutesList []*StaticRoute

// StaticRoutesAncestor is the interface that an ancestor of a StaticRoute must implement.
// An Ancestor is defined as an entity that has StaticRoute as a descendant.
// An Ancestor can get a list of its child StaticRoutes, but not necessarily create one.
type StaticRoutesAncestor interface {
	StaticRoutes(*bambou.FetchingInfo) (StaticRoutesList, *bambou.Error)
}

// StaticRoutesParent is the interface that a parent of a StaticRoute must implement.
// A Parent is defined as an entity that has StaticRoute as a child.
// A Parent is an Ancestor which can create a StaticRoute.
type StaticRoutesParent interface {
	StaticRoutesAncestor
	CreateStaticRoute(*StaticRoute) *bambou.Error
}

// StaticRoute represents the model of a staticroute
type StaticRoute struct {
	ID                 string `json:"ID,omitempty"`
	ParentID           string `json:"parentID,omitempty"`
	ParentType         string `json:"parentType,omitempty"`
	Owner              string `json:"owner,omitempty"`
	BFDEnabled         bool   `json:"BFDEnabled"`
	IPType             string `json:"IPType,omitempty"`
	IPv6Address        string `json:"IPv6Address,omitempty"`
	LastUpdatedBy      string `json:"lastUpdatedBy,omitempty"`
	Address            string `json:"address,omitempty"`
	Netmask            string `json:"netmask,omitempty"`
	NextHopIp          string `json:"nextHopIp,omitempty"`
	EntityScope        string `json:"entityScope,omitempty"`
	RouteDistinguisher string `json:"routeDistinguisher,omitempty"`
	AssociatedSubnetID string `json:"associatedSubnetID,omitempty"`
	ExternalID         string `json:"externalID,omitempty"`
	Type               string `json:"type,omitempty"`
}

// NewStaticRoute returns a new *StaticRoute
func NewStaticRoute() *StaticRoute {

	return &StaticRoute{
		BFDEnabled: false,
	}
}

// Identity returns the Identity of the object.
func (o *StaticRoute) Identity() bambou.Identity {

	return StaticRouteIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *StaticRoute) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *StaticRoute) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the StaticRoute from the server
func (o *StaticRoute) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the StaticRoute into the server
func (o *StaticRoute) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the StaticRoute from the server
func (o *StaticRoute) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the StaticRoute
func (o *StaticRoute) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the StaticRoute
func (o *StaticRoute) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the StaticRoute
func (o *StaticRoute) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the StaticRoute
func (o *StaticRoute) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// EventLogs retrieves the list of child EventLogs of the StaticRoute
func (o *StaticRoute) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
