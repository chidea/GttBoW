cat>~/.init-ssh-agent<<'EOF'
#!/usr/bin/env bash
if [[ $(pgrep ssh-agent) = "" ]]
  then eval `ssh-agent > ~/.ssh/ssh-agent.sh`
  source ~/.ssh/ssh-agent.sh > /dev/null
  for KEY in `ls ~/.ssh/*.pub | sed s/\.pub//`
  do ssh-add $KEY
  done
else
  source ~/.ssh/ssh-agent.sh > /dev/null
fi
EOF

chmod +x ~/.init-ssh-agent
