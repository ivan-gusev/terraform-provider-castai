data "google_client_config" "default" {}

data "google_container_cluster" "this" {
  name =  var.cluster_name
  location = var.cluster_location
  project = var.project_id
}

provider "helm" {
  kubernetes {
    host                   = "https://${data.google_container_cluster.this.endpoint}"
    token                  = data.google_client_config.default.access_token
    cluster_ca_certificate = base64decode(data.google_container_cluster.this.master_auth[0].cluster_ca_certificate)
  }
}

provider "castai" {
  api_token = var.castai_api_token
}

module "castai-gke-iam" {
  source = "castai/gke-iam/castai"

  project_id       = var.project_id
  gke_cluster_name = var.cluster_name

}

module "castai-gke-cluster" {
  source = "castai/gke-cluster/castai"

  project_id         = var.project_id
  gke_cluster_name   = var.cluster_name
  gke_cluster_location = var.cluster_location

  gke_credentials            = module.castai-gke-iam.private_key
  delete_nodes_on_disconnect = true

  # Full schema can be found here https://api.cast.ai/v1/spec/#/PoliciesAPI/PoliciesAPIUpsertClusterPolicies
  autoscaler_policies_json = <<-EOT
    {
        "enabled": true,
        "isScopedMode": true,
        "unschedulablePods": {
            "enabled": true
        },
        "spotInstances": {
            "enabled": true,
            "clouds": ["gcp"],
            "spotBackups": {
                "enabled": true
            }
        },
        "nodeDownscaler": {
            "emptyNodes": {
                "enabled": true
            }
        }
    }
  EOT
}
