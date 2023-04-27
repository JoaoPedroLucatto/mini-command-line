# mini-command-line
Mini application command line
<!---Esses são exemplos. Veja https://shields.io para outras pessoas ou para personalizar este conjunto de escudos. Você pode querer incluir dependências, status do projeto e informações de licença aqui--->

## 💻 Prerequisites

* Have docker installed **[docker:Docker](https://www.docker.com/)**

## ☕ Running

* Clone the project.
* Run the following command to build the image.
```console
docker build . -t <name-image> 
```
* To run the image along with the program. Example:
```console
docker run  <name-image> /line-code ip --host amazon.com.br
```
```console
docker run  <name-image> /line-code server --host amazon.com.br
```

## 😎 Technologies used

![image](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![image](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
