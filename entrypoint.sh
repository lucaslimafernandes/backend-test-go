#!/bin/sh

# Função para iniciar e monitorar processos
start_processes() {
    # Inicia o primeiro binário em background
    /bin/main &
    pid1=$!

    # Inicia o segundo binário em background
    /bin/services &
    pid2=$!

    # Aguarda até que qualquer um dos processos termine
    wait -n $pid1 $pid2

    # Captura o código de saída do primeiro processo que terminar
    exit_code=$?

    # Encerra ambos os processos
    kill $pid1 $pid2

    # Retorna o código de saída do processo que terminou primeiro
    return $exit_code
}

# Chama a função para iniciar e monitorar processos
start_processes
