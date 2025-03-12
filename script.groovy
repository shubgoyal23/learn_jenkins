def buildApp() {
    echo 'Building the project'
}
def testApp() {
    echo 'testing the project'
}
def deployeApp() {
    echo 'deploying the project'
    echo "project version ${params.VERSION}"
}
