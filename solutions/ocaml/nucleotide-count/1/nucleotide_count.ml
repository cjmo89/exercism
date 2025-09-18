open Base

let empty = Map.empty (module Char)
let valid_nuceleotides = [ 'A'; 'G'; 'T'; 'C' ]

let count_nucleotide (s : string) (c : char) =
  if not (List.mem valid_nuceleotides c ~equal:Char.equal) then Error c
  else
    let rec traverse (i : int) (acc : int) =
      if i < 0 then Ok acc
      else if Char.equal s.[i] c then traverse (i - 1) (acc + 1)
      else if not (List.mem valid_nuceleotides s.[i] ~equal:Char.equal) then
        Error s.[i]
      else traverse (i - 1) acc
    in
    traverse (String.length s - 1) 0

let count_nucleotides s =
  if String.length s = 0 then Ok empty
  else
    let rec aux i m =
      if i < 0 then Ok m
      else if List.mem valid_nuceleotides s.[i] ~equal:Char.equal then
        let map =
          Map.set m ~key:s.[i]
            ~data:
              (match count_nucleotide s s.[i] with Error _e -> 0 | Ok v -> v)
        in
        aux (i - 1) map
      else Error s.[i]
    in
    aux (String.length s - 1) empty
