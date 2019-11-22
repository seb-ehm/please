echo "$#"
if [ "$#" -gt 0 ]
then
  str="'$*'"
  echo "$str"
else
  last_command=`fc -ln -2 -2`
  echo "Last Command: ${last_command}"
  suggestion=`please "${last_command}"`
  echo "Usage: ..."
  exit 1
fi


