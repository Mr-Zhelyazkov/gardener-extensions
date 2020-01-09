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

package cmd

import (
	"errors"
	"io/ioutil"

	"github.com/gardener/gardener-extensions/controllers/extension-certificate-service/pkg/controller/certservice"
	"github.com/gardener/gardener-extensions/controllers/extension-certificate-service/pkg/controller/lifecycle"
	"github.com/gardener/gardener-extensions/pkg/controller/cmd"

	"github.com/gardener/gardener-extensions/controllers/extension-certificate-service/pkg/apis/config/v1alpha1"
	"github.com/gardener/gardener-extensions/controllers/extension-certificate-service/pkg/apis/config/validation"

	"github.com/gardener/gardener-extensions/controllers/extension-certificate-service/pkg/apis/config"

	controllerconfig "github.com/gardener/gardener-extensions/controllers/extension-certificate-service/pkg/controller/config"
	"github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var (
	scheme  *runtime.Scheme
	decoder runtime.Decoder
)

func init() {
	scheme = runtime.NewScheme()
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(config.AddToScheme(scheme))

	decoder = serializer.NewCodecFactory(scheme).UniversalDecoder()
}

// CertificateServiceOptions holds options related to the certificate service.
type CertificateServiceOptions struct {
	ConfigLocation string
	config         *CertificateServiceConfig
}

// AddFlags implements Flagger.AddFlags.
func (o *CertificateServiceOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.ConfigLocation, "config", "", "Path to certificate service configuration")
}

// Complete implements Completer.Complete.
func (o *CertificateServiceOptions) Complete() error {
	if o.ConfigLocation == "" {
		return errors.New("config location is not set")
	}
	data, err := ioutil.ReadFile(o.ConfigLocation)
	if err != nil {
		return err
	}

	config := config.Configuration{}
	_, _, err = decoder.Decode(data, nil, &config)
	if err != nil {
		return err
	}

	if errs := validation.ValidateConfiguration(&config); len(errs) > 0 {
		return errs.ToAggregate()
	}

	o.config = &CertificateServiceConfig{
		config: config,
	}

	return nil
}

// Completed returns the decoded CertificatesServiceConfiguration instance. Only call this if `Complete` was successful.
func (o *CertificateServiceOptions) Completed() *CertificateServiceConfig {
	return o.config
}

// CertificateServiceConfig contains configuration information about the certificate service.
type CertificateServiceConfig struct {
	config config.Configuration
}

// Apply applies the CertificateServiceOptions to the passed ControllerOptions instance.
func (c *CertificateServiceConfig) Apply(config *controllerconfig.Config) {
	config.Configuration = c.config
}

// ControllerSwitches are the cmd.SwitchOptions for the provider controllers.
func ControllerSwitches() *cmd.SwitchOptions {
	return cmd.NewSwitchOptions(
		cmd.Switch(lifecycle.ControllerName, lifecycle.AddToManager),
		cmd.Switch(certservice.ControllerName, certservice.AddToManager),
	)
}