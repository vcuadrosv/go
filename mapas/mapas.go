package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string, 5)
	fmt.Println(paises)

	paises["Mexico"] = "D.F"
	paises["Colombia"] = "Bogota"

	fmt.Println(paises)
	fmt.Println(paises["Colombia"])

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 47,
		"Boca Junior": 67,
	}
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s tiene un puntaje %d\n ", equipo, puntaje)
	}

	delete(campeonato, "Boca Junior")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["juventus"]
	fmt.Printf("El puntaje es %d y el campeonato existe =  %t\n", puntaje, existe)
}
