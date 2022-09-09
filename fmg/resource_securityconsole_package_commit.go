// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Install policies to device from preview cache. Only to be used when a preview cache is previously generated by install/package command.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSecurityconsolePackageCommit() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityconsolePackageCommitUpdate,
		Read:   resourceSecurityconsolePackageCommitRead,
		Update: resourceSecurityconsolePackageCommitUpdate,
		Delete: resourceSecurityconsolePackageCommitDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"fmgadom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSecurityconsolePackageCommitUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	obj, err := getObjectSecurityconsolePackageCommit(d)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsolePackageCommit resource while getting object: %v", err)
	}

	_, err = c.UpdateSecurityconsolePackageCommit(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SecurityconsolePackageCommit resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SecurityconsolePackageCommit")

	return resourceSecurityconsolePackageCommitRead(d, m)
}

func resourceSecurityconsolePackageCommitDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceSecurityconsolePackageCommitRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenSecurityconsolePackageCommitAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsolePackageCommitScope(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenSecurityconsolePackageCommitScopeName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "SecurityconsolePackageCommit-Scope-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenSecurityconsolePackageCommitScopeVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "SecurityconsolePackageCommit-Scope-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSecurityconsolePackageCommitScopeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSecurityconsolePackageCommitScopeVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSecurityconsolePackageCommit(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fmgadom", flattenSecurityconsolePackageCommitAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "SecurityconsolePackageCommit-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("scope", flattenSecurityconsolePackageCommitScope(o["scope"], d, "scope")); err != nil {
			if vv, ok := fortiAPIPatch(o["scope"], "SecurityconsolePackageCommit-Scope"); ok {
				if err = d.Set("scope", vv); err != nil {
					return fmt.Errorf("Error reading scope: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("scope"); ok {
			if err = d.Set("scope", flattenSecurityconsolePackageCommitScope(o["scope"], d, "scope")); err != nil {
				if vv, ok := fortiAPIPatch(o["scope"], "SecurityconsolePackageCommit-Scope"); ok {
					if err = d.Set("scope", vv); err != nil {
						return fmt.Errorf("Error reading scope: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading scope: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSecurityconsolePackageCommitFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSecurityconsolePackageCommitAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsolePackageCommitScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandSecurityconsolePackageCommitScopeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandSecurityconsolePackageCommitScopeVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSecurityconsolePackageCommitScopeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSecurityconsolePackageCommitScopeVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSecurityconsolePackageCommit(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("adom") {
		t, err := expandSecurityconsolePackageCommitAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("scope"); ok || d.HasChange("scope") {
		t, err := expandSecurityconsolePackageCommitScope(d, v, "scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}

	return &obj, nil
}
