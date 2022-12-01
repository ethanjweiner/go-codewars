# Two To One

# Problem

- Input: Two strings, `str1` and `str2`
- Output: A singular string containing all distinct letters from both strings,
in sorted order

# Example

# Algorithm

## Sort later

Given `str1` and `str2`:
- Declare a rune slice to store characters
- Iterate over the strings, adding non-duplicates (includes?)
- Sort the rune slice
- Join the sorted slice & return

## Using string library

Given `str1` and `str2`:
- Initialize slice of all chars in strings
- Sort the slice (string sort)
- Remove duplicates
- Return string w/ duplicates removed

## Remove duplicates

Given a slice of `chars`:
- Initialize a `uniqueString` to empty
- Iterate over `chars`:
  - if `uniqueString` contains current char -> don't include
- Return `uniqueString`ß