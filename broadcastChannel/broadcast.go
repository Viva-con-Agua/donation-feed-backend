package broadcastChannel

import "golang.org/x/exp/slices"

// BroadcastChannel is a wrapper around native channels that allows multiple receivers to dynamically subscribe and
// unsubscribe themselves to a single source channel.
// All of these receivers will then get a copy of the value from the source channel making the emitted value effectively
// being broadcast to all receivers.
type BroadcastChannel[T any] struct {
	source    chan T
	receivers []chan T
}

// NewBroadcastChannel creates a new BroadcastChannel that broadcasts message from the given source to any number
// of dynamically subscribed receivers.
func NewBroadcastChannel[T any](source chan T) *BroadcastChannel[T] {
	result := new(BroadcastChannel[T])
	result.receivers = make([]chan T, 0)
	result.source = source
	go result.runBroadcaster()
	return result
}

// Subscribe to events sent from this broadcast
func (c *BroadcastChannel[T]) Subscribe() chan T {
	receiver := make(chan T)
	c.receivers = append(c.receivers, receiver)
	return receiver
}

// Run a loop that relays messages from the BroadcastChannels source to all receivers
func (c *BroadcastChannel[T]) runBroadcaster() {
	// iterate over all data from the source
	for data := range c.source {
		// try to send data to all receivers
		closedChans := make([]chan T, 0)
		for _, receiver := range c.receivers {
			select {
			case receiver <- data:
			default:
				closedChans = append(closedChans, receiver)
			}
		}

		// remove all closed receivers from the broadcaster
		c.receivers = filterSlice(c.receivers, func(elem chan T) bool {
			return !slices.Contains(closedChans, elem)
		})
	}

	// close all receivers since source has now closed as well
	for _, receiver := range c.receivers {
		close(receiver)
	}
}
