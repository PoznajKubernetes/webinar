counts = Hash.new(0)

100.times do
  output = `curl -s http://localhost/pet`
  counts[output] += 1
end

puts counts
