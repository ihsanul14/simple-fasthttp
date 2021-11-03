def dockerName = "simple-fasthttp"
def container = "simple-fasthttp"
def gopath = "D:/Golang"

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
                sh "export $(grep -v '^#' ${gopath}/src/${dockerName}/.env | xargs)"
                sh "docker container kill ${dockerName}"
                sh "docker container rm ${dockerName}"
                sh "docker-compose --env-file ${gopath}/src/${dockerName}/.env up -d"
            }
        }
    }
}