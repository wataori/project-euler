20..1 |> Enum.reduce(fn(n, fact) -> fact*n end) |> Integer.digits |> Enum.reduce(fn(n, sum) -> sum+n end) |> IO.puts
