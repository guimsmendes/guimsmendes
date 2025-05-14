  <div align="center">
    <img src="./github.png" align="center" style="width: 100%" />
  </div>  
  <div align="center">
    <img src="https://komarev.com/ghpvc/?username=guimsmendes&&style=flat-square" align="center" />
    <img src="https://wakatime.com/badge/user/496581e6-44e0-4e81-b292-7c3f7c33a58f/project/125dcb50-b3da-4152-9e34-bd7fb3db88bf.svg" align="center" />
  </div>  

### <div align="center"></div>  

```go
  var gMendes = Profile{
	Name: "Guilherme Mendes", 
	Location: "Amsterdam, NL", 
	Birthplace: "Santos, BR", 
	LivedCountries: []string{"BR", "JP", "NL"}, 
	Hobbies: []string{"Travel", "Cats", "Music", "Photography", "Bike", "Videogame"},
	EagerToLearn: []string{"React", "Python", "LLM"}
	Languages: []Language{
		{ Name: "pt-BR", Level:"native"},
		{ Name: "en-US", Level:"fluent"},
		{ Name: "es-ES", Level:"intermediate"},
		{ Name: "jp-JP", Level:"basic"},
		{ Name: "nl-NL", Level:"basic"},
	}, 
	Experiences: []Job{
		{ Role: "Sofware & DevOps Engineer", Company: "Gain.pro", StartYear: 2023, CurrentEmployer: true, MainStack: []string{"go", "gcp", "pulumi", "postgres"} },
		{ Role: "Senior Software Engineer", Company: "Sniptech", StartYear: 2022, EndYear: 2023,  MainStack: []string{"go", "gcp", "terraform", "kubernetes", "postgres", "elasticsearch", "kafka"} },
		{ Role: "Software Engineer", Company: "CVC", StartYear: 2021, EndYear: 2022, MainStack: []string{"java", "spring", "go", "aws", "mongodb", "elasticsearch", "kafka"} },
		{ Role: "Software Engineer", Company: "Itaú", StartYear: 2017, EndYear: 2021, MainStack: []string{"java", "spring", "sybase", "kafka"} },
	},
	Education: []School{
		{ Name: "Universidade Federal do ABC", Location: "São Paulo, BR", Degree: "Bachelor in Science and Technology", StartYear: 2013, EndYear: 2017 },
		{ Name: "Shibaura Institute of Technology", Location: "Tokyo, JP", Degree: "Bachelor in Science and Technology (Exchange program)", StartYear: 2015, EndYear: 2016 },
	},
  }
  
  func (p Profile) Start() {
	  fmt.Println("wake up")
	  for _, job := range p.Experiences {
		  if job.CurrentEmployer { job.Work() break }
	  }
	  
	  for _, language := range p.Languages {
		  if language.Level == "intermediate" || language.Level == "basic" { language.Practice() }
	  }
	  
	  p.Play(Hobbies[rand.Intn(len(p.Hobbies))])
	  
	  p.Study(EagerToLearn[rand.Intn(len(EagerToLearn))])
	  
	  fmt.Println("sleep")
  }
  
  func main() {
	  gMendes.Start()
  }
```

<br/>  


