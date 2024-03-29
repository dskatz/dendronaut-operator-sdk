// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DendronautJob) DeepCopyInto(out *DendronautJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DendronautJob.
func (in *DendronautJob) DeepCopy() *DendronautJob {
	if in == nil {
		return nil
	}
	out := new(DendronautJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DendronautJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DendronautJobList) DeepCopyInto(out *DendronautJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DendronautJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DendronautJobList.
func (in *DendronautJobList) DeepCopy() *DendronautJobList {
	if in == nil {
		return nil
	}
	out := new(DendronautJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DendronautJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DendronautJobSpec) DeepCopyInto(out *DendronautJobSpec) {
	*out = *in
	in.Cron.DeepCopyInto(&out.Cron)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DendronautJobSpec.
func (in *DendronautJobSpec) DeepCopy() *DendronautJobSpec {
	if in == nil {
		return nil
	}
	out := new(DendronautJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DendronautJobStatus) DeepCopyInto(out *DendronautJobStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DendronautJobStatus.
func (in *DendronautJobStatus) DeepCopy() *DendronautJobStatus {
	if in == nil {
		return nil
	}
	out := new(DendronautJobStatus)
	in.DeepCopyInto(out)
	return out
}
