_pls_suggest () {
    echo "$1" | please
}

_pls_run(){
  arguments="$*"
  if [[ $arguments == "sudo"* ]]
    then
    [ "$UID" -eq 0 ] || exec sudo "$@"
  else
    eval "$arguments"
  fi
}

if [ "$#" -gt 0 ]
then
  arguments="$*"
  echo "$arguments"
  _pls_run $arguments
else
  last_command=`fc -ln -2 -2`
  #trim whitespace, especially \t introduced by fc
  last_command="${last_command#"${last_command%%[![:space:]]*}"}"
  #echo "Last Command: ${last_command}"
  IFSold=$IFS
  IFS=$'\n'
  # get command suggestions
  suggestion=$(_pls_suggest $last_command)
  echo "Do you want to run the following command?"
  echo "${suggestion}"
  read -p $'Y/N\n' -n 1
  echo ""
  if [[ -z "$REPLY" ]] || [[ $REPLY =~ ^[Yy]$ ]]
  then
    _pls_run $suggestion
  fi
  IFS=$IFSold
  exit 1
fi
