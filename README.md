# Projeto em GOlang

Prof Ricardo Mendes | Turma CC5M | 2022/1 |

![cd snippetbox](https://user-images.githubusercontent.com/62394959/164979244-a6a852cb-a867-427e-8a2b-4af52c2625d0.jpg) BY Ricardo Mendes

ps.: só funciona o link que executar  no replit antes, se não fica carredando infinitamente
BY Ricardo Mendes

> PFCC5M-WEB [Snippetbox](https://pfcc5m-web.suellenmayuko.repl.co/)

~~~Bash
cmd executar antes
$ cd snippetbox
$ go run main.go
~~~

~~~Bash
cmd executar depois
$ cd snippetbox
$ go run cmd/web/*
~~~

~~~Bash
cmd line flags
$ go run cmd/web/* -addr=":80"
~~~

~~~ 
Testes
•Certas:

http://localhost:4000/snippet?id=123

•Erros:

http://localhost:4000/snippet

http://localhost:4000/snippet?id=-1

http://localhost:4000/snippet?id=foo
~~~
