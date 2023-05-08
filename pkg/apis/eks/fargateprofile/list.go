package fargateprofile

import (
	"encoding/json"
	"hcp-apiserver/pkg/apis"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/julienschmidt/httprouter"
	"k8s.io/klog"
)

type ListResource struct {
	apis.PostNotSupported
	apis.PutNotSupported
	apis.DeleteNotSupported
}
type FargateProfileListInput struct {
	ClusterName string `json:"clusterName"`
	MaxResults  int64  `json:"maxResults"`
	NextToken   string `json:"nextToken"`
}

func (ListResource) Uri() string {
	return "/eks/fargate-profile/list"
}
func (ListResource) Get(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) apis.Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln(err)
	}
	inputReq := &FargateProfileListInput{}
	err = json.Unmarshal(body, inputReq)
	if err != nil {
		klog.Errorln(err)
	}
	realInput := &eks.ListFargateProfilesInput{
		ClusterName: aws.String(inputReq.ClusterName),
		MaxResults:  aws.Int64(inputReq.MaxResults),
		NextToken:   aws.String(inputReq.NextToken),
	}

	result, err := FargateProfileClient.ListFargateProfiles(realInput)
	if err != nil {
		klog.Errorln(err)
	}
	return apis.Response{Code: 200, Data: result}
}
