package main

func getValidFlags() FlagList {
    verbose := NewBoolFlag(NewFlag("verbose", "-v", "--verbose"), false)
    vverbose := NewBoolFlag(NewFlag("vverbose", "-vv", "--vverbose"), false)
    limit := NewIntFlag(NewFlag("limit", "-l", "--limit"), 0)
    falsee := NewBoolFlag(NewFlag("false", "-f", "--false"), false)

    boolFlags := []BoolFlag{falsee, verbose, vverbose}
    intFlags := []IntFlag{limit}
    stringFlags := []StringFlag{}

    flagList := FlagList{
        BoolFlags: boolFlags,
        IntFlags: intFlags, 
        StringFlags: stringFlags,
    } 

    return flagList 
}

