# Get the input 
annual_salary= float(input("Enter the starting annual salary: "))
monthly_salary= annual_salary/12


# Fixed Parameters
portion_int_list = [i for i in range(1,10001)]
semi_annual_raise=.07 
r=0.04
portion_down_payment = .25
total_cost=1000000
down_payment = total_cost*portion_down_payment

# CheckSavingPortion
def CheckSavingPortion(portion_saved,monthly_salary=monthly_salary):
    month_count=0
    current_savings=0
    while current_savings < down_payment:
        current_savings += current_savings*r/12
        current_savings += monthly_salary*portion_saved
        month_count += 1
        if month_count % 6 == 0:
            monthly_salary += monthly_salary*semi_annual_raise
    return month_count


#Bisection Serch
def BisectionSearch(Bisection_list, Bisection_count):
    month_count= CheckSavingPortion(Bisection_list[int(len(Bisection_list)/2-1)]/10000) 
    if month_count== 36 :
        return Bisection_list[int(len(Bisection_list)/2-1)], Bisection_count+1
    elif  month_count < 36 :
        return BisectionSearch(Bisection_list[:int(len(Bisection_list)/2-1)], Bisection_count+1)
    else:
        return BisectionSearch(Bisection_list[int(len(Bisection_list)/2-1):], Bisection_count+1)

if monthly_salary < down_payment/36:
    print("It is not possible to pay the down payment in three years.")
else:
    saving_rate, steps = BisectionSearch(portion_int_list,0)
    print("Best savings rate:", saving_rate/10000)
    print("Steps in bisection search:", steps)