---
layout: "scaleway"
page_title: "Scaleway: scaleway_baremetal_server_beta"
description: |-
  Manages Scaleway Compute Baremetal servers.
---

# scaleway_baremetal_server_beta

Creates and manages Scaleway Compute Baremetal servers. For more information, see [the documentation](https://developers.scaleway.com/en/products/baremetal/api).

## Examples
    
### Basic

```hcl
data "scaleway_account_ssh_key" "main" {
  name = "main"
}

resource "scaleway_baremetal_server_beta" "base" {
  zone		    = "fr-par-2"
  offer       = "GP-BM1-M"
  os_id       = "d17d6872-0412-45d9-a198-af82c34d3c5c"
  ssh_key_ids = [data.scaleway_account_ssh_key.main]
}
```

## Arguments Reference

The following arguments are supported:

- `offer` - (Required) The offer name or UUID of the baremetal server.
Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-334154) to find the right offer.

~> **Important:** Updates to `offer_id` will recreate the server.

- `os_id` - (Required) The UUID of the base image used by the server.
Use [this endpoint](https://developers.scaleway.com/en/products/baremetal/api/#get-87598a) to find the right OS ID.

~> **Important:** Updates to `os_id` will reinstall the server.

- `ssh_key_ids` - (Required) List of SSH keys allowed to connect to the server.

- `name` - (Optional) The name of the server.

- `description` - (Optional) A description for the server.

~> **Important:** Updates to `ssh_key_ids` will reinstall the server.

- `tags` - (Optional) The tags associated with the server.

- `zone` - (Defaults to [provider](../index.html#zone) `zone`) The [zone](../guides/regions_and_zones.html#zones) in which the server should be created.

- `organization_id` - (Defaults to [provider](../index.html#organization_id) `organization_id`) The ID of the organization the server is associated with.


## Attributes Reference

In addition to all above arguments, the following attributes are exported:

- `id` - The ID of the server.
- `ips` - (List of) The IPs of the server.
  - `id` - The ID of the IP.
  - `address` - The address of the IP.
  - `reverse` - The reverse of the IP.
- `domain` - The domain of the server.

## Import

Baremetal servers can be imported using the `{zone}/{id}`, e.g.

```bash
$ terraform import scaleway_baremetal_server_beta.web fr-par-2/11111111-1111-1111-1111-111111111111
```
