package kubernetes

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceKubernetesStorageClass() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKubernetesStorageClassRead,
		Schema: map[string]*schema.Schema{
			"metadata": metadataSchema("storage class", false),
			"parameters": {
				Type:        schema.TypeMap,
				Description: "The parameters for the provisioner that should create volumes of this storage class",
				Computed:    true,
			},
			"storage_provisioner": {
				Type:        schema.TypeString,
				Description: "Indicates the type of the provisioner",
				Computed:    true,
			},
			"reclaim_policy": {
				Type:        schema.TypeString,
				Description: "Indicates the type of the reclaim policy",
				Computed:    true,
			},
			"allow_volume_expansion": {
				Type:        schema.TypeBool,
				Description: "Indicates whether the storage class allow volume expand",
				Computed:    true,
			},
			"mount_options": {
				Type:        schema.TypeList,
				Description: "Indicates a map of mount options needed for some storage classes",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceKubernetesStorageClassRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("metadata.0.name").(string)
	d.SetId(name)
	return resourceKubernetesStorageClassRead(d, meta)
}
