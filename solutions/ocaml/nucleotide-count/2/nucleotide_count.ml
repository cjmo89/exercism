open Base

let empty = Map.empty (module Char)
let valid_nucleotides = Set.of_list (module Char) [ 'A'; 'G'; 'T'; 'C' ]
let is_valid_nucleotide c = Set.mem valid_nucleotides c

let count_nucleotide s c =
  if not (is_valid_nucleotide c) then Error c
  else
    let len = String.length s in
    let rec count_char i acc =
      if i >= len then Ok acc
      else if not (is_valid_nucleotide s.[i]) then Error s.[i]
      else if Char.equal s.[i] c then count_char (i + 1) (acc + 1)
      else count_char (i + 1) acc
    in
    count_char 0 0

let count_nucleotides s =
  let len = String.length s in
  if len = 0 then Ok empty
  else
    let rec build_counts i map =
      if i >= len then Ok map
      else
        let c = s.[i] in
        if not (is_valid_nucleotide c) then Error c
        else
          let new_count =
            match Map.find map c with Some count -> count + 1 | None -> 1
          in
          let updated_map = Map.set map ~key:c ~data:new_count in
          build_counts (i + 1) updated_map
    in
    build_counts 0 empty
