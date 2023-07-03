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

---{{.reflectMakeChan}}
---@param typ reflectType
---@param buffer number
---@return reflectValue
function reflect.MakeChan(typ, buffer)
end

---{{.reflectMakeSlice}}
---@param typ reflectType
---@param len number
---@param cap number
---@return reflectValue
function reflect.MakeSlice(typ, len, cap)
end

---{{.reflectMakeFunc}}
---@param typ reflectType
---@param fn function
---@return reflectValue
function reflect.MakeFunc(typ, fn)
end

---{{.reflectChanOf}}
---@param dir reflectChanDir
---@param t reflectType
---@return reflectType
function reflect.ChanOf(dir, t)
end

---{{.reflectZero}}
---@param typ reflectType
---@return reflectValue
function reflect.Zero(typ)
end

---{{.reflectValueOf}}
---@param i any
---@return reflectValue
function reflect.ValueOf(i)
end

-- ---{{.reflectArenaNew}}
-- ---@param a arenaArena
-- ---@param typ reflectType
-- ---@return reflectValue
-- function reflect.ArenaNew(a, typ)
-- end

---{{.reflectDeepEqual}}
---@param x any
---@param y any
---@return boolean
function reflect.DeepEqual(x, y)
end

---{{.reflectPointerTo}}
---@param t reflectType
---@return reflectType
function reflect.PointerTo(t)
end

---{{.reflectAppendSlice}}
---@param s reflectValue
---@param t reflectValue
---@return reflectValue
function reflect.AppendSlice(s, t)
end

---{{.reflectStructOf}}
---@param fields any
---@return reflectType
function reflect.StructOf(fields)
end

---{{.reflectPtrTo}}
---@param t reflectType
---@return reflectType
function reflect.PtrTo(t)
end

---{{.reflectCopy}}
---@param dst reflectValue
---@param src reflectValue
---@return number
function reflect.Copy(dst, src)
end

---{{.reflectNew}}
---@param typ reflectType
---@return reflectValue
function reflect.New(typ)
end

---{{.reflectNewAt}}
---@param typ reflectType
---@param p unsafePointer
---@return reflectValue
function reflect.NewAt(typ, p)
end

---{{.reflectIndirect}}
---@param v reflectValue
---@return reflectValue
function reflect.Indirect(v)
end

---{{.reflectSliceOf}}
---@param t reflectType
---@return reflectType
function reflect.SliceOf(t)
end

---{{.reflectAppend}}
---@param s reflectValue
---@vararg reflectValue
---@return reflectValue
function reflect.Append(s, ...)
end

---{{.reflectVisibleFields}}
---@param t reflectType
---@return any
function reflect.VisibleFields(t)
end

---{{.reflectArrayOf}}
---@param length number
---@param elem reflectType
---@return reflectType
function reflect.ArrayOf(length, elem)
end

---{{.reflectMakeMapWithSize}}
---@param typ reflectType
---@param n number
---@return reflectValue
function reflect.MakeMapWithSize(typ, n)
end

---{{.reflectMapOf}}
---@param key reflectType
---@param elem reflectType
---@return reflectType
function reflect.MapOf(key, elem)
end

---{{.reflectTypeOf}}
---@param i any
---@return reflectType
function reflect.TypeOf(i)
end

---{{.reflectMakeMap}}
---@param typ reflectType
---@return reflectValue
function reflect.MakeMap(typ)
end

---{{.reflectSelect}}
---@param cases any
---@return number, reflectValue, boolean
function reflect.Select(cases)
end

---{{.reflectSwapper}}
---@param slice any
---@return any
function reflect.Swapper(slice)
end

---{{.reflectFuncOf}}
---@param in_ any
---@param out_ any
---@param variadic boolean
---@return reflectType
function reflect.FuncOf(in_, out_, variadic)
end

---@class reflectKind
local reflectKind = {}

---{{.reflectKindString}}
---@return string
function reflectKind:String()
end

---@class reflectMethod
---@field Name string
---@field PkgPath string
---@field Type reflectType
---@field Func reflectValue
---@field Index number
local reflectMethod = {}

---{{.reflectMethodIsExported}}
---@return boolean
function reflectMethod:IsExported()
end

---@class reflectStructField
---@field Name string
---@field PkgPath string
---@field Type reflectType
---@field Tag reflectStructTag
---@field Offset any
---@field Index any
---@field Anonymous boolean
local reflectStructField = {}

---{{.reflectStructFieldIsExported}}
---@return boolean
function reflectStructField:IsExported()
end

---@class reflectSelectCase
---@field Dir reflectSelectDir
---@field Chan reflectValue
---@field Send reflectValue
local reflectSelectCase = {}

