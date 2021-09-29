package steps

import (
	"fmt"
	"strconv"
	"zarbat_test/internal/logging"
)

func IShouldGetToSeeMyAccountBalanceMoreThanOrEqual(balance int) error {
	var err error

	AccountInfo, err = AccountPrimaryPort.ViewAccount()
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}
	accountBalance, _ := strconv.ParseFloat(AccountInfo.AccountBalance, 64)
	fBalance := float64(balance)
	if accountBalance < fBalance {
		return fmt.Errorf("Error: Expected account balance should be %d but got %s", balance, AccountInfo.AccountBalance)
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
