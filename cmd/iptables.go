// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ansurfen/yock/util"
)

type FireWareRule interface {
	Name() string
	Proto() string
	Src() string
	Dst() string
	Action() string
}

var (
	_ FireWareRule = (*firewareRuleWindows)(nil)
	_ FireWareRule = (*fireWareRuleLinux)(nil)
	_ FireWareRule = (*fireWareRuleDarwin)(nil)
)

type IPTablesListOpt struct {
	Legacy bool
	Chain  string
	Name   string
}

func IPTablesList(opt IPTablesListOpt) (rules []FireWareRule, err error) {
	switch util.CurPlatform.OS {
	case "windows":
		cmd := "netsh advfirewall firewall show rule name=%s"
		if len(opt.Name) == 0 {
			cmd = fmt.Sprintf(cmd, "all")
		} else {
			cmd = fmt.Sprintf(cmd, fmt.Sprintf(`"%s"`, opt.Name))
		}
		str, err := Exec(ExecOpt{Quiet: true, Terminal: TermCmd}, cmd)
		if err != nil {
			return nil, err
		}
		rs, err := IPTablesListWindows(str)
		if err != nil {
			return nil, err
		}
		for _, rule := range rs {
			rules = append(rules, rule)
		}
	case "linux":
		var args *ArgsBuilder
		if opt.Legacy {
			args = NewArgsBuilder("iptables-legacy -L")
		} else {
			args = NewArgsBuilder("iptables -L")
		}
		args.AddString("%s", opt.Chain)
		str, err := Exec(ExecOpt{Quiet: true}, args.Build())
		if err != nil {
			return nil, err
		}
		rs, err := IPTablesListLinux(str)
		if err != nil {
			return nil, err
		}
		for _, rule := range rs {
			rules = append(rules, rule)
		}
	case "darwin":
		cmd := "pfctl -sr"
		str, err := Exec(ExecOpt{Quiet: true}, cmd)
		if err != nil {
			return nil, err
		}
		rs, err := IPTablesListDarwin(str)
		if err != nil {
			return nil, err
		}
		for _, rule := range rs {
			rules = append(rules, rule)
		}
	}
	return
}

