# helloworld infrastructure and test protocols

## Required Tools
* [Docker](https://www.docker.com/)
* [sewan terraform provider image](https://hub.docker.com/r/sewan/terraform-provider-sewan/)
* A valid token of next.cloud-datacenter.fr session (with chrome or chromium : ctrl-maj-i, then get a token from an HTTP request)

## Test process for 

* clone AirDrumRegressionTests repository (choose a $WORKDIR/)
```
git clone git@gitlab.priv.sewan.fr:rd/AirDrumRegressionTests.git $WORKDIR/
```

* add the session token in variables.tf file with your favorite editor

* get the docker image and create a container from it
```
docker pull sewan/terraform-provider-sewan
docker run -it --name AirDrumRegressionTests sewan/terraform-provider-sewan bash
root@<container generated id>:~
```

* copy test file top the container, from an other terminal :
```
docker cp $WORKDIR/AirDrumRegressionTests/variables.tf <container generated id>:root/
docker cp $WORKDIR/AirDrumRegressionTests/main.tf <container generated id>:root/
```

* rm example file, terraform setup, validate sewan provider plugin is up and running :

```
root@<container generated id>:~ rm example-main.tf
root@<container generated id>:~ terraform init

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.
If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
root@<container generated id>:~ terraform providers
.
└── provider.sewan
```

* resources creation from main.tf infrastructure configuration file

```
root@<container generated id>:~ terraform apply
```

cmd return available under APPENDIX 1

* update configuration files (main.tf + centos7-rd-DC1_Template_override.tf.json) files, then update infrastructure

```
docker cp $WORKDIR/AirDrumRegressionTests/centos7-rd-DC1_Template_override.tf.json <container generated id>:root/centos7-rd-DC1_Template_override.tf.json
docker cp $WORKDIR/AirDrumRegressionTests/main-update.tf <container generated id>:root/main.tf
```
```
root@<container generated id>:~ terraform apply
```

cmd return available under APPENDIX 2

* delete infrastructure

```
root@<container generated id>:~ terraform destroy
```

cmd return available under APPENDIX 3

Notice the occurence of [terraform limitation](https://redmine.priv.sewan.fr/projects/telecomservice/wiki/Terraform-sewan-plugin_and_sewan-sdk-go#VI-Terraform-Limitations) that require a second run of "terraform destroy" command to successfully delete the vdc.

* List of Behavior to validate (need manual operations on next.clouddc + modification on configuration files) :

To process theses manual tests, first handle [terraform getting start doc](https://www.terraform.io/intro/getting-started/build.html)

  * Impossibility to delete a started resource
  * Impossibility to set the OS of a vm created from template
  * Impossibility to reduce the size of a disk
  * Get an error message when trying to create a vm from an non-existing template
  * etc.

## APPENDIX
### APPENDIX 1 : terraform apply

```

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  + sewan_clouddc_vdc.terraform-vdc-template
      id:                       <computed>
      datacenter:               "dc1"
      enterprise:               "sewan-rd-cloud-beta"
      name:                     "terraform-vdc-template"
      slug:                     <computed>
      vdc_resources.#:          "4"
      vdc_resources.0.resource: "ram"
      vdc_resources.0.slug:     <computed>
      vdc_resources.0.total:    "10"
      vdc_resources.0.used:     <computed>
      vdc_resources.1.resource: "cpu"
      vdc_resources.1.slug:     <computed>
      vdc_resources.1.total:    "10"
      vdc_resources.1.used:     <computed>
      vdc_resources.2.resource: "storage_enterprise"
      vdc_resources.2.slug:     <computed>
      vdc_resources.2.total:    "80"
      vdc_resources.2.used:     <computed>
      vdc_resources.3.resource: "storage_high_performance"
      vdc_resources.3.slug:     <computed>
      vdc_resources.3.total:    "10"
      vdc_resources.3.used:     <computed>

  + sewan_clouddc_vm.server[0]
      id:                       <computed>
      backup:                   "backup-no-backup"
      backup_size:              <computed>
      boot:                     "on disk"
      cpu:                      "2"
      disks.#:                  "2"
      disks.0.name:             "add disk test"
      disks.0.size:             "1"
      disks.0.slug:             <computed>
      disks.0.storage_class:    "storage_enterprise"
      disks.0.v_disk:           <computed>
      disks.1.name:             "add disk test2"
      disks.1.size:             "1"
      disks.1.slug:             <computed>
      disks.1.storage_class:    "storage_enterprise"
      disks.1.v_disk:           <computed>
      dynamic_field:            <computed>
      enterprise:               "sewan-rd-cloud-beta"
      name:                     "server-0"
      nics.#:                   "1"
      nics.0.connected:         "false"
      nics.0.mac_address:       <computed>
      nics.0.vlan:              "internal-2412"
      os:                       "Debian"
      outsourcing:              <computed>
      platform_name:            <computed>
      ram:                      "1"
      slug:                     <computed>
      state:                    <computed>
      storage_class:            "storage_enterprise"
      token:                    <computed>
      vdc:                      "${sewan_clouddc_vdc.terraform-vdc-template.slug}"

  + sewan_clouddc_vm.server[1]
      id:                       <computed>
      backup:                   "backup-no-backup"
      backup_size:              <computed>
      boot:                     "on disk"
      cpu:                      "2"
      disks.#:                  "2"
      disks.0.name:             "add disk test"
      disks.0.size:             "1"
      disks.0.slug:             <computed>
      disks.0.storage_class:    "storage_enterprise"
      disks.0.v_disk:           <computed>
      disks.1.name:             "add disk test2"
      disks.1.size:             "1"
      disks.1.slug:             <computed>
      disks.1.storage_class:    "storage_enterprise"
      disks.1.v_disk:           <computed>
      dynamic_field:            <computed>
      enterprise:               "sewan-rd-cloud-beta"
      name:                     "server-1"
      nics.#:                   "1"
      nics.0.connected:         "false"
      nics.0.mac_address:       <computed>
      nics.0.vlan:              "internal-2412"
      os:                       "Debian"
      outsourcing:              <computed>
      platform_name:            <computed>
      ram:                      "1"
      slug:                     <computed>
      state:                    <computed>
      storage_class:            "storage_enterprise"
      token:                    <computed>
      vdc:                      "${sewan_clouddc_vdc.terraform-vdc-template.slug}"

  + sewan_clouddc_vm.template-server[0]
      id:                       <computed>
      backup:                   "backup-no-backup"
      backup_size:              <computed>
      boot:                     "on disk"
      cpu:                      "2"
      dynamic_field:            <computed>
      enterprise:               "sewan-rd-cloud-beta"
      instance_number:          "1"
      name:                     "template-server"
      nics.#:                   "1"
      nics.0.connected:         "false"
      nics.0.mac_address:       <computed>
      nics.0.vlan:              "internal-2412"
      outsourcing:              <computed>
      platform_name:            <computed>
      ram:                      "1"
      slug:                     <computed>
      state:                    <computed>
      storage_class:            "storage_enterprise"
      template:                 "centos7-rd-DC1"
      token:                    <computed>
      vdc:                      "${sewan_clouddc_vdc.terraform-vdc-template.slug}"

  + sewan_clouddc_vm.template-server[1]
      id:                       <computed>
      backup:                   "backup-no-backup"
      backup_size:              <computed>
      boot:                     "on disk"
      cpu:                      "2"
      dynamic_field:            <computed>
      enterprise:               "sewan-rd-cloud-beta"
      instance_number:          "2"
      name:                     "template-server"
      nics.#:                   "1"
      nics.0.connected:         "false"
      nics.0.mac_address:       <computed>
      nics.0.vlan:              "internal-2412"
      outsourcing:              <computed>
      platform_name:            <computed>
      ram:                      "1"
      slug:                     <computed>
      state:                    <computed>
      storage_class:            "storage_enterprise"
      template:                 "centos7-rd-DC1"
      token:                    <computed>
      vdc:                      "${sewan_clouddc_vdc.terraform-vdc-template.slug}"


Plan: 5 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

sewan_clouddc_vdc.terraform-vdc-template: Creating...
  datacenter:               "" => "dc1"
  enterprise:               "" => "sewan-rd-cloud-beta"
  name:                     "" => "terraform-vdc-template"
  slug:                     "" => "<computed>"
  vdc_resources.#:          "0" => "4"
  vdc_resources.0.resource: "" => "ram"
  vdc_resources.0.slug:     "" => "<computed>"
  vdc_resources.0.total:    "" => "10"
  vdc_resources.0.used:     "" => "<computed>"
  vdc_resources.1.resource: "" => "cpu"
  vdc_resources.1.slug:     "" => "<computed>"
  vdc_resources.1.total:    "" => "10"
  vdc_resources.1.used:     "" => "<computed>"
  vdc_resources.2.resource: "" => "storage_enterprise"
  vdc_resources.2.slug:     "" => "<computed>"
  vdc_resources.2.total:    "" => "80"
  vdc_resources.2.used:     "" => "<computed>"
  vdc_resources.3.resource: "" => "storage_high_performance"
  vdc_resources.3.slug:     "" => "<computed>"
  vdc_resources.3.total:    "" => "10"
  vdc_resources.3.used:     "" => "<computed>"
sewan_clouddc_vdc.terraform-vdc-template: Creation complete after 0s (ID: 428)
sewan_clouddc_vm.template-server[0]: Creating...
  backup:             "" => "backup-no-backup"
  backup_size:        "" => "<computed>"
  boot:               "" => "on disk"
  cpu:                "" => "2"
  dynamic_field:      "" => "<computed>"
  enterprise:         "" => "sewan-rd-cloud-beta"
  instance_number:    "" => "1"
  name:               "" => "template-server"
  nics.#:             "0" => "1"
  nics.0.connected:   "" => "false"
  nics.0.mac_address: "" => "<computed>"
  nics.0.vlan:        "" => "internal-2412"
  outsourcing:        "" => "<computed>"
  platform_name:      "" => "<computed>"
  ram:                "" => "1"
  slug:               "" => "<computed>"
  state:              "" => "<computed>"
  storage_class:      "" => "storage_enterprise"
  template:           "" => "centos7-rd-DC1"
  token:              "" => "<computed>"
  vdc:                "" => "terraform-vdc-template"
sewan_clouddc_vm.template-server[1]: Creating...
  backup:             "" => "backup-no-backup"
  backup_size:        "" => "<computed>"
  boot:               "" => "on disk"
  cpu:                "" => "2"
  dynamic_field:      "" => "<computed>"
  enterprise:         "" => "sewan-rd-cloud-beta"
  instance_number:    "" => "2"
  name:               "" => "template-server"
  nics.#:             "0" => "1"
  nics.0.connected:   "" => "false"
  nics.0.mac_address: "" => "<computed>"
  nics.0.vlan:        "" => "internal-2412"
  outsourcing:        "" => "<computed>"
  platform_name:      "" => "<computed>"
  ram:                "" => "1"
  slug:               "" => "<computed>"
  state:              "" => "<computed>"
  storage_class:      "" => "storage_enterprise"
  template:           "" => "centos7-rd-DC1"
  token:              "" => "<computed>"
  vdc:                "" => "terraform-vdc-template"
sewan_clouddc_vm.server[1]: Creating...
  backup:                "" => "backup-no-backup"
  backup_size:           "" => "<computed>"
  boot:                  "" => "on disk"
  cpu:                   "" => "2"
  disks.#:               "0" => "2"
  disks.0.name:          "" => "add disk test"
  disks.0.size:          "" => "1"
  disks.0.slug:          "" => "<computed>"
  disks.0.storage_class: "" => "storage_enterprise"
  disks.0.v_disk:        "" => "<computed>"
  disks.1.name:          "" => "add disk test2"
  disks.1.size:          "" => "1"
  disks.1.slug:          "" => "<computed>"
  disks.1.storage_class: "" => "storage_enterprise"
  disks.1.v_disk:        "" => "<computed>"
  dynamic_field:         "" => "<computed>"
  enterprise:            "" => "sewan-rd-cloud-beta"
  name:                  "" => "server-1"
  nics.#:                "0" => "1"
  nics.0.connected:      "" => "false"
  nics.0.mac_address:    "" => "<computed>"
  nics.0.vlan:           "" => "internal-2412"
  os:                    "" => "Debian"
  outsourcing:           "" => "<computed>"
  platform_name:         "" => "<computed>"
  ram:                   "" => "1"
  slug:                  "" => "<computed>"
  state:                 "" => "<computed>"
  storage_class:         "" => "storage_enterprise"
  token:                 "" => "<computed>"
  vdc:                   "" => "terraform-vdc-template"
sewan_clouddc_vm.server[0]: Creating...
  backup:                "" => "backup-no-backup"
  backup_size:           "" => "<computed>"
  boot:                  "" => "on disk"
  cpu:                   "" => "2"
  disks.#:               "0" => "2"
  disks.0.name:          "" => "add disk test"
  disks.0.size:          "" => "1"
  disks.0.slug:          "" => "<computed>"
  disks.0.storage_class: "" => "storage_enterprise"
  disks.0.v_disk:        "" => "<computed>"
  disks.1.name:          "" => "add disk test2"
  disks.1.size:          "" => "1"
  disks.1.slug:          "" => "<computed>"
  disks.1.storage_class: "" => "storage_enterprise"
  disks.1.v_disk:        "" => "<computed>"
  dynamic_field:         "" => "<computed>"
  enterprise:            "" => "sewan-rd-cloud-beta"
  name:                  "" => "server-0"
  nics.#:                "0" => "1"
  nics.0.connected:      "" => "false"
  nics.0.mac_address:    "" => "<computed>"
  nics.0.vlan:           "" => "internal-2412"
  os:                    "" => "Debian"
  outsourcing:           "" => "<computed>"
  platform_name:         "" => "<computed>"
  ram:                   "" => "1"
  slug:                  "" => "<computed>"
  state:                 "" => "<computed>"
  storage_class:         "" => "storage_enterprise"
  token:                 "" => "<computed>"
  vdc:                   "" => "terraform-vdc-template"
sewan_clouddc_vm.server[0]: Creation complete after 0s (ID: 2118)
sewan_clouddc_vm.server[1]: Creation complete after 0s (ID: 2119)
sewan_clouddc_vm.template-server[0]: Creation complete after 1s (ID: 2120)
sewan_clouddc_vm.template-server[1]: Creation complete after 1s (ID: 2121)

Apply complete! Resources: 5 added, 0 changed, 0 destroyed.
```

### APPENDIX 2 : terraform apply (update infrastructure)

```
sewan_clouddc_vdc.terraform-vdc-template: Refreshing state... (ID: 428)
sewan_clouddc_vm.server[0]: Refreshing state... (ID: 2118)
sewan_clouddc_vm.server[1]: Refreshing state... (ID: 2119)
sewan_clouddc_vm.template-server[1]: Refreshing state... (ID: 2121)
sewan_clouddc_vm.template-server[0]: Refreshing state... (ID: 2120)

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  ~ update in-place

Terraform will perform the following actions:

  ~ sewan_clouddc_vdc.terraform-vdc-template
      vdc_resources.0.total: "10" => "11"
      vdc_resources.1.total: "10" => "11"
      vdc_resources.2.total: "80" => "79"
      vdc_resources.3.total: "10" => "9"

  ~ sewan_clouddc_vm.server[0]
      cpu:                   "2" => "1"
      disks.0.size:          "1" => "2"
      disks.1.name:          "add disk test2" => "updated name"
      disks.1.size:          "1" => "2"
      nics.0.connected:      "false" => "true"

  ~ sewan_clouddc_vm.server[1]
      cpu:                   "2" => "1"
      disks.0.size:          "1" => "2"
      disks.1.name:          "add disk test2" => "updated name"
      disks.1.size:          "1" => "2"
      nics.0.connected:      "false" => "true"

  ~ sewan_clouddc_vm.template-server[0]
      disks.0.size:          "20" => "21"
      ram:                   "1" => "2"

  ~ sewan_clouddc_vm.template-server[1]
      disks.0.size:          "20" => "21"
      ram:                   "1" => "2"


Plan: 0 to add, 5 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

sewan_clouddc_vdc.terraform-vdc-template: Modifying... (ID: 428)
  vdc_resources.0.total: "10" => "11"
  vdc_resources.1.total: "10" => "11"
  vdc_resources.2.total: "80" => "79"
  vdc_resources.3.total: "10" => "9"
sewan_clouddc_vdc.terraform-vdc-template: Modifications complete after 1s (ID: 428)
sewan_clouddc_vm.server[0]: Modifying... (ID: 2118)
  cpu:              "2" => "1"
  disks.0.size:     "1" => "2"
  disks.1.name:     "add disk test2" => "updated name"
  disks.1.size:     "1" => "2"
  nics.0.connected: "false" => "true"
sewan_clouddc_vm.template-server[0]: Modifying... (ID: 2120)
  disks.0.size: "20" => "21"
  ram:          "1" => "2"
sewan_clouddc_vm.server[1]: Modifying... (ID: 2119)
  cpu:              "2" => "1"
  disks.0.size:     "1" => "2"
  disks.1.name:     "add disk test2" => "updated name"
  disks.1.size:     "1" => "2"
  nics.0.connected: "false" => "true"
sewan_clouddc_vm.template-server[1]: Modifying... (ID: 2121)
  disks.0.size: "20" => "21"
  ram:          "1" => "2"
sewan_clouddc_vm.server[0]: Modifications complete after 0s (ID: 2118)
sewan_clouddc_vm.template-server[0]: Modifications complete after 0s (ID: 2120)
sewan_clouddc_vm.server[1]: Modifications complete after 0s (ID: 2119)
sewan_clouddc_vm.template-server[1]: Modifications complete after 0s (ID: 2121)

Apply complete! Resources: 0 added, 5 changed, 0 destroyed.
```

### APPENDIX 3 : terraform destroy (delete infrastructure)

```
sewan_clouddc_vdc.terraform-vdc-template: Refreshing state... (ID: 428)
sewan_clouddc_vm.server[1]: Refreshing state... (ID: 2119)
sewan_clouddc_vm.server[0]: Refreshing state... (ID: 2118)
sewan_clouddc_vm.template-server[0]: Refreshing state... (ID: 2120)
sewan_clouddc_vm.template-server[1]: Refreshing state... (ID: 2121)

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  - destroy

Terraform will perform the following actions:

  - sewan_clouddc_vdc.terraform-vdc-template

  - sewan_clouddc_vm.server[0]

  - sewan_clouddc_vm.server[1]

  - sewan_clouddc_vm.template-server[0]

  - sewan_clouddc_vm.template-server[1]


Plan: 0 to add, 0 to change, 5 to destroy.

Do you really want to destroy?
  Terraform will destroy all your managed infrastructure, as shown above.
  There is no undo. Only 'yes' will be accepted to confirm.

  Enter a value: yes

sewan_clouddc_vm.server[1]: Destroying... (ID: 2119)
sewan_clouddc_vm.template-server[0]: Destroying... (ID: 2120)
sewan_clouddc_vm.template-server[1]: Destroying... (ID: 2121)
sewan_clouddc_vm.server[0]: Destroying... (ID: 2118)
sewan_clouddc_vm.server[1]: Destruction complete after 0s
sewan_clouddc_vm.server[0]: Destruction complete after 0s
sewan_clouddc_vm.template-server[1]: Destruction complete after 0s
sewan_clouddc_vm.template-server[0]: Destruction complete after 0s
sewan_clouddc_vdc.terraform-vdc-template: Destroying... (ID: 428)

Error: Error applying plan:

1 error(s) occurred:

* sewan_clouddc_vdc.terraform-vdc-template (destroy): 1 error(s) occurred:

* sewan_clouddc_vdc.terraform-vdc-template: Wrong response status code,
expected :204
got :406
Full response status : 406 Not Acceptable

Response body error :{"detail":"You still have VM in this VDC"}

Terraform does not automatically rollback in the face of errors.
Instead, your Terraform state file has been partially updated with
any resources that successfully completed. Please address the error
above and apply again to incrementally change your infrastructure.
```
