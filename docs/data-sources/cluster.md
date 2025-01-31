---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "castai_cluster Data Source - terraform-provider-castai"
subcategory: ""
description: |-
  
---

# castai_cluster (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) The ID of this resource.

### Read-Only

- **autoscaler_policies** (List of Object) (see [below for nested schema](#nestedatt--autoscaler_policies))
- **credentials** (Set of String)
- **kubeconfig** (List of Object) (see [below for nested schema](#nestedatt--kubeconfig))
- **name** (String)
- **region** (String)
- **status** (String)

<a id="nestedatt--autoscaler_policies"></a>
### Nested Schema for `autoscaler_policies`

Read-Only:

- **cluster_limits** (List of Object) (see [below for nested schema](#nestedobjatt--autoscaler_policies--cluster_limits))
- **enabled** (Boolean)
- **node_downscaler** (List of Object) (see [below for nested schema](#nestedobjatt--autoscaler_policies--node_downscaler))
- **spot_instances** (List of Object) (see [below for nested schema](#nestedobjatt--autoscaler_policies--spot_instances))
- **unschedulable_pods** (List of Object) (see [below for nested schema](#nestedobjatt--autoscaler_policies--unschedulable_pods))

<a id="nestedobjatt--autoscaler_policies--cluster_limits"></a>
### Nested Schema for `autoscaler_policies.cluster_limits`

Read-Only:

- **cpu** (List of Object) (see [below for nested schema](#nestedobjatt--autoscaler_policies--cluster_limits--cpu))
- **enabled** (Boolean)

<a id="nestedobjatt--autoscaler_policies--cluster_limits--cpu"></a>
### Nested Schema for `autoscaler_policies.cluster_limits.cpu`

Read-Only:

- **max_cores** (Number)
- **min_cores** (Number)



<a id="nestedobjatt--autoscaler_policies--node_downscaler"></a>
### Nested Schema for `autoscaler_policies.node_downscaler`

Read-Only:

- **empty_nodes** (List of Object) (see [below for nested schema](#nestedobjatt--autoscaler_policies--node_downscaler--empty_nodes))

<a id="nestedobjatt--autoscaler_policies--node_downscaler--empty_nodes"></a>
### Nested Schema for `autoscaler_policies.node_downscaler.empty_nodes`

Read-Only:

- **delay_seconds** (Number)
- **enabled** (Boolean)



<a id="nestedobjatt--autoscaler_policies--spot_instances"></a>
### Nested Schema for `autoscaler_policies.spot_instances`

Read-Only:

- **clouds** (List of String)
- **enabled** (Boolean)


<a id="nestedobjatt--autoscaler_policies--unschedulable_pods"></a>
### Nested Schema for `autoscaler_policies.unschedulable_pods`

Read-Only:

- **enabled** (Boolean)
- **headroom** (List of Object) (see [below for nested schema](#nestedobjatt--autoscaler_policies--unschedulable_pods--headroom))
- **node_constraints** (List of Object) (see [below for nested schema](#nestedobjatt--autoscaler_policies--unschedulable_pods--node_constraints))

<a id="nestedobjatt--autoscaler_policies--unschedulable_pods--headroom"></a>
### Nested Schema for `autoscaler_policies.unschedulable_pods.headroom`

Read-Only:

- **cpu_percentage** (Number)
- **enabled** (Boolean)
- **memory_percentage** (Number)


<a id="nestedobjatt--autoscaler_policies--unschedulable_pods--node_constraints"></a>
### Nested Schema for `autoscaler_policies.unschedulable_pods.node_constraints`

Read-Only:

- **enabled** (Boolean)
- **max_node_cpu_cores** (Number)
- **max_node_ram_gib** (Number)
- **min_node_cpu_cores** (Number)
- **min_node_ram_gib** (Number)




<a id="nestedatt--kubeconfig"></a>
### Nested Schema for `kubeconfig`

Read-Only:

- **client_certificate** (String)
- **client_key** (String)
- **cluster_ca_certificate** (String)
- **host** (String)
- **raw_config** (String)


