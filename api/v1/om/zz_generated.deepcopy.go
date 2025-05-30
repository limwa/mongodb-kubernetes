//go:build !ignore_autogenerated

/*
Copyright 2021.

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

// Code generated by controller-gen. DO NOT EDIT.

package om

import (
	"github.com/mongodb/mongodb-kubernetes/api/v1/mdb"
	"github.com/mongodb/mongodb-kubernetes/api/v1/status"
	"github.com/mongodb/mongodb-kubernetes/api/v1/user"
	"github.com/mongodb/mongodb-kubernetes/mongodb-community-operator/api/v1"
	"github.com/mongodb/mongodb-kubernetes/mongodb-community-operator/api/v1/common"
	"github.com/mongodb/mongodb-kubernetes/mongodb-community-operator/pkg/automationconfig"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentVersion) DeepCopyInto(out *AgentVersion) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentVersion.
func (in *AgentVersion) DeepCopy() *AgentVersion {
	if in == nil {
		return nil
	}
	out := new(AgentVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDBConfigurable) DeepCopyInto(out *AppDBConfigurable) {
	*out = *in
	in.AppDBSpec.DeepCopyInto(&out.AppDBSpec)
	in.OpsManager.DeepCopyInto(&out.OpsManager)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDBConfigurable.
func (in *AppDBConfigurable) DeepCopy() *AppDBConfigurable {
	if in == nil {
		return nil
	}
	out := new(AppDBConfigurable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDBSpec) DeepCopyInto(out *AppDBSpec) {
	*out = *in
	if in.PodSpec != nil {
		in, out := &in.PodSpec, &out.PodSpec
		*out = new(mdb.MongoDbPodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.FeatureCompatibilityVersion != nil {
		in, out := &in.FeatureCompatibilityVersion, &out.FeatureCompatibilityVersion
		*out = new(string)
		**out = **in
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(mdb.Security)
		(*in).DeepCopyInto(*out)
	}
	if in.Connectivity != nil {
		in, out := &in.Connectivity, &out.Connectivity
		*out = new(mdb.MongoDBConnectivity)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalAccessConfiguration != nil {
		in, out := &in.ExternalAccessConfiguration, &out.ExternalAccessConfiguration
		*out = new(mdb.ExternalAccessConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalMongodConfig != nil {
		in, out := &in.AdditionalMongodConfig, &out.AdditionalMongodConfig
		*out = (*in).DeepCopy()
	}
	in.AutomationAgent.DeepCopyInto(&out.AutomationAgent)
	in.MonitoringAgent.DeepCopyInto(&out.MonitoringAgent)
	in.ConnectionSpec.DeepCopyInto(&out.ConnectionSpec)
	if in.PasswordSecretKeyRef != nil {
		in, out := &in.PasswordSecretKeyRef, &out.PasswordSecretKeyRef
		*out = new(user.SecretKeyRef)
		**out = **in
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(v1.Prometheus)
		**out = **in
	}
	if in.AutomationConfigOverride != nil {
		in, out := &in.AutomationConfigOverride, &out.AutomationConfigOverride
		*out = new(v1.AutomationConfigOverride)
		(*in).DeepCopyInto(*out)
	}
	if in.MemberConfig != nil {
		in, out := &in.MemberConfig, &out.MemberConfig
		*out = make([]automationconfig.MemberOptions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterSpecList != nil {
		in, out := &in.ClusterSpecList, &out.ClusterSpecList
		*out = make(mdb.ClusterSpecList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDBSpec.
func (in *AppDBSpec) DeepCopy() *AppDBSpec {
	if in == nil {
		return nil
	}
	out := new(AppDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDbBuilder) DeepCopyInto(out *AppDbBuilder) {
	*out = *in
	if in.appDb != nil {
		in, out := &in.appDb, &out.appDb
		*out = new(AppDBSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDbBuilder.
func (in *AppDbBuilder) DeepCopy() *AppDbBuilder {
	if in == nil {
		return nil
	}
	out := new(AppDbBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppDbStatus) DeepCopyInto(out *AppDbStatus) {
	*out = *in
	in.MongoDbStatus.DeepCopyInto(&out.MongoDbStatus)
	if in.ClusterStatusList != nil {
		in, out := &in.ClusterStatusList, &out.ClusterStatusList
		*out = make([]status.ClusterStatusItem, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppDbStatus.
func (in *AppDbStatus) DeepCopy() *AppDbStatus {
	if in == nil {
		return nil
	}
	out := new(AppDbStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStatus) DeepCopyInto(out *BackupStatus) {
	*out = *in
	in.Common.DeepCopyInto(&out.Common)
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]status.Warning, len(*in))
		copy(*out, *in)
	}
	if in.ClusterStatusList != nil {
		in, out := &in.ClusterStatusList, &out.ClusterStatusList
		*out = make([]status.OMClusterStatusItem, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStatus.
func (in *BackupStatus) DeepCopy() *BackupStatus {
	if in == nil {
		return nil
	}
	out := new(BackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecOMItem) DeepCopyInto(out *ClusterSpecOMItem) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.JVMParams != nil {
		in, out := &in.JVMParams, &out.JVMParams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MongoDBOpsManagerExternalConnectivity != nil {
		in, out := &in.MongoDBOpsManagerExternalConnectivity, &out.MongoDBOpsManagerExternalConnectivity
		*out = new(MongoDBOpsManagerServiceDefinition)
		(*in).DeepCopyInto(*out)
	}
	if in.StatefulSetConfiguration != nil {
		in, out := &in.StatefulSetConfiguration, &out.StatefulSetConfiguration
		*out = (*in).DeepCopy()
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(MongoDBOpsManagerBackupClusterSpecItem)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecOMItem.
func (in *ClusterSpecOMItem) DeepCopy() *ClusterSpecOMItem {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecOMItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionSpec) DeepCopyInto(out *ConnectionSpec) {
	*out = *in
	in.SharedConnectionSpec.DeepCopyInto(&out.SharedConnectionSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionSpec.
func (in *ConnectionSpec) DeepCopy() *ConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(ConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataStoreConfig) DeepCopyInto(out *DataStoreConfig) {
	*out = *in
	out.MongoDBResourceRef = in.MongoDBResourceRef
	if in.MongoDBUserRef != nil {
		in, out := &in.MongoDBUserRef, &out.MongoDBUserRef
		*out = new(MongoDBUserRef)
		**out = **in
	}
	if in.AssignmentLabels != nil {
		in, out := &in.AssignmentLabels, &out.AssignmentLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataStoreConfig.
func (in *DataStoreConfig) DeepCopy() *DataStoreConfig {
	if in == nil {
		return nil
	}
	out := new(DataStoreConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Encryption) DeepCopyInto(out *Encryption) {
	*out = *in
	if in.Kmip != nil {
		in, out := &in.Kmip, &out.Kmip
		*out = new(KmipConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Encryption.
func (in *Encryption) DeepCopy() *Encryption {
	if in == nil {
		return nil
	}
	out := new(Encryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemStoreConfig) DeepCopyInto(out *FileSystemStoreConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemStoreConfig.
func (in *FileSystemStoreConfig) DeepCopy() *FileSystemStoreConfig {
	if in == nil {
		return nil
	}
	out := new(FileSystemStoreConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KmipConfig) DeepCopyInto(out *KmipConfig) {
	*out = *in
	out.Server = in.Server
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KmipConfig.
func (in *KmipConfig) DeepCopy() *KmipConfig {
	if in == nil {
		return nil
	}
	out := new(KmipConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logging) DeepCopyInto(out *Logging) {
	*out = *in
	if in.LogBackAccessRef != nil {
		in, out := &in.LogBackAccessRef, &out.LogBackAccessRef
		*out = new(mdb.ConfigMapRef)
		**out = **in
	}
	if in.LogBackRef != nil {
		in, out := &in.LogBackRef, &out.LogBackRef
		*out = new(mdb.ConfigMapRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logging.
func (in *Logging) DeepCopy() *Logging {
	if in == nil {
		return nil
	}
	out := new(Logging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsManager) DeepCopyInto(out *MongoDBOpsManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsManager.
func (in *MongoDBOpsManager) DeepCopy() *MongoDBOpsManager {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBOpsManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsManagerBackup) DeepCopyInto(out *MongoDBOpsManagerBackup) {
	*out = *in
	if in.ExternalServiceEnabled != nil {
		in, out := &in.ExternalServiceEnabled, &out.ExternalServiceEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AssignmentLabels != nil {
		in, out := &in.AssignmentLabels, &out.AssignmentLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HeadDB != nil {
		in, out := &in.HeadDB, &out.HeadDB
		*out = new(common.PersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.JVMParams != nil {
		in, out := &in.JVMParams, &out.JVMParams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.S3OplogStoreConfigs != nil {
		in, out := &in.S3OplogStoreConfigs, &out.S3OplogStoreConfigs
		*out = make([]S3Config, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OplogStoreConfigs != nil {
		in, out := &in.OplogStoreConfigs, &out.OplogStoreConfigs
		*out = make([]DataStoreConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BlockStoreConfigs != nil {
		in, out := &in.BlockStoreConfigs, &out.BlockStoreConfigs
		*out = make([]DataStoreConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.S3Configs != nil {
		in, out := &in.S3Configs, &out.S3Configs
		*out = make([]S3Config, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FileSystemStoreConfigs != nil {
		in, out := &in.FileSystemStoreConfigs, &out.FileSystemStoreConfigs
		*out = make([]FileSystemStoreConfig, len(*in))
		copy(*out, *in)
	}
	if in.StatefulSetConfiguration != nil {
		in, out := &in.StatefulSetConfiguration, &out.StatefulSetConfiguration
		*out = (*in).DeepCopy()
	}
	out.QueryableBackupSecretRef = in.QueryableBackupSecretRef
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(Encryption)
		(*in).DeepCopyInto(*out)
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(Logging)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsManagerBackup.
func (in *MongoDBOpsManagerBackup) DeepCopy() *MongoDBOpsManagerBackup {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsManagerBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsManagerBackupClusterSpecItem) DeepCopyInto(out *MongoDBOpsManagerBackupClusterSpecItem) {
	*out = *in
	if in.AssignmentLabels != nil {
		in, out := &in.AssignmentLabels, &out.AssignmentLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HeadDB != nil {
		in, out := &in.HeadDB, &out.HeadDB
		*out = new(common.PersistenceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.JVMParams != nil {
		in, out := &in.JVMParams, &out.JVMParams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StatefulSetConfiguration != nil {
		in, out := &in.StatefulSetConfiguration, &out.StatefulSetConfiguration
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsManagerBackupClusterSpecItem.
func (in *MongoDBOpsManagerBackupClusterSpecItem) DeepCopy() *MongoDBOpsManagerBackupClusterSpecItem {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsManagerBackupClusterSpecItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsManagerList) DeepCopyInto(out *MongoDBOpsManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongoDBOpsManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsManagerList.
func (in *MongoDBOpsManagerList) DeepCopy() *MongoDBOpsManagerList {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongoDBOpsManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsManagerSecurity) DeepCopyInto(out *MongoDBOpsManagerSecurity) {
	*out = *in
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsManagerSecurity.
func (in *MongoDBOpsManagerSecurity) DeepCopy() *MongoDBOpsManagerSecurity {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsManagerSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsManagerServiceDefinition) DeepCopyInto(out *MongoDBOpsManagerServiceDefinition) {
	*out = *in
	if in.ClusterIP != nil {
		in, out := &in.ClusterIP, &out.ClusterIP
		*out = new(string)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsManagerServiceDefinition.
func (in *MongoDBOpsManagerServiceDefinition) DeepCopy() *MongoDBOpsManagerServiceDefinition {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsManagerServiceDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsManagerSpec) DeepCopyInto(out *MongoDBOpsManagerSpec) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.AppDB.DeepCopyInto(&out.AppDB)
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(Logging)
		(*in).DeepCopyInto(*out)
	}
	if in.JVMParams != nil {
		in, out := &in.JVMParams, &out.JVMParams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(MongoDBOpsManagerBackup)
		(*in).DeepCopyInto(*out)
	}
	if in.InternalConnectivity != nil {
		in, out := &in.InternalConnectivity, &out.InternalConnectivity
		*out = new(MongoDBOpsManagerServiceDefinition)
		(*in).DeepCopyInto(*out)
	}
	if in.MongoDBOpsManagerExternalConnectivity != nil {
		in, out := &in.MongoDBOpsManagerExternalConnectivity, &out.MongoDBOpsManagerExternalConnectivity
		*out = new(MongoDBOpsManagerServiceDefinition)
		(*in).DeepCopyInto(*out)
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(MongoDBOpsManagerSecurity)
		**out = **in
	}
	if in.StatefulSetConfiguration != nil {
		in, out := &in.StatefulSetConfiguration, &out.StatefulSetConfiguration
		*out = (*in).DeepCopy()
	}
	if in.ClusterSpecList != nil {
		in, out := &in.ClusterSpecList, &out.ClusterSpecList
		*out = make([]ClusterSpecOMItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsManagerSpec.
func (in *MongoDBOpsManagerSpec) DeepCopy() *MongoDBOpsManagerSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsManagerStatus) DeepCopyInto(out *MongoDBOpsManagerStatus) {
	*out = *in
	in.OpsManagerStatus.DeepCopyInto(&out.OpsManagerStatus)
	in.AppDbStatus.DeepCopyInto(&out.AppDbStatus)
	in.BackupStatus.DeepCopyInto(&out.BackupStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsManagerStatus.
func (in *MongoDBOpsManagerStatus) DeepCopy() *MongoDBOpsManagerStatus {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBOpsManagerTLS) DeepCopyInto(out *MongoDBOpsManagerTLS) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBOpsManagerTLS.
func (in *MongoDBOpsManagerTLS) DeepCopy() *MongoDBOpsManagerTLS {
	if in == nil {
		return nil
	}
	out := new(MongoDBOpsManagerTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBUserRef) DeepCopyInto(out *MongoDBUserRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBUserRef.
func (in *MongoDBUserRef) DeepCopy() *MongoDBUserRef {
	if in == nil {
		return nil
	}
	out := new(MongoDBUserRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsManagerBuilder) DeepCopyInto(out *OpsManagerBuilder) {
	*out = *in
	in.om.DeepCopyInto(&out.om)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsManagerBuilder.
func (in *OpsManagerBuilder) DeepCopy() *OpsManagerBuilder {
	if in == nil {
		return nil
	}
	out := new(OpsManagerBuilder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsManagerStatus) DeepCopyInto(out *OpsManagerStatus) {
	*out = *in
	in.Common.DeepCopyInto(&out.Common)
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]status.Warning, len(*in))
		copy(*out, *in)
	}
	if in.ClusterStatusList != nil {
		in, out := &in.ClusterStatusList, &out.ClusterStatusList
		*out = make([]status.OMClusterStatusItem, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsManagerStatus.
func (in *OpsManagerStatus) DeepCopy() *OpsManagerStatus {
	if in == nil {
		return nil
	}
	out := new(OpsManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpsManagerVersionMapping) DeepCopyInto(out *OpsManagerVersionMapping) {
	*out = *in
	if in.OpsManager != nil {
		in, out := &in.OpsManager, &out.OpsManager
		*out = make(map[OpsManagerVersion]AgentVersion, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpsManagerVersionMapping.
func (in *OpsManagerVersionMapping) DeepCopy() *OpsManagerVersionMapping {
	if in == nil {
		return nil
	}
	out := new(OpsManagerVersionMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Config) DeepCopyInto(out *S3Config) {
	*out = *in
	if in.MongoDBResourceRef != nil {
		in, out := &in.MongoDBResourceRef, &out.MongoDBResourceRef
		*out = new(user.MongoDBResourceRef)
		**out = **in
	}
	if in.MongoDBUserRef != nil {
		in, out := &in.MongoDBUserRef, &out.MongoDBUserRef
		*out = new(MongoDBUserRef)
		**out = **in
	}
	if in.S3SecretRef != nil {
		in, out := &in.S3SecretRef, &out.S3SecretRef
		*out = new(SecretRef)
		**out = **in
	}
	if in.AssignmentLabels != nil {
		in, out := &in.AssignmentLabels, &out.AssignmentLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomCertificateSecretRefs != nil {
		in, out := &in.CustomCertificateSecretRefs, &out.CustomCertificateSecretRefs
		*out = make([]corev1.SecretKeySelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Config.
func (in *S3Config) DeepCopy() *S3Config {
	if in == nil {
		return nil
	}
	out := new(S3Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSecretRef) DeepCopyInto(out *TLSSecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSecretRef.
func (in *TLSSecretRef) DeepCopy() *TLSSecretRef {
	if in == nil {
		return nil
	}
	out := new(TLSSecretRef)
	in.DeepCopyInto(out)
	return out
}
