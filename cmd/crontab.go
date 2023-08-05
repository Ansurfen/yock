// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ansurfen/yock/util"
	"github.com/beevik/etree"
)

type CronTask interface {
	Expr() string
	Cmd() string
}

var (
	_ CronTask = (*cronTaskPosix)(nil)
	_ CronTask = (*cronTaskWindows)(nil)
)

func CrontabList(expr string) (tasks []CronTask, err error) {
	if util.CurPlatform.OS == "windows" {
		args := NewArgsBuilder("schtasks /query").
			AddString("/tn %s", expr).Add("/XML")
		str, err := Exec(ExecOpt{Quiet: true}, args.Build())
		if err != nil {
			return nil, fmt.Errorf("%s%s", str, err)
		}
		ls := false
		if len(expr) == 0 {
			ls = true
		}
		ts, err := crontabListWindowsV2(str, ls)
		if err != nil {
			return nil, err
		}
		for _, task := range ts {
			tasks = append(tasks, task)
		}
	} else {
		args := NewArgsBuilder("crontab -l")
		str, err := Exec(ExecOpt{Quiet: true}, args.Build())
		if err != nil {
			return nil, fmt.Errorf("%s%s", str, err)
		}
		ts, err := crontabListPosix(str)
		if err != nil {
			return nil, err
		}
		for _, task := range ts {
			tasks = append(tasks, task)
		}
	}
	return
}

type cronTaskPosix struct {
	Cron    string `json:"cron"`
	Command string `json:"command"`
}

func (ct cronTaskPosix) Expr() string {
	return ct.Cron
}

func (ct cronTaskPosix) Cmd() string {
	return ct.Command
}

func crontabListPosix(str string) (tasks []cronTaskPosix, err error) {
	util.ReadLineFromString(str, func(s string) string {
		expr, cmd := CronParse(str)
		suffix := strings.LastIndex(expr.String(), " ")
		tasks = append(tasks, cronTaskPosix{
			Cron:    strings.TrimSpace(expr.String()[1:suffix]),
			Command: strings.TrimSpace(cmd),
		})
		return ""
	})
	return
}

type cronTaskWindows struct {
	Folder          string `json:"folder"`
	Name            string `json:"name"`
	NextRunTime     string `json:"nextRunTime"`
	Status          string `json:"status"`
	CalendarTrigger string `json:"calendarTrigger"`
	Command         string `json:"command"`
}

func (ct cronTaskWindows) Cmd() string {
	return ct.Command
}

func (ct cronTaskWindows) Expr() string {
	return ct.CalendarTrigger
}

func crontabListWindowsV1(str string) (tasks []cronTaskWindows, err error) {
	folder := ""
	skip := false
	re := regexp.MustCompile(`(N/A|\d+/\d+/\d+ \d+:\d+:\d+)`)
	util.ReadLineFromString(str, func(s string) string {
		if len(s) == 0 {
			return ""
		}
		if skip {
			skip = false
			return ""
		}
		if strings.Contains(s, ":") && strings.Contains(s, "\\") {
			kv := strings.Split(s, ":")
			folder = strings.TrimSpace(kv[1])
			skip = true
			return ""
		}
		if strings.Contains(s, "=") {
			return ""
		}
		if re.MatchString(s) {
			loc := re.FindAllStringIndex(s, 1)[0]
			tasks = append(tasks, cronTaskWindows{
				Folder:      folder,
				Name:        strings.TrimSpace(s[:loc[0]]),
				NextRunTime: s[loc[0]:loc[1]],
				Status:      strings.TrimSpace(s[loc[1]:]),
			})
		}
		return ""
	})
	return
}

