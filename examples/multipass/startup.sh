#!/bin/bash
echo "Spinning up machines"

# Get the host RSA key
mykey=`cat ~/.ssh/id_rsa.pub`
cp cloud-init.template.yaml cloud-init.yaml
echo "  - " $mykey >> cloud-init.yaml

# Startup the VMs if needed
MANAGER_EXISTS=`multipass list | grep manager | tr -s ' ' | cut -d ' ' -f 1`
if [ -z "${MANAGER_EXISTS}" ]; then
  multipass launch 18.04 -n manager --memory 4G --disk 10G --cloud-init cloud-init.yaml
else
  echo "Manager VM already exists. Not recreating it"
fi

WORKER_EXISTS=`multipass list | grep worker | tr -s ' ' | cut -d ' ' -f 1`
if [ -z "${WORKER_EXISTS}" ]; then
  multipass launch 18.04 -n worker  --memory 4G --disk 10G --cloud-init cloud-init.yaml
else
  echo "Worker VM already exists. Not recreating it"
fi

# # Get the IP address of the VMs
cp launchpad.template.yaml launchpad.yaml

MANAGER_IP=`multipass info manager --format json | jq .info.manager.ipv4[0]`
sed -i "s/{{MANAGER_IP}}/$MANAGER_IP/g" launchpad.yaml

WORKER_IP=`multipass info worker --format json | jq .info.worker.ipv4[0]`
sed -i "s/{{WORKER_IP}}/$WORKER_IP/g" launchpad.yaml

# Deploy using launchpad
../../bin/launchpad apply -c ./launchpad.yml
