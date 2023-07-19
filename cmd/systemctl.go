// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	yocke "github.com/ansurfen/yock/env"
	"github.com/ansurfen/yock/util"
)

func SystemCtlCreate(name string, opt SCCreateOpt) error {
	switch util.CurPlatform.OS {
	case "windows":
		return systemCtlCreateWindows(name, opt)
	case "linux":
		err := os.Chdir("/lib/systemd/system")
		if err != nil {
			return err
		}
		return systemCtlCreateLinux(name, opt)
	case "darwin":
		err := os.Chdir("~/Library/LaunchAgents")
		if err != nil {
			return err
		}
		return systemCtlCreateDarwin(name, opt)
	default:
		return fmt.Errorf("no implements")
	}
}

func systemCtlCreateWindows(name string, opt SCCreateOpt) error {
	args := NewArgsBuilder("sc create").Add(name).
		AddString(`binPath=""%s""`, opt.Service.ExecStart)
	str, err := OnceScript(fmt.Sprintf(`Dim objShell
Set objShell = WScript.CreateObject("WScript.Shell")
objShell.Run "%s"`, args.Build()))
	if err != nil {
		return fmt.Errorf("%s%s", str, err)
	}
	return nil
}

func systemCtlCreateLinux(name string, opt SCCreateOpt) error {
	err := os.WriteFile(name+".service", []byte(opt.String()), 0666)
	if err != nil {
		return err
	}
	return nil
}

func systemCtlCreateDarwin(name string, opt SCCreateOpt) error {
	fp, err := yocke.CreatePlistFile(fmt.Sprintf("%s.plist", name))
	if err != nil {
		return err
	}
	fp.Set(yocke.MetaMap{
		"Description": opt.Unit.Description,
		"Label":       name,
		"ProgramArguments": yocke.MetaArr{
			"bash", "-c", opt.Service.ExecStart,
		},
		"WorkingDirectory": opt.Service.WorkingDirectory,
	})
	if len(opt.Service.ExecStart) != 0 {
		fp.GetDict().Set("RunAtLoad", true)
	}
	err = fp.Write()
	if err != nil {
		panic(err)
	}
	return nil
}

type SCCreateOpt struct {
	Unit    scCreateOptUnit    `json:"unit"`
	Service scCreateOptService `json:"service"`
	Install scCreateOptInstall `json:"install"`
}

func (s *SCCreateOpt) String() string {
	return fmt.Sprintf(`[Unit]
%s
[Service]
%s
[Install]
%s`, s.Unit, s.Service, s.Install)
}

type scCreateOptUnit struct {
	Description   string `json:"description"`
	Before        string `json:"before"`
	After         string `json:"after"`
	Documentation string `json:"documentation"`
	Wants         string `json:"wants"`
	Requires      string `json:"requires"`
}

func (unit scCreateOptUnit) String() string {
	return fmt.Sprintf(`Description=%s
Before=%s
After=%s
`, unit.Description, unit.Before, unit.After)
}

type scCreateOptService struct {
	Type             string `json:"type"`
	ExecStart        string `json:"execStart"`
	ExecStop         string `json:"execStop"`
	ExecReload       string `json:"execReload"`
	ExecStartPre     string `json:"execStartPre"`
	ExecStartPost    string `json:"execStartPost"`
	ExecStopPost     string `json:"execStopPost"`
	PrivateTmp       bool   `json:"privateTmp"`
	RestartSec       int64  `json:"restartSec"`
	Restart          string `json:"restart"`
	User             string `json:"user"`
	Group            string `json:"group"`
	Environment      string `json:"environment"`
	WorkingDirectory string `json:"workingDirectory"`
}

func (service scCreateOptService) String() string {
	return fmt.Sprintf(`Type=%s
ExecStart=%s
ExecStop=%s
PrivateTmp=%v
RestartSec=%d
Restart=%s
`, service.Type, service.ExecStart, service.ExecStop, service.PrivateTmp, service.RestartSec, service.Restart)
}

type scCreateOptInstall struct {
	WantedBy string `json:"wantedBy"`
}

