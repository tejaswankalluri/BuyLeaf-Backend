pipeline {
    agent any
    tools {
        jdk 'java17'
        maven 'maven3'
    }
    stages {
        stage("Cleanup Workspace"){
            steps {
                cleanWs()
            }
        }
        stage("Checkout from SCM"){
            steps {
                git branch: 'main', credentialsId: 'tejaswankalluri-gh', url: 'https://github.com/tejaswankalluri/BuyLeaf-Backend'
            }
        }
        stage("Build Application"){
            steps {
                sh "mvn clean package"
            }
        }
        stage("Test Aplication"){
            steps {
                sh "mvn test"
            }
        }

    }
}