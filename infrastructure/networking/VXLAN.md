# VXLAN
- VLAN supports only up to 4096 VLANs because of the spare 12 bits to encode the VLAN ID in the Ethernet frame
- VXLAN creates virtual broadcast domains (L2) but nodes have to have L3 (networking) connectivity
- Process:
	- Every VXLAN nodes' outgoing Ethernet frames are captured, wrapped into UDP datagrams (encapsulated), and sent over an L3 network to the destination VXLAN
	- On arrival, Ethernet frames extraced from the UDP packets (decapsulated) and injected into the destination's network interface, this is called tunneling
	- as a result, VXLAN nodes create a virtual L2 segment, hence an L2 broadcast domain

Links:
[[Networking]]
