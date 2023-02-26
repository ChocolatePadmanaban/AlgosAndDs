import random

# Get the input 
annual_salary= float(input("Enter the starting annual salary: "))
monthly_salary= annual_salary/12


# Fixed Parameters
semi_annual_raise=.07 
r=0.04
portion_down_payment = .25
total_cost=1000000
down_payment = total_cost*portion_down_payment

#set problem parameters
epsilon = 100
bisection_search_steps = 0
max_portion = 10000
min_portion = 1

# Check if payment is possible
if monthly_salary < down_payment/36:
    print("It is not possible to pay the down payment in three years.")
else:
    while True:
        portion_saved = random.randint(min_portion,max_portion)
        monthly_salary= annual_salary/12
        bisection_search_steps+=1
        month_count=0
        current_savings=0
        while current_savings < down_payment:
            current_savings += current_savings*r/12
            current_savings += monthly_salary*portion_saved/10000
            month_count += 1
            if month_count % 6 == 0:
                monthly_salary += monthly_salary*semi_annual_raise
        present_epsilon= current_savings - down_payment
        if month_count == 36 and abs(present_epsilon) < epsilon:
            break
        elif month_count > 36 :
            min_portion = portion_saved + 1 
        else :
            max_portion = portion_saved - 1
    print("Best savings rate:", portion_saved/10000)
    print("Steps in bisection search:", bisection_search_steps)