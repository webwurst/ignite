

# v1alpha1
`import "/go/src/github.com/weaveworks/ignite/pkg/apis/ignite/v1alpha1"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
+k8s:deepcopy-gen=package
+k8s:defaulter-gen=TypeMeta




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func SetDefaults_ImageSource(obj *ImageSource)](#SetDefaults_ImageSource)
* [func SetDefaults_PoolSpec(obj *PoolSpec)](#SetDefaults_PoolSpec)
* [func SetDefaults_VMSpec(obj *VMSpec)](#SetDefaults_VMSpec)
* [func SetDefaults_VMStatus(obj *VMStatus)](#SetDefaults_VMStatus)
* [type FileMapping](#FileMapping)
* [type Image](#Image)
* [type ImageClaim](#ImageClaim)
* [type ImageSource](#ImageSource)
* [type ImageSourceType](#ImageSourceType)
* [type ImageSpec](#ImageSpec)
* [type Kernel](#Kernel)
* [type KernelClaim](#KernelClaim)
* [type KernelSpec](#KernelSpec)
* [type Pool](#Pool)
* [type PoolDevice](#PoolDevice)
* [type PoolDeviceType](#PoolDeviceType)
* [type PoolSpec](#PoolSpec)
* [type PoolStatus](#PoolStatus)
* [type PortMapping](#PortMapping)
* [type SSH](#SSH)
* [type VM](#VM)
* [type VMSpec](#VMSpec)
* [type VMState](#VMState)
* [type VMStatus](#VMStatus)


#### <a name="pkg-files">Package files</a>
[defaults.go](/src/target/defaults.go) [doc.go](/src/target/doc.go) [register.go](/src/target/register.go) [types.go](/src/target/types.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    // GroupName is the group name use in this package
    GroupName = "ignite.weave.works"

    // VMKind returns the kind for the VM API type
    VMKind = "VM"
    // KernelKind returns the kind for the Kernel API type
    KernelKind = "Kernel"
    // PoolKind returns the kind for the Pool API type
    PoolKind = "Pool"
    // ImageKind returns the kind for the Image API type
    ImageKind = "Image"
)
```

## <a name="pkg-variables">Variables</a>
``` go
var (
    // SchemeBuilder the schema builder
    SchemeBuilder = runtime.NewSchemeBuilder(
        addKnownTypes,
        addDefaultingFuncs,
    )

    AddToScheme = localSchemeBuilder.AddToScheme
)
```
``` go
var SchemeGroupVersion = schema.GroupVersion{
    Group:   GroupName,
    Version: "v1alpha1",
}
```
SchemeGroupVersion is group version used to register these objects



## <a name="SetDefaults_ImageSource">func</a> [SetDefaults_ImageSource](/src/target/defaults.go?s=263:309#L13)
``` go
func SetDefaults_ImageSource(obj *ImageSource)
```


## <a name="SetDefaults_PoolSpec">func</a> [SetDefaults_PoolSpec](/src/target/defaults.go?s=349:389#L17)
``` go
func SetDefaults_PoolSpec(obj *PoolSpec)
```


## <a name="SetDefaults_VMSpec">func</a> [SetDefaults_VMSpec](/src/target/defaults.go?s=915:951#L39)
``` go
func SetDefaults_VMSpec(obj *VMSpec)
```


## <a name="SetDefaults_VMStatus">func</a> [SetDefaults_VMStatus](/src/target/defaults.go?s=1338:1378#L59)
``` go
func SetDefaults_VMStatus(obj *VMStatus)
```



## <a name="FileMapping">type</a> [FileMapping](/src/target/types.go?s=6878:6973#L188)
``` go
type FileMapping struct {
    HostPath string `json:"hostPath"`
    VMPath   string `json:"vmPath"`
}

```
FileMapping defines mappings between files on the host and VM










## <a name="Image">type</a> [Image](/src/target/types.go?s=285:749#L13)
``` go
type Image struct {
    metav1.TypeMeta `json:",inline"`
    // meta.ObjectMeta is also embedded into the struct, and defines the human-readable name, and the machine-readable ID
    // Name is available at the .metadata.name JSON path
    // ID is available at the .metadata.uid JSON path (the Go type is k8s.io/apimachinery/pkg/types.UID, which is only a typed string)
    meta.ObjectMeta `json:"metadata"`

    Spec ImageSpec `json:"spec"`
}

```
Image represents a cached OCI image ready to be used with Ignite
+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object










