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

--- ParseInLocation is like Parse but differs in two important ways.
--- First, in the absence of time zone information, Parse interprets a time as UTC;
--- ParseInLocation interprets the time as in the given location.
--- Second, when given a zone offset or abbreviation, Parse tries to match it
--- against the Local location; ParseInLocation uses the given location.
---@param layout string
---@param value string
---@param loc timeLocation
---@return timeTime, err
function time.ParseInLocation(layout, value, loc) end

--- Sleep pauses the current goroutine for at least the duration d.
--- A negative or zero duration causes Sleep to return immediately.
---@param d timeDuration
function time.Sleep(d) end

--- NewTicker returns a new Ticker containing a channel that will send
--- the current time on the channel after each tick. The period of the
--- ticks is specified by the duration argument. The ticker will adjust
--- the time interval or drop ticks to make up for slow receivers.
--- The duration d must be greater than zero; if not, NewTicker will
--- panic. Stop the ticker to release associated resources.
---@param d timeDuration
---@return timeTicker
function time.NewTicker(d) end

--- Until returns the duration until t.
--- It is shorthand for t.Sub(time.Now()).
---@param t timeTime
---@return timeDuration
function time.Until(t) end

--- Unix returns the local Time corresponding to the given Unix time,
--- sec seconds and nsec nanoseconds since January 1, 1970 UTC.
--- It is valid to pass nsec outside the range [0, 999999999].
--- Not all sec values have a corresponding time value. One such
--- value is 1<<63-1 (the largest int64 value).
---@param sec number
---@param nsec number
---@return timeTime
function time.Unix(sec, nsec) end

--- Date returns the Time corresponding to
---
---	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
---
--- in the appropriate zone for that time in the given location.
---
--- The month, day, hour, min, sec, and nsec values may be outside
--- their usual ranges and will be normalized during the conversion.
--- For example, October 32 converts to November 1.
---
--- A daylight savings time transition skips or repeats times.
--- For example, in the United States, March 13, 2011 2:15am never occurred,
--- while November 6, 2011 1:15am occurred twice. In such cases, the
--- choice of time zone, and therefore the time, is not well-defined.
--- Date returns a time that is correct in one of the two zones involved
--- in the transition, but it does not guarantee which.
---
--- Date panics if loc is nil.
---@param year number
---@param month timeMonth
---@param day number
---@param hour number
---@param min number
---@param sec number
---@param nsec number
---@param loc timeLocation
---@return timeTime
function time.Date(year, month, day, hour, min, sec, nsec, loc) end

--- Parse parses a formatted string and returns the time value it represents.
--- See the documentation for the constant called Layout to see how to
--- represent the format. The second argument must be parseable using
--- the format string (layout) provided as the first argument.
---
--- The example for Time.Format demonstrates the working of the layout string
--- in detail and is a good reference.
---
--- When parsing (only), the input may contain a fractional second
--- field immediately after the seconds field, even if the layout does not
--- signify its presence. In that case either a comma or a decimal point
--- followed by a maximal series of digits is parsed as a fractional second.
--- Fractional seconds are truncated to nanosecond precision.
---
--- Elements omitted from the layout are assumed to be zero or, when
--- zero is impossible, one, so parsing "3:04pm" returns the time
--- corresponding to Jan 1, year 0, 15:04:00 UTC (note that because the year is
--- 0, this time is before the zero Time).
--- Years must be in the range 0000..9999. The day of the week is checked
--- for syntax but it is otherwise ignored.
---
--- For layouts specifying the two-digit year 06, a value NN >= 69 will be treated
--- as 19NN and a value NN < 69 will be treated as 20NN.
---
--- The remainder of this comment describes the handling of time zones.
---
--- In the absence of a time zone indicator, Parse returns a time in UTC.
---
--- When parsing a time with a zone offset like -0700, if the offset corresponds
--- to a time zone used by the current location (Local), then Parse uses that
--- location and zone in the returned time. Otherwise it records the time as
--- being in a fabricated location with time fixed at the given zone offset.
---
--- When parsing a time with a zone abbreviation like MST, if the zone abbreviation
--- has a defined offset in the current location, then that offset is used.
--- The zone abbreviation "UTC" is recognized as UTC regardless of location.
--- If the zone abbreviation is unknown, Parse records the time as being
--- in a fabricated location with the given zone abbreviation and a zero offset.
--- This choice means that such a time can be parsed and reformatted with the
--- same layout losslessly, but the exact instant used in the representation will
--- differ by the actual zone offset. To avoid such problems, prefer time layouts
--- that use a numeric zone offset, or use ParseInLocation.
---@param layout string
---@param value string
---@return timeTime, err
function time.Parse(layout, value) end

