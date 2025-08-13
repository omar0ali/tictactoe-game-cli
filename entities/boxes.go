package entities

var boxes []*BoxHolder

func SetBoxes(listOfBoxes []*BoxHolder) {
	boxes = listOfBoxes
}

func GetBoxes() []*BoxHolder {
	if boxes != nil {
		return boxes
	}
	return nil
}

func checkBoxes() bool {
	return boxes != nil
}

func DisabledBoxes() bool {
	if !checkBoxes() {
		return false
	}
	for _, box := range boxes {
		box.disable = true
	}
	return true
}

func EnableBoxes() bool {
	if !checkBoxes() {
		return false
	}
	for _, box := range boxes {
		box.disable = false
	}
	return true
}
