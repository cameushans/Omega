pipeline {
    agent any

    stages {
        stage("Checkout") {
            steps {   
                git url:  'https://github.com/cameushans/Omega.git'
            }
        }
        stage('build') {
            steps {
                echo 'Hello World'
            }
        }
    }
}  