# selenium grid docker

https://nitinbhardwaj6.medium.com/selenium-grid-with-docker-c8ecb0d8404
https://nitinbhardwaj6.medium.com/selenium-grid-with-debug-docker-containers-23eceb6ab17f

## Pre-reqs
```
# Allow sudo
sudo ls

# Java SDKs
curl -s "https://get.sdkman.io" | bash
source "/home/vladimir/.sdkman/bin/sdkman-init.sh"
sdk install java 8.0.292-zulu
sdk install maven 3.8.1

# VNC Viewer
wget -O vnc-viewer.deb https://www.realvnc.com/download/file/viewer.files/VNC-Viewer-6.20.529-Linux-x64.deb
sudo dpkg -i vnc-viewer.deb
rm vnc-viewer.deb

# Install Selenium IDE in Chrome (Recorder)
https://chrome.google.com/webstore/detail/selenium-ide/mooikfkahbdckldjjndioackbalphokd?hl=en
```

## Start Selenium Grid
```
cd infra/selenium/docker
docker-compose -f docker-compose.yml up -d
docker ps
docker stats
# docker-compose -f docker-compose.yml down
```

## Check it is running
http://localhost:4444/grid/console

## Run Tests!
mvn clean test -Dsurefire.suiteXmlFiles=testng.xml

## Connect via VNC Viewer
Adress:   localhost:5900 
Password: secret




# examples projects:
https://github.com/ChristianAA/selenium-springboot-example
https://github.com/bo32/spring-cucumber-selenium-sample