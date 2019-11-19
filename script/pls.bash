last_command=`fc -ln -2 -2`

echo "Last Command: ${last_command}"
suggestion=`please "${last_command}"`

echo "Sugggestion: "${suggestion}""