package stringit

import (
    "reflect"
    "strconv"
    "strings"
)


func Format(formatStr string, a... interface{}) string {
    split := strings.SplitAfter(formatStr, "{}")
    for i, arg := range a {
        split[i] = formatPiece(split[i], arg)
    }
    return strings.Join(split, "")
}

func formatPiece(piece string, arg interface{}) string {
    switch reflect.TypeOf(arg).String() {
    case "string":
        str, _ := arg.(string)
        return strings.Replace(piece, "{}", str, 1)
    case "int":
        i, _ := arg.(int)
        return strings.Replace(piece, "{}", strconv.Itoa(i), 1)
    case "uint":
        i, _ := arg.(uint64)
        return strings.Replace(piece, "{}", strconv.FormatUint(i, 10), 1)
    case "float64":
        f, _ := arg.(float64)
        return strings.Replace(piece, "{}", strconv.FormatFloat(f, 'f', 2, 64), 1)
    
    case "bool":
        b, _ := arg.(bool)
        return strings.Replace(piece, "{}", strconv.FormatBool(b), 1)
    }
    return ""
}
