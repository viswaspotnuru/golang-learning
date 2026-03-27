In the video, Srini explains **Methods** starting at [[3:42:21]](https://www.google.com/search?q=https://www.youtube.com/watch%3Fv%3DtgGNwG_UxFo%26t%3D13341s). In Go, a method is simply a function with a special **receiver** argument. It allows you to "attach" behavior to a struct, making it feel like Object-Oriented Programming (OOP) without the complexity of classes.

-----

### 1\. The Receiver (The "This" or "Self" of Go)

The receiver is the variable defined between the `func` keyword and the method name. It tells Go which type this method belongs to.

  * **Instance:** When you create a variable from a struct (e.g., `v := Vertex{3, 4}`), that variable is an **instance**.
  * **Receiver:** The parameter that allows the method to access the instance's data.

-----

### 2\. Value Receiver vs. Pointer Receiver

This is the most critical distinction in Go backend development, as it directly impacts **memory management** and **data consistency**.

#### **Value Receiver (`func (v Vertex) Abs()`)**

  * **How it works:** Go creates a **complete copy** of the instance and passes it to the method.
  * **Memory:** It uses more memory if the struct is large because of the copying process.
  * **Behavior:** Any changes made to the receiver inside the method **do not affect the original instance**. It is "read-only" in practice.

#### **Pointer Receiver (`func (v *Vertex) Scale(f float64)`)**

  * **How it works:** Go passes the **memory address** (8 bytes) of the instance.
  * **Memory:** Extremely efficient. Regardless of how large the struct is, you only pass a tiny pointer.
  * **Behavior:** Changes made inside the method **modify the original instance**.

-----

### 3\. Comparison Table: Where to use what?

| Feature | Value Receiver (`T`) | Pointer Receiver (`*T`) |
| :--- | :--- | :--- |
| **Modification** | Cannot modify original. | **Can modify original.** |
| **Efficiency** | Slow for large structs (copies data). | **Fast** (only copies the address). |
| **Consistency** | Use for small, immutable data. | **Standard for most API/Backend logic.** |
| **Nil Safety** | Will panic if called on nil. | Can be written to handle `nil` gracefully. |

-----

### 4\. Memory Management: Behind the Scenes

As Srini highlights in the [methods section](https://www.google.com/search?q=https://www.youtube.com/watch%3Fv%3DtgGNwG_UxFo%26t%3D13341s), Go's compiler is smart about how it handles these:

  * **Automatic Referencing:** If you have a value `v` and call a pointer method `v.Scale(5)`, Go automatically converts it to `(&v).Scale(5)`.
  * **Automatic Dereferencing:** If you have a pointer `p` and call a value method `p.Abs()`, Go interprets it as `(*p).Abs()`.
  * **Escape Analysis:** If a method returns a pointer to a local variable, Go’s compiler will move that variable from the **Stack** to the **Heap** so it doesn't disappear when the function ends.

-----

### 5\. Practical Example

```go
type Vertex struct {
    X, Y float64
}

// VALUE RECEIVER: Used for calculation, doesn't change X or Y.
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// POINTER RECEIVER: Used to update the actual coordinates.
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}
```

**Rule of Thumb:** If you are unsure, use a **Pointer Receiver**. It is the industry standard for Go because it avoids unnecessary copying and allows for consistency across your application.

For a deeper look at the real-world codebase examples Srini uses to explain these, check out the [methods chapter at 3:42:21](https://www.google.com/search?q=https://www.youtube.com/watch%3Fv%3DtgGNwG_UxFo%26t%3D13341s).

Would you like to see how these methods interact with **Interfaces**, which Srini explains right after this?