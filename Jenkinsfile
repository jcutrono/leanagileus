node {
	// Mark the code checkout 'stage'....
	stage ('checkout') {
		// Get some code from a GitHub repository
		git url: 'https://github.com/jcutrono/itpalooza16'
		sh 'git clean -fdx; sleep 4;'
	}
	
	// Get the maven tool.
	// ** NOTE: This 'mvn' maven tool must be configured
	// **       in the global configuration.	
	def goTool = tool 'go'	
	
	stage ('build') {
		// set the version of the build artifact to the Jenkins BUILD_NUMBER so you can
		// map artifacts to Jenkins builds		
		env.GOPATH="${env.HOME}/go-plugins"
		
		sh "go get ./..."
		sh "go build -i"
	}
	
	stage ('test') {
		parallel 'test': {
			sh "go test -v"
		}, 'verify': {
			sh "${mvnHome}/bin/mvn verify; sleep 3"
		}
	}

	stage ('archive') {
		archive 'target/*.jar'
	}
}


node {
	stage ('deploy Canary') {
		sh 'echo "write your deploy code here"; sleep 5;'
	}
	
	stage ('deploy Production') {
		input 'Proceed?'
		sh 'echo "write your deploy code here"; sleep 6;'
		archive 'target/*.jar'
	}
}

// Add whichever params you think you'd most want to have
// replace the slackURL below with the hook url provided by
// slack when you configure the webhook
def notifySlack(text) {
    def slackURL = 'https://hooks.slack.com/services/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'
    def payload = JsonOutput.toJson([text      : text,
                                     channel   : "#ci",
                                     username  : "jenkins",
                                     icon_emoji: ":jenkins:"])
    sh "curl -X POST --data-urlencode \'payload=${payload}\' ${slackURL}"
}