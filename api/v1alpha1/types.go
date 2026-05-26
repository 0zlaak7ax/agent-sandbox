/*
Copyright 2024 The agent-sandbox Authors.

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

// Package v1alpha1 contains API Schema definitions for the agent-sandbox v1alpha1 API group.
// +groupName=sandbox.k8s.io
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AgentSandboxSpec defines the desired state of AgentSandbox.
type AgentSandboxSpec struct {
	// Image is the container image to run as the sandbox agent.
	// +kubebuilder:validation:Required
	Image string `json:"image"`

	// RuntimeClassName optionally specifies the RuntimeClass to use for the sandbox pod.
	// If not set, the cluster default is used.
	// +optional
	RuntimeClassName *string `json:"runtimeClassName,omitempty"`

	// Resources describes the compute resource requirements for the sandbox container.
	// +optional
	Resources ResourceRequirements `json:"resources,omitempty"`

	// Timeout is the maximum duration the sandbox is allowed to run.
	// Defaults to 30 minutes if not specified (increased from upstream default of 10m
	// to better accommodate longer-running agent tasks in my use case).
	// +optional
	Timeout *metav1.Duration `json:"timeout,omitempty"`

	// Env is a list of environment variables to set in the sandbox container.
	// +optional
	Env []EnvVar `json:"env,omitempty"`

	// ServiceAccountName is the name of the ServiceAccount to use for the sandbox pod.
	// Useful when the agent needs to interact with the Kubernetes API.
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
}

// ResourceRequirements describes the compute resource requirements.
type ResourceRequirements struct {
	// CPU is the CPU resource limit (e.g. "500m", "2").
	// +optional
	CPU string `json:"cpu,omitempty"`

	// Memory is the memory resource limit (e.g. "128Mi", "1Gi").
	// +optional
	Memory string `json:"memory,omitempty"`

	// EphemeralStorage is the ephemeral storage limit (e.g. "1Gi", "10Gi").
	// Useful for agent tasks that generate large intermediate files.
	// +optional
	EphemeralStorage string `json:"ephemeralStorage,omitempty"`
}

// EnvVar represents an environment variable present in the sandbox container.
type EnvVar struct {
	// Name of the environment variable.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Value of the environment variable.
	// +optional
	Value string `json:"value,omitempty"`

	// SecretRef optionally references a Kubernetes Secret key to populate this variable.
	// When set, Value is ignored. Handy for injecting API keys without hardcoding them.
	// +optional
	SecretRef *SecretKeyRef `json:"secretRef,omitempty"`
}

// SecretKeyRef selects a key from a Kubernetes Secret.
type SecretKeyRef struct {
	// Name is the name of the Secret.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Key is the key within the Secret to use.
	// +kubebuilder:validation:Required
	Key string `json:"key"`
}

// AgentSandboxPhase is a label for the condition of an AgentSandbox at the current time.
// +kubebuilder:validation:Enum=Pending;Running
