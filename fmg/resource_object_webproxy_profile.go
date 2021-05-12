// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure web proxy profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWebProxyProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWebProxyProfileCreate,
		Read:   resourceObjectWebProxyProfileRead,
		Update: resourceObjectWebProxyProfileUpdate,
		Delete: resourceObjectWebProxyProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
					"inherit",
				}, false),
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"header_client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_front_end_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_via_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_via_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_x_authenticated_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_x_authenticated_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_x_forwarded_for": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"add_option": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"base64_encoding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dstaddr6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"log_header_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"strip_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectWebProxyProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWebProxyProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWebProxyProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWebProxyProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWebProxyProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebProxyProfileRead(d, m)
}

func resourceObjectWebProxyProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWebProxyProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebProxyProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWebProxyProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWebProxyProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWebProxyProfileRead(d, m)
}

func resourceObjectWebProxyProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWebProxyProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWebProxyProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWebProxyProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWebProxyProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebProxyProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWebProxyProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWebProxyProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWebProxyProfileHeaderClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "pass",
			1: "add",
			2: "remove",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWebProxyProfileHeaderFrontEndHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "pass",
			1: "add",
			2: "remove",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWebProxyProfileHeaderViaRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "pass",
			1: "add",
			2: "remove",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWebProxyProfileHeaderViaResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "pass",
			1: "add",
			2: "remove",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWebProxyProfileHeaderXAuthenticatedGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "pass",
			1: "add",
			2: "remove",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWebProxyProfileHeaderXAuthenticatedUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "pass",
			1: "add",
			2: "remove",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWebProxyProfileHeaderXForwardedFor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "pass",
			1: "add",
			2: "remove",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWebProxyProfileHeaders(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenObjectWebProxyProfileHeadersAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "ObjectWebProxyProfile-Headers-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_option"
		if _, ok := i["add-option"]; ok {
			v := flattenObjectWebProxyProfileHeadersAddOption(i["add-option"], d, pre_append)
			tmp["add_option"] = fortiAPISubPartPatch(v, "ObjectWebProxyProfile-Headers-AddOption")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := i["base64-encoding"]; ok {
			v := flattenObjectWebProxyProfileHeadersBase64Encoding(i["base64-encoding"], d, pre_append)
			tmp["base64_encoding"] = fortiAPISubPartPatch(v, "ObjectWebProxyProfile-Headers-Base64Encoding")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := i["content"]; ok {
			v := flattenObjectWebProxyProfileHeadersContent(i["content"], d, pre_append)
			tmp["content"] = fortiAPISubPartPatch(v, "ObjectWebProxyProfile-Headers-Content")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := i["dstaddr"]; ok {
			v := flattenObjectWebProxyProfileHeadersDstaddr(i["dstaddr"], d, pre_append)
			tmp["dstaddr"] = fortiAPISubPartPatch(v, "ObjectWebProxyProfile-Headers-Dstaddr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := i["dstaddr6"]; ok {
			v := flattenObjectWebProxyProfileHeadersDstaddr6(i["dstaddr6"], d, pre_append)
			tmp["dstaddr6"] = fortiAPISubPartPatch(v, "ObjectWebProxyProfile-Headers-Dstaddr6")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWebProxyProfileHeadersId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWebProxyProfile-Headers-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWebProxyProfileHeadersName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWebProxyProfile-Headers-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			v := flattenObjectWebProxyProfileHeadersProtocol(i["protocol"], d, pre_append)
			tmp["protocol"] = fortiAPISubPartPatch(v, "ObjectWebProxyProfile-Headers-Protocol")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWebProxyProfileHeadersAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "add-to-request",
			1: "add-to-response",
			2: "remove-from-request",
			3: "remove-from-response",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWebProxyProfileHeadersAddOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "append",
			2: "new-on-not-found",
			3: "new",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWebProxyProfileHeadersBase64Encoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectWebProxyProfileHeadersContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyProfileHeadersDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyProfileHeadersDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyProfileHeadersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyProfileHeadersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyProfileHeadersProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "https",
			2: "http",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenObjectWebProxyProfileLogHeaderChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenObjectWebProxyProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWebProxyProfileStripEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func refreshObjectObjectWebProxyProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("header_client_ip", flattenObjectWebProxyProfileHeaderClientIp(o["header-client-ip"], d, "header_client_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-client-ip"], "ObjectWebProxyProfile-HeaderClientIp"); ok {
			if err = d.Set("header_client_ip", vv); err != nil {
				return fmt.Errorf("Error reading header_client_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_client_ip: %v", err)
		}
	}

	if err = d.Set("header_front_end_https", flattenObjectWebProxyProfileHeaderFrontEndHttps(o["header-front-end-https"], d, "header_front_end_https")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-front-end-https"], "ObjectWebProxyProfile-HeaderFrontEndHttps"); ok {
			if err = d.Set("header_front_end_https", vv); err != nil {
				return fmt.Errorf("Error reading header_front_end_https: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_front_end_https: %v", err)
		}
	}

	if err = d.Set("header_via_request", flattenObjectWebProxyProfileHeaderViaRequest(o["header-via-request"], d, "header_via_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-via-request"], "ObjectWebProxyProfile-HeaderViaRequest"); ok {
			if err = d.Set("header_via_request", vv); err != nil {
				return fmt.Errorf("Error reading header_via_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_via_request: %v", err)
		}
	}

	if err = d.Set("header_via_response", flattenObjectWebProxyProfileHeaderViaResponse(o["header-via-response"], d, "header_via_response")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-via-response"], "ObjectWebProxyProfile-HeaderViaResponse"); ok {
			if err = d.Set("header_via_response", vv); err != nil {
				return fmt.Errorf("Error reading header_via_response: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_via_response: %v", err)
		}
	}

	if err = d.Set("header_x_authenticated_groups", flattenObjectWebProxyProfileHeaderXAuthenticatedGroups(o["header-x-authenticated-groups"], d, "header_x_authenticated_groups")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-x-authenticated-groups"], "ObjectWebProxyProfile-HeaderXAuthenticatedGroups"); ok {
			if err = d.Set("header_x_authenticated_groups", vv); err != nil {
				return fmt.Errorf("Error reading header_x_authenticated_groups: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_x_authenticated_groups: %v", err)
		}
	}

	if err = d.Set("header_x_authenticated_user", flattenObjectWebProxyProfileHeaderXAuthenticatedUser(o["header-x-authenticated-user"], d, "header_x_authenticated_user")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-x-authenticated-user"], "ObjectWebProxyProfile-HeaderXAuthenticatedUser"); ok {
			if err = d.Set("header_x_authenticated_user", vv); err != nil {
				return fmt.Errorf("Error reading header_x_authenticated_user: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_x_authenticated_user: %v", err)
		}
	}

	if err = d.Set("header_x_forwarded_for", flattenObjectWebProxyProfileHeaderXForwardedFor(o["header-x-forwarded-for"], d, "header_x_forwarded_for")); err != nil {
		if vv, ok := fortiAPIPatch(o["header-x-forwarded-for"], "ObjectWebProxyProfile-HeaderXForwardedFor"); ok {
			if err = d.Set("header_x_forwarded_for", vv); err != nil {
				return fmt.Errorf("Error reading header_x_forwarded_for: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading header_x_forwarded_for: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("headers", flattenObjectWebProxyProfileHeaders(o["headers"], d, "headers")); err != nil {
			if vv, ok := fortiAPIPatch(o["headers"], "ObjectWebProxyProfile-Headers"); ok {
				if err = d.Set("headers", vv); err != nil {
					return fmt.Errorf("Error reading headers: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading headers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("headers"); ok {
			if err = d.Set("headers", flattenObjectWebProxyProfileHeaders(o["headers"], d, "headers")); err != nil {
				if vv, ok := fortiAPIPatch(o["headers"], "ObjectWebProxyProfile-Headers"); ok {
					if err = d.Set("headers", vv); err != nil {
						return fmt.Errorf("Error reading headers: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading headers: %v", err)
				}
			}
		}
	}

	if err = d.Set("log_header_change", flattenObjectWebProxyProfileLogHeaderChange(o["log-header-change"], d, "log_header_change")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-header-change"], "ObjectWebProxyProfile-LogHeaderChange"); ok {
			if err = d.Set("log_header_change", vv); err != nil {
				return fmt.Errorf("Error reading log_header_change: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_header_change: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWebProxyProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWebProxyProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("strip_encoding", flattenObjectWebProxyProfileStripEncoding(o["strip-encoding"], d, "strip_encoding")); err != nil {
		if vv, ok := fortiAPIPatch(o["strip-encoding"], "ObjectWebProxyProfile-StripEncoding"); ok {
			if err = d.Set("strip_encoding", vv); err != nil {
				return fmt.Errorf("Error reading strip_encoding: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading strip_encoding: %v", err)
		}
	}

	return nil
}

func flattenObjectWebProxyProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWebProxyProfileHeaderClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeaderFrontEndHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeaderViaRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeaderViaResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeaderXAuthenticatedGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeaderXAuthenticatedUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeaderXForwardedFor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeaders(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandObjectWebProxyProfileHeadersAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_option"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["add-option"], _ = expandObjectWebProxyProfileHeadersAddOption(d, i["add_option"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["base64-encoding"], _ = expandObjectWebProxyProfileHeadersBase64Encoding(d, i["base64_encoding"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["content"], _ = expandObjectWebProxyProfileHeadersContent(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dstaddr"], _ = expandObjectWebProxyProfileHeadersDstaddr(d, i["dstaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dstaddr6"], _ = expandObjectWebProxyProfileHeadersDstaddr6(d, i["dstaddr6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectWebProxyProfileHeadersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectWebProxyProfileHeadersName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandObjectWebProxyProfileHeadersProtocol(d, i["protocol"], pre_append)
		} else {
			tmp["protocol"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWebProxyProfileHeadersAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeadersAddOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeadersBase64Encoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeadersContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeadersDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeadersDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeadersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeadersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileHeadersProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWebProxyProfileLogHeaderChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWebProxyProfileStripEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWebProxyProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("header_client_ip"); ok {
		t, err := expandObjectWebProxyProfileHeaderClientIp(d, v, "header_client_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-client-ip"] = t
		}
	}

	if v, ok := d.GetOk("header_front_end_https"); ok {
		t, err := expandObjectWebProxyProfileHeaderFrontEndHttps(d, v, "header_front_end_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-front-end-https"] = t
		}
	}

	if v, ok := d.GetOk("header_via_request"); ok {
		t, err := expandObjectWebProxyProfileHeaderViaRequest(d, v, "header_via_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-via-request"] = t
		}
	}

	if v, ok := d.GetOk("header_via_response"); ok {
		t, err := expandObjectWebProxyProfileHeaderViaResponse(d, v, "header_via_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-via-response"] = t
		}
	}

	if v, ok := d.GetOk("header_x_authenticated_groups"); ok {
		t, err := expandObjectWebProxyProfileHeaderXAuthenticatedGroups(d, v, "header_x_authenticated_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-authenticated-groups"] = t
		}
	}

	if v, ok := d.GetOk("header_x_authenticated_user"); ok {
		t, err := expandObjectWebProxyProfileHeaderXAuthenticatedUser(d, v, "header_x_authenticated_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-authenticated-user"] = t
		}
	}

	if v, ok := d.GetOk("header_x_forwarded_for"); ok {
		t, err := expandObjectWebProxyProfileHeaderXForwardedFor(d, v, "header_x_forwarded_for")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-forwarded-for"] = t
		}
	}

	if v, ok := d.GetOk("headers"); ok {
		t, err := expandObjectWebProxyProfileHeaders(d, v, "headers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["headers"] = t
		}
	}

	if v, ok := d.GetOk("log_header_change"); ok {
		t, err := expandObjectWebProxyProfileLogHeaderChange(d, v, "log_header_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-header-change"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectWebProxyProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("strip_encoding"); ok {
		t, err := expandObjectWebProxyProfileStripEncoding(d, v, "strip_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-encoding"] = t
		}
	}

	return &obj, nil
}