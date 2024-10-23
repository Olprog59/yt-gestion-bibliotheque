# **Exercice : Gestion d'une Bibliothèque**

## **Objectif :**

Créer un programme en Go pour gérer une petite bibliothèque de livres.

**Instructions :**

1. **Définir une struct `Livre`** avec les champs suivants :
    - `Titre` (string)
    - `Auteur` (string)
    - `Pages` (int)
    - `Disponible` (bool)
2. **Créer une slice de `Livre`** pour stocker plusieurs livres dans la bibliothèque.
3. **Initialiser la bibliothèque** avec au moins 5 livres.
4. **Écrire une fonction `AfficherLivres`** qui parcourt la slice et affiche les informations de chaque livre.
5. **Écrire une fonction `EmprunterLivre`** qui :
    - Prend en entrée le titre du livre à emprunter.
    - Vérifie si le livre est disponible.
    - Met à jour le statut `Disponible` à `false` si le livre est emprunté avec succès.
    - Affiche un message approprié si le livre n'est pas disponible.
6. **Écrire une fonction `RetournerLivre`** qui :
    - Prend en entrée le titre du livre à retourner.
    - Met à jour le statut `Disponible` à `true`.
    - Affiche un message confirmant le retour.
7. **Dans la fonction `main`** :
    - Afficher un **menu** avec les options :
        - Afficher tous les livres
        - Emprunter un livre
        - Retourner un livre
        - Quitter le programme
    - Utiliser une boucle pour permettre à l'utilisateur de choisir une option jusqu'à ce qu'il décide de quitter.
    - Utiliser des **conditions** pour exécuter les actions en fonction du choix de l'utilisateur.
8. **Gestion des erreurs** :
    - Gérer les cas où le livre demandé n'existe pas.
    - Assurer que l'utilisateur ne peut pas emprunter un livre déjà emprunté ou retourner un livre non emprunté.

**Bonus :**

- Utiliser une **map** pour associer les titres des livres à leurs indices dans la slice, afin d'optimiser la recherche.
- Implémenter des **fonctions supplémentaires** pour ajouter ou supprimer des livres de la bibliothèque.
- Utiliser des **array** si nécessaire pour des fonctionnalités spécifiques.


```go
bibliotheque := []Livre{
  {Titre: "Le Petit Prince", Auteur: "Antoine de Saint-Exupéry", Pages: 96, Disponible: true},
  {Titre: "1984", Auteur: "George Orwell", Pages: 328, Disponible: true},
  {Titre: "Les Misérables", Auteur: "Victor Hugo", Pages: 1232, Disponible: true},
  {Titre: "L'Étranger", Auteur: "Albert Camus", Pages: 123, Disponible: true},
  {Titre: "Le Seigneur des Anneaux", Auteur: "J.R.R. Tolkien", Pages: 1216, Disponible: true},
}
```
