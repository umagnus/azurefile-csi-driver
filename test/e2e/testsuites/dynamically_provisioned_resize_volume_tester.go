/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testsuites

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/test/e2e/framework"

	"sigs.k8s.io/azurefile-csi-driver/pkg/azurefile"
	volumehelper "sigs.k8s.io/azurefile-csi-driver/pkg/util"
	"sigs.k8s.io/azurefile-csi-driver/test/e2e/driver"
	"sigs.k8s.io/azurefile-csi-driver/test/utils/azure"
	"sigs.k8s.io/azurefile-csi-driver/test/utils/credentials"
)

// DynamicallyProvisionedResizeVolumeTest will provision required StorageClass(es), PVC(s) and Pod(s)
// Waiting for the PV provisioner to resize the PV
// Testing if the PV is resized successfully.
type DynamicallyProvisionedResizeVolumeTest struct {
	CSIDriver              driver.DynamicPVTestDriver
	Pods                   []PodDetails
	StorageClassParameters map[string]string
}

var secretNameFormatRE = regexp.MustCompile(`azure-storage-account-(.+)-secret`)

func getStorageAccountName(secretName string) (string, error) {
	matches := secretNameFormatRE.FindStringSubmatch(secretName)
	if len(matches) != 2 {
		return "", fmt.Errorf("could not get account name from %s, correct format: %s", secretName, secretNameFormatRE)
	}
	return matches[1], nil
}

func (t *DynamicallyProvisionedResizeVolumeTest) Run(ctx context.Context, client clientset.Interface, namespace *v1.Namespace) {
	for _, pod := range t.Pods {
		tpod, cleanup := pod.SetupWithDynamicVolumes(ctx, client, namespace, t.CSIDriver, t.StorageClassParameters)
		for i := range cleanup {
			defer cleanup[i](ctx)
		}

		ginkgo.By("deploying the pod")
		tpod.Create(ctx)
		defer tpod.Cleanup(ctx)
		ginkgo.By("checking that the pods command exits with no error")
		tpod.WaitForSuccess(ctx)

		pvcName := tpod.pod.Spec.Volumes[0].VolumeSource.PersistentVolumeClaim.ClaimName
		pvc, err := client.CoreV1().PersistentVolumeClaims(namespace.Name).Get(ctx, pvcName, metav1.GetOptions{})
		if err != nil {
			framework.ExpectNoError(err, fmt.Sprintf("fail to get original pvc(%s): %v", pvcName, err))
		}

		originalSize := pvc.Spec.Resources.Requests["storage"]
		delta := resource.Quantity{}
		delta.Set(volumehelper.GiBToBytes(1))
		originalSize.Add(delta)
		pvc.Spec.Resources.Requests["storage"] = originalSize

		ginkgo.By("resizing the pvc")
		updatedPvc, err := client.CoreV1().PersistentVolumeClaims(namespace.Name).Update(ctx, pvc, metav1.UpdateOptions{})
		if err != nil {
			framework.ExpectNoError(err, fmt.Sprintf("fail to resize pvc(%s): %v", pvcName, err))
		}
		updatedSize := updatedPvc.Spec.Resources.Requests["storage"]

		ginkgo.By("sleep 30s waiting for resize complete")
		time.Sleep(30 * time.Second)

		ginkgo.By("checking the resizing result")
		newPvc, err := client.CoreV1().PersistentVolumeClaims(namespace.Name).Get(ctx, tpod.pod.Spec.Volumes[0].VolumeSource.PersistentVolumeClaim.ClaimName, metav1.GetOptions{})
		if err != nil {
			framework.ExpectNoError(err, fmt.Sprintf("fail to get new pvc(%s): %v", pvcName, err))
		}
		newSize := newPvc.Spec.Resources.Requests["storage"]
		if !newSize.Equal(updatedSize) {
			framework.Failf("newSize(%+v) is not equal to updatedSize(%+v)", newSize, updatedSize)
		}

		ginkgo.By("checking the resizing PV result")
		newPv, _ := client.CoreV1().PersistentVolumes().Get(ctx, updatedPvc.Spec.VolumeName, metav1.GetOptions{})
		newPvSize := newPv.Spec.Capacity["storage"]
		newPvSizeStr := newPvSize.String() + "Gi"

		if !strings.Contains(newPvSizeStr, newSize.String()) {
			framework.Failf("newPVCSize(%+v) is not equal to newPVSize(%+v)", newSize.String(), newPvSizeStr)
		}

		ginkgo.By("checking the resizing azurefile result")

		creds, err := credentials.CreateAzureCredentialFile(false)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		azureClient, err := azure.GetAzureClient(creds.Cloud, creds.SubscriptionID, creds.AADClientID, creds.TenantID, creds.AADClientSecret)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())

		var resourceGroup string
		var accountName string
		var shareName string
		if newPv.Spec.PersistentVolumeSource.CSI != nil {
			resourceGroup, accountName, shareName, _, _, _, err = azurefile.GetFileShareInfo(newPv.Spec.PersistentVolumeSource.CSI.VolumeHandle)
			framework.ExpectNoError(err, fmt.Sprintf("Error getting filesource for azurefile %v", err))
		} else if newPv.Spec.PersistentVolumeSource.AzureFile != nil {
			resourceGroup = creds.ResourceGroup
			azureSource := newPv.Spec.PersistentVolumeSource.AzureFile
			accountName, err = getStorageAccountName(azureSource.SecretName)
			if err != nil {
				framework.Failf("getStorageAccountName(%s) returned with error: %v", azureSource.SecretName, err)
			}
			shareName = azureSource.ShareName
		}

		//get file information
		fileshareClient, err := azureClient.GetAzureFilesClient()
		framework.ExpectNoError(err, fmt.Sprintf("Error getting client for azurefile %v", err))
		share, err := fileshareClient.Get(ctx, resourceGroup, accountName, shareName)
		framework.ExpectNoError(err, fmt.Sprintf("Error getting file for azurefile %v", err))
		newfileSize := strconv.Itoa(int(*share.FileShareProperties.ShareQuota)) + "Gi"
		if !(newSize.String() == newfileSize) {
			framework.Failf("newPVCSize(%+v) is not equal to new azurefileSize(%+v)", newSize.String(), newfileSize)
		}
	}
}
