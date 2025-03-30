// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package alert

import (
	"github.com/bborbe/k8s"

	v1 "github.com/bborbe/alert/k8s/apis/monitoring.benjamin-borbe.de/v1"
)

type AlertEventHandler k8s.EventHandler[v1.Alert]

func NewAlertEventHandler() AlertEventHandler {
	return k8s.NewEventHandler[v1.Alert]()
}
