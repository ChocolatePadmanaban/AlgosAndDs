#include <stdio.h>
#include <math.h>

// constants 
    float total_costs = 1000000.0;
    float semi_annual_raise = 0.07;
    float portion_down_payment = .25;
    float r = 0.04;
    float down_payment = 250000;

float calculate_dp_savings_diff(int savings_rate,float salary){
    float current_savings = 0;
    int months = 0;

    while(months < 36){
        float return_on_investment = current_savings * r /12;
        float portion_monthly_salary = (float)savings_rate*salary/120000;
        current_savings += return_on_investment + portion_monthly_salary;
        months++;
        if(months%6==0){
            salary *= 1+ semi_annual_raise;
        }
    }
    return down_payment - current_savings;
}


int main() {
    printf("Enter your annual salary: ");
    float annual_salary;
    if (scanf("%f",&annual_salary) != 1){
        printf("Invalid input. Please enter a valid floating-point number.\n");
        return 1;
    }

    // bisection search endpoints 
    int min_savings_rate = 0;
    int max_savings_rate = 10000;

    // constants 
    float total_costs = 1000000.0;
    float semi_annual_raise = 0.07;
    float portion_down_payment = .25;
    float r = 0.04;
    float down_payment = portion_down_payment*total_costs;

    //accumulators 
    int steps = 0;
    float dp_savings_diff = down_payment;

    if (calculate_dp_savings_diff(max_savings_rate,annual_salary)>0) {
        printf("It is not possible to pay the down payment in three years.\n");
        return 0;
    }

    // start finding a savings rate when the savings amount is within $100 from
	// the down payment, or the algorithm cannot find a more precise savings rate
	// since the savings rate is limited to 2 decimal places.
    int mp_savings_rate;
    float min_diff , mp_diff;
    while (dp_savings_diff> 100 && max_savings_rate-min_savings_rate>1){
        mp_savings_rate= (int)((min_savings_rate+max_savings_rate)/2);
        mp_diff = calculate_dp_savings_diff(mp_savings_rate,annual_salary);

        if(mp_diff > 0) {
            min_savings_rate= mp_savings_rate;
        }else{
            max_savings_rate = mp_savings_rate;
        }
        steps++;
        dp_savings_diff = fabs(mp_diff);
    }
    printf("Best savings rate: %.4f\n", (float)(mp_savings_rate)/10000);
    printf("Steps in bisection search: %d\n",steps);
    return 0;
}