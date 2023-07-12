// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// FmtTimestamp return a time string whom format is 2006-01-02
func FmtTimestamp(ts int64) time.Time {
	return time.Unix(ts, 0)
}

// NowTimestamp returm current timestamp
func NowTimestamp() int64 {
	return time.Now().Unix()
}

// NowTimestamp returm current timestamp string
func NowTimestampByString() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

type Duration struct {
	td time.Duration
}

func ParseDuration(s string) *Duration {
	if strings.Count(s, ":") == 2 {
		str := strings.Split(s, ":")
		if len(str) != 3 {
			return &Duration{}
		}
		s = fmt.Sprintf("%sh%sm%ss", str[0], str[1], str[2])
	}
	d, _ := time.ParseDuration(s)
	return &Duration{td: d}
}

func (d *Duration) Hour() int {
	return int(d.td.Hours())
}

func (d *Duration) AddHour(i float64) {
	d.td = time.Duration((d.td.Hours() + i) * float64(time.Hour))
}

func (d *Duration) SubHour(i float64) {
	d.td = time.Duration((d.td.Hours() - i) * float64(time.Hour))
}

func (d *Duration) AddMinute(i float64) {
	d.td = time.Duration((d.td.Minutes() + i) * float64(time.Minute))
}

func (d *Duration) SubMinute(i float64) {
	d.td = time.Duration((d.td.Minutes() - i) * float64(time.Minute))
}

func (d *Duration) AddSecond(i float64) {
	d.td = time.Duration((d.td.Seconds() + i) * float64(time.Second))
}

func (d *Duration) SubSecond(i float64) {
	d.td = time.Duration((d.td.Seconds() - i) * float64(time.Second))
}

func (d *Duration) Format(sep string) string {
	if len(sep) == 0 {
		return d.String()
	}

	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(d.td.String(), -1)
	str := make([]string, 3)
	for i := 0; i < 3; i++ {
		str[i] = "00"
	}
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return fmt.Sprintf(`00%s00%s00`, sep, sep)
		}
		numstr := strconv.Itoa(num)
		if len(numstr) == 1 {
			numstr = "0" + numstr
		}
		str = append(str, numstr)
	}
	return strings.Join(str[len(str)-3:], sep)
}

func (d *Duration) Interval() string {
	seconds := d.td.Seconds()
	minutes := seconds / 60
	hours := minutes / 60
	days := hours / 24
	months := days / 30
	if months > 0 {
		return fmt.Sprintf("%f month ago", months)
	} else if days > 0 {
		return fmt.Sprintf("%f day ago", days)
	} else if hours > 0 {
		return fmt.Sprintf("%f hour ago", hours)
	} else if minutes > 0 {
		return fmt.Sprintf("%f minute ago", minutes)
	} else {
		return "just now"
	}
}

func (d *Duration) Clone() *Duration {
	return &Duration{td: d.td}
}

func (d *Duration) String() string {
	return d.td.String()
}

type Time struct {
	t      time.Time
	layout string
}

func ParseTime(layout, value string) *Time {
	t, _ := time.Parse(layout, value)
	return &Time{t: t, layout: layout}
}

func (tm *Time) Year() int {
	return tm.t.Year()
}

func (tm *Time) Month() int {
	return int(tm.t.Month())
}

func (tm *Time) Day() int {
	return tm.t.Day()
}

func (tm *Time) AddDate(years, months, days int) {
	tm.t = tm.t.AddDate(years, months, days)
}

func (tm *Time) Diff(t *Time) *Duration {
	return &Duration{td: tm.t.Sub(t.t)}
}

func (tm *Time) DiffNow() *Duration {
	return tm.Diff(&Time{t: time.Now()})
}

func (tm *Time) Format(layout string) string {
	return tm.t.Format(layout)
}

func (tm *Time) String() string {
	// if len(tm.layout) != 0 {
	// 	return tm.Format(tm.layout)
	// }
	return tm.t.String()
}

func (tm Time) Clone() *Time {
	return &Time{t: tm.t, layout: tm.layout}
}
