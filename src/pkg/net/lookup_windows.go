// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"syscall"
	"unsafe"
	"os"
	"sync"
)

var hostentLock sync.Mutex
var serventLock sync.Mutex

func LookupHost(name string) (addrs []string, err os.Error) {
	ips, err := LookupIP(name)
	if err != nil {
		return
	}
	addrs = make([]string, 0, len(ips))
	for _, ip := range ips {
		addrs = append(addrs, ip.String())
	}
	return
}

func LookupIP(name string) (addrs []IP, err os.Error) {
	hostentLock.Lock()
	defer hostentLock.Unlock()
	h, e := syscall.GetHostByName(name)
	if e != 0 {
		return nil, os.NewSyscallError("GetHostByName", e)
	}
	switch h.AddrType {
	case syscall.AF_INET:
		i := 0
		addrs = make([]IP, 100) // plenty of room to grow
		for p := (*[100](*[4]byte))(unsafe.Pointer(h.AddrList)); i < cap(addrs) && p[i] != nil; i++ {
			addrs[i] = IPv4(p[i][0], p[i][1], p[i][2], p[i][3])
		}
		addrs = addrs[0:i]
	default: // TODO(vcc): Implement non IPv4 address lookups.
		return nil, os.NewSyscallError("LookupHost", syscall.EWINDOWS)
	}
	return addrs, nil
}

func LookupPort(network, service string) (port int, err os.Error) {
	switch network {
	case "tcp4", "tcp6":
		network = "tcp"
	case "udp4", "udp6":
		network = "udp"
	}
	serventLock.Lock()
	defer serventLock.Unlock()
	s, e := syscall.GetServByName(service, network)
	if e != 0 {
		return 0, os.NewSyscallError("GetServByName", e)
	}
	return int(syscall.Ntohs(s.Port)), nil
}

func LookupCNAME(name string) (cname string, err os.Error) {
	var r *syscall.DNSRecord
	e := syscall.DnsQuery(name, syscall.DNS_TYPE_CNAME, 0, nil, &r, nil)
	if int(e) != 0 {
		return "", os.NewSyscallError("LookupCNAME", int(e))
	}
	defer syscall.DnsRecordListFree(r, 1)
	if r != nil && r.Type == syscall.DNS_TYPE_CNAME {
		v := (*syscall.DNSPTRData)(unsafe.Pointer(&r.Data[0]))
		cname = syscall.UTF16ToString((*[256]uint16)(unsafe.Pointer(v.Host))[:]) + "."
	}
	return
}

func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err os.Error) {
	var r *syscall.DNSRecord
	target := "_" + service + "._" + proto + "." + name
	e := syscall.DnsQuery(target, syscall.DNS_TYPE_SRV, 0, nil, &r, nil)
	if int(e) != 0 {
		return "", nil, os.NewSyscallError("LookupSRV", int(e))
	}
	defer syscall.DnsRecordListFree(r, 1)
	addrs = make([]*SRV, 100)
	i := 0
	for p := r; p != nil && p.Type == syscall.DNS_TYPE_SRV; p = p.Next {
		v := (*syscall.DNSSRVData)(unsafe.Pointer(&p.Data[0]))
		addrs[i] = &SRV{syscall.UTF16ToString((*[256]uint16)(unsafe.Pointer(v.Target))[:]), v.Port, v.Priority, v.Weight}
		i++
	}
	addrs = addrs[0:i]
	return name, addrs, nil
}

// TODO(brainman): implement LookupMX and LookupAddr.

func LookupMX(name string) (mx []*MX, err os.Error) {
	return nil, os.NewSyscallError("LookupMX", syscall.EWINDOWS)
}

func LookupAddr(addr string) (name []string, err os.Error) {
	return nil, os.NewSyscallError("LookupAddr", syscall.EWINDOWS)
}
