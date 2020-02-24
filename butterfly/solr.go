/*
Package butterfly : The library for butterfly

Copyright(c) 2020 Star Inc. All Rights Reserved.
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package butterfly

import (
	"github.com/supersonictw/butterfly-solr/solr"
)

// SolrHandle :
type SolrHandle struct {
	Client *solr.SolrInterface
}

// NewSolrClient :
func NewSolrClient() *SolrHandle {
	handle := new(SolrHandle)
	client, err := solr.NewSolrInterface(Config.Solr.URI, Config.Solr.Collection)
	DeBug("Get Solr Client", err)
	handle.Client = client
	return handle
}

// Query :
func (handle *SolrHandle) Query(request string) *solr.Collection {
	query := solr.NewQuery()
	query.Q(request)
	response := handle.Client.Search(query)
	output, _ := response.Result(nil)
	return output.Results
}

// Update :
func (handle *SolrHandle) Update(data *VioletDataStruct) {
	_, err := handle.Client.UpdateDoc(data)
	DeBug("Solr Update", err)
	handle.Client.Commit()
}
