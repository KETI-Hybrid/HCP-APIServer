package cluster

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/docs/util"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/container/v1"
)

type StartIpRotationResource struct {
	docs.DeleteNotSupported
	docs.GetNotSupported
	docs.PutNotSupported
}

type StartIpRotation struct {
	ProjectName       string `json:"projectName"`
	LocationName      string `json:"locationName"`
	ClusterName       string `json:"clusterName"`
	RotateCredentials bool   `json:"rotateCredentials"`
}

func (StartIpRotationResource) Uri() string {
	return "/gke/zone/cluster/startIpRotation"
}

func (StartIpRotationResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	request, response := util.DocWithReq(StartIpRotation{}, container.Operation{})

	resp := docs.ForDoc{
		Req:  request,
		Resp: response,
	}
	return docs.Response{Code: 200, Data: resp}
}
