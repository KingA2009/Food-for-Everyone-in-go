package main

import "fmt"


type Eat struct{
	name,like_eat,dislike_eat string
}


var eats []Eat
func main(){
	var p1 Eat
	p1.name = "Sam"
	p1.like_eat = "ice cream"
	p1.dislike_eat = "carrots"	
	eats = append(eats, p1)
	taste(p1.name,p1.dislike_eat)
}

func taste(name string,eat string){
	if eat == eats[0].dislike_eat{
		fmt.Println(name +" eats the " + eat+ " and hates it!")
	}else if eat == eats[0].like_eat{
		fmt.Println(name +" eats the " + eat+ " and loves it!")
	}else{
		fmt.Println(name +" eats the " + eat)
	}
}