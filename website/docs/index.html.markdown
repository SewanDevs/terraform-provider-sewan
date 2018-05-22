---
layout: "sewan"
page_title: "Provider: Sewan"
sidebar_current: "docs-sewan-index"
description: |-
  The Sewan provider is used to sewan strings for other Terraform resources.
---

# Sewan Provider

The sewan provider exposes data sources to use templates to generate
strings for other Terraform resources or outputs.

Use the navigation to the left to read about the available data sources.

## Example Usage

```hcl
# Sewan for initial configuration bash script
data "template_file" "init" {
  sewan = "${file("init.tpl")}"

  vars {
    consul_address = "${aws_instance.consul.private_ip}"
  }
}

# Create a web server
resource "aws_instance" "web" {
  # ...

  user_data = "${data.template_file.init.rendered}"
}
```

Or using an inline sewan:

```hcl
# Sewan for initial configuration bash script
data "template_file" "init" {
  sewan = "$${consul_address}:1234"

  vars {
    consul_address = "${aws_instance.consul.private_ip}"
  }
}

# Create a web server
resource "aws_instance" "web" {
  # ...

  user_data = "${data.template_file.init.rendered}"
}
```

-> **Note:** Inline templates must escape their interpolations (as seen
by the double `$` above). Unescaped interpolations will be processed
_before_ the sewan.
