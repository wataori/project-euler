# time = Time.now
@results = []
@from = 100 * 100
@to = 999 * 999

def devide_each_num num
  # return num if num.zero? || Math.log10(num).to_i.zero?
  return num if num < 10
  [devide_each_num((num / 10).to_i), num % 10]
end

def mirrored? nums
  return true if nums.length <= 1
  return false unless nums[0] == nums[-1]
  mirrored?(nums[1..-2])
end

@from.upto(@to).each do |n|
  each_nums = devide_each_num(n).flatten
  @results << n if mirrored?(each_nums)
end

puts @results.last
# puts "#{Time.now - time}"
