# HugoShop
Introduction is a waste...

## Prerequisites
First install Chocolatey (package manager for Windows). Open a PowerShell windows, and run the script.
```
Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
```
Verify that it is installed, by running the command `choco` in PowerShell. It should display the version installed.

Second, install Hugo.
```
choco install hugo -confirm
choco install hugo-extended
```
Verify the installation by running the command `hugo version` in PowerShell. It should display the version installed.

Third, install GO.
```
choco install golang
```
Verify the installation by running the command `go version` in PowerShell. It should display the version installed. Sometimes you might have to close and reopen the PowerShell in order for it to work.

Fourth and last, install NodeJs and NPM.
```
choco install nodejs.install
```
Verify the installation by running the command `node -v` and `npm -v` in PowerShell. You might have to close and reopen PowerShell before it is working.

## Install
Find a suitable location on your file system, and clone the repository like this. 
```
git clone https://github.com/moemar/hugoshop.git
```

Move into the newly created directory. 
```
cd hugoshop
```

Install NPM modules 
```
npm install
```

## Run
Run the solution.
```
hugo server -D
hugo serve -D --ignoreCache
```
If everything is working as expected, the site will be available on `http://localhost:1313/`

## Edit
Open the project in Visual Studio Code, if you dont have it, install it, period. Then run the following command.
```
code .
```

## Changes
Whenever you work with the project, remember to check inn and pull changes.

### Pull
```
git pull
```

### Commit
```
git add .
git commit -m "Write a message"
git push origin
