
let rec diff = function
    | a :: b :: tail -> (b - a) :: diff (b :: tail)
    | _ -> []

let rec all_zeros = function
  | x :: tail -> x == 0 && all_zeros tail
  | _ -> true

let last arr = List.hd (List.rev arr)

let rec forecast arr =
  if all_zeros arr then (0, 0)
  else
    let (a, b) = forecast (diff arr) in
    ((List.hd arr) - a, (last arr + b))

let () =
  let (a, b) = forecast [10;  13;  16;  21;  30] in
  let result = (Int.to_string a) ^ " " ^ (Int.to_string b) in
  print_string result
