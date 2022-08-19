package chainer

func CheckStepHandle(currStep int, defaultStep StepHandle, stepHandles ...StepHandle) int {
	for _, sh := range stepHandles {
		if currStep == int(sh) {
			return currStep
		}
	}

	return int(defaultStep)
}