## <a name="ImageClaim">type</a> [ImageClaim](/src/target/types.go?s=6329:6502#L168)
``` go
type ImageClaim struct {
    Type ImageSourceType `json:"type"`
    Ref  string          `json:"ref"`
    // TODO: Temporary ID for the old metadata handling
    ID string `json:"ID"`
}

```
ImageClaim specifies a claim to import an image










## <a name="ImageSource">type</a> [ImageSource](/src/target/types.go?s=1152:1509#L38)
``` go
type ImageSource struct {
    // Type defines how the image was imported
    Type ImageSourceType `json:"type"`
    // ID defines the source's ID (e.g. the Docker image ID)
    ID string `json:"id"`
    // Name defines the user-friendly name of the imported source
    Name string `json:"name"`
    // Size defines the size of the source in bytes
    Size meta.Size `json:"size"`
}

```
ImageSource defines where the image was imported from










## <a name="ImageSourceType">type</a> [ImageSourceType](/src/target/types.go?s=932:959#L30)
``` go
type ImageSourceType string
```
ImageSourceType is an enum of different supported Image Source Types


``` go
const (
    // ImageSourceTypeDocker defines that the image is imported from Docker
    ImageSourceTypeDocker ImageSourceType = "Docker"
)
```









## <a name="ImageSpec">type</a> [ImageSpec](/src/target/types.go?s=797:858#L25)
``` go
type ImageSpec struct {
    Source ImageSource `json:"source"`
}

```
ImageSpec declares what the image contains










## <a name="Kernel">type</a> [Kernel](/src/target/types.go?s=4071:4538#L114)
``` go
type Kernel struct {
    metav1.TypeMeta `json:",inline"`
    // meta.ObjectMeta is also embedded into the struct, and defines the human-readable name, and the machine-readable ID
    // Name is available at the .metadata.name JSON path
    // ID is available at the .metadata.uid JSON path (the Go type is k8s.io/apimachinery/pkg/types.UID, which is only a typed string)
    meta.ObjectMeta `json:"metadata"`

    Spec KernelSpec `json:"spec"`
}

```
Kernel is a serializable object that caches information about imported kernels
This file is stored in /var/lib/firecracker/kernels/{oci-image-digest}/metadata.json
+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object










## <a name="KernelClaim">type</a> [KernelClaim](/src/target/types.go?s=6560:6648#L176)
``` go
type KernelClaim struct {
    ID      string `json:"ID"`
    CmdLine string `json:"cmdline"`
}

```
TODO: Temporary helper for the old metadata handling










## <a name="KernelSpec">type</a> [KernelSpec](/src/target/types.go?s=4591:4797#L126)
``` go
type KernelSpec struct {
    Version string      `json:"version"`
    Source  ImageSource `json:"source"`
}

```
KernelSpec describes the properties of a kernel










## <a name="Pool">type</a> [Pool](/src/target/types.go?s=1965:2147#L59)
``` go
type Pool struct {
    metav1.TypeMeta `json:",inline"`

    Spec   PoolSpec   `json:"spec"`
    Status PoolStatus `json:"status"`
}

```
Pool defines device mapper pool database
This file is managed by the snapshotter part of Ignite, and the file (existing as a singleton)
is present at /var/lib/firecracker/snapshotter/pool.json
+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object










## <a name="PoolDevice">type</a> [PoolDevice](/src/target/types.go?s=3438:3828#L101)
``` go
type PoolDevice struct {
    Size   meta.Size `json:"size"`
    Parent meta.DMID `json:"parent"`
    // Type specifies the type of the contents of the device
    Type PoolDeviceType `json:"type"`
    // MetadataPath points to the JSON/YAML file with metadata about this device
    // This is most often of the format /var/lib/firecracker/{type}/{id}/metadata.json
    MetadataPath string `json:"metadataPath"`
}

```
PoolDevice defines one device in the pool










## <a name="PoolDeviceType">type</a> [PoolDeviceType](/src/target/types.go?s=3167:3193#L91)
``` go
type PoolDeviceType string
```

``` go
const (
    PoolDeviceTypeImage  PoolDeviceType = "Image"
    PoolDeviceTypeResize PoolDeviceType = "Resize"
    PoolDeviceTypeKernel PoolDeviceType = "Kernel"
    PoolDeviceTypeVM     PoolDeviceType = "VM"
)
```









