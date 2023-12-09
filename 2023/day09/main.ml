
let rec diff = function
    | a :: b :: tail ->
        (b - a) :: diff (b :: tail)
    | _ -> []

let rec all_zeros = function
  | 0 :: tail -> all_zeros tail
  | [] -> true
  | _ -> false

let last arr = List.hd (List.rev arr)

let rec forecast arr =
  if all_zeros arr
    then 0
  else
    last arr + forecast (diff arr)

let () =
  forecast [10;  13;  16;  21;  30] |>
  Int.to_string |>
  print_string ;

  print_newline () ;

  forecast (List.rev [10;  13;  16;  21;  30]) |>
  Int.to_string |>
  print_string ;

  print_newline () ;

