package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDedicatedHostClusters invokes the ecs.DescribeDedicatedHostClusters API synchronously
func (client *Client) DescribeDedicatedHostClusters(request *DescribeDedicatedHostClustersRequest) (response *DescribeDedicatedHostClustersResponse, err error) {
	response = CreateDescribeDedicatedHostClustersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDedicatedHostClustersWithChan invokes the ecs.DescribeDedicatedHostClusters API asynchronously
func (client *Client) DescribeDedicatedHostClustersWithChan(request *DescribeDedicatedHostClustersRequest) (<-chan *DescribeDedicatedHostClustersResponse, <-chan error) {
	responseChan := make(chan *DescribeDedicatedHostClustersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDedicatedHostClusters(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeDedicatedHostClustersWithCallback invokes the ecs.DescribeDedicatedHostClusters API asynchronously
func (client *Client) DescribeDedicatedHostClustersWithCallback(request *DescribeDedicatedHostClustersRequest, callback func(response *DescribeDedicatedHostClustersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDedicatedHostClustersResponse
		var err error
		defer close(result)
		response, err = client.DescribeDedicatedHostClusters(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeDedicatedHostClustersRequest is the request struct for api DescribeDedicatedHostClusters
type DescribeDedicatedHostClustersRequest struct {
	*requests.RpcRequest
	DedicatedHostClusterName string                              `position:"Query" name:"DedicatedHostClusterName"`
	ResourceOwnerId          requests.Integer                    `position:"Query" name:"ResourceOwnerId"`
	DedicatedHostClusterIds  string                              `position:"Query" name:"DedicatedHostClusterIds"`
	PageNumber               requests.Integer                    `position:"Query" name:"PageNumber"`
	ResourceGroupId          string                              `position:"Query" name:"ResourceGroupId"`
	LockReason               string                              `position:"Query" name:"LockReason"`
	PageSize                 requests.Integer                    `position:"Query" name:"PageSize"`
	Tag                      *[]DescribeDedicatedHostClustersTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount     string                              `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string                              `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer                    `position:"Query" name:"OwnerId"`
	ZoneId                   string                              `position:"Query" name:"ZoneId"`
	Status                   string                              `position:"Query" name:"Status"`
}

// DescribeDedicatedHostClustersTag is a repeated param struct in DescribeDedicatedHostClustersRequest
type DescribeDedicatedHostClustersTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeDedicatedHostClustersResponse is the response struct for api DescribeDedicatedHostClusters
type DescribeDedicatedHostClustersResponse struct {
	*responses.BaseResponse
	PageSize              int                   `json:"PageSize" xml:"PageSize"`
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	PageNumber            int                   `json:"PageNumber" xml:"PageNumber"`
	TotalCount            int                   `json:"TotalCount" xml:"TotalCount"`
	DedicatedHostClusters DedicatedHostClusters `json:"DedicatedHostClusters" xml:"DedicatedHostClusters"`
}

// CreateDescribeDedicatedHostClustersRequest creates a request to invoke DescribeDedicatedHostClusters API
func CreateDescribeDedicatedHostClustersRequest() (request *DescribeDedicatedHostClustersRequest) {
	request = &DescribeDedicatedHostClustersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeDedicatedHostClusters", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDedicatedHostClustersResponse creates a response to parse from DescribeDedicatedHostClusters response
func CreateDescribeDedicatedHostClustersResponse() (response *DescribeDedicatedHostClustersResponse) {
	response = &DescribeDedicatedHostClustersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}