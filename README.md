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

### Installation
 Run `install.bat` as administrator privilege.
 Installation script will ask for some options.
 You can use ssh-agent with simple modification in `install.bat`. Also, check the note below.

### Note for GCC
 - To use `gcc.bat`, do `sudo apt install gcc-mingw-w64-x86-64` in ubuntu first. `install.bat` asks if you want it installed.

### Note for ssh-agent
Run following script into powershell or bash. You can simply run `install-ssh-agent.ps1` instead.
```
bash -c 'cat<<''EOF''>~/.init-ssh-agent
#!/usr/bin/env bash
if [[ \$(pgrep ssh-agent) = \"\" ]]
  then eval \`ssh-agent > ~/.ssh/ssh-agent.sh\`
  source ~/.ssh/ssh-agent.sh
  for KEY in \`ls ~/.ssh/*.pub | sed s/\.pub//\`
  do ssh-add \$KEY
  done
else
  source ~/.ssh/ssh-agent.sh > /dev/null
fi
EOF
chmod +x ~/.init-ssh-agent'
```
Line below is the content to be placed in `~/.bashrc` or `~/.zshrc`
```
. ~/.init-ssh-agent
```
For example of bashrc,
```
bash -c 'echo . \~/.init-ssh-agent >> ~/.bashrc'
```

### Note for piping
To use piping, do it like `bbash 'echo 1 > _tmp'` or `bbash "echo 1 > _tmp"` which is composed of `bbash` and only one argument wrapped with quotes.
Piping directly from windows to ubuntu is available since bd5aa169fc075a69a7a74a45959edcb2219f606a

### Currrently on testing
 - [x] live standard stream support for apps like `ssh`
 - [x] git
  - [x] git commit/push
  - [x] git clone to windows directory and path
 - [x] GCC build test
 - [x] Plug on neovim on windows
  - [x] PlugInstall
 - [x] Golang
  - [x] git related get/install
  - [ ] gcc related build (cgo)
 - [ ] Xming
  - [ ] ssh.exe replacement for remote X11 forwarding
