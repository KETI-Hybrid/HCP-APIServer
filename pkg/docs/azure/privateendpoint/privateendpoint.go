package privateendpoint

import (
	"hcp-apiserver/pkg/docs"
	"hcp-apiserver/pkg/types"

	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/julienschmidt/httprouter"
)

var PrivateEndpointConnectionsClient *armcontainerservice.PrivateEndpointConnectionsClient

func PrivateEndpointConnectionsResourceAttach(router *httprouter.Router) {
	docs.AddResource(router, new(ListResource))
	docs.AddResource(router, new(DeleteResource))
	docs.AddResource(router, new(GetResource))
	docs.AddResource(router, new(UpdateResource))
	privateEndpointConnectionsClientInit()
}

func privateEndpointConnectionsClientInit() {
	sess := types.GetAKSClient()
	PrivateEndpointConnectionsClient = sess.PrivateEndpointClient
}
