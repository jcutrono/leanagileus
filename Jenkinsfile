#!groovy

import groovy.json.JsonOutput

node {
	// Mark the code checkout 'stage'....
	stage ('checkout') {
		// Get some code from a GitHub repository
		git url: 'https://github.com/jcutrono/leanagileus', branch: 'develop'
		sh 'git clean -fdx; sleep 4;'
	}
	
	def goTool = tool 'go'	
	
	stage ('build') {	
		env.GOPATH="${env.HOME}/go-plugins"
		
		sh "go get github.com/gorilla/mux"
		sh "go get gopkg.in/mgo.v2"
		sh "go build -i"
	}
	
	stage ('test') {		
		sh "go test"
	}
	
	notifySlack("build succeeded")

	stage ('merge to master') {
		withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'mylogin',
                    usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
			sh "git checkout master"
			sh "git pull origin master"
			sh "git merge develop"
			sh "git push https://$USERNAME:$PASSWORD@myrepository.biz/file.git master"
		}
	}
}


def notifySlack(text) {
    def slackURL = 'https://hooks.slack.com/services/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'
    def payload = JsonOutput.toJson([text      : text,
                                     channel   : "#ci",
                                     username  : "jenkins",
                                     icon_emoji: ":jenkins:"])
    sh "curl -X POST --data-urlencode \'payload=${payload}\' ${slackURL}"
}