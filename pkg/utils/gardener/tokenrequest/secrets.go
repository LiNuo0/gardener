// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package tokenrequest

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientcmdlatest "k8s.io/client-go/tools/clientcmd/api/latest"
	clientcmdv1 "k8s.io/client-go/tools/clientcmd/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	resourcesv1alpha1 "github.com/gardener/gardener/pkg/apis/resources/v1alpha1"
	"github.com/gardener/gardener/pkg/client/kubernetes"
	"github.com/gardener/gardener/pkg/utils/flow"
	gardenerutils "github.com/gardener/gardener/pkg/utils/gardener"
	secretsutils "github.com/gardener/gardener/pkg/utils/secrets"
	secretsmanager "github.com/gardener/gardener/pkg/utils/secrets/manager"
)

var tokenRequestorRequirement *labels.Requirement

func init() {
	var err error
	tokenRequestorRequirement, err = labels.NewRequirement(resourcesv1alpha1.ResourceManagerPurpose, selection.Equals, []string{resourcesv1alpha1.LabelPurposeTokenRequest})
	utilruntime.Must(err)
}

// GenerateGenericTokenKubeconfig generates a generic token kubeconfig in the given namespace for the given kube-apiserver address.
// In case of a rotation, the old kubeconfig is kept in the cluster.
func GenerateGenericTokenKubeconfig(ctx context.Context, secretsManager secretsmanager.Interface, namespace, kubeAPIServerAddress string) (*corev1.Secret, error) {
	clusterCABundleSecret, found := secretsManager.Get(v1beta1constants.SecretNameCACluster)
	if !found {
		return nil, fmt.Errorf("secret %q not found", v1beta1constants.SecretNameCACluster)
	}

	config := &secretsutils.KubeconfigSecretConfig{
		Name:        v1beta1constants.SecretNameGenericTokenKubeconfig,
		ContextName: namespace,
		Cluster: clientcmdv1.Cluster{
			Server:                   kubeAPIServerAddress,
			CertificateAuthorityData: clusterCABundleSecret.Data[secretsutils.DataKeyCertificateBundle],
		},
		AuthInfo: clientcmdv1.AuthInfo{
			TokenFile: gardenerutils.PathShootToken,
		},
	}

	// Keep old kubeconfig secret to give components outside `gardener/gardener` the chance to update their secret refs.
	return secretsManager.Generate(ctx, config, secretsmanager.Rotate(secretsmanager.KeepOld))
}

// RenewAccessSecrets drops the serviceaccount.resources.gardener.cloud/token-renew-timestamp annotation from all
// access secrets selected by the given list options.
// This will make the token-requestor controller in gardener-resource-manager/gardenlet issue new tokens immediately.
func RenewAccessSecrets(ctx context.Context, c client.Client, opts ...client.ListOption) error {
	listOptions := &client.ListOptions{}
	listOptions.ApplyOptions(opts)

	// Add resources.gardener.cloud/purpose=token-requestor selector requirement.
	// We cannot use MatchingLabels directly, as it would overwrite other label selectors given in opts.
	if listOptions.LabelSelector == nil {
		listOptions.LabelSelector = labels.NewSelector()
	}
	listOptions.LabelSelector = listOptions.LabelSelector.Add(*tokenRequestorRequirement)

	secretList := &corev1.SecretList{}
	if err := c.List(ctx, secretList, listOptions); err != nil {
		return err
	}

	var fns []flow.TaskFn

	for _, obj := range secretList.Items {
		secret := obj

		fns = append(fns, func(ctx context.Context) error {
			patch := client.MergeFrom(secret.DeepCopy())
			delete(secret.Annotations, resourcesv1alpha1.ServiceAccountTokenRenewTimestamp)
			return c.Patch(ctx, &secret, patch)
		})
	}

	return flow.Parallel(fns...)(ctx)
}

// IsTokenPopulated checks if a `kubeconfig` secret already contains a token.
func IsTokenPopulated(secret *corev1.Secret) (bool, error) {
	kubeconfig := &clientcmdv1.Config{}
	if _, _, err := clientcmdlatest.Codec.Decode(secret.Data[kubernetes.KubeConfig], nil, kubeconfig); err != nil {
		return false, err
	}

	var userName string
	for _, namedContext := range kubeconfig.Contexts {
		if namedContext.Name == kubeconfig.CurrentContext {
			userName = namedContext.Context.AuthInfo
			break
		}
	}

	for _, users := range kubeconfig.AuthInfos {
		if users.Name == userName {
			if len(users.AuthInfo.Token) > 0 {
				return true, nil
			}
			return false, nil
		}
	}

	return false, nil
}
