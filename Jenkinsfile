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
                sh "docker container kill ${dockerName}"
                sh "docker container rm ${dockerName}"
                sh "docker-compose --env-file C:/Users/hrd/sample_secret/.env up -d"
            }
        }
    }
}