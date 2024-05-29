package exercicio

import (
	"encoding/json"
	"fmt"
)

type Actor struct {
	First   string   `json:"first"`
	Last    string   `json:"last"`
	Age     int      `json:"age"`
	Sayings []string `json:"sayings"`
}

func Exercicio() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	jsonBlob := []byte(s)

	var actores []Actor

	err := json.Unmarshal(jsonBlob, &actores)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(actores)
	fmt.Println(actores[0])
	fmt.Println(actores[0].Last)

}
