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

// ModifyRouterInterfaceSpec invokes the ecs.ModifyRouterInterfaceSpec API synchronously
func (client *Client) ModifyRouterInterfaceSpec(request *ModifyRouterInterfaceSpecRequest) (response *ModifyRouterInterfaceSpecResponse, err error) {
	response = CreateModifyRouterInterfaceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyRouterInterfaceSpecWithChan invokes the ecs.ModifyRouterInterfaceSpec API asynchronously
func (client *Client) ModifyRouterInterfaceSpecWithChan(request *ModifyRouterInterfaceSpecRequest) (<-chan *ModifyRouterInterfaceSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyRouterInterfaceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyRouterInterfaceSpec(request)
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

// ModifyRouterInterfaceSpecWithCallback invokes the ecs.ModifyRouterInterfaceSpec API asynchronously
func (client *Client) ModifyRouterInterfaceSpecWithCallback(request *ModifyRouterInterfaceSpecRequest, callback func(response *ModifyRouterInterfaceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyRouterInterfaceSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyRouterInterfaceSpec(request)
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

// ModifyRouterInterfaceSpecRequest is the request struct for api ModifyRouterInterfaceSpec
type ModifyRouterInterfaceSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Spec                 string           `position:"Query" name:"Spec"`
	UserCidr             string           `position:"Query" name:"UserCidr"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	RouterInterfaceId    string           `position:"Query" name:"RouterInterfaceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyRouterInterfaceSpecResponse is the response struct for api ModifyRouterInterfaceSpec
type ModifyRouterInterfaceSpecResponse struct {
	*responses.BaseResponse
	Spec      string `json:"Spec" xml:"Spec"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyRouterInterfaceSpecRequest creates a request to invoke ModifyRouterInterfaceSpec API
func CreateModifyRouterInterfaceSpecRequest() (request *ModifyRouterInterfaceSpecRequest) {
	request = &ModifyRouterInterfaceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyRouterInterfaceSpec", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyRouterInterfaceSpecResponse creates a response to parse from ModifyRouterInterfaceSpec response
func CreateModifyRouterInterfaceSpecResponse() (response *ModifyRouterInterfaceSpecResponse) {
	response = &ModifyRouterInterfaceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
