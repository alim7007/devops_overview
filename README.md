Add the Jenkins repo using the following command:
sudo wget -O /etc/yum.repos.d/jenkins.repo \
    https://pkg.jenkins.io/redhat-stable/jenkins.repo

Import a key file from Jenkins-CI to enable installation from the package:
sudo rpm --import https://pkg.jenkins.io/redhat-stable/jenkins.io-2023.key

sudo yum upgrade

Install Java:
sudo yum install java-17-amazon-corretto -y

Install Jenkins:
sudo yum install jenkins -y

Enable the Jenkins service to start at boot:
sudo systemctl enable jenkins

Start Jenkins as a service:
sudo systemctl start jenkins

You can check the status of the Jenkins service using the command:
sudo systemctl status jenkins


Open port 8080 in EC2 security group
Go to AWS console → EC2 → Security Groups → Edit inbound rules.

Add:

Custom TCP

Port 8080

Source 0.0.0.0/0 (⚠️ For test only; restrict in production).


Connect to http://<your_server_public_DNS>:8080


Try using HTTP instead of HTTPS
By default Jenkins runs on HTTP, not HTTPS.

sudo cat /var/lib/jenkins/secrets/initialAdminPassword



####
later we need to seperate it.
In this lab EC2 is “all-in-one”,
but real CI/CD splits:
Controller ➜ Agents ➜ SCM.


Install Git
sudo yum update -y
sudo yum install git -y
git --version

Install Docker
sudo yum update -y
sudo yum install docker -y

Start Docker
sudo service docker start

Add Jenkins user to Docker group. This lets Jenkins run docker commands without sudo.
sudo usermod -aG docker jenkins


After that: restart Jenkins to apply new group permissions:
sudo systemctl restart jenkins

Also check that Docker works:
docker version



#############################
if wanna work localy with docker localy:

docker network create jenkins

docker run \
  --name jenkins-docker \
  --rm -d \
  --privileged \
  --network jenkins \
  --env DOCKER_TLS_CERTDIR=/certs \
  -v jenkins-docker-certs:/certs/client \
  -v jenkins-data:/var/jenkins_home \
  docker:dind --storage-driver overlay2

docker run \
  --name jenkins-blueocean \
  --rm -d \
  --network jenkins \
  --env DOCKER_HOST=tcp://jenkins-docker:2376 \
  --env DOCKER_TLS_VERIFY=1 \
  --env DOCKER_CERT_PATH=/certs/client \
  -v jenkins-data:/var/jenkins_home \
  -v jenkins-docker-certs:/certs/client:ro \
  -p 8080:8080 -p 50000:50000 \
  jenkins/jenkins:lts-jdk11
#################################