package main

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetNotes() []Note {

	table := []Note{
		{1, "To Do", "- Cleaning"},
		{2, "Iternary", "- Hagia Sophia"},
		{3, "Exercise", "- Pushup 25x"},
	}

	return table

}
