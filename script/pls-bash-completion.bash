_pls_suggest () {
    echo "$1" | please
}

_pls_complete()
{
    local current_word last_command suggestion

    current_word="${COMP_WORDS[COMP_CWORD]}"

    last_command=`fc -ln -0`
    #trim whitespace, especially \t introduced by fc
    last_command="${last_command#"${last_command%%[![:space:]]*}"}"

    local IFS=$'\n'
    # get command suggestions
    suggestion=$(_pls_suggest $last_command)
    # escape single and double quotes
    suggestion=${suggestion//\"/\\\"}
    suggestion=${suggestion//\'/\\\'}
    #filter list of suggestions using the current word
    suggestions=($(compgen -W "${suggestion[*]}" -- "$current_word"))

    if [ ${#suggestions[*]} -eq 0 ]; then
        COMPREPLY=()
    else
        COMPREPLY=($(printf '%s\n' "${suggestions[@]}"))
    return 0
    fi
}

complete -F _pls_complete pls