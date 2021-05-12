// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Enable/disable logging of FortiGuard antivirus and IPS update packages received by FortiManager's built-in FortiGuard.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdateAvIpsAdvancedLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateAvIpsAdvancedLogUpdate,
		Read:   resourceFmupdateAvIpsAdvancedLogRead,
		Update: resourceFmupdateAvIpsAdvancedLogUpdate,
		Delete: resourceFmupdateAvIpsAdvancedLogDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"log_fortigate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateAvIpsAdvancedLogUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateAvIpsAdvancedLog(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateAvIpsAdvancedLog resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateAvIpsAdvancedLog(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateAvIpsAdvancedLog resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateAvIpsAdvancedLog")

	return resourceFmupdateAvIpsAdvancedLogRead(d, m)
}

func resourceFmupdateAvIpsAdvancedLogDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateAvIpsAdvancedLog(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateAvIpsAdvancedLog resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateAvIpsAdvancedLogRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateAvIpsAdvancedLog(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateAvIpsAdvancedLog resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateAvIpsAdvancedLog(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateAvIpsAdvancedLog resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateAvIpsAdvancedLogLogFortigate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenFmupdateAvIpsAdvancedLogLogServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectFmupdateAvIpsAdvancedLog(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("log_fortigate", flattenFmupdateAvIpsAdvancedLogLogFortigate(o["log-fortigate"], d, "log_fortigate")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-fortigate"], "FmupdateAvIpsAdvancedLog-LogFortigate"); ok {
			if err = d.Set("log_fortigate", vv); err != nil {
				return fmt.Errorf("Error reading log_fortigate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_fortigate: %v", err)
		}
	}

	if err = d.Set("log_server", flattenFmupdateAvIpsAdvancedLogLogServer(o["log-server"], d, "log_server")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-server"], "FmupdateAvIpsAdvancedLog-LogServer"); ok {
			if err = d.Set("log_server", vv); err != nil {
				return fmt.Errorf("Error reading log_server: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_server: %v", err)
		}
	}

	return nil
}

func flattenFmupdateAvIpsAdvancedLogFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateAvIpsAdvancedLogLogFortigate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateAvIpsAdvancedLogLogServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateAvIpsAdvancedLog(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("log_fortigate"); ok {
		t, err := expandFmupdateAvIpsAdvancedLogLogFortigate(d, v, "log_fortigate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-fortigate"] = t
		}
	}

	if v, ok := d.GetOk("log_server"); ok {
		t, err := expandFmupdateAvIpsAdvancedLogLogServer(d, v, "log_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-server"] = t
		}
	}

	return &obj, nil
}
