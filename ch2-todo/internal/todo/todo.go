package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

var (
	outstanding = false
	verbosity   = false
)

// Outstanding sets the CLI to only list outstandingItems items
func (l *List) Outstanding(outstandingItems bool) {
	outstanding = outstandingItems
}

// Verbose sets the CLI to verbose mode
func (l *List) Verbosity(verbose bool) {
	verbosity = verbose
}

// Add creates a new ToDo item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Compete marks a ToDo item as completed by setting
// Done = true and CompletedAt to the current time.
func (l *List) Complete(i int) error {
	ls := *l
	if i < 1 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// Adjusting index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete method deletes a ToDo item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i < 1 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	*l = append(ls[:i-1], ls[i:]...)

	return nil

}

// Save method encodes the list as JSON and saves it
// using the provided file name
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, js, 0644)
}

// Get method opens the provided filename, decodes
// the JSON data and parses it into a List
func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := " "

		if t.Done {
			if outstanding {
				continue
			}

			prefix = "X"
		}

		formatted += fmt.Sprintf("%s %d: %s", prefix, k+1, t.Task)

		if verbosity {
			formatted += fmt.Sprintf("%s %s", t.CreatedAt, t.CompletedAt)
		}

		formatted += "\n"
	}

	return formatted
}
