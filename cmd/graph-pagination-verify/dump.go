package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func dumpPayload(dumpFile io.Writer, key string, payload []byte) error {

	out := responseDump{
		Key:  key,
		Data: payload,
	}

	encoded, err := json.Marshal(out)
	if err != nil {
		return fmt.Errorf("could not encode payload: %w", err)
	}

	_, err = dumpFile.Write(encoded)
	if err != nil {
		return fmt.Errorf("could not dump payload: %w", err)
	}
	_, err = dumpFile.Write([]byte("\n"))
	if err != nil {
		return fmt.Errorf("could not write newline to dump file: %w", err)
	}

	return nil
}
