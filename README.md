# Optional Go

<p align="center">
   <img src="/resources/optional_gopher.png" alt="Optional Gopher"/>
</p>

```
optString := optional.Of("foo")

optNum := optional.Of(5)

if ok, num := optNum.Value(); ok {
	fmt.Println("we have a value of %d", num)
}
```