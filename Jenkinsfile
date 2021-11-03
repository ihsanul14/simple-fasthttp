def dockerName = "simple-fasthttp"
def container = "simple-fasthttp"
def port = "30001"

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
                sh "docker-compose --env-file /.env up"
            }
        }
    }
}