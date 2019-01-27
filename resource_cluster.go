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
	"errors"
	"strconv"

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
			"provider_client_config_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"aci_profile_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"kubernetes_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ssh_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"masters": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"workers": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ssh_user": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"is_harbor_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"is_istio_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"loadbalancer_ip_num": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ingress_vip_addr_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ingress_vip_pool_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_plugin": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
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
						"details": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"worker_node_pool": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vcpus": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"memory": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"template": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"master_node_pool": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vcpus": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"memory": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"template": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"infra": &schema.Schema{
				Type:     schema.TypeMap,
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
					},
				},
			},
			"deployer": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider_type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"providers": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vsphere_datacenter": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"vsphere_datastore": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"vsphere_client_config_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"vsphere_working_dir": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceClusterCreate(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	networkPlugins := d.Get("network_plugin").(map[string]interface{})

	networkPlugin := ccp.NetworkPlugin{
		Name:    ccp.String(networkPlugins["name"].(string)),
		Status:  ccp.String(networkPlugins["status"].(string)),
		Details: ccp.String(networkPlugins["details"].(string)),
	}

	providers := d.Get("providers").(map[string]interface{})

	provider := ccp.Provider{
		VsphereDataCenter:       ccp.String(providers["vsphere_datacenter"].(string)),
		VsphereDatastore:        ccp.String(providers["vsphere_datastore"].(string)),
		VsphereClientConfigUUID: ccp.String(providers["vsphere_client_config_uuid"].(string)),
		VsphereWorkingDir:       ccp.String(providers["vsphere_working_dir"].(string)),
	}

	deployers := d.Get("deployer").(map[string]interface{})

	deployer := ccp.Deployer{
		ProviderType: ccp.String(deployers["provider_type"].(string)),
		Provider:     &provider,
	}

	workerNodePools := d.Get("worker_node_pool").(map[string]interface{})

	// it seems like ints are converted to strings when using the schema:map[string] above. therefore need to convert back to an int64
	vcpus, err := strconv.ParseInt(workerNodePools["vcpus"].(string), 10, 64)
	memory, err := strconv.ParseInt(workerNodePools["memory"].(string), 10, 64)

	workerNodePool := ccp.WorkerNodePool{
		VCPUs:    ccp.Int64(vcpus),
		Memory:   ccp.Int64(memory),
		Template: ccp.String(workerNodePools["template"].(string)),
	}

	masterNodePools := d.Get("master_node_pool").(map[string]interface{})

	vcpus, err = strconv.ParseInt(masterNodePools["vcpus"].(string), 10, 64)
	memory, err = strconv.ParseInt(masterNodePools["memory"].(string), 10, 64)

	masterNodePool := ccp.MasterNodePool{
		VCPUs:    ccp.Int64(vcpus),
		Memory:   ccp.Int64(memory),
		Template: ccp.String(masterNodePools["template"].(string)),
	}

	networks := []string{}
	for _, network := range d.Get("networks").([]interface{}) {
		networks = append(networks, network.(string))
	}

	infrastructure := d.Get("infra").(map[string]interface{})

	infra := ccp.Infra{
		Networks:     &networks,
		Datacenter:   ccp.String(infrastructure["datacenter"].(string)),
		Cluster:      ccp.String(infrastructure["cluster"].(string)),
		ResourcePool: ccp.String(infrastructure["resource_pool"].(string)),
		Datastore:    ccp.String(infrastructure["datastore"].(string)),
	}

	newCluster := ccp.Cluster{

		ProviderClientConfigUUID: ccp.String(d.Get("provider_client_config_uuid").(string)),
		Name:              ccp.String(d.Get("name").(string)),
		KubernetesVersion: ccp.String(d.Get("kubernetes_version").(string)),
		SSHKey:            ccp.String(d.Get("ssh_key").(string)),
		Networks:          &networks,
		Datacenter:        ccp.String(infrastructure["datacenter"].(string)),
		Cluster:           ccp.String(infrastructure["cluster"].(string)),
		ResourcePool:      ccp.String(infrastructure["resource_pool"].(string)),
		Datastore:         ccp.String(infrastructure["datastore"].(string)),
		Masters:           ccp.Int64(int64(d.Get("masters").(int))),
		Workers:           ccp.Int64(int64(d.Get("workers").(int))),
		SSHUser:           ccp.String(d.Get("ssh_user").(string)),
		Type:              ccp.Int64(int64(d.Get("type").(int))),
		//DeployerType:      ccp.String(d.Get("deployer_type").(string)),
		IsHarborEnabled:   ccp.Bool(d.Get("is_harbor_enabled").(bool)),
		IsIstioEnabled:    ccp.Bool(d.Get("is_istio_enabled").(bool)),
		LoadBalancerIPNum: ccp.Int64(int64(d.Get("loadbalancer_ip_num").(int))),
		Infra:             &infra,
		NetworkPlugin:     &networkPlugin,
		Deployer:          &deployer,
		MasterNodePool:    &masterNodePool,
		WorkerNodePool:    &workerNodePool,
		IngressVIPPoolID:  ccp.String(d.Get("ingress_vip_pool_id").(string)),
	}

	cluster, err := client.AddCluster(&newCluster)

	if err != nil {
		return errors.New(err.Error())
	}

	uuid := *cluster.UUID
	d.SetId(uuid)

	cluster, err = client.GetCluster(d.Get("name").(string))

	if err != nil {
		return errors.New(err.Error())
	}

	return setClusterResourceData(d, cluster)
}

func resourceClusterRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

	cluster, err := client.GetCluster(d.Get("name").(string))

	if err != nil {
		return errors.New("UNABLE TO RETRIEVE DETAILS FOR CLUSTER: " + d.Get("name").(string))
	}

	return setClusterResourceData(d, cluster)

}

func resourceClusterUpdate(d *schema.ResourceData, m interface{}) error {

	client := m.(*ccp.Client)

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

	return setClusterResourceData(d, cluster)

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

	if err := d.Set("uuid", u.UUID); err != nil {
		return errors.New("CANNOT SET UUID")
	}
	if err := d.Set("provider_client_config_uuid", u.ProviderClientConfigUUID); err != nil {
		return errors.New("CANNOT SET PROVIDER CLIENT CONFIG UUID")
	}
	if err := d.Set("aci_profile_uuid", u.ACIProfileUUID); err != nil {
		return errors.New("CANNOT SET ACI PROFILE UUID")
	}
	if err := d.Set("name", u.Name); err != nil {
		return errors.New("CANNOT SET NAME")
	}
	if err := d.Set("description", u.Description); err != nil {
		return errors.New("CANNOT SET DESCRIPTION")
	}
	if err := d.Set("kubernetes_version", u.KubernetesVersion); err != nil {
		return errors.New("CANNOT SET KUBERNETES VERSION")
	}
	if err := d.Set("ssh_key", u.SSHKey); err != nil {
		return errors.New("CANNOT SET SSH KEY")
	}
	if err := d.Set("masters", u.Masters); err != nil {
		return errors.New("CANNOT SET NUMBER OF MASTERS")
	}
	if err := d.Set("workers", u.Workers); err != nil {
		return errors.New("CANNOT SET NUMBER OF WORKERS")
	}
	if err := d.Set("ssh_user", u.SSHUser); err != nil {
		return errors.New("CANNOT SET SSH USER")
	}
	if err := d.Set("type", u.Type); err != nil {
		return errors.New("CANNOT SET TYPE")
	}
	if err := d.Set("is_harbor_enabled", u.IsHarborEnabled); err != nil {
		return errors.New("CANNOT SET 'IS HARBOR ENABLED' FIELD")
	}
	if err := d.Set("is_istio_enabled", u.IsIstioEnabled); err != nil {
		return errors.New("CANNOT SET 'IS ISTIO ENABLED' FIELD")
	}
	if err := d.Set("loadbalancer_ip_num", u.LoadBalancerIPNum); err != nil {
		return errors.New("CANNOT SET LOAD BALANCER IP NUMBER")
	}
	if err := d.Set("network_plugin", d.Get("network_plugin").(map[string]interface{})); err != nil {
		return errors.New("CANNOT SET NETWORK PLUGIN")
	}
	if err := d.Set("worker_node_pool", d.Get("worker_node_pool").(map[string]interface{})); err != nil {
		return errors.New("CANNOT SET WORKER NODE POOL")
	}
	if err := d.Set("master_node_pool", d.Get("master_node_pool").(map[string]interface{})); err != nil {
		return errors.New("CANNOT SET MASTER NODE POOL")
	}
	if err := d.Set("deployer", d.Get("deployer").(map[string]interface{})); err != nil {
		return errors.New("CANNOT SET DEPLOYER")
	}
	if err := d.Set("providers", d.Get("providers").(map[string]interface{})); err != nil {
		return errors.New("CANNOT SET PROVIDER")
	}
	if err := d.Set("infra", d.Get("infra").(map[string]interface{})); err != nil {
		return errors.New("CANNOT SET INFRA")
	}
	if err := d.Set("ingress_vip_addr_id", u.IngressVIPAddrID); err != nil {
		return errors.New("CANNOT SET INGRESS VIP ADDRESS ID")
	}
	if err := d.Set("ingress_vip_pool_id", u.IngressVIPPoolID); err != nil {
		return errors.New("CANNOT SET INGRESS VIP POOL ID")
	}
	return nil
}
