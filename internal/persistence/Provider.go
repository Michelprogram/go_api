package persistence

import (
	"fmt"
	"internal/persistence/daobolt"
	"internal/persistence/daomemory"
	"internal/persistence/interfaces"
	"os"
	"strings"
)

func ValidityOfArgs() bool {

	var args []string = os.Args

	var daos []string = []string{"memory", "bolt", "mongo"}

	var size int = len(args)

	if size < 2 || size > 2 {
		fmt.Println("Le nombre d'arguments est incohérent(s).")
		return false
	}

	args = strings.Split(args[1], "=")

	if args[0] != "--dao" {
		fmt.Println("L'argument doit être --dao")
		return false
	}

	if len(args) < 2 {
		fmt.Println("Le nombre de paramètres est incohérent(s).")
		return false
	}

	if !includes(args[1], daos) {
		fmt.Println("Dao supportés sont bolt, memory et mongo")
		return false
	}

	return true
}

func GetDaoLanguage() interfaces.LanguageDao {

	var dao string = strings.Split(os.Args[1], "=")[1]

	switch dao {
	case "mongo":

		return nil

	case "bolt":
		return daobolt.NewLanguageDaoBolt()

	case "memory":
		return daomemory.NewLanguageDaoMemory()

	default:
		return daomemory.NewLanguageDaoMemory()
	}

}

func GetDaoStudent() interfaces.StudentDao {

	var dao string = strings.Split(os.Args[1], "=")[1]

	switch dao {
	case "mongo":

		return nil

	case "bolt":
		return daobolt.NewStudentDaoBolt()

	case "memory":
		return daomemory.NewStudentDaoMemory()

	default:
		return daomemory.NewStudentDaoMemory()
	}

}

func includes(lookingFor string, list []string) bool {

	for _, word := range list {
		if word == lookingFor {
			return true
		}
	}

	return false
}
