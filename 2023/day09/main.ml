
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

let solve arr =
  forecast arr |>
  Int.to_string |>
  print_string ;

  print_newline () ;

  forecast (List.rev arr) |>
  Int.to_string |>
  print_string ;

  print_newline ()

let () =
  solve [10;  13;  16;  21;  30]
