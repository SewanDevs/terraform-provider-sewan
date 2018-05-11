package sewan

import (
	"bytes"
	"github.com/hashicorp/terraform/helper/schema"
	"net/http"
	"io/ioutil"
)

func resourceVM() *schema.Resource {
	return &schema.Resource{
		Create: resourceVMCreate,
		Read:   resourceVMRead,
		Update: resourceVMUpdate,
		Delete: resourceVMDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdc": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ram": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"disk_image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"boot": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			//"nic": &schema.Schema{
			//	Type: schema.TypeString,
			//	Required: true,
			//},
			"template": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdc_resource_disk": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"backup": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

type VM struct {
	Name       string
	Vdc        string
	RAM        string
	CPU        string
	Disk_image string
	Boot       string
	//Nic					string
	Template          string
	Vdc_resource_disk string
	Backup            string
}

func resourceVMCreate(d *schema.ResourceData, m interface{}) error {

	VMInstance := VM{
		Name:       d.Get("name").(string),
		Vdc:        d.Get("vdc").(string),
		CPU:        d.Get("cpu").(string),
		RAM:        d.Get("ram").(string),
		Disk_image: d.Get("disk_image").(string),
		Boot:       d.Get("boot").(string),
		//Nic:				d.Get("Nic").(string),
		Template:          d.Get("template").(string),
		Vdc_resource_disk: d.Get("vdc_resource_disk").(string),
		Backup:            d.Get("backup").(string),
	}

	var requestBody bytes.Buffer
	var responseBody string
	client := &http.Client{}
	logger := loggerCreate("resourceVMCreate_" + VMInstance.Vdc + "_" + VMInstance.Name + ".log")

	// NB : The following 2 vars will be deleted when the provider config will be handled
	var dest string
	var argToken string

	requestBody.WriteString("{\"name\":\"" + VMInstance.Name + "\",")
	requestBody.WriteString("\"vdc\":\"" + VMInstance.Vdc + "\",")
	requestBody.WriteString("\"ram\":" + VMInstance.RAM + ",")
	requestBody.WriteString("\"cpu\":" + VMInstance.RAM + ",")
	requestBody.WriteString("\"disk_image\":\"" + VMInstance.Disk_image + "\",")
	requestBody.WriteString("\"boot\":\"" + VMInstance.Boot + "\",")
	//requestBody.WriteString("\"nics\":" + VMInstance.Nic + ",")
	requestBody.WriteString("\"template\":\"" + VMInstance.Template + "\",")
	requestBody.WriteString("\"vdc_resource_disk\":\"" + VMInstance.Vdc_resource_disk + "\",")
	requestBody.WriteString("\"backup\":\"" + VMInstance.Backup + "\"}")

	dest = "https://next.cloud-datacenter.fr/api/clouddc/vm/"
	argToken = "53278f61bad8d42ecb75cff430427747794cc4dd"

	req, _ := http.NewRequest("POST", dest, &requestBody)

	req.Header.Add("Host", "next.cloud-datacenter.fr")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("authorization", "Token "+argToken)
	req.Header.Add("Origin", "https://next.cloud-datacenter.fr")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Referer", "https://next.cloud-datacenter.fr/fr/datacenter/enterprise/sewan-rd-cloud-beta/vdc/view/sewan-rd-cloud-beta-dc1-terraf/vm/create")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")

	logger.Println("Creation of ", VMInstance.Name, "request Header = ", req.Header)
	logger.Println("Creation of ", VMInstance.Name, "request body = ", req.Body)

	resp, create_err := client.Do(req)
	defer resp.Body.Close()

	bodyBytes, create_resp_body_read_err := ioutil.ReadAll(resp.Body)
	responseBody = string(bodyBytes)

	if create_err != nil {
		logger.Println("Creation of ", VMInstance.Name, " response reception error : ", create_err)
	} else {
		defer d.SetId(VMInstance.Name)
	}

	if create_resp_body_read_err != nil {
		logger.Println("Creation of ", VMInstance.Name, " response body read error ", create_resp_body_read_err)
	}

	logger.Println("Creation of ", VMInstance.Name, " response status = ", resp.Status)
	logger.Println("Creation of ", VMInstance.Name, " response body = ", responseBody)

	return nil
}

func resourceVMRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceVMUpdate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceVMDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