type fireWareRuleLinux struct {
	Chain       string `json:"chain"`
	Target      string `json:"target"`
	Protocol    string `json:"protocol"`
	Option      string `json:"option"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

func (rule fireWareRuleLinux) Name() string {
	return rule.Target
}

func (rule fireWareRuleLinux) Proto() string {
	return rule.Protocol
}

func (rule fireWareRuleLinux) Src() string {
	return rule.Source
}

func (rule fireWareRuleLinux) Dst() string {
	return rule.Destination
}

func (rule fireWareRuleLinux) Action() string {
	return rule.Chain
}

func IPTablesListLinux(str string) (rules []fireWareRuleLinux, err error) {
	chain := regexp.MustCompile(`Chain (.*)`)
	re := regexp.MustCompile(`\s+`)
	curChain := ""
	rule := fireWareRuleLinux{}
	util.ReadLineFromString(str, func(s string) string {
		if chain.MatchString(s) {
			curChain = chain.FindStringSubmatch(s)[1]
			return ""
		}
		if strings.Contains(s, "target") {
			return ""
		}
		rule.Chain = curChain
		s = re.ReplaceAllString(s, " ")
		str := strings.SplitN(s, " ", 5)
		if len(str) == 5 {
			rules = append(rules, fireWareRuleLinux{
				Chain:       curChain,
				Target:      str[0],
				Protocol:    str[1],
				Option:      str[2],
				Source:      str[3],
				Destination: strings.TrimSuffix(str[4], " "),
			})
		}
		return ""
	})
	return
}

type fireWareRuleDarwin struct {
	Type      string `json:"type"`
	Interface string `json:"interface"`
	From      string `json:"from"`
	To        string `json:"to"`
	Act       string `json:"action"`
}

func (rule fireWareRuleDarwin) Name() string {
	return rule.Interface
}

func (rule fireWareRuleDarwin) Proto() string {
	return ""
}

func (rule fireWareRuleDarwin) Src() string {
	return rule.From
}

func (rule fireWareRuleDarwin) Dst() string {
	return rule.To
}

func (rule fireWareRuleDarwin) Action() string {
	return rule.Act
}

func IPTablesListDarwin(str string) (rules []fireWareRuleDarwin, err error) {
	natReg := regexp.MustCompile(`^nat\s+on\s+(\S+)\s+inet\s+from\s+(\S+)\s+to\s+(\S+)\s+->\s+\((\S+)\)\s+(\S+)$`)
	util.ReadLineFromString(str, func(s string) string {
		s = strings.TrimSpace(s)
		if len(s) == 0 || strings.HasPrefix(s, "#") || strings.Contains(s, "ALTQ") {
			return ""
		}
		if ss := natReg.FindStringSubmatch(s); len(ss) > 1 {
			rules = append(rules, fireWareRuleDarwin{
				Type:      "nat",
				Interface: ss[1],
				From:      ss[2],
				To:        ss[3],
				Act:       ss[5],
			})
			return ""
		}
		if ss := strings.SplitN(s, " ", 3); len(ss) == 3 {
			rules = append(rules, fireWareRuleDarwin{
				Type:      ss[0],
				Interface: ss[1],
				Act:       ss[2],
			})
		}
		return ""
	})
	return
}

type firewareRuleWindows struct {
	RuleName      string `json:"name"`
	Enabled       string `json:"enabled"`
	Direction     string `json:"direction"`
	Profiles      string `json:"profiles"`
	Grouping      string `json:"grouping"`
	LocalIP       string `json:"localIP"`
	RemoteIP      string `json:"remoteIP"`
	Protocol      string `json:"protocol"`
	LocalPort     string `json:"localPort"`
	RemotePort    string `json:"remotePort"`
	EdgeTraversal string `json:"edgeTraversal"`
	Act           string `json:"action"`
}

func (rule firewareRuleWindows) Name() string {
	return rule.RuleName
}

func (rule firewareRuleWindows) Proto() string {
	return rule.Protocol
}

func (rule firewareRuleWindows) Src() string {
	return ""
}

func (rule firewareRuleWindows) Dst() string {
	return ""
}

func (rule firewareRuleWindows) Action() string {
	return rule.Act
}

func IPTablesListWindows(str string) (rules []firewareRuleWindows, err error) {
	rule := firewareRuleWindows{}
	banner := false
	line := 0
	util.ReadLineFromString(str, func(s string) string {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			return ""
		}
		if strings.HasPrefix(s, "---") {
			banner = true
			line = 0
			return ""
		}
		kv := strings.SplitN(s, ":", 2)
		if len(kv) != 2 {
			if banner {
				line++
			}
			return ""
		}
		value := strings.TrimSpace(kv[1])
		if banner {
			switch line {
			case 0:
				rule.Enabled = value
			case 1:
				rule.Direction = value
			case 2:
				rule.Profiles = value
			case 3:
				rule.Grouping = value
			case 4:
				rule.LocalIP = value
			case 5:
				rule.RemoteIP = value
			case 6:
				rule.Protocol = value
			case 7:
				rule.LocalPort = value
			case 8:
				rule.RemotePort = value
			case 9:
				rule.EdgeTraversal = value
			case 10:
				rule.Act = value
				rules = append(rules, rule)
				rule = firewareRuleWindows{}
				banner = false
				line = 0
				return ""
			}
			line++
		} else {
			rule.RuleName = value
		}
		return ""
	})
	return
}

type IPTablesOpOpt struct {
	Chain       string `json:"chain"`
	Name        string `json:"name"`
	Protocol    string `json:"protocol"`
	Destination string `json:"dest"`
	// linux: ACCEPT, DROP, REJECT, LOG, SNAT, DNAT, MASQUEREAD
	Action string          `json:"action"`
	Legacy bool            `json:"legacy"`
	Op     IPTablesOperate `json:"op"`
}

type IPTablesOperate byte

const (
	iPTablesUnknown IPTablesOperate = iota
	IPTablesAdd
	IPTablesDel
)

func (op IPTablesOperate) value() IPTablesOperate {
	switch op {
	case IPTablesAdd, IPTablesDel:
		return op
	}
	panic("invalid operator")
}

func (opt IPTablesOpOpt) ToWindows() string {
	return iPTablesOpOptWindows{
		LocalPort: opt.Destination,
		Protocol:  opt.Protocol,
		Name:      opt.Name,
		Dir:       winDir(opt.Chain),
		Action:    winAction(opt.Action),
		Op:        opt.Op,
	}.String()
}

func (opt IPTablesOpOpt) ToLinux() string {
	var cmd *ArgsBuilder
	op := ""
	switch opt.Op.value() {
	case IPTablesAdd:
		op = "-A"
	case IPTablesDel:
		op = "-D"
	}
	if opt.Legacy {
		cmd = NewArgsBuilder(fmt.Sprintf("iptables-legacy %s", op))
	} else {
		cmd = NewArgsBuilder(fmt.Sprintf("iptables %s", op))
	}
	cmd.AddString("%s", opt.Chain).AddString("-p %s", opt.Protocol).
		AddString("--dport %s", opt.Destination)
	return cmd.Build()
}

func IPTablesOp(opt IPTablesOpOpt) error {
	var term *Terminal
	switch util.CurPlatform.OS {
	case "windows":
		term = cmdTerm(opt.ToWindows())
	case "linux":
		term = bashTerm(opt.ToLinux())
	case "darwin":
	}
	str, err := term.Exec(&ExecOpt{Quiet: true})
	if err != nil {
		return fmt.Errorf("%s%s", string(str), err)
	}
	return nil
}

type iPTablesOpOptWindows struct {
	Name      string          `json:"name"`
	Dir       winDir          `json:"dir"`
	Action    winAction       `json:"action"`
	Protocol  string          `json:"protocol"`
	LocalPort string          `json:"localPort"`
	Op        IPTablesOperate `json:"op"`
}

func (opt iPTablesOpOptWindows) String() string {
	var cmd *ArgsBuilder
	switch opt.Op.value() {
	case IPTablesAdd:
		cmd = NewArgsBuilder("netsh advfirewall firewall add rule").
			AddString(`name="%s"`, opt.Name).AddString("dir=%s", opt.Dir.String()).
			AddString("action=%s", opt.Action.String()).AddString("protocol=%s", opt.Protocol).
			AddString("localport=%s", opt.LocalPort)
	case IPTablesDel:
		cmd = NewArgsBuilder("netsh advfirewall firewall delete rule").
			AddString(`name="%s"`, opt.Name)
	}
	return cmd.Build()
}

type winDir string

func (dir winDir) String() string {
	switch v := strings.ToLower(string(dir)); v {
	case "input":
		return "in"
	case "output":
		return "out"
	// netsh
	case "in", "out", "fw":
		return v
	}
	panic("unreachable")
}

type winAction string

func (act winAction) String() string {
	switch v := strings.ToLower(string(act)); v {
	case "accept":
		return "allow"
	case "drop":
		return "block"
	// netsh
	case "allow", "block", "bypass", "custom", "default", "notconfigured", "reset":
		return v
	}
	panic("unreachable")
}
