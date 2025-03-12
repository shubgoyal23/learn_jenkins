pipeline{
    agent any
    environment{
        NEW_VERSION = '1.0.5' // creating our own env variables
        SERVER_CRED = credentials('server-cred') // we get credentials from jenkins credential manager
    }
    parameters{
        choice(name: "VERSION", choices: ['1.1.0', '1.2.0', '1.3.0'], description: '')
        booleanParam(name: "executeTests", defaultValue: true, description: "")
    }
    stages{ // build stages

        stage('Build'){
            steps{
                echo 'Building the project'
                echo "Building version ${NEW_VERSION}" // using 
            }
        }

        stage('Test'){
            when { // this is condition check if its fullfill this stage will execute
                expression {
                    BRANCH_NAME=='dev' && param.executeTests // this branchname will come from env variables
                }
            }
            steps{
                echo 'Testing the project'
            }
        }

        stage('Deploy'){
            steps{
                echo 'Deploying the project'
                withcredentials([usernamePassword(credentialsId: 'server-cred', passwordVariable: 'PASSWORD', usernameVariable: 'USERNAME')]){
                    sh 'echo $USERNAME'
                } // we can use wrapper like this to get env variables
            }
        }
    }
}