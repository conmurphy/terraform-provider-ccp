/*
    These are the credentials used to login to CCP
*/

variable "username" {
    type = "string"
    default="my_ccp_admin_account"
}

variable "password" {
    type = "string"
    default="my_ccp_password"
}

variable "base_url" {
    type = "string"
    default="https://my_ccp_url:ccp_port"
}

provider "ccp" {
    username = "${var.username}"
    password = "${var.password}"
    base_url = "${var.base_url}"
}

/*
    This will create a new local user within CCP
*/
resource "ccp_user" "user" {
    firstname       = "Terrafom"
    lastname        = "Plugin"
    password        = "myPassword"
    username        = "builtByTerraform"
    role            = "Administrator"
}


/*
    This will create a new cluster within CCP
*/

variable "networks" {
   type = "list"
   default = ["my_ccp_network","k8-priv-iscsivm-network"]
}
      
variable "worker_node_pool" {
   type = "map"
   default = 
      {
        vcpus = 2,
        memory = 4096,
        template = "ccp-tenant-image-1.10.1-ubuntu16-1.5.0"
      }
}

variable "master_node_pool" {
   type = "map"
   default = 
      {
        vcpus = 2,
        memory = 4192,
        template = "ccp-tenant-image-1.10.1-ubuntu16-1.5.0"
      }
}

variable "providers" {
   type = "map"
   default = 
      {
        vsphere_datacenter          = "my_ccp_datacenter",
        vsphere_datastore           = "my_ccp_datastore",
        vsphere_client_config_uuid  = "aaa123-bbb123-ccc123-ddd123-eee123"
        vsphere_working_dir         = "ccp_working_directory"
      }
}

variable "deployer" {
   type = "map"
   default = 
      {
        provider_type = "vsphere"
      }  
}

variable "network_plugin" {
   type = "map"
   default = 
      {
        name = "contiv-vpp",
        status = "",
        details = ""
      }
}

variable "infra" {
   type = "map"
   default = 
      {
        datacenter	                = "my_vsphere_datacenter"
        cluster	                    = "my_vsphere_cluster"
        datastore	                  = "my_vsphere_datastore"
        resource_pool               = "my_vsphere_resource_pool"
      }
}

resource "ccp_cluster" "cluster" {
    provider_client_config_uuid = "123aaa-123bbb-123ccc-123ddd-123eee"
    name	                      = "builtByTerraform"
    is_harbor_enabled	          = false
    is_istio_enabled	          = true
    kubernetes_version	        = "1.10.1"
    loadbalancer_ip_num	        = 1
    masters                 	  = 1
    workers	                    = 3
    ssh_user	                  = "ccpuser"
    ssh_key	                    = "ssh-rsa AAA123bbb123CCC123 my_username@localhost"
    type	                      = 1
    master_node_pool            = "${var.master_node_pool}"
    worker_node_pool            = "${var.worker_node_pool}"
    networks                    = "${var.networks}"
    deployer                    = "${var.deployer}"
    providers                   = "${var.providers}"
    network_plugin              = "${var.network_plugin}"
    infra                       = "${var.infra}"
    ingress_vip_pool_id         = "aaa-bbb-ccc-ddd-eee"      
}