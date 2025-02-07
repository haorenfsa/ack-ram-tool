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

// DescribeStorageCapacityUnits invokes the ecs.DescribeStorageCapacityUnits API synchronously
func (client *Client) DescribeStorageCapacityUnits(request *DescribeStorageCapacityUnitsRequest) (response *DescribeStorageCapacityUnitsResponse, err error) {
	response = CreateDescribeStorageCapacityUnitsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStorageCapacityUnitsWithChan invokes the ecs.DescribeStorageCapacityUnits API asynchronously
func (client *Client) DescribeStorageCapacityUnitsWithChan(request *DescribeStorageCapacityUnitsRequest) (<-chan *DescribeStorageCapacityUnitsResponse, <-chan error) {
	responseChan := make(chan *DescribeStorageCapacityUnitsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStorageCapacityUnits(request)
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

// DescribeStorageCapacityUnitsWithCallback invokes the ecs.DescribeStorageCapacityUnits API asynchronously
func (client *Client) DescribeStorageCapacityUnitsWithCallback(request *DescribeStorageCapacityUnitsRequest, callback func(response *DescribeStorageCapacityUnitsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStorageCapacityUnitsResponse
		var err error
		defer close(result)
		response, err = client.DescribeStorageCapacityUnits(request)
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

// DescribeStorageCapacityUnitsRequest is the request struct for api DescribeStorageCapacityUnits
type DescribeStorageCapacityUnitsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer                   `position:"Query" name:"ResourceOwnerId"`
	PageNumber            requests.Integer                   `position:"Query" name:"PageNumber"`
	Capacity              requests.Integer                   `position:"Query" name:"Capacity"`
	StorageCapacityUnitId *[]string                          `position:"Query" name:"StorageCapacityUnitId"  type:"Repeated"`
	PageSize              requests.Integer                   `position:"Query" name:"PageSize"`
	Tag                   *[]DescribeStorageCapacityUnitsTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount  string                             `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string                             `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer                   `position:"Query" name:"OwnerId"`
	Name                  string                             `position:"Query" name:"Name"`
	AllocationType        string                             `position:"Query" name:"AllocationType"`
	Status                *[]string                          `position:"Query" name:"Status"  type:"Repeated"`
}

// DescribeStorageCapacityUnitsTag is a repeated param struct in DescribeStorageCapacityUnitsRequest
type DescribeStorageCapacityUnitsTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeStorageCapacityUnitsResponse is the response struct for api DescribeStorageCapacityUnits
type DescribeStorageCapacityUnitsResponse struct {
	*responses.BaseResponse
	PageSize             int                  `json:"PageSize" xml:"PageSize"`
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	PageNumber           int                  `json:"PageNumber" xml:"PageNumber"`
	TotalCount           int                  `json:"TotalCount" xml:"TotalCount"`
	StorageCapacityUnits StorageCapacityUnits `json:"StorageCapacityUnits" xml:"StorageCapacityUnits"`
}

// CreateDescribeStorageCapacityUnitsRequest creates a request to invoke DescribeStorageCapacityUnits API
func CreateDescribeStorageCapacityUnitsRequest() (request *DescribeStorageCapacityUnitsRequest) {
	request = &DescribeStorageCapacityUnitsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeStorageCapacityUnits", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeStorageCapacityUnitsResponse creates a response to parse from DescribeStorageCapacityUnits response
func CreateDescribeStorageCapacityUnitsResponse() (response *DescribeStorageCapacityUnitsResponse) {
	response = &DescribeStorageCapacityUnitsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
