/*
Copyright 2022 The Kubeflow Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/kubeflow/katib/pkg/client/controller/applyconfiguration/common/v1beta1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// TrialSpecApplyConfiguration represents an declarative configuration of the TrialSpec type for use
// with apply.
type TrialSpecApplyConfiguration struct {
	Objective            *v1beta1.ObjectiveSpecApplyConfiguration        `json:"objective,omitempty"`
	ParameterAssignments []v1beta1.ParameterAssignmentApplyConfiguration `json:"parameterAssignments,omitempty"`
	EarlyStoppingRules   []v1beta1.EarlyStoppingRuleApplyConfiguration   `json:"earlyStoppingRules,omitempty"`
	RunSpec              *unstructured.Unstructured                      `json:"runSpec,omitempty"`
	RetainRun            *bool                                           `json:"retainRun,omitempty"`
	MetricsCollector     *v1beta1.MetricsCollectorSpecApplyConfiguration `json:"metricsCollector,omitempty"`
	PrimaryPodLabels     map[string]string                               `json:"primaryPodLabels,omitempty"`
	PrimaryContainerName *string                                         `json:"primaryContainerName,omitempty"`
	SuccessCondition     *string                                         `json:"successCondition,omitempty"`
	FailureCondition     *string                                         `json:"failureCondition,omitempty"`
	Labels               map[string]string                               `json:"labels,omitempty"`
}

// TrialSpecApplyConfiguration constructs an declarative configuration of the TrialSpec type for use with
// apply.
func TrialSpec() *TrialSpecApplyConfiguration {
	return &TrialSpecApplyConfiguration{}
}

// WithObjective sets the Objective field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Objective field is set to the value of the last call.
func (b *TrialSpecApplyConfiguration) WithObjective(value *v1beta1.ObjectiveSpecApplyConfiguration) *TrialSpecApplyConfiguration {
	b.Objective = value
	return b
}

// WithParameterAssignments adds the given value to the ParameterAssignments field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ParameterAssignments field.
func (b *TrialSpecApplyConfiguration) WithParameterAssignments(values ...*v1beta1.ParameterAssignmentApplyConfiguration) *TrialSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithParameterAssignments")
		}
		b.ParameterAssignments = append(b.ParameterAssignments, *values[i])
	}
	return b
}

// WithEarlyStoppingRules adds the given value to the EarlyStoppingRules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EarlyStoppingRules field.
func (b *TrialSpecApplyConfiguration) WithEarlyStoppingRules(values ...*v1beta1.EarlyStoppingRuleApplyConfiguration) *TrialSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithEarlyStoppingRules")
		}
		b.EarlyStoppingRules = append(b.EarlyStoppingRules, *values[i])
	}
	return b
}

// WithRunSpec sets the RunSpec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RunSpec field is set to the value of the last call.
func (b *TrialSpecApplyConfiguration) WithRunSpec(value unstructured.Unstructured) *TrialSpecApplyConfiguration {
	b.RunSpec = &value
	return b
}

// WithRetainRun sets the RetainRun field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RetainRun field is set to the value of the last call.
func (b *TrialSpecApplyConfiguration) WithRetainRun(value bool) *TrialSpecApplyConfiguration {
	b.RetainRun = &value
	return b
}

// WithMetricsCollector sets the MetricsCollector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MetricsCollector field is set to the value of the last call.
func (b *TrialSpecApplyConfiguration) WithMetricsCollector(value *v1beta1.MetricsCollectorSpecApplyConfiguration) *TrialSpecApplyConfiguration {
	b.MetricsCollector = value
	return b
}

// WithPrimaryPodLabels puts the entries into the PrimaryPodLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the PrimaryPodLabels field,
// overwriting an existing map entries in PrimaryPodLabels field with the same key.
func (b *TrialSpecApplyConfiguration) WithPrimaryPodLabels(entries map[string]string) *TrialSpecApplyConfiguration {
	if b.PrimaryPodLabels == nil && len(entries) > 0 {
		b.PrimaryPodLabels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.PrimaryPodLabels[k] = v
	}
	return b
}

// WithPrimaryContainerName sets the PrimaryContainerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PrimaryContainerName field is set to the value of the last call.
func (b *TrialSpecApplyConfiguration) WithPrimaryContainerName(value string) *TrialSpecApplyConfiguration {
	b.PrimaryContainerName = &value
	return b
}

// WithSuccessCondition sets the SuccessCondition field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SuccessCondition field is set to the value of the last call.
func (b *TrialSpecApplyConfiguration) WithSuccessCondition(value string) *TrialSpecApplyConfiguration {
	b.SuccessCondition = &value
	return b
}

// WithFailureCondition sets the FailureCondition field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailureCondition field is set to the value of the last call.
func (b *TrialSpecApplyConfiguration) WithFailureCondition(value string) *TrialSpecApplyConfiguration {
	b.FailureCondition = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *TrialSpecApplyConfiguration) WithLabels(entries map[string]string) *TrialSpecApplyConfiguration {
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}
