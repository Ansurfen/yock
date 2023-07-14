-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@class reflect
---@field Invalid any
---@field Bool any
---@field Int any
---@field Int8 any
---@field Int16 any
---@field Int32 any
---@field Int64 any
---@field Uint any
---@field Uint8 any
---@field Uint16 any
---@field Uint32 any
---@field Uint64 any
---@field Uintptr any
---@field Float32 any
---@field Float64 any
---@field Complex64 any
---@field Complex128 any
---@field Array any
---@field Chan any
---@field Func any
---@field Interface any
---@field Map any
---@field Pointer any
---@field Slice any
---@field String any
---@field Struct any
---@field UnsafePointer any
---@field Ptr any
---@field RecvDir any
---@field SendDir any
---@field BothDir any
---@field SelectSend any
---@field SelectRecv any
---@field SelectDefault any
reflect = {}

--- ArrayOf returns the array type with the given length and element type.
--- For example, if t represents int, ArrayOf(5, t) represents [5]int.
---
--- If the resulting type would be larger than the available address space,
--- ArrayOf panics.
---@param length number
---@param elem reflectType
---@return reflectType
function reflect.ArrayOf(length, elem) end

--- ArenaNew returns a Value representing a pointer to a new zero value for the
--- specified type, allocating storage for it in the provided arena. That is,
--- the returned Value's Type is PointerTo(typ).
---@param a arenaArena
---@param typ reflectType
---@return reflectValue
function reflect.ArenaNew(a, typ) end

