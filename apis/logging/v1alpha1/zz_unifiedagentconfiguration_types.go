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

type DestinationObservation struct {
}

type DestinationParameters struct {

	// +crossplane:generate:reference:type=Log
	// +kubebuilder:validation:Optional
	LogObjectID *string `json:"logObjectId,omitempty" tf:"log_object_id,omitempty"`

	// +kubebuilder:validation:Optional
	LogObjectIDRef *v1.Reference `json:"logObjectIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LogObjectIDSelector *v1.Selector `json:"logObjectIdSelector,omitempty" tf:"-"`
}

type GroupAssociationObservation struct {
}

type GroupAssociationParameters struct {

	// +kubebuilder:validation:Optional
	GroupList []*string `json:"groupList,omitempty" tf:"group_list,omitempty"`
}

type ParserObservation struct {
}

type ParserParameters struct {

	// +kubebuilder:validation:Optional
	Delimiter *string `json:"delimiter,omitempty" tf:"delimiter,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	FieldTimeKey *string `json:"fieldTimeKey,omitempty" tf:"field_time_key,omitempty"`

	// +kubebuilder:validation:Optional
	Format []*string `json:"format,omitempty" tf:"format,omitempty"`

	// +kubebuilder:validation:Optional
	FormatFirstline *string `json:"formatFirstline,omitempty" tf:"format_firstline,omitempty"`

	// +kubebuilder:validation:Optional
	GrokFailureKey *string `json:"grokFailureKey,omitempty" tf:"grok_failure_key,omitempty"`

	// +kubebuilder:validation:Optional
	GrokNameKey *string `json:"grokNameKey,omitempty" tf:"grok_name_key,omitempty"`

	// +kubebuilder:validation:Optional
	IsEstimateCurrentEvent *bool `json:"isEstimateCurrentEvent,omitempty" tf:"is_estimate_current_event,omitempty"`

	// +kubebuilder:validation:Optional
	IsKeepTimeKey *bool `json:"isKeepTimeKey,omitempty" tf:"is_keep_time_key,omitempty"`

	// +kubebuilder:validation:Optional
	IsNullEmptyString *bool `json:"isNullEmptyString,omitempty" tf:"is_null_empty_string,omitempty"`

	// +kubebuilder:validation:Optional
	IsSupportColonlessIdent *bool `json:"isSupportColonlessIdent,omitempty" tf:"is_support_colonless_ident,omitempty"`

	// +kubebuilder:validation:Optional
	IsWithPriority *bool `json:"isWithPriority,omitempty" tf:"is_with_priority,omitempty"`

	// +kubebuilder:validation:Optional
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`

	// +kubebuilder:validation:Optional
	MessageFormat *string `json:"messageFormat,omitempty" tf:"message_format,omitempty"`

	// +kubebuilder:validation:Optional
	MessageKey *string `json:"messageKey,omitempty" tf:"message_key,omitempty"`

	// +kubebuilder:validation:Optional
	MultiLineStartRegexp *string `json:"multiLineStartRegexp,omitempty" tf:"multi_line_start_regexp,omitempty"`

	// +kubebuilder:validation:Optional
	NullValuePattern *string `json:"nullValuePattern,omitempty" tf:"null_value_pattern,omitempty"`

	// +kubebuilder:validation:Required
	ParserType *string `json:"parserType" tf:"parser_type,omitempty"`

	// +kubebuilder:validation:Optional
	Patterns []PatternsParameters `json:"patterns,omitempty" tf:"patterns,omitempty"`

	// +kubebuilder:validation:Optional
	Rfc5424TimeFormat *string `json:"rfc5424timeFormat,omitempty" tf:"rfc5424time_format,omitempty"`

	// +kubebuilder:validation:Optional
	SyslogParserType *string `json:"syslogParserType,omitempty" tf:"syslog_parser_type,omitempty"`

	// +kubebuilder:validation:Optional
	TimeFormat *string `json:"timeFormat,omitempty" tf:"time_format,omitempty"`

	// +kubebuilder:validation:Optional
	TimeType *string `json:"timeType,omitempty" tf:"time_type,omitempty"`

	// +kubebuilder:validation:Optional
	TimeoutInMilliseconds *float64 `json:"timeoutInMilliseconds,omitempty" tf:"timeout_in_milliseconds,omitempty"`

	// +kubebuilder:validation:Optional
	Types map[string]*string `json:"types,omitempty" tf:"types,omitempty"`
}

type PatternsObservation struct {
}

type PatternsParameters struct {

	// +kubebuilder:validation:Optional
	FieldTimeFormat *string `json:"fieldTimeFormat,omitempty" tf:"field_time_format,omitempty"`

	// +kubebuilder:validation:Optional
	FieldTimeKey *string `json:"fieldTimeKey,omitempty" tf:"field_time_key,omitempty"`

	// +kubebuilder:validation:Optional
	FieldTimeZone *string `json:"fieldTimeZone,omitempty" tf:"field_time_zone,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type ServiceConfigurationObservation struct {
}

