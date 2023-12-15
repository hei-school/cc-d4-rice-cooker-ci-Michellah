require 'rspec'

require_relative 'main'

RSpec.describe '#check_power' do 
    context 'when there is power' do
        it 'does not raise an error' do
            expect { check_power(true) }.not_to raise_error
        end
    end

    context 'when there is no power' do
        it 'raise an error' do
            expect { check_power(false) }.to raise_error('No power, unable to cook!')
        end
    end
end     

RSpec.describe '#add_water' do 
    context 'when water is null' do
        it 'raise an error' do
            expect { add_water(nil) }.to raise_error(RuntimeError, 'Please specify the amount of water!')
        end
    end
    context 'when water specify is zero' do
        it 'raise an error' do
            expect { add_water(0) }.to raise_error(RuntimeError, 'Water amount should be greater than zero!')
        end
    end
    context 'when water specify > 0 and not null' do
        it 'does not raise an error' do
            expect { add_water(2) }.not_to raise_error
        end
    end
end  

RSpec.describe '#check_item_and_timer' do 
    context 'when item to cook is null' do
        it 'raise an error' do
            expect { check_item_and_timer(nil, 2) }.to raise_error(RuntimeError, 'Please specify the item to cook!')
        end
    end
    context 'when time to cook rice is insufficient' do
        it 'raise an error' do
            expect { check_item_and_timer('rice', 2) }.to raise_error(RuntimeError, 'Time to cook is insufficient')
        end
    end
    context 'when item to cook and timer is ok' do
        it 'does notraise an error' do
            expect { check_item_and_timer('tomato', 4) }.not_to raise_error
        end
    end
end  

RSpec.describe '#start_cooking' do 
    context 'when item to cook is < 0' do
        it 'raise an error' do
            expect { start_cooking('tomato', 2, 0) }.to raise_error(RuntimeError, 'Timer must be greater than 0!')
        end
    end
    context 'outputs messages when starting cooking' do
        it 'does not raise an error' do
            item_to_cook = 'tomato'
            water = 2
            timer = 3
            expect { start_cooking(item_to_cook, water, timer) }.to output(
                "Power is available!\nAdding #{water} cups of water...\nStarting #{item_to_cook} cooking...\n#{item_to_cook} is cooked! Time to eat.\n"
            ).to_stdout
        end    
    end        
end
