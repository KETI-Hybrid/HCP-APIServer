package nodepool

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type SetSizeResource struct {
	docs.GetNotSupported
	docs.DeleteNotSupported
	docs.PutNotSupported
}

type Size struct {
	ProjectName  string `json:"projectName"`
	LocationName string `json:"locationName"`
	ClusterName  string `json:"clusterName"`
	NodePoolName string `json:"nodepoolName"`
	NodeCount    int64  `json:"nodeCount"`
}

func (SetSizeResource) Uri() string {
	return "/gke/zone/cluster/nodePools/setSize"
}

func (SetSizeResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(Size{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
