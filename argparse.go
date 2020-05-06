package main

import (
    "fmt"
    "strconv"
    "os"
)

type Flag struct {
    Action string
    Name   string
    Short  string
    Long   string
    Value  string
}

func NewFlag(name, action, short, long, value string) Flag {
    var new_flag Flag = Flag{
        Action: action,
        Name: name,
        Short: short,
        Long: long,
        Value: value,
    }

    return new_flag
}

func GetFlag(flag_name string, flag_list *[]Flag) *Flag {
    for i := 0; i < len(*flag_list); i++ {
        if compareStringsCaseInsensitive(flag_name, (*flag_list)[i].Name) {
            return &(*flag_list)[i]
        }
    }

    return nil 
}

func GetFlagValue(f *Flag) string {
    return (*f).Value
}

func ValidateFlag(flag_name string) (is_valid bool, flag Flag) {
    var valid_flags []Flag = getValidFlags()

    for i := 0; i < len(valid_flags); i++ {
        if flag_name == valid_flags[i].Short || flag_name == valid_flags[i].Long {
            is_valid = true
            flag = valid_flags[i]
            return
        }
    }

    return
}

/****************************************************************************** 
 *
 * Project specific functions
 *
******************************************************************************/
func getValidFlags() []Flag {
    var verbose Flag = NewFlag("verbose", "exists", "-v", "--verbose", "")
    var vverbose Flag = NewFlag("vverbose", "exists", "-vv", "--vverbose", "")
    var limit Flag = NewFlag("limit", "store", "-l", "--limit", "")

    var flags []Flag = []Flag{limit, verbose, vverbose}

    return flags 
}

func getExistsFlagValue(flag_name string, flags *[]Flag) bool {
    if len(*flags) == 0 {
        return false
    }
    
    var flag *Flag = GetFlag(flag_name, flags)

    if flag != nil {
        return true
    }
    
    return false
}

func getStoreIntFlagValue(flag_name string, flags *[]Flag) int {
    if len(*flags) == 0 {
        return 0
    }
    
    var flag *Flag = GetFlag(flag_name, flags)
    
    if flag != nil {
        var flag_value string = GetFlagValue(flag)
        value, err := strconv.ParseInt(flag_value, 10, 0)
        if err != nil {
            fmt.Printf("Limit value must be an integer. Value: %v\n", flag_value)
            fmt.Printf("%v\n", err)
            os.Exit(1)
        }

        return int(value)
    }

    return 0
}