--- AfterFunc waits for the duration to elapse and then calls f
--- in its own goroutine. It returns a Timer that can
--- be used to cancel the call using its Stop method.
---@param d timeDuration
---@param f function
---@return timeTimer
function time.AfterFunc(d, f) end

--- Tick is a convenience wrapper for NewTicker providing access to the ticking
--- channel only. While Tick is useful for clients that have no need to shut down
--- the Ticker, be aware that without a way to shut it down the underlying
--- Ticker cannot be recovered by the garbage collector; it "leaks".
--- Unlike NewTicker, Tick will return nil if d <= 0.
---@param d timeDuration
---@return any
function time.Tick(d) end

--- UnixMilli returns the local Time corresponding to the given Unix time,
--- msec milliseconds since January 1, 1970 UTC.
---@param msec number
---@return timeTime
function time.UnixMilli(msec) end

--- Now returns the current local time.
---@return timeTime
function time.Now() end

--- LoadLocation returns the Location with the given name.
---
--- If the name is "" or "UTC", LoadLocation returns UTC.
--- If the name is "Local", LoadLocation returns Local.
---
--- Otherwise, the name is taken to be a location name corresponding to a file
--- in the IANA Time Zone database, such as "America/New_York".
---
--- LoadLocation looks for the IANA Time Zone database in the following
--- locations in order:
---
---   - the directory or uncompressed zip file named by the ZONEINFO environment variable
---   - on a Unix system, the system standard installation location
---   - $GOROOT/lib/time/zoneinfo.zip
---   - the time/tzdata package, if it was imported
---@param name string
---@return timeLocation, err
function time.LoadLocation(name) end

--- NewTimer creates a new Timer that will send
--- the current time on its channel after at least duration d.
---@param d timeDuration
---@return timeTimer
function time.NewTimer(d) end

--- UnixMicro returns the local Time corresponding to the given Unix time,
--- usec microseconds since January 1, 1970 UTC.
---@param usec number
---@return timeTime
function time.UnixMicro(usec) end

--- FixedZone returns a Location that always uses
--- the given zone name and offset (seconds east of UTC).
---@param name string
---@param offset number
---@return timeLocation
function time.FixedZone(name, offset) end

--- LoadLocationFromTZData returns a Location with the given name
--- initialized from the IANA Time Zone database-formatted data.
--- The data should be in the format of a standard IANA time zone file
--- (for example, the content of /etc/localtime on Unix systems).
---@param name string
---@param data byte[]
---@return timeLocation, err
function time.LoadLocationFromTZData(name, data) end

--- Since returns the time elapsed since t.
--- It is shorthand for time.Now().Sub(t).
---@param t timeTime
---@return timeDuration
function time.Since(t) end

--- After waits for the duration to elapse and then sends the current time
--- on the returned channel.
--- It is equivalent to NewTimer(d).C.
--- The underlying Timer is not recovered by the garbage collector
--- until the timer fires. If efficiency is a concern, use NewTimer
--- instead and call Timer.Stop if the timer is no longer needed.
---@param d timeDuration
---@return any
function time.After(d) end

--- ParseDuration parses a duration string.
--- A duration string is a possibly signed sequence of
--- decimal numbers, each with optional fraction and a unit suffix,
--- such as "300ms", "-1.5h" or "2h45m".
--- Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
---@param s string
---@return timeDuration, err
function time.ParseDuration(s) end

--- ParseError describes a problem parsing a time string.
---@class netParseError
---@field Layout string
---@field Value string
---@field LayoutElem string
---@field ValueElem string
---@field Message string
local netParseError = {}

