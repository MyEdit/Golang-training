package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var percent float32 = 0

    switch {
        case balance < 0:
        	percent = 3.213
        case balance >= 0 && balance < 1000:
        	percent = 0.5
        case balance >= 1000 && balance < 5000:
        	percent = 1.621
        default:
        	percent = 2.475
    }

    return percent
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance / 100 * float64(InterestRate(balance))
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0

    for ; balance < targetBalance; years++ {
    	balance = AnnualBalanceUpdate(balance)
	}

    return years
}
