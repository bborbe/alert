// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1_test

import (
	"context"

	"github.com/bborbe/k8s"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "github.com/bborbe/alert/k8s/apis/monitoring.benjamin-borbe.de/v1"
)

var _ = Describe("Alert", func() {
	var ctx context.Context
	var err error
	var alertA *v1.Alert
	var alertB *v1.Alert
	BeforeEach(func() {
		ctx = context.Background()
		alertA = &v1.Alert{
			ObjectMeta: apismetav1.ObjectMeta{
				Name:      "alert-a",
				Namespace: "alertA",
			},
			Spec: v1.AlertSpec{
				Name:       "MyAlertA",
				Expression: "bar > 0",
				Annotations: map[string]string{
					"foo": "bar",
				},
				Labels: map[string]string{
					"foo": "bar",
				},
			},
		}
		alertB = &v1.Alert{
			ObjectMeta: apismetav1.ObjectMeta{
				Name:      "alert-b",
				Namespace: "alertB",
			},
			Spec: v1.AlertSpec{
				Name:       "MyAlertB",
				Expression: "foo > 0",
				Annotations: map[string]string{
					"foo": "bar",
				},
				Labels: map[string]string{
					"foo": "bar",
				},
			},
		}
	})
	Context("Equal", func() {
		var result bool
		JustBeforeEach(func() {
			result = alertA.Equal(alertB)
		})
		Context("not equal", func() {
			It("returns false", func() {
				Expect(result).To(BeFalse())
			})
		})
		Context("equal", func() {
			BeforeEach(func() {
				alertA = alertB
			})
			It("returns true", func() {
				Expect(result).To(BeTrue())
			})
		})
	})
	Context("Validate", func() {
		JustBeforeEach(func() {
			err = alertA.Validate(ctx)
		})
		It("is valid", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("Identifier", func() {
		var id k8s.Identifier
		JustBeforeEach(func() {
			id = alertA.Identifier()
		})
		It("is valid", func() {
			Expect(id).To(Equal(k8s.Identifier("alert-a")))
		})
	})
	Context("String", func() {
		var str string
		JustBeforeEach(func() {
			str = alertA.String()
		})
		It("is valid", func() {
			Expect(str).To(Equal("alert-a"))
		})
	})
})