func crontabListWindowsV2(str string, ls bool) (tasks []cronTaskWindows, err error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(str); err != nil {
		return nil, err
	}
	var docs []*etree.Element
	if ls {
		docs = doc.SelectElement("Tasks").SelectElements("Task")
	} else {
		docs = doc.SelectElements("Task")
	}
	for _, task := range docs {
		if action := task.SelectElement("Actions"); action != nil {
			if exec := action.SelectElement("Exec"); exec != nil {
				if cmd := exec.SelectElement("Command"); cmd != nil {
					tasks = append(tasks, cronTaskWindows{
						Command: cmd.Text(),
					})
				}
			}
		}
	}
	return
}

func CrontabAdd(expr string, names ...string) error {
	switch util.CurPlatform.OS {
	case "windows":
		// 创建bat脚本
	default:
		str, err := OnceScript(
			fmt.Sprintf(`(crontab -l 2>/dev/null; echo "%s") | crontab -`, expr))
		if err != nil {
			return fmt.Errorf("%s%s", str, err)
		}
	}
	return nil
}

func CrontabDel(expr string, names ...string) error {
	return nil
}

type cronExpr struct {
	sec  cronAnchor
	min  cronAnchor
	hour cronAnchor
	day  cronAnchor
	mon  cronAnchor
	week cronAnchor
	year cronAnchor
}

func (expr cronExpr) String() string {
	str := []string{
		expr.sec.String(),
		expr.min.String(),
		expr.hour.String(),
		expr.day.String(),
		expr.mon.String(),
		expr.week.String(),
		expr.year.String(),
	}
	return strings.Join(str, " ")
}

type timeline struct {
	schdule   string
	mo        cronTime
	startTime *util.Duration
	endTime   *util.Duration
	startDate *util.Time
	endDate   *util.Time
	day       string
}

func (s *timeline) String() string {
	st := s.startTime.Format(":")
	if st == "00:00:00" {
		st = ""
	} else {
		st = st[:len(st)-3]
	}
	et := s.endTime.Format(":")
	if et == "00:00:00" {
		et = ""
	} else {
		et = et[:len(et)-3]
	}
	sd := s.startDate.Format("2006/01/02")
	if sd == "01/01/0001" {
		sd = ""
	}
	ed := s.endDate.Format("2006/01/02")
	if ed == "01/01/0001" {
		ed = ""
	} else if s.endDate.Year() == 1 {
		s.endDate.AddDate(9998, 0, 0)
		ed = s.endDate.Format("2006/01/02")
	}
	mo := s.mo.Str()
	if s.mo == 1 || s.mo == 0 {
		mo = ""
	}
	return NewArgsBuilder().AddString("/SC %s", strings.ToUpper(s.schdule)).
		AddString("/MO %s", mo).AddString("/D %s", s.day).
		AddString("/ST %s", st).AddString("/ET %s", et).
		AddString("/SD %s", sd).AddString("/ED %s", ed).Build()
}

func (s *timeline) Clone() *timeline {
	tl := &timeline{schdule: s.schdule, mo: s.mo, day: s.day}
	if s.startTime != nil {
		tl.startTime = s.startTime.Clone()
	}
	if s.endTime != nil {
		tl.endTime = s.endTime.Clone()
	}
	if s.startDate != nil {
		tl.startDate = s.startDate.Clone()
	}
	if s.endDate != nil {
		tl.endDate = s.endDate.Clone()
	}
	return tl
}

func newTimeline() *timeline {
	return &timeline{
		mo:        cronTime(1),
		startTime: util.ParseDuration(""),
		endTime:   util.ParseDuration(""),
		startDate: util.ParseTime("2006", ""),
		endDate:   util.ParseTime("2006", ""),
	}
}

type schtasks struct {
	tls     []*timeline
	sc      string
	delaySc string
	week    bool
}

func (st *schtasks) hasSc() bool {
	return len(st.sc) != 0
}

func (st *schtasks) rrange(fn func(tl *timeline) error) error {
	if len(st.tls) == 0 {
		st.tls = append(st.tls, newTimeline())
	}
	for _, tl := range st.tls {
		if err := fn(tl); err != nil {
			return err
		}
	}
	return nil
}

