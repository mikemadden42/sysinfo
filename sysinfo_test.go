package main

import (
	"bytes"
	"encoding/json"
	"runtime"
	"strings"
	"testing"
)

func TestCollect(t *testing.T) {
	info, errs := collect()
	for _, err := range errs {
		t.Logf("collect: %v", err)
	}
	if info.CPUs < 1 {
		t.Errorf("CPUs = %d, want >= 1", info.CPUs)
	}
	if info.OS != runtime.GOOS {
		t.Errorf("OS = %q, want %q", info.OS, runtime.GOOS)
	}
	for _, iface := range info.Interfaces {
		for _, addr := range iface.Addrs {
			if strings.HasPrefix(addr, "127.") || addr == "::1/128" {
				t.Errorf("loopback address %s on %s should be filtered", addr, iface.Name)
			}
		}
	}
}

var sample = Info{
	CPUs:      4,
	Hostname:  "host",
	OS:        "linux",
	Arch:      "amd64",
	TempDir:   "/tmp",
	GoVersion: "go1.23.0",
	UID:       "1000",
	GID:       "1000",
	Interfaces: []Interface{
		{Name: "eth0", HardwareAddr: "00:11:22:33:44:55", Addrs: []string{"10.0.0.2/24"}},
		{Name: "tun0", Addrs: []string{"10.8.0.1/24"}},
	},
}

func TestWriteText(t *testing.T) {
	var buf bytes.Buffer
	if err := writeText(&buf, sample); err != nil {
		t.Fatal(err)
	}
	want := `CPUs: 4
Hostname: host
OS: linux
OS arch: amd64
Temp dir: /tmp
Go version: go1.23.0
UID: 1000
GID: 1000
eth0 00:11:22:33:44:55
eth0 -> 10.0.0.2/24
tun0 -> 10.8.0.1/24
`
	if got := buf.String(); got != want {
		t.Errorf("writeText output mismatch\ngot:\n%s\nwant:\n%s", got, want)
	}
}

func TestWriteJSON(t *testing.T) {
	var buf bytes.Buffer
	if err := writeJSON(&buf, sample); err != nil {
		t.Fatal(err)
	}
	var got Info
	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if got.Hostname != sample.Hostname || len(got.Interfaces) != len(sample.Interfaces) {
		t.Errorf("round trip mismatch: %+v", got)
	}
}
