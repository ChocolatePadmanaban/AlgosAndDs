# Get the inputs:
annual_salary= float(input("Enter your annual salary: "))
portion_saved= float(input("Enter the percent of your salary to save, as a decimal: "))
total_cost= float(input("Enter the cost of your dream home: "))


# Fixed Parameters
portion_down_payment = .25
r=.04
current_savings=0
monthly_salary=annual_salary/12
down_payment = total_cost*portion_down_payment

#Months to make downpayment 
month_count=0
while current_savings < down_payment:
    current_savings += current_savings*r/12
    current_savings += monthly_salary*portion_saved
    month_count += 1
print("Number of months:", month_count)