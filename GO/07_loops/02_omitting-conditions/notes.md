# OMITTING CONDITIONS FROM A FOR LOOP IN GO

Loops in Go can omit secetions of a for loop. For example, the `CONDITION`(middle part) can be omitted which causes the loop to run forever.

    for INITIAL; ; AFTER {
        // do something forever
    }
