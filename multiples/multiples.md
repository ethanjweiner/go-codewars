# Multiples

# Algorithm

N.B. Guard is unnecessary

Given an `upperBound`:
- Initialize a `sum`
- Iterate from 3 until `upperBound` (exclusive), tracking `num`:
  - If `num` is divisible by 3 or 5:
    - Add `num` to `sum`
- Return `sum`