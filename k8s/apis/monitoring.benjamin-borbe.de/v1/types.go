// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"context"
	"reflect"

	"github.com/bborbe/errors"
	"github.com/bborbe/k8s"
	"github.com/bborbe/validation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Alerts []Alert

func (a Alerts) Contains(name string) bool {
	for _, aa := range a {
		if aa.Name == name {
			return true
		}
	}
	return false
}

func (a Alerts) Specs() AlertSpecs {
	var result AlertSpecs
	for _, aa := range a {
		result = append(result, aa.Spec)
	}
	return result
}

var _ k8s.Type = Alert{}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Alert describes a database.
type Alert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AlertSpec `json:"spec"`
}

func (a Alert) Equal(other k8s.Type) bool {
	switch alert := other.(type) {
	case Alert:
		return a.Spec.Equal(alert.Spec)
	case *Alert:
		return a.Spec.Equal(alert.Spec)
	}
	return false
}

func (a Alert) Validate(ctx context.Context) error {
	return a.Spec.Validate(ctx)
}

func (a Alert) Identifier() k8s.Identifier {
	return k8s.Identifier(k8s.BuildName(a.Namespace, a.Name))
}

func (a Alert) String() string {
	return a.Name
}

/*
Sample Alert:

alert: K8sPodStatus
annotations:
  description: 'Failed K8s Pods for more than 1h'
  summary: K8s Pod status
expr: sum(max_over_time(octopus_k8s_pod_status_counter{status="Waiting"}[600s])) by (name)  > 0
for: 1h
labels:
  severity: warning
*/

type AlertSpecs []AlertSpec

func (a AlertSpecs) Contains(name string) bool {
	for _, aa := range a {
		if aa.Name == name {
			return true
		}
	}
	return false
}

// AlertSpec is the spec for a Foo resource
type AlertSpec struct {
	Name        string            `json:"name"                  yaml:"name"`
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Expression  string            `json:"expression,omitempty"  yaml:"expression,omitempty"`
	For         string            `json:"for,omitempty"         yaml:"for,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"      yaml:"labels,omitempty"`
}

func (a AlertSpec) Equal(alert AlertSpec) bool {
	if a.Name != alert.Name {
		return false
	}
	if a.Expression != alert.Expression {
		return false
	}
	if a.For != alert.For {
		return false
	}
	if !reflect.DeepEqual(a.Annotations, alert.Annotations) {
		return false
	}
	if !reflect.DeepEqual(a.Labels, alert.Labels) {
		return false
	}
	return true
}

func (a AlertSpec) Validate(ctx context.Context) error {
	if a.Name == "" {
		return errors.Wrap(ctx, validation.Error, "Name is empty")
	}
	if a.Expression == "" {
		return errors.Wrap(ctx, validation.Error, "Expression is empty")
	}
	if len(a.Annotations) == 0 {
		return errors.Wrap(ctx, validation.Error, "Annotations is empty")
	}
	if len(a.Labels) == 0 {
		return errors.Wrap(ctx, validation.Error, "Labels is empty")
	}
	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlertList is a list of Alert resources
type AlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Alert `json:"items"`
}
