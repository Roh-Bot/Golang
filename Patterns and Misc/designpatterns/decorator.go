package designpatterns

import "log"

type DB interface {
	Store(string) error
	Delete(string) error
}

type Store struct {
	UserName string
	Mobile   int
}

type Read struct {
	UserName string
	Password string
}

func (s *Store) Store(string) error {
	log.Print("Storing username: ", s.UserName)
	log.Print("Storing mobile: ", s.Mobile)
	return nil
}

func (s *Store) Delete(string) error {
	log.Print("Deleting username: ", s.UserName)
	log.Print("Deleting mobile: ", s.Mobile)
	return nil
}

func (s *Read) Store(string) error {
	log.Print("Storing username: ", s.UserName)
	log.Print("Storing mobile: ", s.Password)
	return nil
}

func (s *Read) Delete(string) error {
	log.Print("Deleting username: ", s.UserName)
	log.Print("Deleting mobile: ", s.Password)
	return nil
}

func myExecuteFunc(db DB) ExecuteFunc {
	return func(s string) {
		log.Println("My Execute Func: ", s)
		db.Store(s)
		db.Delete(s)
	}
}

type ExecuteFunc func(string)

func Execute(fn ExecuteFunc) {
	fn("Executed function type...")
}

func DecoratorExecute() {
	log.Print("Executing Decorator Pattern...")
	Execute(myExecuteFunc(&Store{
		UserName: "Dhebug Bhai",
		Mobile:   1234567890,
	}))
	Execute(myExecuteFunc(&Read{
		UserName: "Dhebug Bhai",
		Password: "1234567890",
	}))
	log.Println("Decorator Pattern Executed...")
}
