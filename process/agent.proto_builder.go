// THIS IS A GENERATED FILE
// DO NOT EDIT
package process

import (
	bytes "bytes"
	protowire "google.golang.org/protobuf/encoding/protowire"
	io "io"
	math "math"
)

type ResCollectorBuilder struct {
	writer                     io.Writer
	buf                        bytes.Buffer
	scratch                    []byte
	resCollector_HeaderBuilder ResCollector_HeaderBuilder
	collectorStatusBuilder     CollectorStatusBuilder
}

func NewResCollectorBuilder(writer io.Writer) *ResCollectorBuilder {
	return &ResCollectorBuilder{
		writer: writer,
	}
}
func (x *ResCollectorBuilder) SetHeader(cb func(w *ResCollector_HeaderBuilder)) {
	x.buf.Reset()
	x.resCollector_HeaderBuilder.writer = &x.buf
	x.resCollector_HeaderBuilder.scratch = x.scratch
	cb(&x.resCollector_HeaderBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ResCollectorBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResCollectorBuilder) SetStatus(cb func(w *CollectorStatusBuilder)) {
	x.buf.Reset()
	x.collectorStatusBuilder.writer = &x.buf
	x.collectorStatusBuilder.scratch = x.scratch
	cb(&x.collectorStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ResCollector_HeaderBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewResCollector_HeaderBuilder(writer io.Writer) *ResCollector_HeaderBuilder {
	return &ResCollector_HeaderBuilder{
		writer: writer,
	}
}
func (x *ResCollector_HeaderBuilder) SetType(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type CollectorProcBuilder struct {
	writer            io.Writer
	buf               bytes.Buffer
	scratch           []byte
	processBuilder    ProcessBuilder
	hostBuilder       HostBuilder
	systemInfoBuilder SystemInfoBuilder
	containerBuilder  ContainerBuilder
}

func NewCollectorProcBuilder(writer io.Writer) *CollectorProcBuilder {
	return &CollectorProcBuilder{
		writer: writer,
	}
}
func (x *CollectorProcBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorProcBuilder) SetNetworkId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x5a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorProcBuilder) AddProcesses(cb func(w *ProcessBuilder)) {
	x.buf.Reset()
	x.processBuilder.writer = &x.buf
	x.processBuilder.scratch = x.scratch
	cb(&x.processBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcBuilder) SetInfo(cb func(w *SystemInfoBuilder)) {
	x.buf.Reset()
	x.systemInfoBuilder.writer = &x.buf
	x.systemInfoBuilder.scratch = x.scratch
	cb(&x.systemInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorProcBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorProcBuilder) AddContainers(cb func(w *ContainerBuilder)) {
	x.buf.Reset()
	x.containerBuilder.writer = &x.buf
	x.containerBuilder.scratch = x.scratch
	cb(&x.containerBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcBuilder) SetContainerHostType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x60)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *CollectorProcBuilder) SetHintMask(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x70)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type CollectorProcDiscoveryBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	processDiscoveryBuilder ProcessDiscoveryBuilder
	hostBuilder             HostBuilder
}

func NewCollectorProcDiscoveryBuilder(writer io.Writer) *CollectorProcDiscoveryBuilder {
	return &CollectorProcDiscoveryBuilder{
		writer: writer,
	}
}
func (x *CollectorProcDiscoveryBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorProcDiscoveryBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorProcDiscoveryBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorProcDiscoveryBuilder) AddProcessDiscoveries(cb func(w *ProcessDiscoveryBuilder)) {
	x.buf.Reset()
	x.processDiscoveryBuilder.writer = &x.buf
	x.processDiscoveryBuilder.scratch = x.scratch
	cb(&x.processDiscoveryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorProcDiscoveryBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorRealTimeBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	processStatBuilder   ProcessStatBuilder
	containerStatBuilder ContainerStatBuilder
}

func NewCollectorRealTimeBuilder(writer io.Writer) *CollectorRealTimeBuilder {
	return &CollectorRealTimeBuilder{
		writer: writer,
	}
}
func (x *CollectorRealTimeBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) AddStats(cb func(w *ProcessStatBuilder)) {
	x.buf.Reset()
	x.processStatBuilder.writer = &x.buf
	x.processStatBuilder.scratch = x.scratch
	cb(&x.processStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorRealTimeBuilder) SetHostId(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) SetOrgId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) SetNumCpus(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) SetTotalMemory(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRealTimeBuilder) AddContainerStats(cb func(w *ContainerStatBuilder)) {
	x.buf.Reset()
	x.containerStatBuilder.writer = &x.buf
	x.containerStatBuilder.scratch = x.scratch
	cb(&x.containerStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorRealTimeBuilder) SetContainerHostType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x58)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type CollectorContainerBuilder struct {
	writer            io.Writer
	buf               bytes.Buffer
	scratch           []byte
	systemInfoBuilder SystemInfoBuilder
	containerBuilder  ContainerBuilder
	hostBuilder       HostBuilder
}

func NewCollectorContainerBuilder(writer io.Writer) *CollectorContainerBuilder {
	return &CollectorContainerBuilder{
		writer: writer,
	}
}
func (x *CollectorContainerBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerBuilder) SetNetworkId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x5a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerBuilder) SetInfo(cb func(w *SystemInfoBuilder)) {
	x.buf.Reset()
	x.systemInfoBuilder.writer = &x.buf
	x.systemInfoBuilder.scratch = x.scratch
	cb(&x.systemInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorContainerBuilder) AddContainers(cb func(w *ContainerBuilder)) {
	x.buf.Reset()
	x.containerBuilder.writer = &x.buf
	x.containerBuilder.scratch = x.scratch
	cb(&x.containerBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorContainerBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorContainerBuilder) SetContainerHostType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x48)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type CollectorContainerRealTimeBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	containerStatBuilder ContainerStatBuilder
}

func NewCollectorContainerRealTimeBuilder(writer io.Writer) *CollectorContainerRealTimeBuilder {
	return &CollectorContainerRealTimeBuilder{
		writer: writer,
	}
}
func (x *CollectorContainerRealTimeBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerRealTimeBuilder) AddStats(cb func(w *ContainerStatBuilder)) {
	x.buf.Reset()
	x.containerStatBuilder.writer = &x.buf
	x.containerStatBuilder.scratch = x.scratch
	cb(&x.containerStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorContainerRealTimeBuilder) SetNumCpus(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerRealTimeBuilder) SetTotalMemory(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerRealTimeBuilder) SetHostId(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerRealTimeBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerRealTimeBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorContainerRealTimeBuilder) SetContainerHostType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x40)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type CollectorReqStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCollectorReqStatusBuilder(writer io.Writer) *CollectorReqStatusBuilder {
	return &CollectorReqStatusBuilder{
		writer: writer,
	}
}
func (x *CollectorReqStatusBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorPodBuilder struct {
	writer      io.Writer
	buf         bytes.Buffer
	scratch     []byte
	podBuilder  PodBuilder
	hostBuilder HostBuilder
}

func NewCollectorPodBuilder(writer io.Writer) *CollectorPodBuilder {
	return &CollectorPodBuilder{
		writer: writer,
	}
}
func (x *CollectorPodBuilder) SetHostName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPodBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPodBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPodBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPodBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPodBuilder) AddPods(cb func(w *PodBuilder)) {
	x.buf.Reset()
	x.podBuilder.writer = &x.buf
	x.podBuilder.scratch = x.scratch
	cb(&x.podBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorPodBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorPodBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorReplicaSetBuilder struct {
	writer            io.Writer
	buf               bytes.Buffer
	scratch           []byte
	replicaSetBuilder ReplicaSetBuilder
}

func NewCollectorReplicaSetBuilder(writer io.Writer) *CollectorReplicaSetBuilder {
	return &CollectorReplicaSetBuilder{
		writer: writer,
	}
}
func (x *CollectorReplicaSetBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorReplicaSetBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorReplicaSetBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorReplicaSetBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorReplicaSetBuilder) AddReplicaSets(cb func(w *ReplicaSetBuilder)) {
	x.buf.Reset()
	x.replicaSetBuilder.writer = &x.buf
	x.replicaSetBuilder.scratch = x.scratch
	cb(&x.replicaSetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorReplicaSetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorDeploymentBuilder struct {
	writer            io.Writer
	buf               bytes.Buffer
	scratch           []byte
	deploymentBuilder DeploymentBuilder
}

func NewCollectorDeploymentBuilder(writer io.Writer) *CollectorDeploymentBuilder {
	return &CollectorDeploymentBuilder{
		writer: writer,
	}
}
func (x *CollectorDeploymentBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorDeploymentBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorDeploymentBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorDeploymentBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorDeploymentBuilder) AddDeployments(cb func(w *DeploymentBuilder)) {
	x.buf.Reset()
	x.deploymentBuilder.writer = &x.buf
	x.deploymentBuilder.scratch = x.scratch
	cb(&x.deploymentBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorDeploymentBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorServiceBuilder struct {
	writer         io.Writer
	buf            bytes.Buffer
	scratch        []byte
	serviceBuilder ServiceBuilder
}

func NewCollectorServiceBuilder(writer io.Writer) *CollectorServiceBuilder {
	return &CollectorServiceBuilder{
		writer: writer,
	}
}
func (x *CollectorServiceBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceBuilder) AddServices(cb func(w *ServiceBuilder)) {
	x.buf.Reset()
	x.serviceBuilder.writer = &x.buf
	x.serviceBuilder.scratch = x.scratch
	cb(&x.serviceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorServiceBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorNodeBuilder struct {
	writer                                     io.Writer
	buf                                        bytes.Buffer
	scratch                                    []byte
	nodeBuilder                                NodeBuilder
	collectorNode_HostAliasMappingEntryBuilder CollectorNode_HostAliasMappingEntryBuilder
}

func NewCollectorNodeBuilder(writer io.Writer) *CollectorNodeBuilder {
	return &CollectorNodeBuilder{
		writer: writer,
	}
}
func (x *CollectorNodeBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNodeBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNodeBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorNodeBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorNodeBuilder) AddNodes(cb func(w *NodeBuilder)) {
	x.buf.Reset()
	x.nodeBuilder.writer = &x.buf
	x.nodeBuilder.scratch = x.scratch
	cb(&x.nodeBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorNodeBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNodeBuilder) AddHostAliasMapping(cb func(w *CollectorNode_HostAliasMappingEntryBuilder)) {
	x.buf.Reset()
	x.collectorNode_HostAliasMappingEntryBuilder.writer = &x.buf
	x.collectorNode_HostAliasMappingEntryBuilder.scratch = x.scratch
	cb(&x.collectorNode_HostAliasMappingEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorNode_HostAliasMappingEntryBuilder struct {
	writer      io.Writer
	buf         bytes.Buffer
	scratch     []byte
	hostBuilder HostBuilder
}

func NewCollectorNode_HostAliasMappingEntryBuilder(writer io.Writer) *CollectorNode_HostAliasMappingEntryBuilder {
	return &CollectorNode_HostAliasMappingEntryBuilder{
		writer: writer,
	}
}
func (x *CollectorNode_HostAliasMappingEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNode_HostAliasMappingEntryBuilder) SetValue(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorClusterBuilder struct {
	writer         io.Writer
	buf            bytes.Buffer
	scratch        []byte
	clusterBuilder ClusterBuilder
}

func NewCollectorClusterBuilder(writer io.Writer) *CollectorClusterBuilder {
	return &CollectorClusterBuilder{
		writer: writer,
	}
}
func (x *CollectorClusterBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterBuilder) SetCluster(cb func(w *ClusterBuilder)) {
	x.buf.Reset()
	x.clusterBuilder.writer = &x.buf
	x.clusterBuilder.scratch = x.scratch
	cb(&x.clusterBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorClusterBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorManifestBuilder struct {
	writer          io.Writer
	buf             bytes.Buffer
	scratch         []byte
	manifestBuilder ManifestBuilder
}

func NewCollectorManifestBuilder(writer io.Writer) *CollectorManifestBuilder {
	return &CollectorManifestBuilder{
		writer: writer,
	}
}
func (x *CollectorManifestBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorManifestBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorManifestBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorManifestBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorManifestBuilder) AddManifests(cb func(w *ManifestBuilder)) {
	x.buf.Reset()
	x.manifestBuilder.writer = &x.buf
	x.manifestBuilder.scratch = x.scratch
	cb(&x.manifestBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorManifestCRDBuilder struct {
	writer                   io.Writer
	buf                      bytes.Buffer
	scratch                  []byte
	collectorManifestBuilder CollectorManifestBuilder
}

func NewCollectorManifestCRDBuilder(writer io.Writer) *CollectorManifestCRDBuilder {
	return &CollectorManifestCRDBuilder{
		writer: writer,
	}
}
func (x *CollectorManifestCRDBuilder) SetManifest(cb func(w *CollectorManifestBuilder)) {
	x.buf.Reset()
	x.collectorManifestBuilder.writer = &x.buf
	x.collectorManifestBuilder.scratch = x.scratch
	cb(&x.collectorManifestBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorManifestCRBuilder struct {
	writer                   io.Writer
	buf                      bytes.Buffer
	scratch                  []byte
	collectorManifestBuilder CollectorManifestBuilder
}

func NewCollectorManifestCRBuilder(writer io.Writer) *CollectorManifestCRBuilder {
	return &CollectorManifestCRBuilder{
		writer: writer,
	}
}
func (x *CollectorManifestCRBuilder) SetManifest(cb func(w *CollectorManifestBuilder)) {
	x.buf.Reset()
	x.collectorManifestBuilder.writer = &x.buf
	x.collectorManifestBuilder.scratch = x.scratch
	cb(&x.collectorManifestBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CollectorNamespaceBuilder struct {
	writer           io.Writer
	buf              bytes.Buffer
	scratch          []byte
	namespaceBuilder NamespaceBuilder
}

func NewCollectorNamespaceBuilder(writer io.Writer) *CollectorNamespaceBuilder {
	return &CollectorNamespaceBuilder{
		writer: writer,
	}
}
func (x *CollectorNamespaceBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNamespaceBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorNamespaceBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorNamespaceBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorNamespaceBuilder) AddNamespaces(cb func(w *NamespaceBuilder)) {
	x.buf.Reset()
	x.namespaceBuilder.writer = &x.buf
	x.namespaceBuilder.scratch = x.scratch
	cb(&x.namespaceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorNamespaceBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorJobBuilder struct {
	writer     io.Writer
	buf        bytes.Buffer
	scratch    []byte
	jobBuilder JobBuilder
}

func NewCollectorJobBuilder(writer io.Writer) *CollectorJobBuilder {
	return &CollectorJobBuilder{
		writer: writer,
	}
}
func (x *CollectorJobBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorJobBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorJobBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorJobBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorJobBuilder) AddJobs(cb func(w *JobBuilder)) {
	x.buf.Reset()
	x.jobBuilder.writer = &x.buf
	x.jobBuilder.scratch = x.scratch
	cb(&x.jobBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorJobBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorCronJobBuilder struct {
	writer         io.Writer
	buf            bytes.Buffer
	scratch        []byte
	cronJobBuilder CronJobBuilder
}

func NewCollectorCronJobBuilder(writer io.Writer) *CollectorCronJobBuilder {
	return &CollectorCronJobBuilder{
		writer: writer,
	}
}
func (x *CollectorCronJobBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorCronJobBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorCronJobBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorCronJobBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorCronJobBuilder) AddCronJobs(cb func(w *CronJobBuilder)) {
	x.buf.Reset()
	x.cronJobBuilder.writer = &x.buf
	x.cronJobBuilder.scratch = x.scratch
	cb(&x.cronJobBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorCronJobBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorDaemonSetBuilder struct {
	writer           io.Writer
	buf              bytes.Buffer
	scratch          []byte
	daemonSetBuilder DaemonSetBuilder
}

func NewCollectorDaemonSetBuilder(writer io.Writer) *CollectorDaemonSetBuilder {
	return &CollectorDaemonSetBuilder{
		writer: writer,
	}
}
func (x *CollectorDaemonSetBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorDaemonSetBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorDaemonSetBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorDaemonSetBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorDaemonSetBuilder) AddDaemonSets(cb func(w *DaemonSetBuilder)) {
	x.buf.Reset()
	x.daemonSetBuilder.writer = &x.buf
	x.daemonSetBuilder.scratch = x.scratch
	cb(&x.daemonSetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorDaemonSetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorStatefulSetBuilder struct {
	writer             io.Writer
	buf                bytes.Buffer
	scratch            []byte
	statefulSetBuilder StatefulSetBuilder
}

func NewCollectorStatefulSetBuilder(writer io.Writer) *CollectorStatefulSetBuilder {
	return &CollectorStatefulSetBuilder{
		writer: writer,
	}
}
func (x *CollectorStatefulSetBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorStatefulSetBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorStatefulSetBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorStatefulSetBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorStatefulSetBuilder) AddStatefulSets(cb func(w *StatefulSetBuilder)) {
	x.buf.Reset()
	x.statefulSetBuilder.writer = &x.buf
	x.statefulSetBuilder.scratch = x.scratch
	cb(&x.statefulSetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorStatefulSetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorPersistentVolumeBuilder struct {
	writer                  io.Writer
	buf                     bytes.Buffer
	scratch                 []byte
	persistentVolumeBuilder PersistentVolumeBuilder
}

func NewCollectorPersistentVolumeBuilder(writer io.Writer) *CollectorPersistentVolumeBuilder {
	return &CollectorPersistentVolumeBuilder{
		writer: writer,
	}
}
func (x *CollectorPersistentVolumeBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeBuilder) AddPersistentVolumes(cb func(w *PersistentVolumeBuilder)) {
	x.buf.Reset()
	x.persistentVolumeBuilder.writer = &x.buf
	x.persistentVolumeBuilder.scratch = x.scratch
	cb(&x.persistentVolumeBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorPersistentVolumeBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorPersistentVolumeClaimBuilder struct {
	writer                       io.Writer
	buf                          bytes.Buffer
	scratch                      []byte
	persistentVolumeClaimBuilder PersistentVolumeClaimBuilder
}

func NewCollectorPersistentVolumeClaimBuilder(writer io.Writer) *CollectorPersistentVolumeClaimBuilder {
	return &CollectorPersistentVolumeClaimBuilder{
		writer: writer,
	}
}
func (x *CollectorPersistentVolumeClaimBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeClaimBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeClaimBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeClaimBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorPersistentVolumeClaimBuilder) AddPersistentVolumeClaims(cb func(w *PersistentVolumeClaimBuilder)) {
	x.buf.Reset()
	x.persistentVolumeClaimBuilder.writer = &x.buf
	x.persistentVolumeClaimBuilder.scratch = x.scratch
	cb(&x.persistentVolumeClaimBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorPersistentVolumeClaimBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorRoleBuilder struct {
	writer      io.Writer
	buf         bytes.Buffer
	scratch     []byte
	roleBuilder RoleBuilder
}

func NewCollectorRoleBuilder(writer io.Writer) *CollectorRoleBuilder {
	return &CollectorRoleBuilder{
		writer: writer,
	}
}
func (x *CollectorRoleBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBuilder) AddRoles(cb func(w *RoleBuilder)) {
	x.buf.Reset()
	x.roleBuilder.writer = &x.buf
	x.roleBuilder.scratch = x.scratch
	cb(&x.roleBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorRoleBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorRoleBindingBuilder struct {
	writer             io.Writer
	buf                bytes.Buffer
	scratch            []byte
	roleBindingBuilder RoleBindingBuilder
}

func NewCollectorRoleBindingBuilder(writer io.Writer) *CollectorRoleBindingBuilder {
	return &CollectorRoleBindingBuilder{
		writer: writer,
	}
}
func (x *CollectorRoleBindingBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBindingBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBindingBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBindingBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorRoleBindingBuilder) AddRoleBindings(cb func(w *RoleBindingBuilder)) {
	x.buf.Reset()
	x.roleBindingBuilder.writer = &x.buf
	x.roleBindingBuilder.scratch = x.scratch
	cb(&x.roleBindingBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorRoleBindingBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorClusterRoleBuilder struct {
	writer             io.Writer
	buf                bytes.Buffer
	scratch            []byte
	clusterRoleBuilder ClusterRoleBuilder
}

func NewCollectorClusterRoleBuilder(writer io.Writer) *CollectorClusterRoleBuilder {
	return &CollectorClusterRoleBuilder{
		writer: writer,
	}
}
func (x *CollectorClusterRoleBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBuilder) AddClusterRoles(cb func(w *ClusterRoleBuilder)) {
	x.buf.Reset()
	x.clusterRoleBuilder.writer = &x.buf
	x.clusterRoleBuilder.scratch = x.scratch
	cb(&x.clusterRoleBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorClusterRoleBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorClusterRoleBindingBuilder struct {
	writer                    io.Writer
	buf                       bytes.Buffer
	scratch                   []byte
	clusterRoleBindingBuilder ClusterRoleBindingBuilder
}

func NewCollectorClusterRoleBindingBuilder(writer io.Writer) *CollectorClusterRoleBindingBuilder {
	return &CollectorClusterRoleBindingBuilder{
		writer: writer,
	}
}
func (x *CollectorClusterRoleBindingBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBindingBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBindingBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBindingBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorClusterRoleBindingBuilder) AddClusterRoleBindings(cb func(w *ClusterRoleBindingBuilder)) {
	x.buf.Reset()
	x.clusterRoleBindingBuilder.writer = &x.buf
	x.clusterRoleBindingBuilder.scratch = x.scratch
	cb(&x.clusterRoleBindingBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorClusterRoleBindingBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorServiceAccountBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	serviceAccountBuilder ServiceAccountBuilder
}

func NewCollectorServiceAccountBuilder(writer io.Writer) *CollectorServiceAccountBuilder {
	return &CollectorServiceAccountBuilder{
		writer: writer,
	}
}
func (x *CollectorServiceAccountBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceAccountBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceAccountBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceAccountBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorServiceAccountBuilder) AddServiceAccounts(cb func(w *ServiceAccountBuilder)) {
	x.buf.Reset()
	x.serviceAccountBuilder.writer = &x.buf
	x.serviceAccountBuilder.scratch = x.scratch
	cb(&x.serviceAccountBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorServiceAccountBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorIngressBuilder struct {
	writer         io.Writer
	buf            bytes.Buffer
	scratch        []byte
	ingressBuilder IngressBuilder
}

func NewCollectorIngressBuilder(writer io.Writer) *CollectorIngressBuilder {
	return &CollectorIngressBuilder{
		writer: writer,
	}
}
func (x *CollectorIngressBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorIngressBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorIngressBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorIngressBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorIngressBuilder) AddIngresses(cb func(w *IngressBuilder)) {
	x.buf.Reset()
	x.ingressBuilder.writer = &x.buf
	x.ingressBuilder.scratch = x.scratch
	cb(&x.ingressBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorIngressBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorVerticalPodAutoscalerBuilder struct {
	writer                       io.Writer
	buf                          bytes.Buffer
	scratch                      []byte
	verticalPodAutoscalerBuilder VerticalPodAutoscalerBuilder
}

func NewCollectorVerticalPodAutoscalerBuilder(writer io.Writer) *CollectorVerticalPodAutoscalerBuilder {
	return &CollectorVerticalPodAutoscalerBuilder{
		writer: writer,
	}
}
func (x *CollectorVerticalPodAutoscalerBuilder) SetClusterName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorVerticalPodAutoscalerBuilder) SetClusterId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CollectorVerticalPodAutoscalerBuilder) SetGroupId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorVerticalPodAutoscalerBuilder) SetGroupSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorVerticalPodAutoscalerBuilder) AddVerticalPodAutoscalers(cb func(w *VerticalPodAutoscalerBuilder)) {
	x.buf.Reset()
	x.verticalPodAutoscalerBuilder.writer = &x.buf
	x.verticalPodAutoscalerBuilder.scratch = x.scratch
	cb(&x.verticalPodAutoscalerBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CollectorVerticalPodAutoscalerBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CollectorStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCollectorStatusBuilder(writer io.Writer) *CollectorStatusBuilder {
	return &CollectorStatusBuilder{
		writer: writer,
	}
}
func (x *CollectorStatusBuilder) SetActiveClients(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CollectorStatusBuilder) SetInterval(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ProcessBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	hostBuilder            HostBuilder
	commandBuilder         CommandBuilder
	processUserBuilder     ProcessUserBuilder
	memoryStatBuilder      MemoryStatBuilder
	cPUStatBuilder         CPUStatBuilder
	containerBuilder       ContainerBuilder
	iOStatBuilder          IOStatBuilder
	processNetworksBuilder ProcessNetworksBuilder
}

func NewProcessBuilder(writer io.Writer) *ProcessBuilder {
	return &ProcessBuilder{
		writer: writer,
	}
}
func (x *ProcessBuilder) SetKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetPid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetNsPid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetCommand(cb func(w *CommandBuilder)) {
	x.buf.Reset()
	x.commandBuilder.writer = &x.buf
	x.commandBuilder.scratch = x.scratch
	cb(&x.commandBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetUser(cb func(w *ProcessUserBuilder)) {
	x.buf.Reset()
	x.processUserBuilder.writer = &x.buf
	x.processUserBuilder.scratch = x.scratch
	cb(&x.processUserBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetMemory(cb func(w *MemoryStatBuilder)) {
	x.buf.Reset()
	x.memoryStatBuilder.writer = &x.buf
	x.memoryStatBuilder.scratch = x.scratch
	cb(&x.memoryStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetCpu(cb func(w *CPUStatBuilder)) {
	x.buf.Reset()
	x.cPUStatBuilder.writer = &x.buf
	x.cPUStatBuilder.scratch = x.scratch
	cb(&x.cPUStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetCreateTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetContainer(cb func(w *ContainerBuilder)) {
	x.buf.Reset()
	x.containerBuilder.writer = &x.buf
	x.containerBuilder.scratch = x.scratch
	cb(&x.containerBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetOpenFdCount(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x58)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetState(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x60)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ProcessBuilder) SetIoStat(cb func(w *IOStatBuilder)) {
	x.buf.Reset()
	x.iOStatBuilder.writer = &x.buf
	x.iOStatBuilder.scratch = x.scratch
	cb(&x.iOStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x6a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetContainerId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x72)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetContainerKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x78)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetVoluntaryCtxSwitches(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x80)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetInvoluntaryCtxSwitches(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x88)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessBuilder) SetByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x92)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetContainerByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x9a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessBuilder) SetNetworks(cb func(w *ProcessNetworksBuilder)) {
	x.buf.Reset()
	x.processNetworksBuilder.writer = &x.buf
	x.processNetworksBuilder.scratch = x.scratch
	cb(&x.processNetworksBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xaa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ProcessDiscoveryBuilder struct {
	writer             io.Writer
	buf                bytes.Buffer
	scratch            []byte
	hostBuilder        HostBuilder
	commandBuilder     CommandBuilder
	processUserBuilder ProcessUserBuilder
}

func NewProcessDiscoveryBuilder(writer io.Writer) *ProcessDiscoveryBuilder {
	return &ProcessDiscoveryBuilder{
		writer: writer,
	}
}
func (x *ProcessDiscoveryBuilder) SetPid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessDiscoveryBuilder) SetNsPid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessDiscoveryBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessDiscoveryBuilder) SetCommand(cb func(w *CommandBuilder)) {
	x.buf.Reset()
	x.commandBuilder.writer = &x.buf
	x.commandBuilder.scratch = x.scratch
	cb(&x.commandBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessDiscoveryBuilder) SetUser(cb func(w *ProcessUserBuilder)) {
	x.buf.Reset()
	x.processUserBuilder.writer = &x.buf
	x.processUserBuilder.scratch = x.scratch
	cb(&x.processUserBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessDiscoveryBuilder) SetCreateTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessDiscoveryBuilder) SetByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CommandBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCommandBuilder(writer io.Writer) *CommandBuilder {
	return &CommandBuilder{
		writer: writer,
	}
}
func (x *CommandBuilder) AddArgs(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CommandBuilder) SetCwd(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CommandBuilder) SetRoot(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CommandBuilder) SetOnDisk(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x28)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *CommandBuilder) SetPpid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CommandBuilder) SetPgroup(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CommandBuilder) SetExe(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ProcessUserBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewProcessUserBuilder(writer io.Writer) *ProcessUserBuilder {
	return &ProcessUserBuilder{
		writer: writer,
	}
}
func (x *ProcessUserBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ProcessUserBuilder) SetUid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessUserBuilder) SetGid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessUserBuilder) SetEuid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessUserBuilder) SetEgid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessUserBuilder) SetSuid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessUserBuilder) SetSgid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ProcessNetworksBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewProcessNetworksBuilder(writer io.Writer) *ProcessNetworksBuilder {
	return &ProcessNetworksBuilder{
		writer: writer,
	}
}
func (x *ProcessNetworksBuilder) SetConnectionRate(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x9)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessNetworksBuilder) SetBytesRate(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}

type ContainerAddrBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewContainerAddrBuilder(writer io.Writer) *ContainerAddrBuilder {
	return &ContainerAddrBuilder{
		writer: writer,
	}
}
func (x *ContainerAddrBuilder) SetIp(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerAddrBuilder) SetPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerAddrBuilder) SetProtocol(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x18)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type ContainerBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	hostBuilder          HostBuilder
	containerAddrBuilder ContainerAddrBuilder
}

func NewContainerBuilder(writer io.Writer) *ContainerBuilder {
	return &ContainerBuilder{
		writer: writer,
	}
}
func (x *ContainerBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetImage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetCpuLimit(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x29)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetMemoryLimit(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetState(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x40)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ContainerBuilder) SetHealth(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x48)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ContainerBuilder) SetCreated(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x50)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetRbps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x59)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetWbps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x61)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x68)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetNetRcvdPs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x71)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetNetSentPs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x79)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetNetRcvdBps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x81)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetNetSentBps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x89)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetUserPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x91)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetSystemPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x99)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetTotalPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa1)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetMemRss(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetMemCache(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xb0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xba)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerBuilder) SetStarted(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xc0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xca)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xd2)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) AddAddresses(cb func(w *ContainerAddrBuilder)) {
	x.buf.Reset()
	x.containerAddrBuilder.writer = &x.buf
	x.containerAddrBuilder.scratch = x.scratch
	cb(&x.containerAddrBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xda)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerBuilder) SetThreadCount(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xe0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetThreadLimit(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xe8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetMemUsage(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xf0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetCpuUsageNs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xf9)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerBuilder) SetMemAccounted(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x100)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ProcessStatBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	memoryStatBuilder      MemoryStatBuilder
	cPUStatBuilder         CPUStatBuilder
	iOStatBuilder          IOStatBuilder
	processNetworksBuilder ProcessNetworksBuilder
}

func NewProcessStatBuilder(writer io.Writer) *ProcessStatBuilder {
	return &ProcessStatBuilder{
		writer: writer,
	}
}
func (x *ProcessStatBuilder) SetPid(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetCreateTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetMemory(cb func(w *MemoryStatBuilder)) {
	x.buf.Reset()
	x.memoryStatBuilder.writer = &x.buf
	x.memoryStatBuilder.scratch = x.scratch
	cb(&x.memoryStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessStatBuilder) SetCpu(cb func(w *CPUStatBuilder)) {
	x.buf.Reset()
	x.cPUStatBuilder.writer = &x.buf
	x.cPUStatBuilder.scratch = x.scratch
	cb(&x.cPUStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessStatBuilder) SetNice(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetThreads(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetOpenFdCount(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerState(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x58)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ProcessStatBuilder) SetProcessState(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x60)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ProcessStatBuilder) SetIoStat(cb func(w *IOStatBuilder)) {
	x.buf.Reset()
	x.iOStatBuilder.writer = &x.buf
	x.iOStatBuilder.scratch = x.scratch
	cb(&x.iOStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x9a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessStatBuilder) SetNetworks(cb func(w *ProcessNetworksBuilder)) {
	x.buf.Reset()
	x.processNetworksBuilder.writer = &x.buf
	x.processNetworksBuilder.scratch = x.scratch
	cb(&x.processNetworksBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xe2)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessStatBuilder) SetContainerHealth(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x78)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ProcessStatBuilder) SetContainerRbps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x81)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerWbps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x89)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x90)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerNetRcvdPs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa1)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerNetSentPs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa9)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerNetRcvdBps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xb1)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetContainerNetSentBps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xb9)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetVoluntaryCtxSwitches(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xc0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetInvoluntaryCtxSwitches(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xc8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcessStatBuilder) SetByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xd2)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ProcessStatBuilder) SetContainerByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xda)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ProcStatsWithPermBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewProcStatsWithPermBuilder(writer io.Writer) *ProcStatsWithPermBuilder {
	return &ProcStatsWithPermBuilder{
		writer: writer,
	}
}
func (x *ProcStatsWithPermBuilder) SetOpenFDCount(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcStatsWithPermBuilder) SetReadCount(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcStatsWithPermBuilder) SetWriteCount(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcStatsWithPermBuilder) SetReadBytes(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcStatsWithPermBuilder) SetWriteBytes(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ProcStatsWithPermByPIDBuilder struct {
	writer                                        io.Writer
	buf                                           bytes.Buffer
	scratch                                       []byte
	procStatsWithPermByPID_StatsByPIDEntryBuilder ProcStatsWithPermByPID_StatsByPIDEntryBuilder
}

func NewProcStatsWithPermByPIDBuilder(writer io.Writer) *ProcStatsWithPermByPIDBuilder {
	return &ProcStatsWithPermByPIDBuilder{
		writer: writer,
	}
}
func (x *ProcStatsWithPermByPIDBuilder) AddStatsByPID(cb func(w *ProcStatsWithPermByPID_StatsByPIDEntryBuilder)) {
	x.buf.Reset()
	x.procStatsWithPermByPID_StatsByPIDEntryBuilder.writer = &x.buf
	x.procStatsWithPermByPID_StatsByPIDEntryBuilder.scratch = x.scratch
	cb(&x.procStatsWithPermByPID_StatsByPIDEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ProcStatsWithPermByPID_StatsByPIDEntryBuilder struct {
	writer                   io.Writer
	buf                      bytes.Buffer
	scratch                  []byte
	procStatsWithPermBuilder ProcStatsWithPermBuilder
}

func NewProcStatsWithPermByPID_StatsByPIDEntryBuilder(writer io.Writer) *ProcStatsWithPermByPID_StatsByPIDEntryBuilder {
	return &ProcStatsWithPermByPID_StatsByPIDEntryBuilder{
		writer: writer,
	}
}
func (x *ProcStatsWithPermByPID_StatsByPIDEntryBuilder) SetKey(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ProcStatsWithPermByPID_StatsByPIDEntryBuilder) SetValue(cb func(w *ProcStatsWithPermBuilder)) {
	x.buf.Reset()
	x.procStatsWithPermBuilder.writer = &x.buf
	x.procStatsWithPermBuilder.scratch = x.scratch
	cb(&x.procStatsWithPermBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ContainerStatBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewContainerStatBuilder(writer io.Writer) *ContainerStatBuilder {
	return &ContainerStatBuilder{
		writer: writer,
	}
}
func (x *ContainerStatBuilder) SetId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetUserPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetSystemPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x19)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetTotalPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x21)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetCpuLimit(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x29)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetMemRss(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetMemCache(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetMemLimit(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetRbps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x49)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetWbps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x51)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetNetRcvdPs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x59)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetNetSentPs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x61)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetNetRcvdBps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x69)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetNetSentBps(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x71)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetState(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x78)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ContainerStatBuilder) SetHealth(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x80)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *ContainerStatBuilder) SetKey(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x88)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetStarted(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x90)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetByteKey(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x9a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerStatBuilder) SetThreadCount(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetThreadLimit(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetMemUsage(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xb0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetCpuUsageNs(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xb9)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatBuilder) SetMemAccounted(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xc0)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type SystemInfoBuilder struct {
	writer         io.Writer
	buf            bytes.Buffer
	scratch        []byte
	oSInfoBuilder  OSInfoBuilder
	cPUInfoBuilder CPUInfoBuilder
}

func NewSystemInfoBuilder(writer io.Writer) *SystemInfoBuilder {
	return &SystemInfoBuilder{
		writer: writer,
	}
}
func (x *SystemInfoBuilder) SetUuid(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *SystemInfoBuilder) SetOs(cb func(w *OSInfoBuilder)) {
	x.buf.Reset()
	x.oSInfoBuilder.writer = &x.buf
	x.oSInfoBuilder.scratch = x.scratch
	cb(&x.oSInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *SystemInfoBuilder) AddCpus(cb func(w *CPUInfoBuilder)) {
	x.buf.Reset()
	x.cPUInfoBuilder.writer = &x.buf
	x.cPUInfoBuilder.scratch = x.scratch
	cb(&x.cPUInfoBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *SystemInfoBuilder) SetTotalMemory(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type OSInfoBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewOSInfoBuilder(writer io.Writer) *OSInfoBuilder {
	return &OSInfoBuilder{
		writer: writer,
	}
}
func (x *OSInfoBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *OSInfoBuilder) SetPlatform(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *OSInfoBuilder) SetFamily(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *OSInfoBuilder) SetVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *OSInfoBuilder) SetKernelVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type IOStatBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewIOStatBuilder(writer io.Writer) *IOStatBuilder {
	return &IOStatBuilder{
		writer: writer,
	}
}
func (x *IOStatBuilder) SetReadRate(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x9)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *IOStatBuilder) SetWriteRate(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *IOStatBuilder) SetReadBytesRate(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x19)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *IOStatBuilder) SetWriteBytesRate(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x21)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}

type MemoryStatBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewMemoryStatBuilder(writer io.Writer) *MemoryStatBuilder {
	return &MemoryStatBuilder{
		writer: writer,
	}
}
func (x *MemoryStatBuilder) SetRss(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetVms(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetSwap(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetShared(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetText(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetLib(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetData(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MemoryStatBuilder) SetDirty(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type CPUStatBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	singleCPUStatBuilder SingleCPUStatBuilder
}

func NewCPUStatBuilder(writer io.Writer) *CPUStatBuilder {
	return &CPUStatBuilder{
		writer: writer,
	}
}
func (x *CPUStatBuilder) SetLastCpu(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) SetTotalPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) SetUserPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x19)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) SetSystemPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x21)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) SetNumThreads(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) AddCpus(cb func(w *SingleCPUStatBuilder)) {
	x.buf.Reset()
	x.singleCPUStatBuilder.writer = &x.buf
	x.singleCPUStatBuilder.scratch = x.scratch
	cb(&x.singleCPUStatBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CPUStatBuilder) SetNice(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) SetUserTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CPUStatBuilder) SetSystemTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type SingleCPUStatBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewSingleCPUStatBuilder(writer io.Writer) *SingleCPUStatBuilder {
	return &SingleCPUStatBuilder{
		writer: writer,
	}
}
func (x *SingleCPUStatBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *SingleCPUStatBuilder) SetTotalPct(v float32) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11)
	x.scratch = protowire.AppendFixed32(x.scratch, math.Float32bits(v))
	x.writer.Write(x.scratch)
}

type CPUInfoBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCPUInfoBuilder(writer io.Writer) *CPUInfoBuilder {
	return &CPUInfoBuilder{
		writer: writer,
	}
}
func (x *CPUInfoBuilder) SetNumber(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetVendor(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetFamily(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetModel(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetPhysicalId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetCoreId(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetCores(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetMhz(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CPUInfoBuilder) SetCacheSize(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type HostBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewHostBuilder(writer io.Writer) *HostBuilder {
	return &HostBuilder{
		writer: writer,
	}
}
func (x *HostBuilder) SetId(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) SetOrgId(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) AddAllTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) SetNumCpus(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) SetTotalMemory(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) SetTagIndex(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HostBuilder) SetTagsModified(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x50)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ClusterBuilder struct {
	writer                                io.Writer
	buf                                   bytes.Buffer
	scratch                               []byte
	cluster_KubeletVersionsEntryBuilder   Cluster_KubeletVersionsEntryBuilder
	cluster_ApiServerVersionsEntryBuilder Cluster_ApiServerVersionsEntryBuilder
	resourceMetricsBuilder                ResourceMetricsBuilder
}

func NewClusterBuilder(writer io.Writer) *ClusterBuilder {
	return &ClusterBuilder{
		writer: writer,
	}
}
func (x *ClusterBuilder) SetNodeCount(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) AddKubeletVersions(cb func(w *Cluster_KubeletVersionsEntryBuilder)) {
	x.buf.Reset()
	x.cluster_KubeletVersionsEntryBuilder.writer = &x.buf
	x.cluster_KubeletVersionsEntryBuilder.scratch = x.scratch
	cb(&x.cluster_KubeletVersionsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterBuilder) AddApiServerVersions(cb func(w *Cluster_ApiServerVersionsEntryBuilder)) {
	x.buf.Reset()
	x.cluster_ApiServerVersionsEntryBuilder.writer = &x.buf
	x.cluster_ApiServerVersionsEntryBuilder.scratch = x.scratch
	cb(&x.cluster_ApiServerVersionsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterBuilder) SetPodCapacity(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetPodAllocatable(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetMemoryAllocatable(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetMemoryCapacity(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetCpuAllocatable(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetCpuCapacity(v uint64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetResourceVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetCreationTimestamp(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x58)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x62)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x6a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type Cluster_KubeletVersionsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCluster_KubeletVersionsEntryBuilder(writer io.Writer) *Cluster_KubeletVersionsEntryBuilder {
	return &Cluster_KubeletVersionsEntryBuilder{
		writer: writer,
	}
}
func (x *Cluster_KubeletVersionsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *Cluster_KubeletVersionsEntryBuilder) SetValue(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type Cluster_ApiServerVersionsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCluster_ApiServerVersionsEntryBuilder(writer io.Writer) *Cluster_ApiServerVersionsEntryBuilder {
	return &Cluster_ApiServerVersionsEntryBuilder{
		writer: writer,
	}
}
func (x *Cluster_ApiServerVersionsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *Cluster_ApiServerVersionsEntryBuilder) SetValue(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type MetadataBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	ownerReferenceBuilder OwnerReferenceBuilder
}

func NewMetadataBuilder(writer io.Writer) *MetadataBuilder {
	return &MetadataBuilder{
		writer: writer,
	}
}
func (x *MetadataBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) SetNamespace(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) SetUid(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) SetCreationTimestamp(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) SetDeletionTimestamp(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) AddLabels(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) AddAnnotations(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) AddOwnerReferences(cb func(w *OwnerReferenceBuilder)) {
	x.buf.Reset()
	x.ownerReferenceBuilder.writer = &x.buf
	x.ownerReferenceBuilder.scratch = x.scratch
	cb(&x.ownerReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *MetadataBuilder) SetResourceVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *MetadataBuilder) AddFinalizers(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type OwnerReferenceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewOwnerReferenceBuilder(writer io.Writer) *OwnerReferenceBuilder {
	return &OwnerReferenceBuilder{
		writer: writer,
	}
}
func (x *OwnerReferenceBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *OwnerReferenceBuilder) SetUid(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *OwnerReferenceBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ObjectReferenceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewObjectReferenceBuilder(writer io.Writer) *ObjectReferenceBuilder {
	return &ObjectReferenceBuilder{
		writer: writer,
	}
}
func (x *ObjectReferenceBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ObjectReferenceBuilder) SetNamespace(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ObjectReferenceBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ObjectReferenceBuilder) SetUid(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ObjectReferenceBuilder) SetApiVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ObjectReferenceBuilder) SetResourceVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ObjectReferenceBuilder) SetFieldPath(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ServicePortBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewServicePortBuilder(writer io.Writer) *ServicePortBuilder {
	return &ServicePortBuilder{
		writer: writer,
	}
}
func (x *ServicePortBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServicePortBuilder) SetProtocol(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServicePortBuilder) SetPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ServicePortBuilder) SetTargetPort(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServicePortBuilder) SetNodePort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ServiceSessionAffinityConfigBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewServiceSessionAffinityConfigBuilder(writer io.Writer) *ServiceSessionAffinityConfigBuilder {
	return &ServiceSessionAffinityConfigBuilder{
		writer: writer,
	}
}
func (x *ServiceSessionAffinityConfigBuilder) SetClientIPTimeoutSeconds(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type NodeBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	metadataBuilder        MetadataBuilder
	taintBuilder           TaintBuilder
	nodeStatusBuilder      NodeStatusBuilder
	hostBuilder            HostBuilder
	resourceMetricsBuilder ResourceMetricsBuilder
}

func NewNodeBuilder(writer io.Writer) *NodeBuilder {
	return &NodeBuilder{
		writer: writer,
	}
}
func (x *NodeBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeBuilder) SetPodCIDR(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeBuilder) AddPodCIDRs(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeBuilder) SetUnschedulable(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *NodeBuilder) AddTaints(cb func(w *TaintBuilder)) {
	x.buf.Reset()
	x.taintBuilder.writer = &x.buf
	x.taintBuilder.scratch = x.scratch
	cb(&x.taintBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeBuilder) SetStatus(cb func(w *NodeStatusBuilder)) {
	x.buf.Reset()
	x.nodeStatusBuilder.writer = &x.buf
	x.nodeStatusBuilder.scratch = x.scratch
	cb(&x.nodeStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeBuilder) AddRoles(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeBuilder) SetProviderID(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x5a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x62)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type NodeStatusBuilder struct {
	writer                               io.Writer
	buf                                  bytes.Buffer
	scratch                              []byte
	nodeStatus_CapacityEntryBuilder      NodeStatus_CapacityEntryBuilder
	nodeStatus_AllocatableEntryBuilder   NodeStatus_AllocatableEntryBuilder
	nodeStatus_NodeAddressesEntryBuilder NodeStatus_NodeAddressesEntryBuilder
	nodeConditionBuilder                 NodeConditionBuilder
	containerImageBuilder                ContainerImageBuilder
}

func NewNodeStatusBuilder(writer io.Writer) *NodeStatusBuilder {
	return &NodeStatusBuilder{
		writer: writer,
	}
}
func (x *NodeStatusBuilder) AddCapacity(cb func(w *NodeStatus_CapacityEntryBuilder)) {
	x.buf.Reset()
	x.nodeStatus_CapacityEntryBuilder.writer = &x.buf
	x.nodeStatus_CapacityEntryBuilder.scratch = x.scratch
	cb(&x.nodeStatus_CapacityEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeStatusBuilder) AddAllocatable(cb func(w *NodeStatus_AllocatableEntryBuilder)) {
	x.buf.Reset()
	x.nodeStatus_AllocatableEntryBuilder.writer = &x.buf
	x.nodeStatus_AllocatableEntryBuilder.scratch = x.scratch
	cb(&x.nodeStatus_AllocatableEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeStatusBuilder) AddNodeAddresses(cb func(w *NodeStatus_NodeAddressesEntryBuilder)) {
	x.buf.Reset()
	x.nodeStatus_NodeAddressesEntryBuilder.writer = &x.buf
	x.nodeStatus_NodeAddressesEntryBuilder.scratch = x.scratch
	cb(&x.nodeStatus_NodeAddressesEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeStatusBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) SetKubeletVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) AddConditions(cb func(w *NodeConditionBuilder)) {
	x.buf.Reset()
	x.nodeConditionBuilder.writer = &x.buf
	x.nodeConditionBuilder.scratch = x.scratch
	cb(&x.nodeConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeStatusBuilder) AddImages(cb func(w *ContainerImageBuilder)) {
	x.buf.Reset()
	x.containerImageBuilder.writer = &x.buf
	x.containerImageBuilder.scratch = x.scratch
	cb(&x.containerImageBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeStatusBuilder) SetKubeProxyVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) SetOperatingSystem(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) SetArchitecture(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) SetKernelVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x5a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) SetOsImage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x62)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatusBuilder) SetContainerRuntimeVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x6a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NodeStatus_CapacityEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewNodeStatus_CapacityEntryBuilder(writer io.Writer) *NodeStatus_CapacityEntryBuilder {
	return &NodeStatus_CapacityEntryBuilder{
		writer: writer,
	}
}
func (x *NodeStatus_CapacityEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatus_CapacityEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type NodeStatus_AllocatableEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewNodeStatus_AllocatableEntryBuilder(writer io.Writer) *NodeStatus_AllocatableEntryBuilder {
	return &NodeStatus_AllocatableEntryBuilder{
		writer: writer,
	}
}
func (x *NodeStatus_AllocatableEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatus_AllocatableEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type NodeStatus_NodeAddressesEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewNodeStatus_NodeAddressesEntryBuilder(writer io.Writer) *NodeStatus_NodeAddressesEntryBuilder {
	return &NodeStatus_NodeAddressesEntryBuilder{
		writer: writer,
	}
}
func (x *NodeStatus_NodeAddressesEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeStatus_NodeAddressesEntryBuilder) SetValue(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NodeConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewNodeConditionBuilder(writer io.Writer) *NodeConditionBuilder {
	return &NodeConditionBuilder{
		writer: writer,
	}
}
func (x *NodeConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *NodeConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NodeConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ContainerImageBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewContainerImageBuilder(writer io.Writer) *ContainerImageBuilder {
	return &ContainerImageBuilder{
		writer: writer,
	}
}
func (x *ContainerImageBuilder) AddNames(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerImageBuilder) SetSizeBytes(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type TaintBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewTaintBuilder(writer io.Writer) *TaintBuilder {
	return &TaintBuilder{
		writer: writer,
	}
}
func (x *TaintBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *TaintBuilder) SetValue(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *TaintBuilder) SetEffect(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *TaintBuilder) SetTimeAdded(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ServiceSpecBuilder struct {
	writer                              io.Writer
	buf                                 bytes.Buffer
	scratch                             []byte
	servicePortBuilder                  ServicePortBuilder
	labelSelectorRequirementBuilder     LabelSelectorRequirementBuilder
	serviceSessionAffinityConfigBuilder ServiceSessionAffinityConfigBuilder
}

func NewServiceSpecBuilder(writer io.Writer) *ServiceSpecBuilder {
	return &ServiceSpecBuilder{
		writer: writer,
	}
}
func (x *ServiceSpecBuilder) AddPorts(cb func(w *ServicePortBuilder)) {
	x.buf.Reset()
	x.servicePortBuilder.writer = &x.buf
	x.servicePortBuilder.scratch = x.scratch
	cb(&x.servicePortBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceSpecBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceSpecBuilder) SetClusterIP(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) AddExternalIPs(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetSessionAffinity(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetLoadBalancerIP(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) AddLoadBalancerSourceRanges(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetExternalName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetExternalTrafficPolicy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x52)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetHealthCheckNodePort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x58)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ServiceSpecBuilder) SetPublishNotReadyAddresses(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x60)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *ServiceSpecBuilder) SetSessionAffinityConfig(cb func(w *ServiceSessionAffinityConfigBuilder)) {
	x.buf.Reset()
	x.serviceSessionAffinityConfigBuilder.writer = &x.buf
	x.serviceSessionAffinityConfigBuilder.scratch = x.scratch
	cb(&x.serviceSessionAffinityConfigBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x6a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceSpecBuilder) SetIpFamily(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x72)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ServiceStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewServiceStatusBuilder(writer io.Writer) *ServiceStatusBuilder {
	return &ServiceStatusBuilder{
		writer: writer,
	}
}
func (x *ServiceStatusBuilder) AddLoadBalancerIngress(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ServiceBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	metadataBuilder        MetadataBuilder
	serviceSpecBuilder     ServiceSpecBuilder
	serviceStatusBuilder   ServiceStatusBuilder
	resourceMetricsBuilder ResourceMetricsBuilder
}

func NewServiceBuilder(writer io.Writer) *ServiceBuilder {
	return &ServiceBuilder{
		writer: writer,
	}
}
func (x *ServiceBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceBuilder) SetSpec(cb func(w *ServiceSpecBuilder)) {
	x.buf.Reset()
	x.serviceSpecBuilder.writer = &x.buf
	x.serviceSpecBuilder.scratch = x.scratch
	cb(&x.serviceSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceBuilder) SetStatus(cb func(w *ServiceStatusBuilder)) {
	x.buf.Reset()
	x.serviceStatusBuilder.writer = &x.buf
	x.serviceStatusBuilder.scratch = x.scratch
	cb(&x.serviceStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ServiceBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type DeploymentConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewDeploymentConditionBuilder(writer io.Writer) *DeploymentConditionBuilder {
	return &DeploymentConditionBuilder{
		writer: writer,
	}
}
func (x *DeploymentConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentConditionBuilder) SetLastUpdateTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type DeploymentBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	metadataBuilder                 MetadataBuilder
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	resourceRequirementsBuilder     ResourceRequirementsBuilder
	resourceMetricsBuilder          ResourceMetricsBuilder
	deploymentConditionBuilder      DeploymentConditionBuilder
}

func NewDeploymentBuilder(writer io.Writer) *DeploymentBuilder {
	return &DeploymentBuilder{
		writer: writer,
	}
}
func (x *DeploymentBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DeploymentBuilder) SetReplicasDesired(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetDeploymentStrategy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetMaxUnavailable(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetMaxSurge(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetPaused(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x30)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *DeploymentBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DeploymentBuilder) SetReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetUpdatedReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x48)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetReadyReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x50)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetAvailableReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x58)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetUnavailableReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x60)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetConditionMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x6a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x82)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DeploymentBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x72)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DeploymentBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x7a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DeploymentBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DeploymentBuilder) AddConditions(cb func(w *DeploymentConditionBuilder)) {
	x.buf.Reset()
	x.deploymentConditionBuilder.writer = &x.buf
	x.deploymentConditionBuilder.scratch = x.scratch
	cb(&x.deploymentConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x92)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ReplicaSetConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewReplicaSetConditionBuilder(writer io.Writer) *ReplicaSetConditionBuilder {
	return &ReplicaSetConditionBuilder{
		writer: writer,
	}
}
func (x *ReplicaSetConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ReplicaSetBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	metadataBuilder                 MetadataBuilder
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	resourceRequirementsBuilder     ResourceRequirementsBuilder
	resourceMetricsBuilder          ResourceMetricsBuilder
	replicaSetConditionBuilder      ReplicaSetConditionBuilder
}

func NewReplicaSetBuilder(writer io.Writer) *ReplicaSetBuilder {
	return &ReplicaSetBuilder{
		writer: writer,
	}
}
func (x *ReplicaSetBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ReplicaSetBuilder) SetReplicasDesired(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ReplicaSetBuilder) SetReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetBuilder) SetFullyLabeledReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetBuilder) SetReadyReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetBuilder) SetAvailableReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ReplicaSetBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ReplicaSetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ReplicaSetBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x5a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ReplicaSetBuilder) AddConditions(cb func(w *ReplicaSetConditionBuilder)) {
	x.buf.Reset()
	x.replicaSetConditionBuilder.writer = &x.buf
	x.replicaSetConditionBuilder.scratch = x.scratch
	cb(&x.replicaSetConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x62)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type LabelSelectorRequirementBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewLabelSelectorRequirementBuilder(writer io.Writer) *LabelSelectorRequirementBuilder {
	return &LabelSelectorRequirementBuilder{
		writer: writer,
	}
}
func (x *LabelSelectorRequirementBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LabelSelectorRequirementBuilder) SetOperator(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LabelSelectorRequirementBuilder) AddValues(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PodBuilder struct {
	writer                      io.Writer
	buf                         bytes.Buffer
	scratch                     []byte
	metadataBuilder             MetadataBuilder
	containerStatusBuilder      ContainerStatusBuilder
	hostBuilder                 HostBuilder
	resourceRequirementsBuilder ResourceRequirementsBuilder
	resourceMetricsBuilder      ResourceMetricsBuilder
	podConditionBuilder         PodConditionBuilder
}

func NewPodBuilder(writer io.Writer) *PodBuilder {
	return &PodBuilder{
		writer: writer,
	}
}
func (x *PodBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) SetIP(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetNominatedNodeName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetNodeName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetPhase(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetRestartCount(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) AddContainerStatuses(cb func(w *ContainerStatusBuilder)) {
	x.buf.Reset()
	x.containerStatusBuilder.writer = &x.buf
	x.containerStatusBuilder.scratch = x.scratch
	cb(&x.containerStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) AddInitContainerStatuses(cb func(w *ContainerStatusBuilder)) {
	x.buf.Reset()
	x.containerStatusBuilder.writer = &x.buf
	x.containerStatusBuilder.scratch = x.scratch
	cb(&x.containerStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x72)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) SetConditionMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x4a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x5a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetHost(cb func(w *HostBuilder)) {
	x.buf.Reset()
	x.hostBuilder.writer = &x.buf
	x.hostBuilder.scratch = x.scratch
	cb(&x.hostBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x62)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x6a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) SetQOSClass(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x7a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetPriorityClass(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x82)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x8a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PodBuilder) SetStartTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x90)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) SetScheduledTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x98)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodBuilder) AddConditions(cb func(w *PodConditionBuilder)) {
	x.buf.Reset()
	x.podConditionBuilder.writer = &x.buf
	x.podConditionBuilder.scratch = x.scratch
	cb(&x.podConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa2)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PodConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPodConditionBuilder(writer io.Writer) *PodConditionBuilder {
	return &PodConditionBuilder{
		writer: writer,
	}
}
func (x *PodConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodConditionBuilder) SetLastProbeTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PodConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PodConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ContainerStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewContainerStatusBuilder(writer io.Writer) *ContainerStatusBuilder {
	return &ContainerStatusBuilder{
		writer: writer,
	}
}
func (x *ContainerStatusBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerStatusBuilder) SetContainerID(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerStatusBuilder) SetReady(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x18)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *ContainerStatusBuilder) SetRestartCount(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ContainerStatusBuilder) SetState(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerStatusBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ManifestBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewManifestBuilder(writer io.Writer) *ManifestBuilder {
	return &ManifestBuilder{
		writer: writer,
	}
}
func (x *ManifestBuilder) SetType(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) SetResourceVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) SetUid(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) SetContent(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ManifestBuilder) SetContentType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ManifestBuilder) SetVersion(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NamespaceConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewNamespaceConditionBuilder(writer io.Writer) *NamespaceConditionBuilder {
	return &NamespaceConditionBuilder{
		writer: writer,
	}
}
func (x *NamespaceConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NamespaceConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NamespaceConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *NamespaceConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NamespaceConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NamespaceBuilder struct {
	writer                    io.Writer
	buf                       bytes.Buffer
	scratch                   []byte
	metadataBuilder           MetadataBuilder
	namespaceConditionBuilder NamespaceConditionBuilder
}

func NewNamespaceBuilder(writer io.Writer) *NamespaceBuilder {
	return &NamespaceBuilder{
		writer: writer,
	}
}
func (x *NamespaceBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NamespaceBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NamespaceBuilder) SetConditionMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NamespaceBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NamespaceBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *NamespaceBuilder) AddConditions(cb func(w *NamespaceConditionBuilder)) {
	x.buf.Reset()
	x.namespaceConditionBuilder.writer = &x.buf
	x.namespaceConditionBuilder.scratch = x.scratch
	cb(&x.namespaceConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ResourceRequirementsBuilder struct {
	writer                                    io.Writer
	buf                                       bytes.Buffer
	scratch                                   []byte
	resourceRequirements_LimitsEntryBuilder   ResourceRequirements_LimitsEntryBuilder
	resourceRequirements_RequestsEntryBuilder ResourceRequirements_RequestsEntryBuilder
}

func NewResourceRequirementsBuilder(writer io.Writer) *ResourceRequirementsBuilder {
	return &ResourceRequirementsBuilder{
		writer: writer,
	}
}
func (x *ResourceRequirementsBuilder) AddLimits(cb func(w *ResourceRequirements_LimitsEntryBuilder)) {
	x.buf.Reset()
	x.resourceRequirements_LimitsEntryBuilder.writer = &x.buf
	x.resourceRequirements_LimitsEntryBuilder.scratch = x.scratch
	cb(&x.resourceRequirements_LimitsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ResourceRequirementsBuilder) AddRequests(cb func(w *ResourceRequirements_RequestsEntryBuilder)) {
	x.buf.Reset()
	x.resourceRequirements_RequestsEntryBuilder.writer = &x.buf
	x.resourceRequirements_RequestsEntryBuilder.scratch = x.scratch
	cb(&x.resourceRequirements_RequestsEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ResourceRequirementsBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceRequirementsBuilder) SetType(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}

type ResourceRequirements_LimitsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewResourceRequirements_LimitsEntryBuilder(writer io.Writer) *ResourceRequirements_LimitsEntryBuilder {
	return &ResourceRequirements_LimitsEntryBuilder{
		writer: writer,
	}
}
func (x *ResourceRequirements_LimitsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceRequirements_LimitsEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ResourceRequirements_RequestsEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewResourceRequirements_RequestsEntryBuilder(writer io.Writer) *ResourceRequirements_RequestsEntryBuilder {
	return &ResourceRequirements_RequestsEntryBuilder{
		writer: writer,
	}
}
func (x *ResourceRequirements_RequestsEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceRequirements_RequestsEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ResourceMetricsBuilder struct {
	writer                                   io.Writer
	buf                                      bytes.Buffer
	scratch                                  []byte
	resourceMetrics_MetricValuesEntryBuilder ResourceMetrics_MetricValuesEntryBuilder
}

func NewResourceMetricsBuilder(writer io.Writer) *ResourceMetricsBuilder {
	return &ResourceMetricsBuilder{
		writer: writer,
	}
}
func (x *ResourceMetricsBuilder) AddMetricValues(cb func(w *ResourceMetrics_MetricValuesEntryBuilder)) {
	x.buf.Reset()
	x.resourceMetrics_MetricValuesEntryBuilder.writer = &x.buf
	x.resourceMetrics_MetricValuesEntryBuilder.scratch = x.scratch
	cb(&x.resourceMetrics_MetricValuesEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ResourceMetrics_MetricValuesEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewResourceMetrics_MetricValuesEntryBuilder(writer io.Writer) *ResourceMetrics_MetricValuesEntryBuilder {
	return &ResourceMetrics_MetricValuesEntryBuilder{
		writer: writer,
	}
}
func (x *ResourceMetrics_MetricValuesEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceMetrics_MetricValuesEntryBuilder) SetValue(v float64) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11)
	x.scratch = protowire.AppendFixed64(x.scratch, math.Float64bits(v))
	x.writer.Write(x.scratch)
}

type JobSpecBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	resourceRequirementsBuilder     ResourceRequirementsBuilder
}

func NewJobSpecBuilder(writer io.Writer) *JobSpecBuilder {
	return &JobSpecBuilder{
		writer: writer,
	}
}
func (x *JobSpecBuilder) SetParallelism(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobSpecBuilder) SetCompletions(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobSpecBuilder) SetActiveDeadlineSeconds(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobSpecBuilder) SetBackoffLimit(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobSpecBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *JobSpecBuilder) SetManualSelector(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x30)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *JobSpecBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type JobStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewJobStatusBuilder(writer io.Writer) *JobStatusBuilder {
	return &JobStatusBuilder{
		writer: writer,
	}
}
func (x *JobStatusBuilder) SetConditionMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *JobStatusBuilder) SetStartTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobStatusBuilder) SetCompletionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobStatusBuilder) SetActive(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobStatusBuilder) SetSucceeded(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobStatusBuilder) SetFailed(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type JobConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewJobConditionBuilder(writer io.Writer) *JobConditionBuilder {
	return &JobConditionBuilder{
		writer: writer,
	}
}
func (x *JobConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *JobConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *JobConditionBuilder) SetLastProbeTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *JobConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *JobConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type JobBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	metadataBuilder     MetadataBuilder
	jobSpecBuilder      JobSpecBuilder
	jobStatusBuilder    JobStatusBuilder
	jobConditionBuilder JobConditionBuilder
}

func NewJobBuilder(writer io.Writer) *JobBuilder {
	return &JobBuilder{
		writer: writer,
	}
}
func (x *JobBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *JobBuilder) SetSpec(cb func(w *JobSpecBuilder)) {
	x.buf.Reset()
	x.jobSpecBuilder.writer = &x.buf
	x.jobSpecBuilder.scratch = x.scratch
	cb(&x.jobSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *JobBuilder) SetStatus(cb func(w *JobStatusBuilder)) {
	x.buf.Reset()
	x.jobStatusBuilder.writer = &x.buf
	x.jobStatusBuilder.scratch = x.scratch
	cb(&x.jobStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *JobBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *JobBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *JobBuilder) AddConditions(cb func(w *JobConditionBuilder)) {
	x.buf.Reset()
	x.jobConditionBuilder.writer = &x.buf
	x.jobConditionBuilder.scratch = x.scratch
	cb(&x.jobConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CronJobSpecBuilder struct {
	writer                      io.Writer
	buf                         bytes.Buffer
	scratch                     []byte
	resourceRequirementsBuilder ResourceRequirementsBuilder
}

func NewCronJobSpecBuilder(writer io.Writer) *CronJobSpecBuilder {
	return &CronJobSpecBuilder{
		writer: writer,
	}
}
func (x *CronJobSpecBuilder) SetSchedule(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CronJobSpecBuilder) SetStartingDeadlineSeconds(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CronJobSpecBuilder) SetConcurrencyPolicy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CronJobSpecBuilder) SetSuspend(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *CronJobSpecBuilder) SetSuccessfulJobsHistoryLimit(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CronJobSpecBuilder) SetFailedJobsHistoryLimit(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *CronJobSpecBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CronJobStatusBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	objectReferenceBuilder ObjectReferenceBuilder
}

func NewCronJobStatusBuilder(writer io.Writer) *CronJobStatusBuilder {
	return &CronJobStatusBuilder{
		writer: writer,
	}
}
func (x *CronJobStatusBuilder) AddActive(cb func(w *ObjectReferenceBuilder)) {
	x.buf.Reset()
	x.objectReferenceBuilder.writer = &x.buf
	x.objectReferenceBuilder.scratch = x.scratch
	cb(&x.objectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CronJobStatusBuilder) SetLastScheduleTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type CronJobBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	metadataBuilder      MetadataBuilder
	cronJobSpecBuilder   CronJobSpecBuilder
	cronJobStatusBuilder CronJobStatusBuilder
}

func NewCronJobBuilder(writer io.Writer) *CronJobBuilder {
	return &CronJobBuilder{
		writer: writer,
	}
}
func (x *CronJobBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CronJobBuilder) SetSpec(cb func(w *CronJobSpecBuilder)) {
	x.buf.Reset()
	x.cronJobSpecBuilder.writer = &x.buf
	x.cronJobSpecBuilder.scratch = x.scratch
	cb(&x.cronJobSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CronJobBuilder) SetStatus(cb func(w *CronJobStatusBuilder)) {
	x.buf.Reset()
	x.cronJobStatusBuilder.writer = &x.buf
	x.cronJobStatusBuilder.scratch = x.scratch
	cb(&x.cronJobStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CronJobBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CronJobBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type DaemonSetSpecBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	resourceRequirementsBuilder     ResourceRequirementsBuilder
}

func NewDaemonSetSpecBuilder(writer io.Writer) *DaemonSetSpecBuilder {
	return &DaemonSetSpecBuilder{
		writer: writer,
	}
}
func (x *DaemonSetSpecBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DaemonSetSpecBuilder) SetDeploymentStrategy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DaemonSetSpecBuilder) SetMaxUnavailable(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DaemonSetSpecBuilder) SetMinReadySeconds(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetSpecBuilder) SetRevisionHistoryLimit(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x28)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetSpecBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type DaemonSetStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewDaemonSetStatusBuilder(writer io.Writer) *DaemonSetStatusBuilder {
	return &DaemonSetStatusBuilder{
		writer: writer,
	}
}
func (x *DaemonSetStatusBuilder) SetCurrentNumberScheduled(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetStatusBuilder) SetNumberMisscheduled(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetStatusBuilder) SetDesiredNumberScheduled(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetStatusBuilder) SetNumberReady(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetStatusBuilder) SetUpdatedNumberScheduled(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetStatusBuilder) SetNumberAvailable(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x38)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetStatusBuilder) SetNumberUnavailable(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x40)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type DaemonSetConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewDaemonSetConditionBuilder(writer io.Writer) *DaemonSetConditionBuilder {
	return &DaemonSetConditionBuilder{
		writer: writer,
	}
}
func (x *DaemonSetConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DaemonSetConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DaemonSetConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *DaemonSetConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DaemonSetConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type DaemonSetBuilder struct {
	writer                    io.Writer
	buf                       bytes.Buffer
	scratch                   []byte
	metadataBuilder           MetadataBuilder
	daemonSetSpecBuilder      DaemonSetSpecBuilder
	daemonSetStatusBuilder    DaemonSetStatusBuilder
	resourceMetricsBuilder    ResourceMetricsBuilder
	daemonSetConditionBuilder DaemonSetConditionBuilder
}

func NewDaemonSetBuilder(writer io.Writer) *DaemonSetBuilder {
	return &DaemonSetBuilder{
		writer: writer,
	}
}
func (x *DaemonSetBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DaemonSetBuilder) SetSpec(cb func(w *DaemonSetSpecBuilder)) {
	x.buf.Reset()
	x.daemonSetSpecBuilder.writer = &x.buf
	x.daemonSetSpecBuilder.scratch = x.scratch
	cb(&x.daemonSetSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DaemonSetBuilder) SetStatus(cb func(w *DaemonSetStatusBuilder)) {
	x.buf.Reset()
	x.daemonSetStatusBuilder.writer = &x.buf
	x.daemonSetStatusBuilder.scratch = x.scratch
	cb(&x.daemonSetStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DaemonSetBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DaemonSetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DaemonSetBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DaemonSetBuilder) AddConditions(cb func(w *DaemonSetConditionBuilder)) {
	x.buf.Reset()
	x.daemonSetConditionBuilder.writer = &x.buf
	x.daemonSetConditionBuilder.scratch = x.scratch
	cb(&x.daemonSetConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type StatefulSetSpecBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	resourceRequirementsBuilder     ResourceRequirementsBuilder
}

func NewStatefulSetSpecBuilder(writer io.Writer) *StatefulSetSpecBuilder {
	return &StatefulSetSpecBuilder{
		writer: writer,
	}
}
func (x *StatefulSetSpecBuilder) SetDesiredReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *StatefulSetSpecBuilder) AddSelectors(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StatefulSetSpecBuilder) SetServiceName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetSpecBuilder) SetPodManagementPolicy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetSpecBuilder) SetUpdateStrategy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetSpecBuilder) SetPartition(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x30)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *StatefulSetSpecBuilder) AddResourceRequirements(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type StatefulSetStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewStatefulSetStatusBuilder(writer io.Writer) *StatefulSetStatusBuilder {
	return &StatefulSetStatusBuilder{
		writer: writer,
	}
}
func (x *StatefulSetStatusBuilder) SetReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *StatefulSetStatusBuilder) SetReadyReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *StatefulSetStatusBuilder) SetCurrentReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *StatefulSetStatusBuilder) SetUpdatedReplicas(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type StatefulSetConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewStatefulSetConditionBuilder(writer io.Writer) *StatefulSetConditionBuilder {
	return &StatefulSetConditionBuilder{
		writer: writer,
	}
}
func (x *StatefulSetConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *StatefulSetConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type StatefulSetBuilder struct {
	writer                      io.Writer
	buf                         bytes.Buffer
	scratch                     []byte
	metadataBuilder             MetadataBuilder
	statefulSetSpecBuilder      StatefulSetSpecBuilder
	statefulSetStatusBuilder    StatefulSetStatusBuilder
	resourceMetricsBuilder      ResourceMetricsBuilder
	statefulSetConditionBuilder StatefulSetConditionBuilder
}

func NewStatefulSetBuilder(writer io.Writer) *StatefulSetBuilder {
	return &StatefulSetBuilder{
		writer: writer,
	}
}
func (x *StatefulSetBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StatefulSetBuilder) SetSpec(cb func(w *StatefulSetSpecBuilder)) {
	x.buf.Reset()
	x.statefulSetSpecBuilder.writer = &x.buf
	x.statefulSetSpecBuilder.scratch = x.scratch
	cb(&x.statefulSetSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StatefulSetBuilder) SetStatus(cb func(w *StatefulSetStatusBuilder)) {
	x.buf.Reset()
	x.statefulSetStatusBuilder.writer = &x.buf
	x.statefulSetStatusBuilder.scratch = x.scratch
	cb(&x.statefulSetStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StatefulSetBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StatefulSetBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *StatefulSetBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *StatefulSetBuilder) AddConditions(cb func(w *StatefulSetConditionBuilder)) {
	x.buf.Reset()
	x.statefulSetConditionBuilder.writer = &x.buf
	x.statefulSetConditionBuilder.scratch = x.scratch
	cb(&x.statefulSetConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PersistentVolumeBuilder struct {
	writer                        io.Writer
	buf                           bytes.Buffer
	scratch                       []byte
	metadataBuilder               MetadataBuilder
	persistentVolumeSpecBuilder   PersistentVolumeSpecBuilder
	persistentVolumeStatusBuilder PersistentVolumeStatusBuilder
}

func NewPersistentVolumeBuilder(writer io.Writer) *PersistentVolumeBuilder {
	return &PersistentVolumeBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeBuilder) SetSpec(cb func(w *PersistentVolumeSpecBuilder)) {
	x.buf.Reset()
	x.persistentVolumeSpecBuilder.writer = &x.buf
	x.persistentVolumeSpecBuilder.scratch = x.scratch
	cb(&x.persistentVolumeSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeBuilder) SetStatus(cb func(w *PersistentVolumeStatusBuilder)) {
	x.buf.Reset()
	x.persistentVolumeStatusBuilder.writer = &x.buf
	x.persistentVolumeStatusBuilder.scratch = x.scratch
	cb(&x.persistentVolumeStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PersistentVolumeSpecBuilder struct {
	writer                                    io.Writer
	buf                                       bytes.Buffer
	scratch                                   []byte
	persistentVolumeSpec_CapacityEntryBuilder PersistentVolumeSpec_CapacityEntryBuilder
	objectReferenceBuilder                    ObjectReferenceBuilder
	nodeSelectorTermBuilder                   NodeSelectorTermBuilder
	persistentVolumeSourceBuilder             PersistentVolumeSourceBuilder
}

func NewPersistentVolumeSpecBuilder(writer io.Writer) *PersistentVolumeSpecBuilder {
	return &PersistentVolumeSpecBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeSpecBuilder) AddCapacity(cb func(w *PersistentVolumeSpec_CapacityEntryBuilder)) {
	x.buf.Reset()
	x.persistentVolumeSpec_CapacityEntryBuilder.writer = &x.buf
	x.persistentVolumeSpec_CapacityEntryBuilder.scratch = x.scratch
	cb(&x.persistentVolumeSpec_CapacityEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSpecBuilder) SetPersistentVolumeType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpecBuilder) AddAccessModes(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpecBuilder) SetClaimRef(cb func(w *ObjectReferenceBuilder)) {
	x.buf.Reset()
	x.objectReferenceBuilder.writer = &x.buf
	x.objectReferenceBuilder.scratch = x.scratch
	cb(&x.objectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSpecBuilder) SetPersistentVolumeReclaimPolicy(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpecBuilder) SetStorageClassName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpecBuilder) AddMountOptions(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x3a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpecBuilder) SetVolumeMode(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x42)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpecBuilder) AddNodeAffinity(cb func(w *NodeSelectorTermBuilder)) {
	x.buf.Reset()
	x.nodeSelectorTermBuilder.writer = &x.buf
	x.nodeSelectorTermBuilder.scratch = x.scratch
	cb(&x.nodeSelectorTermBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x4a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSpecBuilder) SetPersistentVolumeSource(cb func(w *PersistentVolumeSourceBuilder)) {
	x.buf.Reset()
	x.persistentVolumeSourceBuilder.writer = &x.buf
	x.persistentVolumeSourceBuilder.scratch = x.scratch
	cb(&x.persistentVolumeSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PersistentVolumeSpec_CapacityEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPersistentVolumeSpec_CapacityEntryBuilder(writer io.Writer) *PersistentVolumeSpec_CapacityEntryBuilder {
	return &PersistentVolumeSpec_CapacityEntryBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeSpec_CapacityEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeSpec_CapacityEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type PersistentVolumeSourceBuilder struct {
	writer                                  io.Writer
	buf                                     bytes.Buffer
	scratch                                 []byte
	gCEPersistentDiskVolumeSourceBuilder    GCEPersistentDiskVolumeSourceBuilder
	aWSElasticBlockStoreVolumeSourceBuilder AWSElasticBlockStoreVolumeSourceBuilder
	azureFilePersistentVolumeSourceBuilder  AzureFilePersistentVolumeSourceBuilder
	azureDiskVolumeSourceBuilder            AzureDiskVolumeSourceBuilder
	cSIVolumeSourceBuilder                  CSIVolumeSourceBuilder
}

func NewPersistentVolumeSourceBuilder(writer io.Writer) *PersistentVolumeSourceBuilder {
	return &PersistentVolumeSourceBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeSourceBuilder) SetGcePersistentDisk(cb func(w *GCEPersistentDiskVolumeSourceBuilder)) {
	x.buf.Reset()
	x.gCEPersistentDiskVolumeSourceBuilder.writer = &x.buf
	x.gCEPersistentDiskVolumeSourceBuilder.scratch = x.scratch
	cb(&x.gCEPersistentDiskVolumeSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSourceBuilder) SetAwsElasticBlockStore(cb func(w *AWSElasticBlockStoreVolumeSourceBuilder)) {
	x.buf.Reset()
	x.aWSElasticBlockStoreVolumeSourceBuilder.writer = &x.buf
	x.aWSElasticBlockStoreVolumeSourceBuilder.scratch = x.scratch
	cb(&x.aWSElasticBlockStoreVolumeSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSourceBuilder) SetAzureFile(cb func(w *AzureFilePersistentVolumeSourceBuilder)) {
	x.buf.Reset()
	x.azureFilePersistentVolumeSourceBuilder.writer = &x.buf
	x.azureFilePersistentVolumeSourceBuilder.scratch = x.scratch
	cb(&x.azureFilePersistentVolumeSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSourceBuilder) SetAzureDisk(cb func(w *AzureDiskVolumeSourceBuilder)) {
	x.buf.Reset()
	x.azureDiskVolumeSourceBuilder.writer = &x.buf
	x.azureDiskVolumeSourceBuilder.scratch = x.scratch
	cb(&x.azureDiskVolumeSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeSourceBuilder) SetCsi(cb func(w *CSIVolumeSourceBuilder)) {
	x.buf.Reset()
	x.cSIVolumeSourceBuilder.writer = &x.buf
	x.cSIVolumeSourceBuilder.scratch = x.scratch
	cb(&x.cSIVolumeSourceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type GCEPersistentDiskVolumeSourceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewGCEPersistentDiskVolumeSourceBuilder(writer io.Writer) *GCEPersistentDiskVolumeSourceBuilder {
	return &GCEPersistentDiskVolumeSourceBuilder{
		writer: writer,
	}
}
func (x *GCEPersistentDiskVolumeSourceBuilder) SetPdName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *GCEPersistentDiskVolumeSourceBuilder) SetFsType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *GCEPersistentDiskVolumeSourceBuilder) SetPartition(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *GCEPersistentDiskVolumeSourceBuilder) SetReadOnly(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}

type AWSElasticBlockStoreVolumeSourceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewAWSElasticBlockStoreVolumeSourceBuilder(writer io.Writer) *AWSElasticBlockStoreVolumeSourceBuilder {
	return &AWSElasticBlockStoreVolumeSourceBuilder{
		writer: writer,
	}
}
func (x *AWSElasticBlockStoreVolumeSourceBuilder) SetVolumeID(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AWSElasticBlockStoreVolumeSourceBuilder) SetFsType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AWSElasticBlockStoreVolumeSourceBuilder) SetPartition(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *AWSElasticBlockStoreVolumeSourceBuilder) SetReadOnly(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}

type AzureFilePersistentVolumeSourceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewAzureFilePersistentVolumeSourceBuilder(writer io.Writer) *AzureFilePersistentVolumeSourceBuilder {
	return &AzureFilePersistentVolumeSourceBuilder{
		writer: writer,
	}
}
func (x *AzureFilePersistentVolumeSourceBuilder) SetSecretName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AzureFilePersistentVolumeSourceBuilder) SetShareName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AzureFilePersistentVolumeSourceBuilder) SetReadOnly(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x18)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *AzureFilePersistentVolumeSourceBuilder) SetSecretNamespace(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type AzureDiskVolumeSourceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewAzureDiskVolumeSourceBuilder(writer io.Writer) *AzureDiskVolumeSourceBuilder {
	return &AzureDiskVolumeSourceBuilder{
		writer: writer,
	}
}
func (x *AzureDiskVolumeSourceBuilder) SetDiskName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AzureDiskVolumeSourceBuilder) SetDiskURI(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AzureDiskVolumeSourceBuilder) SetCachingMode(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AzureDiskVolumeSourceBuilder) SetFsType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *AzureDiskVolumeSourceBuilder) SetReadOnly(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x28)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *AzureDiskVolumeSourceBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type CSIVolumeSourceBuilder struct {
	writer                                       io.Writer
	buf                                          bytes.Buffer
	scratch                                      []byte
	cSIVolumeSource_VolumeAttributesEntryBuilder CSIVolumeSource_VolumeAttributesEntryBuilder
	secretReferenceBuilder                       SecretReferenceBuilder
}

func NewCSIVolumeSourceBuilder(writer io.Writer) *CSIVolumeSourceBuilder {
	return &CSIVolumeSourceBuilder{
		writer: writer,
	}
}
func (x *CSIVolumeSourceBuilder) SetDriver(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CSIVolumeSourceBuilder) SetVolumeHandle(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CSIVolumeSourceBuilder) SetReadOnly(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x18)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *CSIVolumeSourceBuilder) SetFsType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CSIVolumeSourceBuilder) AddVolumeAttributes(cb func(w *CSIVolumeSource_VolumeAttributesEntryBuilder)) {
	x.buf.Reset()
	x.cSIVolumeSource_VolumeAttributesEntryBuilder.writer = &x.buf
	x.cSIVolumeSource_VolumeAttributesEntryBuilder.scratch = x.scratch
	cb(&x.cSIVolumeSource_VolumeAttributesEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CSIVolumeSourceBuilder) SetControllerPublishSecretRef(cb func(w *SecretReferenceBuilder)) {
	x.buf.Reset()
	x.secretReferenceBuilder.writer = &x.buf
	x.secretReferenceBuilder.scratch = x.scratch
	cb(&x.secretReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CSIVolumeSourceBuilder) SetNodeStageSecretRef(cb func(w *SecretReferenceBuilder)) {
	x.buf.Reset()
	x.secretReferenceBuilder.writer = &x.buf
	x.secretReferenceBuilder.scratch = x.scratch
	cb(&x.secretReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CSIVolumeSourceBuilder) SetNodePublishSecretRef(cb func(w *SecretReferenceBuilder)) {
	x.buf.Reset()
	x.secretReferenceBuilder.writer = &x.buf
	x.secretReferenceBuilder.scratch = x.scratch
	cb(&x.secretReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x42)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CSIVolumeSourceBuilder) SetControllerExpandSecretRef(cb func(w *SecretReferenceBuilder)) {
	x.buf.Reset()
	x.secretReferenceBuilder.writer = &x.buf
	x.secretReferenceBuilder.scratch = x.scratch
	cb(&x.secretReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x4a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *CSIVolumeSourceBuilder) SetNodeExpandSecretRef(cb func(w *SecretReferenceBuilder)) {
	x.buf.Reset()
	x.secretReferenceBuilder.writer = &x.buf
	x.secretReferenceBuilder.scratch = x.scratch
	cb(&x.secretReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x52)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type CSIVolumeSource_VolumeAttributesEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewCSIVolumeSource_VolumeAttributesEntryBuilder(writer io.Writer) *CSIVolumeSource_VolumeAttributesEntryBuilder {
	return &CSIVolumeSource_VolumeAttributesEntryBuilder{
		writer: writer,
	}
}
func (x *CSIVolumeSource_VolumeAttributesEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *CSIVolumeSource_VolumeAttributesEntryBuilder) SetValue(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type SecretReferenceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewSecretReferenceBuilder(writer io.Writer) *SecretReferenceBuilder {
	return &SecretReferenceBuilder{
		writer: writer,
	}
}
func (x *SecretReferenceBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *SecretReferenceBuilder) SetNamespace(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PersistentVolumeStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPersistentVolumeStatusBuilder(writer io.Writer) *PersistentVolumeStatusBuilder {
	return &PersistentVolumeStatusBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeStatusBuilder) SetPhase(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeStatusBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeStatusBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type NodeSelectorTermBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
}

func NewNodeSelectorTermBuilder(writer io.Writer) *NodeSelectorTermBuilder {
	return &NodeSelectorTermBuilder{
		writer: writer,
	}
}
func (x *NodeSelectorTermBuilder) AddMatchExpressions(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *NodeSelectorTermBuilder) AddMatchFields(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PersistentVolumeClaimBuilder struct {
	writer                             io.Writer
	buf                                bytes.Buffer
	scratch                            []byte
	metadataBuilder                    MetadataBuilder
	persistentVolumeClaimSpecBuilder   PersistentVolumeClaimSpecBuilder
	persistentVolumeClaimStatusBuilder PersistentVolumeClaimStatusBuilder
}

func NewPersistentVolumeClaimBuilder(writer io.Writer) *PersistentVolumeClaimBuilder {
	return &PersistentVolumeClaimBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeClaimBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimBuilder) SetSpec(cb func(w *PersistentVolumeClaimSpecBuilder)) {
	x.buf.Reset()
	x.persistentVolumeClaimSpecBuilder.writer = &x.buf
	x.persistentVolumeClaimSpecBuilder.scratch = x.scratch
	cb(&x.persistentVolumeClaimSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimBuilder) SetStatus(cb func(w *PersistentVolumeClaimStatusBuilder)) {
	x.buf.Reset()
	x.persistentVolumeClaimStatusBuilder.writer = &x.buf
	x.persistentVolumeClaimStatusBuilder.scratch = x.scratch
	cb(&x.persistentVolumeClaimStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PersistentVolumeClaimStatusBuilder struct {
	writer                                           io.Writer
	buf                                              bytes.Buffer
	scratch                                          []byte
	persistentVolumeClaimStatus_CapacityEntryBuilder PersistentVolumeClaimStatus_CapacityEntryBuilder
	persistentVolumeClaimConditionBuilder            PersistentVolumeClaimConditionBuilder
}

func NewPersistentVolumeClaimStatusBuilder(writer io.Writer) *PersistentVolumeClaimStatusBuilder {
	return &PersistentVolumeClaimStatusBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeClaimStatusBuilder) SetPhase(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimStatusBuilder) AddAccessModes(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimStatusBuilder) AddCapacity(cb func(w *PersistentVolumeClaimStatus_CapacityEntryBuilder)) {
	x.buf.Reset()
	x.persistentVolumeClaimStatus_CapacityEntryBuilder.writer = &x.buf
	x.persistentVolumeClaimStatus_CapacityEntryBuilder.scratch = x.scratch
	cb(&x.persistentVolumeClaimStatus_CapacityEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimStatusBuilder) AddConditions(cb func(w *PersistentVolumeClaimConditionBuilder)) {
	x.buf.Reset()
	x.persistentVolumeClaimConditionBuilder.writer = &x.buf
	x.persistentVolumeClaimConditionBuilder.scratch = x.scratch
	cb(&x.persistentVolumeClaimConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type PersistentVolumeClaimStatus_CapacityEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPersistentVolumeClaimStatus_CapacityEntryBuilder(writer io.Writer) *PersistentVolumeClaimStatus_CapacityEntryBuilder {
	return &PersistentVolumeClaimStatus_CapacityEntryBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeClaimStatus_CapacityEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimStatus_CapacityEntryBuilder) SetValue(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type PersistentVolumeClaimSpecBuilder struct {
	writer                           io.Writer
	buf                              bytes.Buffer
	scratch                          []byte
	resourceRequirementsBuilder      ResourceRequirementsBuilder
	labelSelectorRequirementBuilder  LabelSelectorRequirementBuilder
	typedLocalObjectReferenceBuilder TypedLocalObjectReferenceBuilder
}

func NewPersistentVolumeClaimSpecBuilder(writer io.Writer) *PersistentVolumeClaimSpecBuilder {
	return &PersistentVolumeClaimSpecBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeClaimSpecBuilder) AddAccessModes(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimSpecBuilder) SetResources(cb func(w *ResourceRequirementsBuilder)) {
	x.buf.Reset()
	x.resourceRequirementsBuilder.writer = &x.buf
	x.resourceRequirementsBuilder.scratch = x.scratch
	cb(&x.resourceRequirementsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimSpecBuilder) SetVolumeName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimSpecBuilder) AddSelector(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *PersistentVolumeClaimSpecBuilder) SetStorageClassName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimSpecBuilder) SetVolumeMode(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimSpecBuilder) SetDataSource(cb func(w *TypedLocalObjectReferenceBuilder)) {
	x.buf.Reset()
	x.typedLocalObjectReferenceBuilder.writer = &x.buf
	x.typedLocalObjectReferenceBuilder.scratch = x.scratch
	cb(&x.typedLocalObjectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x3a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type TypedLocalObjectReferenceBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewTypedLocalObjectReferenceBuilder(writer io.Writer) *TypedLocalObjectReferenceBuilder {
	return &TypedLocalObjectReferenceBuilder{
		writer: writer,
	}
}
func (x *TypedLocalObjectReferenceBuilder) SetApiGroup(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *TypedLocalObjectReferenceBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *TypedLocalObjectReferenceBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PersistentVolumeClaimConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPersistentVolumeClaimConditionBuilder(writer io.Writer) *PersistentVolumeClaimConditionBuilder {
	return &PersistentVolumeClaimConditionBuilder{
		writer: writer,
	}
}
func (x *PersistentVolumeClaimConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimConditionBuilder) SetLastProbeTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x20)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PersistentVolumeClaimConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PolicyRuleBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPolicyRuleBuilder(writer io.Writer) *PolicyRuleBuilder {
	return &PolicyRuleBuilder{
		writer: writer,
	}
}
func (x *PolicyRuleBuilder) AddVerbs(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PolicyRuleBuilder) AddApiGroups(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PolicyRuleBuilder) AddResources(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PolicyRuleBuilder) AddResourceNames(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PolicyRuleBuilder) AddNonResourceURLs(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type SubjectBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewSubjectBuilder(writer io.Writer) *SubjectBuilder {
	return &SubjectBuilder{
		writer: writer,
	}
}
func (x *SubjectBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *SubjectBuilder) SetApiGroup(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *SubjectBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *SubjectBuilder) SetNamespace(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type RoleBuilder struct {
	writer            io.Writer
	buf               bytes.Buffer
	scratch           []byte
	metadataBuilder   MetadataBuilder
	policyRuleBuilder PolicyRuleBuilder
}

func NewRoleBuilder(writer io.Writer) *RoleBuilder {
	return &RoleBuilder{
		writer: writer,
	}
}
func (x *RoleBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBuilder) AddRules(cb func(w *PolicyRuleBuilder)) {
	x.buf.Reset()
	x.policyRuleBuilder.writer = &x.buf
	x.policyRuleBuilder.scratch = x.scratch
	cb(&x.policyRuleBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type RoleBindingBuilder struct {
	writer                           io.Writer
	buf                              bytes.Buffer
	scratch                          []byte
	metadataBuilder                  MetadataBuilder
	subjectBuilder                   SubjectBuilder
	typedLocalObjectReferenceBuilder TypedLocalObjectReferenceBuilder
}

func NewRoleBindingBuilder(writer io.Writer) *RoleBindingBuilder {
	return &RoleBindingBuilder{
		writer: writer,
	}
}
func (x *RoleBindingBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBindingBuilder) AddSubjects(cb func(w *SubjectBuilder)) {
	x.buf.Reset()
	x.subjectBuilder.writer = &x.buf
	x.subjectBuilder.scratch = x.scratch
	cb(&x.subjectBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBindingBuilder) SetRoleRef(cb func(w *TypedLocalObjectReferenceBuilder)) {
	x.buf.Reset()
	x.typedLocalObjectReferenceBuilder.writer = &x.buf
	x.typedLocalObjectReferenceBuilder.scratch = x.scratch
	cb(&x.typedLocalObjectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBindingBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *RoleBindingBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ClusterRoleBuilder struct {
	writer                          io.Writer
	buf                             bytes.Buffer
	scratch                         []byte
	metadataBuilder                 MetadataBuilder
	policyRuleBuilder               PolicyRuleBuilder
	labelSelectorRequirementBuilder LabelSelectorRequirementBuilder
	resourceMetricsBuilder          ResourceMetricsBuilder
}

func NewClusterRoleBuilder(writer io.Writer) *ClusterRoleBuilder {
	return &ClusterRoleBuilder{
		writer: writer,
	}
}
func (x *ClusterRoleBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBuilder) AddRules(cb func(w *PolicyRuleBuilder)) {
	x.buf.Reset()
	x.policyRuleBuilder.writer = &x.buf
	x.policyRuleBuilder.scratch = x.scratch
	cb(&x.policyRuleBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBuilder) AddAggregationRules(cb func(w *LabelSelectorRequirementBuilder)) {
	x.buf.Reset()
	x.labelSelectorRequirementBuilder.writer = &x.buf
	x.labelSelectorRequirementBuilder.scratch = x.scratch
	cb(&x.labelSelectorRequirementBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ClusterRoleBuilder) SetMetrics(cb func(w *ResourceMetricsBuilder)) {
	x.buf.Reset()
	x.resourceMetricsBuilder.writer = &x.buf
	x.resourceMetricsBuilder.scratch = x.scratch
	cb(&x.resourceMetricsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ClusterRoleBindingBuilder struct {
	writer                           io.Writer
	buf                              bytes.Buffer
	scratch                          []byte
	metadataBuilder                  MetadataBuilder
	subjectBuilder                   SubjectBuilder
	typedLocalObjectReferenceBuilder TypedLocalObjectReferenceBuilder
}

func NewClusterRoleBindingBuilder(writer io.Writer) *ClusterRoleBindingBuilder {
	return &ClusterRoleBindingBuilder{
		writer: writer,
	}
}
func (x *ClusterRoleBindingBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBindingBuilder) AddSubjects(cb func(w *SubjectBuilder)) {
	x.buf.Reset()
	x.subjectBuilder.writer = &x.buf
	x.subjectBuilder.scratch = x.scratch
	cb(&x.subjectBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBindingBuilder) SetRoleRef(cb func(w *TypedLocalObjectReferenceBuilder)) {
	x.buf.Reset()
	x.typedLocalObjectReferenceBuilder.writer = &x.buf
	x.typedLocalObjectReferenceBuilder.scratch = x.scratch
	cb(&x.typedLocalObjectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBindingBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ClusterRoleBindingBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ServiceAccountBuilder struct {
	writer                           io.Writer
	buf                              bytes.Buffer
	scratch                          []byte
	metadataBuilder                  MetadataBuilder
	objectReferenceBuilder           ObjectReferenceBuilder
	typedLocalObjectReferenceBuilder TypedLocalObjectReferenceBuilder
}

func NewServiceAccountBuilder(writer io.Writer) *ServiceAccountBuilder {
	return &ServiceAccountBuilder{
		writer: writer,
	}
}
func (x *ServiceAccountBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceAccountBuilder) AddSecrets(cb func(w *ObjectReferenceBuilder)) {
	x.buf.Reset()
	x.objectReferenceBuilder.writer = &x.buf
	x.objectReferenceBuilder.scratch = x.scratch
	cb(&x.objectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceAccountBuilder) AddImagePullSecrets(cb func(w *TypedLocalObjectReferenceBuilder)) {
	x.buf.Reset()
	x.typedLocalObjectReferenceBuilder.writer = &x.buf
	x.typedLocalObjectReferenceBuilder.scratch = x.scratch
	cb(&x.typedLocalObjectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceAccountBuilder) SetAutomountServiceAccountToken(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x20)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *ServiceAccountBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ServiceAccountBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type IngressServiceBackendBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewIngressServiceBackendBuilder(writer io.Writer) *IngressServiceBackendBuilder {
	return &IngressServiceBackendBuilder{
		writer: writer,
	}
}
func (x *IngressServiceBackendBuilder) SetServiceName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *IngressServiceBackendBuilder) SetPortName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *IngressServiceBackendBuilder) SetPortNumber(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type IngressBackendBuilder struct {
	writer                           io.Writer
	buf                              bytes.Buffer
	scratch                          []byte
	ingressServiceBackendBuilder     IngressServiceBackendBuilder
	typedLocalObjectReferenceBuilder TypedLocalObjectReferenceBuilder
}

func NewIngressBackendBuilder(writer io.Writer) *IngressBackendBuilder {
	return &IngressBackendBuilder{
		writer: writer,
	}
}
func (x *IngressBackendBuilder) SetService(cb func(w *IngressServiceBackendBuilder)) {
	x.buf.Reset()
	x.ingressServiceBackendBuilder.writer = &x.buf
	x.ingressServiceBackendBuilder.scratch = x.scratch
	cb(&x.ingressServiceBackendBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressBackendBuilder) SetResource(cb func(w *TypedLocalObjectReferenceBuilder)) {
	x.buf.Reset()
	x.typedLocalObjectReferenceBuilder.writer = &x.buf
	x.typedLocalObjectReferenceBuilder.scratch = x.scratch
	cb(&x.typedLocalObjectReferenceBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type IngressTLSBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewIngressTLSBuilder(writer io.Writer) *IngressTLSBuilder {
	return &IngressTLSBuilder{
		writer: writer,
	}
}
func (x *IngressTLSBuilder) AddHosts(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *IngressTLSBuilder) SetSecretName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type HTTPIngressPathBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	ingressBackendBuilder IngressBackendBuilder
}

func NewHTTPIngressPathBuilder(writer io.Writer) *HTTPIngressPathBuilder {
	return &HTTPIngressPathBuilder{
		writer: writer,
	}
}
func (x *HTTPIngressPathBuilder) SetPath(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HTTPIngressPathBuilder) SetPathType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HTTPIngressPathBuilder) SetBackend(cb func(w *IngressBackendBuilder)) {
	x.buf.Reset()
	x.ingressBackendBuilder.writer = &x.buf
	x.ingressBackendBuilder.scratch = x.scratch
	cb(&x.ingressBackendBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type IngressRuleBuilder struct {
	writer                 io.Writer
	buf                    bytes.Buffer
	scratch                []byte
	hTTPIngressPathBuilder HTTPIngressPathBuilder
}

func NewIngressRuleBuilder(writer io.Writer) *IngressRuleBuilder {
	return &IngressRuleBuilder{
		writer: writer,
	}
}
func (x *IngressRuleBuilder) SetHost(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *IngressRuleBuilder) AddHttpPaths(cb func(w *HTTPIngressPathBuilder)) {
	x.buf.Reset()
	x.hTTPIngressPathBuilder.writer = &x.buf
	x.hTTPIngressPathBuilder.scratch = x.scratch
	cb(&x.hTTPIngressPathBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type IngressSpecBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	ingressBackendBuilder IngressBackendBuilder
	ingressTLSBuilder     IngressTLSBuilder
	ingressRuleBuilder    IngressRuleBuilder
}

func NewIngressSpecBuilder(writer io.Writer) *IngressSpecBuilder {
	return &IngressSpecBuilder{
		writer: writer,
	}
}
func (x *IngressSpecBuilder) SetDefaultBackend(cb func(w *IngressBackendBuilder)) {
	x.buf.Reset()
	x.ingressBackendBuilder.writer = &x.buf
	x.ingressBackendBuilder.scratch = x.scratch
	cb(&x.ingressBackendBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressSpecBuilder) AddTls(cb func(w *IngressTLSBuilder)) {
	x.buf.Reset()
	x.ingressTLSBuilder.writer = &x.buf
	x.ingressTLSBuilder.scratch = x.scratch
	cb(&x.ingressTLSBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressSpecBuilder) AddRules(cb func(w *IngressRuleBuilder)) {
	x.buf.Reset()
	x.ingressRuleBuilder.writer = &x.buf
	x.ingressRuleBuilder.scratch = x.scratch
	cb(&x.ingressRuleBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressSpecBuilder) SetIngressClassName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type PortStatusBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewPortStatusBuilder(writer io.Writer) *PortStatusBuilder {
	return &PortStatusBuilder{
		writer: writer,
	}
}
func (x *PortStatusBuilder) SetPort(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *PortStatusBuilder) SetProtocol(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *PortStatusBuilder) SetError(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x1a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type LoadBalancerIngressBuilder struct {
	writer            io.Writer
	buf               bytes.Buffer
	scratch           []byte
	portStatusBuilder PortStatusBuilder
}

func NewLoadBalancerIngressBuilder(writer io.Writer) *LoadBalancerIngressBuilder {
	return &LoadBalancerIngressBuilder{
		writer: writer,
	}
}
func (x *LoadBalancerIngressBuilder) SetIp(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LoadBalancerIngressBuilder) SetHostname(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *LoadBalancerIngressBuilder) AddPorts(cb func(w *PortStatusBuilder)) {
	x.buf.Reset()
	x.portStatusBuilder.writer = &x.buf
	x.portStatusBuilder.scratch = x.scratch
	cb(&x.portStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type IngressStatusBuilder struct {
	writer                     io.Writer
	buf                        bytes.Buffer
	scratch                    []byte
	loadBalancerIngressBuilder LoadBalancerIngressBuilder
}

func NewIngressStatusBuilder(writer io.Writer) *IngressStatusBuilder {
	return &IngressStatusBuilder{
		writer: writer,
	}
}
func (x *IngressStatusBuilder) AddIngress(cb func(w *LoadBalancerIngressBuilder)) {
	x.buf.Reset()
	x.loadBalancerIngressBuilder.writer = &x.buf
	x.loadBalancerIngressBuilder.scratch = x.scratch
	cb(&x.loadBalancerIngressBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type IngressBuilder struct {
	writer               io.Writer
	buf                  bytes.Buffer
	scratch              []byte
	metadataBuilder      MetadataBuilder
	ingressSpecBuilder   IngressSpecBuilder
	ingressStatusBuilder IngressStatusBuilder
}

func NewIngressBuilder(writer io.Writer) *IngressBuilder {
	return &IngressBuilder{
		writer: writer,
	}
}
func (x *IngressBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressBuilder) SetSpec(cb func(w *IngressSpecBuilder)) {
	x.buf.Reset()
	x.ingressSpecBuilder.writer = &x.buf
	x.ingressSpecBuilder.scratch = x.scratch
	cb(&x.ingressSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressBuilder) SetStatus(cb func(w *IngressStatusBuilder)) {
	x.buf.Reset()
	x.ingressStatusBuilder.writer = &x.buf
	x.ingressStatusBuilder.scratch = x.scratch
	cb(&x.ingressStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *IngressBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type KafkaRequestHeaderBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewKafkaRequestHeaderBuilder(writer io.Writer) *KafkaRequestHeaderBuilder {
	return &KafkaRequestHeaderBuilder{
		writer: writer,
	}
}
func (x *KafkaRequestHeaderBuilder) SetRequest_type(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *KafkaRequestHeaderBuilder) SetRequest_version(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type KafkaAggregationBuilder struct {
	writer                    io.Writer
	buf                       bytes.Buffer
	scratch                   []byte
	kafkaRequestHeaderBuilder KafkaRequestHeaderBuilder
}

func NewKafkaAggregationBuilder(writer io.Writer) *KafkaAggregationBuilder {
	return &KafkaAggregationBuilder{
		writer: writer,
	}
}
func (x *KafkaAggregationBuilder) SetHeader(cb func(w *KafkaRequestHeaderBuilder)) {
	x.buf.Reset()
	x.kafkaRequestHeaderBuilder.writer = &x.buf
	x.kafkaRequestHeaderBuilder.scratch = x.scratch
	cb(&x.kafkaRequestHeaderBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *KafkaAggregationBuilder) SetTopic(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *KafkaAggregationBuilder) SetCount(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type DataStreamsAggregationsBuilder struct {
	writer                                                  io.Writer
	buf                                                     bytes.Buffer
	scratch                                                 []byte
	dataStreamsAggregations_KafkaProduceAggregationsBuilder DataStreamsAggregations_KafkaProduceAggregationsBuilder
	dataStreamsAggregations_KafkaFetchAggregationsBuilder   DataStreamsAggregations_KafkaFetchAggregationsBuilder
	kafkaAggregationBuilder                                 KafkaAggregationBuilder
}

func NewDataStreamsAggregationsBuilder(writer io.Writer) *DataStreamsAggregationsBuilder {
	return &DataStreamsAggregationsBuilder{
		writer: writer,
	}
}
func (x *DataStreamsAggregationsBuilder) SetKafkaProduceAggregations(cb func(w *DataStreamsAggregations_KafkaProduceAggregationsBuilder)) {
	x.buf.Reset()
	x.dataStreamsAggregations_KafkaProduceAggregationsBuilder.writer = &x.buf
	x.dataStreamsAggregations_KafkaProduceAggregationsBuilder.scratch = x.scratch
	cb(&x.dataStreamsAggregations_KafkaProduceAggregationsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DataStreamsAggregationsBuilder) SetKafkaFetchAggregations(cb func(w *DataStreamsAggregations_KafkaFetchAggregationsBuilder)) {
	x.buf.Reset()
	x.dataStreamsAggregations_KafkaFetchAggregationsBuilder.writer = &x.buf
	x.dataStreamsAggregations_KafkaFetchAggregationsBuilder.scratch = x.scratch
	cb(&x.dataStreamsAggregations_KafkaFetchAggregationsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *DataStreamsAggregationsBuilder) AddKafkaAggregations(cb func(w *KafkaAggregationBuilder)) {
	x.buf.Reset()
	x.kafkaAggregationBuilder.writer = &x.buf
	x.kafkaAggregationBuilder.scratch = x.scratch
	cb(&x.kafkaAggregationBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type DataStreamsAggregations_TopicStatsBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewDataStreamsAggregations_TopicStatsBuilder(writer io.Writer) *DataStreamsAggregations_TopicStatsBuilder {
	return &DataStreamsAggregations_TopicStatsBuilder{
		writer: writer,
	}
}
func (x *DataStreamsAggregations_TopicStatsBuilder) SetTopic(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *DataStreamsAggregations_TopicStatsBuilder) SetCount(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x10)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type DataStreamsAggregations_KafkaProduceAggregationsBuilder struct {
	writer                                    io.Writer
	buf                                       bytes.Buffer
	scratch                                   []byte
	dataStreamsAggregations_TopicStatsBuilder DataStreamsAggregations_TopicStatsBuilder
}

func NewDataStreamsAggregations_KafkaProduceAggregationsBuilder(writer io.Writer) *DataStreamsAggregations_KafkaProduceAggregationsBuilder {
	return &DataStreamsAggregations_KafkaProduceAggregationsBuilder{
		writer: writer,
	}
}
func (x *DataStreamsAggregations_KafkaProduceAggregationsBuilder) AddStats(cb func(w *DataStreamsAggregations_TopicStatsBuilder)) {
	x.buf.Reset()
	x.dataStreamsAggregations_TopicStatsBuilder.writer = &x.buf
	x.dataStreamsAggregations_TopicStatsBuilder.scratch = x.scratch
	cb(&x.dataStreamsAggregations_TopicStatsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type DataStreamsAggregations_KafkaFetchAggregationsBuilder struct {
	writer                                    io.Writer
	buf                                       bytes.Buffer
	scratch                                   []byte
	dataStreamsAggregations_TopicStatsBuilder DataStreamsAggregations_TopicStatsBuilder
}

func NewDataStreamsAggregations_KafkaFetchAggregationsBuilder(writer io.Writer) *DataStreamsAggregations_KafkaFetchAggregationsBuilder {
	return &DataStreamsAggregations_KafkaFetchAggregationsBuilder{
		writer: writer,
	}
}
func (x *DataStreamsAggregations_KafkaFetchAggregationsBuilder) AddStats(cb func(w *DataStreamsAggregations_TopicStatsBuilder)) {
	x.buf.Reset()
	x.dataStreamsAggregations_TopicStatsBuilder.writer = &x.buf
	x.dataStreamsAggregations_TopicStatsBuilder.scratch = x.scratch
	cb(&x.dataStreamsAggregations_TopicStatsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HTTPAggregationsBuilder struct {
	writer           io.Writer
	buf              bytes.Buffer
	scratch          []byte
	hTTPStatsBuilder HTTPStatsBuilder
}

func NewHTTPAggregationsBuilder(writer io.Writer) *HTTPAggregationsBuilder {
	return &HTTPAggregationsBuilder{
		writer: writer,
	}
}
func (x *HTTPAggregationsBuilder) AddEndpointAggregations(cb func(w *HTTPStatsBuilder)) {
	x.buf.Reset()
	x.hTTPStatsBuilder.writer = &x.buf
	x.hTTPStatsBuilder.scratch = x.scratch
	cb(&x.hTTPStatsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HTTP2AggregationsBuilder struct {
	writer           io.Writer
	buf              bytes.Buffer
	scratch          []byte
	hTTPStatsBuilder HTTPStatsBuilder
}

func NewHTTP2AggregationsBuilder(writer io.Writer) *HTTP2AggregationsBuilder {
	return &HTTP2AggregationsBuilder{
		writer: writer,
	}
}
func (x *HTTP2AggregationsBuilder) AddEndpointAggregations(cb func(w *HTTPStatsBuilder)) {
	x.buf.Reset()
	x.hTTPStatsBuilder.writer = &x.buf
	x.hTTPStatsBuilder.scratch = x.scratch
	cb(&x.hTTPStatsBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HTTPStatsBuilder struct {
	writer                                  io.Writer
	buf                                     bytes.Buffer
	scratch                                 []byte
	hTTPStats_DataBuilder                   HTTPStats_DataBuilder
	hTTPStats_StatsByStatusCodeEntryBuilder HTTPStats_StatsByStatusCodeEntryBuilder
}

func NewHTTPStatsBuilder(writer io.Writer) *HTTPStatsBuilder {
	return &HTTPStatsBuilder{
		writer: writer,
	}
}
func (x *HTTPStatsBuilder) SetPath(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *HTTPStatsBuilder) SetMethod(v uint64) {
	if v != 0 {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x28)
		x.scratch = protowire.AppendVarint(x.scratch, v)
		x.writer.Write(x.scratch)
	}
}
func (x *HTTPStatsBuilder) SetFullPath(v bool) {
	if v {
		x.scratch = protowire.AppendVarint(x.scratch[:0], 0x30)
		x.scratch = protowire.AppendVarint(x.scratch, 1)
		x.writer.Write(x.scratch)
	}
}
func (x *HTTPStatsBuilder) AddStatsByResponseStatus(cb func(w *HTTPStats_DataBuilder)) {
	x.buf.Reset()
	x.hTTPStats_DataBuilder.writer = &x.buf
	x.hTTPStats_DataBuilder.scratch = x.scratch
	cb(&x.hTTPStats_DataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HTTPStatsBuilder) AddStatsByStatusCode(cb func(w *HTTPStats_StatsByStatusCodeEntryBuilder)) {
	x.buf.Reset()
	x.hTTPStats_StatsByStatusCodeEntryBuilder.writer = &x.buf
	x.hTTPStats_StatsByStatusCodeEntryBuilder.scratch = x.scratch
	cb(&x.hTTPStats_StatsByStatusCodeEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HTTPStats_StatsByStatusCodeEntryBuilder struct {
	writer                io.Writer
	buf                   bytes.Buffer
	scratch               []byte
	hTTPStats_DataBuilder HTTPStats_DataBuilder
}

func NewHTTPStats_StatsByStatusCodeEntryBuilder(writer io.Writer) *HTTPStats_StatsByStatusCodeEntryBuilder {
	return &HTTPStats_StatsByStatusCodeEntryBuilder{
		writer: writer,
	}
}
func (x *HTTPStats_StatsByStatusCodeEntryBuilder) SetKey(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HTTPStats_StatsByStatusCodeEntryBuilder) SetValue(cb func(w *HTTPStats_DataBuilder)) {
	x.buf.Reset()
	x.hTTPStats_DataBuilder.writer = &x.buf
	x.hTTPStats_DataBuilder.scratch = x.scratch
	cb(&x.hTTPStats_DataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type HTTPStats_DataBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewHTTPStats_DataBuilder(writer io.Writer) *HTTPStats_DataBuilder {
	return &HTTPStats_DataBuilder{
		writer: writer,
	}
}
func (x *HTTPStats_DataBuilder) SetCount(v uint32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *HTTPStats_DataBuilder) SetLatencies(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *HTTPStats_DataBuilder) SetFirstLatencySample(v float64) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x21)
	x.scratch = protowire.AppendFixed64(x.scratch, math.Float64bits(v))
	x.writer.Write(x.scratch)
}

type DNSDatabaseEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewDNSDatabaseEntryBuilder(writer io.Writer) *DNSDatabaseEntryBuilder {
	return &DNSDatabaseEntryBuilder{
		writer: writer,
	}
}
func (x *DNSDatabaseEntryBuilder) AddNameOffsets(v int32) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}

type ResourceListBuilder struct {
	writer                                io.Writer
	buf                                   bytes.Buffer
	scratch                               []byte
	resourceList_MetricValuesEntryBuilder ResourceList_MetricValuesEntryBuilder
}

func NewResourceListBuilder(writer io.Writer) *ResourceListBuilder {
	return &ResourceListBuilder{
		writer: writer,
	}
}
func (x *ResourceListBuilder) AddMetricValues(cb func(w *ResourceList_MetricValuesEntryBuilder)) {
	x.buf.Reset()
	x.resourceList_MetricValuesEntryBuilder.writer = &x.buf
	x.resourceList_MetricValuesEntryBuilder.scratch = x.scratch
	cb(&x.resourceList_MetricValuesEntryBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ResourceList_MetricValuesEntryBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewResourceList_MetricValuesEntryBuilder(writer io.Writer) *ResourceList_MetricValuesEntryBuilder {
	return &ResourceList_MetricValuesEntryBuilder{
		writer: writer,
	}
}
func (x *ResourceList_MetricValuesEntryBuilder) SetKey(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ResourceList_MetricValuesEntryBuilder) SetValue(v float64) {
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x11)
	x.scratch = protowire.AppendFixed64(x.scratch, math.Float64bits(v))
	x.writer.Write(x.scratch)
}

type VerticalPodAutoscalerBuilder struct {
	writer                                io.Writer
	buf                                   bytes.Buffer
	scratch                               []byte
	metadataBuilder                       MetadataBuilder
	verticalPodAutoscalerSpecBuilder      VerticalPodAutoscalerSpecBuilder
	verticalPodAutoscalerStatusBuilder    VerticalPodAutoscalerStatusBuilder
	verticalPodAutoscalerConditionBuilder VerticalPodAutoscalerConditionBuilder
}

func NewVerticalPodAutoscalerBuilder(writer io.Writer) *VerticalPodAutoscalerBuilder {
	return &VerticalPodAutoscalerBuilder{
		writer: writer,
	}
}
func (x *VerticalPodAutoscalerBuilder) SetMetadata(cb func(w *MetadataBuilder)) {
	x.buf.Reset()
	x.metadataBuilder.writer = &x.buf
	x.metadataBuilder.scratch = x.scratch
	cb(&x.metadataBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *VerticalPodAutoscalerBuilder) SetSpec(cb func(w *VerticalPodAutoscalerSpecBuilder)) {
	x.buf.Reset()
	x.verticalPodAutoscalerSpecBuilder.writer = &x.buf
	x.verticalPodAutoscalerSpecBuilder.scratch = x.scratch
	cb(&x.verticalPodAutoscalerSpecBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *VerticalPodAutoscalerBuilder) SetStatus(cb func(w *VerticalPodAutoscalerStatusBuilder)) {
	x.buf.Reset()
	x.verticalPodAutoscalerStatusBuilder.writer = &x.buf
	x.verticalPodAutoscalerStatusBuilder.scratch = x.scratch
	cb(&x.verticalPodAutoscalerStatusBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *VerticalPodAutoscalerBuilder) SetYaml(cb func(b *bytes.Buffer)) {
	x.buf.Reset()
	cb(&x.buf)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *VerticalPodAutoscalerBuilder) AddTags(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerBuilder) AddConditions(cb func(w *VerticalPodAutoscalerConditionBuilder)) {
	x.buf.Reset()
	x.verticalPodAutoscalerConditionBuilder.writer = &x.buf
	x.verticalPodAutoscalerConditionBuilder.scratch = x.scratch
	cb(&x.verticalPodAutoscalerConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x32)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type VerticalPodAutoscalerConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewVerticalPodAutoscalerConditionBuilder(writer io.Writer) *VerticalPodAutoscalerConditionBuilder {
	return &VerticalPodAutoscalerConditionBuilder{
		writer: writer,
	}
}
func (x *VerticalPodAutoscalerConditionBuilder) SetType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerConditionBuilder) SetStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type VerticalPodAutoscalerSpecBuilder struct {
	writer                             io.Writer
	buf                                bytes.Buffer
	scratch                            []byte
	verticalPodAutoscalerTargetBuilder VerticalPodAutoscalerTargetBuilder
	containerResourcePolicyBuilder     ContainerResourcePolicyBuilder
}

func NewVerticalPodAutoscalerSpecBuilder(writer io.Writer) *VerticalPodAutoscalerSpecBuilder {
	return &VerticalPodAutoscalerSpecBuilder{
		writer: writer,
	}
}
func (x *VerticalPodAutoscalerSpecBuilder) SetTarget(cb func(w *VerticalPodAutoscalerTargetBuilder)) {
	x.buf.Reset()
	x.verticalPodAutoscalerTargetBuilder.writer = &x.buf
	x.verticalPodAutoscalerTargetBuilder.scratch = x.scratch
	cb(&x.verticalPodAutoscalerTargetBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0xa)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *VerticalPodAutoscalerSpecBuilder) SetUpdateMode(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerSpecBuilder) AddResourcePolicies(cb func(w *ContainerResourcePolicyBuilder)) {
	x.buf.Reset()
	x.containerResourcePolicyBuilder.writer = &x.buf
	x.containerResourcePolicyBuilder.scratch = x.scratch
	cb(&x.containerResourcePolicyBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type VerticalPodAutoscalerTargetBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewVerticalPodAutoscalerTargetBuilder(writer io.Writer) *VerticalPodAutoscalerTargetBuilder {
	return &VerticalPodAutoscalerTargetBuilder{
		writer: writer,
	}
}
func (x *VerticalPodAutoscalerTargetBuilder) SetKind(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerTargetBuilder) SetName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type ContainerResourcePolicyBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	resourceListBuilder ResourceListBuilder
}

func NewContainerResourcePolicyBuilder(writer io.Writer) *ContainerResourcePolicyBuilder {
	return &ContainerResourcePolicyBuilder{
		writer: writer,
	}
}
func (x *ContainerResourcePolicyBuilder) SetContainerName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerResourcePolicyBuilder) SetMode(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerResourcePolicyBuilder) SetMinAllowed(cb func(w *ResourceListBuilder)) {
	x.buf.Reset()
	x.resourceListBuilder.writer = &x.buf
	x.resourceListBuilder.scratch = x.scratch
	cb(&x.resourceListBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerResourcePolicyBuilder) SetMaxAllowed(cb func(w *ResourceListBuilder)) {
	x.buf.Reset()
	x.resourceListBuilder.writer = &x.buf
	x.resourceListBuilder.scratch = x.scratch
	cb(&x.resourceListBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerResourcePolicyBuilder) AddControlledResource(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerResourcePolicyBuilder) SetControlledValues(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x32)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}

type VerticalPodAutoscalerStatusBuilder struct {
	writer                         io.Writer
	buf                            bytes.Buffer
	scratch                        []byte
	containerRecommendationBuilder ContainerRecommendationBuilder
	vPAConditionBuilder            VPAConditionBuilder
}

func NewVerticalPodAutoscalerStatusBuilder(writer io.Writer) *VerticalPodAutoscalerStatusBuilder {
	return &VerticalPodAutoscalerStatusBuilder{
		writer: writer,
	}
}
func (x *VerticalPodAutoscalerStatusBuilder) SetLastRecommendedDate(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x8)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *VerticalPodAutoscalerStatusBuilder) AddRecommendations(cb func(w *ContainerRecommendationBuilder)) {
	x.buf.Reset()
	x.containerRecommendationBuilder.writer = &x.buf
	x.containerRecommendationBuilder.scratch = x.scratch
	cb(&x.containerRecommendationBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *VerticalPodAutoscalerStatusBuilder) AddConditions(cb func(w *VPAConditionBuilder)) {
	x.buf.Reset()
	x.vPAConditionBuilder.writer = &x.buf
	x.vPAConditionBuilder.scratch = x.scratch
	cb(&x.vPAConditionBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type ContainerRecommendationBuilder struct {
	writer              io.Writer
	buf                 bytes.Buffer
	scratch             []byte
	resourceListBuilder ResourceListBuilder
}

func NewContainerRecommendationBuilder(writer io.Writer) *ContainerRecommendationBuilder {
	return &ContainerRecommendationBuilder{
		writer: writer,
	}
}
func (x *ContainerRecommendationBuilder) SetContainerName(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *ContainerRecommendationBuilder) SetTarget(cb func(w *ResourceListBuilder)) {
	x.buf.Reset()
	x.resourceListBuilder.writer = &x.buf
	x.resourceListBuilder.scratch = x.scratch
	cb(&x.resourceListBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x12)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerRecommendationBuilder) SetLowerBound(cb func(w *ResourceListBuilder)) {
	x.buf.Reset()
	x.resourceListBuilder.writer = &x.buf
	x.resourceListBuilder.scratch = x.scratch
	cb(&x.resourceListBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x1a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerRecommendationBuilder) SetUpperBound(cb func(w *ResourceListBuilder)) {
	x.buf.Reset()
	x.resourceListBuilder.writer = &x.buf
	x.resourceListBuilder.scratch = x.scratch
	cb(&x.resourceListBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x22)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}
func (x *ContainerRecommendationBuilder) SetUncappedTarget(cb func(w *ResourceListBuilder)) {
	x.buf.Reset()
	x.resourceListBuilder.writer = &x.buf
	x.resourceListBuilder.scratch = x.scratch
	cb(&x.resourceListBuilder)
	x.scratch = protowire.AppendVarint(x.scratch[:0], 0x2a)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(x.buf.Len()))
	x.writer.Write(x.scratch)
	x.writer.Write(x.buf.Bytes())
}

type VPAConditionBuilder struct {
	writer  io.Writer
	buf     bytes.Buffer
	scratch []byte
}

func NewVPAConditionBuilder(writer io.Writer) *VPAConditionBuilder {
	return &VPAConditionBuilder{
		writer: writer,
	}
}
func (x *VPAConditionBuilder) SetConditionType(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0xa)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VPAConditionBuilder) SetConditionStatus(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x12)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VPAConditionBuilder) SetLastTransitionTime(v int64) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x18)
	x.scratch = protowire.AppendVarint(x.scratch, uint64(v))
	x.writer.Write(x.scratch)
}
func (x *VPAConditionBuilder) SetReason(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x22)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
func (x *VPAConditionBuilder) SetMessage(v string) {
	x.scratch = x.scratch[:0]
	x.scratch = protowire.AppendVarint(x.scratch, 0x2a)
	x.scratch = protowire.AppendString(x.scratch, v)
	x.writer.Write(x.scratch)
}
