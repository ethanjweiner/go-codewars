# Reverse Large Words

## Algorithm

Given a `string`:
- Initialize a slice of `words`
- Iterate over a string slice of the string split by space
  - If word is greater than 5 in length:
    - Reverse the word + append to `words`
  - Otherwise, append the word to `words`
- Join the words with spaces

### Reverse the word

Given a `word`:
- Transform word to a `[]string`
- Reverse the strings (`sort.Reverse`)
- Join and return

## Implementation

2 options for slice population:
1. Create new slice of words
2. Mutate existing slice of words