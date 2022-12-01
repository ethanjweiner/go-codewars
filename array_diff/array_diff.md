# Array Diff

## Problem

- Input: Slices `a` and `b`
- Output: A new slice w/ all elements from `b` removed from `a`
- *All* instances of elements in `b` which are also in `a` should be removed

## Ideas

- Iterate through every element in `b`; then remove every one in `a`
- Iterate through every element in `a`; remove if in `b`
- Iterate through every element in `a`, include in new slice if not in `b`**

## Examples

See test cases