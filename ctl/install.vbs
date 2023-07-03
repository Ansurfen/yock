Set wshShell = CreateObject("WScript.Shell")
Set env = wshShell.Environment("User")

path = env("PATH")
userProfilePath = wshShell.ExpandEnvironmentStrings("%USERPROFILE%")

newPath = userProfilePath & "\.yock\mnt"
if InStr(1, path, newPath, vbTextCompare) = 0 Then
    path = newPath & ";" & path
End If

env("PATH") = path

MsgBox "Install success!", vbOKOnly, "Yock"
