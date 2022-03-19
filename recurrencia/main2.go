//Generador de  lista de nùmero enviado por canal
package main
import 
(
   "fmt"
  // "time"
)

func Generador(c chan<- int){

     for i := 1; i<= 10; i++{

   	c <- i
      }
      //c <- 1555555
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
	fmt.Println("---El canal de entrada es el siguente ")
	//go Print(generador)
	//time.Sleep(10 * time.Second)
    go Generador(generador)
	//time.Sleep(10 * time.Second)
	go Doble(generador,dobles)
	//time.Sleep(10 * time.Second)
	 Print(dobles)
    //time.Sleep(10 * time.Second)
       
	
	

	


}