--- Error returns the string representation of a ParseError.
---@return string
function netParseError:Error() end


---@class timeMapZone
---@field Other string
---@field Territory string
---@field Type string
local timeMapZone = {}


---@class timeSupplementalData
---@field Zones any
local timeSupplementalData = {}

--- The Timer type represents a single event.
--- When the Timer expires, the current time will be sent on C,
--- unless the Timer was created by AfterFunc.
--- A Timer must be created with NewTimer or AfterFunc.
---@class timeTimer
---@field C any
local timeTimer = {}

--- Stop prevents the Timer from firing.
--- It returns true if the call stops the timer, false if the timer has already
--- expired or been stopped.
--- Stop does not close the channel, to prevent a read from the channel succeeding
--- incorrectly.
---
--- To ensure the channel is empty after a call to Stop, check the
--- return value and drain the channel.
--- For example, assuming the program has not received from t.C already:
---
---	if !t.Stop() {
---		<-t.C
---	}
---
--- This cannot be done concurrent to other receives from the Timer's
--- channel or other calls to the Timer's Stop method.
---
--- For a timer created with AfterFunc(d, f), if t.Stop returns false, then the timer
--- has already expired and the function f has been started in its own goroutine;
--- Stop does not wait for f to complete before returning.
--- If the caller needs to know whether f is completed, it must coordinate
--- with f explicitly.
---@return boolean
function timeTimer:Stop() end

--- Reset changes the timer to expire after duration d.
--- It returns true if the timer had been active, false if the timer had
--- expired or been stopped.
---
--- For a Timer created with NewTimer, Reset should be invoked only on
--- stopped or expired timers with drained channels.
---
--- If a program has already received a value from t.C, the timer is known
--- to have expired and the channel drained, so t.Reset can be used directly.
--- If a program has not yet received a value from t.C, however,
--- the timer must be stopped and—if Stop reports that the timer expired
--- before being stopped—the channel explicitly drained:
---
---	if !t.Stop() {
---		<-t.C
---	}
---	t.Reset(d)
---
--- This should not be done concurrent to other receives from the Timer's
--- channel.
---
--- Note that it is not possible to use Reset's return value correctly, as there
--- is a race condition between draining the channel and the new timer expiring.
--- Reset should always be invoked on stopped or expired channels, as described above.
--- The return value exists to preserve compatibility with existing programs.
---
--- For a Timer created with AfterFunc(d, f), Reset either reschedules
--- when f will run, in which case Reset returns true, or schedules f
--- to run again, in which case it returns false.
--- When Reset returns false, Reset neither waits for the prior f to
--- complete before returning nor does it guarantee that the subsequent
--- goroutine running f does not run concurrently with the prior
--- one. If the caller needs to know whether the prior execution of
--- f is completed, it must coordinate with f explicitly.
---@param d timeDuration
---@return boolean
function timeTimer:Reset(d) end

--- A Ticker holds a channel that delivers “ticks” of a clock
--- at intervals.
---@class timeTicker
---@field C any
local timeTicker = {}

--- Stop turns off a ticker. After Stop, no more ticks will be sent.
--- Stop does not close the channel, to prevent a concurrent goroutine
--- reading from the channel from seeing an erroneous "tick".
function timeTicker:Stop() end

--- Reset stops a ticker and resets its period to the specified duration.
--- The next tick will arrive after the new period elapses. The duration d
--- must be greater than zero; if not, Reset will panic.
---@param d timeDuration
function timeTicker:Reset(d) end

--- A Month specifies a month of the year (January = 1, ...).
---@class timeMonth
local timeMonth = {}

--- String returns the English name of the month ("January", "February", ...).
---@return string
function timeMonth:String() end

--- A Weekday specifies a day of the week (Sunday = 0, ...).
---@class timeWeekday
local timeWeekday = {}

--- String returns the English name of the day ("Sunday", "Monday", ...).
---@return string
function timeWeekday:String() end

