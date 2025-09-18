type nucleotide = A | C | G | T

let hamming_distance s1 s2 =
  if List.length s1 != List.length s2 then
    Error "strands must be of equal length"
  else
    let rec compare l1 l2 acc =
      if List.is_empty l1 then acc
      else if List.hd l1 != List.hd l2 then
        compare (List.tl l1) (List.tl l2) acc + 1
      else compare (List.tl l1) (List.tl l2) acc
    in
    Ok (compare s1 s2 0)
