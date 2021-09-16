package spdg

/*
Status enumerates the values that can be injected into a type to inform the
owner (caller) about the status of their request.
*/
type Status uint

const (
	NO  Status = iota // Rejection
	OK                // Accepted
	ERR               // Something went wrong
)

/*
Reason embelishes Status with an extra dimension of states that can be enumerated.
The whole thing eventually has the ability to come together as a powerful state-machine
even using techniques like auto-state mapping.
*/
type Reason uint

const (
	KEY       Reason = iota // Key was provided
	BUSY                    // Too busy for new workloads
	WORM                    // Write Once Read Many
	READONLY                // Can not write
	WRITEONLY               // Can not read
)
