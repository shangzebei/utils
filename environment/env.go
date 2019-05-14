package environment

import log "github.com/sirupsen/logrus"

type EnvProfile string

const (
	Dev  EnvProfile = "dev"
	Prod            = "prod"
)

var profile EnvProfile

func SetProfiles(envProfile string) {
	if envProfile == "dev" {
		profile = Dev
	}
	if envProfile == "prod" {
		profile = Prod
	}
	log.Printf("application run in %s\n", profile)

}

func GetProfile() EnvProfile {
	return profile
}

func ISDev(f func()) {
	if profile == "" {
		panic("profile not init")
	}
	if profile == Dev {
		f()
	}

}

func ISProd(f func()) {
	if profile == "" {
		panic("profile not init")
	}
	if profile == Prod {
		f()
	}
}