---@class reflectValue
local reflectValue = {}

---{{.reflectValueCanFloat}}
---@return boolean
function reflectValue:CanFloat()
end

---{{.reflectValueCanSet}}
---@return boolean
function reflectValue:CanSet()
end

---{{.reflectValueTrySend}}
---@param x reflectValue
---@return boolean
function reflectValue:TrySend(x)
end

---{{.reflectValueMapKeys}}
---@return any
function reflectValue:MapKeys()
end

---{{.reflectValueFloat}}
---@return number
function reflectValue:Float()
end

---{{.reflectValueOverflowFloat}}
---@param x number
---@return boolean
function reflectValue:OverflowFloat(x)
end

---{{.reflectValuePointer}}
---@return any
function reflectValue:Pointer()
end

---{{.reflectValueCanUint}}
---@return boolean
function reflectValue:CanUint()
end

---{{.reflectValueCanAddr}}
---@return boolean
function reflectValue:CanAddr()
end

---{{.reflectValueField}}
---@param i number
---@return reflectValue
function reflectValue:Field(i)
end

---{{.reflectValueSetPointer}}
---@param x unsafePointer
function reflectValue:SetPointer(x)
end

---{{.reflectValueAddr}}
---@return reflectValue
function reflectValue:Addr()
end

---{{.reflectValueFieldByName}}
---@param name string
---@return reflectValue
function reflectValue:FieldByName(name)
end

---{{.reflectValueIndex}}
---@param i number
---@return reflectValue
function reflectValue:Index(i)
end

---{{.reflectValueIsNil}}
---@return boolean
function reflectValue:IsNil()
end

---{{.reflectValueNumMethod}}
---@return number
function reflectValue:NumMethod()
end

---{{.reflectValueSet}}
---@param x reflectValue
function reflectValue:Set(x)
end

---{{.reflectValueSetBytes}}
---@param x byte[]
function reflectValue:SetBytes(x)
end

---{{.reflectValueElem}}
---@return reflectValue
function reflectValue:Elem()
end

---{{.reflectValueCanConvert}}
---@param t reflectType
---@return boolean
function reflectValue:CanConvert(t)
end

---{{.reflectValueIsZero}}
---@return boolean
function reflectValue:IsZero()
end

---{{.reflectValueSetIterValue}}
---@param iter reflectMapIter
function reflectValue:SetIterValue(iter)
end

---{{.reflectValueMethod}}
---@param i number
---@return reflectValue
function reflectValue:Method(i)
end

---{{.reflectValueSend}}
---@param x reflectValue
function reflectValue:Send(x)
end

---{{.reflectValueCanInterface}}
---@return boolean
function reflectValue:CanInterface()
end

---{{.reflectValueInterface}}
---@return any
function reflectValue:Interface()
end

---{{.reflectValueKind}}
---@return reflectKind
function reflectValue:Kind()
end

---{{.reflectValueOverflowComplex}}
---@param x any
---@return boolean
function reflectValue:OverflowComplex(x)
end

---{{.reflectValueType}}
---@return reflectType
function reflectValue:Type()
end

---{{.reflectValueUint}}
---@return number
function reflectValue:Uint()
end

---{{.reflectValueFieldByIndexErr}}
---@param index any
---@return reflectValue, err
function reflectValue:FieldByIndexErr(index)
end

---{{.reflectValueSlice}}
---@param i number
---@param j number
---@return reflectValue
function reflectValue:Slice(i, j)
end

---{{.reflectValueConvert}}
---@param t reflectType
---@return reflectValue
function reflectValue:Convert(t)
end

---{{.reflectValueCanInt}}
---@return boolean
function reflectValue:CanInt()
end

---{{.reflectValueSlice3}}
---@param i number
---@param j number
---@param k number
---@return reflectValue
function reflectValue:Slice3(i, j, k)
end

---{{.reflectValueTryRecv}}
---@return reflectValue, boolean
function reflectValue:TryRecv()
end

---{{.reflectValueInt}}
---@return number
function reflectValue:Int()
end

---{{.reflectValueFieldByIndex}}
---@param index any
---@return reflectValue
function reflectValue:FieldByIndex(index)
end

---{{.reflectValueSetFloat}}
---@param x number
function reflectValue:SetFloat(x)
end

---{{.reflectValueSetInt}}
---@param x number
function reflectValue:SetInt(x)
end

---{{.reflectValueSetLen}}
---@param n number
function reflectValue:SetLen(n)
end

---{{.reflectValueSetString}}
---@param x string
function reflectValue:SetString(x)
end

---{{.reflectValueCallSlice}}
---@param in_ any
---@return any
function reflectValue:CallSlice(in_)
end

