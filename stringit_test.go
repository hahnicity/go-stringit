package stringit

import (
    "reflect"
    "testing"
)


var (
    test1    int      = 1
    test2    int8     = 1
    test3    int16    = 1
    test4    int32    = 1
    test5    int64    = 1
    test6    uint     = 1
    test7    uint8    = 1
    test8    uint16   = 1
    test9    uint32   = 1
    test10   uint64   = 1
    test11   float32  = 1.1
    test12   float64  = 1.1
    test13   bool     = true
    test14   string   = "baz"
    tests             = []interface{}{
        test1, test2, test3, test4, test5, 
        test6, test7, test8, test9, test10, 
        test11, test12, test13, test14,
    }
    expected          = []string{
        "1", "1", "1", "1", "1",
        "1", "1", "1", "1", "1",
        "1.10000", "1.10000", "true", "baz",
    }
)

func TestAllTypes(t *testing.T) {
    formatStr := "{}"
    result    := ""
    for i, test := range tests {
        result = Format(formatStr, test)
        if result != expected[i] {
            t.Errorf(
                "Testing has failed for %s, expected: %s, result %s", 
                reflect.TypeOf(test).String(),
                expected[i],
                result,   
            )    
        }
    }
}
