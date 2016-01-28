require 'prime'
num = 600851475143
upto = Math.sqrt(num).to_i
puts Prime.each(upto).select{|n| num % n == 0}.last
