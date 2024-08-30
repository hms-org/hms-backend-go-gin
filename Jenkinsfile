pipeline {
    agent any

    environment {
        DOCKER_IMAGE_NAME = ""
        DEPLOY_PORT = ""
    }

    stages {
        stage('Prepare') {
            steps {
                script {
                    if (env.BRANCH_NAME == 'dev') {
                        DOCKER_IMAGE_NAME = "hms-backend-golang-dev"
                        DEPLOY_PORT = "8011"
                    } else if (env.BRANCH_NAME == 'uat') {
                        DOCKER_IMAGE_NAME = "hms-backend-golang-uat"
                        DEPLOY_PORT = "8012"
                    } else if (env.BRANCH_NAME == 'prod') {
                        DOCKER_IMAGE_NAME = "hms-backend-golang-prod"
                        DEPLOY_PORT = "8013"
                    } else {
                        error("Unknown branch for deployment!")
                    }

                    echo "Deploying branch: ${env.BRANCH_NAME}"
                    echo "Using Docker image: ${DOCKER_IMAGE_NAME}"
                    echo "Deploying on port: ${DEPLOY_PORT}"
                }
            }
        }

        stage('Clone repository') {
            steps {
                git branch: "${env.BRANCH_NAME}", url: 'https://github.com/hms-org/hms-backend-golang.git'
            }
        }

        stage('Build Go Application') {
            steps {
                dir('src') { // Change directory to /src
                    script {
                        // Build the Go application locally before Dockerizing
                        sh 'go mod download'
                        sh 'go build -o main .'
                    }
                }
            }
        }

        stage('Remove Previous Docker') {
            steps {
                script {
                    sh "docker stop ${DOCKER_IMAGE_NAME} || true"
                    sh "docker rm ${DOCKER_IMAGE_NAME} || true"
                    sh "docker rmi ${DOCKER_IMAGE_NAME} || true"
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${DOCKER_IMAGE_NAME}:${env.BUILD_ID}")
                }
            }
        }

        stage('Deploy to Server') {
            steps {
                script {
                    sh """
                    docker run -d --name ${DOCKER_IMAGE_NAME} \
                    -p ${DEPLOY_PORT}:8000 \
                    ${DOCKER_IMAGE_NAME}:${env.BUILD_ID}
                    """
                }
            }
        }
    }
    
    post {
        success {
            echo 'Build and deploy successful!'
        }
        failure {
            echo 'Build or deploy failed.'
        }
    }
}