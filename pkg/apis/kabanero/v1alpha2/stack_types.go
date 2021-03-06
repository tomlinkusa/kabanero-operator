package v1alpha2

import (
	"strings"
	
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// NOTE: The +listType=set marker are required by OpenAPI generation for list types.

const (
	// StackDesiredStateActive represents a desired stack active state.
	// It indicates that the stack needs activation.
	StackDesiredStateActive = "active"

	// StackDesiredStateInactive represents a desired stack inactive state.
	// It indicates that the stack needs to be deactivated.
	StackDesiredStateInactive = "inactive"

	// StackStateError represents a stack status error state.
	// It indicates that the stack did not complete an activation process
	StackStateError = "error"

	// Stack digest policy: strict.
	StackPolicyStrictDigest = "strictDigest"

	// Stack digest policy: active.
	StackPolicyActiveDigest = "activeDigest"

	// Stack digest policy: ignore.
	StackPolicyIgnoreDigest = "ignoreDigest"

	// Stack digest policy: none.
	StackPolicyNone = "none"
)

// StackSpec defines the desired composition of a Stack
// +k8s:openapi-gen=true
type StackSpec struct {
	Name string `json:"name,omitempty"`
	// +listType=map
	// +listMapKey=version
	Versions []StackVersion `json:"versions,omitempty"`
}

func (s StackSpec) GetVersions() []ComponentSpecVersion {
	ret := make([]ComponentSpecVersion, len(s.Versions))
	for i, _ := range s.Versions {
		ret[i] = s.Versions[i]
	}
	return ret
}

// StackVersion defines the desired composition of a specific stack version.
type StackVersion struct {
	SkipRegistryCertVerification bool `json:"skipRegistryCertVerification,omitempty"`

	// +listType=map
	// +listMapKey=id
	// +listMapKey=sha256
	Pipelines            []PipelineSpec `json:"pipelines,omitempty"`
	Version              string         `json:"version,omitempty"`
	DesiredState         string         `json:"desiredState,omitempty"`
	SkipCertVerification bool           `json:"skipCertVerification,omitempty"`
	// +listType=map
	// +listMapKey=id
	// +listMapKey=image
	Images               []Image        `json:"images,omitempty"`
	Devfile              string         `json:"devfile,omitempty"`
	Metafile             string         `json:"metafile,omitempty"`
}

func (sv StackVersion) GetVersion() string {
	return sv.Version
}

func (sv StackVersion) GetPipelines() []PipelineSpec {
	// Only return pipelines if the version is active
	if !strings.EqualFold(sv.DesiredState, StackDesiredStateInactive) {
		return sv.Pipelines
	}

	return nil
}

// GitReleaseInfo is all of the GitReleaseSpec information, minus the "skip cert
// verification" information, which is not relevant for status.
type GitReleaseInfo struct {
	Hostname             string `json:"hostname,omitempty"`
	Organization         string `json:"organization,omitempty"`
	Project              string `json:"project,omitempty"`
	Release              string `json:"release,omitempty"`
	AssetName            string `json:"assetName,omitempty"`
}

// Returns true if the user specified all values for the release.
func (gitRelease GitReleaseInfo) IsUsable() bool {
	return len(gitRelease.Hostname) != 0 && len(gitRelease.Organization) != 0 && len(gitRelease.Project) != 0 &&
		len(gitRelease.Release) != 0 && len(gitRelease.AssetName) != 0
}


// RepositoryAssetStatus defines the observed state of a single asset in a pipelines respository.
type RepositoryAssetStatus struct {
	Name          string `json:"assetName,omitempty"`
	Namespace     string `json:"namespace,omitempty"`
	Group         string `json:"group,omitempty"`
	Version       string `json:"version,omitempty"`
	Kind          string `json:"kind,omitempty"`
	Digest        string `json:"assetDigest,omitempty"`
	Status        string `json:"status,omitempty"`
	StatusMessage string `json:"statusMessage,omitempty"`
}

// StackStatus defines the observed state of a stack
// +k8s:openapi-gen=true
type StackStatus struct {
	StatusMessage string `json:"statusMessage,omitempty"`
	// +listType=map
	// +listMapKey=version
	Versions []StackVersionStatus `json:"versions,omitempty"`
	Summary  string               `json:"summary,omitempty"`
}

func (s StackStatus) GetVersions() []ComponentStatusVersion {
	ret := make([]ComponentStatusVersion, len(s.Versions))
	for i, _ := range s.Versions {
		ret[i] = s.Versions[i]
	}
	return ret
}

// StackVersionStatus defines the observed state of a specific stack version.
type StackVersionStatus struct {
	Version  string `json:"version,omitempty"`
	Location string `json:"location,omitempty"`
	// +listType=map
	// +listMapKey=name
	// +listMapKey=digest
	Pipelines     []PipelineStatus `json:"pipelines,omitempty"`
	Status        string           `json:"status,omitempty"`
	StatusMessage string           `json:"statusMessage,omitempty"`
	// +listType=map
	// +listMapKey=id
	// +listMapKey=image
	Images []ImageStatus `json:"images,omitempty"`
}

func (sv StackVersionStatus) GetVersion() string {
	return sv.Version
}

func (sv StackVersionStatus) GetPipelines() []PipelineStatus {
	return sv.Pipelines
}

// Image defines a container image used by a stack
type Image struct {
	Id    string `json:"id,omitempty"`
	Image string `json:"image,omitempty"`
}

// ImageStatus defines a container image status used by a stack
type ImageStatus struct {
	Id     string      `json:"id,omitempty"`
	Image  string      `json:"image,omitempty"`
	Digest ImageDigest `json:"digest,omitempty"`
}

// ImageDigest defines a container image digest used by a stack
type ImageDigest struct {
	Activation string `json:"activation,omitempty"`
	Message    string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Stack is the Schema for the stack API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations."
// +kubebuilder:printcolumn:name="Summary",type="string",JSONPath=".status.summary",description="Stack summary."
// +kubebuilder:resource:path=stacks,scope=Namespaced
type Stack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StackSpec   `json:"spec,omitempty"`
	Status StackStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StackList contains a list of Stacks
type StackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// +listType=set
	Items []Stack `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Stack{}, &StackList{})
}
