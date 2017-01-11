set shell = WScript.CreateObject("WScript.shell")

args = ""
Set objArgs = WScript.Arguments
For I = 0 to objArgs.Count - 1
   args = args & " " &  objArgs(I)
Next
shell.Run "bash.exe" & args , 0, True
