pipeline {
    agent any

        environment {
        DOCKER_IMAGE_FRONTEND = "fayez74/react-frontend:latest"
        DOCKER_IMAGE_BACKEND = "fayez74/go-backend:latest"
        DOCKER = "/opt/homebrew/bin/docker"
    }

    stages {
            stage('Checkout') {
                steps {
                    // Pull latest code from GitHub
                    checkout([$class: 'GitSCM', 
                            branches: [[name: '*/main']], 
                            userRemoteConfigs: [[url: 'https://github.com/Fayez73/go-app-jenkins', credentialsId: 'github-creds']]])
                }
            }
            stage('Check files') {
                steps {
                    sh 'pwd'
                    sh 'ls -la ~'
                }
            }
            stage('Docker Build Front end') {
                steps {
                    script {
                        dir("${env.WORKSPACE}/frontend") {
                            withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                                sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'
                                sh "${DOCKER} build -t ${DOCKER_IMAGE_FRONTEND} ."
                            }
                        }
                    }
                }
            }
            stage('Push to Docker Hub') {
                steps {
                    sh 'cd go-app-jenkins/frontend'
                    withCredentials([usernamePassword(credentialsId: 'docker-credentials', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                        sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'
                        sh "docker push ${DOCKER_IMAGE_FRONTEND}"
                    }
                }
            }
        }

        post {
            success {
                echo "Pipeline completed successfully!"
            }
            failure {
                echo "Pipeline failed. Check logs."
            }
        }
}