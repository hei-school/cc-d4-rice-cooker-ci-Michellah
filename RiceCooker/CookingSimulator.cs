//------------------------------------------------------------------------------
// <summary>
// File Name    : CookingSimulator.cs
// Author       : VotreNom
// Created Date : [Date]
// Description  : Ce fichier contient la logique du simulateur de cuisson.
// Licence      : [Licence du projet]
// </summary>
//------------------------------------------------------------------------------

using System;

namespace CookingApp
{
    public class CookingSimulator
    {
    public static void CheckPower(bool hasPower)
    {
        if (!hasPower)
        {
            throw new Exception("No power, unable to cook!");
        }
    }

    public static void AddWater(string? water)
    {
        if (water == null)
        {
            throw new Exception("Please specify the amount of water!");
        }
    }

    public static void CheckItemAndTimer(string? itemToCook, int timer)
    {
        if (itemToCook == null)
        {
            throw new Exception("Please specify the item to cook!");
        }

        if (itemToCook == "rice" && timer < 3)
        {
            throw new Exception("Time to cook is insufficient");
        }
    }

    public static void StartCooking(string? itemToCook, string? water, int timer)
    {
        if (timer <= 0)
        {
            throw new Exception("Timer must be greater than 0!");
        }

        Console.WriteLine("Power is available!");
        Console.WriteLine($"Adding {water} cups of water...");
        Console.WriteLine($"Starting {itemToCook} cooking...");
        System.Threading.Thread.Sleep(timer * 1000);
        Console.WriteLine($"{itemToCook} is cooked! Time to eat.");
    }

    public static void Cook(int timer, string? itemToCook = null, string? water = null, bool hasPower = false)
    {
        try
        {
            CheckPower(hasPower);
            AddWater(water);
            CheckItemAndTimer(itemToCook, timer);
            StartCooking(itemToCook, water, timer);
        }
        catch (Exception e)
        {
            Console.WriteLine($"An error occurred: {e.Message}");
        }
    }

    public static void Main()
    {
        Cook(10, "rice", null, true);
        Cook(10, null, "2c", true);
        Cook(2, "tomato", "4c", true);
        Cook(2, "rice", "4c", true);
        Cook(10, "rice", "4c", true);
    }
    }
}