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
                sh "docker container kill ${container}"
                sh "docker container rm ${container}"
                sh "docker run -d --name ${container} -p ${port}:${port} ${dockerName}"
            }
        }
    }
}