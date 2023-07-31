tar -xf $1.tar -C .
cd yock
chmod 777 yock
chmod 777 yockd
./yock run install.lua
./yock run ../$2
rm ../$1.tar ../$2
