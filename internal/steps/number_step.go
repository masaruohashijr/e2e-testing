package steps

import (
	"encoding/xml"
	"fmt"
	"strings"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
)

func ConfiguredToDialAndSendDigitsTo(dialerNumber, digits, dialedNumber string) error {
	services.CloseChannel = true
	dialed, _ := Configuration.SelectNumber(dialedNumber)
	n := &domains.Number{
		Value:      dialed,
		SendDigits: digits,
	}
	d := &domains.DialNumber{
		Number: *n,
	}
	ResponseDialNumber.DialNumber = *d
	p := &domains.Hangup{}
	ResponseDialNumber.Hangup = *p
	x, _ := xml.MarshalIndent(ResponseDialNumber, "", "")
	strXML := domains.Header + string(x)
	services.WriteActionXML("number", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(dialerNumber)
	Configuration.VoiceUrl = services.BaseUrl + "/Number"
	NumberPrimaryPort.UpdateNumber()
	println(string(x))
	return nil
}

func ShouldBeReset(number string) error {
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.VoiceUrl = ""
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func IListAllAvailableNumbers() error {
	anumbers, err := NumberSecondaryPort.ListAvailableNumbers()
	AvailableNumbers = anumbers
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to list all available numbers.")
	}
	for _, a := range AvailableNumbers {
		logging.Debug.Println(a)
		println(a)
	}
	return nil
}

func IShouldGetToBuyFromList(amount int) error {
	ok := false
	for i := 0; i < amount; i++ {
		logging.Debug.Println("Buying number is: ", AvailableNumbers[i])
		NumberSecondaryPort.AddNumber(AvailableNumbers[i])
		purchased, _ := NumberSecondaryPort.ListNumbers()
		for _, n := range *purchased {
			if AvailableNumbers[i] == n.PhoneNumber {
				logging.Debug.Println("Purchased number is: ", AvailableNumbers[i])
				ok = true
				break
			}
		}
		if !ok {
			return fmt.Errorf("Error %s", "Not able to list available numbers.")
		}
	}
	return nil
}

func IListMyNumbers() error {
	myNumbers, err := NumberSecondaryPort.ListNumbers()
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to list available numbers.")
	}
	for _, in := range *myNumbers {
		logging.Debug.Println(in.PhoneNumber)
		println(in.PhoneNumber)
	}
	return nil
}

func IReleaseAllMyNumbersExcept(exceptionList string) error {
	myNumbers, err := NumberPrimaryPort.ListNumbers()
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to list my numbers.")
	}
	exList := strings.Split(exceptionList, ",")
	for _, a := range *myNumbers {
		exceptionNumberFound := false
		for _, e := range exList {
			pn, _ := Configuration.SelectNumber(e)
			if pn == a.PhoneNumber {
				exceptionNumberFound = true
				break
			}
		}
		if !exceptionNumberFound {
			println("Releasing " + a.PhoneNumber)
			NumberPrimaryPort.DeleteNumber(a.Sid)
		}
	}
	return nil
}

func IShouldGetNumbersFromMyList(amount int) error {
	ok := false
	myNumbers, err := NumberSecondaryPort.ListAvailableNumbers()
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to list my numbers.")
	}
	for i := 0; i < amount; i++ {
		logging.Debug.Println("Buying number is: ", AvailableNumbers[i])
		NumberSecondaryPort.AddNumber(AvailableNumbers[i])
		purchased, _ := NumberSecondaryPort.ListNumbers()
		for _, n := range *purchased {
			if myNumbers[i] == n.PhoneNumber {
				logging.Debug.Println("Purchased number is: ", IncomingNumbers[i])
				ok = true
				break
			}
		}
		if !ok {
			return fmt.Errorf("Error %s", "Not able to list my numbers.")
		}
	}
	return nil
}

func ConfiguredWithFriendlyNameAs(number, friendlyName string) error {
	services.CloseChannel = true
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.FriendlyName = friendlyName
	Configuration.VoiceUrl = ""
	NumberPrimaryPort.UpdateNumber()
	return nil
}
func IShouldGetFriendlyNameOn(friendlyName, number string) error {
	selectedNumber, sid := Configuration.SelectNumber(number)
	ipn, err := NumberPrimaryPort.ViewNumber(sid)
	IncomingPhoneNumber = ipn
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to view number info.")
	}
	println(IncomingPhoneNumber.FriendlyName)
	if IncomingPhoneNumber != nil {
		if IncomingPhoneNumber.PhoneNumber == selectedNumber {
			if IncomingPhoneNumber.FriendlyName == friendlyName {
				return nil
			}
		}
	}
	return fmt.Errorf("Error %s", "Not able to get friendly name on number.")

}
func IViewInfo(number string) error {
	_, sid := Configuration.SelectNumber(number)
	ipn, err := NumberPrimaryPort.ViewNumber(sid)
	IncomingPhoneNumber = ipn
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to view number info.")
	}
	println(IncomingPhoneNumber.FriendlyName)
	return nil

}

func IShouldListMyNumbersAs(list string) error {
	myNumbers, err := NumberSecondaryPort.ListNumbers()
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to list my numbers.")
	}
	arr := strings.Split(list, ",")
	for _, n := range *myNumbers {
		found := false
		for j := 0; j < len(arr); j++ {
			if n.PhoneNumber == arr[j] {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("Error %s", "List is different than expected.")
		}
	}
	return nil
}
