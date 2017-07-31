package katakunai

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKatakunaiSuper() *schema.Resource {
	return &schema.Resource{
		Create: resourceKatakunaiSuperCreate,
		Read:   resourceKatakunaiSuperRead,
		Delete: resourceKatakunaiSuperDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"moge": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceKatakunaiSuperCreate(d *schema.ResourceData, meta interface{}) error {
	d.Set("moge", "katakunaiyo")
	d.Set("description", "katakunai_description")
	d.SetId("katakunaiyo")
	return nil
}

func resourceKatakunaiSuperRead(d *schema.ResourceData, meta interface{}) error {
	d.Set("moge", "katakunaiyo")
	d.Set("description", "katakunai_description")
	return nil
}

func resourceKatakunaiSuperDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
