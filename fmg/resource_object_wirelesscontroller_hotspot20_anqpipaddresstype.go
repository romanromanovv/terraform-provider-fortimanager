// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IP address type availability.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerHotspot20AnqpIpAddressType() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerHotspot20AnqpIpAddressTypeCreate,
		Read:   resourceObjectWirelessControllerHotspot20AnqpIpAddressTypeRead,
		Update: resourceObjectWirelessControllerHotspot20AnqpIpAddressTypeUpdate,
		Delete: resourceObjectWirelessControllerHotspot20AnqpIpAddressTypeDelete,

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
			"ipv4_address_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_address_type": &schema.Schema{
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
		},
	}
}

func resourceObjectWirelessControllerHotspot20AnqpIpAddressTypeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpIpAddressType(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpIpAddressType resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerHotspot20AnqpIpAddressType(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerHotspot20AnqpIpAddressType resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20AnqpIpAddressTypeRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpIpAddressTypeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerHotspot20AnqpIpAddressType(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpIpAddressType resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerHotspot20AnqpIpAddressType(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerHotspot20AnqpIpAddressType resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerHotspot20AnqpIpAddressTypeRead(d, m)
}

func resourceObjectWirelessControllerHotspot20AnqpIpAddressTypeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerHotspot20AnqpIpAddressType(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerHotspot20AnqpIpAddressType resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerHotspot20AnqpIpAddressTypeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerHotspot20AnqpIpAddressType(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpIpAddressType resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerHotspot20AnqpIpAddressType(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerHotspot20AnqpIpAddressType resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpIpAddressTypeIpv4AddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "not-available",
			2: "not-known",
			3: "public",
			4: "port-restricted",
			5: "single-NATed-private",
			6: "double-NATed-private",
			7: "port-restricted-and-single-NATed",
			8: "port-restricted-and-double-NATed",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpIpAddressTypeIpv6AddressType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "not-available",
			1: "available",
			2: "not-known",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerHotspot20AnqpIpAddressTypeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerHotspot20AnqpIpAddressType(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ipv4_address_type", flattenObjectWirelessControllerHotspot20AnqpIpAddressTypeIpv4AddressType(o["ipv4-address-type"], d, "ipv4_address_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv4-address-type"], "ObjectWirelessControllerHotspot20AnqpIpAddressType-Ipv4AddressType"); ok {
			if err = d.Set("ipv4_address_type", vv); err != nil {
				return fmt.Errorf("Error reading ipv4_address_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv4_address_type: %v", err)
		}
	}

	if err = d.Set("ipv6_address_type", flattenObjectWirelessControllerHotspot20AnqpIpAddressTypeIpv6AddressType(o["ipv6-address-type"], d, "ipv6_address_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ipv6-address-type"], "ObjectWirelessControllerHotspot20AnqpIpAddressType-Ipv6AddressType"); ok {
			if err = d.Set("ipv6_address_type", vv); err != nil {
				return fmt.Errorf("Error reading ipv6_address_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ipv6_address_type: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerHotspot20AnqpIpAddressTypeName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerHotspot20AnqpIpAddressType-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerHotspot20AnqpIpAddressTypeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerHotspot20AnqpIpAddressTypeIpv4AddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpIpAddressTypeIpv6AddressType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerHotspot20AnqpIpAddressTypeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerHotspot20AnqpIpAddressType(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ipv4_address_type"); ok {
		t, err := expandObjectWirelessControllerHotspot20AnqpIpAddressTypeIpv4AddressType(d, v, "ipv4_address_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-address-type"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_address_type"); ok {
		t, err := expandObjectWirelessControllerHotspot20AnqpIpAddressTypeIpv6AddressType(d, v, "ipv6_address_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-address-type"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectWirelessControllerHotspot20AnqpIpAddressTypeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}