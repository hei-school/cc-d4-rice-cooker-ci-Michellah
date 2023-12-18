import { checkPower, addWater, checkItemAndTimer} from './RiceCooker'

describe('checkPower function', () => {
  test('should throw error when hasPower is false', () => {
    expect(() => {
      checkPower(false);
    }).toThrow('No power, unable to cook!');
  });

  test('should not throw error when hasPower is true', () => {
    expect(() => {
      checkPower(true);
    }).not.toThrow();
  });
});

describe('addWater function', () => {
  test('should throw error when water is null', () => {
    expect(() => {
      addWater(null);
    }).toThrow('Please specify the amount of water!');
  });

  test('should not throw error when water is specified', () => {
    expect(() => {
      addWater('2c');
    }).not.toThrow();
  });
});

describe('checkItemAndTimer function', () => {
  test('should throw error when itemToCook is null', () => {
    expect(() => {
      checkItemAndTimer(null, 10);
    }).toThrow('Please specify the item to cook!');
  });

  test('should throw error when cooking rice with insufficient time', () => {
    expect(() => {
      checkItemAndTimer('rice', 2);
    }).toThrow('Time to cook is insufficient');
  });

  test('should not throw error when itemToCook and timer are specified correctly', () => {
    expect(() => {
      checkItemAndTimer('rice', 10);
    }).not.toThrow();
  });
});

