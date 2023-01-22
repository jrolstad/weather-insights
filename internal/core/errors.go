package core

import (
	"errors"
	"fmt"
)

func ConsolidateErrors(toMap []error) error {
	if toMap == nil || len(toMap) == 0 {
		return nil
	}

	hasErrorValue := false

	combinedMessage := ""
	for _, item := range toMap {
		if item != nil {
			combinedMessage = fmt.Sprintf("%v%v;", combinedMessage, item.Error())
			hasErrorValue = true
		}
	}

	if hasErrorValue {
		return errors.New(combinedMessage)
	}

	return nil

}
