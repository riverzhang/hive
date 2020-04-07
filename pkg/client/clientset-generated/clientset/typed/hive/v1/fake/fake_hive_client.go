// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/hive/pkg/client/clientset-generated/clientset/typed/hive/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHiveV1 struct {
	*testing.Fake
}

func (c *FakeHiveV1) Checkpoints(namespace string) v1.CheckpointInterface {
	return &FakeCheckpoints{c, namespace}
}

func (c *FakeHiveV1) ClusterDeployments(namespace string) v1.ClusterDeploymentInterface {
	return &FakeClusterDeployments{c, namespace}
}

func (c *FakeHiveV1) ClusterDeprovisions(namespace string) v1.ClusterDeprovisionInterface {
	return &FakeClusterDeprovisions{c, namespace}
}

func (c *FakeHiveV1) ClusterImageSets() v1.ClusterImageSetInterface {
	return &FakeClusterImageSets{c}
}

func (c *FakeHiveV1) ClusterProvisions(namespace string) v1.ClusterProvisionInterface {
	return &FakeClusterProvisions{c, namespace}
}

func (c *FakeHiveV1) ClusterStates(namespace string) v1.ClusterStateInterface {
	return &FakeClusterStates{c, namespace}
}

func (c *FakeHiveV1) DNSZones(namespace string) v1.DNSZoneInterface {
	return &FakeDNSZones{c, namespace}
}

func (c *FakeHiveV1) HiveConfigs() v1.HiveConfigInterface {
	return &FakeHiveConfigs{c}
}

func (c *FakeHiveV1) MachinePools(namespace string) v1.MachinePoolInterface {
	return &FakeMachinePools{c, namespace}
}

func (c *FakeHiveV1) MachinePoolNameLeases(namespace string) v1.MachinePoolNameLeaseInterface {
	return &FakeMachinePoolNameLeases{c, namespace}
}

func (c *FakeHiveV1) SelectorSyncIdentityProviders() v1.SelectorSyncIdentityProviderInterface {
	return &FakeSelectorSyncIdentityProviders{c}
}

func (c *FakeHiveV1) SelectorSyncSets() v1.SelectorSyncSetInterface {
	return &FakeSelectorSyncSets{c}
}

func (c *FakeHiveV1) SyncIdentityProviders(namespace string) v1.SyncIdentityProviderInterface {
	return &FakeSyncIdentityProviders{c, namespace}
}

func (c *FakeHiveV1) SyncSets(namespace string) v1.SyncSetInterface {
	return &FakeSyncSets{c, namespace}
}

func (c *FakeHiveV1) SyncSetInstances(namespace string) v1.SyncSetInstanceInterface {
	return &FakeSyncSetInstances{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHiveV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
