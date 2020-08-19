/*Copyright (c) 2019 Cisco and/or its affiliates.

This software is licensed to you under the terms of the Cisco Sample
Code License, Version 1.0 (the "License"). You may obtain a copy of the
License at

https://developer.cisco.com/docs/licenses

All use of the material herein must be in accordance with the terms of
the License. All rights not expressly granted by the License are
reserved. Unless required by applicable law or agreed to separately in
writing, software distributed under the License is distributed on an "AS
IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
or implied.*/

package main

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"

	"github.com/ccp-clientlibrary-go/ccp"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,

		Schema: map[string]*schema.Schema{
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"provider_client_config_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"kubernetes_version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"kube_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_allocation_method": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"master_vip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"loadbalancer_ip_num": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"subnet_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ntp_pools": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ntp_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"registries_root_ca": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			/*"registries_self_signed": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},*/
			"registries_insecure": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"docker_proxy_http": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"docker_proxy_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"docker_bip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"infra": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"datacenter": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"cluster": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"datastore": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"resource_pool": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"networks": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"master_node_pool": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"size": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"template": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"vcpus": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"memory": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"gpus": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"ssh_user": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"ssh_key": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"nodes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"status_detail": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"status_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"public_ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"private_ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"phase": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"kubernetes_version": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"worker_node_pools": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"size": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"template": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"vcpus": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"memory": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"gpus": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"ssh_user": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"ssh_key": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"nodes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"status_detail": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"status_reason": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"public_ip": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"private_ip": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"phase": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"kubernetes_version": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"network_plugin": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"details": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pod_cidr": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
			"ingress_as_lb": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"nginx_ingress_class": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"etcd_encrypted": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"skip_management": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"docker_no_proxy": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"routable_cidr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aci_profile_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"aws_iam_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceClusterCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	networkPlugins := d.Get("network_plugin").([]interface{})
	networkPluginsKeys := networkPlugins[0].(map[string]interface{})

	networkPluginName := networkPluginsKeys["name"].(string)
	networkPluginDetails := networkPluginsKeys["details"].([]interface{})
	detailsPodCIDR := networkPluginDetails[0].(map[string]interface{})

	networkPluginDetail := ccp.NetworkPluginDetails{
		PodCIDR: ccp.String(detailsPodCIDR["pod_cidr"].(string)),
	}

	networkPlugin := ccp.NetworkPlugin{
		Name:    ccp.String(networkPluginName),
		Details: &networkPluginDetail,
	}

	infrastructureKeys := d.Get("infra").([]interface{})
	infrastructure := infrastructureKeys[0].(map[string]interface{})
	networksKeys := infrastructure["networks"].([]interface{})

	var networks []string

	for _, network := range networksKeys {
		networks = append(networks, network.(string))
	}

	infra := ccp.Infra{
		Datacenter:   ccp.String(infrastructure["datacenter"].(string)),
		Cluster:      ccp.String(infrastructure["cluster"].(string)),
		ResourcePool: ccp.String(infrastructure["resource_pool"].(string)),
		Datastore:    ccp.String(infrastructure["datastore"].(string)),
		Networks:     &networks,
	}

	ntpPools := []string{}
	for _, pool := range d.Get("ntp_pools").([]interface{}) {
		ntpPools = append(ntpPools, pool.(string))
	}

	ntpServers := []string{}
	for _, server := range d.Get("ntp_servers").([]interface{}) {
		ntpServers = append(ntpServers, server.(string))
	}

	registriesRootCA := []string{}
	for _, registry := range d.Get("registries_root_ca").([]interface{}) {
		registriesRootCA = append(registriesRootCA, registry.(string))
	}

	/*selfSigned := d.Get("registries_self_signed").(map[string]interface{})

	registriesSelfSigned := ccp.RegistriesSelfSigned{
		Cert: ccp.String(selfSigned["cert"].(string)),
	}*/

	registriesInsecure := []string{}
	for _, registry := range d.Get("registries_insecure").([]interface{}) {
		registriesInsecure = append(registriesInsecure, registry.(string))
	}

	masterNodePools := d.Get("master_node_pool").([]interface{})
	masterNode := masterNodePools[0].(map[string]interface{})

	gpuKeys := masterNode["gpus"].([]interface{})

	var gpus []string

	for _, gpu := range gpuKeys {
		gpus = append(gpus, gpu.(string))
	}

	var nodePool []ccp.Node
	nodeKeys := masterNode["nodes"].([]interface{})

	for _, node := range nodeKeys {

		tmpNode := node.(map[string]interface{})

		nodes := ccp.Node{
			Name:         ccp.String(tmpNode["name"].(string)),
			Status:       ccp.String(tmpNode["status"].(string)),
			StatusDetail: ccp.String(tmpNode["status_detail"].(string)),
			StatusReason: ccp.String(tmpNode["status_reason"].(string)),
			PublicIP:     ccp.String(tmpNode["public_ip"].(string)),
			PrivateIP:    ccp.String(tmpNode["private_ip"].(string)),
			Phase:        ccp.String(tmpNode["phase"].(string)),
		}

		nodePool = append(nodePool, nodes)
	}

	masterNodePool := ccp.MasterNodePool{
		Name:     ccp.String(masterNode["name"].(string)),
		Size:     ccp.Int64(int64(masterNode["size"].(int))),
		Template: ccp.String(masterNode["template"].(string)),
		VCPUs:    ccp.Int64(int64(masterNode["vcpus"].(int))),
		Memory:   ccp.Int64(int64(masterNode["memory"].(int))),
		//GPUs:              &gpus,
		SSHUser:           ccp.String(masterNode["ssh_user"].(string)),
		SSHKey:            ccp.String(masterNode["ssh_key"].(string)),
		Nodes:             &nodePool,
		KubernetesVersion: ccp.String(masterNode["kubernetes_version"].(string)),
	}

	var workerPool []ccp.WorkerNodePool

	workerNodePools := d.Get("worker_node_pools").([]interface{})

	for _, workerNode := range workerNodePools {

		worker := workerNode.(map[string]interface{})

		gpuKeys := worker["gpus"].([]interface{})

		var gpus []string

		for _, gpu := range gpuKeys {
			gpus = append(gpus, gpu.(string))
		}

		var nodePool []ccp.Node
		nodeKeys := worker["nodes"].([]interface{})

		for _, node := range nodeKeys {

			tmpNode := node.(map[string]interface{})

			nodes := ccp.Node{
				Name:         ccp.String(tmpNode["name"].(string)),
				Status:       ccp.String(tmpNode["status"].(string)),
				StatusDetail: ccp.String(tmpNode["status_detail"].(string)),
				StatusReason: ccp.String(tmpNode["status_reason"].(string)),
				PublicIP:     ccp.String(tmpNode["public_ip"].(string)),
				PrivateIP:    ccp.String(tmpNode["private_ip"].(string)),
				Phase:        ccp.String(tmpNode["phase"].(string)),
			}

			nodePool = append(nodePool, nodes)
		}

		workerNodePool := ccp.WorkerNodePool{
			Name:     ccp.String(worker["name"].(string)),
			Size:     ccp.Int64(int64(worker["size"].(int))),
			Template: ccp.String(worker["template"].(string)),
			VCPUs:    ccp.Int64(int64(worker["vcpus"].(int))),
			Memory:   ccp.Int64(int64(worker["memory"].(int))),
			//GPUs:              &gpus,
			SSHUser:           ccp.String(worker["ssh_user"].(string)),
			SSHKey:            ccp.String(worker["ssh_key"].(string)),
			Nodes:             &nodePool,
			KubernetesVersion: ccp.String(worker["kubernetes_version"].(string)),
		}

		workerPool = append(workerPool, workerNodePool)
	}

	dockerNoProxy := []string{}
	for _, proxy := range d.Get("docker_no_proxy").([]interface{}) {
		dockerNoProxy = append(dockerNoProxy, proxy.(string))
	}

	newCluster := ccp.Cluster{

		Type:               ccp.String(d.Get("type").(string)),
		Name:               ccp.String(d.Get("name").(string)),
		InfraProviderUUID:  ccp.String(d.Get("provider_client_config_uuid").(string)),
		Status:             ccp.String(d.Get("status").(string)),
		KubernetesVersion:  ccp.String(d.Get("kubernetes_version").(string)),
		KubeConfig:         ccp.String(d.Get("kube_config").(string)),
		IPAllocationMethod: ccp.String(d.Get("ip_allocation_method").(string)),
		MasterVIP:          ccp.String(d.Get("master_vip").(string)),
		LoadBalancerIPNum:  ccp.Int64(int64(d.Get("loadbalancer_ip_num").(int))),
		SubnetUUID:         ccp.String(d.Get("subnet_uuid").(string)),
		NTPPools:           &ntpPools,
		NTPServers:         &ntpServers,
		RegistriesRootCA:   &registriesRootCA,
		/*RegistriesSelfSigned: &registriesSelfSigned,*/
		RegistriesInsecure: &registriesInsecure,
		DockerProxyHTTP:    ccp.String(d.Get("docker_proxy_http").(string)),
		DockerProxyHTTPS:   ccp.String(d.Get("docker_proxy_https").(string)),
		DockerBIP:          ccp.String(d.Get("docker_bip").(string)),
		Infra:              &infra,
		MasterNodePool:     &masterNodePool,
		WorkerNodePool:     &workerPool,
		NetworkPlugin:      &networkPlugin,
		IngressAsLB:        ccp.Bool(d.Get("ingress_as_lb").(bool)),
		NginxIngressClass:  ccp.String(d.Get("nginx_ingress_class").(string)),
		ETCDEncrypted:      ccp.Bool(d.Get("etcd_encrypted").(bool)),
		SkipManagement:     ccp.Bool(d.Get("skip_management").(bool)),
		DockerNoProxy:      &dockerNoProxy,
		RoutableCIDR:       ccp.String(d.Get("routable_cidr").(string)),
		ImagePrefix:        ccp.String(d.Get("image_prefix").(string)),
		ACIProfileUUID:     ccp.String(d.Get("aci_profile_uuid").(string)),
		Description:        ccp.String(d.Get("description").(string)),
		AWSIamEnabled:      ccp.Bool(d.Get("aws_iam_enabled").(bool)),
	}

	cluster, err := client.AddCluster(&newCluster)

	log.Printf(" [DEBUG] ***************** CLUSTER: %+v", err)

	log.Printf(" [DEBUG] ***************** CLUSTER: %+v", cluster)

	if err != nil {
		return errors.New(err.Error())
	}

	uuid := *cluster.UUID

	log.Printf(" [DEBUG] ***************** CLUSTER UUID: %+v", *cluster.UUID)

	d.SetId(uuid)

	cluster, err = client.GetClusterByName(d.Get("name").(string))

	if err != nil {
		return errors.New(err.Error())
	}

	return setClusterResourceData(d, cluster)
}

func resourceClusterRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	cluster, err := client.GetClusterByName(d.Get("name").(string))

	if err != nil {
		return errors.New("UNABLE TO RETRIEVE DETAILS FOR CLUSTER: " + d.Get("name").(string))
	}

	return setClusterResourceData(d, cluster)

}

func resourceClusterUpdate(d *schema.ResourceData, m interface{}) error {

	/*client := m.(*ccp.Client)

	newCluster := ccp.Cluster{
		UUID:              ccp.String(d.Get("uuid").(string)),
		Workers:           ccp.Int64(int64(d.Get("workers").(int))),
		LoadBalancerIPNum: ccp.Int64(int64(d.Get("loadbalancer_ip_num").(int))),
	}

	cluster, err := client.PatchCluster(&newCluster)

	if err != nil {
		return errors.New(err.Error())
	}

	cluster, err = client.GetCluster(d.Get("name").(string))

	if err != nil {
		return errors.New("UNABLE TO RETRIEVE DETAILS FOR CLUSTER: " + d.Get("name").(string))
	}

	return setClusterResourceData(d, cluster)*/

	return nil

}

func resourceClusterDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	err := client.DeleteCluster(d.Get("uuid").(string))

	if err != nil {
		return errors.New(err.Error())
	}

	d.SetId("")
	return nil
}

