if [ "$#" -gt 0 ]
then
  arguments="$*"
  echo "$arguments"
  if [[ $arguments == "sudo"* ]]
  then
  [ "$UID" -eq 0 ] || exec sudo "$@"
  else
  eval "$arguments"
  fi
else
  last_command=`fc -ln -2 -2`
  echo "Last Command: ${last_command}"
  suggestion=`please "${last_command}"`
  echo "Usage: ..."
  exit 1
fi
