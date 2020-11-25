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

// AllocatePublicIpAddress invokes the ecs.AllocatePublicIpAddress API synchronously
func (client *Client) AllocatePublicIpAddress(request *AllocatePublicIpAddressRequest) (response *AllocatePublicIpAddressResponse, err error) {
	response = CreateAllocatePublicIpAddressResponse()
	err = client.DoAction(request, response)
	return
}

// AllocatePublicIpAddressWithChan invokes the ecs.AllocatePublicIpAddress API asynchronously
func (client *Client) AllocatePublicIpAddressWithChan(request *AllocatePublicIpAddressRequest) (<-chan *AllocatePublicIpAddressResponse, <-chan error) {
	responseChan := make(chan *AllocatePublicIpAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocatePublicIpAddress(request)
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

// AllocatePublicIpAddressWithCallback invokes the ecs.AllocatePublicIpAddress API asynchronously
func (client *Client) AllocatePublicIpAddressWithCallback(request *AllocatePublicIpAddressRequest, callback func(response *AllocatePublicIpAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocatePublicIpAddressResponse
		var err error
		defer close(result)
		response, err = client.AllocatePublicIpAddress(request)
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

// AllocatePublicIpAddressRequest is the request struct for api AllocatePublicIpAddress
type AllocatePublicIpAddressRequest struct {
	*requests.RpcRequest
	IpAddress            string           `position:"Query" name:"IpAddress"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	VlanId               string           `position:"Query" name:"VlanId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// AllocatePublicIpAddressResponse is the response struct for api AllocatePublicIpAddress
type AllocatePublicIpAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	IpAddress string `json:"IpAddress" xml:"IpAddress"`
}

// CreateAllocatePublicIpAddressRequest creates a request to invoke AllocatePublicIpAddress API
func CreateAllocatePublicIpAddressRequest() (request *AllocatePublicIpAddressRequest) {
	request = &AllocatePublicIpAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AllocatePublicIpAddress", "", "")
	request.Method = requests.POST
	return
}

// CreateAllocatePublicIpAddressResponse creates a response to parse from AllocatePublicIpAddress response
func CreateAllocatePublicIpAddressResponse() (response *AllocatePublicIpAddressResponse) {
	response = &AllocatePublicIpAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
