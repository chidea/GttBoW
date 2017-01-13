# GttBoW
Transparent Git to the Bash on Windows.

It lets you use the git installed on bash on windows as native windows git, without installing git again with silly mingw ssh tools.
Because it uses raw argument as string by bat file(`%*`), you don't need any quote/doublequote brackets for command-argument list.
So you do use commands directly with any brackets you need to use as if it's an windows command.

### Built with
 - golang
  - Because it produces binary without installing entire visual studio monster.

### WARNING
### It doesn't support 32bit windows yet.

### Automatic Installation
 Run `install.bat` as administrator privilege
 It builds and installs as normal mode. You can use ssh-agent with simple modification in `install.bat`. Also, check the note below.

### Manual Installation
 - Get `gbash.exe` by manual building with `go build gbash.go normal.go`.
  - If you are using ssh-agent, consider using  `go build sbash.go ssh-agent.go`. This lets you share ssh passwords from pre-opened background bash terminal. Check ssh-agent section for more information.
 - Move `gbash.exe` into `%WINDIR%\System32\`(cmd) or `$env:WINDIR\System32\`(powershell) or just under any directory which PATH environment variable directs.
 - Do the same move with `batch/git.bat` and `batch/bbatch.bat` file.
 - Now `git [command]` will work as native windows command.
 - Do the same with other `.bat` files in `batch` directory to use it like native windows binary.

### Note for GCC
 - To use `gcc.bat`, do `sudo apt install gcc-mingw-w64-x86-64` in ubuntu first.

### Note for ssh-agent
This is the content that I place in `~/.bashrc` and `~/.zshrc`.
```
AGENT_TEST=$(pgrep ssh-agent)
if [[ $AGENT_TEST = "" ]]; then ssh-agent > ~/.ssh/ssh-agent.sh; fi
. ~/.ssh/ssh-agent.sh
if [[ $AGENT_TEST = "" ]]; then ssh-add; fi
```

### Note for piping
To use piping, do it like `bbash 'echo 1 > _tmp'` or `bbash "echo 1 > _tmp"` which is composed of `bbash` and only one argument wrapped with quotes.
Piping directly from windows to ubuntu is available since bd5aa169fc075a69a7a74a45959edcb2219f606a

### Currrently on testing
 - [x] git
  - [x] git commit/push
  - [x] git clone to windows directory and path
 - [x] GCC build test
 - [ ] Vundle on gvim on windows
  - [x] PluginInstall
  - [ ] PluginSearch
 - [ ] Golang
  - [ ] git related e.g)`go get`
  - [ ] gcc related (cgo)
