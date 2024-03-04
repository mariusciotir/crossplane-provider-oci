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

type StorageSnapshotObservation struct {

	// The OCID of the file system snapshot policy that created this snapshot.
	FilesystemSnapshotPolicyID *string `json:"filesystemSnapshotPolicyId,omitempty" tf:"filesystem_snapshot_policy_id,omitempty"`

	// The OCID of the snapshot.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether the snapshot has been cloned. See Cloning a File System.
	IsCloneSource *bool `json:"isCloneSource,omitempty" tf:"is_clone_source,omitempty"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details,omitempty"`

	// An OCID identifying the parent from which this snapshot was cloned. If this snapshot was not cloned, then the provenanceId is the same as the snapshot id value. If this snapshot was cloned, then the provenanceId value is the parent's provenanceId. See Cloning a File System.
	ProvenanceID *string `json:"provenanceId,omitempty" tf:"provenance_id,omitempty"`

	// The date and time the snapshot was taken, expressed in RFC 3339 timestamp format. This value might be the same or different from timeCreated depending on the following factors:
	SnapshotTime *string `json:"snapshotTime,omitempty" tf:"snapshot_time,omitempty"`

	// Specifies the generation type of the snapshot.
	SnapshotType *string `json:"snapshotType,omitempty" tf:"snapshot_type,omitempty"`

	// The current state of the snapshot.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The date and time the snapshot was created, expressed in RFC 3339 timestamp format.  Example: 2016-08-25T21:10:29.600Z
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`
}

type StorageSnapshotParameters struct {

	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: {"Operations.CostCenter": "42"}
	// +kubebuilder:validation:Optional
	DefinedTags map[string]*string `json:"definedTags,omitempty" tf:"defined_tags,omitempty"`

	// (Updatable) The time when this snapshot will be deleted.
	// +kubebuilder:validation:Optional
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// The OCID of the file system to take a snapshot of.
	// +crossplane:generate:reference:type=StorageFileSystem
	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// Reference to a StorageFileSystem to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDRef *v1.Reference `json:"fileSystemIdRef,omitempty" tf:"-"`

	// Selector for a StorageFileSystem to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDSelector *v1.Selector `json:"fileSystemIdSelector,omitempty" tf:"-"`

	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags. Example: {"Department": "Finance"}
	// +kubebuilder:validation:Optional
	FreeformTags map[string]*string `json:"freeformTags,omitempty" tf:"freeform_tags,omitempty"`

	// Name of the snapshot. This value is immutable. It must also be unique with respect to all other non-DELETED snapshots on the associated file system.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// StorageSnapshotSpec defines the desired state of StorageSnapshot
type StorageSnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageSnapshotParameters `json:"forProvider"`
}

// StorageSnapshotStatus defines the observed state of StorageSnapshot.
type StorageSnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageSnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StorageSnapshot is the Schema for the StorageSnapshots API. Provides the Snapshot resource in Oracle Cloud Infrastructure File Storage service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,oci}
type StorageSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageSnapshotSpec   `json:"spec"`
	Status            StorageSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageSnapshotList contains a list of StorageSnapshots
type StorageSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageSnapshot `json:"items"`
}

// Repository type metadata.
var (
	StorageSnapshot_Kind             = "StorageSnapshot"
	StorageSnapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StorageSnapshot_Kind}.String()
	StorageSnapshot_KindAPIVersion   = StorageSnapshot_Kind + "." + CRDGroupVersion.String()
	StorageSnapshot_GroupVersionKind = CRDGroupVersion.WithKind(StorageSnapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&StorageSnapshot{}, &StorageSnapshotList{})
}
