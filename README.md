
# Flexões Hoje ![Versão interface de linha de comando (CLI)](https://img.shields.io/badge/CLI-636b2f)

Quantas flexões de braço você fez hoje? **Registre sem sair do terminal!**

![Demonstração das funcionalidades](docs/demo-2026-05-02.gif)

## Utilização

```sh
# Veja quantas flexões você executou hoje
flexoeshoje

# Registre determinado número flexões no dia de hoje
flexoeshoje 10

# Subtraia um número de flexões caso você precise
flexoeshoje 5 --subtrair

# Conhece mais sobre a CLI
flexoeshoje --help
```

## Instalação

1. Entre na [página de releases](https://github.com/kauefraga/flexoeshoje-cli/releases/latest)
2. Instale o binário compatível com a sua plataforma (exemplo: `flexoeshoje.exe` para Windows)
3. Pronto! Comece já a registrar quantas vezes você empurrou a Terra para baixo!

###### Instalação pelo terminal

<details>

<summary>Linux</summary>

```sh
curl -OL https://github.com/kauefraga/flexoeshoje-cli/releases/latest/download/flexoeshoje
```

</details>

<details>

<summary>Windows</summary>

```sh
curl -OL https://github.com/kauefraga/flexoeshoje-cli/releases/latest/download/flexoeshoje.exe
```

</details>

<details>

<summary>MacOS</summary>

```sh
curl -OL https://github.com/kauefraga/flexoeshoje-cli/releases/latest/download/flexoeshoje-darwin
```

</details>

Lembre-se de conceder permissão de execução ao binário e mover para um diretório que esteja no `PATH`.

## Desenvolvimento

Ferramenta feita usando Cobra para construção da interface de linha de comando e SQLite como banco de dados local.

Para adição de mais flexões no mesmo dia foi implementada uma estratégia similar a um ["Ledger"](https://en.wikipedia.org/wiki/General_ledger), onde cada vez que o usuário informa novas repetições de flexão o sistema cria um novo registro, ao invés de atualizar apenas um. Com isso, existe um controle robusto das execuções.

O banco de dados SQLite foi escolhido pois oferece praticidade e máxima eficiência já que o banco se encontra na própria máquina, porém futuramente a ferramenta irá se integrar com a API do projeto, [flexoeshoje-api](https://github.com/kauefraga/flexoeshoje-api).

## Licença

Este projeto está sob a licença MIT - Veja a [LICENÇA](LICENSE) para mais informações.