func (st *schtasks) parseYear(anchor cronAnchor) bool {
	years := []cronTime{}
	if anchor.hasStep() {
		step := anchor.getStep()
		base := anchor.times()[0]
		for i := 0; i < 3; i++ {
			years = append(years, base)
			base = base + step
		}
	} else if anchor.hasRound() {
		year := anchor.times()
		for i := 0; i <= int(year[1]-year[0]); i++ {
			years = append(years, year[0]+cronTime(i))
		}
	} else {
		years = append(years, anchor.times()...)
	}
	results := [][]cronTime{{years[0]}}
	for i := 1; i < len(years); i++ {
		prev := results[len(results)-1]
		curr := years[i]
		if curr-prev[len(prev)-1] == 1 {
			prev = append(prev, curr)
			results[len(results)-1] = prev
		} else {
			results = append(results, []cronTime{curr})
		}
	}
	for _, interval := range results {
		if len(interval) > 1 {
			at := util.ParseTime("2006", interval[0].Str())
			bt := util.ParseTime("2006", interval[len(interval)-1].Str())
			st.tls = append(st.tls, &timeline{
				startDate: at,
				endDate:   bt,
				startTime: util.ParseDuration("0"),
				endTime:   util.ParseDuration("0"),
			})
		}
	}
	for _, interval := range results {
		if len(interval) == 1 {
			at := util.ParseTime("2006", interval[0].Str())
			st.tls = append(st.tls, &timeline{
				startDate: at,
				endDate:   at,
				startTime: util.ParseDuration("0"),
				endTime:   util.ParseDuration("0"),
			})
		}
	}
	return true
}

func (st *schtasks) parseWeek(anchor cronAnchor) bool {
	st.week = true
	if anchor.hasStep() { // * * * * * 0/5 *
		st.sc = "weekly"
		st.rrange(func(tl *timeline) error {
			tl.schdule = "weekly"
			tl.mo = anchor.getStep()
			return nil
		})
	} else { // * * * * * 6#3 *
		t := anchor.times()
		st.rrange(func(tl *timeline) error {
			tl.schdule = "monthly"
			switch len(t) {
			case 2:
				tl.mo = t[0]
				tl.day = t[1].Week()
			case 1:
				tl.day = t[0].Week()
			}
			return nil
		})
	}
	return true
}

func (st *schtasks) parseMonth(anchor cronAnchor) bool {
	if anchor.hasStep() && st.week {
		return false
	}
	if st.hasSc() && anchor.hasStep() {
		tmp := []*timeline{}
		st.rrange(func(tl *timeline) error {
			step := tl.mo.Int()
			for i := tl.startDate.Year(); i < tl.endDate.Year(); i += step {
				s := tl.Clone()
				s.startDate.AddDate(-s.startDate.Year()+i, 0, 0)
				tmp = append(tmp, s)
			}
			return nil
		})
		st.tls = tmp
	}
	if anchor.hasStep() {
		st.sc = "monthly"
	}
	if anchor.hasStep() && anchor.hasRound() { // * * * * 0-10/5 *
		month := anchor.times()
		st.rrange(func(tl *timeline) error {
			tl.schdule = st.sc
			tl.mo = anchor.getStep()
			tl.startDate.AddDate(0, month[0].Int()-1, 0)
			tl.endDate.AddDate(0, month[1].Int()-1, 0)
			return nil
		})
	} else if anchor.hasStep() { // * * * * 0/5 *
		st.rrange(func(tl *timeline) error {
			tl.schdule = st.sc
			tl.mo = anchor.getStep()
			return nil
		})
	} else if anchor.hasRound() { // * * * * 0-10 *
		month := anchor.times()
		st.rrange(func(tl *timeline) error {
			tl.startDate.AddDate(0, month[0].Int()-1, 0)
			tl.endDate.AddDate(0, month[1].Int()-1, 0)
			return nil
		})
	} else { // * * * * 0,1,3
		tmp := []*timeline{}
		for _, m := range anchor.times() {
			st.rrange(func(tl *timeline) error {
				s := tl.Clone()
				s.startDate.AddDate(0, m.Int()-1, 0)
				s.endDate.AddDate(0, m.Int()-1, 0)
				tmp = append(tmp, s)
				return nil
			})
		}
		st.tls = tmp
	}
	return true
}

