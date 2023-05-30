# VMware-CSP-GoProxy

This is a Go developed Proxy to automate authentication with VMware CSP services so you can work with VMC services as in a local environment.

![alt text](https://miro.medium.com/v2/resize:fit:1100/format:webp/1*dj7MWgwuwh9klyDyZ-2vPw.png)


All documentation of the project is available [here](https://medium.com/@albert.rabassa/vmware-csp-api-automating-authentication-with-go-963e530f9eb2)


Just edit your `config.json`file accordingly:

```
{
    "csp_url": "https://console.cloud.vmware.com",
    "service_url": "https://vcdr-99-99-99-299.app.vcdr.vmware.com",
    "refreshtoken": "myApiRefreshToken"
  }
```

Run the proxy:

```
% ./vmc-csp-goproxy
2023/05/29 23:20:36 VMC CSP Go Proxy listening on localhost:8889
```

Curl to my local proxy:

```
% curl -XGET -ks https://localhost:8889/api/vcdr/v1alpha/cloud-file-systems  | jq .
{
  "cloud_file_systems": [
    {
      "id": "4d738096-1eca-4a0d-9ad8-fbba33d912ea",
      "name": "My SCFS"
    }
  ]
}
```

Checking the flow of the CSP Proxy in the logs:
```
% ./vmc-csp-goproxy
2023/05/29 11:25:37 VMC CSP Go Proxy listening on localhost:8889
2023/05/29 11:25:43 CSP API Login HTTP Response Status: 200 OK
2023/05/29 11:25:44 HTTP GET https://vcdr-99-99-99-299.app.vcdr.vmware.com/api/vcdr/v1alpha/cloud-file-systems
2023/05/29 11:25:44 RESPONSE body : {"cloud_file_systems":[{"id":"4d738096-1eca-4a0d-9ad8-fbba33d912ea","name":"My SCFS"}]}
2023/05/29 11:25:44 VMC CSP Go Proxy forwarding response to localhost:8889
