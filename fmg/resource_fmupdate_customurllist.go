// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure the URL database for rating and filtering.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdateCustomUrlList() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateCustomUrlListUpdate,
		Read:   resourceFmupdateCustomUrlListRead,
		Update: resourceFmupdateCustomUrlListUpdate,
		Delete: resourceFmupdateCustomUrlListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"db_selection": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateCustomUrlListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateCustomUrlList(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateCustomUrlList resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateCustomUrlList(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateCustomUrlList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateCustomUrlList")

	return resourceFmupdateCustomUrlListRead(d, m)
}

func resourceFmupdateCustomUrlListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateCustomUrlList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateCustomUrlList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateCustomUrlListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateCustomUrlList(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateCustomUrlList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateCustomUrlList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateCustomUrlList resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateCustomUrlListDbSelection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "both",
			2: "custom-url",
			4: "fortiguard-db",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func refreshObjectFmupdateCustomUrlList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("db_selection", flattenFmupdateCustomUrlListDbSelection(o["db_selection"], d, "db_selection")); err != nil {
		if vv, ok := fortiAPIPatch(o["db_selection"], "FmupdateCustomUrlList-DbSelection"); ok {
			if err = d.Set("db_selection", vv); err != nil {
				return fmt.Errorf("Error reading db_selection: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading db_selection: %v", err)
		}
	}

	return nil
}

func flattenFmupdateCustomUrlListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateCustomUrlListDbSelection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectFmupdateCustomUrlList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("db_selection"); ok {
		t, err := expandFmupdateCustomUrlListDbSelection(d, v, "db_selection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["db_selection"] = t
		}
	}

	return &obj, nil
}
