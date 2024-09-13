### maps
- representation of a hash table in go
- `var ages map[string]int` declared using `map[K]V`
- K is some comparable type which can be compared using == operator (string, int) avoid float because float comaprison is scuffed
- A struct can also be used as a Key, since they're also comparable types 
- Reason being the map needs to test if the given key is equal to the existing key or not
- V has to type restriction

**nil maps**
- maps declared like this `var ages map[string]int` are not only empty but they're also nil
- maps declared using make() function `ages := make(make[string]int)` are just empty but they have some memory allocalted to them so they're not nil
- if a map is nil, storing value will cause panic, therefore always make sure to allocate memory to it before storing values

- we cannot take address of a map element. `addr = &ages["bob"]` because addresses of map element are not fixed. As the map grows, they could be rehashed and their previous mem. addr would become invalid.

**accessing a map element by subscripting**
- `a := ages['alice']` doing this always produce some value and never an error (even if the element doesn't exist in the map). If the elment doesn't exist it returns an empty value of the corresponding type
- `a , ok := ages['alice']`
- The second variable ok is a boolean , it tells if the element in exists in the map or not
- Because if the elemnts are of int type, by default they all will have a 0 empty value, therefore it becomes important to check if the actual value is 0 or is it due to the face that the element never exisited in the map
-
