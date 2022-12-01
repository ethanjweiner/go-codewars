# Digital Root

## Problem

- Input: An `integer`
- Output: The recursive sum of the digits in `integer`
  - Recursively fold digits until left w/ one-digit integer

## Algorithm

Given an `integer`:
- If the `integer` is one digit: return the digit
- Compute the `sum` of the digits in `integer`
- Compute the `digitalRoot` of the new `sum`

### Is integer one digit?

Given an `integer`:
- Return true if the `integer` / 10 (as float) is less than 1
- Else return false

### Digit Sum

1523

quotient / 10 | remainder % 10
--------------+----------------
152             3
15              2
1               5
0               1

Given an `integer`:
- Initialize a `remaining` to `integer`
- Initialize a `sum` to 0
- While `remaining` is greater than zero:
  - Add `remaining % 10` to `sum`
  - Reassign `remaining` to `remaining / 10` (coerced to integer)
- Return the `sum`


## Implementation

- Determining whether digit is length 1: Divide by 10 -> less than 0?