pipeline {
  agent any
  options {
    buildDiscarder(logRotator(numToKeepStr:'5', artifactNumToKeepStr:'5', artifactDaysToKeepStr:'7'))
    durabilityHint('PERFORMANCE_OPTIMIZED')
    retry(1)
    skipDefaultCheckout()
    timestamps()
  }
  stages {
    stage ('Checkout') {
      environment {
        def TAG = sh returnStdout: true, script: "git tag -l | tail -n1"
      }
      steps {
        checkout([$class:'GitSCM', branches: [[name: '*/master'], [name: '*/develop'], [name: '*/release']], 
        doGenerateSubmoduleConfigurations:false, extensions:[], submoduleCfg:[],
        userRemoteConfigs:[[ url:'https://github.com/samfil-technohub/samuel-nwoye-website.git']]])
        sh('''
            git config user.name 'knoxknot'
            git config user.email 'samuel.nwoye@yahoo.com'
            git checkout -B develop
        ''')
        echo "Using Git Tag: ${env.GIT_BRANCH}"   
        sh 'printenv' 
      }
    }
    stage ('Test') {
      agent { 
        docker { 
          image 'golang' 
          args ' -e GOCACHE=/tmp/.cache -e GO111MODULE=on -e GOOS=linux -e GOARCH=amd64 '
        } 
      }
      steps {
        sh 'go version'
        sh 'go mod download'
        sh 'go test -v'
      }
    }
    stage ('Build') {
      agent { 
        docker { 
          image 'golang'
          args ' -e GOCACHE=/tmp/.cache -e GO111MODULE=on -e GOOS=linux -e GOARCH=amd64 '
        } 
      }
      steps {
        sh 'go mod download'
        sh 'go build main.go'
        withCredentials([usernamePassword(credentialsId: 'github', passwordVariable: 'github_password', usernameVariable: 'github_username')]) {
          // sh("git tag -a some_tag -m 'Jenkins'")
          script {
            env.encodedPass=URLEncoder.encode(github_password, "UTF-8")
          }
          sh('''
            git pull origin develop
            // git add .
          ''')
          sh("git commit -am 'update: build ${env.BUILD_NUMBER} is successful'")
          sh('git push https://${github_username}:${encodedPass}@github.com/samfil-technohub/samuel-nwoye-website.git')
        } 
      }
    }
    // stage('Push') {
    //   environment { 
    //     GIT_AUTH = credentials('github') 
    //   }
    //   steps {
    //     sh('''
    //         git config --local credential.helper "!f() { echo username=\\$GIT_AUTH_USR; echo password=\\$GIT_AUTH_PSW; }; f"
    //         git commit -am "update: successful go build for ${env.BUILD_NUMBER}"
    //         git push
    //     ''')
    //   }
    // }
    // stage('Deliver') {
    //   steps {
    //     withCredentials([usernamePassword(credentialsId: 'github', passwordVariable: 'github_password', usernameVariable: 'github_username')]) {
    //       // sh("git tag -a some_tag -m 'Jenkins'")
    //       script {
    //         env.encodedPass=URLEncoder.encode(github_password, "UTF-8")
    //       }
    //       sh('git add .')
    //       sh("git commit -m 'update: build ${env.BUILD_NUMBER} is successful'")
    //       sh('git push https://${github_username}:${encodedPass}@github.com/samfil-technohub/samuel-nwoye-website.git')
    //     }
    //   }
    // }
    stage ('Clean Up'){
      steps {
        cleanWs()
      }
    }
    stage('Deploy') {
      when {
        branch 'master' 
      }
      steps {
        echo "Deploying to Production"
      }
    }
  }
  post {
    success {
      slackSend (channel: '#mymonitor', color: 'good', message: "*${currentBuild.currentResult}:* _Job_ ${env.JOB_NAME}; _Build_ ${env.BUILD_NUMBER}\n *Visit:* ${env.BUILD_URL}")
    }
    failure {
      slackSend (channel: '#mymonitor', color: 'danger', message: "*${currentBuild.currentResult}:* _Job_ ${env.JOB_NAME}; _Build_ ${env.BUILD_NUMBER}\n *Visit:* ${env.BUILD_URL}")
    }
  }
}