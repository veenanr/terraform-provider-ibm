// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	kp "github.com/IBM/keyprotect-go-client"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func suppressKMSInstanceIDDiff(k, old, new string, d *schema.ResourceData) bool {
	// TF currently uses GUID. So just check when instance crn is passed as input it has same GUID in it.
	return old == getInstanceIDFromCRN(new)
}

// Get Instance ID from CRN
func getInstanceIDFromCRN(crn string) string {
	crnSegments := strings.Split(crn, ":")
	if len(crnSegments) > 3 {
		return crnSegments[len(crnSegments)-3]
	}
	return crn
}

func ResourceIBMKmskey() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMKmsKeyCreate,
		Read:     resourceIBMKmsKeyRead,
		Update:   resourceIBMKmsKeyUpdate,
		Delete:   resourceIBMKmsKeyDelete,
		Exists:   resourceIBMKmsKeyExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				Description:      "Key protect or hpcs instance GUID or CRN",
				DiffSuppressFunc: suppressKMSInstanceIDDiff,
			},
			"key_ring_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "default",
				Description: "Key Ring for the Key",
			},
			"key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Key ID",
			},
			"key_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Key name",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "type of service hs-crypto or kms",
			},
			"endpoint_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"public", "private"}),
				Description:  "public or private",
				ForceNew:     true,
			},
			"standard_key": {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				ForceNew:    true,
				Description: "Standard key type",
			},
			"payload": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			"encrypted_nonce": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Only for imported root key",
			},
			"iv_value": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Only for imported root key",
			},
			"force_delete": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "set to true to force delete the key",
				ForceNew:    false,
				Default:     false,
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Crn of the key",
			},
			"expiration_date": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The date the key material expires. The date format follows RFC 3339. You can set an expiration date on any key on its creation. A key moves into the Deactivated state within one hour past its expiration date, if one is assigned. If you create a key without specifying an expiration date, the key does not expire",
				ForceNew:    true,
			},
			"instance_crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Key protect or hpcs instance CRN",
			},
			flex.ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			flex.ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			flex.ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			flex.ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
			flex.ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about the resource",
			},
		},
	}
}

func resourceIBMKmsKeyCreate(d *schema.ResourceData, meta interface{}) error {
	instanceID := getInstanceIDFromCRN(d.Get("instance_id").(string))
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return err
	}

	kpAPI.Config.KeyRing = d.Get("key_ring_id").(string)

	name := d.Get("key_name").(string)
	standardKey := d.Get("standard_key").(bool)

	var expiration *time.Time
	if es, ok := d.GetOk("expiration_date"); ok {
		expiration_string := es.(string)
		// parse string to required time format
		expiration_time, err := time.Parse(time.RFC3339, expiration_string)
		if err != nil {
			return fmt.Errorf("[ERROR] Invalid time format (the date format follows RFC 3339): %s", err)
		}
		expiration = &expiration_time
	} else {
		expiration = nil
	}

	var keyCRN string
	if standardKey {
		if v, ok := d.GetOk("payload"); ok {
			//import standard key
			payload := v.(string)
			stkey, err := kpAPI.CreateImportedStandardKey(context.Background(), name, expiration, payload)
			if err != nil {
				return fmt.Errorf("[ERROR] Error while creating standard key with payload: %s", err)
			}
			keyCRN = stkey.CRN
			d.SetId(keyCRN)

		} else {
			//create standard key
			stkey, err := kpAPI.CreateStandardKey(context.Background(), name, expiration)
			if err != nil {
				return fmt.Errorf("[ERROR] Error while creating standard key: %s", err)
			}
			keyCRN = stkey.CRN
			d.SetId(keyCRN)

		}
	} else {
		if v, ok := d.GetOk("payload"); ok {
			payload := v.(string)
			encryptedNonce := d.Get("encrypted_nonce").(string)
			iv := d.Get("iv_value").(string)
			stkey, err := kpAPI.CreateImportedRootKey(context.Background(), name, expiration, payload, encryptedNonce, iv)
			if err != nil {
				return fmt.Errorf("[ERROR] Error while creating Root key with payload: %s", err)
			}
			keyCRN = stkey.CRN
			d.SetId(keyCRN)

		} else {
			stkey, err := kpAPI.CreateRootKey(context.Background(), name, expiration)
			if err != nil {
				return fmt.Errorf("[ERROR] Error while creating Root key: %s", err)
			}
			keyCRN = stkey.CRN
			d.SetId(keyCRN)
		}
	}
	return resourceIBMKmsKeyUpdate(d, meta)
}

func resourceIBMKmsKeyRead(d *schema.ResourceData, meta interface{}) error {
	instanceCRN, instanceID, keyid := getInstanceAndKeyDataFromCRN(d.Id())

	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return err
	}
	// keyid := d.Id()
	key, err := kpAPI.GetKey(context.Background(), keyid)
	if err != nil {
		kpError := err.(*kp.Error)
		if kpError.StatusCode == 404 || kpError.StatusCode == 409 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Get Key failed with error while reading Key: %s", err)
	} else if key.State == 5 { //Refers to Deleted state of the Key
		d.SetId("")
		return nil
	}

	err = setKeyDetails(d, meta, instanceID, instanceCRN, key, kpAPI)
	if err != nil {
		return err
	}

	return nil

}

