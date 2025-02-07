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

// ModifyAutoProvisioningGroup invokes the ecs.ModifyAutoProvisioningGroup API synchronously
func (client *Client) ModifyAutoProvisioningGroup(request *ModifyAutoProvisioningGroupRequest) (response *ModifyAutoProvisioningGroupResponse, err error) {
	response = CreateModifyAutoProvisioningGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAutoProvisioningGroupWithChan invokes the ecs.ModifyAutoProvisioningGroup API asynchronously
func (client *Client) ModifyAutoProvisioningGroupWithChan(request *ModifyAutoProvisioningGroupRequest) (<-chan *ModifyAutoProvisioningGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyAutoProvisioningGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAutoProvisioningGroup(request)
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

// ModifyAutoProvisioningGroupWithCallback invokes the ecs.ModifyAutoProvisioningGroup API asynchronously
func (client *Client) ModifyAutoProvisioningGroupWithCallback(request *ModifyAutoProvisioningGroupRequest, callback func(response *ModifyAutoProvisioningGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAutoProvisioningGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyAutoProvisioningGroup(request)
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

// ModifyAutoProvisioningGroupRequest is the request struct for api ModifyAutoProvisioningGroup
type ModifyAutoProvisioningGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                  requests.Integer                                   `position:"Query" name:"ResourceOwnerId"`
	TerminateInstancesWithExpiration requests.Boolean                                   `position:"Query" name:"TerminateInstancesWithExpiration"`
	DefaultTargetCapacityType        string                                             `position:"Query" name:"DefaultTargetCapacityType"`
	ExcessCapacityTerminationPolicy  string                                             `position:"Query" name:"ExcessCapacityTerminationPolicy"`
	LaunchTemplateConfig             *[]ModifyAutoProvisioningGroupLaunchTemplateConfig `position:"Query" name:"LaunchTemplateConfig"  type:"Repeated"`
	ResourceOwnerAccount             string                                             `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                     string                                             `position:"Query" name:"OwnerAccount"`
	OwnerId                          requests.Integer                                   `position:"Query" name:"OwnerId"`
	AutoProvisioningGroupId          string                                             `position:"Query" name:"AutoProvisioningGroupId"`
	PayAsYouGoTargetCapacity         string                                             `position:"Query" name:"PayAsYouGoTargetCapacity"`
	TotalTargetCapacity              string                                             `position:"Query" name:"TotalTargetCapacity"`
	SpotTargetCapacity               string                                             `position:"Query" name:"SpotTargetCapacity"`
	MaxSpotPrice                     requests.Float                                     `position:"Query" name:"MaxSpotPrice"`
	AutoProvisioningGroupName        string                                             `position:"Query" name:"AutoProvisioningGroupName"`
}

// ModifyAutoProvisioningGroupLaunchTemplateConfig is a repeated param struct in ModifyAutoProvisioningGroupRequest
type ModifyAutoProvisioningGroupLaunchTemplateConfig struct {
	VSwitchId        string `name:"VSwitchId"`
	MaxPrice         string `name:"MaxPrice"`
	Priority         string `name:"Priority"`
	InstanceType     string `name:"InstanceType"`
	WeightedCapacity string `name:"WeightedCapacity"`
}

// ModifyAutoProvisioningGroupResponse is the response struct for api ModifyAutoProvisioningGroup
type ModifyAutoProvisioningGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAutoProvisioningGroupRequest creates a request to invoke ModifyAutoProvisioningGroup API
func CreateModifyAutoProvisioningGroupRequest() (request *ModifyAutoProvisioningGroupRequest) {
	request = &ModifyAutoProvisioningGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyAutoProvisioningGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyAutoProvisioningGroupResponse creates a response to parse from ModifyAutoProvisioningGroup response
func CreateModifyAutoProvisioningGroupResponse() (response *ModifyAutoProvisioningGroupResponse) {
	response = &ModifyAutoProvisioningGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
