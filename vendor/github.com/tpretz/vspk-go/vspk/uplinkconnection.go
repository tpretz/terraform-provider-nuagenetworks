/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// UplinkConnectionIdentity represents the Identity of the object
var UplinkConnectionIdentity = bambou.Identity{
	Name:     "uplinkconnection",
	Category: "uplinkconnections",
}

// UplinkConnectionsList represents a list of UplinkConnections
type UplinkConnectionsList []*UplinkConnection

// UplinkConnectionsAncestor is the interface that an ancestor of a UplinkConnection must implement.
// An Ancestor is defined as an entity that has UplinkConnection as a descendant.
// An Ancestor can get a list of its child UplinkConnections, but not necessarily create one.
type UplinkConnectionsAncestor interface {
	UplinkConnections(*bambou.FetchingInfo) (UplinkConnectionsList, *bambou.Error)
}

// UplinkConnectionsParent is the interface that a parent of a UplinkConnection must implement.
// A Parent is defined as an entity that has UplinkConnection as a child.
// A Parent is an Ancestor which can create a UplinkConnection.
type UplinkConnectionsParent interface {
	UplinkConnectionsAncestor
	CreateUplinkConnection(*UplinkConnection) *bambou.Error
}

// UplinkConnection represents the model of a uplinkconnection
type UplinkConnection struct {
	ID                      string  `json:"ID,omitempty"`
	ParentID                string  `json:"parentID,omitempty"`
	ParentType              string  `json:"parentType,omitempty"`
	Owner                   string  `json:"owner,omitempty"`
	PATEnabled              bool    `json:"PATEnabled"`
	DNSAddress              string  `json:"DNSAddress,omitempty"`
	DNSAddressV6            string  `json:"DNSAddressV6,omitempty"`
	Password                string  `json:"password,omitempty"`
	Gateway                 string  `json:"gateway,omitempty"`
	GatewayV6               string  `json:"gatewayV6,omitempty"`
	Address                 string  `json:"address,omitempty"`
	AddressFamily           string  `json:"addressFamily,omitempty"`
	AddressV6               string  `json:"addressV6,omitempty"`
	AdvertisementCriteria   string  `json:"advertisementCriteria,omitempty"`
	SecondaryAddress        string  `json:"secondaryAddress,omitempty"`
	Netmask                 string  `json:"netmask,omitempty"`
	VlanId                  string  `json:"vlanId,omitempty"`
	UnderlayEnabled         bool    `json:"underlayEnabled"`
	Inherited               bool    `json:"inherited"`
	InstallerManaged        bool    `json:"installerManaged"`
	InterfaceConnectionType string  `json:"interfaceConnectionType,omitempty"`
	Mode                    string  `json:"mode,omitempty"`
	Role                    string  `json:"role,omitempty"`
	RoleOrder               string  `json:"roleOrder,omitempty"`
	PortName                string  `json:"portName,omitempty"`
	DownloadRateLimit       float64 `json:"downloadRateLimit,omitempty"`
	UplinkID                int     `json:"uplinkID,omitempty"`
	Username                string  `json:"username,omitempty"`
	AssocUnderlayID         string  `json:"assocUnderlayID,omitempty"`
	AssociatedBGPNeighborID string  `json:"associatedBGPNeighborID,omitempty"`
	AssociatedUnderlayName  string  `json:"associatedUnderlayName,omitempty"`
	AuxiliaryLink           bool    `json:"auxiliaryLink"`
}

// NewUplinkConnection returns a new *UplinkConnection
func NewUplinkConnection() *UplinkConnection {

	return &UplinkConnection{
		PATEnabled:              true,
		AddressFamily:           "IPV4",
		UnderlayEnabled:         true,
		Inherited:               false,
		InstallerManaged:        false,
		InterfaceConnectionType: "AUTOMATIC",
		Mode:              "Dynamic",
		Role:              "PRIMARY",
		DownloadRateLimit: 8,
		AuxiliaryLink:     false,
	}
}

// Identity returns the Identity of the object.
func (o *UplinkConnection) Identity() bambou.Identity {

	return UplinkConnectionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *UplinkConnection) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *UplinkConnection) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the UplinkConnection from the server
func (o *UplinkConnection) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the UplinkConnection into the server
func (o *UplinkConnection) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the UplinkConnection from the server
func (o *UplinkConnection) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// BFDSessions retrieves the list of child BFDSessions of the UplinkConnection
func (o *UplinkConnection) BFDSessions(info *bambou.FetchingInfo) (BFDSessionsList, *bambou.Error) {

	var list BFDSessionsList
	err := bambou.CurrentSession().FetchChildren(o, BFDSessionIdentity, &list, info)
	return list, err
}

// CreateBFDSession creates a new child BFDSession under the UplinkConnection
func (o *UplinkConnection) CreateBFDSession(child *BFDSession) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Underlays retrieves the list of child Underlays of the UplinkConnection
func (o *UplinkConnection) Underlays(info *bambou.FetchingInfo) (UnderlaysList, *bambou.Error) {

	var list UnderlaysList
	err := bambou.CurrentSession().FetchChildren(o, UnderlayIdentity, &list, info)
	return list, err
}

// CustomProperties retrieves the list of child CustomProperties of the UplinkConnection
func (o *UplinkConnection) CustomProperties(info *bambou.FetchingInfo) (CustomPropertiesList, *bambou.Error) {

	var list CustomPropertiesList
	err := bambou.CurrentSession().FetchChildren(o, CustomPropertyIdentity, &list, info)
	return list, err
}

// CreateCustomProperty creates a new child CustomProperty under the UplinkConnection
func (o *UplinkConnection) CreateCustomProperty(child *CustomProperty) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
