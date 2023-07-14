// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package timelib

import (
	"time"

	yocki "github.com/ansurfen/yock/interface"
)

func LoadTime(yocks yocki.YockScheduler) {
	lib := yocks.CreateLib("time")
	lib.SetField(map[string]any{
		// functions
		"LoadLocation":           time.LoadLocation,
		"FixedZone":              time.FixedZone,
		"LoadLocationFromTZData": time.LoadLocationFromTZData,
		"Parse":                  time.Parse,
		"AfterFunc":              time.AfterFunc,
		"UnixMicro":              time.UnixMicro,
		"Since":                  time.Since,
		"Now":                    time.Now,
		"ParseDuration":          time.ParseDuration,
		"ParseInLocation":        time.ParseInLocation,
		"Date":                   time.Date,
		"After":                  time.After,
		"Unix":                   time.Unix,
		"Until":                  time.Until,
		"Sleep":                  time.Sleep,
		"NewTimer":               time.NewTimer,
		"NewTicker":              time.NewTicker,
		"Tick":                   time.Tick,
		"UnixMilli":              time.UnixMilli,
		// constants
		"Layout":      time.Layout,
		"ANSIC":       time.ANSIC,
		"UnixDate":    time.UnixDate,
		"RubyDate":    time.RubyDate,
		"RFC822":      time.RFC822,
		"RFC822Z":     time.RFC822Z,
		"RFC850":      time.RFC850,
		"RFC1123":     time.RFC1123,
		"RFC1123Z":    time.RFC1123Z,
		"RFC3339":     time.RFC3339,
		"RFC3339Nano": time.RFC3339Nano,
		"Kitchen":     time.Kitchen,
		"Stamp":       time.Stamp,
		"StampMilli":  time.StampMilli,
		"StampMicro":  time.StampMicro,
		"StampNano":   time.StampNano,
		"DateTime":    time.DateTime,
		"DateOnly":    time.DateOnly,
		"TimeOnly":    time.TimeOnly,
		"January":     time.January,
		"February":    time.February,
		"March":       time.March,
		"April":       time.April,
		"May":         time.May,
		"June":        time.June,
		"July":        time.July,
		"August":      time.August,
		"September":   time.September,
		"October":     time.October,
		"November":    time.November,
		"December":    time.December,
		"Sunday":      time.Sunday,
		"Monday":      time.Monday,
		"Tuesday":     time.Tuesday,
		"Wednesday":   time.Wednesday,
		"Thursday":    time.Thursday,
		"Friday":      time.Friday,
		"Saturday":    time.Saturday,
		"Nanosecond":  time.Nanosecond,
		"Microsecond": time.Microsecond,
		"Millisecond": time.Millisecond,
		"Second":      time.Second,
		"Minute":      time.Minute,
		"Hour":        time.Hour,
		// variable
		"UTC":   time.UTC,
		"Local": time.Local,
	})
}
