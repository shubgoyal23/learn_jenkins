def gv

pipeline {
    agent any
    parameters {
        choice(name: 'VERSION', choices: ['1.1.0', '1.2.0', '1.3.0'], description: '')
        booleanParam(name: 'executeTests', defaultValue: true, description: '')
    }
    stages {
        stage('init') {
            steps {
                script {
                    gv = load 'script.groovy' // this load groovy script
                }
            }
        }
        stage('Build') {
            steps {
                script {
                    gv.buildApp
                }
            }
        }
        stage('Test') {
            when { // this is condition check if its fullfill this stage will execute
                expression {
                    BRANCH_NAME == 'dev' && params.executeTests // this branchname will come from env variables
                }
            }
            steps {
                script {
                    gv.testApp
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    gv.deployeApp
                }
            }
        }
    }
}
