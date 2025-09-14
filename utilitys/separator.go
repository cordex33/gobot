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

func SeparatorSD(list []string) [][]string {
	//este lo usamos en caso de que piden sin desarrolador, ya que la lista es menor
	//la única diferencia es que el tamanó de la lista es menor
	//función que toma una lista de x cantidad. Esta la separa en más listas para poder hacer un búcle al modelo
	var listoflist [][]string
	for x := 0; x < len(list); x += 9 {
		final := x + 9
		if final > len(list) {
			final = len(list)
		}
		listoflist = append(listoflist, list[x:final])
	}
	return listoflist

	//ejemplo: list = [a b c d e f g h i j k l m n] el listoflist := [a ... j] [k ... etc]
}
