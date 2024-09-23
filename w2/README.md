---
marp: true
---

# Sentencias de Control y Repeticiones en Go y Java

---

## Sentencias de control

---

### 1. if y else

**Java**

```java
int number = 10;
if (number > 0) {
    System.out.println("Positive");
} else {
    System.out.println("Negative or zero");
}
```

**Go**

```go
number := 10
if number > 0 {
    fmt.Println("Positive")
} else {
    fmt.Println("Negative or zero")
}
```

---

## 2. switch

**Java**

```java
int day = 3;
switch (day) {
    case 1:
        System.out.println("Monday");
        break;
    case 2:
        System.out.println("Tuesday");
        break;
    default:
        System.out.println("Another day");
}
```

---

**Go**

```go
day := 3
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
default:
    fmt.Println("Another day")
}
```

---

## 3. For

**Java**

```java
for (int i = 0; i < 5; i++) {
    System.out.println(i);
}
```

**Go**

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

---

### 4. Ciclo while

**Java**

```java
int i = 0;
while (i < 5) {
    System.out.println(i);
    i++;
}
```

**Go**

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

---

### 5. Ciclo do-while

**Java**

```java
int i = 0;
do {
    System.out.println(i);
    i++;
} while (i < 5);
```

**Go**
No hay un ciclo do-while en Go, pero puede simularse usando un for infinito con una condición break.

```go
i := 0
for {
    fmt.Println(i)
    i++
    if i >= 5 {
        break
    }
}
```
