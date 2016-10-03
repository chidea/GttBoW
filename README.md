# GttBoW
Transparent Git to the Bash on Windows.

It lets you use the git installed on bash on windows as native windows git, without installing git again with silly mingw ssh tools.

# Built with
 - golang
  - Because it produces binary without installing entire visual studio monster.

# Installation
 - Get `git.exe` by manual building or downloading prebuilt binary release
 - Move `git.exe` into `%WINDIR%\System32\` or `$env:WINDIR\System32\` or just any PATH environment variable directs.

# Why not simpler batch script?
 - It's totally doable but because of BoW's bug, it'll produce errors on pipe use cases which are used by many tools for automation.
 - 