func (install scCreateOptInstall) String() string {
	return fmt.Sprintf(`WantedBy=%s
`, install.WantedBy)
}

type SystemCtlStatusOpt struct {
	Type   string `json:"type"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

type ServiceInfo interface {
	PID() int32
	Name() string
	Status() string
}

const (
	statusRunning = "running"
	statusStopped = "stopped"
	statusUnknown = "unknown"
)

var (
	_ ServiceInfo = (*windowsServiceInfo)(nil)
	_ ServiceInfo = (*linuxServiceInfo)(nil)
	_ ServiceInfo = (*darwinServiceInfo)(nil)
)

func SystemCtlStatus(opt SystemCtlStatusOpt) (infos []ServiceInfo, err error) {
	switch util.CurPlatform.OS {
	case "windows":
		args := NewArgsBuilder("sc queryex").AddString("%s", opt.Name).
			AddString("type=%s", opt.Type).AddString("state=%s", opt.Status)
		str, err := Exec(ExecOpt{Terminal: TermCmd, Quiet: true}, args.Build())
		if err != nil {
			return nil, fmt.Errorf("%s%s", str, err)
		}
		winInfos, err := systemCtlStatusWindows(str)
		if err != nil {
			return nil, err
		}
		for _, info := range winInfos {
			infos = append(infos, info)
		}
	case "linux":
		status := "--all"
		if len(opt.Status) > 0 {
			status = "--state=%s"
		}
		args := NewArgsBuilder("systemctl list-units").
			AddString("--type=%s", opt.Type).AddString(status, opt.Status)
		str, err := Exec(ExecOpt{Terminal: TermCmd, Quiet: true}, args.Build())
		if err != nil {
			return nil, fmt.Errorf("%s%s", str, err)
		}
		linuxInfos, err := systemCtlStatusLinux(str)
		if err != nil {
			return nil, err
		}
		for _, info := range linuxInfos {
			infos = append(infos, info)
		}
	case "darwin":
		args := NewArgsBuilder("launchctl list")
		str, err := Exec(ExecOpt{Quiet: true}, args.Build())
		if err != nil {
			return nil, fmt.Errorf("%s%s", str, err)
		}
		darwinInfos, err := systemCtlStatusDarwin(str)
		if err != nil {
			return nil, err
		}
		for _, info := range darwinInfos {
			infos = append(infos, info)
		}
	default:
		panic("no implements")
	}
	return
}

type windowsServiceInfo struct {
	ServiceName     string `json:"serviceName"`
	DisplayName     string `json:"displayName"`
	Type            string `json:"type"`
	State           int    `json:"state"`
	Win32ExitCode   uint32 `json:"win32ExitCode"`
	ServiceExitCode uint32 `json:"serviceExitCode"`
	Checkpoint      uint32 `json:"checkpoint"`
	WaitHint        uint32 `json:"waitHint"`
	Pid             int32  `json:"pid"`
	FLAGS           string `json:"flags"`
}

func (info windowsServiceInfo) PID() int32 {
	return info.Pid
}

func (info windowsServiceInfo) Name() string {
	return info.ServiceName
}

func (info windowsServiceInfo) Status() string {
	switch info.State {
	case 4:
		return statusRunning
	case 1:
		return statusStopped
	}
	return statusUnknown
}

func systemCtlStatusWindows(str string) (infos []windowsServiceInfo, err error) {
	infos = []windowsServiceInfo{}
	curService := windowsServiceInfo{}
	util.ReadLineFromString(str, func(s string) string {
		parts := strings.SplitN(s, ":", 2)
		if len(parts) != 2 {
			if len(strings.TrimSpace(s)) == 0 && len(curService.ServiceName) != 0 {
				infos = append(infos, curService)
				curService = windowsServiceInfo{}
			}
			return ""
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		switch key {
		case "SERVICE_NAME":
			curService.ServiceName = value
		case "DISPLAY_NAME":
			curService.DisplayName = value
		case "TYPE":
			curService.Type = value
		case "STATE":
			stateParts := strings.SplitN(value, " ", 2)
			if len(stateParts) == 2 {
				stateStr := strings.TrimSpace(stateParts[0])
				stateInt, err := strconv.Atoi(stateStr)
				if err == nil {
					curService.State = stateInt
				}
			}
		case "WIN32_EXIT_CODE":
			exitCodeParts := strings.SplitN(value, " ", 3)
			if len(exitCodeParts) == 3 {
				exitCodeStr := strings.TrimSpace(exitCodeParts[1])
				exitCode, err := strconv.ParseUint(exitCodeStr, 0, 32)
				if err == nil {
					curService.Win32ExitCode = uint32(exitCode)
				}
			}
		case "SERVICE_EXIT_CODE":
			exitCodeParts := strings.SplitN(value, " ", 3)
			if len(exitCodeParts) == 3 {
				exitCodeStr := strings.TrimSpace(exitCodeParts[1])
				exitCode, err := strconv.ParseUint(exitCodeStr, 0, 32)
				if err == nil {
					curService.ServiceExitCode = uint32(exitCode)
				}
			}
		case "CHECKPOINT":
			checkpointStr := strings.TrimPrefix(value, "0x")
			checkpoint, err := strconv.ParseUint(checkpointStr, 16, 32)
			if err == nil {
				curService.Checkpoint = uint32(checkpoint)
			}
		case "WAIT_HINT":
			waitHintStr := strings.TrimPrefix(value, "0x")
			waitHint, err := strconv.ParseUint(waitHintStr, 16, 32)
			if err == nil {
				curService.WaitHint = uint32(waitHint)
			}
		case "PID":
			pidStr := strings.TrimSpace(value)
			pid, err := strconv.ParseInt(pidStr, 10, 32)
			if err == nil {
				curService.Pid = int32(pid)
			}
		case "FLAGS":
			curService.FLAGS = value
		}
		return ""
	})
	if len(curService.ServiceName) != 0 {
		infos = append(infos, curService)
	}
	return
}

type linuxServiceInfo struct {
	Unit        string `json:"unit"`
	Load        string `json:"load"`
	Active      string `json:"active"`
	Sub         string `json:"sub"`
	Description string `json:"description"`
	Pid         int32  `json:"pid"`
}

func (info linuxServiceInfo) PID() int32 {
	return info.Pid
}

func (info linuxServiceInfo) Name() string {
	return info.Unit
}

func (info linuxServiceInfo) Status() string {
	return statusUnknown
}

func systemCtlStatusLinux(str string) (infos []linuxServiceInfo, err error) {
	idx := strings.Index(str, "\n")
	if idx == -1 {
		return
	}
	re := regexp.MustCompile(`\s+`)
	util.ReadLineFromString(str[idx:], func(s string) string {
		s = strings.TrimSpace(s)
		if len(s) == 0 || s == "0 loaded units listed." {
			return ""
		}
		s = re.ReplaceAllString(s, " ")
		info := linuxServiceInfo{}
		step := 0
		for i := 0; i < len(s); i++ {
			ch := s[i]
			if ch == ' ' {
				step++
				if step <= 3 {
					continue
				}
			}
			switch step {
			case 0:
				info.Unit += string(ch)
			case 1:
				info.Load += string(ch)
			case 2:
				info.Active += string(ch)
			case 3:
				info.Sub += string(ch)
			default:
				info.Description += string(ch)
			}
		}
		info.Description = strings.TrimPrefix(info.Description, " ")
		infos = append(infos, info)
		return ""
	})
	return
}

type darwinServiceInfo struct {
	Pid   string `json:"pid"`
	State string `json:"status"`
	Label string `json:"label"`
}

func (info darwinServiceInfo) PID() int32 {
	if info.Pid == "-" {
		return 0
	}
	i, err := strconv.Atoi(info.Pid)
	if err != nil {
		return 0
	}
	return int32(i)
}

func (info darwinServiceInfo) Name() string {
	return info.Label
}

func (info darwinServiceInfo) Status() string {
	if info.State == "0" {
		return statusRunning
	} else if len(info.State) > 0 && info.State[0] == '-' {
		return statusStopped
	}
	return statusUnknown
}

func systemCtlStatusDarwin(str string) (infos []darwinServiceInfo, err error) {
	idx := strings.Index(str, "\n")
	if idx == -1 {
		return
	}
	re := regexp.MustCompile(`\s+`)
	util.ReadLineFromString(str[idx:], func(s string) string {
		if len(s) == 0 {
			return ""
		}
		s = re.ReplaceAllString(s, " ")
		res := strings.SplitN(s, " ", 3)
		infos = append(infos, darwinServiceInfo{
			Pid:   res[0],
			State: res[1],
			Label: res[2],
		})
		return ""
	})
	return
}

func SystemCtlIsEnable(name string) bool {
	switch util.CurPlatform.OS {
	case "windows":
	case "linux":
		args := NewArgsBuilder("systemctl is-enabled").AddString("%s", name)
		str, err := Exec(ExecOpt{}, args.Build())
		if err != nil {
			return false
		}
		if str == "enabled" {
			return true
		}
	case "darwin":
	}
	return false
}

func SystemCtlStart(server string) error {
	cmd := ""
	switch util.CurPlatform.OS {
	case "windows":
		cmd = "sc start %s"
	case "linux":
		cmd = "systemctl start %s"
	case "darwin":
		err := os.Chdir("~/Library/LaunchAgents")
		if err != nil {
			return err
		}
		cmd = fmt.Sprintf("launchctl load %s && launchctl start %%s", server)
	default:
		panic("no support")
	}
	_, err := Exec(ExecOpt{Quiet: true}, fmt.Sprintf(cmd, server))
	return err
}

func SystemCtlStop(server string) error {
	cmd := ""
	switch util.CurPlatform.OS {
	case "windows":
		cmd = "sc stop %s"
	case "linux":
		cmd = "systemctl stop %s"
	case "darwin":
		cmd = "launchctl stop %s"
	default:
		panic("no support")
	}
	_, err := Exec(ExecOpt{Quiet: true}, fmt.Sprintf(cmd, server))
	return err
}

func SystemCtlDisable(server string) error {
	cmd := ""
	switch util.CurPlatform.OS {
	case "windows":
		cmd = "sc config %s start=disabled"
	case "linux":
		cmd = "systemctl disable %s"
	case "darwin":
		cmd = "launchctl unload %s"
	default:
		panic("no support")
	}
	_, err := Exec(ExecOpt{Quiet: true}, fmt.Sprintf(cmd, server))
	return err
}

func SystemCtlDelete(server string) error {
	switch util.CurPlatform.OS {
	case "windows":
		str, err := Exec(ExecOpt{Quiet: true, Terminal: TermCmd}, fmt.Sprintf("sc delete %s", server))
		if err != nil {
			return fmt.Errorf("%s%s", str, err)
		}
		return nil
	case "linux":
		err := SystemCtlStop(server)
		if err != nil {
			return err
		}
		err = SystemCtlDisable(server)
		if err != nil {
			return err
		}
		err = os.Chdir("/lib/systemd/system")
		if err != nil {
			return err
		}
		err = os.Remove(server + ".service")
		if err != nil {
			return err
		}
	case "darwin":
		err := os.Chdir("~/Library/LaunchAgents")
		if err != nil {
			return err
		}
		if err := os.Remove(fmt.Sprintf("%s.plist", server)); err != nil {
			return err
		}
	default:
		panic("no support")
	}
	return nil
}

func SystemCtlRelaod(server string) error {
	cmd := ""
	switch util.CurPlatform.OS {
	case "windows":
		cmd = fmt.Sprintf("sc stop %s & sc start %s", server, server)
	case "linux":
		cmd = fmt.Sprintf("systemctl restart %s", server)
	case "darwin":
		cmd = fmt.Sprintf("launchctl stop %s & launchctl start %s", server, server)
	default:
		panic("no support")
	}
	_, err := Exec(ExecOpt{Quiet: true}, cmd)
	return err
}
