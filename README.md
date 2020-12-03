# Desolation
De-isolation network of VDI, RDP, VNC and VMs through display and keyboard

With desolation you can create a connection beteween the remote/guest and the host through the screen and keyboard.

## Use Case

### On guest screen (windows, linux or mac):

```shell
ssh -o "StrictHostKeyChecking=no" -o 'CommandProxy desolation guest %h %p' -g -D 5005 -R 5006 host command
```

Yet on guest, open a browser to be operated by the host: http://localhost:5010/

You should see a initial qrcode on top of a input text.

### On host screen (linux only):

You will start the operator on host screen:

```shell
desolation host
```

You have 5 seconds to position the cursor in the input text of guest.

The **desolation host** process will connect with the ssh server. And it will send commands to guest using text input and receive replies through qrcodes. 

### In the end

You have a socks proxy in guest `-D 5005` forwarded through the host and another socks proxy in host `-R 5006` reverse-forwarded through the guest.

With `-g` you make these socks proxies available on `0.0.0.0`.
