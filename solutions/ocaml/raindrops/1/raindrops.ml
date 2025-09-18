let raindrop n =
  let pling = if n mod 3 = 0 then "Pling" else "" in
  let plang = if n mod 5 = 0 then "Plang" else "" in
  let plong = if n mod 7 = 0 then "Plong" else "" in
  let result = pling ^ plang ^ plong in
  if result = "" then string_of_int n else result
