/*
   No Header

   to add a custom header, add a file
   named __code_header in your sdk user vanilla
   and add whatever header you like
   Don't forget to make it a comment
*/

package vspk

import (
	"crypto/tls"
	"fmt"
	"github.com/tpretz/go-bambou/bambou"
	"strings"
)

var (
	urlpostfix string
)

// Returns a new Session -- authentication using username + password
func NewSession(username, password, organization, url string) (*bambou.Session, *Me) {

	root := NewMe()
	url += urlpostfix

	session := bambou.NewSession(username, password, organization, url, root)

	return session, root
}

// Returns a new Session -- authentication using X509 certificate.
func NewX509Session(cert *tls.Certificate, url string) (*bambou.Session, *Me) {

	root := NewMe()
	url += urlpostfix

	session := bambou.NewX509Session(cert, url, root)

	return session, root
}

func init() {

	urlpostfix = "/" + SDKAPIPrefix + "/v" + strings.Replace(fmt.Sprintf("%.1f", SDKAPIVersion), ".", "_", 100)
}
