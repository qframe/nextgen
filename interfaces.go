package nextgen


type Edge interface {

}


type Plugin interface {

}

// InputPlugin generates
type InputPlugin interface {
	Plugin
	New(Config, Sink) Input
}

// Input recevices data from an edge
type Input interface {


}

// Sink receives data
type Sink interface {
	Write(Message)
}

type Trans interface {
	Sink

}

type TransPlugin interface {
	Plugin
	New(Config, Sink) Trans
}

type Output interface {
	Sink
}

type OutputPlugin interface {
	Plugin
	New(Config) Output
}

type Cache interface {
	Sink
	Query(QueryRequest) QueryResponse
}

type QueryRequest interface {

}

type QueryResponse interface {

}

type Message interface {

}

type Config interface {

}

type SignalReceiver interface {
	ReceiveSignal()
}

type Signal struct {}