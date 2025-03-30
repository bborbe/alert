// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "github.com/bborbe/alert/k8s/apis/monitoring.benjamin-borbe.de/v1"
	internal "github.com/bborbe/alert/k8s/client/applyconfiguration/internal"
	monitoringbenjaminborbedev1 "github.com/bborbe/alert/k8s/client/applyconfiguration/monitoring.benjamin-borbe.de/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=monitoring, Version=v1
	case v1.SchemeGroupVersion.WithKind("Alert"):
		return &monitoringbenjaminborbedev1.AlertApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AlertSpec"):
		return &monitoringbenjaminborbedev1.AlertSpecApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
