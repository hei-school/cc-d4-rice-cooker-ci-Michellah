using NUnit.Framework;
using System;

namespace CookingApp.Tests
{
    [TestFixture]
    public class CookingSimulatorTests
    {
        [Test]
        public void CheckPower_NoPower_ThrowsException()
        {
            Assert.Throws<Exception>(() => CookingSimulator.CheckPower(false));
        }

        [Test]
        public void AddWater_NullInput_ThrowsException()
        {
            Assert.Throws<Exception>(() => CookingSimulator.AddWater(null));
        }

        [Test]
        public void CheckItemAndTimer_NullItem_ThrowsException()
        {
            Assert.Throws<Exception>(() => CookingSimulator.CheckItemAndTimer(null, 5));
        }

        [Test]
        public void CheckItemAndTimer_InsufficientTimeForRice_ThrowsException()
        {
            Assert.Throws<Exception>(() => CookingSimulator.CheckItemAndTimer("rice", 2));
        }

        [Test]
        public void StartCooking_NonPositiveTimer_ThrowsException()
        {
            Assert.Throws<Exception>(() => CookingSimulator.StartCooking("rice", "2c", 0));
        }

        [Test]
        public void Cook_RiceWithNullWater_ThrowsException()
        {
            Assert.Throws<Exception>(() => CookingSimulator.Cook(10, "rice", null, true));
        }

        [Test]
        public void Cook_NullItemWithWater_ThrowsException()
        {
            Assert.Throws<Exception>(() => CookingSimulator.Cook(10, null, "2c", true));
        }

        [Test]
        public void Cook_TomatoWithShortTimer_ThrowsException()
        {
            Assert.Throws<Exception>(() => CookingSimulator.Cook(2, "tomato", "4c", true));
        }

        [Test]
        public void Cook_RiceWithShortTimer_ThrowsException()
        {
            Assert.Throws<Exception>(() => CookingSimulator.Cook(2, "rice", "4c", true));
        }

        [Test]
        public void Cook_RiceWithValidInputs_NoExceptionThrown()
        {
            Assert.DoesNotThrow(() => CookingSimulator.Cook(10, "rice", "4c", true));
        }
    }
}
