// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package alert

import (
	"context"

	"github.com/bborbe/errors"
	"github.com/bborbe/k8s"
	"github.com/golang/glog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "github.com/bborbe/alert/k8s/apis/monitoring.benjamin-borbe.de/v1"
	"github.com/bborbe/alert/k8s/client/clientset/versioned"
)

//counterfeiter:generate -o mocks/alert-deployer.go --fake-name AlertDeployer . AlertDeployer
type AlertDeployer interface {
	Deploy(ctx context.Context, alert v1.Alert) error
	Undeploy(ctx context.Context, namespace k8s.Namespace, name string) error
}

func NewAlertDeployer(
	clientset *versioned.Clientset,
) AlertDeployer {
	return &alertDeployer{
		clientset: clientset,
	}
}

type alertDeployer struct {
	clientset *versioned.Clientset
}

func (a *alertDeployer) Deploy(ctx context.Context, alert v1.Alert) error {
	currentAlert, err := a.clientset.MonitoringV1().Alerts(alert.Namespace).Get(ctx, alert.Name, metav1.GetOptions{})
	if err != nil {
		_, err = a.clientset.MonitoringV1().Alerts(alert.Namespace).Create(ctx, &alert, metav1.CreateOptions{})
		if err != nil {
			return errors.Wrap(ctx, err, "create alert failed")
		}
		glog.V(3).Infof("alert %s created successful", alert.Name)
		return nil
	}
	updateAlert := mergeAlert(*currentAlert, alert)
	_, err = a.clientset.MonitoringV1().Alerts(alert.Namespace).Update(ctx, &updateAlert, metav1.UpdateOptions{})
	if err != nil {
		return errors.Wrap(ctx, err, "update alert failed")
	}
	glog.V(3).Infof("update alert %s successful", alert.Name)
	return nil
}

func (a *alertDeployer) Undeploy(ctx context.Context, namespace k8s.Namespace, name string) error {
	_, err := a.clientset.MonitoringV1().Alerts(namespace.String()).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		glog.V(4).Infof("alert '%s' not found => skip", name)
		return nil
	}
	if err := a.clientset.MonitoringV1().Alerts(namespace.String()).Delete(ctx, name, metav1.DeleteOptions{}); err != nil {
		return err
	}
	glog.V(3).Infof("delete alert %s successful", name)
	return nil
}

func mergeAlert(current, new v1.Alert) v1.Alert {
	new.ObjectMeta.ResourceVersion = current.ObjectMeta.ResourceVersion
	return new
}
