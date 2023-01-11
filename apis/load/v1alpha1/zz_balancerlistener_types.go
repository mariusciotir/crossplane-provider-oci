/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BalancerListenerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type BalancerListenerParameters struct {

	// +kubebuilder:validation:Optional
	ConnectionConfiguration []ConnectionConfigurationParameters `json:"connectionConfiguration,omitempty" tf:"connection_configuration,omitempty"`

	// +crossplane:generate:reference:type=BalancerBackendSet
	// +kubebuilder:validation:Optional
	DefaultBackendSetName *string `json:"defaultBackendSetName,omitempty" tf:"default_backend_set_name,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultBackendSetNameRef *v1.Reference `json:"defaultBackendSetNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefaultBackendSetNameSelector *v1.Selector `json:"defaultBackendSetNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=BalancerHostname
	// +kubebuilder:validation:Optional
	HostnameNames []*string `json:"hostnameNames,omitempty" tf:"hostname_names,omitempty"`

	// +kubebuilder:validation:Optional
	HostnameNamesRefs []v1.Reference `json:"hostnameNamesRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	HostnameNamesSelector *v1.Selector `json:"hostnameNamesSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=BalancerLoadBalancer
	// +kubebuilder:validation:Optional
	LoadBalancerID *string `json:"loadBalancerId,omitempty" tf:"load_balancer_id,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerIDRef *v1.Reference `json:"loadBalancerIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LoadBalancerIDSelector *v1.Selector `json:"loadBalancerIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=BalancerPathRouteSet
	// +kubebuilder:validation:Optional
	PathRouteSetName *string `json:"pathRouteSetName,omitempty" tf:"path_route_set_name,omitempty"`

	// +kubebuilder:validation:Optional
	PathRouteSetNameRef *v1.Reference `json:"pathRouteSetNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PathRouteSetNameSelector *v1.Selector `json:"pathRouteSetNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +crossplane:generate:reference:type=BalancerLoadBalancerRoutingPolicy
	// +kubebuilder:validation:Optional
	RoutingPolicyName *string `json:"routingPolicyName,omitempty" tf:"routing_policy_name,omitempty"`

	// +kubebuilder:validation:Optional
	RoutingPolicyNameRef *v1.Reference `json:"routingPolicyNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoutingPolicyNameSelector *v1.Selector `json:"routingPolicyNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=BalancerRuleSet
	// +kubebuilder:validation:Optional
	RuleSetNames []*string `json:"ruleSetNames,omitempty" tf:"rule_set_names,omitempty"`

	// +kubebuilder:validation:Optional
	RuleSetNamesRefs []v1.Reference `json:"ruleSetNamesRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RuleSetNamesSelector *v1.Selector `json:"ruleSetNamesSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SSLConfiguration []BalancerListenerSSLConfigurationParameters `json:"sslConfiguration,omitempty" tf:"ssl_configuration,omitempty"`
}

type BalancerListenerSSLConfigurationObservation struct {
}

type BalancerListenerSSLConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CertificateIds []*string `json:"certificateIds,omitempty" tf:"certificate_ids,omitempty"`

	// +kubebuilder:validation:Optional
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	// +kubebuilder:validation:Optional
	CipherSuiteName *string `json:"cipherSuiteName,omitempty" tf:"cipher_suite_name,omitempty"`

	// +kubebuilder:validation:Optional
	Protocols []*string `json:"protocols,omitempty" tf:"protocols,omitempty"`

	// +kubebuilder:validation:Optional
	ServerOrderPreference *string `json:"serverOrderPreference,omitempty" tf:"server_order_preference,omitempty"`

	// +kubebuilder:validation:Optional
	TrustedCertificateAuthorityIds []*string `json:"trustedCertificateAuthorityIds,omitempty" tf:"trusted_certificate_authority_ids,omitempty"`

	// +kubebuilder:validation:Optional
	VerifyDepth *float64 `json:"verifyDepth,omitempty" tf:"verify_depth,omitempty"`

	// +kubebuilder:validation:Optional
	VerifyPeerCertificate *bool `json:"verifyPeerCertificate,omitempty" tf:"verify_peer_certificate,omitempty"`
}

type ConnectionConfigurationObservation struct {
}

type ConnectionConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	BackendTCPProxyProtocolVersion *float64 `json:"backendTcpProxyProtocolVersion,omitempty" tf:"backend_tcp_proxy_protocol_version,omitempty"`

	// +kubebuilder:validation:Required
	IdleTimeoutInSeconds *string `json:"idleTimeoutInSeconds" tf:"idle_timeout_in_seconds,omitempty"`
}

// BalancerListenerSpec defines the desired state of BalancerListener
type BalancerListenerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BalancerListenerParameters `json:"forProvider"`
}

// BalancerListenerStatus defines the observed state of BalancerListener.
type BalancerListenerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BalancerListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerListener is the Schema for the BalancerListeners API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type BalancerListener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BalancerListenerSpec   `json:"spec"`
	Status            BalancerListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BalancerListenerList contains a list of BalancerListeners
type BalancerListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BalancerListener `json:"items"`
}

// Repository type metadata.
var (
	BalancerListener_Kind             = "BalancerListener"
	BalancerListener_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BalancerListener_Kind}.String()
	BalancerListener_KindAPIVersion   = BalancerListener_Kind + "." + CRDGroupVersion.String()
	BalancerListener_GroupVersionKind = CRDGroupVersion.WithKind(BalancerListener_Kind)
)

func init() {
	SchemeBuilder.Register(&BalancerListener{}, &BalancerListenerList{})
}