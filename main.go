package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Store struct {
	ID       int
	Name     string
	Branches []*Branch
}

type Branch struct {
	ID          int
	Name        string
	PhoneNumber []string
	Addres      []*Address
	Vacancies   []*Vacancy
}

type Address struct {
	ID         int
	City       string
	StreetName string
}

type Vacancy struct {
	ID     int
	Name   string
	Salary float64
}

type Respons struct {
	Stores []*Store
}

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "ismoiljon12", "12", "storedb")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to connect:", err)
	}

	defer db.Close()

	VacanciesInfo(db)

	_ = []Store{
		{
			Name: "Korzinka",
			Branches: []*Branch{
				{
					Name: "Mirzo Ulug'bek",
					PhoneNumber: []string{
						"+99893 587 47 54",
						"+99871 457 58 58",
					},
					Addres: []*Address{
						{
							City:       "Toshkent, Novza",
							StreetName: "Novza street",
						},
					},
					Vacancies: []*Vacancy{
						{
							Name:   "Manager",
							Salary: 4500000,
						},
						{
							Name:   "Cleaner",
							Salary: 3000000,
						},
					},
				},
				{
					Name: "Kohinur",
					PhoneNumber: []string{
						"+99894 558 45 45",
						"+99871 875 98 89",
					},
					Addres: []*Address{
						{
							City:       "Toshkent,Beruniy",
							StreetName: "Berunit street",
						},
					},
					Vacancies: []*Vacancy{
						{
							Name:   "Cashier",
							Salary: 5000000,
						},
						{
							Name:   "Manager",
							Salary: 4000000,
						},
					},
				},
			},
		},
		{
			Name: "Havas",
			Branches: []*Branch{
				{
					Name: "Beshagach",
					PhoneNumber: []string{
						"+99871 205 55 55",
						"+99894 548 87 87",
					},
					Addres: []*Address{
						{
							City:       "Toshkent,Yunusobos",
							StreetName: "Qodiriy street",
						},
					},
					Vacancies: []*Vacancy{
						{
							Name:   "Guard",
							Salary: 3500000,
						},
						{
							Name:   "Cashier",
							Salary: 4000000,
						},
					},
				},
				{
					Name: "Humo Arena",
					PhoneNumber: []string{
						"+99893 547 56 56",
						"+998971 584 78 78",
					},
					Addres: []*Address{
						{
							City:       "Toshkent,Chilonzor",
							StreetName: "Shuhrat street",
						},
					},
					Vacancies: []*Vacancy{
						{
							Name:   "Cashier",
							Salary: 4500000,
						},
						{
							Name:   "Manager",
							Salary: 5000000,
						},
					},
				},
			},
		},
	}

	// tx, err := db.Begin()
	// if err != nil {
	// 	tx.Rollback()
	// 	fmt.Println("Failed to begin:", err)
	// 	return
	// }

	// for _, store := range stores {
	// 	var storeId int
	// 	err := tx.QueryRow("INSERT INTO stores(name) VALUES($1) RETURNING id", store.Name).Scan(&storeId)
	// 	if err != nil {
	// 		// tx.Rollback()
	// 		fmt.Println("Failed to insert stores:", err)
	// 	}

	// 	for _, branch := range store.Branches {
	// 		var branchID int
	// 		err := tx.QueryRow("INSERT INTO branches(name,phone_numbers,store_id) VALUES ($1,$2,$3) RETURNING id", branch.Name, pq.Array(branch.PhoneNumber), storeId).Scan(&branchID)
	// 		if err != nil {
	// 			tx.Rollback()
	// 			fmt.Println("Failed to INSERT branches:", err)
	// 			return
	// 		}

	// 		for _, address := range branch.Addres {
	// 			_, err = tx.Exec("INSERT INTO addresses(city,steet_name,branch_id) VALUES($1,$2,$3)", address.City, address.StreetName, branchID)
	// 			if err != nil {
	// 				tx.Rollback()
	// 				fmt.Println("Failed to INSERT addresses:", err)
	// 				return
	// 			}
	// 		}
	// 		for _, vacancy := range branch.Vacancies {
	// 			var vacancyID int
	// 			err := tx.QueryRow("INSERT INTO vacancies(name,salary) VALUES($1,$2) RETURNING id", vacancy.Name, vacancy.Salary).Scan(&vacancyID)
	// 			if err != nil {
	// 				tx.Rollback()
	// 				fmt.Println("Failed to INSERT vacancies", err)
	// 				return
	// 			}

	// 			_, err = tx.Exec("INSERT INTO branches_vacancies(branch_id,vacancy_id) VALUES ($1,$2)", branchID, vacancyID)
	// 			if err != nil {
	// 				tx.Rollback()
	// 				fmt.Println("Failed to INSERT branches_vacancies", err)
	// 				return
	// 			}
	// 		}
	// 	}
	// }
	// err = tx.Commit()
	// if err != nil {
	// 	tx.Rollback()
	// 	fmt.Println("Failed to commit: ", err)
	// 	return
	// }

	// 	resp := Respons{}

	// 	// Store getAll
	// 	storeRows, err := db.Query("SELECT * FROM stores")
	// 	if err != nil {
	// 		fmt.Println("Failed to SELECT store:", err)
	// 		return
	// 	}

	// 	for storeRows.Next() {
	// 		store := Store{}
	// 		err := storeRows.Scan(
	// 			&store.ID,
	// 			&store.Name,
	// 		)
	// 		if err != nil {
	// 			fmt.Println("Failed to Scan store", err)
	// 			return
	// 		}
	// 		// Branch Getall
	// 		branchRows, err := db.Query("SELECT id,name,phone_numbers  FROM branches WHERE store_id=$1", store.ID)
	// 		if err != nil {
	// 			fmt.Println("Failed to Select branches", err)
	// 			return
	// 		}

	// 		for branchRows.Next() {
	// 			branch := Branch{}
	// 			err := branchRows.Scan(
	// 				&branch.ID,
	// 				&branch.Name,
	// 				pq.Array(&branch.PhoneNumber),
	// 			)
	// 			if err != nil {
	// 				fmt.Println("Failed to Scan branch", err)
	// 				return
	// 			}

	// 			vacancyRows, err := db.Query(`SELECT v.id,v.name,v.salary FROM vacancies v JOIN branches_vacancies br ON v.id = br.vacancy_id JOIN
	// 			branches b ON b.id = br.branch_id WHERE b.id = $1`, branch.ID)
	// 			if err != nil {
	// 				fmt.Println("Failed to select vacancies", err)
	// 				return
	// 			}

	// 			for vacancyRows.Next() {
	// 				vacancy := Vacancy{}
	// 				err := vacancyRows.Scan(
	// 					&vacancy.ID,
	// 					&vacancy.Name,
	// 					&vacancy.Salary,
	// 				)
	// 				if err != nil {
	// 					fmt.Println("Failed to Scan vacancies", err)
	// 					return
	// 				}

	// 				branch.Vacancies = append(branch.Vacancies, &vacancy)
	// 			}

	// 			addressRows, err := db.Query("SELECT id,city,street_name FROM addresses WHERE branch_id = $1", branch.ID)
	// 			if err != nil {
	// 				fmt.Println("Failed to select addreses", err)
	// 				return
	// 			}
	// 			for addressRows.Next() {
	// 				addres := Address{}
	// 				err := addressRows.Scan(
	// 					&addres.ID,
	// 					&addres.City,
	// 					&addres.StreetName,
	// 				)
	// 				if err != nil {
	// 					fmt.Println("Failed to Scan addresses", err)
	// 					return
	// 				}

	// 				branch.Addres = &addres
	// 			}

	// 			store.Branches = append(store.Branches, &branch)
	// 		}
	// 		resp.Stores = append(resp.Stores, &store)

	// 		for _, store := range resp.Stores {
	// 			fmt.Println(store)

	// 			for _, branch := range store.Branches {
	// 				fmt.Println(branch)

	// 				for _, vacancy := range branch.Vacancies {
	// 					fmt.Println(vacancy)
	// 					fmt.Println(branch.Addres)
	// 				}

	// 			}

	// 		}
	// 	}

}

