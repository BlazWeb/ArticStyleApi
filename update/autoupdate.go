package update

import (
	"log"
	"time"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const version = "0.1.6"

func ConfirmAndSelfUpdate() {
	v := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(v, "Artic-Dev/ArticStyleApi")
	if err != nil {
		log.Println("ArticDev -> Binary update failed:", err)
		return
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		log.Println("ArticDev -> Estas en la última versión la ", version)
	} else {
		log.Println("ArticDev -> Actualización completada con éxito", latest.Version)
		log.Println("ArticDev -> Notas del parche:\n", latest.ReleaseNotes)
		time.Sleep(10 * time.Second)
		log.Fatal("ArticDev -> Cerrando aplicacion...")
	}

}
