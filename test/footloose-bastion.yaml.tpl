cluster:
  name: mke
  privateKey: ./id_rsa_launchpad
machines:
- count: 1
  backend: docker
  spec:
    image: $LINUX_IMAGE
    name: manager%d
    privileged: true
    volumes:
    - type: bind
      source: /lib/modules
      destination: /lib/modules
    - type: volume
      destination: /var/lib/containerd
    - type: volume
      destination: /var/lib/docker
    - type: volume
      destination: /var/lib/kubelet
    networks:
    - footloose-cluster
    - containerPort: 22
      hostPort: 9122
- count: 1
  backend: docker
  spec:
    image: $LINUX_IMAGE
    name: worker%d
    privileged: true
    volumes:
    - type: bind
      source: /lib/modules
      destination: /lib/modules
    - type: volume
      destination: /var/lib/containerd
    - type: volume
      destination: /var/lib/docker
    - type: volume
      destination: /var/lib/kubelet
    networks:
    - footloose-cluster
    - containerPort: 22
      hostPort: 9122
- count: 1
  backend: docker
  spec:
    image: $LINUX_IMAGE
    name: bastion%d
    privileged: true
    volumes:
    - type: bind
      source: /lib/modules
      destination: /lib/modules
    - type: volume
      destination: /var/lib/containerd
    - type: volume
      destination: /var/lib/docker
    - type: volume
      destination: /var/lib/kubelet
    networks:
    - footloose-cluster
    portMappings:
    - containerPort: 22
      hostPort: 9022
