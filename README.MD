# GoFrame Template For SingleRepo

Quick Start: 
- https://goframe.org/pages/viewpage.action?pageId=1114399

#gen pb
#gf gen pb -p manifest/protobuf/taskreport/v1 -a api/taskreport/v1 -c internal/app/server/controller/
#gen service
#gf gen service -s internal/app/server/logic -d internal/app/server/service
#gen dao
#gf gen dao -p internal/app/server