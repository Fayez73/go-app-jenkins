pipeline {
    agent any

        environment {
        DOCKER_IMAGE_FRONTEND = "fayez74/react-frontend:latest"
        DOCKER_IMAGE_BACKEND = "fayez74/go-backend:latest"
        DOCKER = "docker"
        PATH = "${env.HOME}/.local/bin:${env.PATH}"
        KUBECONFIG = "${env.HOME}/.kube/config"
    }

    stages {
            stage('Cleanup') {
                steps {
                    cleanWs()
                }
            }
            stage('Install kubectl') {
                steps {
                    sh '''
                    # Make a local bin directory in Jenkins home
                    mkdir -p $HOME/.local/bin

                    # Download kubectl (fixed version)
                    curl -LO "https://dl.k8s.io/release/v1.27.6/bin/linux/amd64/kubectl"

                    # Make it executable
                    chmod +x kubectl

                    # Move to local bin
                    mv kubectl $HOME/.local/bin/

                    # Add local bin to PATH for this shell
                    export PATH=$HOME/.local/bin:$PATH

                    # Verify installation
                    kubectl version --client
                    '''
                }
            }
            stage('Fix Minikube kubeconfig paths') {
                steps {
                    sh '''
                    kubectl get nodes
                    '''
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

