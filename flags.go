package main

import (
    "fmt"
    "os"
)

type Flag struct {
    Name   string
    Short  string
    Long   string
}

type StringFlag struct {
    Flag   Flag
    Value  string
}

type BoolFlag struct {
    Flag   Flag
    Value  bool
}

type IntFlag struct {
    Flag   Flag
    Value  int
}

type FlagList struct {
    BoolFlags   []BoolFlag
    IntFlags    []IntFlag
    StringFlags []StringFlag
}

func NewFlag(name, short, long string) Flag {
    f := Flag{
        Name: name,
        Short: short,
        Long: long,
    }

    return f
}


func NewBoolFlag(flag Flag, value bool) BoolFlag {
    bf := BoolFlag{
        Flag: flag,
        Value: value,
    }

    return bf
}

func NewIntFlag(flag Flag, value int) IntFlag {
    intf := IntFlag{
        Flag: flag,
        Value: value,
    }

    return intf
}

func NewStringFlag(flag Flag, value string) StringFlag {
    sf := StringFlag{
        Flag: flag,
        Value: value,
    }

    return sf
}

func GetBoolFlagValue(name string, validFlags FlagList) bool {
    for i := 0; i < len(validFlags.BoolFlags); i++ {
        if name == validFlags.BoolFlags[i].Flag.Name {
            return validFlags.BoolFlags[i].Value
        }
    }

    fmt.Printf("Flag \"%s\" doesn't exist.\n", name)
    os.Exit(1)
    
    return false
}

func GetIntFlagValue(name string, validFlags FlagList) int {
    for i := 0; i < len(validFlags.IntFlags); i++ {
        if name == validFlags.IntFlags[i].Flag.Name {
            return validFlags.IntFlags[i].Value
        }
    }

    fmt.Printf("Flag \"%s\" doesn't exist.\n", name)
    os.Exit(1)

    return 0
}

func GetStringFlagValue(name string, validFlags FlagList) string {
    for i := 0; i < len(validFlags.StringFlags); i++ {
        if name == validFlags.StringFlags[i].Flag.Name {
            return validFlags.StringFlags[i].Value
        }
    }
    
    fmt.Printf("Flag \"%s\" doesn't exist.\n", name)
    os.Exit(1)
    
    return ""
}

func GetFlagType(flagShortOrLong string, validFlags FlagList) string {
    for i := 0; i < len(validFlags.BoolFlags); i++ {
        if flagShortOrLong == validFlags.BoolFlags[i].Flag.Short ||
           flagShortOrLong == validFlags.BoolFlags[i].Flag.Long {
            return "bool"
        }
    }
    
    for i := 0; i < len(validFlags.IntFlags); i++ {
        if flagShortOrLong == validFlags.IntFlags[i].Flag.Short ||
           flagShortOrLong == validFlags.IntFlags[i].Flag.Long {
            return "int"
        }
    }
    
    for i := 0; i < len(validFlags.StringFlags); i++ {
        if flagShortOrLong == validFlags.StringFlags[i].Flag.Short ||
           flagShortOrLong == validFlags.StringFlags[i].Flag.Long {
            return "string"
        }
    }

    return ""
}


func IsFlagValid(flagName string, validFlags FlagList) bool {
    if GetFlagType(flagName, validFlags) != "none" {
        return true
    }

    return false
}
