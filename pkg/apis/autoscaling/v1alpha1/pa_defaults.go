/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"

	"github.com/knative/pkg/apis"
	"github.com/knative/serving/pkg/apis/autoscaling"
)

func (r *PodAutoscaler) SetDefaults(ctx context.Context) {
	r.Spec.SetDefaults(apis.WithinSpec(ctx))
	if r.Annotations == nil {
		r.Annotations = make(map[string]string)
	}
	if _, ok := r.Annotations[autoscaling.ClassAnnotationKey]; !ok {
		// Default class to KPA.
		r.Annotations[autoscaling.ClassAnnotationKey] = autoscaling.KPA
	}
	// Default metric per class
	switch r.Class() {
	case autoscaling.KPA:
		if _, ok := r.Annotations[autoscaling.MetricAnnotationKey]; !ok {
			r.Annotations[autoscaling.MetricAnnotationKey] = autoscaling.Concurrency
		}
	case autoscaling.HPA:
		if _, ok := r.Annotations[autoscaling.MetricAnnotationKey]; !ok {
			r.Annotations[autoscaling.MetricAnnotationKey] = autoscaling.CPU
		}
	}
}

func (rs *PodAutoscalerSpec) SetDefaults(ctx context.Context) {}
