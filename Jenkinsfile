pipeline {
    agent any
    stages {
        stage('Lint') {
            steps {
                sh 'golangci-lint run ./...'
            }
        }
        stage('Test') {
            steps {
                sh 'go test ./... -v'
            }
        }
        stage('Docker Build') {
            steps {
                sh 'docker build -t payments-api:${BUILD_NUMBER} .'
            }
        }
    }
}