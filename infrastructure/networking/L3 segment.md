# L3 segment
- IP subnetwork
- [[L2 segment]] only requires MAC addresses

### How IP Packets are sent across L3?
- communications between L3 segments requires at least one **router**
- when a node SENDS an IP packet to a node that resides in another L3 segment (diff IP subnet), it sends it to its gateway router instead.
- since nodes can only talk directly with other nodes on the same L2 segment, once of the router's interfaces has to reside on the sender's L2 segment

### How IP Packets are sent within L3?
- IP packets are wrapped into Ethernet frames (assuming L2 uses Ethernet protocol)
- IP protocol data units (packets) are encapsulated in the Ethernet protocol data units (frames)

- Send an Ethernet frame with the IP packet inside to the L2 Segment's node that owns the destination IP. 
	- So the sender node needs to learn the destination's MAC address first. L3 (IP) to L2 (MAC) address translation mechanism is required. 
	- Neighbor Discovery Protocol (ARP for IPv4, NDP for IPv6) that relies on L2 broadcast capabilities.
- When IP to MAC translation is not known, sender broadcast L2 Frame asking who has the destination IP. Once the destination MAC is known, sender just wraps the IP packet into an L2 frame destined to the MAC address
- Normally, L3 and L2 has 1:1 mapping but it doesnt prevent you from creating L3 segments over L2 (connected machines in one network switch)
- you can also create an L3 over multiple L2 segments using [[Proxy ARP]]
- 

Questions:
- are there other protocol for [[L2 segment]]?
