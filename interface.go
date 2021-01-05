package main

import "fmt"

type net_device_stats struct {
	rx_packets    uint64 /* total packets received       */
	tx_packets    uint64 /* total packets transmitted    */
	rx_bytes      uint64 /* total bytes received         */
	tx_bytes      uint64 /* total bytes transmitted      */
	rx_errors     uint64 /* bad packets received         */
	tx_errors     uint64 /* packet transmit problems     */
	rx_dropped    uint64 /* no space in linux buffers    */
	tx_dropped    uint64 /* no space available in linux  */
	rx_multicast  uint64 /* multicast packets received   */
	rx_compressed uint64
	tx_compressed uint64
	collisions    uint64

	/* detailed rx_errors: */
	rx_length_errors uint64
	rx_over_errors   uint64 /* receiver ring buff overflow  */
	rx_crc_errors    uint64 /* recved pkt with crc error    */
	rx_frame_errors  uint64 /* recv'd frame alignment error */
	rx_fifo_errors   uint64 /* recv'r fifo overrun          */
	rx_missed_errors uint64 /* receiver missed packet     */
	/* detailed tx_errors */
	tx_aborted_errors   uint64
	tx_carrier_errors   uint64
	tx_fifo_errors      uint64
	tx_heartbeat_errors uint64
	tx_window_errors    uint64
}
type iface struct {
	name       string
	iftype     int
	flags      int
	metric     int
	mtu        int
	txQueueLen int
	stats      net_device_stats
}

type ifaceList []*iface

var interfaces ifaceList

func addiface(inface *iface) {
	for i, curr := range interfaces {
		if curr.name == inface.name {
			interfaces[i] = inface
			return
		}
	}
	interfaces = append(interfaces, inface)
}

func removeiface(inface *iface) {
	for i, curr := range interfaces {
		if *curr == *inface {
			copy(interfaces[i:], interfaces[i+1:])
			interfaces[len(interfaces)-1] = nil
			interfaces = interfaces[:len(interfaces)-1]
		}
	}
}

func listiface() {
	for _, curr := range interfaces {
		fmt.Println("interface name is", curr.name)
	}
}
