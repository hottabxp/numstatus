go build numstatus.go
mv numstatus app/usr/bin/
dpkg-deb --build app numstatus_0.1-1_amd64.deb