---{{.reflectValueLen}}
---@return number
function reflectValue:Len()
end

---{{.reflectValueMapIndex}}
---@param key reflectValue
---@return reflectValue
function reflectValue:MapIndex(key)
end

---{{.reflectValueSetIterKey}}
---@param iter reflectMapIter
function reflectValue:SetIterKey(iter)
end

---{{.reflectValueUnsafePointer}}
---@return unsafePointer
function reflectValue:UnsafePointer()
end

---{{.reflectValueBytes}}
---@return byte[]
function reflectValue:Bytes()
end

---{{.reflectValueIsValid}}
---@return boolean
function reflectValue:IsValid()
end

---{{.reflectValueSetZero}}
function reflectValue:SetZero()
end

---{{.reflectValueNumField}}
---@return number
function reflectValue:NumField()
end

---{{.reflectValueOverflowInt}}
---@param x number
---@return boolean
function reflectValue:OverflowInt(x)
end

---{{.reflectValueRecv}}
---@return reflectValue, boolean
function reflectValue:Recv()
end

---{{.reflectValueSetCap}}
---@param n number
function reflectValue:SetCap(n)
end

---{{.reflectValueSetMapIndex}}
---@param key reflectValue
---@param elem reflectValue
function reflectValue:SetMapIndex(key, elem)
end

---{{.reflectValueCanComplex}}
---@return boolean
function reflectValue:CanComplex()
end

---{{.reflectValueClose}}
function reflectValue:Close()
end

---{{.reflectValueComparable}}
---@return boolean
function reflectValue:Comparable()
end

---{{.reflectValueEqual}}
---@param u reflectValue
---@return boolean
function reflectValue:Equal(u)
end

---{{.reflectValueBool}}
---@return boolean
function reflectValue:Bool()
end

---{{.reflectValueMapRange}}
---@return reflectMapIter
function reflectValue:MapRange()
end

---{{.reflectValueSetComplex}}
---@param x any
function reflectValue:SetComplex(x)
end

---{{.reflectValueGrow}}
---@param n number
function reflectValue:Grow(n)
end

---{{.reflectValueCall}}
---@param in_ any
---@return any
function reflectValue:Call(in_)
end

---{{.reflectValueCap}}
---@return number
function reflectValue:Cap()
end

---{{.reflectValueFieldByNameFunc}}
---@param match any
---@return reflectValue
function reflectValue:FieldByNameFunc(match)
end

---{{.reflectValueOverflowUint}}
---@param x number
---@return boolean
function reflectValue:OverflowUint(x)
end

---{{.reflectValueSetBool}}
---@param x boolean
function reflectValue:SetBool(x)
end

---{{.reflectValueUnsafeAddr}}
---@return any
function reflectValue:UnsafeAddr()
end

---{{.reflectValueInterfaceData}}
---@return any
function reflectValue:InterfaceData()
end

---{{.reflectValueMethodByName}}
---@param name string
---@return reflectValue
function reflectValue:MethodByName(name)
end

---{{.reflectValueSetUint}}
---@param x number
function reflectValue:SetUint(x)
end

---{{.reflectValueString}}
---@return string
function reflectValue:String()
end

---{{.reflectValueComplex}}
---@return any
function reflectValue:Complex()
end

---@class reflectMapIter
local reflectMapIter = {}

---{{.reflectMapIterKey}}
---@return reflectValue
function reflectMapIter:Key()
end

---{{.reflectMapIterValue}}
---@return reflectValue
function reflectMapIter:Value()
end

---{{.reflectMapIterNext}}
---@return boolean
function reflectMapIter:Next()
end

---{{.reflectMapIterReset}}
---@param v reflectValue
function reflectMapIter:Reset(v)
end

---@class reflectStructTag
local reflectStructTag = {}

---{{.reflectStructTagGet}}
---@param key string
---@return string
function reflectStructTag:Get(key)
end

---{{.reflectStructTagLookup}}
---@param key string
---@return string, boolean
function reflectStructTag:Lookup(key)
end

---@class reflectType
local reflectType = {}

---@class reflectChanDir
local reflectChanDir = {}

---{{.reflectChanDirString}}
---@return string
function reflectChanDir:String()
end

---@class reflectSliceHeader
---@field Data any
---@field Len number
---@field Cap number
local reflectSliceHeader = {}

---@class reflectSelectDir
local reflectSelectDir = {}

---@class reflectValueError
---@field Method string
---@field Kind reflectKind
local reflectValueError = {}

---{{.reflectValueErrorError}}
---@return string
function reflectValueError:Error()
end

---@class reflectStringHeader
---@field Data any
---@field Len number
local reflectStringHeader = {}
