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

// ModifyInstanceNetworkSpec invokes the ecs.ModifyInstanceNetworkSpec API synchronously
func (client *Client) ModifyInstanceNetworkSpec(request *ModifyInstanceNetworkSpecRequest) (response *ModifyInstanceNetworkSpecResponse, err error) {
	response = CreateModifyInstanceNetworkSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceNetworkSpecWithChan invokes the ecs.ModifyInstanceNetworkSpec API asynchronously
func (client *Client) ModifyInstanceNetworkSpecWithChan(request *ModifyInstanceNetworkSpecRequest) (<-chan *ModifyInstanceNetworkSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceNetworkSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceNetworkSpec(request)
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

// ModifyInstanceNetworkSpecWithCallback invokes the ecs.ModifyInstanceNetworkSpec API asynchronously
func (client *Client) ModifyInstanceNetworkSpecWithCallback(request *ModifyInstanceNetworkSpecRequest, callback func(response *ModifyInstanceNetworkSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceNetworkSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceNetworkSpec(request)
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

// ModifyInstanceNetworkSpecRequest is the request struct for api ModifyInstanceNetworkSpec
type ModifyInstanceNetworkSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken             string           `position:"Query" name:"ClientToken"`
	ISP                     string           `position:"Query" name:"ISP"`
	InternetMaxBandwidthOut requests.Integer `position:"Query" name:"InternetMaxBandwidthOut"`
	StartTime               string           `position:"Query" name:"StartTime"`
	AutoPay                 requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	EndTime                 string           `position:"Query" name:"EndTime"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId              string           `position:"Query" name:"InstanceId"`
	NetworkChargeType       string           `position:"Query" name:"NetworkChargeType"`
	InternetMaxBandwidthIn  requests.Integer `position:"Query" name:"InternetMaxBandwidthIn"`
	AllocatePublicIp        requests.Boolean `position:"Query" name:"AllocatePublicIp"`
}

// ModifyInstanceNetworkSpecResponse is the response struct for api ModifyInstanceNetworkSpec
type ModifyInstanceNetworkSpecResponse struct {
	*responses.BaseResponse
	OrderId   string `json:"OrderId" xml:"OrderId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceNetworkSpecRequest creates a request to invoke ModifyInstanceNetworkSpec API
func CreateModifyInstanceNetworkSpecRequest() (request *ModifyInstanceNetworkSpecRequest) {
	request = &ModifyInstanceNetworkSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceNetworkSpec", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyInstanceNetworkSpecResponse creates a response to parse from ModifyInstanceNetworkSpec response
func CreateModifyInstanceNetworkSpecResponse() (response *ModifyInstanceNetworkSpecResponse) {
	response = &ModifyInstanceNetworkSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
