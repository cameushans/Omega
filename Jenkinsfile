pipeline {
    agent any
    tools {
        go 'go'
    }

    stages {
        stage("Checkout") {
            steps {   
                git url: 'https://github.com/cameushans/Omega.git'
            }
        }
        stage('build') {
            steps {
                sh 'go run ${WORKSPACE}/main.go'
            }
        }
    }

    post { always { echo 'This is your best pipeline'}}
}  