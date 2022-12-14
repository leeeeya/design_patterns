# Цепочка обязанностей (Chain of Responsibility)

Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков. Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

## Применимость

Как и многие другие поведенческие паттерны, Цепочка обязанностей базируется на том, чтобы превратить отдельные поведения в объекты. В случае данного паттерна каждая проверка станет отдельным типом с единственным методом выполнения. Данные запроса, над которым происходит проверка, будут передаваться в метод как аргументы.

Паттерн предлагает связать объекты обработчиков в одну цепь. Каждый из них будет иметь ссылку на следующий обработчик в цепи. Таким образом, при получении запроса обработчик сможет не только сам что-то с ним сделать, но и передать обработку следующему объекту в цепочке.

При передаче запросов в первый обработчик цепочки, все объекты в цепи всегда смогут его обработать. При этом длина цепочки не имеет никакого значения.

Обработчик не обязательно должен передавать запрос дальше.

Очень важно, чтобы все объекты цепочки имели общий интерфейс. Обычно каждому конкретному обработчику достаточно знать только то, что следующий объект в цепи имеет метод для выполнения. Благодаря этому связи между объектами цепочки будут более гибкими. Кроме того, можно формировать цепочки на лету из разнообразных объектов, не привязываясь к конкретным типам.

## Плюсы

- Уменьшает зависимость между клиентом и обработчиками.
- Реализует принцип single responsibility.
- Реализует принцип open-closed.

## Минусы

- Запрос может остаться никем не обработанным.

## Примеры использования

- Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.

    С помощью Цепочки обязанностей вы можете связать потенциальных обработчиков в одну цепь и при получении запроса поочерёдно спрашивать каждого из них, не хочет ли он обработать запрос.

- Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.

    Цепочка обязанностей позволяет запускать обработчиков последовательно один за другим в том порядке, в котором они находятся в цепочке.

- Когда набор объектов, способных обработать запрос, должен задаваться динамически.

    В любой момент вы можете вмешаться в существующую цепочку и переназначить связи так, чтобы убрать или добавить новое звено.