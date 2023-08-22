-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class strings
strings = {}

---@deprecated
---Please use table.contact() to instead of it
---
---Join concatenates the elements of its first argument to create a single string. The separator
---string sep is placed between elements in the resulting string.
---@param elems string[]
---@param sep string
---@return string
function strings.Join(elems, sep) end

--- HasPrefix tests whether the string s begins with prefix.
---@param s string
---@param prefix string
---@return boolean
function strings.HasPrefix(s, prefix) end

--- HasSuffix tests whether the string s ends with suffix.
---@param s string
---@param suffix string
---@return boolean
function strings.HasSuffix(s, suffix) end

--- Cut slices s around the first instance of sep,
--- returning the text before and after sep.
--- The found result reports whether sep appears in s.
--- If sep does not appear in s, cut returns s, "", false.
---@param s string
---@param sep string
---@return string, string, boolean
function strings.Cut(s, sep) end

--- CutPrefix returns s without the provided leading prefix string
--- and reports whether it found the prefix.
--- If s doesn't start with prefix, CutPrefix returns s, false.
--- If prefix is the empty string, CutPrefix returns s, true.
---@param s string
---@param prefix string
---@return string, boolean
function strings.CutPrefix(s, prefix) end

--- CutSuffix returns s without the provided ending suffix string
--- and reports whether it found the suffix.
--- If s doesn't end with suffix, CutSuffix returns s, false.
--- If suffix is the empty string, CutSuffix returns s, true.
---@param s string
---@param suffix string
---@return string, boolean
function strings.CutSuffix(s, suffix) end

--- Contains reports whether substr is within s.
---@param s string
---@param substr string
---@return boolean
function strings.Contains(s, substr) end

--- ContainsAny reports whether any Unicode code points in chars are within s.
---@param s string
---@param chars string
---@return boolean
function strings.ContainsAny(s, chars) end

--- ContainsRune reports whether the Unicode code point r is within s.
---@param s string
---@param r any
---@return boolean
function strings.ContainsRune(s, r) end

--- Count counts the number of non-overlapping instances of substr in s.
--- If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
---@param s string
---@param substr string
---@return number
function strings.Count(s, substr) end

--- Replace returns a copy of the string s with the first n
--- non-overlapping instances of old replaced by new.
--- If old is empty, it matches at the beginning of the string
--- and after each UTF-8 sequence, yielding up to k+1 replacements
--- for a k-rune string.
--- If n < 0, there is no limit on the number of replacements.
---@param s string
---@param old string
---@param new string
---@param n number
---@return string
function strings.Replace(s, old, new, n) end

--- ReplaceAll returns a copy of the string s with all
--- non-overlapping instances of old replaced by new.
--- If old is empty, it matches at the beginning of the string
--- and after each UTF-8 sequence, yielding up to k+1 replacements
--- for a k-rune string.
---@param s string
---@param old string
---@param new string
---@return string
function strings.ReplaceAll(s, old, new) end

--- Clone returns a fresh copy of s.
--- It guarantees to make a copy of s into a new allocation,
--- which can be important when retaining only a small substring
--- of a much larger string. Using Clone can help such programs
--- use less memory. Of course, since using Clone makes a copy,
--- overuse of Clone can make programs use more memory.
--- Clone should typically be used only rarely, and only when
--- profiling indicates that it is needed.
--- For strings of length zero the string "" will be returned
--- and no allocation is made.
---@param s string
---@return string
function strings.Clone(s) end

--- Compare returns an integer comparing two strings lexicographically.
--- The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
---
--- Compare is included only for symmetry with package bytes.
--- It is usually clearer and always faster to use the built-in
--- string comparison operators ==, <, >, and so on.
---@param a string
---@param b string
---@return number
function strings.Compare(a, b) end

-- Split slices s into all substrings separated by sep and returns a slice of
-- the substrings between those separators.
--
-- If s does not contain sep and sep is not empty, Split returns a
-- slice of length 1 whose only element is s.
--
-- If sep is empty, Split splits after each UTF-8 sequence. If both s
-- and sep are empty, Split returns an empty slice.
--
-- It is equivalent to SplitN with a count of -1.
--
-- To split around the first instance of a separator, see Cut.
---@param s string
---@param sep string
---@return string[]
function strings.Split(s, sep) end

--- SplitN slices s into substrings separated by sep and returns a slice of
--- the substrings between those separators.
---
--- The count determines the number of substrings to return:
---
---	n > 0: at most n substrings; the last substring will be the unsplit remainder.
---	n == 0: the result is nil (zero substrings)
---	n < 0: all substrings
---
--- Edge cases for s and sep (for example, empty strings) are handled
--- as described in the documentation for Split.
---
--- To split around the first instance of a separator, see Cut.
---@param s string
---@param sep string
---@param n number
---@return string[]
function strings.SplitN(s, sep, n) end

--- NewReader returns a new Reader reading from s.
--- It is similar to bytes.NewBufferString but more efficient and read-only.
---@param s string
---@return stringsReader
function strings.NewReader(s) end

--- TrimSpace returns a slice of the string s, with all leading
--- and trailing white space removed, as defined by Unicode.
---@param s string
---@return string
function strings.TrimSpace(s) end

--- LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
---@param s string
---@param substr string
---@return integer
function strings.LastIndex(s, substr) end

--- IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
---@param s string
---@param c byte
---@return integer
function strings.IndexByte(s, c) end

--- IndexRune returns the index of the first instance of the Unicode code point
--- r, or -1 if rune is not present in s.
--- If r is utf8.RuneError, it returns the first instance of any
--- invalid UTF-8 byte sequence.
---@param s string
---@param r any
---@return integer
function strings.IndexRune(s, r) end

