# Анализатор размера диска (Go)

[![hexlet-check](https://github.com/hazhdev/go-from-scratch-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/hazhdev/go-from-scratch-project-242/actions)

Консольная утилита, которая считает размер файла или директории с флагами рекурсивного обхода, человекочитаемого формата и учёта скрытых файлов.

Учебный проект Хекслета: https://ru.hexlet.io/programs/go

## Стек

- Go

## Установка

```bash
git clone https://github.com/hazhdev/go-from-scratch-project-242.git
cd go-from-scratch-project-242
make build
```

## Использование

```bash
./bin/hexlet-path-size -h
./bin/hexlet-path-size testdata
./bin/hexlet-path-size -a testdata
./bin/hexlet-path-size -H -r -a .
```

---

<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml` — не удаляйте и не переименовывайте ни его, ни репозиторий.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования: авторские программы обучения с практикой, поддержкой наставников и реальными проектами, которые остаются в резюме. Этот репозиторий — один из таких проектов.
