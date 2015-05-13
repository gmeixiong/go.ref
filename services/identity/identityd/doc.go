// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated via go generate.
// DO NOT UPDATE MANUALLY

/*
Command identityd runs a daemon HTTP server that uses OAuth to create
security.Blessings objects.

Starts an HTTP server that brokers blessings after authenticating through OAuth.

To generate TLS certificates so the HTTP server can use SSL:
  go run $(go list -f {{.Dir}} "crypto/tls")/generate_cert.go --host <IP address>

To use Google as an OAuth provider the -google-config-* flags must be set to
point to the a JSON file obtained after registering the application with the
Google Developer Console at https://cloud.google.com/console

More details on Google OAuth at:
  https://developers.google.com/accounts/docs/OAuth2Login

More details on the design of identityd at:
  https://v.io/designdocs/identity-service.html

Usage:
   identityd [flags]

The identityd flags are:
 -assets-prefix=
   Host serving the web assets for the identity server.
 -external-http-addr=
   External address on which the HTTP server listens on.  If none is provided
   the server will only listen on -http-addr.
 -google-config-android=
   Path to the JSON-encoded OAuth client configuration for Android applications
   that obtain blessings from this server (via the
   OAuthBlesser.BlessUsingAccessToken RPC) from this server.
 -google-config-chrome=
   Path to the JSON-encoded OAuth client configuration for Chrome browser
   applications that obtain blessings from this server (via the
   OAuthBlesser.BlessUsingAccessToken RPC) from this server.
 -google-config-web=
   Path to JSON-encoded OAuth client configuration for the web application that
   renders the audit log for blessings provided by this provider.
 -http-addr=localhost:8125
   Address on which the HTTP server listens on.
 -mount-prefix=identity
   Mount name prefix to use.  May be rooted.
 -sql-config=
   Path to file containing a json object of the following form:
      {
       "dataSourceName": "[username[:password]@][protocol[(address)]]/dbname", (the connection string required by go-sql-driver)
       "tlsServerName": "serverName", (the domain name of the sql server for ssl)
       "rootCertPath": "/path/server-ca.pem", (the root certificate of the sql server for ssl)
       "clientCertPath": "/path/client-cert.pem", (the client certificate for ssl)
       "clientKeyPath": "/path/client-key.pem" (the client private key for ssl)
      }
 -tls-config=
   Comma-separated list of TLS certificate and private key files, in that order.
   This must be provided.

The global flags are:
 -alsologtostderr=true
   log to standard error as well as files
 -log_backtrace_at=:0
   when logging hits line file:N, emit a stack trace
 -log_dir=
   if non-empty, write log files to this directory
 -logtostderr=false
   log to standard error instead of files
 -max_stack_buf_size=4292608
   max size in bytes of the buffer to use for logging stack traces
 -stderrthreshold=2
   logs at or above this threshold go to stderr
 -v=0
   log level for V logs
 -v23.credentials=
   directory to use for storing security credentials
 -v23.i18n-catalogue=
   18n catalogue files to load, comma separated
 -v23.metadata=<just specify -v23.metadata to activate>
   Displays metadata for the program and exits.
 -v23.namespace.root=[/(dev.v.io/role/vprod/service/mounttabled)@ns.dev.v.io:8101]
   local namespace root; can be repeated to provided multiple roots
 -v23.permissions.file=map[]
   specify a perms file as <name>:<permsfile>
 -v23.permissions.literal=
   explicitly specify the runtime perms as a JSON-encoded access.Permissions.
   Overrides all --v23.permissions.file flags.
 -v23.proxy=
   object name of proxy service to use to export services across network
   boundaries
 -v23.tcp.address=
   address to listen on
 -v23.tcp.protocol=wsh
   protocol to listen with
 -v23.vtrace.cache-size=1024
   The number of vtrace traces to store in memory.
 -v23.vtrace.collect-regexp=
   Spans and annotations that match this regular expression will trigger trace
   collection.
 -v23.vtrace.dump-on-shutdown=true
   If true, dump all stored traces on runtime shutdown.
 -v23.vtrace.sample-rate=0
   Rate (from 0.0 to 1.0) to sample vtrace traces.
 -vmodule=
   comma-separated list of pattern=N settings for file-filtered logging
*/
package main