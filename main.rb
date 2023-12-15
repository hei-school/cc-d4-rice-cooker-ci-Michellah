# frozen_string_literal: true

def check_power(has_power)
  raise 'No power, unable to cook!' unless has_power
end

def add_water(water)
  raise 'Please specify the amount of water!' if water.nil?
  raise 'Water amount should be greater than zero!' if water.zero?
end

def check_item_and_timer(item_to_cook, timer)
  raise 'Please specify the item to cook!' if item_to_cook.nil?
  raise 'Time to cook is insufficient' if item_to_cook == 'rice' && timer < 3
end

def start_cooking(item_to_cook, water, timer)
  raise 'Timer must be greater than 0!' if timer <= 0

  puts 'Power is available!'
  puts "Adding #{water} cups of water..."
  puts "Starting #{item_to_cook} cooking..."
  sleep(timer)
  puts "#{item_to_cook} is cooked! Time to eat."
end

def cook(timer, item_to_cook = nil, water = nil, has_power: false)
  check_power(has_power)
  add_water(water)
  check_item_and_timer(item_to_cook, timer)
  start_cooking(item_to_cook, water, timer)
rescue StandardError => e
  puts "An error occurred: #{e.message}"
end

cook(10, 'rice', nil, has_power: true)
cook(10, nil, 2, has_power: true)
cook(2, 'tomato', 4, has_power: true)
cook(2, 'rice', 4, has_power: true)
cook(10, 'rice', 4, has_power: true)