--- IndexAny returns the index of the first instance of any Unicode code point
--- from chars in s, or -1 if no Unicode code point from chars is present in s.
---@param s string
---@param chars string
---@return integer
function strings.IndexAny(s, chars) end

--- LastIndexAny returns the index of the last instance of any Unicode code
--- point from chars in s, or -1 if no Unicode code point from chars is
--- present in s.
---@param s string
---@param chars string
---@return number
function strings.LastIndexAny(s, chars) end

--- LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
---@param s string
---@param c byte
---@return integer
function strings.LastIndexByte(s, c) end

--- SplitN slices s into substrings separated by sep and returns a slice of
--- the substrings between those separators.
---
--- The count determines the number of substrings to return:
---
---	n > 0: at most n substrings; the last substring will be the unsplit remainder.
---	n == 0: the result is nil (zero substrings)
---	n < 0: all substrings
---
--- Edge cases for s and sep (for example, empty strings) are handled
--- as described in the documentation for Split.
---
--- To split around the first instance of a separator, see Cut.
---@param s string
---@param sep string
---@param n number
---@return string[]
function strings.SplitN(s, sep, n) end

--- SplitAfterN slices s into substrings after each instance of sep and
--- returns a slice of those substrings.
---
--- The count determines the number of substrings to return:
---
---	n > 0: at most n substrings; the last substring will be the unsplit remainder.
---	n == 0: the result is nil (zero substrings)
---	n < 0: all substrings
---
--- Edge cases for s and sep (for example, empty strings) are handled
--- as described in the documentation for SplitAfter.
---@param s string
---@param sep string
---@param n number
---@return string[]
function strings.SplitAfterN(s, sep, n) end

--- SplitAfter slices s into all substrings after each instance of sep and
--- returns a slice of those substrings.
---
--- If s does not contain sep and sep is not empty, SplitAfter returns
--- a slice of length 1 whose only element is s.
---
--- If sep is empty, SplitAfter splits after each UTF-8 sequence. If
--- both s and sep are empty, SplitAfter returns an empty slice.
---
--- It is equivalent to SplitAfterN with a count of -1.
---@param s string
---@param sep string
---@return string[]
function strings.SplitAfter(s, sep) end

--- Fields splits the string s around each instance of one or more consecutive white space
--- characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
--- empty slice if s contains only white space.
---@param s string
---@return string[]
function strings.Fields(s) end

--- Repeat returns a new string consisting of count copies of the string s.
---
--- It panics if count is negative or if the result of (len(s) * count)
--- overflows.
---@param s string
---@param count integer
---@return string
function strings.Repeat(s, count) end

--- ToUpper returns s with all Unicode letters mapped to their upper case.
---@param s string
---@return string
function strings.ToUpper(s) end

--- ToLower returns s with all Unicode letters mapped to their lower case.
---@param s string
---@return string
function strings.ToLower(s) end

--- ToTitle returns a copy of the string s with all Unicode letters mapped to
--- their Unicode title case.
---@param s string
---@return string
function strings.ToTitle(s) end

--- FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c)
--- and returns an array of slices of s. If all code points in s satisfy f(c) or the
--- string is empty, an empty slice is returned.
---
--- FieldsFunc makes no guarantees about the order in which it calls f(c)
--- and assumes that f always returns the same value for a given c.
---@param s string
---@param f fun(r: integer): boolean
---@return string[]
function strings.FieldsFunc(s, f) end

--- Map returns a copy of the string s with all its characters modified
--- according to the mapping function. If mapping returns a negative value, the character is
--- dropped from the string with no replacement.
---@param mapping fun(r: integer): integer
---@param s string
---@return string
function strings.Map(mapping, s) end

--- TrimLeftFunc returns a slice of the string s with all leading
--- Unicode code points c satisfying f(c) removed.
---@param s string
---@param f fun(r: integer): boolean
---@return string
function strings.TrimLeftFunc(s, f) end

--- TrimRightFunc returns a slice of the string s with all trailing
--- Unicode code points c satisfying f(c) removed.
---@param s string
---@param f fun(r: integer): boolean
---@return string
function strings.TrimRightFunc(s, f) end

--- TrimFunc returns a slice of the string s with all leading
--- and trailing Unicode code points c satisfying f(c) removed.
---@param s string
---@param f fun(r: integer): boolean
---@return string
function strings.TrimFunc(s, f) end

--- IndexFunc returns the index into s of the first Unicode
--- code point satisfying f(c), or -1 if none do.
---@param s string
---@param f fun(r: integer): boolean
---@return integer
function strings.IndexFunc(s, f) end

--- LastIndexFunc returns the index into s of the last
--- Unicode code point satisfying f(c), or -1 if none do.
---@param s string
---@param f fun(r: integer): boolean
---@return integer
function strings.LastIndexFunc(s, f) end

--- TrimPrefix returns s without the provided leading prefix string.
--- If s doesn't start with prefix, s is returned unchanged.
---@param s string
---@param prefix string
---@return string
function strings.TrimPrefix(s, prefix) end

--- TrimSuffix returns s without the provided trailing suffix string.
--- If s doesn't end with suffix, s is returned unchanged.
---@param s string
---@param suffix string
---@return string
function strings.TrimSuffix(s, suffix) end

--- A Reader implements the io.Reader, io.ReaderAt, io.ByteReader, io.ByteScanner,
--- io.RuneReader, io.RuneScanner, io.Seeker, and io.WriterTo interfaces by reading
--- from a string.
--- The zero value for Reader operates like a Reader of an empty string.
---@class stringsReader
local stringsReader = {}