type VacancyResp struct {
	ID       int
	Name     string
	Salary   float64
	Branches []*VacancyBranch
}

type VacancyBranch struct {
	ID           int
	Name         string
	PhoneNumbers []string
	Address      *Address
	Store        *Store
}

type VacanciesResponse struct {
	Vacancies []*VacancyResp
}

func VacanciesInfo(db *sql.DB) {
	vacResponses := VacanciesResponse{}
	vacancyRows, err := db.Query("SELECT id,name,salary FROM vacancies")
	if err != nil {
		log.Println("Failed to select vacancies", err)
	}
	for vacancyRows.Next() {
		vacancy := Vacancy{}
		err := vacancyRows.Scan(
			&vacancy.ID,
			&vacancy.Name,
			&vacancy.Salary,
		)
		if err != nil {
			log.Println("Failed to scan vacancy")
		}
		vacResp := VacancyResp{}

		vacResp.ID = vacancy.ID
		vacResp.Name = vacancy.Name
		vacResp.Salary = vacancy.Salary
		branchRows, err := db.Query("SELECT b.id,b.name,b.phone_numbers FROM branches b JOIN branches_vacancies br ON b.id = br.branch_id JOIN vacancies v ON v.id = br.vacancy_id WHERE v.id = $1", vacancy.ID)
		if err != nil {
			log.Println("Failed to select branches: ", err)
		}
		for branchRows.Next() {
			branch := Branch{}
			err := branchRows.Scan(
				&branch.ID,
				&branch.Name,
				pq.Array(&branch.PhoneNumber),
			)
			if err != nil {
				log.Println("Failed to scan branches: ", err)
			}
			vacBranch := VacancyBranch{}

			vacBranch.ID = branch.ID
			vacBranch.Name = branch.Name
			vacBranch.PhoneNumbers = branch.PhoneNumber

			addresRows, err := db.Query("SELECT id,city,street_name FROM addresses")
			if err != nil {
				log.Println("Failed to select addresses: ", err)
			}
			for addresRows.Next() {
				addres := Address{}
				err := addresRows.Scan(
					&addres.ID,
					&addres.City,
					&addres.StreetName,
				)
				if err != nil {
					log.Println("Failed to scan addresses: ", err)
				}
				vacBranch.Address = &addres
			}

			storeRows, err := db.Query("SELECT s.id, s.name FROM stores s JOIN branches b ON s.id = b.store_id WHERE b.id = $1", branch.ID)
			if err != nil {
				log.Println("Failed to select store: ", err)
			}

			for storeRows.Next() {
				store := Store{}
				err := storeRows.Scan(
					&store.ID,
					&store.Name,
				)
				if err != nil {
					log.Println("Failed to scan store: ", err)
				}
				vacBranch.Store = &store

			}
			vacResp.Branches = append(vacResp.Branches, &vacBranch)

		}
		vacResponses.Vacancies = append(vacResponses.Vacancies, &vacResp)

	}

	for _, vacancies := range vacResponses.Vacancies {
		fmt.Printf("Id: %d\n", vacancies.ID)
		fmt.Printf("Name: %s\n", vacancies.Name)
		for _, branches := range vacancies.Branches {
			fmt.Printf("\tId: %d\n", branches.ID)
			fmt.Printf("\tName: %s\n", branches.Name)
			fmt.Printf("\tPhoneNumbers: %s\n", branches.PhoneNumbers)
			fmt.Printf("\t\tAddress Id: %d\n", branches.Address.ID)
			fmt.Printf("\t\tAddress City: %s\n", branches.Address.City)
			fmt.Printf("\t\tAddress StreetName: %s\n", branches.Address.StreetName)
			fmt.Printf("\t\t\tStore Id: %d\n", branches.Store.ID)
			fmt.Printf("\t\t\tStore Name: %s\n", branches.Store.Name)
		}
	}

}
