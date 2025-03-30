// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package alert

import (
	"context"

	"github.com/bborbe/k8s"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/bborbe/alert/k8s/apis/monitoring.benjamin-borbe.de/v1"
)

func NewAlertResourceEventHandler(
	ctx context.Context,
	eventHandlerAlert k8s.EventHandler[v1.Alert],
) cache.ResourceEventHandler {
	return k8s.NewResourceEventHandler[v1.Alert](
		ctx,
		eventHandlerAlert,
	)
}
