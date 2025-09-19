pipeline {
    agent any

        environment {
        DOCKER_IMAGE_FRONTEND = "fayez74/react-frontend:latest"
        DOCKER_IMAGE_BACKEND = "fayez74/go-backend:latest"
        DOCKER = "docker"
    }

    stages {
        stage('Cleanup') {
            steps {
                cleanWs()
            }
        }
            stage('Checkout') {
                steps {
                    checkout([$class: 'GitSCM',
                            branches: [[name: '*/main']],
                            userRemoteConfigs: [[url: 'https://github.com/Fayez73/go-app-jenkins.git']]])
                }
            }
            stage('Docker Build Front end') {
                steps {
                    script {
                        dir("${env.WORKSPACE}/frontend") {
                            withCredentials([usernamePassword(credentialsId: 'docker-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                                sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'
                                sh "${DOCKER} build -t ${DOCKER_IMAGE_FRONTEND} ."
                            }
                        }
                    }
                }

            }
            stage('Docker Build Backend end') {
                steps {
                    script {
                        dir("${env.WORKSPACE}/backend") {
                            withCredentials([usernamePassword(credentialsId: 'docker-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                                sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'
                                sh "${DOCKER} build -t ${DOCKER_IMAGE_BACKEND} ."
                            }
                        }
                    }
                }

            }
            stage('Push to Docker Hub') {
                steps {
                    script{
                        dir("${env.WORKSPACE}/frontend") {
                            withCredentials([usernamePassword(credentialsId: 'docker-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                                sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'
                                sh "docker push ${DOCKER_IMAGE_FRONTEND}"
                            }
                        }
                        dir("${env.WORKSPACE}/backend") {
                            withCredentials([usernamePassword(credentialsId: 'docker-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                                sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'
                                sh "docker push ${DOCKER_IMAGE_BACKEND}"
                            }
                        }
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

