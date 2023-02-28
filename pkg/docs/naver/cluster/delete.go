package cluster

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type DeleteResource struct {
	docs.PutNotSupported
	docs.PostNotSupported
	docs.GetNotSupported
}

type Delete struct {
	ClusterUUID string `json:"clusteruuid"`
}

type DeleteResp struct {
	ClusterUUID string `json:"clusteruuid"`
	Status      bool   `json:"status"`
}

func (DeleteResource) Uri() string {
	return "/nks/cluster/delete"
}

func (DeleteResource) Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) docs.Response {
	client := types.GetNKSClient()
	containerService := client.Client.V2Api
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &Delete{}
	resp := &DeleteResp{}
	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	err = containerService.ClustersUuidDelete(ctx, &inputRequest.ClusterUUID)
	if err != nil {
		klog.Errorln(err)
	}
	resp.ClusterUUID = inputRequest.ClusterUUID
	resp.Status = true
	return docs.Response{Code: 200, Data: resp}
}
