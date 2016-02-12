# multi = 1..1001 |> Enum.reduce(fn(n, mul) -> 2*mul end)
multi = trunc :math.pow(2, 1000)

ans = Integer.digits(multi) |> Enum.reduce(fn(n, sum) -> n+sum end)

IO.puts ans

# IO.puts Integer.digits(trunc :math.pow(2, 1000)) |> Enum.reduce(fn(n, sum) -> n+sum end) #oneliner
