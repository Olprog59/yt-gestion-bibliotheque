package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Livre struct {
	Titre      string
	Auteur     string
	Pages      int
	Disponible bool
}

func main() {
	bibliotheque := []*Livre{
		{Titre: "Le Petit Prince", Auteur: "Antoine de Saint-Exupéry", Pages: 96, Disponible: true},
		{Titre: "1984", Auteur: "George Orwell", Pages: 328, Disponible: true},
		{Titre: "Les Misérables", Auteur: "Victor Hugo", Pages: 1232, Disponible: true},
		{Titre: "L'Étranger", Auteur: "Albert Camus", Pages: 123, Disponible: true},
		{Titre: "Le Seigneur des Anneaux", Auteur: "J.R.R. Tolkien", Pages: 1216, Disponible: true},
	}

	scanner := bufio.NewScanner(os.Stdin)

iter:
	for {
		fmt.Println("\n--- Bibliothèque ---")
		fmt.Println("1. Afficher tous les livres")
		fmt.Println("2. Emprunter un livre")
		fmt.Println("3. Retourner un livre")
		fmt.Println("4. Quitter le programme")
		fmt.Print("Choisissez une option (1-4): ")

		scanner.Scan()
		choix := scanner.Text()

		switch choix {
		case "1":
			afficherLivres(bibliotheque)
		case "2":
			fmt.Println("Donnez moi le titre recherche")
			scanner.Scan()
			titre := scanner.Text()
			livre := rechercheLivre(bibliotheque, titre)
			emprunterLivre(livre)
		case "4":
			break iter
		default:
			fmt.Println("Je ne connais pas ce choix")
		}
	}

	fmt.Println("a bientot")
}

func afficherLivres(bibi []*Livre) {
	for _, livre := range bibi {
		fmt.Printf("Titre %s | Auteur %s | Pages %d | Disponible %t\n",
			livre.Titre,
			livre.Auteur,
			livre.Pages,
			livre.Disponible,
		)
	}
}

func rechercheLivre(bibi []*Livre, titre string) *Livre {
	for _, livre := range bibi {
		if strings.EqualFold(livre.Titre, titre) {
			return livre
		}
	}
	return nil
}

func emprunterLivre(livre *Livre) {
	if livre.Disponible {
		livre.Disponible = false
		fmt.Println("Bonne lecture")
		return
	} else {
		fmt.Println("Le livre n'est pas disponible")
		return
	}
}

func retournerLivre(livre *Livre) {
	if !livre.Disponible {
		livre.Disponible = true
		fmt.Println("Merci")
		return
	} else {
		fmt.Println("Le livre est disponible")
		return
	}
}
