def docker_hub = [
  usernamePassword(
    usernameVariable: 'REGISTRY_USERNAME',
    passwordVariable: 'REGISTRY_PASSWORD',
    credentialsId   : 'docker-hub-generic-up',
  )
]


pipeline {
  agent {
    kubernetes {
      yaml '''
        apiVersion: v1
        kind: Pod
        spec:
          containers:
            - name: jnlp
              image: docker:latest
              command:
                - apk add --update make
                - cat
              tty: true
              resources:
                requests:
                  memory: "512Mi"
                  cpu: "1g"
              volumeMounts:
                - name: docker-socket
                  mountPath: /var/run
            - name: docker-daemon
              image: docker:dind
              resources:
                limits: 
                  memory: "32Gi"
                requests:
                  memory: "16Gi"
              securityContext:
                privileged: true
              volumeMounts:
                - name: docker-socket
                  mountPath: /var/run
              imagePullSecrets:
                - name: "regcred-msr"
          volumes:
            - name: docker-socket
              emptyDir: {}
      '''
    }
  }

  stages {
    stage("Build") {
      steps {
        sh "make unit-test"
        sh "make lint || echo 'Lint failed'"
        sh "make build-all"
      }
    }

/*    stage("Smoke test") {
     parallel {
       stage("'Register' subcommand") {
         stages {
           stage("Register") {
             steps {
               container("runner") {
                 sh "make bin/launchpad"
                 sh "make smoke-register-test"
               }
             }
           }
         }
       }
     }
   } */
  }
}
