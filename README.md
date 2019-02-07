<img src="/docs/assets/artie.svg" width="80">

# Artie

Modular repository manager written in go.

__Right now, POC of Maven repo using Minio/S3 backend.__

The goal for this project is to write a completely open-source modular Repository Manager which should be fully scalable.

## Dependencies

* golang

## Installation

`go get -u github.com/artheus/artie`

## Usage

Right now (in the POC), everything is hard-coded. So to try it out you will have to run a minio-instance on your localhost.
In the minio-instance there needs to be a bucket named `mymusic` and the access key must be `minio` and secret key `minio123`

This minio setup is due to the examples in https://github.com/minio/minio-go

Whenever you've got minio up and running. You can add this to your pom.xml

```xml
<distributionManagement>
	    <repository>
		    <id>internal.repo</id>
		    <name>local repo</name>
		    <url>http://localhost:8000/maven-private</url>
	</repository>
</distributionManagement>
```

And run `go run .` to start up `artie`. Then test a deploy, running `mvn deploy` in your maven project. Hopefully it will work! ;)

## Todo

- [x] ~~POC Maven repository~~
- [] Support repotsitory modules
- [] Support backend modules
- [] Support configuration with artie.yml file
- [] Support proxied repositories
- [] Support repository grouping
- [] Write additional repository modules (eg, docker, yum, npm, nuget, git, etc..)