func (st *schtasks) parseDay(anchor cronAnchor) bool {
	st.delaySc = "monthly"
	if anchor.hasStep() && st.week {
		return false
	}
	if anchor.hasStep() && st.hasSc() {
		tmp := []*timeline{}
		st.rrange(func(tl *timeline) error {
			step := tl.mo.Int()
			for i := tl.startDate.Month(); i < tl.endDate.Month(); i += step {
				s := tl.Clone()
				s.startDate.AddDate(0, -s.startDate.Month()+i, 0)
				tmp = append(tmp, s)
			}
			return nil
		})
		st.tls = tmp
	}
	if anchor.hasStep() {
		st.sc = "daily"
	}
	if anchor.hasStep() && anchor.hasRound() {
		day := anchor.times()
		st.rrange(func(tl *timeline) error {
			tl.schdule = st.sc
			tl.mo = anchor.getStep()
			tl.startDate.AddDate(0, 0, day[0].Int()-1)
			tl.endDate.AddDate(0, 0, day[1].Int()-1)
			return nil
		})
	} else if anchor.hasStep() {
		st.rrange(func(tl *timeline) error {
			tl.schdule = st.sc
			tl.mo = anchor.getStep()
			return nil
		})
	} else if anchor.hasRound() {
		day := anchor.times()
		st.rrange(func(tl *timeline) error {
			tl.startDate.AddDate(0, 0, day[0].Int()-1)
			tl.endDate.AddDate(0, 0, day[1].Int()-1)
			return nil
		})
	} else {
		tmp := []*timeline{}
		for _, m := range anchor.times() {
			st.rrange(func(tl *timeline) error {
				s := tl.Clone()
				s.startDate.AddDate(0, 0, m.Int()-1)
				s.endDate.AddDate(0, 0, m.Int()-1)
				tmp = append(tmp, s)
				return nil
			})
		}
		st.tls = tmp
	}
	return true
}

func (st *schtasks) parseHour(anchor cronAnchor) bool {
	st.delaySc = "daily"
	if anchor.hasStep() && st.week {
		return false
	}
	if anchor.hasStep() && st.hasSc() { // * * 0-10/2 0/5
		tmp := []*timeline{}
		st.rrange(func(tl *timeline) error {
			step := tl.mo.Int()
			for i := tl.startDate.Day(); i < tl.endDate.Day(); i += step {
				s := tl.Clone()
				s.startDate.AddDate(0, 0, -s.startDate.Day()+i)
				tmp = append(tmp, s)
			}
			return nil
		})
		st.tls = tmp
	}
	if anchor.hasStep() {
		st.sc = "hourly"
	}
	if anchor.hasStep() && anchor.hasRound() {
		hour := anchor.times()
		st.rrange(func(tl *timeline) error {
			tl.schdule = st.sc
			tl.mo = anchor.getStep()
			tl.startTime.AddHour(hour[0].Float64())
			tl.endTime.AddHour(hour[1].Float64())
			return nil
		})
	} else if anchor.hasStep() {
		st.rrange(func(tl *timeline) error {
			tl.schdule = st.sc
			tl.mo = anchor.getStep()
			return nil
		})
	} else if anchor.hasRound() {
		hour := anchor.times()
		st.rrange(func(tl *timeline) error {
			tl.startTime.AddHour(hour[0].Float64())
			tl.endTime.AddHour(hour[1].Float64())
			return nil
		})
	} else {
		tmp := []*timeline{}
		for _, m := range anchor.times() {
			st.rrange(func(tl *timeline) error {
				s := tl.Clone()
				s.startTime.AddHour(m.Float64())
				tmp = append(tmp, s)
				return nil
			})
		}
		st.tls = tmp
	}
	return true
}

