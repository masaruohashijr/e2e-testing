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

func IShouldViewTheTotalCostUsageMoreThan(expectedCost int) error {
	usages, err1 := UsagePrimaryPort.ListUsage()
	if err1 != nil {
		return fmt.Errorf("Error found in list Usages.")
	}
	if len(usages) == 0 {
		return fmt.Errorf("No usages available.")
	}
	total, _ := strconv.Atoi(usages[0].TotalCost)
	if total < expectedCost {
		return fmt.Errorf("Total cost expected is %d and got %d.", expectedCost, total)
	}
	return nil
}
