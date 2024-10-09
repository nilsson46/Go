package main

import (
	"fmt"

	"github.com/cheynewallace/tabby"
)

type Player struct {
	Name         string
	Age          int
	JerseyNumber int
	Team         string
}

/*func createFoppa() *Player {
	foppa := Player{
		Name:         "Peter Forsberg",
		Age:          47,
		JerseyNumber: 21,
		Team:         "Colorado Avalanche",
	}
	return &foppa

} */

/*func changeOnePlayer(players [3]Player) { //Player blir en kopia.
	players[0].Name = "Foo"
}
*/

func changeOnePlayer(players []Player) { //Player blir en kopia.
	players[0].Name = "Foo"
	Kalle := Player{
		Name:         "Kalle",
		Age:          47,
		JerseyNumber: 31,
		Team:         "New york Rangers",
	}
	players = append(players, Kalle)
	for _, player := range players {
		fmt.Println(player.Name)
	}

}
func main() {

	//simon := employee.New("Simon", 47) //Gentlemanna constructor
	//salary := simon.CalculateSalary()  // go style
	//fmt.Println(salary)
	//simon.CalculateSalary()

	foppa := Player{
		Name:         "Peter Forsberg",
		Age:          47,
		JerseyNumber: 21,
		Team:         "Colorado Avalanche",
	}
	Lundqvist := Player{
		Name:         "Henrik Lundqvist",
		Age:          47,
		JerseyNumber: 32,
		Team:         "New york Rangers",
	}
	Lundqvist2 := Player{
		Name:         "Henrik Lundqvist",
		Age:          47,
		JerseyNumber: 33,
		Team:         "New york Rangers",
	}
	/*var players = [3]Player{foppa, Lundqvist, Lundqvist2}
	changeOnePlayer(players) // . Detta betyder att en kopia av spelarna skapas och ändras.

	players[0].Name = "Foo"
	*/

	var players = []Player{foppa, Lundqvist, Lundqvist2}
	changeOnePlayer(players) // . Detta betyder att en kopia av spelarna skapas och ändras.

	newPlayer := Player{
		Name:         "Henrik Lundqvist",
		Age:          47,
		JerseyNumber: 334,
		Team:         "New york Rangers",
	}

	//motsvarande players = realloc(players, ...)

	players = append(players, newPlayer)

	fmt.Println("Original slice - After appending newPlayer:")
	for _, player := range players {
		fmt.Println(player.Name)
	}

	fmt.Println("***********")
	tab := tabby.New()
	tab.AddHeader("Name", "Age", "JerseyNumber", "Team")
	for _, player := range players {
		tab.AddLine(player.Name, player.Age, player.JerseyNumber, player.Team)
		//fmt.Println(player.Name)
	}
	tab.Print()
}