func (st *schtasks) parseMinute(anchor cronAnchor) bool {
	st.delaySc = "hourly"
	if anchor.hasStep() && st.week {
		return false
	}
	if anchor.hasStep() && st.hasSc() { // * 0-10/2 0/5 *
		tmp := []*timeline{}
		st.rrange(func(tl *timeline) error {
			step := tl.mo.Int()
			for i := tl.startTime.Hour(); i < tl.endTime.Hour(); i += step {
				s := tl.Clone()
				s.startTime.AddHour(float64(-tl.startTime.Hour() + i))
				tmp = append(tmp, s)
			}
			return nil
		})
		st.tls = tmp
	}
	if anchor.hasStep() {
		st.sc = "minute"
	}
	if anchor.hasStep() && anchor.hasRound() { // * 0-5/10 0/10
		minute := anchor.times()
		st.rrange(func(tl *timeline) error {
			tl.schdule = st.sc
			tl.mo = anchor.getStep()
			tl.startTime.AddMinute(minute[0].Float64())
			tl.endTime = tl.startTime.Clone()
			tl.endTime.AddMinute(minute[1].Float64())
			return nil
		})
	} else if anchor.hasStep() { // * 0/10 0
		st.rrange(func(tl *timeline) error {
			tl.schdule = st.sc
			tl.mo = anchor.getStep()
			return nil
		})
	} else if anchor.hasRound() { // * 0-5/1
		minute := anchor.times()
		st.rrange(func(tl *timeline) error {
			tl.startTime.AddMinute(minute[0].Float64())
			tl.endTime = tl.startTime.Clone()
			tl.endTime.AddMinute(minute[1].Float64())
			return nil
		})
	} else { // * 0,1,2
		tmp := []*timeline{}
		for _, m := range anchor.times() {
			st.rrange(func(tl *timeline) error {
				s := tl.Clone()
				s.startTime.AddMinute(m.Float64())
				tmp = append(tmp, s)
				return nil
			})
		}
		st.tls = tmp
	}
	return true
}

func (st *schtasks) parseSecond(anchor cronAnchor) bool {
	st.delaySc = "minute"
	return true
}

func (expr cronExpr) toSchtasks() (ret []string) {
	st := &schtasks{}
	if !expr.year.isNil() && !expr.year.hasInvalid() && !st.parseYear(expr.year) {
		return
	}
	if !expr.week.isNil() && !expr.week.hasInvalid() && !st.parseWeek(expr.week) {
		return
	}
	if !expr.mon.isNil() && !expr.mon.hasInvalid() && !st.parseMonth(expr.mon) {
		return
	}
	if !expr.day.isNil() && !expr.day.hasInvalid() && !st.parseDay(expr.day) {
		return
	}
	if !expr.hour.isNil() && !expr.hour.hasInvalid() && !st.parseHour(expr.hour) {
		return
	}
	if !expr.min.isNil() && !expr.min.hasInvalid() && !st.parseMinute(expr.min) {
		return
	}
	if !expr.sec.isNil() && !expr.sec.hasInvalid() && !st.parseSecond(expr.sec) {
		return
	}
	for _, tl := range st.tls {
		ret = append(ret, tl.String())
		if len(tl.schdule) == 0 {
			tl.schdule = st.delaySc
		}
		// fmt.Println(tl)
	}
	return
}

const (
	A_LAST    = 1
	A_STEP    = 2
	A_ROUND   = 4
	A_NO      = 8
	A_INVALID = 16
	A_ENUM    = 32
)

