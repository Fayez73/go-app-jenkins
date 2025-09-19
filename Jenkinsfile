pipeline {
    agent any

        environment {
        DOCKER_IMAGE_FRONTEND = "fayez74/react-frontend:latest"
        DOCKER_IMAGE_BACKEND = "fayez74/go-backend:latest"
        DOCKER = "docker"
        KUBECONFIG = "/var/jenkins_home/.kube/config"
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
            stage('Install kubectl') {
                steps {
                    sh '''
                        KUBECTL_VERSION=$(curl -s https://dl.k8s.io/release/stable.txt)
                        curl -LO "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl"
                        chmod +x kubectl
                        mv kubectl /usr/local/bin/
                        kubectl version --client

                    '''
                }
            }
            stage('Deploy to Minikube') {
                steps {
                    sh '''
                    export KUBECONFIG=$KUBECONFIG

                    # Apply backend and frontend manifests
                    kubectl apply -f k8s/backend.yaml
                    kubectl apply -f k8s/frontend.yaml
                    '''
                }
            }
            stage('Verify Deployment') {
                steps {
                    sh 'kubectl --kubeconfig=$KUBECONFIG get pods,svc'
                }
            }
            stage('Test Backend Connectivity') {
                steps {
                    sh '''
                    export KUBECONFIG=$KUBECONFIG

                    # Run a temporary curl pod
                    kubectl run curl-test --rm -i --tty --image=curlimages/curl --restart=Never -- \
                    curl -s http://backend-service:4000/api/test
                    '''
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

