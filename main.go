package main
import (
	//"errors"
	"fmt"
)
/* const passingMark = 50

type address struct{
	City string
	Country string 
}

type Student struct{
	ID string
	Name string
	Age int
	Address address
	Scores map[string]float64
}

func displayStudents(db []Student){
	for _,student := range db{
		fmt.Println(student.ID,student.Name,student.Age,student.Address,student.Scores)
	}
}

func findStudent(students []Student, ID string)(Student ,error){
	for _,student:= range students{
		if  student.ID == ID{
			return student, nil
		}
	}
	return Student{},errors.New("student with ID")

}

func calculateAverage(scores map[string]float64)(float64,error){

		
		if len(scores) == 0 {
			return 0 ,errors.New("no scores recorded")
		}
        total := 0.0
		for _, score := range scores {
			total += score
		}

		return total/float64(len(scores)), nil
	}

	




func hasPassed(average float64)bool{
	if average >= passingMark{
		return true
	}
	return false
}



func displayStudentsPassed(students []Student){
	for _,student := range students{
		average,err := calculateAverage(student.Scores)
		if err != nil{
			fmt.Println("skipping", student.Name, "-",err)
			continue
		}
		if !hasPassed(average){
			continue
		}
		

        fmt.Println(student.Name , average)
	}

}

func validateStudent(students []Student, s Student)error{
	if s.Name == ""{
         return  errors.New("name cannot be empty")
	}
	if s.Age < 0{
		return  errors.New("age cannot be negative")
	}
	if s.ID == "" {
	return errors.New("ID cannot be empty")
}

    for _, existing := range students {
	   if existing.ID == s.ID {
		  return fmt.Errorf("ID %s is already in use", s.ID)
	}
}
	return nil
	
}

func addStudent(students []Student, s Student)([]Student,error){
	
     err := validateStudent(students,s)
	 if err != nil{
		return students, err
	 }
	return  append(students,s), nil


}



func main(){

	students := []Student{
		{
			ID:   "STU_001",
			Name: "Nicholas",
			Age:  24,
			Address: address{
				City:    "Kampala",
				Country: "Uganda",
			},
			Scores: map[string]float64{
				"Mathematics": 85,
				"English":     90,
				"Science":     92,
			},
		},
		{
			ID:   "STU_002",
			Name: "Amara",
			Age:  22,
			Address: address{
				City:    "Nairobi",
				Country: "Kenya",
			},
			Scores: map[string]float64{
				"Mathematics": 78,
				"English":     88,
				"Science":     84,
			},
		},
		{
			ID:   "STU_003",
			Name: "David",
			Age:  26,
			Address: address{
				City:    "Accra",
				Country: "Ghana",
			},
			Scores: map[string]float64{
				"Mathematics": 95,
				"English":     91,
				"Science":     100,
			},
},

	}

fmt.Println("All Students")
displayStudents(students)

fmt.Println("Students who passed")

displayStudentsPassed(students)


fmt.Println("--- Search and Calculate Individual Average ---")
targetID := "STU_001"

findStudents,err := findStudent(students, targetID)

if err != nil{
	fmt.Println("the student wasnt found ",err)

}else{
	fmt.Printf("Found Student:", findStudents.Name)

	avg,err := calculateAverage(findStudents.Scores)

	if err != nil {
       fmt.Println("Error calculating average:", err)
	}else{
		fmt.Printf("Average Score:", avg)
		fmt.Printf("Passed ", hasPassed(avg))
	}

}


fmt.Println("\n--- Adding a New Student ---")
	newStudent := Student{
		ID:   "STU_004",
		Name: "Kofi",
		Age:  23,
		Address: address{
			City:    "Kumasi",
			Country: "Ghana",
		},
		Scores: map[string]float64{
			"Mathematics": 60,
			"English":     70,
		},
	}

studentz,err := addStudent(students,newStudent)


if err != nil{
	fmt.Println("Failed to add a student",err)
}else{
	fmt.Println("Successfully added a sudent and the Total is:", len(studentz))
}






}  */

// pointers



func addTen(num int){
	num = num + 10
}
func addTenPtr(num *int){
	*num = *num + 10
}



// interface


type Notiers interface{
	Notify(message string)
}

type Email struct{
	EmailAdress string 
}

type SMS struct{
	PhoneNumber string 
}


func (e Email)Notify(message string){

		fmt.Printf("Sending Email to:", e.EmailAdress, message)
	}

func (s SMS)Notify(message string){
		fmt.Printf("Sending SMS to:", s.PhoneNumber, message)
		
	}


func main(){
	val := 5 
	addTen(val)
	fmt.Println(val)

	addTenPtr(&val)
	fmt.Println(val)



}
