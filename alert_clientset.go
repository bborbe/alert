// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package alert

import (
	"github.com/golang/glog"
	"github.com/pkg/errors"
	k8s_rest "k8s.io/client-go/rest"
	k8s_clientcmd "k8s.io/client-go/tools/clientcmd"

	"github.com/bborbe/alert/k8s/client/clientset/versioned"
)

func CreateClientset(kubeconfig string) (versioned.Interface, error) {
	config, err := CreateConfig(kubeconfig)
	if err != nil {
		return nil, errors.Wrap(err, "create k8s config failed")
	}
	return versioned.NewForConfig(config)
}

func CreateConfig(kubeconfig string) (*k8s_rest.Config, error) {
	if len(kubeconfig) > 0 {
		glog.V(4).Infof("create kube config from flags")
		return k8s_clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	glog.V(4).Infof("create in cluster kube config")
	return k8s_rest.InClusterConfig()
}
