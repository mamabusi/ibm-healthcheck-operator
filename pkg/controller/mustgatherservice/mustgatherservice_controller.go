//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package mustgatherservice

import (
	"context"

	operatorv1alpha1 "github.com/IBM/ibm-healthcheck-operator/pkg/apis/operator/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_mustgatherservice")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new MustGatherService Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMustGatherService{client: mgr.GetClient(), reader: mgr.GetAPIReader(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("mustgatherservice-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource MustGatherService
	err = c.Watch(&source.Kind{Type: &operatorv1alpha1.MustGatherService{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Pods and requeue the owner HealthService
	// StatefulSet
	err = c.Watch(&source.Kind{Type: &appsv1.StatefulSet{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &operatorv1alpha1.MustGatherService{},
	})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner MustGatherService
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &operatorv1alpha1.MustGatherService{},
	})
	if err != nil {
		return err
	}

	// Ingress
	err = c.Watch(&source.Kind{Type: &networkingv1.Ingress{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &operatorv1alpha1.MustGatherService{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileMustGatherService implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileMustGatherService{}

// ReconcileMustGatherService reconciles a MustGatherService object
type ReconcileMustGatherService struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	reader client.Reader
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a MustGatherService object and makes changes based on the state read
// and what is in the MustGatherService.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMustGatherService) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling MustGatherService")

	// Fetch the MustGatherService instance
	instance := &operatorv1alpha1.MustGatherService{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	if err = r.createOrUpdateMustGatherServicePVC(instance); err != nil {
		reqLogger.Error(err, "Failed to create or update PVC for mustgather service")
		return reconcile.Result{}, err
	}

	if err = r.createOrUpdateMustGatherServiceStatefulSet(instance); err != nil {
		reqLogger.Error(err, "Failed to create or update StatefulSet for mustgather service")
		return reconcile.Result{}, err
	}

	if err = r.createOrUpdateMustGatherServiceService(instance); err != nil {
		reqLogger.Error(err, "Failed to create or update Service for mustgather service")
		return reconcile.Result{}, err
	}

	if err = r.createOrUpdateMustGatherServiceIngress(instance); err != nil {
		reqLogger.Error(err, "Failed to create or update Ingress for mustgather service")
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}
