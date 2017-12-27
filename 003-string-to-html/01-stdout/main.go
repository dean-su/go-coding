package main

import "fmt"

func main() {
	name := "Dean Su"
	tpl := `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>emmet</title>
</head>
<body>
<!--p{example from emmet}-->
<p>Example from emmet</p>
<h1 title="item1">`+ name +`</h1>
<h2 title="item2">Header 2</h2>
<h3 title="item3">Header 3</h3>

</body>
</html>
	`
	fmt.Println(tpl)
}
