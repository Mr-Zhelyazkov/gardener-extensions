// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package infrastructure

import (
	"fmt"

	"github.com/golang/mock/gomock"

	gcpv1alpha1 "github.com/gardener/gardener-extensions/controllers/provider-gcp/pkg/apis/gcp/v1alpha1"
	"github.com/gardener/gardener-extensions/controllers/provider-gcp/pkg/internal"
	"github.com/gardener/gardener-extensions/pkg/controller"
	mockterraformer "github.com/gardener/gardener-extensions/pkg/mock/gardener-extensions/terraformer"

	gardencorev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var _ = Describe("Terraform", func() {
	var (
		infra              *extensionsv1alpha1.Infrastructure
		config             *gcpv1alpha1.InfrastructureConfig
		cluster            *controller.Cluster
		projectID          string
		serviceAccountData []byte
		serviceAccount     *internal.ServiceAccount

		minPortsPerVM = int32(2048)

		podsCIDR     = "11.0.0.0/16"
		servicesCIDR = "12.0.0.0/16"

		ctrl = gomock.NewController(GinkgoT())
		tf   = mockterraformer.NewMockTerraformer(ctrl)
	)

	BeforeEach(func() {
		tf = mockterraformer.NewMockTerraformer(ctrl)

		internalCIDR := "192.168.0.0/16"

		config = &gcpv1alpha1.InfrastructureConfig{
			Networks: gcpv1alpha1.NetworkConfig{
				VPC: &gcpv1alpha1.VPC{
					Name: "vpc",
					CloudRouter: &gcpv1alpha1.CloudRouter{
						Name: "cloudrouter",
					},
				},
				Internal: &internalCIDR,
				Worker:   "10.1.0.0/16",
			},
		}

		infra = &extensionsv1alpha1.Infrastructure{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "foo",
				Name:      "bar",
			},

			Spec: extensionsv1alpha1.InfrastructureSpec{
				Region: "eu-west-1",
				SecretRef: corev1.SecretReference{
					Namespace: "foo",
					Name:      "gcp-credentials",
				},
				ProviderConfig: &runtime.RawExtension{
					Object: config,
				},
			},
		}

		cluster = &controller.Cluster{
			Shoot: &gardencorev1alpha1.Shoot{
				Spec: gardencorev1alpha1.ShootSpec{
					Networking: gardencorev1alpha1.Networking{
						Pods:     &podsCIDR,
						Services: &servicesCIDR,
					},
				},
			},
		}

		projectID = "project"
		serviceAccountData = []byte(fmt.Sprintf(`{"project_id": "%s"}`, projectID))
		serviceAccount = &internal.ServiceAccount{ProjectID: projectID, Raw: serviceAccountData}
	})

	Describe("#ExtractTerraformState", func() {
		It("should return correct state when cloudRouter name is specified", func() {
			var (
				cloudRouterName             = "test"
				vpcWithoutCloudRouterConfig = &gcpv1alpha1.InfrastructureConfig{
					Networks: gcpv1alpha1.NetworkConfig{
						VPC: &gcpv1alpha1.VPC{
							Name:        "vpc",
							CloudRouter: &gcpv1alpha1.CloudRouter{Name: cloudRouterName},
						},
						Worker: "10.1.0.0/16",
					},
				}

				outputKeys = []string{
					TerraformerOutputKeyVPCName,
					TerraformerOutputKeySubnetNodes,
					TerraformerOutputKeyServiceAccountEmail,
					TerraformOutputKeyCloudRouter,
					TerraformOutputKeyCloudNAT,
				}
			)

			var (
				vpcName             = "vpc"
				subnetNodes         = "subnet"
				serviceAccountEmail = "email"
				cloudNATName        = "cloudnat"
			)

			tf.EXPECT().GetStateOutputVariables(outputKeys).DoAndReturn(func(_ ...string) (map[string]string, error) {
				return map[string]string{
					TerraformerOutputKeyVPCName:             vpcName,
					TerraformerOutputKeySubnetNodes:         subnetNodes,
					TerraformerOutputKeyServiceAccountEmail: serviceAccountEmail,
					TerraformOutputKeyCloudRouter:           cloudRouterName,
					TerraformOutputKeyCloudNAT:              cloudNATName,
				}, nil
			})

			state, err := ExtractTerraformState(tf, vpcWithoutCloudRouterConfig)
			Expect(err).NotTo(HaveOccurred())
			Expect(state).To(Equal(&TerraformState{
				VPCName:             vpcName,
				SubnetNodes:         subnetNodes,
				ServiceAccountEmail: serviceAccountEmail,
				CloudRouterName:     cloudRouterName,
				CloudNATName:        cloudNATName,
			}))
		})
		It("should return correct state when cloudRouter name is NOT specified", func() {
			var (
				vpcWithoutCloudRouterConfig = &gcpv1alpha1.InfrastructureConfig{
					Networks: gcpv1alpha1.NetworkConfig{
						VPC: &gcpv1alpha1.VPC{
							Name: "vpc",
						},
						Worker: "10.1.0.0/16",
					},
				}

				outputKeys = []string{
					TerraformerOutputKeyVPCName,
					TerraformerOutputKeySubnetNodes,
					TerraformerOutputKeyServiceAccountEmail,
				}
			)

			var (
				vpcName             = "vpc"
				subnetNodes         = "subnet"
				serviceAccountEmail = "email"
			)

			tf.EXPECT().GetStateOutputVariables(outputKeys).DoAndReturn(func(_ ...string) (map[string]string, error) {
				return map[string]string{
					TerraformerOutputKeyVPCName:             vpcName,
					TerraformerOutputKeySubnetNodes:         subnetNodes,
					TerraformerOutputKeyServiceAccountEmail: serviceAccountEmail,
				}, nil
			})

			state, err := ExtractTerraformState(tf, vpcWithoutCloudRouterConfig)
			Expect(err).NotTo(HaveOccurred())
			Expect(state).To(Equal(&TerraformState{
				VPCName:             vpcName,
				SubnetNodes:         subnetNodes,
				ServiceAccountEmail: serviceAccountEmail,
			}))
		})
	})

	Describe("#ComputeTerraformerChartValues", func() {
		It("should correctly compute the terraformer chart values", func() {
			values := ComputeTerraformerChartValues(infra, serviceAccount, config, cluster)

			Expect(values).To(Equal(map[string]interface{}{
				"google": map[string]interface{}{
					"region":  infra.Spec.Region,
					"project": projectID,
				},
				"create": map[string]interface{}{
					"vpc":         false,
					"cloudRouter": false,
				},
				"vpc": map[string]interface{}{
					"name": config.Networks.VPC.Name,
					"cloudRouter": map[string]interface{}{
						"name": "cloudrouter",
					},
				},
				"clusterName": infra.Namespace,
				"networks": map[string]interface{}{
					"pods":     podsCIDR,
					"services": servicesCIDR,
					"worker":   config.Networks.Worker,
					"internal": config.Networks.Internal,
					"cloudNAT": map[string]interface{}{
						"minPortsPerVM": minPortsPerVM,
					},
				},
				"outputKeys": map[string]interface{}{
					"vpcName":             TerraformerOutputKeyVPCName,
					"cloudNAT":            TerraformOutputKeyCloudNAT,
					"cloudRouter":         TerraformOutputKeyCloudRouter,
					"serviceAccountEmail": TerraformerOutputKeyServiceAccountEmail,
					"subnetNodes":         TerraformerOutputKeySubnetNodes,
					"subnetInternal":      TerraformerOutputKeySubnetInternal,
				},
			}))
		})

		It("should correctly compute the terraformer chart values with vpc creation", func() {
			config.Networks.VPC = nil
			values := ComputeTerraformerChartValues(infra, serviceAccount, config, cluster)

			Expect(values).To(Equal(map[string]interface{}{
				"google": map[string]interface{}{
					"region":  infra.Spec.Region,
					"project": projectID,
				},
				"create": map[string]interface{}{
					"vpc":         true,
					"cloudRouter": true,
				},
				"vpc": map[string]interface{}{
					"name": DefaultVPCName,
				},
				"clusterName": infra.Namespace,
				"networks": map[string]interface{}{
					"pods":     podsCIDR,
					"services": servicesCIDR,
					"worker":   config.Networks.Worker,
					"internal": config.Networks.Internal,
					"cloudNAT": map[string]interface{}{
						"minPortsPerVM": minPortsPerVM,
					},
				},
				"outputKeys": map[string]interface{}{
					"vpcName":             TerraformerOutputKeyVPCName,
					"cloudNAT":            TerraformOutputKeyCloudNAT,
					"cloudRouter":         TerraformOutputKeyCloudRouter,
					"serviceAccountEmail": TerraformerOutputKeyServiceAccountEmail,
					"subnetNodes":         TerraformerOutputKeySubnetNodes,
					"subnetInternal":      TerraformerOutputKeySubnetInternal,
				},
			}))
		})
	})

	Describe("#StatusFromTerraformState", func() {
		var (
			serviceAccountEmail string
			vpcName             string
			cloudRouterName     string
			cloudNATName        string
			subnetNodes         string
			subnetInternal      string

			state *TerraformState
		)

		BeforeEach(func() {
			serviceAccountEmail = "gardener@cloud"
			vpcName = "vpc-name"
			cloudRouterName = "cloudrouter-name"
			cloudNATName = "cloudnat-name"
			subnetNodes = "nodes-subnet"
			subnetInternal = "internal"

			state = &TerraformState{
				VPCName:             vpcName,
				CloudRouterName:     cloudRouterName,
				CloudNATName:        cloudNATName,
				ServiceAccountEmail: serviceAccountEmail,
				SubnetNodes:         subnetNodes,
				SubnetInternal:      &subnetInternal,
			}
		})

		It("should correctly compute the status", func() {
			status := StatusFromTerraformState(state)

			Expect(status).To(Equal(&gcpv1alpha1.InfrastructureStatus{
				TypeMeta: StatusTypeMeta,
				Networks: gcpv1alpha1.NetworkStatus{
					VPC: gcpv1alpha1.VPC{
						Name:        vpcName,
						CloudRouter: &gcpv1alpha1.CloudRouter{Name: cloudRouterName},
					},
					Subnets: []gcpv1alpha1.Subnet{
						{
							Purpose: gcpv1alpha1.PurposeNodes,
							Name:    subnetNodes,
						},
						{
							Purpose: gcpv1alpha1.PurposeInternal,
							Name:    subnetInternal,
						},
					},
				},
				ServiceAccountEmail: serviceAccountEmail,
			}))
		})

		It("should correctly compute the status without internal subnet", func() {
			state.SubnetInternal = nil
			status := StatusFromTerraformState(state)

			Expect(status).To(Equal(&gcpv1alpha1.InfrastructureStatus{
				TypeMeta: StatusTypeMeta,
				Networks: gcpv1alpha1.NetworkStatus{
					VPC: gcpv1alpha1.VPC{
						Name:        vpcName,
						CloudRouter: &gcpv1alpha1.CloudRouter{Name: cloudRouterName},
					},
					Subnets: []gcpv1alpha1.Subnet{
						{
							Purpose: gcpv1alpha1.PurposeNodes,
							Name:    subnetNodes,
						},
					},
				},
				ServiceAccountEmail: serviceAccountEmail,
			}))
		})
	})
})
