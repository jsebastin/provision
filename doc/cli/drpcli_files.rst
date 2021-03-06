drpcli files
============

Access CLI commands relating to files

Synopsis
--------

Access CLI commands relating to files

Options
-------

::

      -h, --help   help for files

Options inherited from parent commands
--------------------------------------

::

      -d, --debug               Whether the CLI should run in debug mode
      -E, --endpoint string     The Digital Rebar Provision API endpoint to talk to (default "https://127.0.0.1:8092")
      -f, --force               When needed, attempt to force the operation - used on some update/patch calls
      -F, --format string       The serialzation we expect for output.  Can be "json" or "yaml" (default "json")
      -P, --password string     password of the Digital Rebar Provision user (default "r0cketsk8ts")
      -r, --ref string          A reference object for update commands that can be a file name, yaml, or json blob
      -T, --token string        token of the Digital Rebar Provision access
      -t, --trace string        The log level API requests should be logged at on the server side
      -Z, --traceToken string   A token that individual traced requests should report in the server logs
      -U, --username string     Name of the Digital Rebar Provision user to talk to (default "rocketskates")

SEE ALSO
--------

-  `drpcli <drpcli.html>`__ - A CLI application for interacting with the
   DigitalRebar Provision API
-  `drpcli files destroy <drpcli_files_destroy.html>`__ - Delete the
   files [item] on the DRP server
-  `drpcli files download <drpcli_files_download.html>`__ - Download the
   files named [item] to [dest]
-  `drpcli files list <drpcli_files_list.html>`__ - List all files
-  `drpcli files upload <drpcli_files_upload.html>`__ - Upload the files
   [src] as [dest]