## <a name="PoolSpec">type</a> [PoolSpec](/src/target/types.go?s=2194:2907#L69)
``` go
type PoolSpec struct {
    // MetadataSize specifies the size of the pool's metadata
    MetadataSize meta.Size `json:"metadataSize"`
    // DataSize specifies the size of the pool's data
    DataSize meta.Size `json:"dataSize"`
    // AllocationSize specifies the smallest size that can be allocated at a time
    AllocationSize meta.Size `json:"allocationSize"`
    // MetadataPath points to the file where device mapper stores all metadata information
    // Defaults to constants.SNAPSHOTTER_METADATA_PATH
    MetadataPath string `json:"metadataPath"`
    // DataPath points to the backing physical device or sparse file (to be loop mounted) for the pool
    // Defaults to constants.SNAPSHOTTER_DATA_PATH
    DataPath string `json:"dataPath"`
}

```
PoolSpec defines the Pool's specification










## <a name="PoolStatus">type</a> [PoolStatus](/src/target/types.go?s=2957:3165#L85)
``` go
type PoolStatus struct {
    // The Devices array needs to contain pointers to accommodate "holes" in the mapping
    // Where devices have been deleted, the pointer is nil
    Devices []*PoolDevice `json:"devices"`
}

```
PoolStatus defines the Pool's current status










## <a name="PortMapping">type</a> [PortMapping](/src/target/types.go?s=6716:6811#L182)
``` go
type PortMapping struct {
    HostPort uint64 `json:"hostPort"`
    VMPort   uint64 `json:"vmPort"`
}

```
PortMapping defines a port mapping between the VM and the host










## <a name="SSH">type</a> [SSH](/src/target/types.go?s=7036:7092#L194)
``` go
type SSH struct {
    PublicKey string `json:"publicKey"`
}

```
SSH specifies different ways to connect via SSH to the VM










## <a name="VM">type</a> [VM](/src/target/types.go?s=4999:5456#L136)
``` go
type VM struct {
    metav1.TypeMeta `json:",inline"`
    // meta.ObjectMeta is also embedded into the struct, and defines the human-readable name, and the machine-readable ID
    // Name is available at the .metadata.name JSON path
    // ID is available at the .metadata.uid JSON path (the Go type is k8s.io/apimachinery/pkg/types.UID, which is only a typed string)
    meta.ObjectMeta `json:"metadata"`

    Spec   VMSpec   `json:"spec"`
    Status VMStatus `json:"status"`
}

```
VM represents a virtual machine run by Firecracker
These files are stored in /var/lib/firecracker/vm/{vm-id}/metadata.json
+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object










## <a name="VMSpec">type</a> [VMSpec](/src/target/types.go?s=5504:6276#L148)
``` go
type VMSpec struct {
    Image *ImageClaim `json:"image"`
    // TODO: Temporary ID for the old metadata handling
    Kernel   *KernelClaim  `json:"kernel"`
    CPUs     uint64        `json:"cpus"`
    Memory   meta.Size     `json:"memory"`
    DiskSize meta.Size     `json:"diskSize"`
    Ports    []PortMapping `json:"ports,omitempty"`
    // This will be done at either "ignite start" or "ignite create" time
    // TODO: We might to revisit this later
    CopyFiles []FileMapping `json:"copyFiles,omitempty"`
    // SSH specifies how the SSH setup should be done
    // SSH appends to CopyFiles when active
    // nil here means "don't do anything special"
    // An empty struct means "generate a new SSH key and copy it in"
    // Specifying a path means "use this public key"
    SSH *SSH `json:"ssh,omitempty"`
}

```
VMSpec describes the configuration of a VM










## <a name="VMState">type</a> [VMState](/src/target/types.go?s=7145:7164#L199)
``` go
type VMState string
```
VMState defines different states a VM can be in


``` go
const (
    VMStateCreated VMState = "Created"
    VMStateRunning VMState = "Running"
    VMStateStopped VMState = "Stopped"
)
```









## <a name="VMStatus">type</a> [VMStatus](/src/target/types.go?s=7324:7428#L208)
``` go
type VMStatus struct {
    State       VMState  `json:"state"`
    IPAddresses []net.IP `json:"ipAddresses"`
}

```
VMStatus defines the status of a VM














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)