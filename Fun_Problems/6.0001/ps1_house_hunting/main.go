package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getFloatInput() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	float_str, err := reader.ReadString('\n')
	if err != nil {
		return 0.0, err
	}
	float_str = strings.TrimSpace(float_str)

	float_value, err := strconv.ParseFloat(float_str, 64)
	if err != nil {
		return 0.0, err
	}
	return float_value, nil
}

func ps1a() {
	// Prompt the user for input
	fmt.Print("Enter your annual salary: ")

	// Get Annuual Salary
	annual_salary, err := getFloatInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the float value
	monthly_salary := annual_salary / 12

	// Prompt the user for input
	fmt.Print("Enter the percent of your salary to save, as a decimal: ")

	// get The Portion if salary to be saved
	portion_saved, err := getFloatInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Prompt the user for input
	fmt.Print("Enter the cost of your dream home: ")

	// get The cost of your Dream Home
	total_cost, err := getFloatInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	portion_down_payment := .25 * total_cost
	current_savings := 0.0
	r := 0.04
	months := 0

	// do the iteration to calculate number of months recuired for downpayment
	for current_savings < portion_down_payment {
		current_savings = current_savings + (current_savings * r / 12) + (portion_saved * monthly_salary)
		months += 1
	}
	fmt.Println("Number of months:", months)
}

func ps1b() {
	// Prompt the user for input
	fmt.Print("Enter your annual salary: ")

	// Get Annuual Salary
	annual_salary, err := getFloatInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the float value
	monthly_salary := annual_salary / 12

	// Prompt the user for input
	fmt.Print("Enter the percent of your salary to save, as a decimal: ")

	// get The Portion if salary to be saved
	portion_saved, err := getFloatInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Prompt the user for input
	fmt.Print("Enter the cost of your dream home: ")

	// get The cost of your Dream Home
	total_cost, err := getFloatInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Enter the semiannual raise, as a decimal: ")
	semi_annual_raise, err := getFloatInput()
	if err != nil {
		fmt.Println("Error:", err)
	}

	portion_down_payment := .25 * total_cost
	current_savings := 0.0
	r := 0.04
	months := 0

	// do the iteration to calculate number of months recuired for downpayment
	for current_savings < portion_down_payment {
		current_savings = current_savings + (current_savings * r / 12) + (portion_saved * monthly_salary)
		months += 1
		if months%6 == 0 {
			monthly_salary = monthly_salary * (1 + semi_annual_raise)
		}
	}
	fmt.Println("Number of months:", months)
}

func ps1c() {
	// Prompt the user for input
	fmt.Print("Enter your annual salary: ")

	// Get Annuual Salary
	annual_salary, err := getFloatInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// bisection searach endpoints
	min_savings_rate := 0
	max_savings_rate := 10000

	// constants
	total_cost := 1000000.0
	semi_annual_raise := 0.07
	portion_down_payment := .25
	r := 0.04
	down_payment := portion_down_payment * total_cost

	// accumulators
	steps := 0
	dp_savings_diff := down_payment

	calulate_dp_savings_diff := func(savings_rate int, salary float64) float64 {
		current_savings := 0.0
		months := 0

		for months < 36 {
			return_on_investment := current_savings * r / 12
			portion_monthly_salary := float64(savings_rate) * salary / 120000
			current_savings += return_on_investment + portion_monthly_salary
			months += 1
			if months%6 == 0 {
				salary *= 1 + semi_annual_raise
			}
		}
		return down_payment - current_savings
	}

	if calulate_dp_savings_diff(max_savings_rate, annual_salary) > 0 {
		fmt.Println("It is not possible to pay the down payment in three years.")
		return
	}

	// start finding a savings rate when the savings amount is within $100 from
	// the down payment, or the algorithm cannot find a more precise savings rate
	// since the savings rate is limited to 2 decimal places.
	var mp_savings_rate int
	var min_diff, mp_diff float64
	for dp_savings_diff > 100 && max_savings_rate-min_savings_rate > 1 {
		mp_savings_rate = int((min_savings_rate + max_savings_rate) / 2)
		min_diff = calulate_dp_savings_diff(min_savings_rate, annual_salary)
		mp_diff = calulate_dp_savings_diff(mp_savings_rate, annual_salary)

		if (min_diff > 0) == (mp_diff > 0) {
			min_savings_rate = mp_savings_rate
		} else {
			max_savings_rate = mp_savings_rate
		}
		steps += 1
		dp_savings_diff = math.Abs(mp_diff)
	}
	fmt.Printf("Best svaings rate: %.4f\n", float64(mp_savings_rate)/10000)
	fmt.Println("Steps in bisection search: ", steps)
}
func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "ps1a":
			ps1a()
		case "ps1b":
			ps1b()
		case "ps1c":
			ps1c()
		default:
			fmt.Println("Proper arguments are ps1a or ps1b or ps1c")
		}
	} else {
		fmt.Println("Proper arguments are ps1a or ps1b or ps1c")
	}
}
