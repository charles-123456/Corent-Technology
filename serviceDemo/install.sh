dir=$(dirname $0)
cd $dir
startinstall() 
{
cd ..
chmod -R 755 bin
cd bin
./serviceDemo -service=install
estatus=$?
if [ $estatus -ne 0 ];then
echo "*************** serviceDemo Installation Failed **************"
exit 1
else
echo "*************** serviceDemo Installed Successfully **************"
fi
./serviceDemo -service=start
estatus=$?
if [ $estatus -ne 0 ];then
echo "*************** Starting serviceDemo Failed **************"
exit 1
fi
echo "*************** serviceDemo Started Successfully **************"
exit 0
}
startinstall
