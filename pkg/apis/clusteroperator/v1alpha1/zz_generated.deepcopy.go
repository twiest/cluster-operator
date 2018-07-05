// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	cluster_v1alpha1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterProviderConfig) DeepCopyInto(out *AWSClusterProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Hardware.DeepCopyInto(&out.Hardware)
	in.OpenShiftConfig.DeepCopyInto(&out.OpenShiftConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterProviderConfig.
func (in *AWSClusterProviderConfig) DeepCopy() *AWSClusterProviderConfig {
	if in == nil {
		return nil
	}
	out := new(AWSClusterProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSClusterProviderConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSClusterSpec) DeepCopyInto(out *AWSClusterSpec) {
	*out = *in
	if in.Defaults != nil {
		in, out := &in.Defaults, &out.Defaults
		if *in == nil {
			*out = nil
		} else {
			*out = new(MachineSetAWSHardwareSpec)
			**out = **in
		}
	}
	out.AccountSecret = in.AccountSecret
	out.SSHSecret = in.SSHSecret
	out.SSLSecret = in.SSLSecret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSClusterSpec.
func (in *AWSClusterSpec) DeepCopy() *AWSClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AWSClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineProviderStatus) DeepCopyInto(out *AWSMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.InstanceState != nil {
		in, out := &in.InstanceState, &out.InstanceState
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.LastELBSync != nil {
		in, out := &in.LastELBSync, &out.LastELBSync
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineProviderStatus.
func (in *AWSMachineProviderStatus) DeepCopy() *AWSMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AWSMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSMachineProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSRegionAMIs) DeepCopyInto(out *AWSRegionAMIs) {
	*out = *in
	if in.MasterAMI != nil {
		in, out := &in.MasterAMI, &out.MasterAMI
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSRegionAMIs.
func (in *AWSRegionAMIs) DeepCopy() *AWSRegionAMIs {
	if in == nil {
		return nil
	}
	out := new(AWSRegionAMIs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVMImages) DeepCopyInto(out *AWSVMImages) {
	*out = *in
	if in.RegionAMIs != nil {
		in, out := &in.RegionAMIs, &out.RegionAMIs
		*out = make([]AWSRegionAMIs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMImages.
func (in *AWSVMImages) DeepCopy() *AWSVMImages {
	if in == nil {
		return nil
	}
	out := new(AWSVMImages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigSpec) DeepCopyInto(out *ClusterConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigSpec.
func (in *ClusterConfigSpec) DeepCopy() *ClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDeployment) DeepCopyInto(out *ClusterDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDeployment.
func (in *ClusterDeployment) DeepCopy() *ClusterDeployment {
	if in == nil {
		return nil
	}
	out := new(ClusterDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDeploymentCondition) DeepCopyInto(out *ClusterDeploymentCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDeploymentCondition.
func (in *ClusterDeploymentCondition) DeepCopy() *ClusterDeploymentCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterDeploymentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDeploymentList) DeepCopyInto(out *ClusterDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDeploymentList.
func (in *ClusterDeploymentList) DeepCopy() *ClusterDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ClusterDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDeploymentSpec) DeepCopyInto(out *ClusterDeploymentSpec) {
	*out = *in
	in.Hardware.DeepCopyInto(&out.Hardware)
	out.Config = in.Config
	if in.DefaultHardwareSpec != nil {
		in, out := &in.DefaultHardwareSpec, &out.DefaultHardwareSpec
		if *in == nil {
			*out = nil
		} else {
			*out = new(MachineSetHardwareSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.MachineSets != nil {
		in, out := &in.MachineSets, &out.MachineSets
		*out = make([]ClusterMachineSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ClusterVersionRef = in.ClusterVersionRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDeploymentSpec.
func (in *ClusterDeploymentSpec) DeepCopy() *ClusterDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDeploymentStatus) DeepCopyInto(out *ClusterDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterDeploymentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDeploymentStatus.
func (in *ClusterDeploymentStatus) DeepCopy() *ClusterDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterHardwareSpec) DeepCopyInto(out *ClusterHardwareSpec) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		if *in == nil {
			*out = nil
		} else {
			*out = new(AWSClusterSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterHardwareSpec.
func (in *ClusterHardwareSpec) DeepCopy() *ClusterHardwareSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterHardwareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMachineSet) DeepCopyInto(out *ClusterMachineSet) {
	*out = *in
	in.MachineSetConfig.DeepCopyInto(&out.MachineSetConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMachineSet.
func (in *ClusterMachineSet) DeepCopy() *ClusterMachineSet {
	if in == nil {
		return nil
	}
	out := new(ClusterMachineSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterProviderStatus) DeepCopyInto(out *ClusterProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AdminKubeconfig != nil {
		in, out := &in.AdminKubeconfig, &out.AdminKubeconfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.LocalObjectReference)
			**out = **in
		}
	}
	if in.NodeConfigInstalledTime != nil {
		in, out := &in.NodeConfigInstalledTime, &out.NodeConfigInstalledTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ClusterAPIInstalledTime != nil {
		in, out := &in.ClusterAPIInstalledTime, &out.ClusterAPIInstalledTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ClusterVersionRef != nil {
		in, out := &in.ClusterVersionRef, &out.ClusterVersionRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.ObjectReference)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterProviderStatus.
func (in *ClusterProviderStatus) DeepCopy() *ClusterProviderStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVersion) DeepCopyInto(out *ClusterVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVersion.
func (in *ClusterVersion) DeepCopy() *ClusterVersion {
	if in == nil {
		return nil
	}
	out := new(ClusterVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVersionImages) DeepCopyInto(out *ClusterVersionImages) {
	*out = *in
	if in.OpenshiftAnsibleImage != nil {
		in, out := &in.OpenshiftAnsibleImage, &out.OpenshiftAnsibleImage
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.OpenshiftAnsibleImagePullPolicy != nil {
		in, out := &in.OpenshiftAnsibleImagePullPolicy, &out.OpenshiftAnsibleImagePullPolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PullPolicy)
			**out = **in
		}
	}
	if in.ClusterAPIImage != nil {
		in, out := &in.ClusterAPIImage, &out.ClusterAPIImage
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.ClusterAPIImagePullPolicy != nil {
		in, out := &in.ClusterAPIImagePullPolicy, &out.ClusterAPIImagePullPolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PullPolicy)
			**out = **in
		}
	}
	if in.MachineControllerImage != nil {
		in, out := &in.MachineControllerImage, &out.MachineControllerImage
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.MachineControllerImagePullPolicy != nil {
		in, out := &in.MachineControllerImagePullPolicy, &out.MachineControllerImagePullPolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PullPolicy)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVersionImages.
func (in *ClusterVersionImages) DeepCopy() *ClusterVersionImages {
	if in == nil {
		return nil
	}
	out := new(ClusterVersionImages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVersionList) DeepCopyInto(out *ClusterVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVersionList.
func (in *ClusterVersionList) DeepCopy() *ClusterVersionList {
	if in == nil {
		return nil
	}
	out := new(ClusterVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVersionReference) DeepCopyInto(out *ClusterVersionReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVersionReference.
func (in *ClusterVersionReference) DeepCopy() *ClusterVersionReference {
	if in == nil {
		return nil
	}
	out := new(ClusterVersionReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVersionSpec) DeepCopyInto(out *ClusterVersionSpec) {
	*out = *in
	in.Images.DeepCopyInto(&out.Images)
	in.VMImages.DeepCopyInto(&out.VMImages)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVersionSpec.
func (in *ClusterVersionSpec) DeepCopy() *ClusterVersionSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVersionStatus) DeepCopyInto(out *ClusterVersionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVersionStatus.
func (in *ClusterVersionStatus) DeepCopy() *ClusterVersionStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CombinedCluster) DeepCopyInto(out *CombinedCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.AWSClusterProviderConfig != nil {
		in, out := &in.AWSClusterProviderConfig, &out.AWSClusterProviderConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(AWSClusterProviderConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ClusterSpec != nil {
		in, out := &in.ClusterSpec, &out.ClusterSpec
		if *in == nil {
			*out = nil
		} else {
			*out = new(cluster_v1alpha1.ClusterSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ClusterStatus != nil {
		in, out := &in.ClusterStatus, &out.ClusterStatus
		if *in == nil {
			*out = nil
		} else {
			*out = new(cluster_v1alpha1.ClusterStatus)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ClusterProviderStatus != nil {
		in, out := &in.ClusterProviderStatus, &out.ClusterProviderStatus
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterProviderStatus)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CombinedCluster.
func (in *CombinedCluster) DeepCopy() *CombinedCluster {
	if in == nil {
		return nil
	}
	out := new(CombinedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CombinedCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSetAWSHardwareSpec) DeepCopyInto(out *MachineSetAWSHardwareSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSetAWSHardwareSpec.
func (in *MachineSetAWSHardwareSpec) DeepCopy() *MachineSetAWSHardwareSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSetAWSHardwareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSetConfig) DeepCopyInto(out *MachineSetConfig) {
	*out = *in
	if in.Hardware != nil {
		in, out := &in.Hardware, &out.Hardware
		if *in == nil {
			*out = nil
		} else {
			*out = new(MachineSetHardwareSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.NodeLabels != nil {
		in, out := &in.NodeLabels, &out.NodeLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSetConfig.
func (in *MachineSetConfig) DeepCopy() *MachineSetConfig {
	if in == nil {
		return nil
	}
	out := new(MachineSetConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSetHardwareSpec) DeepCopyInto(out *MachineSetHardwareSpec) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		if *in == nil {
			*out = nil
		} else {
			*out = new(MachineSetAWSHardwareSpec)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSetHardwareSpec.
func (in *MachineSetHardwareSpec) DeepCopy() *MachineSetHardwareSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSetHardwareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSetProviderConfigSpec) DeepCopyInto(out *MachineSetProviderConfigSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.MachineSetSpec.DeepCopyInto(&out.MachineSetSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSetProviderConfigSpec.
func (in *MachineSetProviderConfigSpec) DeepCopy() *MachineSetProviderConfigSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSetProviderConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineSetProviderConfigSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSetSpec) DeepCopyInto(out *MachineSetSpec) {
	*out = *in
	in.MachineSetConfig.DeepCopyInto(&out.MachineSetConfig)
	in.ClusterHardware.DeepCopyInto(&out.ClusterHardware)
	out.ClusterVersionRef = in.ClusterVersionRef
	in.VMImage.DeepCopyInto(&out.VMImage)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSetSpec.
func (in *MachineSetSpec) DeepCopy() *MachineSetSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftConfigVersion) DeepCopyInto(out *OpenShiftConfigVersion) {
	*out = *in
	in.Images.DeepCopyInto(&out.Images)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftConfigVersion.
func (in *OpenShiftConfigVersion) DeepCopy() *OpenShiftConfigVersion {
	if in == nil {
		return nil
	}
	out := new(OpenShiftConfigVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftDistributionConfig) DeepCopyInto(out *OpenShiftDistributionConfig) {
	*out = *in
	out.ClusterConfigSpec = in.ClusterConfigSpec
	in.Version.DeepCopyInto(&out.Version)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftDistributionConfig.
func (in *OpenShiftDistributionConfig) DeepCopy() *OpenShiftDistributionConfig {
	if in == nil {
		return nil
	}
	out := new(OpenShiftDistributionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMImage) DeepCopyInto(out *VMImage) {
	*out = *in
	if in.AWSImage != nil {
		in, out := &in.AWSImage, &out.AWSImage
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMImage.
func (in *VMImage) DeepCopy() *VMImage {
	if in == nil {
		return nil
	}
	out := new(VMImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMImages) DeepCopyInto(out *VMImages) {
	*out = *in
	if in.AWSImages != nil {
		in, out := &in.AWSImages, &out.AWSImages
		if *in == nil {
			*out = nil
		} else {
			*out = new(AWSVMImages)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMImages.
func (in *VMImages) DeepCopy() *VMImages {
	if in == nil {
		return nil
	}
	out := new(VMImages)
	in.DeepCopyInto(out)
	return out
}