## Connect  
  <div align="center">
    <a href="https://linkedin.com/in/guilherme-mendes-b5381555" target="_blank">
    <img src=https://img.shields.io/badge/linkedin-%231E77B5.svg?&style=for-the-badge&logo=linkedin&logoColor=white alt=linkedin style="margin-bottom: 5px;" />
    </a>
    <a href = "mailto:guimsmendes@gmail.com">
    <img src="https://img.shields.io/badge/-Gmail-%23333?style=for-the-badge&logo=gmail&logoColor=white" target="_blank">
    </a>
    <a href="https://instagram.com/guimsmendes" target="_blank">
    <img src=https://img.shields.io/badge/instagram-%23000000.svg?&style=for-the-badge&logo=instagram&logoColor=white alt=instagram style="margin-bottom: 5px;" />
    </a>
    <a href="https://profile.playstation.com/Gui_Mendes95" target="_blank">
    <img src="https://img.shields.io/badge/Playstation-003791?style=for-the-badge&logo=playstation&logoColor=white" alt=playstation style="margin-bottom: 5px;" />
    </a>
    <a href="https://www.hackerrank.com/profile/guimsmendes" target="_blank">
    <img src="https://img.shields.io/badge/-Hackerrank-2EC866?style=for-the-badge&logo=HackerRank&logoColor=white" alt=hackerrank style="margin-bottom: 5px;" />
    </a>
    <a href="https://dev.to/guimsmendes" target="_blank">
    <img src=https://img.shields.io/badge/dev.to-%2308090A.svg?&style=for-the-badge&logo=dev.to&logoColor=white alt=devto style="margin-bottom: 5px;" />
    </a>
    <a href="https://medium.com/@guilhermemendes_57138" target="_blank">
    <img src=https://img.shields.io/badge/medium-%23292929.svg?&style=for-the-badge&logo=medium&logoColor=white alt=medium style="margin-bottom: 5px;" />
    </a>
    <a href = "https://soundcloud.com/yunka-dj" target="_blank">
    <img src="https://img.shields.io/badge/SoundCloud-FF3300?style=for-the-badge&logo=soundcloud&logoColor=white" target="_blank">
    </a>
    <a href = "https://www.duolingo.com/profile/guimsmende1" target="_blank">
    <img src="https://img.shields.io/badge/Duolingo-%234DC730.svg?style=for-the-badge&logo=Duolingo&logoColor=white" target="_blank">
    </a>

   
  </div>  
  

<br/>  



<h2>Stack </h2> 
<table align="center" style="width:100%;  text-align:center;table-layout:auto;">
  <tr>
    <td valign="center" style="width:50%; text-align:center; padding: 1000px;">
      <p align="center" style="font-weight: bold; font-size: 56px;">
         BackEnd
      </p>
      <div align="center">
        <table>
          <tr>
            <td style="padding: 20px;"><a href="https://go.dev" target="_blank"><img height="40" width="40" src="https://cdn.simpleicons.org/go" /></a></td>
            <td style="padding: 20px;"><a href="https://www.java.com/en/" target="_blank"><img height="40" width="40" src="https://cdn4.iconfinder.com/data/icons/logos-and-brands/512/181_Java_logo_logos-512.png" alt="Java" /></a></td>
            <td style="padding: 20px;"><a href="https://spring.io" target="_blank"><img height="40" width="40" src="https://profilinator.rishav.dev/skills-assets/springio-icon.svg" alt="Spring"/></a></td>
            <td style="padding: 20px;"><a href="https://quarkus.io" target="_blank"><img height="40" width="40" src="https://www.svgrepo.com/show/354245/quarkus-icon.svg" alt="Quarkus"/></a></td>
            <td style="padding: 20px;"><a href="https://kafka.apache.org" target="_blank"><img height="40" width="40" src="https://openwhisk.apache.org/images/icons/icon-kafka-white-trans.png" alt="Kafka" /> </a></td>
          </tr>
          <tr>
            <td style="padding: 20px;"><a href="https://www.postgresql.org" target="_blank"><img height="40" width="40" src="https://cdn.simpleicons.org/postgresql" /> </a></td>
            <td style="padding: 20px;"><a href="https://www.mongodb.com" target="_blank"><img height="40" width="40" src="https://cdn.simpleicons.org/mongodb" /></a></td>
            <td style="padding: 20px;"><a href="https://www.elastic.co" target="_blank"><img height="40" width="40" src="https://profilinator.rishav.dev/skills-assets/elasticsearch.png" alt="Elastic Search" /></a></td>
            <td style="padding: 20px;"><a href="https://redis.io" target="_blank"><img height="40" width="40" src="https://www.svgrepo.com/show/303460/redis-logo.svg" alt="Redis" /></a></td>
            <td style="padding: 20px;"><a href="https://www.rabbitmq.com" target="_blank"><img height="40" width="40" src="https://cdn.simpleicons.org/rabbitmq" /></a></td>
          </tr>
        </table>
      </div>
    </td>
    <td valign="top" style="width:50%; text-align:center; padding: 50px;"">
      <p align="center" style="font-weight: bold; font-size: 56px;">
         DevOps
      </p>
      <div align="center">
        <table>
          <tr>
            <td style="padding: 20px;"><a href="https://cloud.google.com/" target="_blank"><img height="40" width="40" src="https://profilinator.rishav.dev/skills-assets/google_cloud-icon.svg" alt="GCP" /></a></td>
            <td style="padding: 20px;"><a href="https://www.pulumi.com/" target="_blank"><img height="40" width="40" src="https://www.pulumi.com/logos/brand/avatar-on-black.svg" alt="Pulumi" /></a></td>
            <td style="padding: 20px;"><a href="https://www.docker.com/" target="_blank"><img height="40" width="40" src="https://www.svgrepo.com/show/373553/docker.svg" alt="Docker"/></a></td>
            <td style="padding: 20px;"><a href="https://kubernetes.io/" target="_blank"><img height="40" width="40" src="https://profilinator.rishav.dev/skills-assets/kubernetes-icon.svg" alt="Kubernetes" /></a></td>
            <td style="padding: 20px;"><a href="https://www.terraform.io/" target="_blank"><img height="40" width="40" src="https://profilinator.rishav.dev/skills-assets/terraformio-icon.svg" alt="Terraform" /></a></td>
          </tr>
          <tr>      
            <td style="padding: 20px;"><a href="https://github.com/" target="_blank"><img height="40" width="40" src="https://img.icons8.com/ios_filled/200/FFFFFF/github.png" alt="Git" /></a></td>
            <td style="padding: 20px;"><a href="https://aws.amazon.com/" target="_blank"><img height="40" width="40" src="https://static-00.iconduck.com/assets.00/aws-icon-2048x2048-ptyrjxdo.png" alt="AWS" /></a></td>
            <td style="padding: 20px;"><a href="https://www.jenkins.io/" target="_blank"><img height="40" width="40" src="https://profilinator.rishav.dev/skills-assets/jenkins-icon.svg" alt="Jenkins" /></a></td>
            <td style="padding: 20px;"><a href="https://grafana.com/" target="_blank"><img height="40" width="40" src="https://upload.wikimedia.org/wikipedia/commons/thumb/3/3b/Grafana_icon.svg/351px-Grafana_icon.svg.png?20230113182558" alt="Grafana" /></a></td>
            <td style="padding: 20px;"><a href="https://www.elastic.co/kibana/" target="_blank"><img height="40" width="40" src="https://profilinator.rishav.dev/skills-assets/kibana.png" alt="Kibana" /></a></td>
          </tr>
        </table>
      </div>
    </td>
  </tr>
