// Command sysinfo prints basic information about the host it runs on.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/user"
	"runtime"
)

// Interface describes a non-loopback network interface.
type Interface struct {
	Name         string   `json:"name"`
	HardwareAddr string   `json:"hardware_addr,omitempty"`
	Addrs        []string `json:"addrs,omitempty"`
}

// Info holds the collected system information.
type Info struct {
	CPUs       int         `json:"cpus"`
	Hostname   string      `json:"hostname"`
	OS         string      `json:"os"`
	Arch       string      `json:"arch"`
	TempDir    string      `json:"temp_dir"`
	GoVersion  string      `json:"go_version"`
	Username   string      `json:"username,omitempty"`
	UID        string      `json:"uid,omitempty"`
	GID        string      `json:"gid,omitempty"`
	Interfaces []Interface `json:"interfaces"`
}

func main() {
	jsonOut := flag.Bool("json", false, "print output as JSON")
	flag.Parse()

	info, errs := collect()
	for _, err := range errs {
		fmt.Fprintln(os.Stderr, "sysinfo:", err)
	}

	var err error
	if *jsonOut {
		err = writeJSON(os.Stdout, info)
	} else {
		err = writeText(os.Stdout, info)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "sysinfo:", err)
		os.Exit(1)
	}
}

// collect gathers system information. Partial results are returned
// alongside any errors encountered.
func collect() (Info, []error) {
	var errs []error
	info := Info{
		CPUs:      runtime.NumCPU(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		TempDir:   os.TempDir(),
		GoVersion: runtime.Version(),
	}

	hostname, err := os.Hostname()
	if err != nil {
		errs = append(errs, fmt.Errorf("hostname: %w", err))
	}
	info.Hostname = hostname

	if u, err := user.Current(); err != nil {
		errs = append(errs, fmt.Errorf("user: %w", err))
	} else {
		info.Username = u.Username
		info.UID = u.Uid
		info.GID = u.Gid
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		errs = append(errs, fmt.Errorf("interfaces: %w", err))
	}
	info.Interfaces = make([]Interface, 0, len(ifaces))
	for _, iface := range ifaces {
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		i := Interface{Name: iface.Name}
		if len(iface.HardwareAddr) > 0 {
			i.HardwareAddr = iface.HardwareAddr.String()
		}
		addrs, err := iface.Addrs()
		if err != nil {
			errs = append(errs, fmt.Errorf("addrs for %s: %w", iface.Name, err))
		}
		for _, addr := range addrs {
			i.Addrs = append(i.Addrs, addr.String())
		}
		info.Interfaces = append(info.Interfaces, i)
	}

	return info, errs
}

func writeText(w io.Writer, info Info) error {
	lines := []struct{ label, value string }{
		{"CPUs", fmt.Sprint(info.CPUs)},
		{"Hostname", info.Hostname},
		{"OS", info.OS},
		{"OS arch", info.Arch},
		{"Temp dir", info.TempDir},
		{"Go version", info.GoVersion},
		{"User", info.Username},
		{"UID", info.UID},
		{"GID", info.GID},
	}
	for _, l := range lines {
		if l.value == "" {
			continue
		}
		if _, err := fmt.Fprintf(w, "%s: %s\n", l.label, l.value); err != nil {
			return err
		}
	}
	for _, iface := range info.Interfaces {
		if iface.HardwareAddr != "" {
			if _, err := fmt.Fprintln(w, iface.Name, iface.HardwareAddr); err != nil {
				return err
			}
		}
		for _, addr := range iface.Addrs {
			if _, err := fmt.Fprintln(w, iface.Name, "->", addr); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeJSON(w io.Writer, info Info) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(info)
}