--- MapOf returns the map type with the given key and element types.
--- For example, if k represents int and e represents string,
--- MapOf(k, e) represents map[int]string.
---
--- If the key type is not a valid map key type (that is, if it does
--- not implement Go's == operator), MapOf panics.
---@param key reflectType
---@param elem reflectType
---@return reflectType
function reflect.MapOf(key, elem) end

--- AppendSlice appends a slice t to a slice s and returns the resulting slice.
--- The slices s and t must have the same element type.
---@param s reflectValue
---@param t reflectValue
---@return reflectValue
function reflect.AppendSlice(s, t) end

--- SliceOf returns the slice type with element type t.
--- For example, if t represents int, SliceOf(t) represents []int.
---@param t reflectType
---@return reflectType
function reflect.SliceOf(t) end

--- Indirect returns the value that v points to.
--- If v is a nil pointer, Indirect returns a zero Value.
--- If v is not a pointer, Indirect returns v.
---@param v reflectValue
---@return reflectValue
function reflect.Indirect(v) end

--- Select executes a select operation described by the list of cases.
--- Like the Go select statement, it blocks until at least one of the cases
--- can proceed, makes a uniform pseudo-random choice,
--- and then executes that case. It returns the index of the chosen case
--- and, if that case was a receive operation, the value received and a
--- boolean indicating whether the value corresponds to a send on the channel
--- (as opposed to a zero value received because the channel is closed).
--- Select supports a maximum of 65536 cases.
---@param cases any
---@return number, reflectValue, boolean
function reflect.Select(cases) end

--- MakeMap creates a new map with the specified type.
---@param typ reflectType
---@return reflectValue
function reflect.MakeMap(typ) end

--- VisibleFields returns all the visible fields in t, which must be a
--- struct type. A field is defined as visible if it's accessible
--- directly with a FieldByName call. The returned fields include fields
--- inside anonymous struct members and unexported fields. They follow
--- the same order found in the struct, with anonymous fields followed
--- immediately by their promoted fields.
---
--- For each element e of the returned slice, the corresponding field
--- can be retrieved from a value v of type t by calling v.FieldByIndex(e.Index).
---@param t reflectType
---@return any
function reflect.VisibleFields(t) end

--- DeepEqual reports whether x and y are “deeply equal,” defined as follows.
--- Two values of identical type are deeply equal if one of the following cases applies.
--- Values of distinct types are never deeply equal.
---
--- Array values are deeply equal when their corresponding elements are deeply equal.
---
--- Struct values are deeply equal if their corresponding fields,
--- both exported and unexported, are deeply equal.
---
--- Func values are deeply equal if both are nil; otherwise they are not deeply equal.
---
--- Interface values are deeply equal if they hold deeply equal concrete values.
---
--- Map values are deeply equal when all of the following are true:
--- they are both nil or both non-nil, they have the same length,
--- and either they are the same map object or their corresponding keys
--- (matched using Go equality) map to deeply equal values.
---
--- Pointer values are deeply equal if they are equal using Go's == operator
--- or if they point to deeply equal values.
---
--- Slice values are deeply equal when all of the following are true:
--- they are both nil or both non-nil, they have the same length,
--- and either they point to the same initial entry of the same underlying array
--- (that is, &x[0] == &y[0]) or their corresponding elements (up to length) are deeply equal.
--- Note that a non-nil empty slice and a nil slice (for example, []byte{} and []byte(nil))
--- are not deeply equal.
---
--- Other values - numbers, bools, strings, and channels - are deeply equal
--- if they are equal using Go's == operator.
---
--- In general DeepEqual is a recursive relaxation of Go's == operator.
--- However, this idea is impossible to implement without some inconsistency.
--- Specifically, it is possible for a value to be unequal to itself,
--- either because it is of func type (uncomparable in general)
--- or because it is a floating-point NaN value (not equal to itself in floating-point comparison),
--- or because it is an array, struct, or interface containing
--- such a value.
--- On the other hand, pointer values are always equal to themselves,
--- even if they point at or contain such problematic values,
--- because they compare equal using Go's == operator, and that
--- is a sufficient condition to be deeply equal, regardless of content.
--- DeepEqual has been defined so that the same short-cut applies
--- to slices and maps: if x and y are the same slice or the same map,
--- they are deeply equal regardless of content.
---
--- As DeepEqual traverses the data values it may find a cycle. The
--- second and subsequent times that DeepEqual compares two pointer
--- values that have been compared before, it treats the values as
--- equal rather than examining the values to which they point.
--- This ensures that DeepEqual terminates.
---@param x any
---@param y any
---@return boolean
function reflect.DeepEqual(x, y) end

--- FuncOf returns the function type with the given argument and result types.
--- For example if k represents int and e represents string,
--- FuncOf([]Type{k}, []Type{e}, false) represents func(int) string.
---
--- The variadic argument controls whether the function is variadic. FuncOf
--- panics if the in[len(in)-1] does not represent a slice and variadic is
--- true.
---@param in_ any
---@param out_ any
---@param variadic boolean
---@return reflectType
function reflect.FuncOf(in_, out_, variadic) end

--- Copy copies the contents of src into dst until either
--- dst has been filled or src has been exhausted.
--- It returns the number of elements copied.
--- Dst and src each must have kind Slice or Array, and
--- dst and src must have the same element type.
---
--- As a special case, src can have kind String if the element type of dst is kind Uint8.
---@param dst reflectValue
---@param src reflectValue
---@return number
function reflect.Copy(dst, src) end

--- NewAt returns a Value representing a pointer to a value of the
--- specified type, using p as that pointer.
---@param typ reflectType
---@param p unsafePointer
---@return reflectValue
function reflect.NewAt(typ, p) end

--- PointerTo returns the pointer type with element t.
--- For example, if t represents type Foo, PointerTo(t) represents *Foo.
---@param t reflectType
---@return reflectType
function reflect.PointerTo(t) end

--- PtrTo returns the pointer type with element t.
--- For example, if t represents type Foo, PtrTo(t) represents *Foo.
---
--- PtrTo is the old spelling of PointerTo.
--- The two functions behave identically.
---@param t reflectType
---@return reflectType
function reflect.PtrTo(t) end

--- Zero returns a Value representing the zero value for the specified type.
--- The result is different from the zero value of the Value struct,
--- which represents no value at all.
--- For example, Zero(TypeOf(42)) returns a Value with Kind Int and value 0.
--- The returned value is neither addressable nor settable.
---@param typ reflectType
---@return reflectValue
function reflect.Zero(typ) end

--- MakeSlice creates a new zero-initialized slice value
--- for the specified slice type, length, and capacity.
---@param typ reflectType
---@param len number
---@param cap number
---@return reflectValue
function reflect.MakeSlice(typ, len, cap) end

--- MakeMapWithSize creates a new map with the specified type
--- and initial space for approximately n elements.
---@param typ reflectType
---@param n number
---@return reflectValue
function reflect.MakeMapWithSize(typ, n) end

--- ChanOf returns the channel type with the given direction and element type.
--- For example, if t represents int, ChanOf(RecvDir, t) represents <-chan int.
---
--- The gc runtime imposes a limit of 64 kB on channel element types.
--- If t's size is equal to or exceeds this limit, ChanOf panics.
---@param dir reflectChanDir
---@param t reflectType
---@return reflectType
function reflect.ChanOf(dir, t) end

--- MakeFunc returns a new function of the given Type
--- that wraps the function fn. When called, that new function
--- does the following:
---
---   - converts its arguments to a slice of Values.
---   - runs results := fn(args).
---   - returns the results as a slice of Values, one per formal result.
---
--- The implementation fn can assume that the argument Value slice
--- has the number and type of arguments given by typ.
--- If typ describes a variadic function, the final Value is itself
--- a slice representing the variadic arguments, as in the
--- body of a variadic function. The result Value slice returned by fn
--- must have the number and type of results given by typ.
---
--- The Value.Call method allows the caller to invoke a typed function
--- in terms of Values; in contrast, MakeFunc allows the caller to implement
--- a typed function in terms of Values.
---
--- The Examples section of the documentation includes an illustration
--- of how to use MakeFunc to build a swap function for different types.
---@param typ reflectType
---@param fn function
---@return reflectValue
function reflect.MakeFunc(typ, fn) end

--- TypeOf returns the reflection Type that represents the dynamic type of i.
--- If i is a nil interface value, TypeOf returns nil.
---@param i any
---@return reflectType
function reflect.TypeOf(i) end

--- ValueOf returns a new Value initialized to the concrete value
--- stored in the interface i. ValueOf(nil) returns the zero Value.
---@param i any
---@return reflectValue
function reflect.ValueOf(i) end

--- Append appends the values x to a slice s and returns the resulting slice.
--- As in Go, each x's value must be assignable to the slice's element type.
---@param s reflectValue
---@vararg reflectValue
---@return reflectValue
function reflect.Append(s, ...) end

--- MakeChan creates a new channel with the specified type and buffer size.
---@param typ reflectType
---@param buffer number
---@return reflectValue
function reflect.MakeChan(typ, buffer) end

--- Swapper returns a function that swaps the elements in the provided
--- slice.
---
--- Swapper panics if the provided interface is not a slice.
---@param slice any
---@return any
function reflect.Swapper(slice) end

--- StructOf returns the struct type containing fields.
--- The Offset and Index fields are ignored and computed as they would be
--- by the compiler.
---
--- StructOf currently does not generate wrapper methods for embedded
--- fields and panics if passed unexported StructFields.
--- These limitations may be lifted in a future version.
---@param fields any
---@return reflectType
function reflect.StructOf(fields) end

--- New returns a Value representing a pointer to a new zero value
--- for the specified type. That is, the returned Value's Type is PointerTo(typ).
---@param typ reflectType
---@return reflectValue
function reflect.New(typ) end

--- Value is the reflection interface to a Go value.
---
--- Not all methods apply to all kinds of values. Restrictions,
--- if any, are noted in the documentation for each method.
--- Use the Kind method to find out the kind of value before
--- calling kind-specific methods. Calling a method
--- inappropriate to the kind of type causes a run time panic.
---
--- The zero Value represents no value.
--- Its IsValid method returns false, its Kind method returns Invalid,
--- its String method returns "<invalid Value>", and all other methods panic.
--- Most functions and methods never return an invalid value.
--- If one does, its documentation states the conditions explicitly.
---
--- A Value can be used concurrently by multiple goroutines provided that
--- the underlying Go value can be used concurrently for the equivalent
--- direct operations.
---
--- To compare two Values, compare the results of the Interface method.
--- Using == on two Values does not compare the underlying values
--- they represent.
---@class reflectValue
local reflectValue = {}

--- CanConvert reports whether the value v can be converted to type t.
--- If v.CanConvert(t) returns true then v.Convert(t) will not panic.
---@param t reflectType
---@return boolean
function reflectValue:CanConvert(t) end

--- FieldByName returns the struct field with the given name.
--- It returns the zero Value if no field was found.
--- It panics if v's Kind is not struct.
---@param name string
---@return reflectValue
function reflectValue:FieldByName(name) end

--- TrySend attempts to send x on the channel v but will not block.
--- It panics if v's Kind is not Chan.
--- It reports whether the value was sent.
--- As in Go, x's value must be assignable to the channel's element type.
---@param x reflectValue
---@return boolean
function reflectValue:TrySend(x) end

--- Close closes the channel v.
--- It panics if v's Kind is not Chan.
function reflectValue:Close() end

--- CanFloat reports whether Float can be used without panicking.
---@return boolean
function reflectValue:CanFloat() end

--- NumField returns the number of fields in the struct v.
--- It panics if v's Kind is not Struct.
---@return number
function reflectValue:NumField() end

--- Recv receives and returns a value from the channel v.
--- It panics if v's Kind is not Chan.
--- The receive blocks until a value is ready.
--- The boolean value ok is true if the value x corresponds to a send
--- on the channel, false if it is a zero value received because the channel is closed.
---@return reflectValue, boolean
function reflectValue:Recv() end

--- Bytes returns v's underlying value.
--- It panics if v's underlying value is not a slice of bytes or
--- an addressable array of bytes.
---@return byte[]
function reflectValue:Bytes() end

--- Complex returns v's underlying value, as a complex128.
--- It panics if v's Kind is not Complex64 or Complex128
---@return any
function reflectValue:Complex() end

--- SetZero sets v to be the zero value of v's type.
--- It panics if CanSet returns false.
function reflectValue:SetZero() end

--- Convert returns the value v converted to type t.
--- If the usual Go conversion rules do not allow conversion
--- of the value v to type t, or if converting v to type t panics, Convert panics.
---@param t reflectType
---@return reflectValue
function reflectValue:Convert(t) end

--- Bool returns v's underlying value.
--- It panics if v's kind is not Bool.
---@return boolean
function reflectValue:Bool() end

--- Call calls the function v with the input arguments in.
--- For example, if len(in) == 3, v.Call(in) represents the Go call v(in[0], in[1], in[2]).
--- Call panics if v's Kind is not Func.
--- It returns the output results as Values.
--- As in Go, each input argument must be assignable to the
--- type of the function's corresponding input parameter.
--- If v is a variadic function, Call creates the variadic slice parameter
--- itself, copying in the corresponding values.
---@param in_ any
---@return any
function reflectValue:Call(in_) end

--- OverflowUint reports whether the uint64 x cannot be represented by v's type.
--- It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.
---@param x number
---@return boolean
function reflectValue:OverflowUint(x) end

--- MapRange returns a range iterator for a map.
--- It panics if v's Kind is not Map.
---
--- Call Next to advance the iterator, and Key/Value to access each entry.
--- Next returns false when the iterator is exhausted.
--- MapRange follows the same iteration semantics as a range statement.
---
--- Example:
---
---	iter := reflect.ValueOf(m).MapRange()
---	for iter.Next() {
---		k := iter.Key()
---		v := iter.Value()
---		...
---	}
---@return reflectMapIter
function reflectValue:MapRange() end

--- MethodByName returns a function value corresponding to the method
--- of v with the given name.
--- The arguments to a Call on the returned function should not include
--- a receiver; the returned function will always use v as the receiver.
--- It returns the zero Value if no method was found.
---@param name string
---@return reflectValue
function reflectValue:MethodByName(name) end

--- Grow increases the slice's capacity, if necessary, to guarantee space for
--- another n elements. After Grow(n), at least n elements can be appended
--- to the slice without another allocation.
---
--- It panics if v's Kind is not a Slice or if n is negative or too large to
--- allocate the memory.
---@param n number
function reflectValue:Grow(n) end

--- Addr returns a pointer value representing the address of v.
--- It panics if CanAddr() returns false.
--- Addr is typically used to obtain a pointer to a struct field
--- or slice element in order to call a method that requires a
--- pointer receiver.
---@return reflectValue
function reflectValue:Addr() end

--- FieldByIndexErr returns the nested field corresponding to index.
--- It returns an error if evaluation requires stepping through a nil
--- pointer, but panics if it must step through a field that
--- is not a struct.
---@param index any
---@return reflectValue, err
function reflectValue:FieldByIndexErr(index) end

--- Set assigns x to the value v.
--- It panics if CanSet returns false.
--- As in Go, x's value must be assignable to v's type and
--- must not be derived from an unexported field.
---@param x reflectValue
function reflectValue:Set(x) end

--- SetBytes sets v's underlying value.
--- It panics if v's underlying value is not a slice of bytes.
---@param x byte[]
function reflectValue:SetBytes(x) end

--- Slice3 is the 3-index form of the slice operation: it returns v[i:j:k].
--- It panics if v's Kind is not Array or Slice, or if v is an unaddressable array,
--- or if the indexes are out of bounds.
---@param i number
---@param j number
---@param k number
---@return reflectValue
function reflectValue:Slice3(i, j, k) end

--- MapIndex returns the value associated with key in the map v.
--- It panics if v's Kind is not Map.
--- It returns the zero Value if key is not found in the map or if v represents a nil map.
--- As in Go, the key's value must be assignable to the map's key type.
---@param key reflectValue
---@return reflectValue
function reflectValue:MapIndex(key) end

--- Send sends x on the channel v.
--- It panics if v's kind is not Chan or if x's type is not the same type as v's element type.
--- As in Go, x's value must be assignable to the channel's element type.
---@param x reflectValue
function reflectValue:Send(x) end

--- SetString sets v's underlying value to x.
--- It panics if v's Kind is not String or if CanSet() is false.
---@param x string
function reflectValue:SetString(x) end

--- UnsafePointer returns v's value as a [unsafe.Pointer].
--- It panics if v's Kind is not Chan, Func, Map, Pointer, Slice, or UnsafePointer.
---
--- If v's Kind is Func, the returned pointer is an underlying
--- code pointer, but not necessarily enough to identify a
--- single function uniquely. The only guarantee is that the
--- result is zero if and only if v is a nil func Value.
---
--- If v's Kind is Slice, the returned pointer is to the first
--- element of the slice. If the slice is nil the returned value
--- is nil.  If the slice is empty but non-nil the return value is non-nil.
---@return unsafePointer
function reflectValue:UnsafePointer() end

--- CanInt reports whether Int can be used without panicking.
---@return boolean
function reflectValue:CanInt() end

--- Len returns v's length.
--- It panics if v's Kind is not Array, Chan, Map, Slice, String, or pointer to Array.
---@return number
function reflectValue:Len() end

--- MapKeys returns a slice containing all the keys present in the map,
--- in unspecified order.
--- It panics if v's Kind is not Map.
--- It returns an empty slice if v represents a nil map.
---@return any
function reflectValue:MapKeys() end

--- Cap returns v's capacity.
--- It panics if v's Kind is not Array, Chan, Slice or pointer to Array.
---@return number
function reflectValue:Cap() end

--- Elem returns the value that the interface v contains
--- or that the pointer v points to.
--- It panics if v's Kind is not Interface or Pointer.
--- It returns the zero Value if v is nil.
---@return reflectValue
function reflectValue:Elem() end

--- Int returns v's underlying value, as an int64.
--- It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.
---@return number
function reflectValue:Int() end

--- Kind returns v's Kind.
--- If v is the zero Value (IsValid returns false), Kind returns Invalid.
---@return reflectKind
function reflectValue:Kind() end

--- SetCap sets v's capacity to n.
--- It panics if v's Kind is not Slice or if n is smaller than the length or
--- greater than the capacity of the slice.
---@param n number
function reflectValue:SetCap(n) end

--- SetPointer sets the [unsafe.Pointer] value v to x.
--- It panics if v's Kind is not UnsafePointer.
---@param x unsafePointer
function reflectValue:SetPointer(x) end

--- String returns the string v's underlying value, as a string.
--- String is a special case because of Go's String method convention.
--- Unlike the other getters, it does not panic if v's Kind is not String.
--- Instead, it returns a string of the form "<T value>" where T is v's type.
--- The fmt package treats Values specially. It does not call their String
--- method implicitly but instead prints the concrete values they hold.
---@return string
function reflectValue:String() end

--- UnsafeAddr returns a pointer to v's data, as a uintptr.
--- It panics if v is not addressable.
---
--- It's preferred to use uintptr(Value.Addr().UnsafePointer()) to get the equivalent result.
---@return any
function reflectValue:UnsafeAddr() end

--- Field returns the i'th field of the struct v.
--- It panics if v's Kind is not Struct or i is out of range.
---@param i number
---@return reflectValue
function reflectValue:Field(i) end

--- CanInterface reports whether Interface can be used without panicking.
---@return boolean
function reflectValue:CanInterface() end

--- OverflowInt reports whether the int64 x cannot be represented by v's type.
--- It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.
---@param x number
---@return boolean
function reflectValue:OverflowInt(x) end

--- SetInt sets v's underlying value to x.
--- It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64, or if CanSet() is false.
---@param x number
function reflectValue:SetInt(x) end

--- SetUint sets v's underlying value to x.
--- It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64, or if CanSet() is false.
---@param x number
function reflectValue:SetUint(x) end

--- CanUint reports whether Uint can be used without panicking.
---@return boolean
function reflectValue:CanUint() end

--- Uint returns v's underlying value, as a uint64.
--- It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.
---@return number
function reflectValue:Uint() end

--- CanAddr reports whether the value's address can be obtained with Addr.
--- Such values are called addressable. A value is addressable if it is
--- an element of a slice, an element of an addressable array,
--- a field of an addressable struct, or the result of dereferencing a pointer.
--- If CanAddr returns false, calling Addr will panic.
---@return boolean
function reflectValue:CanAddr() end

--- CanSet reports whether the value of v can be changed.
--- A Value can be changed only if it is addressable and was not
--- obtained by the use of unexported struct fields.
--- If CanSet returns false, calling Set or any type-specific
--- setter (e.g., SetBool, SetInt) will panic.
---@return boolean
function reflectValue:CanSet() end

--- CallSlice calls the variadic function v with the input arguments in,
--- assigning the slice in[len(in)-1] to v's final variadic argument.
--- For example, if len(in) == 3, v.CallSlice(in) represents the Go call v(in[0], in[1], in[2]...).
--- CallSlice panics if v's Kind is not Func or if v is not variadic.
--- It returns the output results as Values.
--- As in Go, each input argument must be assignable to the
--- type of the function's corresponding input parameter.
---@param in_ any
---@return any
function reflectValue:CallSlice(in_) end

--- SetComplex sets v's underlying value to x.
--- It panics if v's Kind is not Complex64 or Complex128, or if CanSet() is false.
---@param x any
function reflectValue:SetComplex(x) end

--- Equal reports true if v is equal to u.
--- For two invalid values, Equal will report true.
--- For an interface value, Equal will compare the value within the interface.
--- Otherwise, If the values have different types, Equal will report false.
--- Otherwise, for arrays and structs Equal will compare each element in order,
--- and report false if it finds non-equal elements.
--- During all comparisons, if values of the same type are compared,
--- and the type is not comparable, Equal will panic.
---@param u reflectValue
---@return boolean
function reflectValue:Equal(u) end

--- IsZero reports whether v is the zero value for its type.
--- It panics if the argument is invalid.
---@return boolean
function reflectValue:IsZero() end

--- SetIterValue assigns to v the value of iter's current map entry.
--- It is equivalent to v.Set(iter.Value()), but it avoids allocating a new Value.
--- As in Go, the value must be assignable to v's type and
--- must not be derived from an unexported field.
---@param iter reflectMapIter
function reflectValue:SetIterValue(iter) end

--- SetBool sets v's underlying value.
--- It panics if v's Kind is not Bool or if CanSet() is false.
---@param x boolean
function reflectValue:SetBool(x) end

--- OverflowFloat reports whether the float64 x cannot be represented by v's type.
--- It panics if v's Kind is not Float32 or Float64.
---@param x number
---@return boolean
function reflectValue:OverflowFloat(x) end

--- Comparable reports whether the value v is comparable.
--- If the type of v is an interface, this checks the dynamic type.
--- If this reports true then v.Interface() == x will not panic for any x,
--- nor will v.Equal(u) for any Value u.
---@return boolean
function reflectValue:Comparable() end

--- FieldByNameFunc returns the struct field with a name
--- that satisfies the match function.
--- It panics if v's Kind is not struct.
--- It returns the zero Value if no field was found.
---@param match any
---@return reflectValue
function reflectValue:FieldByNameFunc(match) end

--- Index returns v's i'th element.
--- It panics if v's Kind is not Array, Slice, or String or i is out of range.
---@param i number
---@return reflectValue
function reflectValue:Index(i) end

--- SetFloat sets v's underlying value to x.
--- It panics if v's Kind is not Float32 or Float64, or if CanSet() is false.
---@param x number
function reflectValue:SetFloat(x) end

--- Interface returns v's current value as an interface{}.
--- It is equivalent to:
---
---	var i interface{} = (v's underlying value)
---
--- It panics if the Value was obtained by accessing
--- unexported struct fields.
---@return any
function reflectValue:Interface() end

--- IsValid reports whether v represents a value.
--- It returns false if v is the zero Value.
--- If IsValid returns false, all other methods except String panic.
--- Most functions and methods never return an invalid Value.
--- If one does, its documentation states the conditions explicitly.
---@return boolean
function reflectValue:IsValid() end

--- IsNil reports whether its argument v is nil. The argument must be
--- a chan, func, interface, map, pointer, or slice value; if it is
--- not, IsNil panics. Note that IsNil is not always equivalent to a
--- regular comparison with nil in Go. For example, if v was created
--- by calling ValueOf with an uninitialized interface variable i,
--- i==nil will be true but v.IsNil will panic as v will be the zero
--- Value.
---@return boolean
function reflectValue:IsNil() end

--- SetMapIndex sets the element associated with key in the map v to elem.
--- It panics if v's Kind is not Map.
--- If elem is the zero Value, SetMapIndex deletes the key from the map.
--- Otherwise if v holds a nil map, SetMapIndex will panic.
--- As in Go, key's elem must be assignable to the map's key type,
--- and elem's value must be assignable to the map's elem type.
---@param key reflectValue
---@param elem reflectValue
function reflectValue:SetMapIndex(key, elem) end

--- TryRecv attempts to receive a value from the channel v but will not block.
--- It panics if v's Kind is not Chan.
--- If the receive delivers a value, x is the transferred value and ok is true.
--- If the receive cannot finish without blocking, x is the zero Value and ok is false.
--- If the channel is closed, x is the zero value for the channel's element type and ok is false.
---@return reflectValue, boolean
function reflectValue:TryRecv() end

--- FieldByIndex returns the nested field corresponding to index.
--- It panics if evaluation requires stepping through a nil
--- pointer or a field that is not a struct.
---@param index any
---@return reflectValue
function reflectValue:FieldByIndex(index) end

--- Float returns v's underlying value, as a float64.
--- It panics if v's Kind is not Float32 or Float64
---@return number
function reflectValue:Float() end

--- NumMethod returns the number of methods in the value's method set.
---
--- For a non-interface type, it returns the number of exported methods.
---
--- For an interface type, it returns the number of exported and unexported methods.
---@return number
function reflectValue:NumMethod() end

--- OverflowComplex reports whether the complex128 x cannot be represented by v's type.
--- It panics if v's Kind is not Complex64 or Complex128.
---@param x any
---@return boolean
function reflectValue:OverflowComplex(x) end

--- Slice returns v[i:j].
--- It panics if v's Kind is not Array, Slice or String, or if v is an unaddressable array,
--- or if the indexes are out of bounds.
---@param i number
---@param j number
---@return reflectValue
function reflectValue:Slice(i, j) end

--- Type returns v's type.
---@return reflectType
function reflectValue:Type() end

--- SetIterKey assigns to v the key of iter's current map entry.
--- It is equivalent to v.Set(iter.Key()), but it avoids allocating a new Value.
--- As in Go, the key must be assignable to v's type and
--- must not be derived from an unexported field.
---@param iter reflectMapIter
function reflectValue:SetIterKey(iter) end

--- Method returns a function value corresponding to v's i'th method.
--- The arguments to a Call on the returned function should not include
--- a receiver; the returned function will always use v as the receiver.
--- Method panics if i is out of range or if v is a nil interface value.
---@param i number
---@return reflectValue
function reflectValue:Method(i) end

--- SetLen sets v's length to n.
--- It panics if v's Kind is not Slice or if n is negative or
--- greater than the capacity of the slice.
---@param n number
function reflectValue:SetLen(n) end

--- CanComplex reports whether Complex can be used without panicking.
---@return boolean
function reflectValue:CanComplex() end

--- InterfaceData returns a pair of unspecified uintptr values.
--- It panics if v's Kind is not Interface.
---
--- In earlier versions of Go, this function returned the interface's
--- value as a uintptr pair. As of Go 1.4, the implementation of
--- interface values precludes any defined use of InterfaceData.
---
--- Deprecated: The memory representation of interface values is not
--- compatible with InterfaceData.
---@return any
function reflectValue:InterfaceData() end

--- Pointer returns v's value as a uintptr.
--- It panics if v's Kind is not Chan, Func, Map, Pointer, Slice, or UnsafePointer.
---
--- If v's Kind is Func, the returned pointer is an underlying
--- code pointer, but not necessarily enough to identify a
--- single function uniquely. The only guarantee is that the
--- result is zero if and only if v is a nil func Value.
---
--- If v's Kind is Slice, the returned pointer is to the first
--- element of the slice. If the slice is nil the returned value
--- is 0.  If the slice is empty but non-nil the return value is non-zero.
---
--- It's preferred to use uintptr(Value.UnsafePointer()) to get the equivalent result.
---@return any
function reflectValue:Pointer() end

--- SliceHeader is the runtime representation of a slice.
--- It cannot be used safely or portably and its representation may
--- change in a later release.
--- Moreover, the Data field is not sufficient to guarantee the data
--- it references will not be garbage collected, so programs must keep
--- a separate, correctly typed pointer to the underlying data.
---
--- In new code, use unsafe.Slice or unsafe.SliceData instead.
---@class reflectSliceHeader
---@field Data any
---@field Len number
---@field Cap number
local reflectSliceHeader = {}

--- A SelectDir describes the communication direction of a select case.
---@class reflectSelectDir
local reflectSelectDir = {}

--- A StructField describes a single field in a struct.
---@class reflectStructField
---@field Name string
---@field PkgPath string
---@field Type reflectType
---@field Tag reflectStructTag
---@field Offset any
---@field Index any
---@field Anonymous boolean
local reflectStructField = {}

--- IsExported reports whether the field is exported.
---@return boolean
function reflectStructField:IsExported() end

--- Method represents a single method.
---@class reflectMethod
---@field Name string
---@field PkgPath string
---@field Type reflectType
---@field Func reflectValue
---@field Index number
local reflectMethod = {}

--- IsExported reports whether the method is exported.
---@return boolean
function reflectMethod:IsExported() end

--- A Kind represents the specific kind of type that a Type represents.
--- The zero Kind is not a valid kind.
---@class reflectKind
local reflectKind = {}

--- String returns the name of k.
---@return string
function reflectKind:String() end

--- A MapIter is an iterator for ranging over a map.
--- See Value.MapRange.
---@class reflectMapIter
local reflectMapIter = {}

--- Key returns the key of iter's current map entry.
---@return reflectValue
function reflectMapIter:Key() end

--- Value returns the value of iter's current map entry.
---@return reflectValue
function reflectMapIter:Value() end

--- Next advances the map iterator and reports whether there is another
--- entry. It returns false when iter is exhausted; subsequent
--- calls to Key, Value, or Next will panic.
---@return boolean
function reflectMapIter:Next() end

--- Reset modifies iter to iterate over v.
--- It panics if v's Kind is not Map and v is not the zero Value.
--- Reset(Value{}) causes iter to not to refer to any map,
--- which may allow the previously iterated-over map to be garbage collected.
---@param v reflectValue
function reflectMapIter:Reset(v) end

--- A ValueError occurs when a Value method is invoked on
--- a Value that does not support it. Such cases are documented
--- in the description of each method.
---@class reflectValueError
---@field Method string
---@field Kind reflectKind
local reflectValueError = {}


---@return string
function reflectValueError:Error() end

--- StringHeader is the runtime representation of a string.
--- It cannot be used safely or portably and its representation may
--- change in a later release.
--- Moreover, the Data field is not sufficient to guarantee the data
--- it references will not be garbage collected, so programs must keep
--- a separate, correctly typed pointer to the underlying data.
---
--- In new code, use unsafe.String or unsafe.StringData instead.
---@class reflectStringHeader
---@field Data any
---@field Len number
local reflectStringHeader = {}

--- A SelectCase describes a single case in a select operation.
--- The kind of case depends on Dir, the communication direction.
---
--- If Dir is SelectDefault, the case represents a default case.
--- Chan and Send must be zero Values.
---
--- If Dir is SelectSend, the case represents a send operation.
--- Normally Chan's underlying value must be a channel, and Send's underlying value must be
--- assignable to the channel's element type. As a special case, if Chan is a zero Value,
--- then the case is ignored, and the field Send will also be ignored and may be either zero
--- or non-zero.
---
--- If Dir is SelectRecv, the case represents a receive operation.
--- Normally Chan's underlying value must be a channel and Send must be a zero Value.
--- If Chan is a zero Value, then the case is ignored, but Send must still be a zero Value.
--- When a receive operation is selected, the received Value is returned by Select.
---@class reflectSelectCase
---@field Dir reflectSelectDir
---@field Chan reflectValue
---@field Send reflectValue
local reflectSelectCase = {}

--- ChanDir represents a channel type's direction.
---@class reflectChanDir
local reflectChanDir = {}


---@return string
function reflectChanDir:String() end

--- Type is the representation of a Go type.
---
--- Not all methods apply to all kinds of types. Restrictions,
--- if any, are noted in the documentation for each method.
--- Use the Kind method to find out the kind of type before
--- calling kind-specific methods. Calling a method
--- inappropriate to the kind of type causes a run-time panic.
---
--- Type values are comparable, such as with the == operator,
--- so they can be used as map keys.
--- Two Type values are equal if they represent identical types.
---@class reflectType
local reflectType = {}

--- A StructTag is the tag string in a struct field.
---
--- By convention, tag strings are a concatenation of
--- optionally space-separated key:"value" pairs.
--- Each key is a non-empty string consisting of non-control
--- characters other than space (U+0020 ' '), quote (U+0022 '"'),
--- and colon (U+003A ':').  Each value is quoted using U+0022 '"'
--- characters and Go string literal syntax.
---@class reflectStructTag
local reflectStructTag = {}

--- Get returns the value associated with key in the tag string.
--- If there is no such key in the tag, Get returns the empty string.
--- If the tag does not have the conventional format, the value
--- returned by Get is unspecified. To determine whether a tag is
--- explicitly set to the empty string, use Lookup.
---@param key string
---@return string
function reflectStructTag:Get(key) end

--- Lookup returns the value associated with key in the tag string.
--- If the key is present in the tag the value (which may be empty)
--- is returned. Otherwise the returned value will be the empty string.
--- The ok return value reports whether the value was explicitly set in
--- the tag string. If the tag does not have the conventional format,
--- the value returned by Lookup is unspecified.
---@param key string
---@return string, boolean
function reflectStructTag:Lookup(key) end