--- A Time represents an instant in time with nanosecond precision.
---
--- Programs using times should typically store and pass them as values,
--- not pointers. That is, time variables and struct fields should be of
--- type time.Time, not *time.Time.
---
--- A Time value can be used by multiple goroutines simultaneously except
--- that the methods GobDecode, UnmarshalBinary, UnmarshalJSON and
--- UnmarshalText are not concurrency-safe.
---
--- Time instants can be compared using the Before, After, and Equal methods.
--- The Sub method subtracts two instants, producing a Duration.
--- The Add method adds a Time and a Duration, producing a Time.
---
--- The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
--- As this time is unlikely to come up in practice, the IsZero method gives
--- a simple way of detecting a time that has not been initialized explicitly.
---
--- Each Time has associated with it a Location, consulted when computing the
--- presentation form of the time, such as in the Format, Hour, and Year methods.
--- The methods Local, UTC, and In return a Time with a specific location.
--- Changing the location in this way changes only the presentation; it does not
--- change the instant in time being denoted and therefore does not affect the
--- computations described in earlier paragraphs.
---
--- Representations of a Time value saved by the GobEncode, MarshalBinary,
--- MarshalJSON, and MarshalText methods store the Time.Location's offset, but not
--- the location name. They therefore lose information about Daylight Saving Time.
---
--- In addition to the required “wall clock” reading, a Time may contain an optional
--- reading of the current process's monotonic clock, to provide additional precision
--- for comparison or subtraction.
--- See the “Monotonic Clocks” section in the package documentation for details.
---
--- Note that the Go == operator compares not just the time instant but also the
--- Location and the monotonic clock reading. Therefore, Time values should not
--- be used as map or database keys without first guaranteeing that the
--- identical Location has been set for all values, which can be achieved
--- through use of the UTC or Local method, and that the monotonic clock reading
--- has been stripped by setting t = t.Round(0). In general, prefer t.Equal(u)
--- to t == u, since t.Equal uses the most accurate comparison available and
--- correctly handles the case when only one of its arguments has a monotonic
--- clock reading.
---@class timeTime
local timeTime = {}

--- Before reports whether the time instant t is before u.
---@param u timeTime
---@return boolean
function timeTime:Before(u) end

--- IsZero reports whether t represents the zero time instant,
--- January 1, year 1, 00:00:00 UTC.
---@return boolean
function timeTime:IsZero() end

--- Local returns t with the location set to local time.
---@return timeTime
function timeTime:Local() end

--- UnixMilli returns t as a Unix time, the number of milliseconds elapsed since
--- January 1, 1970 UTC. The result is undefined if the Unix time in
--- milliseconds cannot be represented by an int64 (a date more than 292 million
--- years before or after 1970). The result does not depend on the
--- location associated with t.
---@return number
function timeTime:UnixMilli() end

--- Unix returns t as a Unix time, the number of seconds elapsed
--- since January 1, 1970 UTC. The result does not depend on the
--- location associated with t.
--- Unix-like operating systems often record time as a 32-bit
--- count of seconds, but since the method here returns a 64-bit
--- value it is valid for billions of years into the past or future.
---@return number
function timeTime:Unix() end

--- Year returns the year in which t occurs.
---@return number
function timeTime:Year() end

--- Clock returns the hour, minute, and second within the day specified by t.
---@return number
function timeTime:Clock() end

--- Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
--- value that can be stored in a Duration, the maximum (or minimum) duration
--- will be returned.
--- To compute t-d for a duration d, use t.Add(-d).
---@param u timeTime
---@return timeDuration
function timeTime:Sub(u) end

--- UTC returns t with the location set to UTC.
---@return timeTime
function timeTime:UTC() end

--- Location returns the time zone information associated with t.
---@return timeLocation
function timeTime:Location() end

--- Minute returns the minute offset within the hour specified by t, in the range [0, 59].
---@return number
function timeTime:Minute() end

--- In returns a copy of t representing the same time instant, but
--- with the copy's location information set to loc for display
--- purposes.
---
--- In panics if loc is nil.
---@param loc timeLocation
---@return timeTime
function timeTime:In(loc) end

--- Zone computes the time zone in effect at time t, returning the abbreviated
--- name of the zone (such as "CET") and its offset in seconds east of UTC.
---@return string, number
function timeTime:Zone() end

--- MarshalBinary implements the encoding.BinaryMarshaler interface.
---@return byte[], err
function timeTime:MarshalBinary() end

