# Cisco Container Platform Provider Plugin for Terraform

This is a provider plugin for Terraform which allows Terraform to interact with various Cisco Container Platform (CCP) resources. 

It is currently a __Proof of Concept__ and has been developed and tested against Cisco Container Platform 1.5 with Go version 1.10 and Terraform version v0.10.2

Table of Contents
=================

  * [CCP Terraform Provider Plugin](#ccp-terraform-provider-plugin)
      * [Quick Start](#quick-start)
      * [Resources](#resources)
         * [Cluster](#cluster)
         * [User](#user)
       
Created by [gh-md-toc](https://github.com/ekalinin/github-markdown-toc)

## Quick Start

```golang

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

```

## Resources

- [Cluster](#cluster)
- [User](#user)

### User

```go
	token     *string 
	username  *string 
	disable   *bool  
	role      *string 
	firstname *string
	lastname  *string
	password  *string
```

##### __Required Fields__
* Username
* Role

##### __Important Notes__
* If ```Username``` is updated the resource will be destroyed and recreated within CCP

##### Example
```go  
resource "ccp_user" "user" {
    firstname       = "Terrafom"
    lastname        = "Plugin"
    password        = "myPassword"
    username        = "builtByTerraform"
    role            = "Administrator"
}
```

### Cluster

```go
type Cluster struct {
	uuid                          *string  
	provider_client_config_uuid   *string 
	aci_profile_uuid              *string
	name                          *string 
	description                   *string  
	workers                       *int64  
	masters                       *int64
	resource_pool                 *string               
	networks                      *[]string            
	vcpus                         *int64               
	memory                        *int64                
	type                          *int64          
	datacenter                    *string            
	cluster                       *string              
	datastore                     *string          
	state                         *string 
	template                      *string
	ssh_user                      *string 
	ssh_password                  *string 
	ssh_key                       *string 
	labels                        *[]label 
	nodes                         *[]node   
	deployer                      *kube_adm              
	kubernetes_version            *string               
	cluster_env_url               *string               
	cluster_dashboard_url         *string               
	network_plugin                *network_plugin
	ccp_private_ssh_key           *string              
	ccp_public_ssh_key            *string              
	ntp_pools                     *[]string       
	netp_servers                  *[]string      
	is_control_cluster            *bool             
	is_adopt                      *bool              
	registries_self_signed        *[]string           
	registries_insecure           *[]string            
	registries_root_ca            *[]string          
	ingress_vip_pool_id           *string             
	ingress_vip_addr_id           *string              
	ingress_vips                  *[]string             
	keepalived_vrid               *int64              
	helm_charts                   *[]helm_chart    
	master_vip_addr_id            *string          
	master_vip                    *string        
	master_mac_addresses          *[]string      
	cluster_health_status         *string       
	auth_list                     *[]string 
	is_harbor_enabled             *bool           
	harbor_admin_server_password  *string        
	harbor_registry_size          *string        
	load_balancer_ip_num          *int64          
	is_istio_enabled              *bool          
	worker_node_pool              *worker_node_pool  
	master_node_pool              *master_node_pool  
	infra                         *infra 
}

type infra struct {
	datacenter                 *string   
	datastore                  *string  
	cluster                    *string   
	networks                   *[]string
	resource_pool              *string   
}

type label struct {
	key                        *string  
	value                      *string  
}

type node struct {
	uuid                       *string   
	name                       *string   
	public_ip                  *string    
	private_ip     		         *string   
	is_master    		           *bool  
	state     	               *string   
	cloud_init_data  		       *string    
	kubernetes_version         *string   
	error_log        	         *string   
	template       	           *string   
	mac_addresses              *[]string  
}

type deployer struct {
	proxy_cmd                  *string    
	provider_type              *string   
	provider                   *provider 
	ip                         *string  
	port                       *int64   
	username                   *string    
	password                   *string    

type network_plugin struct {
	name   			               *string  
	status 			               *string  
	details			               *string  
}

type helm_chart struct {
	helmchart_uuid		         *string  
	cluster_UUID 		           *string  
	chart_url    		           *string  
	name         		           *string  
	options     		           *string  
}	

type provider struct {
	vsphere_datacenter            *string             
	vsphere_datastore             *string             
	vsphere_scsi_controller_type  *string           
	vsphere_working_dir           *string           
	vsphere_client_config_uuid    *string          
	client_config                 *vsphere_client_config  
}

type vsphere_client_config struct {
	ip       		               *string  
	port     		               *int64  
	username 		               *string  
	password 		               *string  
}

type worker_node_pool struct {
	vcpus   		               *int64   
	memory  		               *int64   
	template		               *string  
}

type master_node_pool struct {
	vcpus    		               *int64   
	memory   		               *int64   
	template 		               *string  
}
```

##### __Required Fields__
* ProviderClientConfigUUID
* Name
* KubernetesVersion
* ResourcePool
* Networks
* SSHKey
* Datacenter
* Cluster
* Datastore
* Workers
* SSHUser
* Type
* Masters
* Deployer
  * ProviderType
* Provider 
    * VsphereDataCenter
    * VsphereClientConfigUUID
    * VsphereDatastore
    * VsphereWorkingDir
* NetworkPlugin
  * Name 
  * Status
  * Details
* IsHarborEnabled         
* LoadBalancerIPNum                
* IsIstioEnabled             
* WorkerNodePool    
  * VCPUs    
  * Memory  
  * Template 
* MasterNodePool           
  * VCPUs    
  * Memory  
  * Template 

##### __Important Notes__
* Only the following fields can be changed once the resource has been created
  * workers
  * loadbalancer_ip_num
* If additional fields are changed they will not be updated using ```terraform apply```

##### Example
```go  
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
```


DISCLAIMER:

These scripts are meant for educational/proof of concept purposes only. Any use of these scripts and tools is at your own risk. There is no guarantee that they have been through thorough testing in a comparable environment and we are not responsible for any damage or data loss incurred with their use.
