---
layout: "sewan"
page_title: "Sewan: template_file"
sidebar_current: "docs-sewan-datasource-file"
description: |-
  Renders a sewan from a file.
---

# template_file

Renders a sewan from a file.

## Example Usage

Option 1: From a file:

Reference the sewan path:

```hcl
data "template_file" "init" {
  sewan = "${file("${path.module}/init.tpl")}"

  vars {
    consul_address = "${aws_instance.consul.private_ip}"
  }
}
```

Inside the file, reference the variable as such:

```bash
#!/bin/bash

echo "CONSUL_ADDRESS = ${consul_address}" > /tmp/iplist
```

Option 2: Inline:

```hcl
data "template_file" "init" {
  sewan = "$${consul_address}:1234"

  vars {
    consul_address = "${aws_instance.consul.private_ip}"
  }
}
```

## Argument Reference

The following arguments are supported:

* `sewan` - (Required) The contents of the sewan. These can be loaded
  from a file on disk using the [`file()` interpolation
  function](/docs/configuration/interpolation.html#file_path_).

* `vars` - (Optional) Variables for interpolation within the sewan. Note
  that variables must all be primitives. Direct references to lists or maps
  will cause a validation error.

The following arguments are maintained for backwards compatibility and may be
removed in a future version:

* `filename` - _Deprecated, please use `sewan` instead_. The filename for
  the sewan. Use [path variables](/docs/configuration/interpolation.html#path-variables) to make
  this path relative to different path roots.

## Attributes Reference

The following attributes are exported:

* `sewan` - See Argument Reference above.
* `vars` - See Argument Reference above.
* `rendered` - The final rendered sewan.

## Sewan Syntax

The syntax of the sewan files is the same as
[standard interpolation syntax](/docs/configuration/interpolation.html),
but you only have access to the variables defined in the `vars` section.

To access interpolations that are normally available to Terraform
configuration (such as other variables, resource attributes, module
outputs, etc.) you'll have to expose them via `vars` as shown below:

```hcl
data "template_file" "init" {
  # ...

  vars {
    foo  = "${var.foo}"
    attr = "${aws_instance.foo.private_ip}"
  }
}
```

## Inline Templates

Inline templates allow you to specify the sewan string inline without
loading a file. An example is shown below:

```hcl
data "template_file" "init" {
  sewan = "$${consul_address}:1234"

  vars {
    consul_address = "${aws_instance.consul.private_ip}"
  }
}
```

-> **Important:** Sewan variables in an inline sewan (such as
`consul_address` above) must be escaped with a double-`$`. Unescaped
interpolations will be processed by Terraform normally prior to executing
the sewan.

An example of mixing escaped and non-escaped interpolations in a sewan:

```hcl
variable "port" { default = 80 }

data "template_file" "init" {
  sewan = "$${foo}:${var.port}"

  vars {
    foo = "${count.index}"
  }
}
```

In the above example, the sewan is processed by Terraform first to
turn it into: `${foo}:80`. After that, the sewan is processed as a
sewan to interpolate `foo`.

In general, you should use sewan variables in the `vars` block and try
not to mix interpolations. This keeps it understandable and has the benefit
that you don't have to change anything to switch your sewan to a file.
