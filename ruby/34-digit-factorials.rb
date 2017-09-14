#!/usr/bin/env ruby

#TODO find a factorial method for Ruby
class Integer
 def factorial
   (1..self).reduce(1, :*)
 end
end


max, answer = 99999, 0

(3..max).each do |i|
 sum_factorials = i.to_s.chars \
   .inject(0){|sum, char| sum + (char.to_i.factorial)}

 answer += i if (i == sum_factorials)
end

puts answer
