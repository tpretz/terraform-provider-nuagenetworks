/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// UserContextIdentity represents the Identity of the object
var UserContextIdentity = bambou.Identity{
	Name:     "usercontext",
	Category: "usercontexts",
}

// UserContextsList represents a list of UserContexts
type UserContextsList []*UserContext

// UserContextsAncestor is the interface that an ancestor of a UserContext must implement.
// An Ancestor is defined as an entity that has UserContext as a descendant.
// An Ancestor can get a list of its child UserContexts, but not necessarily create one.
type UserContextsAncestor interface {
	UserContexts(*bambou.FetchingInfo) (UserContextsList, *bambou.Error)
}

// UserContextsParent is the interface that a parent of a UserContext must implement.
// A Parent is defined as an entity that has UserContext as a child.
// A Parent is an Ancestor which can create a UserContext.
type UserContextsParent interface {
	UserContextsAncestor
	CreateUserContext(*UserContext) *bambou.Error
}

// UserContext represents the model of a usercontext
type UserContext struct {
	ID                     string `json:"ID,omitempty"`
	ParentID               string `json:"parentID,omitempty"`
	ParentType             string `json:"parentType,omitempty"`
	Owner                  string `json:"owner,omitempty"`
	AARFlowStatsInterval   int    `json:"AARFlowStatsInterval,omitempty"`
	AARProbeStatsInterval  int    `json:"AARProbeStatsInterval,omitempty"`
	VSSFeatureEnabled      bool   `json:"VSSFeatureEnabled"`
	VSSStatsInterval       int    `json:"VSSStatsInterval,omitempty"`
	PageSize               string `json:"pageSize,omitempty"`
	FlowCollectionEnabled  bool   `json:"flowCollectionEnabled"`
	GoogleMapsAPIKey       string `json:"googleMapsAPIKey,omitempty"`
	StatisticsEnabled      bool   `json:"statisticsEnabled"`
	StatsTSDBServerAddress string `json:"statsTSDBServerAddress,omitempty"`
}

// NewUserContext returns a new *UserContext
func NewUserContext() *UserContext {

	return &UserContext{
		AARFlowStatsInterval:  30,
		AARProbeStatsInterval: 30,
		VSSFeatureEnabled:     false,
		VSSStatsInterval:      30,
	}
}

// Identity returns the Identity of the object.
func (o *UserContext) Identity() bambou.Identity {

	return UserContextIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *UserContext) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *UserContext) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the UserContext from the server
func (o *UserContext) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the UserContext into the server
func (o *UserContext) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the UserContext from the server
func (o *UserContext) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}
