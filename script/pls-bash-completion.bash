_pls_complete()
{
    local current_word last_command suggestion

    current_word="${COMP_WORDS[COMP_CWORD]}"

    last_command=`fc -ln -0`
    last_command="${last_command#"${last_command%%[![:space:]]*}"}"
    echo "Last Command: $last_command" >> "bashruns.txt"
    local IFS=$'\n'
    suggestion=`please "$last_command"`
    echo "$suggestion" >> "bashruns.txt"
    COMPREPLY=( $(compgen -W "${suggestion}" -- $current_word ) );
    return 0
}

complete -F _pls_complete pls