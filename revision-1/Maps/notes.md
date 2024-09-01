## Difference b/w Maps and Structs
- Maps are refrence types (just like slices ), Structs are value types
- In Maps, type of all keys and values must be known before hand. In Structs, we don't need to know their types
- In Structs, we must know all the properties of a struct beforehand. In Maps, we don't need to know all key-values pairs beforehand. We can add/remove them dynamically
- Logically, Structs are used to represent some "thing" with multiple different properites. Maps are used to represent a collection of closely related values
- Maps support indexing of keys, meaning we can iterate over them. Structs are not iteratable
