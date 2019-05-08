// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	config "github.com/gardener/gardener-extensions/controllers/extension-certificate-service/pkg/apis/config"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ACME)(nil), (*config.ACME)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ACME_To_config_ACME(a.(*ACME), b.(*config.ACME), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ACME)(nil), (*ACME)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ACME_To_v1alpha1_ACME(a.(*config.ACME), b.(*ACME), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CloudDNS)(nil), (*config.CloudDNS)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudDNS_To_config_CloudDNS(a.(*CloudDNS), b.(*config.CloudDNS), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.CloudDNS)(nil), (*CloudDNS)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_CloudDNS_To_v1alpha1_CloudDNS(a.(*config.CloudDNS), b.(*CloudDNS), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Configuration)(nil), (*config.Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Configuration_To_config_Configuration(a.(*Configuration), b.(*config.Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Configuration)(nil), (*Configuration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Configuration_To_v1alpha1_Configuration(a.(*config.Configuration), b.(*Configuration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigurationSpec)(nil), (*config.ConfigurationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ConfigurationSpec_To_config_ConfigurationSpec(a.(*ConfigurationSpec), b.(*config.ConfigurationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ConfigurationSpec)(nil), (*ConfigurationSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ConfigurationSpec_To_v1alpha1_ConfigurationSpec(a.(*config.ConfigurationSpec), b.(*ConfigurationSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DNSProviders)(nil), (*config.DNSProviders)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DNSProviders_To_config_DNSProviders(a.(*DNSProviders), b.(*config.DNSProviders), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DNSProviders)(nil), (*DNSProviders)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DNSProviders_To_v1alpha1_DNSProviders(a.(*config.DNSProviders), b.(*DNSProviders), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Route53)(nil), (*config.Route53)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Route53_To_config_Route53(a.(*Route53), b.(*config.Route53), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Route53)(nil), (*Route53)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Route53_To_v1alpha1_Route53(a.(*config.Route53), b.(*Route53), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ACME_To_config_ACME(in *ACME, out *config.ACME, s conversion.Scope) error {
	out.Email = in.Email
	out.Server = in.Server
	out.PrivateKey = (*string)(unsafe.Pointer(in.PrivateKey))
	return nil
}

// Convert_v1alpha1_ACME_To_config_ACME is an autogenerated conversion function.
func Convert_v1alpha1_ACME_To_config_ACME(in *ACME, out *config.ACME, s conversion.Scope) error {
	return autoConvert_v1alpha1_ACME_To_config_ACME(in, out, s)
}

func autoConvert_config_ACME_To_v1alpha1_ACME(in *config.ACME, out *ACME, s conversion.Scope) error {
	out.Email = in.Email
	out.Server = in.Server
	out.PrivateKey = (*string)(unsafe.Pointer(in.PrivateKey))
	return nil
}

// Convert_config_ACME_To_v1alpha1_ACME is an autogenerated conversion function.
func Convert_config_ACME_To_v1alpha1_ACME(in *config.ACME, out *ACME, s conversion.Scope) error {
	return autoConvert_config_ACME_To_v1alpha1_ACME(in, out, s)
}

func autoConvert_v1alpha1_CloudDNS_To_config_CloudDNS(in *CloudDNS, out *config.CloudDNS, s conversion.Scope) error {
	out.Domains = *(*[]string)(unsafe.Pointer(&in.Domains))
	out.Name = in.Name
	out.Project = in.Project
	out.ServiceAccount = in.ServiceAccount
	return nil
}

// Convert_v1alpha1_CloudDNS_To_config_CloudDNS is an autogenerated conversion function.
func Convert_v1alpha1_CloudDNS_To_config_CloudDNS(in *CloudDNS, out *config.CloudDNS, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudDNS_To_config_CloudDNS(in, out, s)
}

func autoConvert_config_CloudDNS_To_v1alpha1_CloudDNS(in *config.CloudDNS, out *CloudDNS, s conversion.Scope) error {
	out.Domains = *(*[]string)(unsafe.Pointer(&in.Domains))
	out.Name = in.Name
	out.Project = in.Project
	out.ServiceAccount = in.ServiceAccount
	return nil
}

// Convert_config_CloudDNS_To_v1alpha1_CloudDNS is an autogenerated conversion function.
func Convert_config_CloudDNS_To_v1alpha1_CloudDNS(in *config.CloudDNS, out *CloudDNS, s conversion.Scope) error {
	return autoConvert_config_CloudDNS_To_v1alpha1_CloudDNS(in, out, s)
}

