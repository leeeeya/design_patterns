# Фабричный метод (Factory Method)

Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов, позволяя изменять их тип.

## Применимость

Паттерн Фабричный метод предлагает создавать объекты не напрямую, а через вызов особого фабричного метода.

Фабричный метод можно переопределить в структуре, чтобы изменить тип создаваемого продукта.

Чтобы эта система работала, все возвращаемые объекты должны иметь общий интерфейс. "Фабрика" сможет производить объекты различных типов, следующих одному и тому же интерфейсу.

## Плюсы

- Выделяет код производства продуктов в одно место, упрощая поддержку кода.
- Упрощает добавление новых продуктов в программу.
- Реализует принцип open-closed.

## Минусы

- Может привести к созданию больших параллельных иерархий классов (структур), так как для каждого типа продукта надо создать свою фабрику.

## Примеры использования

- Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать код.

    Фабричный метод отделяет код производства продуктов от остального кода, который эти продукты использует.
    Благодаря этому, код производства можно расширять, не трогая основной. Так, чтобы добавить поддержку нового продукта, вам нужно создать новую фабрику, возвращая оттуда экземпляр нового продукта.

- Когда нужно дать возможность пользователям расширять части какого-либо фреймворка или библиотеки.

    Пользователи могут расширять структуры фреймворка через наследование.
    Для этого нужно дать пользователям возможность расширять не только желаемые компоненты, но и структуры и методы, которые создают эти компоненты. А для этого создающие фабрики должны иметь конкретные создающие методы, которые можно определить.

- Когда нужно экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых.

    Такая проблема обычно возникает при работе с тяжёлыми ресурсоёмкими объектами, такими, как подключение к базе данных, файловой системе и т. д.
    Самым удобным способом для создания таких объектов был бы конструктор, но конструктор всегда создаёт новые объекты, он не может вернуть существующий экземпляр.
    Значит, нужен другой метод, который бы отдавал как существующие, так и новые объекты. Им и станет фабричный метод.