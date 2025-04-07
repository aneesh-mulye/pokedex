# pokedex
boot.dev pokedex project

## Notes on various things

Wow, WTF, Golang? I have to do this bizarrely convoluted way of going around
and updating fields in a map? But more importantly, WTF if up the the circular
reference logic? There's perfectly valid code that should compile: a map of
command names to commands structs, wherein each struct has a name, a
description, and a callback field that's that command's callback (of type
`func() error`; and one command callback which references this command registry
map to generate the name-description pairs for display in a `help` command.
This apparently doesn't work, because why even what?!; WTF, Go?
