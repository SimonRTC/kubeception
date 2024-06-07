package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.0.0

type StorageBackend struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of a StorageBackend.
	// +optional
	Spec StorageBackendSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Current status of a StorageBackend.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Status StorageBackendStatus `json:"status" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.0.0

type StorageBackendList struct {
	metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of StorageBackend.
	Items []StorageBackend `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type StorageBackendSpec struct {
	// Certificates used to secure etcd communication.
	// +optional
	Endpoints []string `json:"endpoints" protobuf:"bytes,1,opt,name=endpoints"`

	// Certificates used to secure etcd communication.
	// +optional
	Certificates StorageBackendCertificates `json:"certificates" protobuf:"bytes,2,opt,name=certificates"`
}

type StorageBackendCertificates struct {
	// Certificate Authority used to secure etcd communication.
	CA string `json:"ca" protobuf:"bytes,1,opt,name=ca"`

	// Certification file used to secure etcd communication.
	Cert string `json:"cert" protobuf:"bytes,2,opt,name=cert"`

	// Private key used to secure etcd communication.
	Key string `json:"key" protobuf:"bytes,3,opt,name=key"`
}

// StorageBackendStatus represents the current state of a storage backend.
type StorageBackendStatus struct {
	// The latest available observations of an object's current state. When a StorageBackend
	// fails, one of the conditions will have type "Failed" and status true. When
	// a StorageBackend is unhealthy, one of the conditions will have type "Healthy" and
	// status false; when the StorageBackend is resumed, the status of this condition will
	// become false.
	//
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=atomic
	Conditions []StorageBackendCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

type StorageBackendConditionType string

// These are built-in conditions of a StorageBackend.
const (
	// StorageBackendActive means the etcd is active (only affect new clusters)
	StorageBackendActive StorageBackendConditionType = "Active"

	// StorageBackendHealthy means the etcd is ready and healthy
	StorageBackendHealthy StorageBackendConditionType = "Healthy"

	// StorageBackendError means the backend is encounter an issue
	StorageBackendError StorageBackendConditionType = "Error"
)

// StorageBackendCondition describes current state of a backend.
type StorageBackendCondition struct {
	// Type of backend condition, Complete or Failed.
	Type StorageBackendConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=StorageBackendConditionType"`

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
