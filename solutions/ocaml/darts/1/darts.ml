let lies_inside_circle x y r = (x *. x) +. (y *. y) <= r *. r

let score (x : float) (y : float) : int =
  let outer_radius = 10.0 in
  let middle_radius = 5.0 in
  let inner_radius = 1.0 in
  if lies_inside_circle x y inner_radius then 10
  else if lies_inside_circle x y middle_radius then 5
  else if lies_inside_circle x y outer_radius then 1
  else 0
