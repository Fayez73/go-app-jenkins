pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                // Pull latest code from GitHub
                checkout([$class: 'GitSCM', 
                          branches: [[name: '*/main']], 
                          userRemoteConfigs: [[url: 'https://github.com/Fayez73/go-app-jenkins', credentialsId: 'github-creds']]])
            }
        }

        stage('Run Command') {
            steps {
                // Replace this with whatever command you want to run
                sh 'echo "Hello from Jenkins! Running commands on pushed code..."'
                sh 'ls -la'  // lists the repo files
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
