from ubuntu:latest

RUN apt-get update
RUN apt-get upgrade -y

#install developer dependencies
RUN apt-get install clang -y
RUN apt-get install cmake -y
RUN apt-get install git -y
RUN apt-get install curl -y
RUN apt-get install zip -y
RUN apt-get install gdb -y
RUN apt-get install pkg-config -y
RUN apt-get install libssl-dev -y
RUN apt-get -y install python3-pip
RUN apt install protobuf-compiler -y
RUN apt-get install jq -y;
RUN apt-get install systemctl -y
RUN apt-get install systemd -y
RUN apt-get install wget -y;

#install vcpkg package manager
RUN git clone --depth 1 https://github.com/microsoft/vcpkg
RUN ./vcpkg/bootstrap-vcpkg.sh

#install packages for project

RUN /vcpkg/vcpkg install nlohmann-json
# RUN /vcpkg/vcpkg install cpr

# Create local user to avoid file permission issues
ARG USERNAME=developer 

ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN apt install sudo -y


RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    #
    # add sudo support
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

USER $USERNAME

RUN pip3 install requests

RUN sudo chown -R $USERNAME /vcpkg 
# set work directory for project
WORKDIR /Project

RUN sudo chown -R $USERNAME /Project 

WORKDIR /tmp


RUN wget https://golang.org/dl/go1.15.12.linux-amd64.tar.gz
RUN sudo tar -C /usr/local -xzf go1.15.12.linux-amd64.tar.gz
WORKDIR /

#COPY ./authserver /home/$USERNAME/go/src/authserver
RUN mkdir /home/$USERNAME/go
RUN echo $USERNAME

RUN sudo chown -R $USERNAME:$USERNAME /home/$USERNAME/go

RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
RUN echo "export GOPATH=~/.go" >> ~/.profile

WORKDIR /temp

COPY ./authserver /temp
RUN sudo /usr/local/go/bin/go build -o authserver

WORKDIR /server
RUN sudo cp /temp/authserver /server/authserver
RUN sudo rm -rf /temp
WORKDIR /Project

# Couldn't get systemd to work, just gonna start a forked process in docker-compose
#COPY ./authserver/authserver.service /etc/systemd/system/authserver.service