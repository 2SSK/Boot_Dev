# STANDARD IN

**"Standard Input"**, usually called "standard in" or "stdin", is the default place where programs read their input. It's just a stream of data that programs can read from as they run.

All programming languages have a simple way to read from stdin. In Python, it's the **input** function:

    # execution stops untill the user types
    # something (in this case "Lane") and presses enter
    name = input("What is your name? ")

    print("Hello,", name)
    # Hello, Lane!
