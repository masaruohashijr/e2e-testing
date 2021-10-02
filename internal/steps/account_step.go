package steps

import (
	"fmt"
	"strconv"
	"zarbat_test/internal/logging"
)

func IShouldGetToSeeMyAccountBalanceMoreThanOrEqualTo(balance float64) error {
	var err error

	AccountInfo, err = AccountPrimaryPort.ViewAccount()
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	accountBalance, _ := strconv.ParseFloat(AccountInfo.AccountBalance, 64)
	if accountBalance < balance {
		return fmt.Errorf("Expected account balance should be %f but got %f", balance, accountBalance)
	}

	return nil
}

func IShouldGetToSeeAsTheFriendlyNameForMyAccount(friendlyName string) error {
	var err error

	AccountInfo, err = AccountPrimaryPort.ViewAccount()
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}
	if AccountInfo.FriendlyName != friendlyName {
		return fmt.Errorf("Error: Expected friendly name %s but got %s", friendlyName, AccountInfo.FriendlyName)
	}

	return nil
}

func IUpdateTheFriendlyNameForMyAccountTo(friendlyName string) error {
	err := AccountPrimaryPort.UpdateAccount(friendlyName)
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}
	return nil
}

func IViewMyAccountInformation() error {
	var err error
	AccountInfo, err = AccountPrimaryPort.ViewAccount()
	logging.Debug.Printf("My friendly name is %s\n", AccountInfo.FriendlyName)
	if err != nil {
		return fmt.Errorf("Error %s", err.Error())
	}
	return nil
}