type cronAnchor struct {
	nums []cronTime
	attr int32
}

func (anchor cronAnchor) hasLast() bool {
	return (anchor.attr>>0)&1 == 1
}

func (anchor cronAnchor) hasInvalid() bool {
	return (anchor.attr>>4)&1 == 1
}

func (anchor cronAnchor) hasRound() bool {
	return (anchor.attr>>2)&1 == 1
}

func (anchor cronAnchor) hasNo() bool {
	return (anchor.attr>>3)&1 == 1
}

func (anchor cronAnchor) hasStep() bool {
	return (anchor.attr>>1)&1 == 1
}

func (anchor cronAnchor) getStep() cronTime {
	if anchor.hasStep() {
		return anchor.nums[0]
	}
	return -1
}

func (anchor cronAnchor) isNil() bool {
	return anchor.attr == 0 && len(anchor.nums) == 0
}

func (anchor cronAnchor) times() []cronTime {
	idx := 0
	if anchor.hasStep() {
		idx = 1
	}
	return anchor.nums[idx:]
}

type cronTime int

func (tm cronTime) Int() int {
	return int(tm)
}

func (tm cronTime) Week() string {
	switch tm {
	case 1:
		return "MON"
	case 2:
		return "TUE"
	case 3:
		return "WED"
	case 4:
		return "THU"
	case 5:
		return "FRI"
	case 6:
		return "SAT"
	case 7:
		return "SUN"
	}
	return tm.Str()
}

func (tm cronTime) Strf() string {
	switch tm {
	case 1:
		return "FIRST"
	case 2:
		return "SECOND"
	case 3:
		return "THIRD"
	case 4:
		return "FOURTH"
	case 5:
		return "FIFTH"
	case 6:
		return "SIXTH"
	case 7:
		return "SEVENTH"
	case 8:
		return "EIGHTH"
	case 9:
		return "NINTH"
	case 10:
		return "TENTH"
	case 11:
		return "ELEVENTH"
	case 12:
		return "TWELFTH"
	}
	return tm.Str()
}

func (tm cronTime) Hour() string {
	return tm.Str() + "h"
}

func (tm cronTime) Str() string {
	return strconv.Itoa(int(tm))
}

func (tm cronTime) Float32() float32 {
	return float32(tm)
}

func (tm cronTime) Float64() float64 {
	return float64(tm)
}

func (anchor cronAnchor) String() string {
	if anchor.isNil() {
		return "*"
	}
	if anchor.hasInvalid() {
		return "?"
	}
	if len(anchor.nums) == 0 {
		anchor.nums = append(anchor.nums, 0)
	}
	if anchor.hasRound() && len(anchor.nums) >= 2 {
		if anchor.hasStep() {
			return fmt.Sprintf("%d-%d/%d", anchor.nums[1], anchor.nums[2], anchor.nums[0])
		}
		return anchor.nums[0].Str() + "-" + anchor.nums[1].Str()
	}
	if anchor.hasNo() {
		return anchor.nums[0].Str() + "#" + anchor.nums[1].Str()
	}
	if anchor.hasLast() {
		ret := "L"
		if v := anchor.nums[0]; v != 0 {
			ret = v.Str() + "L"
		}
		if len(anchor.nums) > 1 {
			ret += "-" + anchor.nums[1].Str()
		}
		return ret
	}
	if anchor.hasStep() {
		if len(anchor.nums) <= 1 {
			if v := anchor.nums[0]; v == 0 {
				return "0"
			} else {
				return fmt.Sprintf("0/%d", v)
			}
		}
		return anchor.nums[1].Str() + "/" + anchor.nums[0].Str()
	}
	str := []string{}
	for _, num := range anchor.nums {
		str = append(str, num.Str())
	}
	return strings.Join(str, ",")
}

var week = map[string]int{
	"MON": 1,
	"TUE": 2,
	"WED": 3,
	"THU": 4,
	"FRI": 5,
	"SAT": 6,
	"SUN": 7,
}

