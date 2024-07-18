package guimsmendes

import (
	"fmt"
	"math/rand"
)

var gMendes = Profile{
	Name:           "Guilherme Mendes",
	Location:       "Amsterdam, NL",
	Birthplace:     "Santos, BR",
	LivedCountries: []string{"BR", "JP", "NL"},
	Hobbies:        []string{"Travel", "Cats", "Music", "Photography", "Videogame", "Skateboard"},
	Languages: []Language{
		{Name: "pt-BR", Level: "native"},
		{Name: "en-US", Level: "fluent"},
		{Name: "es-ES", Level: "intermediate"},
		{Name: "jp-JP", Level: "basic"},
		{Name: "nl-NL", Level: "basic"},
	},
	Experiences: []Job{
		{Role: "Sofware & DevOps Engineer", Company: "Gain.pro", StartYear: 2023, CurrentEmployer: true, MainStack: []string{"go", "gcp", "pulumi", "postgres"}},
		{Role: "Senior Software Engineer", Company: "Sniptech", StartYear: 2022, EndYear: 2023, MainStack: []string{"go", "gcp", "terraform", "kubernetes", "postgres", "elasticsearch", "kafka"}},
		{Role: "Software Engineer", Company: "CVC", StartYear: 2021, EndYear: 2022, MainStack: []string{"java", "spring", "go", "aws", "mongodb", "elasticsearch", "kafka"}},
		{Role: "Software Engineer", Company: "Itaú", StartYear: 2017, EndYear: 2021, MainStack: []string{"java", "spring", "sybase", "kafka"}},
	},
	Education: []School{
		{Name: "Universidade Federal do ABC", Location: "São Paulo, BR", Degree: "Bachelor in Science and Technology", StartYear: 2013, EndYear: 2017},
		{Name: "Shibaura Institute of Technology", Location: "Tokyo, JP", Degree: "Bachelor in Science and Technology (Exchange program)", StartYear: 2015, EndYear: 2016},
	},
}

type Profile struct {
	Name           string
	Location       string
	Birthplace     string
	LivedCountries []string
	Hobbies        []string
	Languages      []Language
	Experiences    []Job
	Education      []School
}

type Language struct {
	Name  string
	Level string
}

type Job struct {
	Role            string
	Company         string
	StartYear       int
	EndYear         int
	CurrentEmployer bool
	MainStack       []string
}

type School struct {
	Name      string
	Location  string
	Degree    string
	StartYear int
	EndYear   int
}

func (p Profile) Start() {
	fmt.Println("wake up")
	for _, job := range p.Experiences {
		if job.CurrentEmployer {
			job.Work()
			break
		}
	}

	for _, language := range p.Languages {
		if language.Level == "intermediate" || language.Level == "basic" {
			language.Practice()
		}
	}

	p.Hobbies[rand.Intn(len(p.Hobbies))].Play()

	fmt.Println("sleep")
}

func main() {
	gMendes.Start()
}
