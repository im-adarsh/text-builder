# text-builder
This mini util takes in the text template, replace the text template with variables and return the actual text.

Usage : 

```
	str := txtBuilder.NewTextBuilder().Template("hello %v").Params([]string{"world"}).GetText()
	fmt.Println(str)
```

