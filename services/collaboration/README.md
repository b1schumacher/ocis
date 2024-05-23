# Collaboration

The collaboration service connects ocis with document servers such as Collabora and ONLYOFFICE using the WOPI protocol.

Since this service requires an external document server, it won't start by default when using `ocis server`. You must start it manually with the `ocis collaboration server` command, or setting `OCIS_ADD_RUN_SERVICES=collaboration`.

oCIS will automatically routes all `/wopi` traffic to the collaboration service if it is running.

For demo purposes, more than one collaboration service can be started by running a dedicated `ocis collaboration server` instance on a different domain configured by setting e.g. `COLLABORATION_WOPIAPP_WOPISRC=https://otheroffice-wopi.example.com`. It does not require a proxy service and can directly handle requests.

## Requirements

The collaboration service requires the target document server (ONLYOFFICE, Collabora, etc.) to be up and running. Additionally, some Infinite Scale services are also required to be running in order to register the GRPC service for the `open in app` action in the webUI. The following internal and external services need to be available:

* External document server.
* The gateway service.
* The app-provider service.

If any of the named services above have not been started or are not reachable, the collaboration service won't start. For the binary or the docker release of Infinite Scale, check with the `ocis list` command if they have been started. If not, you must start them manually upfront before starting the collaboration service.

## WOPI Configuration

There are a few variables that you need to set:

* `COLLABORATION_WOPIAPP_ADDR`:\
  The URL of the WOPI app (onlyoffice, collabora, etc).\
  For example: `https://office.example.com`.

* `COLLABORATION_WOPIAPP_INSECURE`:\
  In case you are using a self signed certificate for the WOPI app you can tell the collaboration service to allow an insecure connection.

* `COLLABORATION_WOPIAPP_WOPISRC`:\
  The external address of the collaboration service. The target app (onlyoffice, collabora, etc) will use this address to read and write files from Infinite Scale. It defaults to `OCIS_URL` and only needs to be changed when running a dedicated collaboration service to handle requests for a different WOPI app.\
  For example: `https://otheroffice-wopi.example.com`.

The rest of the configuration options available can be left with the default values.
