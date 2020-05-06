package main

type Flag struct {
    Action string
    Name   string
    Short  string
    Long   string
    Value  string
}

func constructFlag(name, action, short, long, value string) Flag {
    var new_flag Flag = Flag{
        Action: action,
        Name: name,
        Short: short,
        Long: long,
        Value: value,
    }

    return new_flag
}

func getValidFlags() []Flag {
    var verbose Flag = constructFlag("verbose", "exists", "-v", "--verbose", "")
    var vverbose Flag = constructFlag("vverbose", "exists", "-vv", "--vverbose", "")
    var limit Flag = constructFlag("limit", "store", "-l", "--limit", "")

    var flags []Flag = []Flag{limit, verbose, vverbose}

    return flags 
}

func GetFlag(flag_name string, flag_list []Flag) *Flag {
    for i := 0; i < len(flag_list); i++ {
        if compareStringsCaseInsensitive(flag_name, flag_list[i].Name) {
            return &flag_list[i]
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

