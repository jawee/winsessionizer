# winsessionizer
terminal app to navigate between projects

powershell bind command to shortcut
```powershell
Set-PSReadlineKeyHandler -Chord Ctrl+f -ScriptBlock {
    [Microsoft.PowerShell.PSConsoleReadLine]::RevertLine()
    [Microsoft.PowerShell.PSConsoleReadLine]::Insert('work')
    [Microsoft.PowerShell.PSConsoleReadLine]::AcceptLine()
}
```

