package observer

type Observable interface {
	AddObserver(name string, o Observer)
	RemoveObserver(name string)
	NotifyObservers()
}
