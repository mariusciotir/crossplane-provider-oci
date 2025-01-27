/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NetworkSecurityGroupSecurityRuleIcmpOptionsObservation struct {
}

type NetworkSecurityGroupSecurityRuleIcmpOptionsParameters struct {

	// The ICMP code .
	// +kubebuilder:validation:Optional
	Code *float64 `json:"code,omitempty" tf:"code,omitempty"`

	// The ICMP type.
	// +kubebuilder:validation:Required
	Type *float64 `json:"type" tf:"type,omitempty"`
}

type NetworkSecurityGroupSecurityRuleObservation struct {

	// An Oracle-assigned identifier for the security rule. You specify this ID when you want to update or delete the rule.  Example: 04ABEC
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether the rule is valid. The value is True when the rule is first created. If the rule's source or destination is a network security group, the value changes to False if that network security group is deleted.
	IsValid *bool `json:"isValid,omitempty" tf:"is_valid,omitempty"`

	// The date and time the security rule was created. Format defined by RFC3339.
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type NetworkSecurityGroupSecurityRuleParameters struct {

	// An optional description of your choice for the rule. Avoid entering confidential information.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Conceptually, this is the range of IP addresses that a packet originating from the instance can go to.
	// +kubebuilder:validation:Optional
	Destination *string `json:"destination,omitempty" tf:"destination,omitempty"`

	// Type of destination for the rule. Required if direction = EGRESS.
	// +kubebuilder:validation:Optional
	DestinationType *string `json:"destinationType,omitempty" tf:"destination_type,omitempty"`

	// Direction of the security rule. Set to EGRESS for rules to allow outbound IP packets, or INGRESS for rules to allow inbound IP packets.
	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code as defined in:
	// +kubebuilder:validation:Optional
	IcmpOptions []NetworkSecurityGroupSecurityRuleIcmpOptionsParameters `json:"icmpOptions,omitempty" tf:"icmp_options,omitempty"`

	// The OCID of the network security group.
	// +crossplane:generate:reference:type=NetworkSecurityGroup
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupID *string `json:"networkSecurityGroupId,omitempty" tf:"network_security_group_id,omitempty"`

	// Reference to a NetworkSecurityGroup to populate networkSecurityGroupId.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIDRef *v1.Reference `json:"networkSecurityGroupIdRef,omitempty" tf:"-"`

	// Selector for a NetworkSecurityGroup to populate networkSecurityGroupId.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIDSelector *v1.Selector `json:"networkSecurityGroupIdSelector,omitempty" tf:"-"`

	// The transport protocol. Specify either all or an IPv4 protocol number as defined in Protocol Numbers. Options are supported only for ICMP ("1"), TCP ("6"), UDP ("17"), and ICMPv6 ("58").
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Conceptually, this is the range of IP addresses that a packet coming into the instance can come from.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Type of source for the rule. Required if direction = INGRESS.
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`

	// A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if egress traffic allows TCP destination port 80, there should be an ingress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic.
	// +kubebuilder:validation:Optional
	Stateless *bool `json:"stateless,omitempty" tf:"stateless,omitempty"`

	// Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed.
	// +kubebuilder:validation:Optional
	TCPOptions []NetworkSecurityGroupSecurityRuleTCPOptionsParameters `json:"tcpOptions,omitempty" tf:"tcp_options,omitempty"`

	// Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed.
	// +kubebuilder:validation:Optional
	UDPOptions []NetworkSecurityGroupSecurityRuleUDPOptionsParameters `json:"udpOptions,omitempty" tf:"udp_options,omitempty"`
}

type NetworkSecurityGroupSecurityRuleTCPOptionsDestinationPortRangeObservation struct {
}

type NetworkSecurityGroupSecurityRuleTCPOptionsDestinationPortRangeParameters struct {

	// The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// The minimum port number, which must not be greater than the maximum port number.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type NetworkSecurityGroupSecurityRuleTCPOptionsObservation struct {
}

type NetworkSecurityGroupSecurityRuleTCPOptionsParameters struct {

	// +kubebuilder:validation:Optional
	DestinationPortRange []NetworkSecurityGroupSecurityRuleTCPOptionsDestinationPortRangeParameters `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// +kubebuilder:validation:Optional
	SourcePortRange []NetworkSecurityGroupSecurityRuleTCPOptionsSourcePortRangeParameters `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`
}

type NetworkSecurityGroupSecurityRuleTCPOptionsSourcePortRangeObservation struct {
}

type NetworkSecurityGroupSecurityRuleTCPOptionsSourcePortRangeParameters struct {

	// The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// The minimum port number, which must not be greater than the maximum port number.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type NetworkSecurityGroupSecurityRuleUDPOptionsDestinationPortRangeObservation struct {
}

type NetworkSecurityGroupSecurityRuleUDPOptionsDestinationPortRangeParameters struct {

	// The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// The minimum port number, which must not be greater than the maximum port number.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

type NetworkSecurityGroupSecurityRuleUDPOptionsObservation struct {
}

type NetworkSecurityGroupSecurityRuleUDPOptionsParameters struct {

	// +kubebuilder:validation:Optional
	DestinationPortRange []NetworkSecurityGroupSecurityRuleUDPOptionsDestinationPortRangeParameters `json:"destinationPortRange,omitempty" tf:"destination_port_range,omitempty"`

	// +kubebuilder:validation:Optional
	SourcePortRange []NetworkSecurityGroupSecurityRuleUDPOptionsSourcePortRangeParameters `json:"sourcePortRange,omitempty" tf:"source_port_range,omitempty"`
}

type NetworkSecurityGroupSecurityRuleUDPOptionsSourcePortRangeObservation struct {
}

type NetworkSecurityGroupSecurityRuleUDPOptionsSourcePortRangeParameters struct {

	// The maximum port number, which must not be less than the minimum port number. To specify a single port number, set both the min and max to the same value.
	// +kubebuilder:validation:Required
	Max *float64 `json:"max" tf:"max,omitempty"`

	// The minimum port number, which must not be greater than the maximum port number.
	// +kubebuilder:validation:Required
	Min *float64 `json:"min" tf:"min,omitempty"`
}

// NetworkSecurityGroupSecurityRuleSpec defines the desired state of NetworkSecurityGroupSecurityRule
type NetworkSecurityGroupSecurityRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkSecurityGroupSecurityRuleParameters `json:"forProvider"`
}

// NetworkSecurityGroupSecurityRuleStatus defines the observed state of NetworkSecurityGroupSecurityRule.
type NetworkSecurityGroupSecurityRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkSecurityGroupSecurityRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkSecurityGroupSecurityRule is the Schema for the NetworkSecurityGroupSecurityRules API. Provides the Network Security Group Security Rule resource in Oracle Cloud Infrastructure Core service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type NetworkSecurityGroupSecurityRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityGroupSecurityRuleSpec   `json:"spec"`
	Status            NetworkSecurityGroupSecurityRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkSecurityGroupSecurityRuleList contains a list of NetworkSecurityGroupSecurityRules
type NetworkSecurityGroupSecurityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityGroupSecurityRule `json:"items"`
}

// Repository type metadata.
var (
	NetworkSecurityGroupSecurityRule_Kind             = "NetworkSecurityGroupSecurityRule"
	NetworkSecurityGroupSecurityRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkSecurityGroupSecurityRule_Kind}.String()
	NetworkSecurityGroupSecurityRule_KindAPIVersion   = NetworkSecurityGroupSecurityRule_Kind + "." + CRDGroupVersion.String()
	NetworkSecurityGroupSecurityRule_GroupVersionKind = CRDGroupVersion.WithKind(NetworkSecurityGroupSecurityRule_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkSecurityGroupSecurityRule{}, &NetworkSecurityGroupSecurityRuleList{})
}
