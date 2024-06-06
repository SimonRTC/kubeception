package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.0.0

type Cluster struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of a cluster.
	// +optional
	Spec ClusterSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.0.0

type ClusterList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of clusters.
	Items []Cluster `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type ClusterSpec struct {
	// Version of the cluster.
	// +optional
	Version string `json:"version" protobuf:"bytes,1,opt,name=version"`

	// Current status of a cluster.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Status ClusterStatus `json:"status" protobuf:"bytes,2,opt,name=status"`
}

// ClusterStatus represents the current state of a cluster.
type ClusterStatus struct {
	// The latest available observations of an object's current state. When a Cluster
	// fails, one of the conditions will have type "Failed" and status true. When
	// a Cluster is suspended, one of the conditions will have type "Suspended" and
	// status true; when the Cluster is resumed, the status of this condition will
	// become false.
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=atomic
	Conditions []ClusterCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

type ClusterConditionType string

// These are built-in conditions of a cluster.
const (
	// ClusterActive means the cluster is active
	ClusterActive ClusterConditionType = "Active"

	// ClusterCreating means the cluster is being created
	ClusterCreating ClusterConditionType = "Creating"

	// ClusterUpdating means the cluster is being updated
	ClusterUpdating ClusterConditionType = "Updating"

	// ClusterDeleting means the cluster is being deleted
	ClusterDeleting ClusterConditionType = "Deleting"

	// ClusterSuspending means the cluster is being suspended
	ClusterSuspending ClusterConditionType = "Suspending"

	// ClusterSuspended means the cluster is suspended
	ClusterSuspended ClusterConditionType = "Suspended"

	// ClusterError means the cluster is encounter an issue
	ClusterError ClusterConditionType = "Error"
)

// ClusterCondition describes current state of a cluster.
type ClusterCondition struct {
	// Type of cluster condition, Complete or Failed.
	Type ClusterConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=ClusterConditionType"`

	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=k8s.io/api/core/v1.ConditionStatus"`

	// Last time the condition was checked.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty" protobuf:"bytes,3,opt,name=lastProbeTime"`

	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`

	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`

	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}
