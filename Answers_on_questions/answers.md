# Ответы на вопросы

1. Какой самый эффективный способ конкатенации строк?
- Cамый эффективный способ strings.Builder в библиотеке "strings"
 
```go
    var str strings.Builder
    str.WriteString("Welcome")
    str.WriteString(" to") // Welcome to
```
 Остальные способы:
 - Оператор +
 - fmt.Sprintf
 - strings.Join
 - bytes.Buffer


2. Что такое интерфейсы, как они применяются в Go?

- Интерфейс (interface) - это абстрактный тип данных, который определяет набор методов (функций),
которые конкретный тип должен реализовать. 

- Интерфейсы являются наборами методов: Они определяют список методов, которые 
конкретный тип должен реализовать. Интерфейс не содержит полей или данных, 
только сигнатуры методов.

- Реализация интерфейса: Любой тип, который имеет все методы, определенные в интерфейсе,
автоматически считается реализующим этот интерфейс. В Go нет неявного объявления реализации
интерфейса, и типы не должны явно объявлять, что они реализуют интерфейс.

- Полиморфизм: Использование интерфейсов позволяет достичь полиморфизма в Go. Это означает, 
что вы можете создавать функции и методы, которые работают с объектами разных типов,
реализующих один и тот же интерфейс.

3. Чем отличаются RWMutex от Mutex?
  - RWMutex и Mutex - это два блокировщика, используемы для обеспечение безопасности при конкурентным доступе к данным.
  - Mutex - это простая блокировка, которая позволяет только одной горутине войти в защищенный участок кода одновременно (подходит для эксклюзивному изменению данных)
  - RWMutex - это мьютекс с чтением/записью, которая различает операции чтения и записи к данным. Её особенность в том, что горутины могут получать доступ к данным параллельно(без блокировки для чтения), но только одна может горутина может получить доступ к данным на запись в данный момент времени (с блокировкой для записи)

4. Чем отличаются буферизированные и не буферизированные каналы?
  - Не буферизированные каналы не имеют буфера (0) (синхронные). Это значит, что отправка в не буф. канал блокирует отправителя, пока данные не будут приняты получателем. Получатель также блокируется, пока не будет получено сообщение.

  - Буфе. каналы асинхронны. Позволяют отправителю отправлять данные без блокировки, пока данные буфера не будет заполнится. При этом получатель может забирать данные из канала без блокировки, пока канал не станет пустым


5. Какой размер у структуры struct{}{}?

- В go структура struct{}{} не имеет полей и поэтому не занимает памяти.
  Размер экземпляра структуры struct{} в памяти равен нулю


6. Есть ли в Go перегрузка методов или операторов?


7. В какой последовательности будут выведены элементы map[int]int? Пример: m[0]=1, m[1]=124, m[2]=281
 - Порядок вывода элементов map[int]int не определен и может быть различным при каждом запуске программы.


8. В чем разница make и new?

  - make используется для создания slice,map и chan. Синтаксис make(type, len, cap) для slice и chan, а make(type, cap) - для map. 
  - new используется для выделения памяти для нового объекта указанного типа. Синтаксис new(Type), возвращает указатель на новый объект, инициализированный нулевым значеним для своего типа.


9. Сколько существует способов задать переменную типа slice или map?


10. Что выведет данная программа и почему?


func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}

11. Что выведет данная программа и почему?


func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}

12. Что выведет данная программа и почему?


func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}


13. Что выведет данная программа и почему?


func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}


14. Что выведет данная программа и почему?


func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
