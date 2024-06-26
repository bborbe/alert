// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1 "github.com/bborbe/alert/k8s/apis/monitoring.benjamin-borbe.de/v1"
	monitoringbenjaminborbedev1 "github.com/bborbe/alert/k8s/client/applyconfiguration/monitoring.benjamin-borbe.de/v1"
)

// FakeAlerts implements AlertInterface
type FakeAlerts struct {
	Fake *FakeMonitoringV1
	ns   string
}

var alertsResource = v1.SchemeGroupVersion.WithResource("alerts")

var alertsKind = v1.SchemeGroupVersion.WithKind("Alert")

// Get takes name of the alert, and returns the corresponding alert object, and an error if there is any.
func (c *FakeAlerts) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Alert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alertsResource, c.ns, name), &v1.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Alert), err
}

// List takes label and field selectors, and returns the list of Alerts that match those selectors.
func (c *FakeAlerts) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AlertList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alertsResource, alertsKind, c.ns, opts), &v1.AlertList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.AlertList{ListMeta: obj.(*v1.AlertList).ListMeta}
	for _, item := range obj.(*v1.AlertList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alerts.
func (c *FakeAlerts) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alertsResource, c.ns, opts))

}

// Create takes the representation of a alert and creates it.  Returns the server's representation of the alert, and an error, if there is any.
func (c *FakeAlerts) Create(ctx context.Context, alert *v1.Alert, opts metav1.CreateOptions) (result *v1.Alert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alertsResource, c.ns, alert), &v1.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Alert), err
}

// Update takes the representation of a alert and updates it. Returns the server's representation of the alert, and an error, if there is any.
func (c *FakeAlerts) Update(ctx context.Context, alert *v1.Alert, opts metav1.UpdateOptions) (result *v1.Alert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alertsResource, c.ns, alert), &v1.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Alert), err
}

// Delete takes name of the alert and deletes it. Returns an error if one occurs.
func (c *FakeAlerts) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(alertsResource, c.ns, name, opts), &v1.Alert{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlerts) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alertsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.AlertList{})
	return err
}

// Patch applies the patch and returns the patched alert.
func (c *FakeAlerts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Alert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alertsResource, c.ns, name, pt, data, subresources...), &v1.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Alert), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied alert.
func (c *FakeAlerts) Apply(ctx context.Context, alert *monitoringbenjaminborbedev1.AlertApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Alert, err error) {
	if alert == nil {
		return nil, fmt.Errorf("alert provided to Apply must not be nil")
	}
	data, err := json.Marshal(alert)
	if err != nil {
		return nil, err
	}
	name := alert.Name
	if name == nil {
		return nil, fmt.Errorf("alert.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alertsResource, c.ns, *name, types.ApplyPatchType, data), &v1.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Alert), err
}
