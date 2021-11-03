def dockerName = "simple-fasthttp"
def container = "simple-fasthttp"

pipeline {
    agent any
    stages {
        stage('Build to Production') {
            when {
                branch 'main'
            }
            steps {
                sh "docker build . -t ${dockerName}"
            }
        }
        stage('Deploy to Production') {
            steps {
                sh "cat %GOPATH%/src/${dockerName}/.env"
                sh "docker container kill ${dockerName}"
                sh "docker container rm ${dockerName}"
                sh "docker-compose --env-file D:/Golang/src/${dockerName}/.env up -d"
            }
        }
    }
}