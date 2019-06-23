/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// ProxyARPFilterIdentity represents the Identity of the object
var ProxyARPFilterIdentity = bambou.Identity{
	Name:     "proxyarpfilter",
	Category: "proxyarpfilters",
}

// ProxyARPFiltersList represents a list of ProxyARPFilters
type ProxyARPFiltersList []*ProxyARPFilter

// ProxyARPFiltersAncestor is the interface that an ancestor of a ProxyARPFilter must implement.
// An Ancestor is defined as an entity that has ProxyARPFilter as a descendant.
// An Ancestor can get a list of its child ProxyARPFilters, but not necessarily create one.
type ProxyARPFiltersAncestor interface {
	ProxyARPFilters(*bambou.FetchingInfo) (ProxyARPFiltersList, *bambou.Error)
}

// ProxyARPFiltersParent is the interface that a parent of a ProxyARPFilter must implement.
// A Parent is defined as an entity that has ProxyARPFilter as a child.
// A Parent is an Ancestor which can create a ProxyARPFilter.
type ProxyARPFiltersParent interface {
	ProxyARPFiltersAncestor
	CreateProxyARPFilter(*ProxyARPFilter) *bambou.Error
}

// ProxyARPFilter represents the model of a proxyarpfilter
type ProxyARPFilter struct {
	ID            string `json:"ID,omitempty"`
	ParentID      string `json:"parentID,omitempty"`
	ParentType    string `json:"parentType,omitempty"`
	Owner         string `json:"owner,omitempty"`
	IPType        string `json:"IPType,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	MaxAddress    string `json:"maxAddress,omitempty"`
	MinAddress    string `json:"minAddress,omitempty"`
	EntityScope   string `json:"entityScope,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
}

// NewProxyARPFilter returns a new *ProxyARPFilter
func NewProxyARPFilter() *ProxyARPFilter {

	return &ProxyARPFilter{
		IPType: "IPV4",
	}
}

// Identity returns the Identity of the object.
func (o *ProxyARPFilter) Identity() bambou.Identity {

	return ProxyARPFilterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ProxyARPFilter) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ProxyARPFilter) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the ProxyARPFilter from the server
func (o *ProxyARPFilter) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the ProxyARPFilter into the server
func (o *ProxyARPFilter) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the ProxyARPFilter from the server
func (o *ProxyARPFilter) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// EventLogs retrieves the list of child EventLogs of the ProxyARPFilter
func (o *ProxyARPFilter) EventLogs(info *bambou.FetchingInfo) (EventLogsList, *bambou.Error) {

	var list EventLogsList
	err := bambou.CurrentSession().FetchChildren(o, EventLogIdentity, &list, info)
	return list, err
}
