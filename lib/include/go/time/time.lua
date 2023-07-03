-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class time
---@field Layout any
---@field ANSIC any
---@field UnixDate any
---@field RubyDate any
---@field RFC822 any
---@field RFC822Z any
---@field RFC850 any
---@field RFC1123 any
---@field RFC1123Z any
---@field RFC3339 any
---@field RFC3339Nano any
---@field Kitchen any
---@field Stamp any
---@field StampMilli any
---@field StampMicro any
---@field StampNano any
---@field DateTime any
---@field DateOnly any
---@field TimeOnly any
---@field January any
---@field February any
---@field March any
---@field April any
---@field May any
---@field June any
---@field July any
---@field August any
---@field September any
---@field October any
---@field November any
---@field December any
---@field Sunday any
---@field Monday any
---@field Tuesday any
---@field Wednesday any
---@field Thursday any
---@field Friday any
---@field Saturday any
---@field Nanosecond any
---@field Microsecond any
---@field Millisecond any
---@field Second any
---@field Minute any
---@field Hour any
---@field UTC any
---@field Local any
time = {}

---{{.timeParseInLocation}}
---@param layout string
---@param value string
---@param loc timeLocation
---@return timeTime, err
function time.ParseInLocation(layout, value, loc)
end

---{{.timeDate}}
---@param year number
---@param month timeMonth
---@param day number
---@param hour number
---@param min number
---@param sec number
---@param nsec number
---@param loc timeLocation
---@return timeTime
function time.Date(year, month, day, hour, min, sec, nsec, loc)
end

---{{.timeParseDuration}}
---@param s string
---@return timeDuration, err
function time.ParseDuration(s)
end

---{{.timeUnix}}
---@param sec number
---@param nsec number
---@return timeTime
function time.Unix(sec, nsec)
end

---{{.timeUntil}}
---@param t timeTime
---@return timeDuration
function time.Until(t)
end

---{{.timeAfter}}
---@param d timeDuration
---@return any
function time.After(d)
end

---{{.timeNewTimer}}
---@param d timeDuration
---@return timeTimer
function time.NewTimer(d)
end

---{{.timeNewTicker}}
---@param d timeDuration
---@return timeTicker
function time.NewTicker(d)
end

---{{.timeTick}}
---@param d timeDuration
---@return any
function time.Tick(d)
end

---{{.timeUnixMilli}}
---@param msec number
---@return timeTime
function time.UnixMilli(msec)
end

---{{.timeSleep}}
---@param d timeDuration
function time.Sleep(d)
end

---{{.timeAfterFunc}}
---@param d timeDuration
---@param f function
---@return timeTimer
function time.AfterFunc(d, f)
end

---{{.timeUnixMicro}}
---@param usec number
---@return timeTime
function time.UnixMicro(usec)
end

---{{.timeSince}}
---@param t timeTime
---@return timeDuration
function time.Since(t)
end

---{{.timeNow}}
---@return timeTime
function time.Now()
end

---{{.timeLoadLocation}}
---@param name string
---@return timeLocation, err
function time.LoadLocation(name)
end

---{{.timeFixedZone}}
---@param name string
---@param offset number
---@return timeLocation
function time.FixedZone(name, offset)
end

---{{.timeLoadLocationFromTZData}}
---@param name string
---@param data byte[]
---@return timeLocation, err
function time.LoadLocationFromTZData(name, data)
end

---{{.timeParse}}
---@param layout string
---@param value string
---@return timeTime, err
function time.Parse(layout, value)
end

---@class timeTimer
---@field C any
local timeTimer = {}

---{{.timeTimerReset}}
---@param d timeDuration
---@return boolean
function timeTimer:Reset(d)
end

---{{.timeTimerStop}}
---@return boolean
function timeTimer:Stop()
end

---@class timeTicker
---@field C any
local timeTicker = {}

---{{.timeTickerStop}}
function timeTicker:Stop()
end

---{{.timeTickerReset}}
---@param d timeDuration
function timeTicker:Reset(d)
end

---@class timeMonth
local timeMonth = {}

---{{.timeMonthString}}
---@return string
function timeMonth:String()
end

---@class timeWeekday
local timeWeekday = {}

---{{.timeWeekdayString}}
---@return string
function timeWeekday:String()
end

---@class timeLocation
local timeLocation = {}

---{{.timeLocationString}}
---@return string
function timeLocation:String()
end

---@class timeTime
local timeTime = {}

---{{.timeTimeMinute}}
---@return number
function timeTime:Minute()
end

---{{.timeTimeYearDay}}
---@return number
function timeTime:YearDay()
end

---{{.timeTimeIsDST}}
---@return boolean
function timeTime:IsDST()
end

---{{.timeTimeEqual}}
---@param u timeTime
---@return boolean
function timeTime:Equal(u)
end

---{{.timeTimeCompare}}
---@param u timeTime
---@return number
function timeTime:Compare(u)
end

---{{.timeTimeDate}}
---@return number, timeMonth, number
function timeTime:Date()
end

