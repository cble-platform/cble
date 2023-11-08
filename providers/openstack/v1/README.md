# Openstack CBLE Provider (v1)

## Generating Provider Config

First, download the `<project_name>-rc.sh` file from your Openstack deployment. Then pass it into the `generate_config.sh` script:

```shell
$ ./generate_config.sh path/to/some-rc.sh
Generating config from path/to/some-rc.sh...
Please enter your OpenStack Password for project <project name> as user <username>:
Generated Openstack provider config config.yaml!
```

## Example Blueprint

```yaml
# vars

main_subnet: 10.10.0.0/24
router_ip: 10.10.0.254
host_ip: 10.10.0.1
```

```yaml
# blueprint

version: "1.0"

# Host 1
host1:
  resource: openstack.v1.host
  config:
    hostname: host1
    image: ubuntu22.04
    flavor: l2-micro
    disk_size: 10240
    networks:
      network1:
        dhcp: false
        ip: "{{ .host_ip }}"
# Host 2
host2:
  resource: openstack.v1.host
  config:
    hostname: host2
    image: ubuntu22.04
    flavor: l2-micro
    disk_size: 10240
    networks:
      network1:
        dhcp: true
# Network 1
network1:
  resource: openstack.v1.network
  config:
    subnet: "{{ .main_subnet }}"
    gateway: "{{ .router_ip }}"
    dhcp:
      - start: 10.10.0.10
        end: 10.10.0.100
# Router 1
router1:
  resource: openstack.v1.router
  config:
    external_network: MAIN NET
    networks:
      network1:
        dhcp: false
        ip: "{{ .router_ip }}"
```
