/*
Copyright 2019 The OpenShift Authors.

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

package azure

import (
	"context"
	"fmt"

	minterv1 "github.com/openshift/cloud-credential-operator/pkg/apis/cloudcredential/v1"
	"github.com/openshift/cloud-credential-operator/pkg/operator/constants"
	"github.com/openshift/cloud-credential-operator/pkg/operator/credentialsrequest/actuator"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

var RootSecretKey = client.ObjectKey{Name: constants.AzureCloudCredSecretName, Namespace: constants.CloudCredSecretNamespace}

type clientWrapper struct {
	client.Client
	RootCredClient client.Client
}

func newClientWrapper(c, rootC client.Client) *clientWrapper {
	return &clientWrapper{Client: c, RootCredClient: rootC}
}

func (cw *clientWrapper) RootSecret(ctx context.Context) (*secret, error) {
	secret, err := cw.secret(ctx, cw.RootCredClient, RootSecretKey)
	if err != nil {
		return nil, err
	}

	if !secret.HasAnnotation() {
		return nil, &actuator.ActuatorError{
			ErrReason: minterv1.CredentialsProvisionFailure,
			Message:   fmt.Sprintf("cannot proceed without cloud cred secret annotation %+v", secret),
		}
	}

	return secret, nil
}

func (cw *clientWrapper) Secret(ctx context.Context, key client.ObjectKey) (*secret, error) {
	return cw.secret(ctx, cw.Client, key)
}

func (cw *clientWrapper) secret(ctx context.Context, c client.Client, key client.ObjectKey) (*secret, error) {
	s := &secret{}
	if err := c.Get(ctx, key, &s.Secret); err != nil {
		return nil, err
	}
	return s, nil
}

func (cw *clientWrapper) Mode(ctx context.Context) (string, error) {
	rs, err := cw.RootSecret(ctx)
	if err != nil {
		return "", err
	}

	return rs.Annotations[constants.AnnotationKey], nil
}
