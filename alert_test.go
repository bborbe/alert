// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package alert_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	v1 "github.com/bborbe/alert/k8s/apis/monitoring.benjamin-borbe.de/v1"
)

const alertJSON = `
{
  "apiVersion": "monitoring.benjamin-borbe.de/v1",
  "kind": "Alert",
  "metadata": {
    "name": "my-alert",
    "namespace": "monitoring"
  },
  "spec": {
    "name": "K8sPodStatus",
    "annotations": {
      "description": "Failed K8s Pods for more than 1h",
      "summary": "K8s Pod status"
    },
    "expression": "sum(max_over_time(octopus_k8s_pod_status_counter{status=\"Waiting\"}[600s])) by (name) > 0",
    "for": "1h",
    "labels": {
      "severity": "warning"
    }
  }
}
`

var _ = Describe("Alert", func() {
	var alert v1.Alert
	var err error

	BeforeEach(func() {
		alert = v1.Alert{}
		err = json.Unmarshal([]byte(alertJSON), &alert)
	})
	It("returns no error", func() {
		Expect(err).To(BeNil())
	})
	It("contains name", func() {
		Expect(alert.Spec.Name).To(Equal("K8sPodStatus"))
	})
	It("contains annotations", func() {
		Expect(
			alert.Spec.Annotations,
		).To(HaveKeyWithValue("description", "Failed K8s Pods for more than 1h"))
		Expect(alert.Spec.Annotations).To(HaveKeyWithValue("summary", "K8s Pod status"))
	})
	It("contains expression", func() {
		Expect(
			alert.Spec.Expression,
		).To(Equal(`sum(max_over_time(octopus_k8s_pod_status_counter{status="Waiting"}[600s])) by (name) > 0`))
	})
	It("contains for", func() {
		Expect(alert.Spec.For).To(Equal("1h"))
	})
	It("contains labels", func() {
		Expect(alert.Spec.Labels).To(HaveKeyWithValue("severity", "warning"))
	})
})
