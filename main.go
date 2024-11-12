package main

//Uso das bibliotecas Time e do Package Screen do Inancgumus
import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	type reservado [5]string //cria um array com 5 string de elementos

	zero := reservado{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := reservado{
		"██",
		" █",
		" █",
		" █",
		"███",
	}

	two := reservado{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := reservado{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := reservado{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := reservado{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := reservado{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := reservado{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := reservado{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := reservado{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colon := reservado{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}

	digits := [...]reservado{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	screen.Clear() //lib para limpar a tela do prompt
	for {          //criando um loop para att o relogio a cada segundo

		screen.MoveTopLeft() //Move o cursor ao topo

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := [...]reservado{ //criando o array do relogio
			//"retorna 1 digito da hora atual e o segundo
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}
		// Exibe cada linha de todos os dígitos ao lado uns dos outros
		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		fmt.Println()

		time.Sleep(time.Second)
	}

}
