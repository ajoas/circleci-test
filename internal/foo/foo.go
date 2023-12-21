package foo

func SubFoo(x int) string {
  if x == 7 {
    return "Seven"
  }
  return "Sad"
}

func SubBar(x int) string {
  if x == 3 {
    return "Three"
  }
  return "Sad"
}

func Foo(input int) string {
  subfoo := SubFoo(input)
  return subfoo
}

func Bar(input int) string {
  subbar := SubFoo(input)
  return subbar
}
