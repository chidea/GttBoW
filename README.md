# GttBoW
Transparent Git to the Bash on Windows.

It lets you use the git installed on bash on windows as native windows git, without installing git again with silly mingw ssh tools.
Because it uses raw argument as string by bat file(`%*`), you don't need any quote/doublequote brackets for command-argument list.
So you do use commands directly with any brackets you need to use as if it's an windows command.

# Built with
 - golang
  - Because it produces binary without installing entire visual studio monster.

# Installation
 - Get `gbash.exe` by manual building with `go build gbash.go normal.go`.
  - If you are using ssh-agent, consider using  `go build sbash.go ssh-agent.go`. This lets you share ssh passwords from pre-opened background bash terminal. Check ssh-agent section for more information.
 - Move `gbash.exe` into `%WINDIR%\System32\`(cmd) or `$env:WINDIR\System32\`(powershell) or just under any directory which PATH environment variable directs.
 - Do the same move with `batch/git.bat` and `batch/bbatch.bat` file.
 - Now `git [command]` will work as native windows command.
 - Do the same with other `.bat` files in `batch` directory to use it like native windows binary.

# Note for GCC
 - To use `gcc.bat`, do `sudo apt install gcc-mingw-w64-x86-64` in ubuntu first.

# Note for ssh-agent
This is the content that I place in `~/.bashrc` and `~/.zshrc`.
```
AGENT_TEST=$(pgrep ssh-agent)
if [[ $AGENT_TEST = "" ]]; then ssh-agent > ~/.ssh/ssh-agent.sh; fi
. ~/.ssh/ssh-agent.sh
if [[ $AGENT_TEST = "" ]]; then ssh-add; fi
```

# Currrently on testing
 - Building go-sqlite3 on windows

# Testing
 - Git

| Command | Works |
|:--------------------------:|:-----:|
| clone to windows directory | x |
| others | o |
 
 - Vundle on GVim on windows

| Command | Works |
|:--------------------------:|:-------:|
| PluginInstall | o |
| PluginSearch | x |

 - Go on windows

| Command | Works |
|:--------------------------:|:-------:|
| go get | x |
| go install | ? |
