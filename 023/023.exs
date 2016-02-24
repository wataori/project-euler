sum = fn(n) ->
        sqrt = :math.sqrt(n)
        Enum.reduce 1..(trunc(sqrt)), fn(x, sum) ->
          cond do
            x == 1 ->
              sum + x
            rem(n, x) == 0 && :math.pow(x, 2) == n ->
              sum + x
            rem(n, x) == 0 ->
              sum + x + div(n, x)
            true ->
              sum
          end
        end
      end

1..28123 |> Enum.filter(fn(x) -> sum.(x) < x end) |> Enum.reduce(0, &+/2) |> IO.puts
