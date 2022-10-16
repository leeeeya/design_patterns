# Стратегия (Strategy)

Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый из них в собственную структуру, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

## Применимость

Паттерн Стратегия предлагает определить семейство схожих алгоритмов, которые часто изменяются или расширяются, и вынести их в собственные типы, называемые стратегиями.

Вместо того, чтобы изначальный тип сам выполнял тот или иной алгоритм, он будет играть роль контекста, ссылаясь на одну из стратегий и делегируя ей выполнение работы. Чтобы сменить алгоритм, будет достаточно подставить в контекст другой объект-стратегию.

Важно, чтобы все стратегии имели общий интерфейс. Используя этот интерфейс, контекст будет независимым от конкретных объектов стратегий. С другой стороны, можно изменять и добавлять новые виды алгоритмов, не трогая код контекста.

## Плюсы

- Замена алгоритмов на лету.
- Изолирует код и данные алгоритмов от остальных типов.
- Уход от наследования к делегированию.
- Реализует принцип open-closed.

## Минусы

- Усложняет программу за счёт дополнительных объектов.
- Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

## Примеры использования

- Когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта.

    Стратегия позволяет варьировать поведение объекта во время выполнения программы, подставляя в него различные объекты-поведения (например, отличающиеся балансом скорости и потребления ресурсов).

- Когда есть множество похожих объектов, отличающихся только некоторым поведением.

    Стратегия позволяет вынести отличающееся поведение в отдельную иерархию объектов, а затем свести первоначальные объекты к одному, сделав его поведение настраиваемым.

- Когда нужно скрыть детали реализации алгоритмов для других объектов.

    Стратегия позволяет изолировать код, данные и зависимости алгоритмов от других объектов, скрыв эти детали внутри объектов-стратегий.

- Когда различные вариации алгоритмов реализованы в виде большого условного оператора. Каждая ветка такого оператора представляет собой вариацию алгоритма.

    Стратегия помещает каждый кейс такого оператора в отдельный объект-стратегию. Затем контекст получает определённый объект-стратегию от клиента и делегирует ему работу. Если вдруг понадобится сменить алгоритм, в контекст можно подать другую стратегию.