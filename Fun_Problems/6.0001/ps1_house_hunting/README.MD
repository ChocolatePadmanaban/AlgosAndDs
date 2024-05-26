---
layout: default
---

# House Hunting


## Part A : House Hunting Basic

You have graduated from MIT and now have a great job! You move to the San Francisco Bay Area and
decide that you want to start saving to buy a house. As housing prices are very high in the Bay Area,
you realize you are going to have to save for several years before you can afford to make the down
payment on a house. In Part A, we are going to determine how long it will take you to save enough
money to make the down payment given the following assumptions:

1. **​total_cost** - Call the cost of your dream home ​.
2. **​portion_down_payment** -  Call the portion of the cost needed for a down payment ​. For
simplicity, assume that portion_down_payment = 0.25 (25%).
3. **current_savings​** - Call the amount that you have saved thus far ​. You start with a current
savings of $0.
4. **​r** - Assume that you invest your current savings wisely, with an annual return of **​r** ​(in other words,
at the end of each month, you receive an additional ​**current_savings*r/12**​ funds to put into
your savings – the 12 is because ​r​ is an annual rate). Assume that your investments earn a
return of r = 0.04 (4%).
5. **annual_salary** - Assume your annual salary is ​**annual_salary**​.
6. **​portion_saved​** - Assume you are going to dedicate a certain amount of your salary each month to saving for
the down payment. Call that **​portion_saved​**. This variable should be in decimal form (i.e. 0.1
for 10%).
7. **​monthly salary** - At the end of each month, your savings will be increased by the return on your investment,
plus a percentage of your **​monthly salary** ​(annual salary / 12).



Write a program to calculate how many months it will take you to save up enough money for a down
payment. You will want your main variables to be floats, so you should cast user inputs to floats.
Your program should ask the user to enter the following variables:
1. The starting annual salary (annual_salary)
2. The portion of salary to be saved (portion_saved)
3. The cost of your dream home (total_cost)

Try different inputs and see how long it takes to save for a down payment. **​Please make your
program print results in the format shown in the test cases below.**

**Test Case 1**
```
Enter your annual salary:​ 120000
Enter the percent of your salary to save, as a decimal:​ .10
Enter the cost of your dream home:​ 1000000
Number of months:​ 183
```

**Test Case 2**
```
Enter your annual salary:​ 80000
Enter the percent of your salary to save, as a decimal:​ .15
Enter the cost of your dream home:​ 500000
Number of months:​ 105
```

## Part B : Saving, with a Raise 

**Background**

In Part A, we unrealistically assumed that your salary didn’t change. But you are an MIT graduate, and
clearly you are going to be worth more to your company over time! So we are going to build on your
solution to Part A by factoring in a raise every six months.

copy your solution to Part A (as we are going to reuse much of that machinery). Modify
your program to include the following
1. Have the user input a semi-annual salary raise **​semi_annual_raise​** (as a decimal percentage)
2. After the 6​th​ month, increase your salary by that percentage. Do the same after the 12th
month, the 18​th​ month, and so on.


Write a program to calculate how many months it will take you save up enough money for a down
payment. LIke before, assume that your investments earn a return of ​r​ = 0.04 (or 4%) and the
required down payment percentage is 0.25 (or 25%). Have the user enter the following variables:
1. The starting annual salary (annual_salary)
2. The portion of salary to be saved (portion_saved)
3. The cost of your dream home (total_cost)
4. The semi­annual salary raise (semi_annual_raise)

**Test Case 1**
```
Enter your starting annual salary:​ 120000
Enter the percent of your salary to save, as a decimal:​ .05
Enter the cost of your dream home: ​500000
Enter the semi­annual raise, as a decimal:​ .03
Number of months:​ 142
```

**Test Case 2**
```
Enter your starting annual salary:​ 80000
Enter the percent of your salary to save, as a decimal:​ .1
Enter the cost of your dream home: ​800000
Enter the semi­annual raise, as a decimal:​ .03
Number of months:​ 159
```

**Test Case 3**
```
Enter your starting annual salary:​ 75000
Enter the percent of your salary to save, as a decimal:​ .05
Enter the cost of your dream home:​ 1500000
Enter the semi­annual raise, as a decimal:​ .05
Number of months:​ 261
```

## Part C: Finding the right amount to save away

In Part B, you had a chance to explore how both the percentage of your salary that you save each month
and your annual raise affect how long it takes you to save for a down payment. This is nice, but
suppose you want to set a particular goal, e.g. to be able to afford the down payment in three years.
How much should you save each month to achieve this? In this problem, you are going to write a
program to answer that question. To simplify things, assume:
1. Your semi-annual raise is .07(7%)
2. Your investments have an annual return of 0.04 (4%)
3. The down payment is .25 (25%) of the cost of the house 
4. The Cost of the house that you are saving is $1 M

You are now going to try to find the best rate of savings to achieve a down payment on a $1M house in
36 months. Since hitting this exactly is a challenge, we simply want your savings to be within $100 of
the required down payment.

write a program to calculate the best savings rate, as a function of your starting salary.
You should use ​bisection search​ to help you do this efficiently. You should keep track of the number of
steps it takes your bisections search to finish. You should be able to reuse some of the code you wrote
for part B in this problem.

Because we are searching for a value that is in principle a float, we are going to limit ourselves to two
decimals of accuracy (i.e., we may want to save at 7.04% ­­ or 0.0704 in decimal – but we are not
going to worry about the difference between 7.041% and 7.039%). This means we can search for an
integer​ between 0 and 10000 (using integer division), and then convert it to a decimal percentage
(using float division) to use when we are calculating the ​current_savings​ after 36 months. By using
this range, there are only a finite number of numbers that we are searching over, as opposed to the
infinite number of decimals between 0 and 1. This range will help prevent infinite loops. The reason we
use 0 to 10000 is to account for two additional decimal places in the range 0% to 100%. Your code
should print out a decimal (e.g. 0.0704 for 7.04%).




**Test Case 1**
```
Enter the starting salary:​ 150000
Best savings rate:​ 0.4411
Steps in bisection search:​ 12
```

**Test Case 2**
```
Enter the starting salary:​ 300000
Best savings rate:​ 0.2206
Steps in bisection search:​ 9
```

**Test Case 3**
```
Enter the starting salary:​ 10000
It is not possible to pay the down payment in three years.
```