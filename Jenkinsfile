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
                git branch: 'main', credentialsId: 'github', url: 'https://github.com/tejaswankalluri/BuyLeaf-Backend'
            }
        }
    }
}