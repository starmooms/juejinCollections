ARG BASE_IMAGE=mcr.microsoft.com/windows/servercore:ltsc2019
FROM $BASE_IMAGE

ENV chocolateyVersion=1.4.0

# Install Chocolatey
RUN powershell -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'));" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"
