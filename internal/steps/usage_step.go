package steps

import (
	"fmt"
	"strconv"
)

func IShouldListAtLeastUsage(number int) error {
	usages, err1 := UsagePrimaryPort.ListUsage()
	if err1 != nil {
		return fmt.Errorf("Error found in list Usages.")
	}
	if len(usages) < number {
		return fmt.Errorf("Error. Minimum number of usages expected is %d and found %d.", number, len(usages))
	}
	return nil
}

func IShouldViewTheTotalCostUsageMoreThan(expectedCost float64) error {
	usages, err1 := UsagePrimaryPort.ListUsage()
	if err1 != nil {
		return fmt.Errorf("Error found in list Usages.")
	}
	if len(usages) == 0 {
		return fmt.Errorf("No usages available.")
	}
	total, _ := strconv.ParseFloat(usages[0].TotalCost, 64)
	if total < expectedCost {
		return fmt.Errorf("Total cost expected is %f and got %f.", expectedCost, total)
	}
	return nil
}
