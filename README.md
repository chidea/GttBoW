# GttBoW
Transparent Git to the Bash on Windows.

It lets you use the git installed on bash on windows as native windows git, without installing git again with silly mingw ssh tools.

# Built with
 - golang
  - Because it produces binary without installing entire visual studio monster.

# Installation
 - Get `sbash.exe` by manual building with `go build sbash.exe`. If you use ssh-agent, consider using `sabash.go` and renaming its output as `sbash.exe`.
 - Move `sbash.exe` into `%WINDIR%\System32\` or `$env:WINDIR\System32\` or just any PATH environment variable directs.
 - Do the same with `batch/git.bat` file.
 - Now `git [command]` will work as native windows command.

# Simpler Batch Usage
 - Copy all `sbash.bat` file under `batch` directory to `%WINDIR%\system32\`.
 - Do the same with `git.bat` or other `.bat` files to use it like native windows binary.
 - Feel free to open `git.bat` and create link batch to any other bash binaries. It's clean short and easy.
 - Check out `sabash.bat` if you're using `ssh-agent`. Open a bash shell and `ssh-agent` in background will let you skip typing ssh passwords.
 - Commands like `ssh` without inline commands which waits for user input are available.

# ETC.
 - To use `gcc.bat`, do `sudo apt install gcc-mingw-w64-x86-64` in ubuntu first.

# Currrently on testing
 - Building go-sqlite3 on windows

# Testing
 - Git

|          Command           | Works |
|:--------------------------:|:-----:|
| clone to windows directory |   x   |
| others                     |   o   |
 
 - Vundle on GVim on windows

| Command | Works |
|:--------------------------:|:-------:|
| PluginInstall | o |
| PluginSearch | o |

 - Go on windows

| Command | Works |
|:--------------------------:|:-------:|
| go get | x |
| go install | ? |

