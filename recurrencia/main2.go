//Generador de  lista de nùmero enviado por canal
package main
import 
(
   "fmt"
   "time"
)

func Generador(c chan<- int){

     for i := 1; i<= 10; i++{

   	c <- i
      }
      c <- 1555555
	close(c)

}

func Doble(in <-chan int,out chan<- int){
		for value:= range in{
 			out <-2*value
			}
		close(out)
}
func Print( c <-chan int){

		for value :=  range c{
			fmt.Println(value)
			}
}

func  main(){
   
        fmt.Println("---Programa de uso de canales ")
  	generador := make(chan int)
  	dobles	    := make(chan int)
	fmt.Println("---Elcanal de entrada es elsiguente ")
	go Print(generador)
        go Generador(generador)
	go Doble(generador,dobles)
	go Print(dobles)
        time.Sleep(10 * time.Second)
       
	
	

	


}
