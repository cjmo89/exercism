let rec sum n = if n <= 1 then n else n + sum (n - 1)

let square_of_sum n =
  let simple_sum = sum n in
  simple_sum * simple_sum

let rec sum_of_squares n = if n <= 1 then n else (n * n) + sum_of_squares (n - 1)
let difference_of_squares n = square_of_sum n - sum_of_squares n
