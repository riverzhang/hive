package clusterdeployment

import (
	hivev1 "github.com/openshift/hive/pkg/apis/hive/v1"
	"github.com/openshift/hive/pkg/test/generic"
)

// Option defines a function signature for any function that wants to be passed into Build
type Option func(*hivev1.ClusterDeployment)

// Build runs each of the functions passed in to generate the object.
func Build(opts ...Option) *hivev1.ClusterDeployment {
	retval := &hivev1.ClusterDeployment{}
	for _, o := range opts {
		o(retval)
	}

	return retval
}

// Generic allows common functions applicable to all objects to be used as Options to Build
func Generic(opt generic.Option) Option {
	return func(clusterDeployment *hivev1.ClusterDeployment) {
		opt(clusterDeployment)
	}
}

// WithCondition adds the specified condition to the ClusterDeployment
func WithCondition(cond hivev1.ClusterDeploymentCondition) Option {
	return func(clusterDeployment *hivev1.ClusterDeployment) {
		for i, c := range clusterDeployment.Status.Conditions {
			if c.Type == cond.Type {
				clusterDeployment.Status.Conditions[i] = cond
				return
			}
		}
		clusterDeployment.Status.Conditions = append(clusterDeployment.Status.Conditions, cond)
	}
}