func CronParse(str string) (expr cronExpr, command string) {
	anchors := strings.Fields(str)
	re := regexp.MustCompile(`(\d+-\d+/\d+|\*|\?|L|\d+|\d+\#\d+|(\dL|L)\-\d+|MON|TUE|WED|THU|FRI|SAT|SUN)`)
	res := []cronAnchor{}
	for _, anchor := range anchors {
		if !re.MatchString(anchor) {
			command += anchor + " "
			continue
		}
		if anchor == "*" {
			res = append(res, cronAnchor{})
			continue
		}
		if anchor == "?" {
			res = append(res, cronAnchor{attr: A_INVALID})
			continue
		}
		if strings.Contains(anchor, ",") {
			nums := strings.Split(anchor, ",")
			ca := cronAnchor{attr: A_ENUM}
			for _, num := range nums {
				i, _ := strconv.Atoi(num)
				ca.nums = append(ca.nums, cronTime(i))
			}
			res = append(res, ca)
			continue
		}
		if strings.Contains(anchor, "/") {
			if strings.Contains(anchor, "-") {
				if a, b, ok := strings.Cut(anchor, "-"); ok {
					left, _ := strconv.Atoi(a)
					if aa, bb, ok := strings.Cut(b, "/"); ok {
						right, _ := strconv.Atoi(aa)
						step, _ := strconv.Atoi(bb)
						res = append(res, cronAnchor{
							nums: []cronTime{
								cronTime(step),
								cronTime(left),
								cronTime(right)},
							attr: A_STEP | A_ROUND,
						})
					}
				}
			} else {
				if a, b, ok := strings.Cut(anchor, "/"); ok {
					left, _ := strconv.Atoi(a)
					right, _ := strconv.Atoi(b)
					res = append(res, cronAnchor{nums: []cronTime{
						cronTime(right),
						cronTime(left),
					}, attr: A_STEP})
				}
			}
		} else if strings.Contains(anchor, "-") {
			if a, b, ok := strings.Cut(anchor, "-"); ok {
				left, _ := strconv.Atoi(a)
				right, _ := strconv.Atoi(b)
				res = append(res, cronAnchor{nums: []cronTime{
					cronTime(left),
					cronTime(right)},
					attr: A_ROUND})
			}
		} else if strings.Contains(anchor, "L") {
			idx := strings.Index(anchor, "L")
			left, _ := strconv.Atoi(anchor[:idx])
			// right := anchor[idx+1:]
			// if n := strings.Index(right, "-"); n != -1 {

			// }
			res = append(res, cronAnchor{nums: []cronTime{cronTime(left)}, attr: A_LAST})
		} else if strings.Contains(anchor, "#") {
			kv := strings.Split(anchor, "#")
			key, _ := strconv.Atoi(kv[0])
			value, _ := strconv.Atoi(kv[1])
			if len(kv) == 2 {
				res = append(res, cronAnchor{nums: []cronTime{
					cronTime(key),
					cronTime(value),
				}, attr: A_NO})
			}
		} else if week[anchor] > 0 {
			res = append(res, cronAnchor{nums: []cronTime{cronTime(week[anchor])}, attr: A_ENUM})
		} else {
			i, _ := strconv.Atoi(anchor)
			res = append(res, cronAnchor{nums: []cronTime{cronTime(i)}, attr: A_ENUM})
		}
	}
	switch n := len(res); n {
	case 5:
		expr.min = res[0]
		expr.hour = res[1]
		expr.day = res[2]
		expr.mon = res[3]
		expr.week = res[4]
	case 7:
		expr.year = res[6]
		fallthrough
	case 6:
		expr.sec = res[0]
		expr.min = res[1]
		expr.hour = res[2]
		expr.day = res[3]
		expr.mon = res[4]
		expr.week = res[5]
	default:
	}
	return
}
