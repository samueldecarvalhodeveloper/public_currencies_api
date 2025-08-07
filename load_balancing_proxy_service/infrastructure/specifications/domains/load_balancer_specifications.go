package specifications

func IsCurrentServerTheLast(
	currentServerIndex int,
	listOfServersToBeBalanced []string) bool {
	return currentServerIndex == len(listOfServersToBeBalanced)-1
}
