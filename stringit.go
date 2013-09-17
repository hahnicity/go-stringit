package stringit

import (
    "reflect"
    "strconv"
    "strings"
)

type Formatter struct {
    intBase, uintBase, floatPrec int
}

func (f *Formatter) FormatStr(formatStr string, a... interface{}) string {
    split := strings.SplitAfter(formatStr, "{}")
    for i, arg := range a {
        split[i] = f.formatPiece(split[i], arg)
    }
    return strings.Join(split, "")
}

func (f *Formatter) formatPiece(piece string, arg interface{}) string {
    switch reflect.TypeOf(arg).String() {
    case "string":
        str, _ := arg.(string)
        return strings.Replace(piece, "{}", str, 1)
    case "int":
        i, _ := arg.(int)
        return strings.Replace(piece, "{}", strconv.FormatInt(int64(i), f.intBase), 1)
    case "int8":
        i, _ := arg.(int8)
        return strings.Replace(piece, "{}", strconv.FormatInt(int64(i), f.intBase), 1)
    case "int16":
        i, _ := arg.(int16)
        return strings.Replace(piece, "{}", strconv.FormatInt(int64(i), f.intBase), 1)
    case "int32":
        i, _ := arg.(int32)
        return strings.Replace(piece, "{}", strconv.FormatInt(int64(i), f.intBase), 1)
    case "int64":
        i, _ := arg.(int64)
        return strings.Replace(piece, "{}", strconv.FormatInt(i, f.intBase), 1)
    case "uint":
        i, _ := arg.(uint)
        return strings.Replace(piece, "{}", strconv.FormatUint(uint64(i), f.uintBase), 1)
    case "uint8":
        i, _ := arg.(uint8)
        return strings.Replace(piece, "{}", strconv.FormatUint(uint64(i), f.uintBase), 1)
    case "uint16":
        i, _ := arg.(uint16)
        return strings.Replace(piece, "{}", strconv.FormatUint(uint64(i), f.uintBase), 1)
    case "uint32":
        i, _ := arg.(uint32)
        return strings.Replace(piece, "{}", strconv.FormatUint(uint64(i), f.uintBase), 1)
    case "uint64":
        i, _ := arg.(uint64)
        return strings.Replace(piece, "{}", strconv.FormatUint(i, f.uintBase), 1)
    case "float32":
        fl, _ := arg.(float32)
        return strings.Replace(piece, "{}", strconv.FormatFloat(float64(fl), 'f', f.floatPrec, 32), 1)
    case "float64":
        fl, _ := arg.(float64)
        return strings.Replace(piece, "{}", strconv.FormatFloat(fl, 'f', f.floatPrec, 64), 1)
    case "bool":
        b, _ := arg.(bool)
        return strings.Replace(piece, "{}", strconv.FormatBool(b), 1)
    }
    return ""
}

// Given a string with brackets like "{} foo {}" and a number of inputs equivalent
// to the number of brackets in the string, like "too", "baz", this function will
// return a string such that an item is matched with its corresponding bracket.
// So the function stringit.Format("{} foo {}", "too", "baz") wil yield "too foo baz"
func Format(formatStr string, a... interface{}) string {
    f := &Formatter{
        intBase: 10,
        uintBase: 10,
        floatPrec: 5,    
    }
    return f.FormatStr(formatStr, a...)
}

