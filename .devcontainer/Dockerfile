FROM mcr.microsoft.com/devcontainers/python:1-3.10-bullseye

# Here I use the USER from the FROM image
ARG USERNAME=vscode
ARG GROUPNAME=vscode

# Here I use the UID/GID from _my_ computer, as I'm _not_ using Docker Desktop
ARG USER_UID=501
ARG USER_GID=20

RUN groupmod --gid $USER_GID -o $GROUPNAME \
    && usermod --uid $USER_UID --gid $USER_GID $USERNAME \
    && chown -R $USER_UID:$USER_GID /home/$USERNAME

# General setup which is "cached" for convenience
RUN pip install -U pip setuptools
RUN pip install -U poetry
RUN pip install -U "ml-metadata==1.14.0"
RUN pip install -U robotframework
RUN pip install -U robotframework-requests
RUN pip install -U PyYAML