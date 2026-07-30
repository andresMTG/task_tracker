# task_tracker

Um gerenciador simples de tarefas via linha de comando.

## Requisitos
- Go 1.24+ instalado

## Como executar
No diretório do projeto, você pode rodar o app diretamente com:

```bash
go run ./cmd add "Comprar pão"
```

Ou criar um binário e executá-lo:

```bash
go build -o task-tracker.exe ./cmd
./task-tracker.exe add "Comprar pão"
```

> No Windows, o comando acima pode ser executado com `./task-tracker.exe` em PowerShell.

## Comandos disponíveis

- `add <descrição>`: cria uma nova tarefa
- `update <id> <nova descrição>`: atualiza a descrição de uma tarefa
- `delete <id>`: remove uma tarefa
- `mark-in-progress <id>`: marca uma tarefa como em andamento
- `mark-in-done <id>`: marca uma tarefa como concluída
- `list`: lista todas as tarefas
- `list <todo|in-progress|done>`: lista tarefas filtradas por status
- `help`: mostra a ajuda do aplicativo

## Exemplo de uso

```bash
go run ./cmd add "Estudar Go"
go run ./cmd list
go run ./cmd mark-in-progress 1
go run ./cmd mark-in-done 1
```

## Armazenamento
As tarefas são salvas em um arquivo chamado `listTask.json` na raiz do projeto. Se ele não existir, o aplicativo cria automaticamente.

## Motivador
https://roadmap.sh/projects/task-tracker

