module Solve1
  def solve
    puts (0..Float::INFINITY).step(20).lazy.select{|n| detect(n.to_i, 20)}.first
  end

  def detect(num, divisor)
    return true if divisor == 2 && num % divisor == 0
    return false if num < 20
    num % divisor == 0 ? detect(num, divisor - 1) : false
  end
end