--- Round returns the result of rounding t to the nearest multiple of d (since the zero time).
--- The rounding behavior for halfway values is to round up.
--- If d <= 0, Round returns t stripped of any monotonic clock reading but otherwise unchanged.
---
--- Round operates on the time as an absolute duration since the
--- zero time; it does not operate on the presentation form of the
--- time. Thus, Round(Hour) may return a time with a non-zero
--- minute, depending on the time's Location.
---@param d timeDuration
---@return timeTime
function timeTime:Round(d) end

--- Second returns the second offset within the minute specified by t, in the range [0, 59].
---@return number
function timeTime:Second() end

--- Add returns the time t+d.
---@param d timeDuration
---@return timeTime
function timeTime:Add(d) end

--- Compare compares the time instant t with u. If t is before u, it returns -1;
--- if t is after u, it returns +1; if they're the same, it returns 0.
---@param u timeTime
---@return number
function timeTime:Compare(u) end

--- UnixMicro returns t as a Unix time, the number of microseconds elapsed since
--- January 1, 1970 UTC. The result is undefined if the Unix time in
--- microseconds cannot be represented by an int64 (a date before year -290307 or
--- after year 294246). The result does not depend on the location associated
--- with t.
---@return number
function timeTime:UnixMicro() end

--- UnixNano returns t as a Unix time, the number of nanoseconds elapsed
--- since January 1, 1970 UTC. The result is undefined if the Unix time
--- in nanoseconds cannot be represented by an int64 (a date before the year
--- 1678 or after 2262). Note that this means the result of calling UnixNano
--- on the zero Time is undefined. The result does not depend on the
--- location associated with t.
---@return number
function timeTime:UnixNano() end

--- MarshalText implements the encoding.TextMarshaler interface.
--- The time is formatted in RFC 3339 format with sub-second precision.
--- If the timestamp cannot be represented as valid RFC 3339
--- (e.g., the year is out of range), then an error is reported.
---@return byte[], err
function timeTime:MarshalText() end

--- Truncate returns the result of rounding t down to a multiple of d (since the zero time).
--- If d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise unchanged.
---
--- Truncate operates on the time as an absolute duration since the
--- zero time; it does not operate on the presentation form of the
--- time. Thus, Truncate(Hour) may return a time with a non-zero
--- minute, depending on the time's Location.
---@param d timeDuration
---@return timeTime
function timeTime:Truncate(d) end

--- Hour returns the hour within the day specified by t, in the range [0, 23].
---@return number
function timeTime:Hour() end

--- YearDay returns the day of the year specified by t, in the range [1,365] for non-leap years,
--- and [1,366] in leap years.
---@return number
function timeTime:YearDay() end

--- MarshalJSON implements the json.Marshaler interface.
--- The time is a quoted string in the RFC 3339 format with sub-second precision.
--- If the timestamp cannot be represented as valid RFC 3339
--- (e.g., the year is out of range), then an error is reported.
---@return byte[], err
function timeTime:MarshalJSON() end

--- AddDate returns the time corresponding to adding the
--- given number of years, months, and days to t.
--- For example, AddDate(-1, 2, 3) applied to January 1, 2011
--- returns March 4, 2010.
---
--- AddDate normalizes its result in the same way that Date does,
--- so, for example, adding one month to October 31 yields
--- December 1, the normalized form for November 31.
---@param years number
---@param months number
---@param days number
---@return timeTime
function timeTime:AddDate(years, months, days) end

--- GobDecode implements the gob.GobDecoder interface.
---@param data byte[]
---@return err
function timeTime:GobDecode(data) end

--- UnmarshalText implements the encoding.TextUnmarshaler interface.
--- The time must be in the RFC 3339 format.
---@param data byte[]
---@return err
function timeTime:UnmarshalText(data) end

--- UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
---@param data byte[]
---@return err
function timeTime:UnmarshalBinary(data) end

--- After reports whether the time instant t is after u.
---@param u timeTime
---@return boolean
function timeTime:After(u) end

