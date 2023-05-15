package archive

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

type TypeArchive struct {
	types map[string]*TypeRecord
}

func (*TypeArchive) Save(name string, record *TypeRecord) {
	archive.types[name] = record
}

func (archive *TypeArchive) BatchSave(t map[string]*TypeRecord) {
	archive.types = t
}

func (archive *TypeArchive) Lookup(typeName string) *TypeRecord {
	if record, ok := archive.types[typeName]; ok {
		return record
	}
	return archive.types["any"]
}

type TypeRecord struct {
	valueType lua.LValueType
}

func (record *TypeRecord) CheckType() lua.LValueType {
	return record.valueType
}

func (record *TypeRecord) Type(ident string) string {
	switch record.valueType {
	case lua.LTNil:
		return "lua.LNil"
	case lua.LTBool:
		return fmt.Sprintf("if %s {\n  l.Push(lua.LTrue)\n} else {\n  l.Push(lua.LFalse)\n}}", ident)
	case lua.LTNumber:
		return fmt.Sprintf("lua.LNumber(%s)", ident)
	case lua.LTString:
		return fmt.Sprintf("lua.LString(%s)", ident)
	case lua.LTFunction:
		fallthrough
	case lua.LTUserData:
		fallthrough
	case lua.LTThread:
		fallthrough
	case lua.LTTable:
		fallthrough
	case lua.LTChannel:
		fallthrough
	default:
		return fmt.Sprintf("luar.New(l, %s)", ident)
	}
}

func (record *TypeRecord) Check(i int) string {
	switch record.valueType {
	case lua.LTNil:
		return ""
	case lua.LTBool:
		return fmt.Sprintf("l.CheckBool(%d)", i)
	case lua.LTNumber:
		return fmt.Sprintf("l.CheckNumber(%d)", i)
	case lua.LTString:
		return fmt.Sprintf("l.CheckString(%d)", i)
	case lua.LTFunction:
		return fmt.Sprintf("l.CheckFunction(%d)", i)
	case lua.LTUserData:
		return fmt.Sprintf("l.CheckUserData(%d)", i)
	case lua.LTThread:
		fallthrough
	case lua.LTTable:
		fallthrough
	case lua.LTChannel:
		fallthrough
	default:
		return fmt.Sprintf("l.CheckAny(%d)", i)
	}
}

var archive *TypeArchive = &TypeArchive{}

func GetArchive() *TypeArchive {
	return archive
}

func init() {
	archive.BatchSave(map[string]*TypeRecord{
		"string": {valueType: lua.LTString},
		"int":    {valueType: lua.LTNumber},
		"bool":   {valueType: lua.LTBool},
		"float":  {valueType: lua.LTNumber},
		"any":    {valueType: lua.LTUserData},
	})
}
