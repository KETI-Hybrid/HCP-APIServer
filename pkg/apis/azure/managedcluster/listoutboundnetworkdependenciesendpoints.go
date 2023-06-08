package managedcluster

import (
	"context"
	"encoding/json"
	"hcp-apiserver/pkg/apis"
	"io/ioutil"
	"net/http"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type ListOutboundNetworkDependenciesEndpointsResource struct {
	apis.DeleteNotSupported
	apis.GetNotSupported
	apis.PutNotSupported
}

// resourceGroupName - The name of the resource group. The name is case insensitive.
// resourceName - The name of the managed cluster resource.
// agentPoolName - The name of the agent pool.
type ListOutboundNetworkDependenciesEndpoints struct {
	ResourceGroupName string `json:"resourceGroupName"`
	ClusterName       string `json:"clusterName"`
}

func (ListOutboundNetworkDependenciesEndpointsResource) Uri() string {
	return "/aks/managedClusters/listOutboundNetworkDependenciesEndpoints"
}
func (ListOutboundNetworkDependenciesEndpointsResource) Post(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputRequest := &ListOutboundNetworkDependenciesEndpoints{}

	err = json.Unmarshal(body, inputRequest)
	if err != nil {
		klog.Errorln(err)
	}
	ctx := context.Background()
	items := ManagedClustersClient.NewListOutboundNetworkDependenciesEndpointsPager(inputRequest.ResourceGroupName, inputRequest.ClusterName, nil)
	result := make([]armcontainerservice.ManagedClustersClientListOutboundNetworkDependenciesEndpointsResponse, 0)
	for items.More() {
		tmp, err := items.NextPage(ctx)
		if err != nil {
			klog.Errorln(err)
		}

		result = append(result, tmp)
	}
	return apis.Response{Code: 200, Data: result}
}
