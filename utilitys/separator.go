package tools

func Separator(list []string) [][]string {

	//función que toma una lista de x cantidad. Esta la separa en más listas para poder hacer un búcle al modelo
	var listoflist [][]string
	for x := 0; x < len(list); x += 10 {
		final := x + 10
		if final > len(list) {
			final = len(list)
		}
		listoflist = append(listoflist, list[x:final])
	}
	return listoflist

	//ejemplo: list = [a b c d e f g h i j k l m n] el listoflist := [a ... j] [k ... etc]
}
