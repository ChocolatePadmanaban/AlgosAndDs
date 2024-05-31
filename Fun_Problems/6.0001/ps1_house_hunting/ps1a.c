#include <stdio.h>

int main() {
    float annual_salary,portion_saved , total_cost;
    printf("Enter your annual salary: ");
    if (scanf("%f",&annual_salary) != 1){
        printf("Invalid input. Please enter a valid floating-point number.\n");
        return 1;
    }
    printf("Enter the percent of your salary to save, as a decimal: ");
    if (scanf("%f", &portion_saved) != 1) {
        printf("Invalid input. Please enter a valid floating-point number.\n");
        return 1;   
    }
    printf("Enter the cost of your dream house: ");
    if (scanf("%f", &total_cost) != 1) {
        printf("Invalid input. Please enter a valid floating-point number.\n");
        return 1;   
    }
    float current_savings , portion_down_payment , monthly_salary;
    current_savings = 0 ;
    portion_down_payment =.25 * total_cost ;
    monthly_salary = annual_salary/12;
    int months = 0;
    float r = 0.04; //intrest rate
    while (current_savings <= portion_down_payment) {
        months += 1;
        current_savings = current_savings + (current_savings * r/12) + (portion_saved * monthly_salary);
    }
    printf("Number of months: %d\n", months);
    return 0;
}