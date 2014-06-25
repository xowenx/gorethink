package gorethink

import (
	p "github.com/dancannon/gorethink/ql2"
)

// Returns a time object representing the current time in UTC
func Now(args ...interface{}) RqlTerm {
	return newRqlTerm("Now", p.Term_NOW, args, map[string]interface{}{})
}

// Create a time object for a specific time
func Time(args ...interface{}) RqlTerm {
	return newRqlTerm("Time", p.Term_TIME, args, map[string]interface{}{})
}

// Returns a time object based on seconds since epoch
func EpochTime(args ...interface{}) RqlTerm {
	return newRqlTerm("EpochTime", p.Term_EPOCH_TIME, args, map[string]interface{}{})
}

type ISO8601Opts struct {
	DefaultTimezone interface{} `gorethink:"default_timezone,omitempty"`
}

func (o *ISO8601Opts) toMap() map[string]interface{} {
	return optArgsToMap(o)
}

// Returns a time object based on an ISO8601 formatted date-time string
//
// Optional arguments (see http://www.rethinkdb.com/api/#js:dates_and_times-iso8601 for more information):
// "default_timezone" (string)
func ISO8601(date interface{}, optArgs ...ISO8601Opts) RqlTerm {

	opts := map[string]interface{}{}
	if len(optArgs) >= 1 {
		opts = optArgs[0].toMap()
	}
	return newRqlTerm("ISO8601", p.Term_ISO8601, []interface{}{date}, opts)
}

// Returns a new time object with a different time zone. While the time
// stays the same, the results returned by methods such as hours() will
// change since they take the timezone into account. The timezone argument
// has to be of the ISO 8601 format.
func (t RqlTerm) InTimezone(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "InTimezone", p.Term_IN_TIMEZONE, args, map[string]interface{}{})
}

// Returns the timezone of the time object
func (t RqlTerm) Timezone(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Timezone", p.Term_TIMEZONE, args, map[string]interface{}{})
}

type DuringOpts struct {
	LeftBound  interface{} `gorethink:"left_bound,omitempty"`
	RightBound interface{} `gorethink:"right_bound,omitempty"`
}

func (o *DuringOpts) toMap() map[string]interface{} {
	return optArgsToMap(o)
}

// Returns true if a time is between two other times
// (by default, inclusive for the start, exclusive for the end).
//
// Optional arguments (see http://www.rethinkdb.com/api/#js:dates_and_times-during for more information):
// "left_bound" and "right_bound" ("open" for exclusive or "closed" for inclusive)
func (t RqlTerm) During(startTime, endTime interface{}, optArgs ...DuringOpts) RqlTerm {
	opts := map[string]interface{}{}
	if len(optArgs) >= 1 {
		opts = optArgs[0].toMap()
	}
	return newRqlTermFromPrevVal(t, "During", p.Term_DURING, []interface{}{startTime, endTime}, opts)
}

// Return a new time object only based on the day, month and year
// (ie. the same day at 00:00).
func (t RqlTerm) Date(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Date", p.Term_DATE, args, map[string]interface{}{})
}

// Return the number of seconds elapsed since the beginning of the
// day stored in the time object.
func (t RqlTerm) TimeOfDay(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "TimeOfDay", p.Term_TIME_OF_DAY, args, map[string]interface{}{})
}

// Return the year of a time object.
func (t RqlTerm) Year(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Year", p.Term_YEAR, args, map[string]interface{}{})
}

// Return the month of a time object as a number between 1 and 12.
// For your convenience, the terms r.January(), r.February() etc. are
// defined and map to the appropriate integer.
func (t RqlTerm) Month(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Month", p.Term_MONTH, args, map[string]interface{}{})
}

// Return the day of a time object as a number between 1 and 31.
func (t RqlTerm) Day(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Day", p.Term_DAY, args, map[string]interface{}{})
}

// Return the day of week of a time object as a number between
// 1 and 7 (following ISO 8601 standard). For your convenience,
// the terms r.Monday(), r.Tuesday() etc. are defined and map to
// the appropriate integer.
func (t RqlTerm) DayOfWeek(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "DayOfWeek", p.Term_DAY_OF_WEEK, args, map[string]interface{}{})
}

// Return the day of the year of a time object as a number between
// 1 and 366 (following ISO 8601 standard).
func (t RqlTerm) DayOfYear(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "DayOfYear", p.Term_DAY_OF_YEAR, args, map[string]interface{}{})
}

// Return the hour in a time object as a number between 0 and 23.
func (t RqlTerm) Hours(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Hours", p.Term_HOURS, args, map[string]interface{}{})
}

// Return the minute in a time object as a number between 0 and 59.
func (t RqlTerm) Minutes(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Minutes", p.Term_MINUTES, args, map[string]interface{}{})
}

// Return the seconds in a time object as a number between 0 and
// 59.999 (double precision).
func (t RqlTerm) Seconds(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Seconds", p.Term_SECONDS, args, map[string]interface{}{})
}

// Convert a time object to its iso 8601 format.
func (t RqlTerm) ToISO8601(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "ToISO8601", p.Term_TO_ISO8601, args, map[string]interface{}{})
}

// Convert a time object to its epoch time.
func (t RqlTerm) ToEpochTime(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "ToEpochTime", p.Term_TO_EPOCH_TIME, args, map[string]interface{}{})
}

var (
	// Days
	Monday    = newRqlTerm("Monday", p.Term_MONDAY, []interface{}{}, map[string]interface{}{})
	Tuesday   = newRqlTerm("Tuesday", p.Term_TUESDAY, []interface{}{}, map[string]interface{}{})
	Wednesday = newRqlTerm("Wednesday", p.Term_WEDNESDAY, []interface{}{}, map[string]interface{}{})
	Thursday  = newRqlTerm("Thursday", p.Term_THURSDAY, []interface{}{}, map[string]interface{}{})
	Friday    = newRqlTerm("Friday", p.Term_FRIDAY, []interface{}{}, map[string]interface{}{})
	Saturday  = newRqlTerm("Saturday", p.Term_SATURDAY, []interface{}{}, map[string]interface{}{})
	Sunday    = newRqlTerm("Sunday", p.Term_SUNDAY, []interface{}{}, map[string]interface{}{})

	// Months
	January   = newRqlTerm("January", p.Term_JANUARY, []interface{}{}, map[string]interface{}{})
	February  = newRqlTerm("February", p.Term_FEBRUARY, []interface{}{}, map[string]interface{}{})
	March     = newRqlTerm("March", p.Term_MARCH, []interface{}{}, map[string]interface{}{})
	April     = newRqlTerm("April", p.Term_APRIL, []interface{}{}, map[string]interface{}{})
	May       = newRqlTerm("May", p.Term_MAY, []interface{}{}, map[string]interface{}{})
	June      = newRqlTerm("June", p.Term_JUNE, []interface{}{}, map[string]interface{}{})
	July      = newRqlTerm("July", p.Term_JULY, []interface{}{}, map[string]interface{}{})
	August    = newRqlTerm("August", p.Term_AUGUST, []interface{}{}, map[string]interface{}{})
	September = newRqlTerm("September", p.Term_SEPTEMBER, []interface{}{}, map[string]interface{}{})
	October   = newRqlTerm("October", p.Term_OCTOBER, []interface{}{}, map[string]interface{}{})
	November  = newRqlTerm("November", p.Term_NOVEMBER, []interface{}{}, map[string]interface{}{})
	December  = newRqlTerm("December", p.Term_DECEMBER, []interface{}{}, map[string]interface{}{})
)
