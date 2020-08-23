Set-Location -Path C:\Projects\LogitechKeyboardLED

$ProcessActive = (Get-Process logitechKeyboardLed -ErrorAction SilentlyContinue).Id
if($null -eq $ProcessActive)
{
    Start-Process ./logitechKeyboardLed.exe -NoNewWindow
}
else
{
    Stop-Process $ProcessActive
}
