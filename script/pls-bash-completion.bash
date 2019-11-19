_pls_complete()
{
    local current_word last_command suggestion

    current_word="${COMP_WORDS[COMP_CWORD]}"

    last_command=`fc -ln -0`
    local IFS=$'\n'
    suggestion=`please "${last_command}"`

    COMPREPLY=( $(compgen -W "${suggestion}" -- $current_word ) );
    return 0
}

complete -F _pls_complete pls