--- Equal reports whether t and u represent the same time instant.
--- Two times can be equal even if they are in different locations.
--- For example, 6:00 +0200 and 4:00 UTC are Equal.
--- See the documentation on the Time type for the pitfalls of using == with
--- Time values; most code should use Equal instead.
---@param u timeTime
---@return boolean
function timeTime:Equal(u) end

--- Day returns the day of the month specified by t.
---@return number
function timeTime:Day() end

--- Weekday returns the day of the week specified by t.
---@return timeWeekday
function timeTime:Weekday() end

--- ZoneBounds returns the bounds of the time zone in effect at time t.
--- The zone begins at start and the next zone begins at end.
--- If the zone begins at the beginning of time, start will be returned as a zero Time.
--- If the zone goes on forever, end will be returned as a zero Time.
--- The Location of the returned times will be the same as t.
---@return timeTime
function timeTime:ZoneBounds() end

--- IsDST reports whether the time in the configured location is in Daylight Savings Time.
---@return boolean
function timeTime:IsDST() end

--- Date returns the year, month, and day in which t occurs.
---@return number, timeMonth, number
function timeTime:Date() end

--- Month returns the month of the year specified by t.
---@return timeMonth
function timeTime:Month() end

--- ISOWeek returns the ISO 8601 year and week number in which t occurs.
--- Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to
--- week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1
--- of year n+1.
---@return number
function timeTime:ISOWeek() end

--- Nanosecond returns the nanosecond offset within the second specified by t,
--- in the range [0, 999999999].
---@return number
function timeTime:Nanosecond() end

--- GobEncode implements the gob.GobEncoder interface.
---@return byte[], err
function timeTime:GobEncode() end

--- UnmarshalJSON implements the json.Unmarshaler interface.
--- The time must be a quoted string in the RFC 3339 format.
---@param data byte[]
---@return err
function timeTime:UnmarshalJSON(data) end

--- A Location maps time instants to the zone in use at that time.
--- Typically, the Location represents the collection of time offsets
--- in use in a geographical area. For many Locations the time offset varies
--- depending on whether daylight savings time is in use at the time instant.
---@class timeLocation
local timeLocation = {}

--- String returns a descriptive name for the time zone information,
--- corresponding to the name argument to LoadLocation or FixedZone.
---@return string
function timeLocation:String() end

--- A Duration represents the elapsed time between two instants
--- as an int64 nanosecond count. The representation limits the
--- largest representable duration to approximately 290 years.
---@class timeDuration
local timeDuration = {}

--- String returns a string representing the duration in the form "72h3m0.5s".
--- Leading zero units are omitted. As a special case, durations less than one
--- second format use a smaller unit (milli-, micro-, or nanoseconds) to ensure
--- that the leading digit is non-zero. The zero duration formats as 0s.
---@return string
function timeDuration:String() end

--- Nanoseconds returns the duration as an integer nanosecond count.
---@return number
function timeDuration:Nanoseconds() end

--- Hours returns the duration as a floating point number of hours.
---@return number
function timeDuration:Hours() end

--- Truncate returns the result of rounding d toward zero to a multiple of m.
--- If m <= 0, Truncate returns d unchanged.
---@param m timeDuration
---@return timeDuration
function timeDuration:Truncate(m) end

--- Round returns the result of rounding d to the nearest multiple of m.
--- The rounding behavior for halfway values is to round away from zero.
--- If the result exceeds the maximum (or minimum)
--- value that can be stored in a Duration,
--- Round returns the maximum (or minimum) duration.
--- If m <= 0, Round returns d unchanged.
---@param m timeDuration
---@return timeDuration
function timeDuration:Round(m) end

--- Abs returns the absolute value of d.
--- As a special case, math.MinInt64 is converted to math.MaxInt64.
---@return timeDuration
function timeDuration:Abs() end

--- Microseconds returns the duration as an integer microsecond count.
---@return number
function timeDuration:Microseconds() end

--- Milliseconds returns the duration as an integer millisecond count.
---@return number
function timeDuration:Milliseconds() end

--- Seconds returns the duration as a floating point number of seconds.
---@return number
function timeDuration:Seconds() end

--- Minutes returns the duration as a floating point number of minutes.
---@return number
function timeDuration:Minutes() end