func setClusterResourceData(d *schema.ResourceData, u *ccp.Cluster) error {

	log.Printf(" [DEBUG] ***************** setClusterResourceData: %+v", u)
	log.Printf(" [DEBUG] ***************** setClusterResourceData: %T", u)
	log.Printf(" [DEBUG] ***************** setClusterResourceData: %+v", u.Type)
	log.Printf(" [DEBUG] ***************** setClusterResourceData: %T", u.Type)

	b, _ := json.Marshal(u.NetworkPlugin)

	var f interface{}
	_ = json.Unmarshal(b, &f)

	m := f.(map[string]interface{})

	pairs := [][]string{}
	for key, value := range m {
		log.Printf(" [DEBUG] ***************** keyType, value: %T, %T", key, value)
		log.Printf(" [DEBUG] ***************** key, value: %s, %s", key, value)
		//pairs = append(pairs, []string{key, value.})
	}

	log.Printf(" [DEBUG] ***************** setClusterResourceData: %+v", u.Infra)
	log.Printf(" [DEBUG] ***************** setClusterResourceData: %T", u.Infra)

	log.Printf(" [DEBUG] ***************** setClusterResourceData: %+v", b)
	log.Printf(" [DEBUG] ***************** setClusterResourceData: %T", b)

	log.Printf(" [DEBUG] ***************** setClusterResourceData: %+v", m)
	log.Printf(" [DEBUG] ***************** setClusterResourceData: %T", m)

	log.Printf(" [DEBUG] ***************** setClusterResourceData: %+v", pairs)
	log.Printf(" [DEBUG] ***************** setClusterResourceData: %T", pairs)

	var infra []interface{}
	output := structToMap(u.Infra)

	infra = append(infra, output)

	log.Printf(" [DEBUG] ***************** infra: %T", infra)
	log.Printf(" [DEBUG] ***************** infra: %+v", infra)

	if err := d.Set("uuid", u.UUID); err != nil {
		return errors.New("CANNOT SET UUID")
	}
	if err := d.Set("type", u.Type); err != nil {
		return errors.New("CANNOT SET TYPE")
	}
	if err := d.Set("name", u.Name); err != nil {
		return errors.New("CANNOT SET NAME")
	}
	if err := d.Set("provider_client_config_uuid", u.InfraProviderUUID); err != nil {
		return errors.New("CANNOT SET PROVIDER CLIENT CONFIG UUID")
	}
	if err := d.Set("status", u.Status); err != nil {
		return errors.New("CANNOT SET STATUS")
	}
	if err := d.Set("kubernetes_version", u.KubernetesVersion); err != nil {
		return errors.New("CANNOT SET KUBERNETES VERSION")
	}
	if err := d.Set("kube_config", u.KubeConfig); err != nil {
		return errors.New("CANNOT SET KUBECONFIG")
	}
	if err := d.Set("ip_allocation_method", u.IPAllocationMethod); err != nil {
		return errors.New("CANNOT SET IP ALLOCATION METHOD")
	}
	if err := d.Set("master_vip", u.MasterVIP); err != nil {
		return errors.New("CANNOT SET MASTER VIP")
	}
	if err := d.Set("loadbalancer_ip_num", u.LoadBalancerIPNum); err != nil {
		return errors.New("CANNOT SET NUMBER OF LOAD BALANCERS")
	}
	if err := d.Set("subnet_uuid", u.SubnetUUID); err != nil {
		return errors.New("CANNOT SET SUBNET ID")
	}
	if err := d.Set("ntp_pools", u.NTPPools); err != nil {
		return errors.New("CANNOT SET NTP Pool")
	}
	if err := d.Set("ntp_servers", u.NTPServers); err != nil {
		return errors.New("CANNOT SET NTP SERVERS")
	}
	if err := d.Set("registries_root_ca", u.RegistriesRootCA); err != nil {
		return errors.New("CANNOT SET REGISTRIES ROOT CA")
	}
	/*if err := d.Set("registries_self_signed", d.Get("registries_self_signed").(map[string]interface{})); err != nil {
		return errors.New("CANNOT SET SELF SIGNED REGISTRIES")
	}*/
	if err := d.Set("registries_insecure", u.RegistriesInsecure); err != nil {
		return errors.New("CANNOT SET INSECURE REGISTRIES")
	}
	if err := d.Set("docker_proxy_http", u.DockerProxyHTTP); err != nil {
		return errors.New("CANNOT SET HTTP DOCKER PROXY")
	}
	if err := d.Set("docker_proxy_https", u.DockerProxyHTTPS); err != nil {
		return errors.New("CANNOT SET HTTPS DOCKER PROXY")
	}
	if err := d.Set("docker_bip", u.DockerBIP); err != nil {
		return errors.New("CANNOT SET DOCKER BIP")
	}

	/*if err := d.Set("infra", infra); err != nil {
		log.Printf("[DEBUG] *********** INFRA: %+v", err)
		return errors.New("CANNOT SET INFRA")
	}
	if err := d.Set("master_node_pool", u.MasterNodePool); err != nil {
		return errors.New("CANNOT SET MASTER NODE POOL")
	}
	if err := d.Set("worker_node_pools", u.WorkerNodePool); err != nil {
		return errors.New("CANNOT SET WORKER NODE POOL")
	}
	if err := d.Set("network_plugin", u.NetworkPlugin); err != nil {
		return errors.New("CANNOT SET NETWORK PLUGIN")
	}
	if err := d.Set("networks", u.Infra.Networks); err != nil {
		return errors.New("CANNOT SET NETWORKS")
	}*/

	if err := d.Set("ingress_as_lb", u.IngressAsLB); err != nil {
		return errors.New("CANNOT SET INGRESS AS LB VALUE")
	}
	if err := d.Set("nginx_ingress_class", u.NginxIngressClass); err != nil {
		return errors.New("CANNOT SET NGINX INGRESS CLASS")
	}
	if err := d.Set("etcd_encrypted", u.ETCDEncrypted); err != nil {
		return errors.New("CANNOT SET ETCD ENCRYPTED")
	}
	if err := d.Set("skip_management", u.SkipManagement); err != nil {
		return errors.New("CANNOT SET SKIP MANAGEMENT VALUE")
	}
	if err := d.Set("docker_no_proxy", u.DockerNoProxy); err != nil {
		return errors.New("CANNOT SET DOCKER NO PROXY")
	}
	if err := d.Set("routable_cidr", u.RoutableCIDR); err != nil {
		return errors.New("CANNOT SET ROUTABLE CIDR")
	}
	if err := d.Set("image_prefix", u.ImagePrefix); err != nil {
		return errors.New("CANNOT SET IMAGE PREFIX")
	}
	if err := d.Set("aci_profile_uuid", u.ACIProfileUUID); err != nil {
		return errors.New("CANNOT SET ACI PROFILE UUID")
	}
	if err := d.Set("description", u.Description); err != nil {
		return errors.New("CANNOT SET DESCRIPTION")
	}
	if err := d.Set("aws_iam_enabled", u.AWSIamEnabled); err != nil {
		return errors.New("CANNOT SET AWS IAM VALUE")
	}

	return nil
}

func structToMap(item interface{}) map[string]interface{} {

	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = structToMap(field)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}