func autoConvert_v1alpha1_Configuration_To_config_Configuration(in *Configuration, out *config.Configuration, s conversion.Scope) error {
	if err := Convert_v1alpha1_ConfigurationSpec_To_config_ConfigurationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Configuration_To_config_Configuration is an autogenerated conversion function.
func Convert_v1alpha1_Configuration_To_config_Configuration(in *Configuration, out *config.Configuration, s conversion.Scope) error {
	return autoConvert_v1alpha1_Configuration_To_config_Configuration(in, out, s)
}

func autoConvert_config_Configuration_To_v1alpha1_Configuration(in *config.Configuration, out *Configuration, s conversion.Scope) error {
	if err := Convert_config_ConfigurationSpec_To_v1alpha1_ConfigurationSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_Configuration_To_v1alpha1_Configuration is an autogenerated conversion function.
func Convert_config_Configuration_To_v1alpha1_Configuration(in *config.Configuration, out *Configuration, s conversion.Scope) error {
	return autoConvert_config_Configuration_To_v1alpha1_Configuration(in, out, s)
}

func autoConvert_v1alpha1_ConfigurationSpec_To_config_ConfigurationSpec(in *ConfigurationSpec, out *config.ConfigurationSpec, s conversion.Scope) error {
	out.LifecycleSync = in.LifecycleSync
	out.ServiceSync = in.ServiceSync
	out.IssuerName = in.IssuerName
	out.NamespaceRef = in.NamespaceRef
	out.ResourceNamespace = in.ResourceNamespace
	if err := Convert_v1alpha1_ACME_To_config_ACME(&in.ACME, &out.ACME, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_DNSProviders_To_config_DNSProviders(&in.Providers, &out.Providers, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ConfigurationSpec_To_config_ConfigurationSpec is an autogenerated conversion function.
func Convert_v1alpha1_ConfigurationSpec_To_config_ConfigurationSpec(in *ConfigurationSpec, out *config.ConfigurationSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ConfigurationSpec_To_config_ConfigurationSpec(in, out, s)
}

func autoConvert_config_ConfigurationSpec_To_v1alpha1_ConfigurationSpec(in *config.ConfigurationSpec, out *ConfigurationSpec, s conversion.Scope) error {
	out.LifecycleSync = in.LifecycleSync
	out.ServiceSync = in.ServiceSync
	out.IssuerName = in.IssuerName
	out.NamespaceRef = in.NamespaceRef
	out.ResourceNamespace = in.ResourceNamespace
	if err := Convert_config_ACME_To_v1alpha1_ACME(&in.ACME, &out.ACME, s); err != nil {
		return err
	}
	if err := Convert_config_DNSProviders_To_v1alpha1_DNSProviders(&in.Providers, &out.Providers, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_ConfigurationSpec_To_v1alpha1_ConfigurationSpec is an autogenerated conversion function.
func Convert_config_ConfigurationSpec_To_v1alpha1_ConfigurationSpec(in *config.ConfigurationSpec, out *ConfigurationSpec, s conversion.Scope) error {
	return autoConvert_config_ConfigurationSpec_To_v1alpha1_ConfigurationSpec(in, out, s)
}

func autoConvert_v1alpha1_DNSProviders_To_config_DNSProviders(in *DNSProviders, out *config.DNSProviders, s conversion.Scope) error {
	out.Route53 = *(*[]config.Route53)(unsafe.Pointer(&in.Route53))
	out.CloudDNS = *(*[]config.CloudDNS)(unsafe.Pointer(&in.CloudDNS))
	return nil
}

// Convert_v1alpha1_DNSProviders_To_config_DNSProviders is an autogenerated conversion function.
func Convert_v1alpha1_DNSProviders_To_config_DNSProviders(in *DNSProviders, out *config.DNSProviders, s conversion.Scope) error {
	return autoConvert_v1alpha1_DNSProviders_To_config_DNSProviders(in, out, s)
}

func autoConvert_config_DNSProviders_To_v1alpha1_DNSProviders(in *config.DNSProviders, out *DNSProviders, s conversion.Scope) error {
	out.Route53 = *(*[]Route53)(unsafe.Pointer(&in.Route53))
	out.CloudDNS = *(*[]CloudDNS)(unsafe.Pointer(&in.CloudDNS))
	return nil
}

// Convert_config_DNSProviders_To_v1alpha1_DNSProviders is an autogenerated conversion function.
func Convert_config_DNSProviders_To_v1alpha1_DNSProviders(in *config.DNSProviders, out *DNSProviders, s conversion.Scope) error {
	return autoConvert_config_DNSProviders_To_v1alpha1_DNSProviders(in, out, s)
}

func autoConvert_v1alpha1_Route53_To_config_Route53(in *Route53, out *config.Route53, s conversion.Scope) error {
	out.Domains = *(*[]string)(unsafe.Pointer(&in.Domains))
	out.Name = in.Name
	out.Region = in.Region
	out.AccessKeyID = in.AccessKeyID
	out.SecretAccessKey = in.SecretAccessKey
	return nil
}

// Convert_v1alpha1_Route53_To_config_Route53 is an autogenerated conversion function.
func Convert_v1alpha1_Route53_To_config_Route53(in *Route53, out *config.Route53, s conversion.Scope) error {
	return autoConvert_v1alpha1_Route53_To_config_Route53(in, out, s)
}

func autoConvert_config_Route53_To_v1alpha1_Route53(in *config.Route53, out *Route53, s conversion.Scope) error {
	out.Domains = *(*[]string)(unsafe.Pointer(&in.Domains))
	out.Name = in.Name
	out.Region = in.Region
	out.AccessKeyID = in.AccessKeyID
	out.SecretAccessKey = in.SecretAccessKey
	return nil
}

// Convert_config_Route53_To_v1alpha1_Route53 is an autogenerated conversion function.
func Convert_config_Route53_To_v1alpha1_Route53(in *config.Route53, out *Route53, s conversion.Scope) error {
	return autoConvert_config_Route53_To_v1alpha1_Route53(in, out, s)
}