// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by client-gen. DO NOT EDIT.
package azclient

import (
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/accountclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/availabilitysetclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/blobcontainerclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/blobservicepropertiesclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/deploymentclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/diskclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/fileshareclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/identityclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/interfaceclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/ipgroupclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/loadbalancerclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/managedclusterclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/privateendpointclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatelinkserviceclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatezoneclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/providerclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipaddressclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipprefixclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/registryclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/resourcegroupclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/roleassignmentclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/routetableclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/secretclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/securitygroupclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/snapshotclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/sshpublickeyresourceclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/subnetclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/vaultclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachineclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetvmclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualnetworkclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualnetworklinkclient"
)

type ClientFactory interface {
	GetAccountClient() accountclient.Interface
	GetAccountClientForSub(subscriptionID string) (accountclient.Interface, error)
	GetAvailabilitySetClient() availabilitysetclient.Interface
	GetBlobContainerClient() blobcontainerclient.Interface
	GetBlobContainerClientForSub(subscriptionID string) (blobcontainerclient.Interface, error)
	GetBlobServicePropertiesClient() blobservicepropertiesclient.Interface
	GetDeploymentClient() deploymentclient.Interface
	GetDiskClient() diskclient.Interface
	GetDiskClientForSub(subscriptionID string) (diskclient.Interface, error)
	GetFileShareClient() fileshareclient.Interface
	GetFileShareClientForSub(subscriptionID string) (fileshareclient.Interface, error)
	GetIdentityClient() identityclient.Interface
	GetInterfaceClient() interfaceclient.Interface
	GetIPGroupClient() ipgroupclient.Interface
	GetLoadBalancerClient() loadbalancerclient.Interface
	GetManagedClusterClient() managedclusterclient.Interface
	GetPrivateEndpointClient() privateendpointclient.Interface
	GetPrivateLinkServiceClient() privatelinkserviceclient.Interface
	GetPrivateZoneClient() privatezoneclient.Interface
	GetProviderClient() providerclient.Interface
	GetPublicIPAddressClient() publicipaddressclient.Interface
	GetPublicIPPrefixClient() publicipprefixclient.Interface
	GetRegistryClient() registryclient.Interface
	GetResourceGroupClient() resourcegroupclient.Interface
	GetRoleAssignmentClient() roleassignmentclient.Interface
	GetRouteTableClient() routetableclient.Interface
	GetSecretClient() secretclient.Interface
	GetSecurityGroupClient() securitygroupclient.Interface
	GetSnapshotClient() snapshotclient.Interface
	GetSnapshotClientForSub(subscriptionID string) (snapshotclient.Interface, error)
	GetSSHPublicKeyResourceClient() sshpublickeyresourceclient.Interface
	GetSubnetClient() subnetclient.Interface
	GetVaultClient() vaultclient.Interface
	GetVirtualMachineClient() virtualmachineclient.Interface
	GetVirtualMachineScaleSetClient() virtualmachinescalesetclient.Interface
	GetVirtualMachineScaleSetVMClient() virtualmachinescalesetvmclient.Interface
	GetVirtualNetworkClient() virtualnetworkclient.Interface
	GetVirtualNetworkLinkClient() virtualnetworklinkclient.Interface
}
