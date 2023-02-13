#!/usr/bin/python
# coding=utf-8

import os
import sys

def _check_python_version():
    major_ver = sys.version_info[0]
    if major_ver < 2:
        print ("The python version is %d.%d. But python 2.x+ is required. (Version 2.7 is well tested)\n"
               "Download it here: https://www.python.org/" % (major_ver, sys.version_info[1]))
        return False

    return True

if __name__ == '__main__':
    if not _check_python_version():
        exit()

    PORT = 8000

    if sys.version_info.major == 2:
        import SimpleHTTPServer
        import SocketServer

        httpd = SocketServer.TCPServer(("", PORT), SimpleHTTPServer.SimpleHTTPRequestHandler)

        print "[Python2][local serving at port]", PORT
        httpd.serve_forever()
    else:
        import http.server

        httpd = http.server.HTTPServer(("", PORT), http.server.SimpleHTTPRequestHandler)

        print "[Python3][local serving at port]", PORT
        httpd.serve_forever()