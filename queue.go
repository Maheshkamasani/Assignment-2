package main
import "fmt"

func enqueue(queue[] int, element int) []int {
  queue = append(queue, element); 
  fmt.Println("Enqueued:", element);
  return queue
}

func dequeue(queue[] int) ([]int) {
  element := queue[0]; 
  fmt.Println("Dequeued:", element)
  return queue[1:]; 
}

func main() {
  var queue[] int; 

  queue = enqueue(queue, 30);
  queue = enqueue(queue, 40);
  queue = enqueue(queue, 50);

  fmt.Println("Queue:", queue);

  queue = dequeue(queue);
  fmt.Println("Queue:", queue);

  queue = enqueue(queue, 40);
  fmt.Println("Queue:", queue);
}