</table>  

<br/>  


<h2>Music</h2>
<p align="center">
<br/>
<a href="https://spotify-github-profile.kittinanx.com/api/view?uid=12160458370&redirect=true" target="_blank">
  <img height="400em" src="https://spotify-github-profile.kittinanx.com/api/view?uid=12160458370&cover_image=true&theme=default&show_offline=false&background_color=121212&interchange=false&bar_color=53b14f&bar_color_cover=false" alt="Spotify GitHub Profile" align="center">
</a>
<a href="https://spotify-recently-played-readme.vercel.app/api?user=12160458370&unique=true" target="_blank">
  <img height="400em" src="https://spotify-recently-played-readme.vercel.app/api?user=12160458370&unique=true" alt="Recently Played" align="center">
</a>
<br>
</p>


  <summary><h2>Github Stats</h2></summary>


 <p align="center">
  <img height="250em" src="https://streak-stats.demolab.com?user=guimsmendes&theme=ayu-mirage&border_radius=20&exclude_days=Sun%2CSat&card_height=200"/>
  <img height="250em" src="https://github-readme-stats.vercel.app/api/top-langs/?username=guimsmendes&langs_count=7&theme=ayu-mirage&border_radius=20"/>
</p>
<br/>  


<!--START_SECTION:activity-->

<!--END_SECTION:activity-->


<!--START_SECTION:waka-->

```txt
Go           5 hrs 49 mins   █████████████████████▓░░░   86.59 %
SQL          29 mins         █▓░░░░░░░░░░░░░░░░░░░░░░░   07.25 %
Python       18 mins         █▒░░░░░░░░░░░░░░░░░░░░░░░   04.71 %
JavaScript   1 min           ░░░░░░░░░░░░░░░░░░░░░░░░░   00.39 %
YAML         1 min           ░░░░░░░░░░░░░░░░░░░░░░░░░   00.37 %
```

<!--END_SECTION:waka-->
[![Guimsmendes's github activity graph](https://github-readme-activity-graph.vercel.app/graph?username=guimsmendes&theme=react-dark)](https://github.com/ashutosh00710/github-readme-activity-graph)