type ServiceConfigurationParameters struct {

	// +kubebuilder:validation:Required
	ConfigurationType *string `json:"configurationType" tf:"configuration_type,omitempty"`

	// +kubebuilder:validation:Required
	Destination []DestinationParameters `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Required
	Sources []SourcesParameters `json:"sources" tf:"sources,omitempty"`
}

type SourcesObservation struct {
}

type SourcesParameters struct {

	// +kubebuilder:validation:Optional
	Channels []*string `json:"channels,omitempty" tf:"channels,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parser []ParserParameters `json:"parser,omitempty" tf:"parser,omitempty"`

	// +kubebuilder:validation:Optional
	Paths []*string `json:"paths,omitempty" tf:"paths,omitempty"`

	// +kubebuilder:validation:Required
	SourceType *string `json:"sourceType" tf:"source_type,omitempty"`
}

type UnifiedAgentConfigurationObservation struct {
	ConfigurationState *string `json:"configurationState,omitempty" tf:"configuration_state,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	TimeLastModified *string `json:"timeLastModified,omitempty" tf:"time_last_modified,omitempty"`
}

type UnifiedAgentConfigurationParameters struct {

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-oci/apis/identity/v1alpha1.Compartment
	// +kubebuilder:validation:Optional
	CompartmentID *string `json:"compartmentId,omitempty" tf:"compartment_id,omitempty"`

	// +kubebuilder:validation:Optional
	CompartmentIDRef *v1.Reference `json:"compartmentIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CompartmentIDSelector *v1.Selector `json:"compartmentIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// +kubebuilder:validation:Optional
	GroupAssociation []GroupAssociationParameters `json:"groupAssociation,omitempty" tf:"group_association,omitempty"`

	// +kubebuilder:validation:Required
	IsEnabled *bool `json:"isEnabled" tf:"is_enabled,omitempty"`

	// +kubebuilder:validation:Required
	ServiceConfiguration []ServiceConfigurationParameters `json:"serviceConfiguration" tf:"service_configuration,omitempty"`
}

// UnifiedAgentConfigurationSpec defines the desired state of UnifiedAgentConfiguration
type UnifiedAgentConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UnifiedAgentConfigurationParameters `json:"forProvider"`
}

// UnifiedAgentConfigurationStatus defines the observed state of UnifiedAgentConfiguration.
type UnifiedAgentConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UnifiedAgentConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UnifiedAgentConfiguration is the Schema for the UnifiedAgentConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ocijet}
type UnifiedAgentConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UnifiedAgentConfigurationSpec   `json:"spec"`
	Status            UnifiedAgentConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UnifiedAgentConfigurationList contains a list of UnifiedAgentConfigurations
type UnifiedAgentConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UnifiedAgentConfiguration `json:"items"`
}

// Repository type metadata.
var (
	UnifiedAgentConfiguration_Kind             = "UnifiedAgentConfiguration"
	UnifiedAgentConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UnifiedAgentConfiguration_Kind}.String()
	UnifiedAgentConfiguration_KindAPIVersion   = UnifiedAgentConfiguration_Kind + "." + CRDGroupVersion.String()
	UnifiedAgentConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(UnifiedAgentConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&UnifiedAgentConfiguration{}, &UnifiedAgentConfigurationList{})
}