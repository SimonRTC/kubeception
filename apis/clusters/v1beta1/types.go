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

	// Current status of a cluster.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Status ClusterStatus `json:"status" protobuf:"bytes,3,opt,name=status"`
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

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.0.0

type HelmTemplate struct {
	metav1.TypeMeta `json:",inline"`

	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Containers image registry (must be accessible from both management and consumer cluster)
	// +optional
	Registry string `json:"registry,omitempty" protobuf:"bytes,2,opt,name=registry"`

	// Kubernetes control plane template (kube-apiserver, kube-controller-manager, ...)
	// +optional
	ControlPlan HelmControlPlaneTemplate `json:"controlPlan,omitempty" protobuf:"bytes,3,opt,name=controlPlan"`

	// Kubernetes consumer cluster template (coredns, cni, csi, ...)
	// +optional
	Consumer HelmConsumerTemplate `json:"consumer,omitempty" protobuf:"bytes,4,opt,name=consumer"`
}

type HelmControlPlaneTemplate struct {

	// The address/hostname on which to advertise the kube-apiserver of the consumer cluster (must be accessible from both management and consumer cluster)
	Advertise string `json:"advertise,omitempty" protobuf:"bytes,1,opt,name=advertise"`

	// Open ID Connect (Single Sign-On)
	// +optional
	OpenID OpenID `json:"oidc,omitempty" protobuf:"bytes,2,opt,name=oidc"`
}

type OpenID struct {
	// The URL of the OpenID issuer, only HTTPS scheme will be accepted. If set, it will be used to verify the OIDC JSON Web Token (JWT).
	Issuer string `json:"issuer,omitempty" protobuf:"bytes,1,opt,name=issuer"`

	// Comma-separated list of allowed JOSE asymmetric signing algorithms. JWTs with a supported 'alg' header values are: RS256, RS384, RS512, ES256, ES384, ES512, PS256, PS384, PS512. Values are defined by RFC 7518 https://tools.ietf.org/html/rfc7518#section-3.1.
	Algs string `json:"algs,omitempty" protobuf:"bytes,2,opt,name=algs"`

	// Name of the secret projected inside of the management cluster (automatically created when OpenID is configured)
	SecretName string `json:"secretName,omitempty" protobuf:"bytes,3,opt,name=secretName"`

	// The OpenID claim to use as the user name. Note that claims other than the default ('sub') is not guaranteed to be unique and immutable.
	Username OpenIDPrefix `json:"username,omitempty" protobuf:"bytes,4,opt,name=username"`

	// If provided, the name of a custom OpenID Connect claim for specifying user groups.
	Groups OpenIDPrefix `json:"groups,omitempty" protobuf:"bytes,5,opt,name=groups"`
}

type OpenIDPrefix struct {
	// The OpenID claim to use
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`

	// Prefix is use to avoid collision with existing internal role bindings (group `example` become `oidc:example`).
	Prefix string `json:"prefix,omitempty" protobuf:"bytes,2,opt,name=prefix"`
}

type HelmConsumerTemplate struct {
	// Must be implemented soon
}
