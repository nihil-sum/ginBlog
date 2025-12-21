package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type JsonFormatTime struct {
	time.Time
}

func (t JsonFormatTime) MarshalJSON() ([]byte, error) {
	if (t == JsonFormatTime{}) {
		formatted := fmt.Sprintf("\"%s\"", "")
		return []byte(formatted), nil
	} else {
		formatted := fmt.Sprintf("\"%s\"", t.Format("2025-12-17 16:59:00"))
		return []byte(formatted), nil
	}
}
