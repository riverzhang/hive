package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterPoolSpec defines the desired state of the ClusterPool.
type ClusterPoolSpec struct {

	// Platform encompasses the desired platform for the cluster.
	// +required
	Platform Platform `json:"platform"`

	// PullSecretRef is the reference to the secret to use when pulling images.
	// +optional
	PullSecretRef *corev1.LocalObjectReference `json:"pullSecretRef,omitempty"`

	// Size is the default number of clusters that we should keep provisioned and waiting for use.
	// +kubebuilder:validation:Minimum=0
	// +required
	Size int32 `json:"size"`

	// MaxSize is the maximum number of clusters that will be provisioned including clusters that have been claimed
	// and ones waiting to be used.
	// By default there is no limit.
	// +optional
	MaxSize *int32 `json:"maxSize,omitempty"`

	// MaxConcurrent is the maximum number of clusters that will be provisioned or deprovisioned at an time.
	// By default there is no limit.
	// +optional
	MaxConcurrent *int32 `json:"maxConcurrent,omitempty"`

	// BaseDomain is the base domain to use for all clusters created in this pool.
	// +required
	BaseDomain string `json:"baseDomain"`

	// ImageSetRef is a reference to a ClusterImageSet. The release image specified in the ClusterImageSet will be used
	// by clusters created for this cluster pool.
	ImageSetRef ClusterImageSetReference `json:"imageSetRef"`

	// Labels to be applied to new ClusterDeployments created for the pool. ClusterDeployments that have already been
	// claimed will not be affected when this value is modified.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// HibernateAfter will be applied to new ClusterDeployments created for the pool. HibernateAfter will transition
	// clusters in the clusterpool to hibernating power state after it has been running for the given duration. The time
	// that a cluster has been running is the time since the cluster was installed or the time since the cluster last came
	// out of hibernation.
	// +optional
	HibernateAfter *metav1.Duration `json:"hibernateAfter,omitempty"`
}

// ClusterPoolStatus defines the observed state of ClusterPool
type ClusterPoolStatus struct {
	// Size is the number of unclaimed clusters that have been created for the pool.
	Size int32 `json:"size"`

	// Ready is the number of unclaimed clusters that have been installed and are ready to be claimed.
	Ready int32 `json:"ready"`

	// Conditions includes more detailed status for the cluster pool
	// +optional
	Conditions []ClusterPoolCondition `json:"conditions,omitempty"`
}

// ClusterPoolCondition contains details for the current condition of a cluster pool
type ClusterPoolCondition struct {
	// Type is the type of the condition.
	Type ClusterPoolConditionType `json:"type"`
	// Status is the status of the condition.
	Status corev1.ConditionStatus `json:"status"`
	// LastProbeTime is the last time we probed the condition.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// LastTransitionTime is the last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// Reason is a unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Message is a human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// ClusterPoolConditionType is a valid value for ClusterPoolCondition.Type
type ClusterPoolConditionType string

const (
	// ClusterPoolMissingDependenciesCondition is set when a cluster pool is missing dependencies required to create a
	// cluster. Dependencies include resources such as the ClusterImageSet and the credentials Secret.
	ClusterPoolMissingDependenciesCondition ClusterPoolConditionType = "MissingDependencies"
	// ClusterPoolCapacityAvailableCondition is set to provide information on whether the cluster pool has capacity
	// available to create more clusters for the pool.
	ClusterPoolCapacityAvailableCondition ClusterPoolConditionType = "CapacityAvailable"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterPool represents a pool of clusters that should be kept ready to be given out to users. Clusters are removed
// from the pool once claimed and then automatically replaced with a new one.
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=.spec.size,statuspath=.status.size
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.ready"
// +kubebuilder:printcolumn:name="Size",type="string",JSONPath=".spec.size"
// +kubebuilder:printcolumn:name="BaseDomain",type="string",JSONPath=".spec.baseDomain"
// +kubebuilder:printcolumn:name="ImageSet",type="string",JSONPath=".spec.imageSetRef.name"
// +kubebuilder:resource:path=clusterpools,shortName=cp
type ClusterPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterPoolSpec   `json:"spec"`
	Status ClusterPoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterPoolList contains a list of ClusterPools
type ClusterPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterPool{}, &ClusterPoolList{})
}
