// Copyright 2022 Antrea Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azure

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/azure-sdk-for-go/services/resourcegraph/mgmt/2021-03-01/resourcegraph"
)

type azureNwIntfWrapper interface {
	createOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string,
		parameters armnetwork.Interface) (armnetwork.Interface, error)
	listAllComplete(ctx context.Context) ([]armnetwork.Interface, error)
}
type azureNwIntfWrapperImpl struct {
	nwIntfAPIClient armnetwork.InterfacesClient
}

func (nwIntf *azureNwIntfWrapperImpl) createOrUpdate(ctx context.Context, resourceGroupName string, networkIntfName string,
	parameters armnetwork.Interface) (armnetwork.Interface, error) {
	var nwInterface armnetwork.Interface
	nwIntfClient := nwIntf.nwIntfAPIClient
	future, err := nwIntfClient.BeginCreateOrUpdate(ctx, resourceGroupName, networkIntfName, parameters, nil)
	if err != nil {
		return nwInterface, fmt.Errorf("cannot create %v, reason: %v", networkIntfName, err)
	}

	resp, err := future.PollUntilDone(ctx, nil)
	if err != nil {
		return nwInterface, fmt.Errorf("cannot get network-interface create or update future response: %v", err)
	}

	return resp.Interface, nil
}

func (nwIntf *azureNwIntfWrapperImpl) listAllComplete(ctx context.Context) ([]armnetwork.Interface, error) {
	var networkInterfaces []armnetwork.Interface

	listResultIterator := nwIntf.nwIntfAPIClient.NewListAllPager(nil)
	for listResultIterator.More() {
		nextResult, err := listResultIterator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to iterate list of network interface, reason %v", err)
		}
		for _, v := range nextResult.Value {
			networkInterfaces = append(networkInterfaces, *v)
		}
	}

	return networkInterfaces, nil
}

type azureNsgWrapper interface {
	createOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string,
		parameters armnetwork.SecurityGroup) (nsg armnetwork.SecurityGroup, err error)
	get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, expand string) (result armnetwork.SecurityGroup,
		err error)
	delete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) error
	listAllComplete(ctx context.Context) ([]armnetwork.SecurityGroup, error)
}
type azureNsgWrapperImpl struct {
	nsgAPIClient armnetwork.SecurityGroupsClient
}

func (sg *azureNsgWrapperImpl) createOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string,
	parameters armnetwork.SecurityGroup) (armnetwork.SecurityGroup, error) {
	var nsg armnetwork.SecurityGroup
	nsgClient := sg.nsgAPIClient
	poller, err := nsgClient.BeginCreateOrUpdate(ctx, resourceGroupName, networkSecurityGroupName, parameters, nil)
	if err != nil {
		return nsg, fmt.Errorf("cannot finish create request for nsg %v, reason: %v", networkSecurityGroupName, err)
	}

	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		return nsg, fmt.Errorf("cannot get nsg create or update future response: %v", err)
	}

	return res.SecurityGroup, nil
}
func (sg *azureNsgWrapperImpl) get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string,
	expand string) (result armnetwork.SecurityGroup, err error) {
	var nsg armnetwork.SecurityGroup
	res, err := sg.nsgAPIClient.Get(ctx, resourceGroupName, networkSecurityGroupName,
		&armnetwork.SecurityGroupsClientGetOptions{Expand: nil})
	if err != nil {
		return nsg, fmt.Errorf("cannot get response for nsg get request %v for asg %s", err, networkSecurityGroupName)
	}
	return res.SecurityGroup, nil
}
func (sg *azureNsgWrapperImpl) delete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) error {
	nsgClient := sg.nsgAPIClient
	poller, err := nsgClient.BeginDelete(ctx, resourceGroupName, networkSecurityGroupName, nil)
	if err != nil {
		return fmt.Errorf("failed to finish delete request for nsg %v, reason: %v", networkSecurityGroupName, err)
	}

	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		return fmt.Errorf("cannot get nsg delete future response: %v", err)
	}

	return nil
}
func (sg *azureNsgWrapperImpl) listAllComplete(ctx context.Context) ([]armnetwork.SecurityGroup, error) {
	var nsgs []armnetwork.SecurityGroup

	pager := sg.nsgAPIClient.NewListAllPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to iterate list of security groups, reason %v", err)
		}
		for _, v := range nextResult.Value {
			nsgs = append(nsgs, *v)
		}
	}

	return nsgs, nil
}

type azureAsgWrapper interface {
	createOrUpdate(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string,
		parameters armnetwork.ApplicationSecurityGroup) (armnetwork.ApplicationSecurityGroup, error)
	get(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string) (armnetwork.ApplicationSecurityGroup, error)
	listComplete(ctx context.Context, resourceGroupName string) ([]armnetwork.ApplicationSecurityGroup, error)
	listAllComplete(ctx context.Context) ([]armnetwork.ApplicationSecurityGroup, error)
	delete(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string) error
}
type azureAsgWrapperImpl struct {
	asgAPIClient armnetwork.ApplicationSecurityGroupsClient
}