func resourceIBMKmsKeyUpdate(d *schema.ResourceData, meta interface{}) error {

	if d.HasChange("force_delete") {
		d.Set("force_delete", d.Get("force_delete").(bool))
	}
	return resourceIBMKmsKeyRead(d, meta)

}

func resourceIBMKmsKeyDelete(d *schema.ResourceData, meta interface{}) error {
	_, instanceID, keyid := getInstanceAndKeyDataFromCRN(d.Id())
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return err
	}

	force := d.Get("force_delete").(bool)
	f := kp.ForceOpt{
		Force: force,
	}

	_, err1 := kpAPI.DeleteKey(context.Background(), keyid, kp.ReturnRepresentation, f)
	if err1 != nil {
		return fmt.Errorf("[ERROR] Error while deleting: %s", err1)
	}
	d.SetId("")
	return nil

}

func resourceIBMKmsKeyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	_, instanceID, keyid := getInstanceAndKeyDataFromCRN(d.Id())

	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return false, err
	}

	_, err = kpAPI.GetKey(context.Background(), keyid)
	if err != nil {
		kpError := err.(*kp.Error)
		if kpError.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil

}

// Populate KP Client using info from schema
func populateKPClient(d *schema.ResourceData, meta interface{}, instanceID string) (kpAPI *kp.Client, instanceCRN *string, err error) {
	kpAPI, err = meta.(conns.ClientSession).KeyManagementAPI()
	if err != nil {
		return nil, nil, err
	}

	endpointType := d.Get("endpoint_type").(string)

	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return nil, nil, err
	}
	resourceInstanceGet := rc.GetResourceInstanceOptions{
		ID: &instanceID,
	}

	instanceData, resp, err := rsConClient.GetResourceInstance(&resourceInstanceGet)
	if err != nil || instanceData == nil {
		return nil, nil, fmt.Errorf("[ERROR] Error retrieving resource instance: %s with resp code: %s", err, resp)
	}
	extensions := instanceData.Extensions
	kpAPI.URL, err = KmsEndpointURL(kpAPI, endpointType, extensions)
	if err != nil {
		return nil, nil, err
	}

	kpAPI.Config.InstanceID = instanceID
	return kpAPI, instanceData.CRN, nil
}

// Set Key Details in the schema
func setKeyDetails(d *schema.ResourceData, meta interface{}, instanceID string, instanceCRN string, key *kp.Key, kpAPI *kp.Client) error {
	d.Set("instance_id", instanceID)
	d.Set("instance_crn", instanceCRN)
	d.Set("key_id", key.ID)
	d.Set("standard_key", key.Extractable)
	d.Set("payload", key.Payload)
	d.Set("encrypted_nonce", key.EncryptedNonce)
	d.Set("iv_value", key.IV)
	d.Set("key_name", key.Name)
	d.Set("crn", key.CRN)
	if strings.Contains((kpAPI.URL).String(), "private") || strings.Contains(kpAPI.Config.BaseURL, "private") {
		d.Set("endpoint_type", "private")
	} else {
		d.Set("endpoint_type", "public")
	}
	d.Set("type", strings.Split(d.Id(), ":")[4])
	if d.Get("force_delete") != nil {
		d.Set("force_delete", d.Get("force_delete").(bool))
	}
	d.Set("key_ring_id", key.KeyRingID)
	if key.Expiration != nil {
		expiration := key.Expiration
		d.Set("expiration_date", expiration.Format(time.RFC3339))
	} else {
		d.Set("expiration_date", "")
	}
	d.Set(flex.ResourceName, key.Name)
	d.Set(flex.ResourceCRN, key.CRN)
	state := key.State
	d.Set(flex.ResourceStatus, strconv.Itoa(state))
	rcontroller, err := flex.GetBaseController(meta)
	if err != nil {
		return err
	}
	id := key.ID
	crn1 := strings.TrimSuffix(key.CRN, ":key:"+id)

	d.Set(flex.ResourceControllerURL, rcontroller+"/services/kms/"+url.QueryEscape(crn1)+"%3A%3A")

	return nil
}

// Extract Instance and Key related info from crn
func getInstanceAndKeyDataFromCRN(crn string) (instanceCRN string, instanceID string, keyID string) {
	crnData := strings.Split(crn, ":")
	instanceCRN = fmt.Sprintf("%s::", strings.Split(crn, ":key:")[0])
	keyID = crnData[len(crnData)-1]
	instanceID = crnData[len(crnData)-3]
	return instanceCRN, instanceID, keyID
}

// Construct KMS URL
func KmsEndpointURL(kpAPI *kp.Client, endpointType string, extensions map[string]interface{}) (*url.URL, error) {

	exturl := extensions["endpoints"].(map[string]interface{})["public"]
	if endpointType == "private" || strings.Contains(kpAPI.Config.BaseURL, "private") {
		exturl = extensions["endpoints"].(map[string]interface{})["private"]
	}
	endpointURL := fmt.Sprintf("%s/api/v2/keys", exturl.(string))

	url1 := conns.EnvFallBack([]string{"IBMCLOUD_KP_API_ENDPOINT"}, endpointURL)
	if !strings.HasSuffix(url1, "/api/v2/keys") {
		url1 = url1 + "/api/v2/keys"
	}
	u, err := url.Parse(url1)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Error Parsing KMS EndpointURL")
	}
	return u, nil
}
