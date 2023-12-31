import (
    "fmt"
    "log/slog"
    "strings"
)

type Person struct {
	Name  string
	Email string
}

func (p Person) LogValue() slog.Value {
	parts := strings.Split(p.Email, "@")
	if len(parts) != 2 {
		return slog.GroupValue(
			slog.String("Name", p.Name),
			slog.String("Email", p.Email),
		)
	}

	localPart := parts[0]
	domain := parts[1]
	maskedLocalPart := strings.Repeat("*", len(localPart))
	return slog.GroupValue(
		slog.String("Name", p.Name),
		slog.String("Email", fmt.Sprintf("%s@%s", maskedLocalPart, domain)),
	)
}
