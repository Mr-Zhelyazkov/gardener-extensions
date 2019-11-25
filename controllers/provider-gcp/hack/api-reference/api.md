<p>Packages:</p>
<ul>
<li>
<a href="#gcp.provider.extensions.gardener.cloud%2fv1alpha1">gcp.provider.extensions.gardener.cloud/v1alpha1</a>
</li>
</ul>
<h2 id="gcp.provider.extensions.gardener.cloud/v1alpha1">gcp.provider.extensions.gardener.cloud/v1alpha1</h2>
<p>
<p>Package v1alpha1 contains the GCP provider API resources.</p>
</p>
Resource Types:
<ul><li>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.CloudProfileConfig">CloudProfileConfig</a>
</li><li>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.ControlPlaneConfig">ControlPlaneConfig</a>
</li><li>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.InfrastructureConfig">InfrastructureConfig</a>
</li></ul>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.CloudProfileConfig">CloudProfileConfig
</h3>
<p>
<p>CloudProfileConfig contains provider-specific configuration that is embedded into Gardener&rsquo;s <code>CloudProfile</code>
resource.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
gcp.provider.extensions.gardener.cloud/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>CloudProfileConfig</code></td>
</tr>
<tr>
<td>
<code>machineImages</code></br>
<em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.MachineImages">
[]MachineImages
</a>
</em>
</td>
<td>
<p>MachineImages is the list of machine images that are understood by the controller. It maps
logical names and versions to provider-specific identifiers.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.ControlPlaneConfig">ControlPlaneConfig
</h3>
<p>
<p>ControlPlaneConfig contains configuration settings for the control plane.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
gcp.provider.extensions.gardener.cloud/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>ControlPlaneConfig</code></td>
</tr>
<tr>
<td>
<code>zone</code></br>
<em>
string
</em>
</td>
<td>
<p>Zone is the GCP zone.</p>
</td>
</tr>
<tr>
<td>
<code>cloudControllerManager</code></br>
<em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.CloudControllerManagerConfig">
CloudControllerManagerConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CloudControllerManager contains configuration settings for the cloud-controller-manager.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.InfrastructureConfig">InfrastructureConfig
</h3>
<p>
<p>InfrastructureConfig infrastructure configuration resource</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
gcp.provider.extensions.gardener.cloud/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>InfrastructureConfig</code></td>
</tr>
<tr>
<td>
<code>networks</code></br>
<em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.NetworkConfig">
NetworkConfig
</a>
</em>
</td>
<td>
<p>Networks is the network configuration (VPC, subnets, etc.)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.CloudControllerManagerConfig">CloudControllerManagerConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.ControlPlaneConfig">ControlPlaneConfig</a>)
</p>
<p>
<p>CloudControllerManagerConfig contains configuration settings for the cloud-controller-manager.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>featureGates</code></br>
<em>
map[string]bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>FeatureGates contains information about enabled feature gates.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.CloudNAT">CloudNAT
</h3>
<p>
(<em>Appears on:</em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>)
</p>
<p>
<p>CloudNAT contains information about the the CloudNAT configuration</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>minPortsPerVM</code></br>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>MinPortsPerVM is the minimum number of ports allocated to a VM in the NAT config.
The default value is 2048 ports.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.CloudRouter">CloudRouter
</h3>
<p>
(<em>Appears on:</em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.VPC">VPC</a>)
</p>
<p>
<p>CloudRouter contains information about the the CloudRouter configuration</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the CloudRouter name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.InfrastructureStatus">InfrastructureStatus
</h3>
<p>
<p>InfrastructureStatus contains information about created infrastructure resources.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>networks</code></br>
<em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.NetworkStatus">
NetworkStatus
</a>
</em>
</td>
<td>
<p>Networks is the status of the networks of the infrastructure.</p>
</td>
</tr>
<tr>
<td>
<code>serviceAccountEmail</code></br>
<em>
string
</em>
</td>
<td>
<p>ServiceAccountEmail is the email address of the service account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.MachineImage">MachineImage
</h3>
<p>
(<em>Appears on:</em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.WorkerStatus">WorkerStatus</a>)
</p>
<p>
<p>MachineImage is a mapping from logical names and versions to GCP-specific identifiers.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the logical name of the machine image.</p>
</td>
</tr>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
<p>Version is the logical version of the machine image.</p>
</td>
</tr>
<tr>
<td>
<code>image</code></br>
<em>
string
</em>
</td>
<td>
<p>Image is the path to the image.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.MachineImageVersion">MachineImageVersion
</h3>
<p>
(<em>Appears on:</em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.MachineImages">MachineImages</a>)
</p>
<p>
<p>MachineImageVersion contains a version and a provider-specific identifier.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
<p>Version is the version of the image.</p>
</td>
</tr>
<tr>
<td>
<code>image</code></br>
<em>
string
</em>
</td>
<td>
<p>Image is the path to the image.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.MachineImages">MachineImages
</h3>
<p>
(<em>Appears on:</em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.CloudProfileConfig">CloudProfileConfig</a>)
</p>
<p>
<p>MachineImages is a mapping from logical names and versions to provider-specific identifiers.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the logical name of the machine image.</p>
</td>
</tr>
<tr>
<td>
<code>versions</code></br>
<em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.MachineImageVersion">
[]MachineImageVersion
</a>
</em>
</td>
<td>
<p>Versions contains versions and a provider-specific identifier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig
</h3>
<p>
(<em>Appears on:</em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.InfrastructureConfig">InfrastructureConfig</a>)
</p>
<p>
<p>NetworkConfig holds information about the Kubernetes and infrastructure networks.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>vpc</code></br>
<em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.VPC">
VPC
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>VPC indicates whether to use an existing VPC or create a new one.</p>
</td>
</tr>
<tr>
<td>
<code>cloudNAT</code></br>
<em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.CloudNAT">
CloudNAT
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CloudNAT contains configation about the the CloudNAT configuration</p>
</td>
</tr>
<tr>
<td>
<code>internal</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Internal is a private subnet (used for internal load balancers).</p>
</td>
</tr>
<tr>
<td>
<code>worker</code></br>
<em>
string
</em>
</td>
<td>
<p>Workers is the worker subnet range to create (used for the VMs).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.NetworkStatus">NetworkStatus
</h3>
<p>
(<em>Appears on:</em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.InfrastructureStatus">InfrastructureStatus</a>)
</p>
<p>
<p>NetworkStatus is the current status of the infrastructure networks.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>vpc</code></br>
<em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.VPC">
VPC
</a>
</em>
</td>
<td>
<p>VPC states the name of the infrastructure VPC.</p>
</td>
</tr>
<tr>
<td>
<code>subnets</code></br>
<em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.Subnet">
[]Subnet
</a>
</em>
</td>
<td>
<p>Subnets are the subnets that have been created.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.Subnet">Subnet
</h3>
<p>
(<em>Appears on:</em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.NetworkStatus">NetworkStatus</a>)
</p>
<p>
<p>Subnet is a subnet that was created.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the name of the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>purpose</code></br>
<em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.SubnetPurpose">
SubnetPurpose
</a>
</em>
</td>
<td>
<p>Purpose is the purpose for which the subnet was created.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.SubnetPurpose">SubnetPurpose
(<code>string</code> alias)</p></h3>
<p>
(<em>Appears on:</em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.Subnet">Subnet</a>)
</p>
<p>
<p>SubnetPurpose is a purpose of a subnet.</p>
</p>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.VPC">VPC
</h3>
<p>
(<em>Appears on:</em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.NetworkConfig">NetworkConfig</a>, 
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.NetworkStatus">NetworkStatus</a>)
</p>
<p>
<p>VPC contains information about the VPC and some related resources.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
<p>Name is the VPC name.</p>
</td>
</tr>
<tr>
<td>
<code>cloudRouter</code></br>
<em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.CloudRouter">
CloudRouter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>CloudRouter indicates whether to use an existing CloudRouter or create a new one</p>
</td>
</tr>
</tbody>
</table>
<h3 id="gcp.provider.extensions.gardener.cloud/v1alpha1.WorkerStatus">WorkerStatus
</h3>
<p>
<p>WorkerStatus contains information about created worker resources.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>machineImages</code></br>
<em>
<a href="#gcp.provider.extensions.gardener.cloud/v1alpha1.MachineImage">
[]MachineImage
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>MachineImages is a list of machine images that have been used in this worker. Usually, the extension controller
gets the mapping from name/version to the provider-specific machine image data in its componentconfig. However, if
a version that is still in use gets removed from this componentconfig it cannot reconcile anymore existing <code>Worker</code>
resources that are still using this version. Hence, it stores the used versions in the provider status to ensure
reconciliation is possible.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
