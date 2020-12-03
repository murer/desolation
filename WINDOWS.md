# Desolation on Windows

## Prepare

Use a PowerShell as Admin

Optional: Enable PowerShell scripting

```ps1
Set-ExecutionPolicy Unrestricted
```

Install openssh client

```ps1
Add-WindowsCapability -Online -Name OpenSSH.Client~~~~0.0.1.0
```

## Download

Use non Admin PowerShell

```ps1
wget "https://github.com/murer/desolation/releases/download/edge/desolation-windows-amd64.exe" -O desolation.exe
```

## Connect

Use non Admin PowerShell

```ps1
ssh -o "ProxyCommand .\desolation.exe guest %h %p" "user@host" command 
```