---{{.timeTimeSub}}
---@param u timeTime
---@return timeDuration
function timeTime:Sub(u)
end

---{{.timeTimeUTC}}
---@return timeTime
function timeTime:UTC()
end

---{{.timeTimeUnmarshalBinary}}
---@param data byte[]
---@return err
function timeTime:UnmarshalBinary(data)
end

---{{.timeTimeIsZero}}
---@return boolean
function timeTime:IsZero()
end

---{{.timeTimeDay}}
---@return number
function timeTime:Day()
end

---{{.timeTimeClock}}
---@return number
function timeTime:Clock()
end

---{{.timeTimeSecond}}
---@return number
function timeTime:Second()
end

---{{.timeTimeLocal}}
---@return timeTime
function timeTime:Local()
end

---{{.timeTimeIn}}
---@param loc timeLocation
---@return timeTime
function timeTime:In(loc)
end

---{{.timeTimeUnixMicro}}
---@return number
function timeTime:UnixMicro()
end

---{{.timeTimeGobDecode}}
---@param data byte[]
---@return err
function timeTime:GobDecode(data)
end

---{{.timeTimeRound}}
---@param d timeDuration
---@return timeTime
function timeTime:Round(d)
end

---{{.timeTimeZoneBounds}}
---@return timeTime
function timeTime:ZoneBounds()
end

---{{.timeTimeUnixMilli}}
---@return number
function timeTime:UnixMilli()
end

---{{.timeTimeBefore}}
---@param u timeTime
---@return boolean
function timeTime:Before(u)
end

---{{.timeTimeMonth}}
---@return timeMonth
function timeTime:Month()
end

---{{.timeTimeISOWeek}}
---@return number
function timeTime:ISOWeek()
end

---{{.timeTimeUnmarshalText}}
---@param data byte[]
---@return err
function timeTime:UnmarshalText(data)
end

---{{.timeTimeTruncate}}
---@param d timeDuration
---@return timeTime
function timeTime:Truncate(d)
end

---{{.timeTimeAfter}}
---@param u timeTime
---@return boolean
function timeTime:After(u)
end

---{{.timeTimeZone}}
---@return string, number
function timeTime:Zone()
end

---{{.timeTimeUnix}}
---@return number
function timeTime:Unix()
end

---{{.timeTimeAdd}}
---@param d timeDuration
---@return timeTime
function timeTime:Add(d)
end

---{{.timeTimeLocation}}
---@return timeLocation
function timeTime:Location()
end

---{{.timeTimeMarshalJSON}}
---@return byte[], err
function timeTime:MarshalJSON()
end

---{{.timeTimeUnmarshalJSON}}
---@param data byte[]
---@return err
function timeTime:UnmarshalJSON(data)
end

---{{.timeTimeNanosecond}}
---@return number
function timeTime:Nanosecond()
end

---{{.timeTimeUnixNano}}
---@return number
function timeTime:UnixNano()
end

---{{.timeTimeMarshalBinary}}
---@return byte[], err
function timeTime:MarshalBinary()
end

---{{.timeTimeGobEncode}}
---@return byte[], err
function timeTime:GobEncode()
end

---{{.timeTimeYear}}
---@return number
function timeTime:Year()
end

---{{.timeTimeAddDate}}
---@param years number
---@param months number
---@param days number
---@return timeTime
function timeTime:AddDate(years, months, days)
end

---{{.timeTimeHour}}
---@return number
function timeTime:Hour()
end

---{{.timeTimeMarshalText}}
---@return byte[], err
function timeTime:MarshalText()
end

---{{.timeTimeWeekday}}
---@return timeWeekday
function timeTime:Weekday()
end

---@class timeParseError
---@field Layout string
---@field Value string
---@field LayoutElem string
---@field ValueElem string
---@field Message string
local timeParseError = {}

---{{.timeParseErrorError}}
---@return string
function timeParseError:Error()
end

---@class timeMapZone
---@field Other string
---@field Territory string
---@field Type string
local timeMapZone = {}

---@class timeSupplementalData
---@field Zones any
local timeSupplementalData = {}

---@class timeDuration
local timeDuration = {}

---{{.timeDurationString}}
---@return string
function timeDuration:String()
end

---{{.timeDurationNanoseconds}}
---@return number
function timeDuration:Nanoseconds()
end

---{{.timeDurationMilliseconds}}
---@return number
function timeDuration:Milliseconds()
end

---{{.timeDurationSeconds}}
---@return number
function timeDuration:Seconds()
end

---{{.timeDurationTruncate}}
---@param m timeDuration
---@return timeDuration
function timeDuration:Truncate(m)
end

---{{.timeDurationMicroseconds}}
---@return number
function timeDuration:Microseconds()
end

---{{.timeDurationMinutes}}
---@return number
function timeDuration:Minutes()
end

---{{.timeDurationHours}}
---@return number
function timeDuration:Hours()
end

---{{.timeDurationRound}}
---@param m timeDuration
---@return timeDuration
function timeDuration:Round(m)
end

---{{.timeDurationAbs}}
---@return timeDuration
function timeDuration:Abs()
end
