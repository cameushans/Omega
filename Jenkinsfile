pipeline {
    agent  {
        docker { image 'golang:latest' } 
    }

    stages {
        stage("Checkout") {
            steps {   
                git url:  'https://github.com/cameushans/Omega.git'
            }
        }
        stage('build') {
            steps {
                sh 'go build /Omega/main.go'
            }
        }
    }
}  