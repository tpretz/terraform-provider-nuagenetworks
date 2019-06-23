/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import "github.com/tpretz/go-bambou/bambou"

// TierIdentity represents the Identity of the object
var TierIdentity = bambou.Identity{
	Name:     "tier",
	Category: "tiers",
}

// TiersList represents a list of Tiers
type TiersList []*Tier

// TiersAncestor is the interface that an ancestor of a Tier must implement.
// An Ancestor is defined as an entity that has Tier as a descendant.
// An Ancestor can get a list of its child Tiers, but not necessarily create one.
type TiersAncestor interface {
	Tiers(*bambou.FetchingInfo) (TiersList, *bambou.Error)
}

// TiersParent is the interface that a parent of a Tier must implement.
// A Parent is defined as an entity that has Tier as a child.
// A Parent is an Ancestor which can create a Tier.
type TiersParent interface {
	TiersAncestor
	CreateTier(*Tier) *bambou.Error
}

// Tier represents the model of a tier
type Tier struct {
	ID                 string `json:"ID,omitempty"`
	ParentID           string `json:"parentID,omitempty"`
	ParentType         string `json:"parentType,omitempty"`
	Owner              string `json:"owner,omitempty"`
	PacketCount        int    `json:"packetCount,omitempty"`
	LastUpdatedBy      string `json:"lastUpdatedBy,omitempty"`
	TierType           string `json:"tierType,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	EntityScope        string `json:"entityScope,omitempty"`
	DownThresholdCount int    `json:"downThresholdCount,omitempty"`
	ProbeInterval      int    `json:"probeInterval,omitempty"`
	ExternalID         string `json:"externalID,omitempty"`
}

// NewTier returns a new *Tier
func NewTier() *Tier {

	return &Tier{
		PacketCount:        1,
		Timeout:            3000,
		DownThresholdCount: 5,
		ProbeInterval:      10,
	}
}

// Identity returns the Identity of the object.
func (o *Tier) Identity() bambou.Identity {

	return TierIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Tier) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Tier) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the Tier from the server
func (o *Tier) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the Tier into the server
func (o *Tier) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the Tier from the server
func (o *Tier) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Destinationurls retrieves the list of child Destinationurls of the Tier
func (o *Tier) Destinationurls(info *bambou.FetchingInfo) (DestinationurlsList, *bambou.Error) {

	var list DestinationurlsList
	err := bambou.CurrentSession().FetchChildren(o, DestinationurlIdentity, &list, info)
	return list, err
}

// CreateDestinationurl creates a new child Destinationurl under the Tier
func (o *Tier) CreateDestinationurl(child *Destinationurl) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// Metadatas retrieves the list of child Metadatas of the Tier
func (o *Tier) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the Tier
func (o *Tier) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the Tier
func (o *Tier) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the Tier
func (o *Tier) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
