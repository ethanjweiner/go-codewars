# Find Odd Int

## Ideas

- [More efficient] Find counts, iterate through counts, return first integer w/ odd counts
  - Finding counts: `n`
  - Iterate through counts: `n`
  - Worst Case: `O(n)`
- [Less efficient] For each integer in array: `n`
  - Determine the count for that integer: `n`
  - Return that integer if it is odd
  - Worst Case: `O(n^2)`

## Algorithm

Given an array (or slice?) of `integers`:
- Create `integerCounts` of all `integers` ##
- Iterate through `integerCounts`
  - Return first integer w/ an odd count

## Tallying

Given a `slice` of any type:
- Initialize a map (of capacity `len(slice)`)
- For each `element` in `slice`:
  - Increment map at element or set to zero
- Return the map