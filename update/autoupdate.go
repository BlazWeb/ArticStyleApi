package update

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const version = "0.0.5"

func ConfirmAndSelfUpdate() {
	latest, found, err := selfupdate.DetectLatest("Artic-Dev/ArticStyleApi")
	if err != nil {
		log.Println("Error occurred while detecting version:", err)
		return
	}

	v := semver.MustParse(version)
	if !found || latest.Version.LTE(v) {
		log.Println("La versión actual es la más reciente " + version)
		log.Println(latest)
		return
	}

	fmt.Print("¿Quieres actualizar a la versión ", latest.Version, "? (y/n): ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil || (input != "y\n" && input != "n\n") {
		log.Println("Solamente introduzca y o n")
		return
	}
	if input == "n\n" {
		return
	}

	exe, err := os.Executable()
	if err != nil {
		log.Println("No se ha podido encontrar la ruta del ejecutable")
		return
	}
	if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
		log.Println("Error al actualizar binario:", err)
		return
	}
	log.Println("Actualizado correctamente a la versión ", latest.Version)
}