func (asg *azureAsgWrapperImpl) createOrUpdate(ctx context.Context, resourceGroupName string,
	applicationSecurityGroupName string, parameters armnetwork.ApplicationSecurityGroup) (armnetwork.ApplicationSecurityGroup, error) {
	var appsg armnetwork.ApplicationSecurityGroup
	azurePluginLogger().Info("Test: starting of create asg function", "rg", resourceGroupName)
	asgClient := asg.asgAPIClient
	poller, err := asgClient.BeginCreateOrUpdate(ctx, resourceGroupName, applicationSecurityGroupName, parameters, nil)

	azurePluginLogger().Info("Test: after BeginCreateOrUpdate", "err", err)
	if err != nil {
		return appsg, fmt.Errorf("cannot finish asg create request, asg %v, reason: %v", applicationSecurityGroupName, err)
	}

	res, err := poller.PollUntilDone(ctx, nil)
	azurePluginLogger().Info("Test: after PollUntilDone", "asg", asg, "err", err)
	if err != nil {
		return appsg, fmt.Errorf("cannot get asg create or update response: %v", err)
	}

	appsg = res.ApplicationSecurityGroup

	return appsg, nil
}
func (asg *azureAsgWrapperImpl) get(ctx context.Context, resourceGroupName string,
	applicationSecurityGroupName string) (armnetwork.ApplicationSecurityGroup, error) {
	var appsg armnetwork.ApplicationSecurityGroup
	//var respErr *azcore.ResponseError
	res, err := asg.asgAPIClient.Get(ctx, resourceGroupName, applicationSecurityGroupName, nil)
	if err != nil {
		return appsg, err
		/*
			if errors.As(err, &respErr) {
				azurePluginLogger().Info("Test: asg get call returns an error", "status code", respErr.StatusCode)
				if respErr.StatusCode != http.StatusNotFound {
					return appsg, err
				}
			}

		*/
	}

	appsg = res.ApplicationSecurityGroup
	return appsg, nil
}
func (asg *azureAsgWrapperImpl) listComplete(ctx context.Context, resourceGroupName string) ([]armnetwork.ApplicationSecurityGroup, error) {
	var asgs []armnetwork.ApplicationSecurityGroup
	listResultIterator := asg.asgAPIClient.NewListPager(resourceGroupName, nil)
	for listResultIterator.More() {
		nextResult, err := listResultIterator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to iterate list of asgs for resource-group: %v, reason %v", resourceGroupName, err)
		}
		for _, v := range nextResult.Value {
			asgs = append(asgs, *v)
		}
	}

	return asgs, nil
}
func (asg *azureAsgWrapperImpl) listAllComplete(ctx context.Context) ([]armnetwork.ApplicationSecurityGroup, error) {
	var asgs []armnetwork.ApplicationSecurityGroup
	listResultIterator := asg.asgAPIClient.NewListAllPager(nil)
	for listResultIterator.More() {
		nextResult, err := listResultIterator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to iterate list of asgs, reason %v", err)
		}
		for _, v := range nextResult.Value {
			asgs = append(asgs, *v)
		}
	}

	return asgs, nil
}
func (asg *azureAsgWrapperImpl) delete(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string) error {
	asgClient := asg.asgAPIClient
	poller, err := asgClient.BeginDelete(ctx, resourceGroupName, applicationSecurityGroupName, nil)
	if err != nil {
		return fmt.Errorf("cannot delete asg %v, reason: %v", applicationSecurityGroupName, err)
	}

	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		return fmt.Errorf("cannot get asg delete poll response: %v", err)
	}

	return nil
}

type azureResourceGraphWrapper interface {
	resources(ctx context.Context, query resourcegraph.QueryRequest) (result resourcegraph.QueryResponse, err error)
}
type azureResourceGraphWrapperImpl struct {
	resourceGraphAPIClient resourcegraph.BaseClient
}

func (rg *azureResourceGraphWrapperImpl) resources(ctx context.Context, query resourcegraph.QueryRequest) (
	result resourcegraph.QueryResponse, err error) {
	return rg.resourceGraphAPIClient.Resources(ctx, query)
}

type azureVirtualNetworksWrapper interface {
	listAllComplete(ctx context.Context) ([]armnetwork.VirtualNetwork, error)
}
type azureVirtualNetworksWrapperImpl struct {
	virtualNetworksClient armnetwork.VirtualNetworksClient
}

func (vnet *azureVirtualNetworksWrapperImpl) listAllComplete(ctx context.Context) ([]armnetwork.VirtualNetwork, error) {
	var VNListResultIterators []armnetwork.VirtualNetwork
	pager := vnet.virtualNetworksClient.NewListAllPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to iterate list of virtual networks: %q", err)
		}
		for _, v := range nextResult.Value {
			VNListResultIterators = append(VNListResultIterators, *v)
		}
	}

	return VNListResultIterators, nil
}
