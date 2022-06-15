package main

import "fmt"

func main(){
	fmt.Println("================================================")
	fmt.Println("|| PROGRAM PENENTUAN RUTE PENYEBARAN DOKUMEN ||")
	fmt.Println("||    ANTAR KECAMATAN DI KABUPATEN BANGLI    ||")
	fmt.Println("================================================")

	// Initialize Data
	//code node
	// A = kantor camat bangli, 
	// B = Kantor camat tembuku,
	// C = Kantor camat kintamani, 
	// D = Kantor camat susut
	var dataJarak = [4][4]float32 {
		//A-B-C-D
		{ 0, 8.0, 28.2, 6.6},	//A
		{ 8.0, 0, 30.1, 9.1},	//B
		{ 28.2, 30.1, 0, 22.7},	//C
		{ 6.6, 9.1, 22.7, 0},	//D
	}
	var isVisited [4]int
	var dataKantor = [4]string{
		"Bangli",
		"Tembuku",
		"Kintamani",
		"Susut",
	}
	fmt.Println("== Daftar Kode Kecamatan ==")
	for i, data := range dataKantor {
		println("|-", i+1,"=",data)
	}
	fmt.Println("===========================")
	fmt.Print("Masukan kode kecamatan untuk menentukan titik awal : ")
	var start int;
	fmt.Scanf("%d", &start)
	fmt.Println("========================================================")
	fmt.Println("Rute penyebaran dokumen dengan titik awal kantor camat ",dataKantor[start-1],": ")
	process(start-1, isVisited, dataJarak, dataKantor)
}

func process(start int, isVisited [4]int, dataJarak [4][4]float32, dataKantor [4]string){
	var(
		min float32
		saveMin int
	) 
	min = 999.00
	var repeat  = false;
	for i := 0; i < len(dataJarak); i++{
	  if (dataJarak[start][i] == 0){
		continue;
	  }else{
		if (isVisited[i] == 0){
		  if (dataJarak[start][i] < min){
			min = dataJarak[start][i];
			saveMin = i;
			repeat = true;
		  }
		}
	  }
	}
	if (repeat){
	  fmt.Print(" -> Kantor Camat ", dataKantor[saveMin]);
	  isVisited[saveMin] = 1;
	  process(start, isVisited, dataJarak, dataKantor);
	}
  
}