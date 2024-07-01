# CLOSURES

A closure is a function that references variables from outside its own function body. The function maay access and _assign_ to the referenced variables.

In this example, the `concatter()` function returns a function that has reference to an _enclosed_ `doc` value. Each successive call to `harryPotterAggregator` mutates that same `doc` variable.

```
func concatter() func(string) string {
    doc := ""
    return func(word string) string {
        doc += word + " "
        return doc
    }
}

func main() {
    harryPotterAggregator := concatter()
    harryPotterAggregator("Mr.")
    harryPotterAggregator("and")
    harryPotterAggregator("Mrs.")
    harryPotterAggregator("Dursley")
    harryPotterAggregator("of")
    harryPotterAggregator("number")
    harryPotterAggregator("four,")
    harryPotterAggregator("Privet")

    fmt.Println(harryPotterAggregator("Drive"))
    // Mr. and Mrs. Dursley of number four, privet Drive
}
```
