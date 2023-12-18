export function checkPower(hasPower: boolean): void {
  if (!hasPower) {
    throw new Error('No power, unable to cook!');
  }
}
  
export function addWater(water: string | null): void {
  if (water === null) {
    throw new Error('Please specify the amount of water!');
  } 
}
  
export function checkItemAndTimer(itemToCook: string | null, timer: number): void {
  if (itemToCook === null) {
    throw new Error('Please specify the item to cook!');
  }
  if (itemToCook === 'rice' && timer < 3) {
    throw new Error('Time to cook is insufficient');
  }
}
  
export function startCooking(itemToCook: string | null, water: string | null, timer: number): void {
  if (timer <= 0) {
    throw new Error('Timer must be greater than 0!');
  }
  
  console.log('Power is available!');
  console.log(`Adding ${water} cups of water...`);
  console.log(`Starting ${itemToCook} cooking...`);
  setTimeout(() => {
    console.log(`${itemToCook} is cooked! Time to eat.`);
  }, timer * 1000);
}
  
function cook(timer: number, itemToCook: string | null = null, water: string | null = null, hasPower: boolean = false): void {
  try {
    checkPower(hasPower);
    addWater(water);
    checkItemAndTimer(itemToCook, timer);
    startCooking(itemToCook, water, timer);
  } catch (e) {
    console.log(`An error occurred: ${(e as Error).message}`);
  }
}
  
cook(10, 'rice', null, true);
cook(10, null, '2c', true);
cook(2, 'tomato', '4c', true);
cook(2, 'rice', '4c', true);
cook(10, 'rice', '4c', true);
  